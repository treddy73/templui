package markdown

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/templui/templui/internal/ui/modules"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	goldmarkhtml "github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
	"go.abhg.dev/goldmark/frontmatter"
)

type Parser struct {
	md goldmark.Markdown
}

func NewParser() *Parser {
	// Custom code block renderer will get context per-request
	codeBlockRenderer := &ShikiCodeBlockRenderer{}

	md := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			extension.Footnote,
			extension.Typographer,
			&frontmatter.Extender{},
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			goldmarkhtml.WithHardWraps(),
			goldmarkhtml.WithXHTML(),
			// Add custom code block renderer
			renderer.WithNodeRenderers(
				util.Prioritized(codeBlockRenderer, 100),
			),
		),
	)

	return &Parser{
		md: md,
	}
}

func (p *Parser) Parse(source []byte) ([]byte, error) {
	var buf bytes.Buffer
	err := p.md.Convert(source, &buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (p *Parser) ParseWithFrontmatter(source []byte) (content []byte, meta map[string]any, err error) {
	context := parser.NewContext()
	var buf bytes.Buffer

	err = p.md.Convert(source, &buf, parser.WithContext(context))
	if err != nil {
		return nil, nil, err
	}

	data := frontmatter.Get(context)
	if data == nil {
		meta = make(map[string]any)
	} else {
		err = data.Decode(&meta)
		if err != nil {
			meta = make(map[string]any)
		}
	}

	return buf.Bytes(), meta, nil
}

// ShikiCodeBlockRenderer renders code blocks with Shiki highlighting
type ShikiCodeBlockRenderer struct{}

func (r *ShikiCodeBlockRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(ast.KindFencedCodeBlock, r.renderFencedCodeBlock)
	reg.Register(ast.KindCodeBlock, r.renderCodeBlock)
}

func (r *ShikiCodeBlockRenderer) renderFencedCodeBlock(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		return ast.WalkContinue, nil
	}

	n := node.(*ast.FencedCodeBlock)
	language := string(n.Language(source))
	if language == "" {
		language = "text"
	}

	// Extract code content
	var code bytes.Buffer
	lines := n.Lines()
	for i := 0; i < lines.Len(); i++ {
		line := lines.At(i)
		code.Write(line.Value(source))
	}

	// Get Shiki highlighted HTML with background context
	highlightedHTML := modules.GetHighlightedHTML(context.Background(), code.String(), language)

	// Wrap with same structure as modules.Code
	_, err := w.WriteString(fmt.Sprintf(`<div class="relative code-highlighting-container" data-code-block=""><div class="[&_pre]:block [&_pre]:overflow-x-auto [&_pre]:overflow-y-auto [&_pre]:max-h-96 [&_pre]:p-4 [&_pre]:rounded-md [&_pre]:text-sm">%s</div></div>`, highlightedHTML))
	return ast.WalkContinue, err
}

func (r *ShikiCodeBlockRenderer) renderCodeBlock(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		return ast.WalkContinue, nil
	}

	n := node.(*ast.CodeBlock)

	// Extract code content
	var code bytes.Buffer
	lines := n.Lines()
	for i := 0; i < lines.Len(); i++ {
		line := lines.At(i)
		code.Write(line.Value(source))
	}

	// Get Shiki highlighted HTML with default language
	highlightedHTML := modules.GetHighlightedHTML(context.Background(), code.String(), "text")

	// Wrap with same structure as modules.Code
	_, err := w.WriteString(fmt.Sprintf(`<div class="relative code-highlighting-container" data-code-block=""><div class="[&_pre]:block [&_pre]:overflow-x-auto [&_pre]:overflow-y-auto [&_pre]:max-h-96 [&_pre]:p-4 [&_pre]:rounded-md [&_pre]:text-sm">%s</div></div>`, highlightedHTML))
	return ast.WalkContinue, err
}

func (p *Parser) ConvertReader(r io.Reader, w io.Writer) error {
	data, err := readAll(r)
	if err != nil {
		return err
	}
	return p.md.Convert(data, w)
}

