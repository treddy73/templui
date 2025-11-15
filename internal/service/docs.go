package service

import (
	"embed"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/templui/templui/internal/markdown"
	"github.com/templui/templui/internal/ui/modules"
)

//go:embed content/docs/*.md
var contentFS embed.FS

type DocsService struct {
	parser *markdown.Parser
}

type DocPage struct {
	Slug        string
	Title       string
	Description string
	Order       int
	Content     string
	TOC         []modules.TableOfContentsItem
}

func NewDocsService() *DocsService {
	return &DocsService{
		parser: markdown.NewParser(),
	}
}

// GetPage loads and parses a markdown document by slug
func (s *DocsService) GetPage(slug string) (*DocPage, error) {
	// Construct file path
	mdPath := filepath.Join("content/docs", slug+".md")

	// Read markdown file from embedded FS
	content, err := contentFS.ReadFile(mdPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read markdown file: %w", err)
	}

	// Parse with frontmatter
	html, meta, err := s.parser.ParseWithFrontmatter(content)
	if err != nil {
		return nil, fmt.Errorf("failed to parse markdown: %w", err)
	}

	// Extract table of contents
	tocItems := s.parser.ExtractTableOfContents(content)

	// Convert to modules.TableOfContentsItem
	moduleTOC := make([]modules.TableOfContentsItem, len(tocItems))
	for i, item := range tocItems {
		moduleTOC[i] = modules.TableOfContentsItem{
			ID:       item.ID,
			Text:     item.Text,
			Children: make([]modules.TableOfContentsItem, len(item.Children)),
		}
		for j, child := range item.Children {
			moduleTOC[i].Children[j] = modules.TableOfContentsItem{
				ID:   child.ID,
				Text: child.Text,
			}
		}
	}

	// Extract metadata
	page := &DocPage{
		Slug:    slug,
		Content: string(html),
		TOC:     moduleTOC,
	}

	if title, ok := meta["title"].(string); ok {
		page.Title = title
	}
	if desc, ok := meta["description"].(string); ok {
		page.Description = desc
	}
	if order, ok := meta["order"].(int); ok {
		page.Order = order
	}

	return page, nil
}

// ListPages returns all available doc pages
func (s *DocsService) ListPages() ([]DocPage, error) {
	var pages []DocPage

	err := fs.WalkDir(contentFS, "content/docs", func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() || !strings.HasSuffix(path, ".md") {
			return err
		}

		// Get relative path and slug
		relPath := strings.TrimPrefix(path, "content/docs/")
		slug := strings.TrimSuffix(relPath, ".md")

		// Read and extract frontmatter only
		content, err := contentFS.ReadFile(path)
		if err != nil {
			return err
		}

		meta := s.parser.ExtractFrontmatter(content)

		page := DocPage{Slug: slug}
		if title, ok := meta["title"].(string); ok {
			page.Title = title
		}
		if desc, ok := meta["description"].(string); ok {
			page.Description = desc
		}
		if order, ok := meta["order"].(int); ok {
			page.Order = order
		}

		pages = append(pages, page)
		return nil
	})

	return pages, err
}
