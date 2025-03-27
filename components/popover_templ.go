// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.856
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"github.com/axzilla/templui/utils"
	"strconv"
)

type PopoverPosition string

const (
	PopoverTop         PopoverPosition = "top"
	PopoverTopStart    PopoverPosition = "top-start"
	PopoverTopEnd      PopoverPosition = "top-end"
	PopoverRight       PopoverPosition = "right"
	PopoverRightStart  PopoverPosition = "right-start"
	PopoverRightEnd    PopoverPosition = "right-end"
	PopoverBottom      PopoverPosition = "bottom"
	PopoverBottomStart PopoverPosition = "bottom-start"
	PopoverBottomEnd   PopoverPosition = "bottom-end"
	PopoverLeft        PopoverPosition = "left"
	PopoverLeftStart   PopoverPosition = "left-start"
	PopoverLeftEnd     PopoverPosition = "left-end"
)

type PopoverTriggerType string

const (
	PopoverTriggerTypeHover PopoverTriggerType = "hover"
	PopoverTriggerTypeClick PopoverTriggerType = "click"
)

type PopoverProps struct {
	Class string
}

type PopoverTriggerProps struct {
	ID          string
	TriggerType PopoverTriggerType
}

type PopoverContentProps struct {
	ID               string
	Class            string
	Attributes       templ.Attributes
	Position         PopoverPosition
	DisableClickAway bool
	DisableESC       bool
	ShowArrow        bool
}