func readAll(r io.Reader) ([]byte, error) {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(r)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (p *Parser) ExtractFrontmatter(source []byte) map[string]any {
	context := parser.NewContext()
	p.md.Parser().Parse(text.NewReader(source), parser.WithContext(context))

	data := frontmatter.Get(context)
	if data == nil {
		return make(map[string]any)
	}

	var meta map[string]any
	err := data.Decode(&meta)
	if err != nil {
		return make(map[string]any)
	}
	return meta
}

// TOCItem represents a table of contents entry
type TOCItem struct {
	ID       string
	Text     string
	Level    int
	Children []TOCItem
}

// ExtractTableOfContents extracts headings from markdown to build a TOC
func (p *Parser) ExtractTableOfContents(source []byte) []TOCItem {
	doc := p.md.Parser().Parse(text.NewReader(source))
	var items []TOCItem

	ast.Walk(doc, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}

		if heading, ok := n.(*ast.Heading); ok {
			// Only extract h2 and h3 for TOC
			if heading.Level < 2 || heading.Level > 3 {
				return ast.WalkContinue, nil
			}

			// Extract heading text - walk all children to get full text
			var textBuf bytes.Buffer
			for child := heading.FirstChild(); child != nil; child = child.NextSibling() {
				if textNode, ok := child.(*ast.Text); ok {
					textBuf.Write(textNode.Segment.Value(source))
				}
			}
			text := textBuf.String()

			// Get heading ID - Goldmark's AutoHeadingID generates it
			id := ""
			if idAttr, ok := heading.Attribute([]byte("id")); ok {
				id = string(idAttr.([]byte))
			}

			// If no ID from attribute, generate one from text (same as Goldmark does)
			if id == "" && text != "" {
				id = generateHeadingID(text)
			}

			if text != "" && id != "" {
				items = append(items, TOCItem{
					ID:    id,
					Text:  text,
					Level: heading.Level,
				})
			}
		}

		return ast.WalkContinue, nil
	})

	// Build hierarchical structure (h2 as parent, h3 as children)
	return buildTOCHierarchy(items)
}

// generateHeadingID creates an ID from heading text (matching Goldmark's behavior)
func generateHeadingID(text string) string {
	// Convert to lowercase
	id := ""
	for _, r := range text {
		if r >= 'A' && r <= 'Z' {
			id += string(r + 32) // Convert to lowercase
		} else if r >= 'a' && r <= 'z' || r >= '0' && r <= '9' {
			id += string(r)
		} else if r == ' ' || r == '-' {
			id += "-"
		}
	}

	// Remove consecutive dashes and trim
	result := ""
	lastDash := false
	for _, r := range id {
		if r == '-' {
			if !lastDash {
				result += string(r)
				lastDash = true
			}
		} else {
			result += string(r)
			lastDash = false
		}
	}

	// Trim dashes from start and end
	for len(result) > 0 && result[0] == '-' {
		result = result[1:]
	}
	for len(result) > 0 && result[len(result)-1] == '-' {
		result = result[:len(result)-1]
	}

	return result
}

func buildTOCHierarchy(items []TOCItem) []TOCItem {
	if len(items) == 0 {
		return items
	}

	var result []TOCItem
	var currentH2 *TOCItem

	for _, item := range items {
		if item.Level == 2 {
			// New h2 - add previous h2 if exists, start new one
			if currentH2 != nil {
				result = append(result, *currentH2)
			}
			currentH2 = &TOCItem{
				ID:       item.ID,
				Text:     item.Text,
				Level:    item.Level,
				Children: []TOCItem{},
			}
		} else if item.Level == 3 && currentH2 != nil {
			// h3 - add as child to current h2
			currentH2.Children = append(currentH2.Children, item)
		}
	}

	// Don't forget the last h2
	if currentH2 != nil {
		result = append(result, *currentH2)
	}

	return result
}