func popoverPortalContainer() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<div id=\"popover-portal-container\" class=\"fixed inset-0 z-[9999] pointer-events-none\"></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func PopoverTrigger(props ...PopoverTriggerProps) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var p PopoverTriggerProps
		if len(props) > 0 {
			p = props[0]
		}
		if p.TriggerType == "" {
			p.TriggerType = PopoverTriggerTypeClick
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "<span data-popover-trigger data-popover-id=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(p.ID)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/popover.templ`, Line: 68, Col: 24}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "\" data-popover-type=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(string(p.TriggerType))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/popover.templ`, Line: 69, Col: 43}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var2.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "</span>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func Popover(props ...PopoverProps) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var5 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var5 == nil {
			templ_7745c5c3_Var5 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var p PopoverProps
		if len(props) > 0 {
			p = props[0]
		}
		var templ_7745c5c3_Var6 = []any{utils.TwMerge("relative inline-block", p.Class)}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var6...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "<div class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var7 string
		templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var6).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/popover.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 7, "\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var5.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 8, "</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = popoverPortalContainer().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func PopoverContent(props ...PopoverContentProps) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var8 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var8 == nil {
			templ_7745c5c3_Var8 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var p PopoverContentProps
		if len(props) > 0 {
			p = props[0]
		}
		if p.Position == "" {
			p.Position = PopoverBottom
		}
		var templ_7745c5c3_Var9 = []any{
			utils.TwMerge(
				"bg-background rounded-lg border text-sm shadow-lg hidden pointer-events-auto absolute z-[9999]",
				p.Class,
			),
		}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var9...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 9, "<div data-popover-content data-popover-id=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var10 string
		templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(p.ID)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/popover.templ`, Line: 96, Col: 24}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 10, "\" data-popover-position=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var11 string
		templ_7745c5c3_Var11, templ_7745c5c3_Err = templ.JoinStringErrs(string(p.Position))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/popover.templ`, Line: 97, Col: 44}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var11))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 11, "\" data-popover-disable-clickaway=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var12 string
		templ_7745c5c3_Var12, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.FormatBool(p.DisableClickAway))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/popover.templ`, Line: 98, Col: 73}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var12))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 12, "\" data-popover-disable-esc=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var13 string
		templ_7745c5c3_Var13, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.FormatBool(p.DisableESC))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/popover.templ`, Line: 99, Col: 61}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var13))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 13, "\" class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var14 string
		templ_7745c5c3_Var14, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var9).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/popover.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var14))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 14, "\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.RenderAttributes(ctx, templ_7745c5c3_Buffer, p.Attributes)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 15, "><div class=\"w-full overflow-hidden\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var8.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 16, "</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if p.ShowArrow {
			var templ_7745c5c3_Var15 = []any{
				utils.TwMerge(
					"absolute h-2.5 w-2.5 rotate-45 bg-background",
					popoverArrowClass(p.Position),
				),
			}
			templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var15...)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 17, "<div class=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var16 string
			templ_7745c5c3_Var16, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var15).String())
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/popover.templ`, Line: 1, Col: 0}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var16))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 18, "\"></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 19, "</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func popoverArrowClass(side PopoverPosition) string {
	switch side {
	case PopoverTop:
		return "bottom-[-5px] left-1/2 -translate-x-1/2 border-b border-r"
	case PopoverTopStart:
		return "bottom-[-5px] left-4 border-b border-r"
	case PopoverTopEnd:
		return "bottom-[-5px] right-4 border-b border-r"
	case PopoverRight:
		return "left-[-5px] top-1/2 -translate-y-1/2 border-b border-l"
	case PopoverRightStart:
		return "left-[-5px] top-2 border-b border-l"
	case PopoverRightEnd:
		return "left-[-5px] bottom-2 border-b border-l"
	case PopoverBottom:
		return "top-[-5px] left-1/2 -translate-x-1/2 border-t border-l"
	case PopoverBottomStart:
		return "top-[-5px] left-4 border-t border-l"
	case PopoverBottomEnd:
		return "top-[-5px] right-4 border-t border-l"
	case PopoverLeft:
		return "right-[-5px] top-1/2 -translate-y-1/2 border-t border-r"
	case PopoverLeftStart:
		return "right-[-5px] top-2 border-t border-r"
	case PopoverLeftEnd:
		return "right-[-5px] bottom-2 border-t border-r"
	default:
		return "top-[-5px] left-1/2 -translate-x-1/2 border-t border-l"
	}
}

func PopoverScript() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var17 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var17 == nil {
			templ_7745c5c3_Var17 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 20, "<script defer nonce=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var18 string
		templ_7745c5c3_Var18, templ_7745c5c3_Err = templ.JoinStringErrs(templ.GetNonce(ctx))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/popover.templ`, Line: 156, Col: 42}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var18))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 21, "\">\n        document.addEventListener('DOMContentLoaded', () => {\n            const portalContainer = document.getElementById('popover-portal-container');\n            if (!portalContainer) return;\n            \n            const triggers = document.querySelectorAll('[data-popover-trigger]');\n            const contents = document.querySelectorAll('[data-popover-content]');\n\n            contents.forEach(content => {\n                portalContainer.appendChild(content);\n            });\n\n            const positionPopover = (trigger, content) => {\n                // Wir müssen das tatsächliche Element finden, auf das sich der Popover bezieht\n                // Das kann der Trigger selbst sein oder ein Kind-Element\n                // Wir analysieren die Boundingbox und nehmen das größte Element\n                let triggerElement = trigger;\n                let largestArea = 0;\n                \n                // Prüfe alle direkten Kinder des Triggers\n                const children = trigger.children;\n                for (let i = 0; i < children.length; i++) {\n                    const child = children[i];\n                    const rect = child.getBoundingClientRect();\n                    const area = rect.width * rect.height;\n                    \n                    // Wenn das Kind eine größere Fläche hat, ist es wahrscheinlich das sichtbare Element\n                    if (area > largestArea) {\n                        largestArea = area;\n                        triggerElement = child;\n                    }\n                }\n                \n                const triggerRect = triggerElement.getBoundingClientRect();\n                const contentRect = content.getBoundingClientRect();\n                const margin = 8;\n                const scrollY = window.scrollY || window.pageYOffset;\n                const scrollX = window.scrollX || window.pageXOffset;\n\n                let top, left;\n                const position = content.dataset.popoverPosition || 'bottom';\n\n                // Get element heights\n                const triggerHeight = triggerRect.height;\n                const contentHeight = contentRect.height;\n                \n                // Definiere Ankerpunkte\n                const triggerTop = triggerRect.top + scrollY;\n                const triggerBottom = triggerRect.bottom + scrollY;\n                const triggerLeft = triggerRect.left + scrollX;\n                const triggerRight = triggerRect.right + scrollX;\n                \n                console.log(`Selected Trigger Element:`, triggerElement);\n                console.log(`Trigger: top=${triggerTop}, bottom=${triggerBottom}, height=${triggerHeight}`);\n                console.log(`Content: height=${contentHeight}`);\n\n                switch (position) {\n                    case 'top':\n                        top = triggerTop - contentHeight - margin;\n                        left = triggerLeft + (triggerRect.width / 2) - (contentRect.width / 2);\n                        break;\n                    case 'top-start':\n                        top = triggerTop - contentHeight - margin;\n                        left = triggerLeft;\n                        break;\n                    case 'top-end':\n                        top = triggerTop - contentHeight - margin;\n                        left = triggerRight - contentRect.width;\n                        break;\n                    case 'right':\n                        top = triggerTop + (triggerHeight / 2) - (contentHeight / 2);\n                        left = triggerRight + margin;\n                        break;\n                    case 'right-start':\n                        // Rechts vom Trigger, oben ausgerichtet\n                        top = triggerTop;\n                        left = triggerRight + margin;\n                        break;\n                    case 'right-end':\n                        // Rechts vom Trigger, unten ausgerichtet\n                        top = triggerBottom - contentHeight;\n                        left = triggerRight + margin;\n                        break;\n                    case 'bottom':\n                        top = triggerBottom + margin;\n                        left = triggerLeft + (triggerRect.width / 2) - (contentRect.width / 2);\n                        break;\n                    case 'bottom-start':\n                        top = triggerBottom + margin;\n                        left = triggerLeft;\n                        break;\n                    case 'bottom-end':\n                        top = triggerBottom + margin;\n                        left = triggerRight - contentRect.width;\n                        break;\n                    case 'left':\n                        top = triggerTop + (triggerHeight / 2) - (contentHeight / 2);\n                        left = triggerLeft - contentRect.width - margin;\n                        break;\n                    case 'left-start':\n                        // Links vom Trigger, oben ausgerichtet\n                        top = triggerTop;\n                        left = triggerLeft - contentRect.width - margin;\n                        break;\n                    case 'left-end':\n                        // Links vom Trigger, unten ausgerichtet\n                        top = triggerBottom - contentHeight;\n                        left = triggerLeft - contentRect.width - margin;\n                        break;\n                    default:\n                        top = triggerBottom + margin;\n                        left = triggerLeft + (triggerRect.width / 2) - (contentRect.width / 2);\n                }\n\n                content.style.top = `${top}px`;\n                content.style.left = `${left}px`;\n                \n                console.log(`Positioning ${position}: top=${top}, left=${left}`);\n            };\n\n            triggers.forEach(trigger => {\n                const popoverId = trigger.dataset.popoverId;\n                const content = document.querySelector(`[data-popover-content][data-popover-id=\"${popoverId}\"]`);\n                if (!content) return;\n\n                const triggerType = trigger.dataset.popoverType;\n\n                if (triggerType === 'click') {\n                    trigger.addEventListener('click', () => {\n                        const isOpen = content.classList.contains('hidden');\n                        content.classList.toggle('hidden', !isOpen);\n                        if (isOpen) positionPopover(trigger, content);\n                    });\n\n                    if (content.dataset.popoverDisableClickaway !== 'true') {\n                        document.addEventListener('click', (e) => {\n                            if (!content.classList.contains('hidden') && \n                                !trigger.contains(e.target) && \n                                !content.contains(e.target)) {\n                                content.classList.add('hidden');\n                            }\n                        });\n                    }\n\n                    if (content.dataset.popoverDisableEsc !== 'true') {\n                        document.addEventListener('keydown', (e) => {\n                            if (e.key === 'Escape' && !content.classList.contains('hidden')) {\n                                content.classList.add('hidden');\n                            }\n                        });\n                    }\n                } else if (triggerType === 'hover') {\n                    trigger.addEventListener('mouseenter', () => {\n                        content.classList.remove('hidden');\n                        positionPopover(trigger, content);\n                    });\n                    trigger.addEventListener('mouseleave', () => {\n                        content.classList.add('hidden');\n                    });\n                }\n            });\n\n            // Global scroll and resize event for all popovers\n            window.addEventListener('scroll', () => {\n                triggers.forEach(trigger => {\n                    const popoverId = trigger.dataset.popoverId;\n                    const content = document.querySelector(`[data-popover-content][data-popover-id=\"${popoverId}\"]`);\n                    if (!content || content.classList.contains('hidden')) return;\n                    \n                    positionPopover(trigger, content);\n                });\n            }, { passive: true }); // Using passive for better scroll performance\n\n            // Identify and attach scroll events to all scrollable parents\n            triggers.forEach(trigger => {\n                let element = trigger;\n                const scrollableParents = [];\n                \n                // Find all scrollable parents\n                while (element) {\n                    if (element.scrollHeight > element.clientHeight) {\n                        scrollableParents.push(element);\n                    }\n                    element = element.parentElement;\n                }\n                \n                // Attach scroll event to each scrollable parent\n                scrollableParents.forEach(scrollParent => {\n                    scrollParent.addEventListener('scroll', () => {\n                        const popoverId = trigger.dataset.popoverId;\n                        const content = document.querySelector(`[data-popover-content][data-popover-id=\"${popoverId}\"]`);\n                        if (!content || content.classList.contains('hidden')) return;\n                        \n                        positionPopover(trigger, content);\n                    }, { passive: true });\n                });\n            });\n\n            window.addEventListener('resize', () => {\n                triggers.forEach(trigger => {\n                    const popoverId = trigger.dataset.popoverId;\n                    const content = document.querySelector(`[data-popover-content][data-popover-id=\"${popoverId}\"]`);\n                    if (!content || content.classList.contains('hidden')) return;\n                    \n                    positionPopover(trigger, content);\n                });\n            });\n        });\n    </script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
