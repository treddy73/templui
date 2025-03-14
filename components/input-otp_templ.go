// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.833
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"github.com/axzilla/templui/utils"
	"strconv"
)

type InputOTPProps struct {
	ID          string
	Name        string
	Value       string
	Length      int
	Type        string
	Placeholder string
	Disabled    bool
	Required    bool
	HasError    bool
	Class       string
	Attributes  templ.Attributes
}

func InputOTPScript() templ.Component {
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
		handle := templ.NewOnceHandle()
		templ_7745c5c3_Var2 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
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
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<script defer nonce=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(templ.GetNonce(ctx))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/input-otp.templ`, Line: 25, Col: 43}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "\">\n\t\t\tdocument.addEventListener('alpine:init', () => {\n\t\t\t\tAlpine.data('inputOTP', () => ({\n\t\t\t\t\totpValue: '',\n\t\t\t\t\tinputRefs: [],\n\t\t\t\t\t\n\t\t\t\t\tinit() {\n\t\t\t\t\t\tif (this.$el.dataset.value) {\n\t\t\t\t\t\t\tthis.otpValue = this.$el.dataset.value;\n\t\t\t\t\t\t\tthis.fillInputsFromValue();\n\t\t\t\t\t\t}\n\t\t\t\t\t},\n\t\t\t\t\t\n\t\t\t\t\tgetInputRefs() {\n\t\t\t\t\t\tconst length = parseInt(this.$el.dataset.length || '6');\n\t\t\t\t\t\tthis.inputRefs = [];\n\t\t\t\t\t\tfor (let i = 0; i < length; i++) {\n\t\t\t\t\t\t\tconst ref = this.$refs[`otpDigit${i}`];\n\t\t\t\t\t\t\tif (ref) this.inputRefs.push(ref);\n\t\t\t\t\t\t}\n\t\t\t\t\t\treturn this.inputRefs;\n\t\t\t\t\t},\n\t\t\t\t\t\n\t\t\t\t\tfocusFirstEmptyInput() {\n\t\t\t\t\t\tconst inputs = this.getInputRefs();\n\t\t\t\t\t\tfor (let i = 0; i < inputs.length; i++) {\n\t\t\t\t\t\t\tif (!inputs[i].value) {\n\t\t\t\t\t\t\t\tinputs[i].focus();\n\t\t\t\t\t\t\t\tbreak;\n\t\t\t\t\t\t\t}\n\t\t\t\t\t\t}\n\t\t\t\t\t},\n\t\t\t\t\t\n\t\t\t\t\tfillInputsFromValue() {\n\t\t\t\t\t\tconst inputs = this.getInputRefs();\n\t\t\t\t\t\tconst valueStr = this.otpValue.toString();\n\t\t\t\t\t\t\n\t\t\t\t\t\tfor (let i = 0; i < inputs.length; i++) {\n\t\t\t\t\t\t\tif (i < valueStr.length) {\n\t\t\t\t\t\t\t\tinputs[i].value = valueStr[i];\n\t\t\t\t\t\t\t} else {\n\t\t\t\t\t\t\t\tinputs[i].value = '';\n\t\t\t\t\t\t\t}\n\t\t\t\t\t\t}\n\t\t\t\t\t},\n\t\t\t\t\t\n\t\t\t\t\tupdateValue() {\n\t\t\t\t\t\tconst inputs = this.getInputRefs();\n\t\t\t\t\t\tlet value = '';\n\t\t\t\t\t\t\n\t\t\t\t\t\tfor (let i = 0; i < inputs.length; i++) {\n\t\t\t\t\t\t\tvalue += inputs[i].value || '';\n\t\t\t\t\t\t}\n\t\t\t\t\t\t\n\t\t\t\t\t\tthis.otpValue = value;\n\t\t\t\t\t},\n\t\t\t\t\t\n\t\t\t\t\thandleInput(event) {\n\t\t\t\t\t\tconst input = event.target;\n\t\t\t\t\t\tconst index = parseInt(input.dataset.inputIndex || 0);\n\t\t\t\t\t\tconst inputs = this.getInputRefs();\n\t\t\t\t\t\t\n\t\t\t\t\t\tif (input.value.length > 1) {\n\t\t\t\t\t\t\tinput.value = input.value.slice(-1);\n\t\t\t\t\t\t}\n\t\t\t\t\t\t\n\t\t\t\t\t\tif (input.value && index < inputs.length - 1) {\n\t\t\t\t\t\t\tinputs[index + 1].focus();\n\t\t\t\t\t\t}\n\t\t\t\t\t\t\n\t\t\t\t\t\tthis.updateValue();\n\t\t\t\t\t},\n\t\t\t\t\t\n\t\t\t\t\thandleKeydown(event) {\n\t\t\t\t\t\tconst input = event.target;\n\t\t\t\t\t\tconst index = parseInt(input.dataset.inputIndex || 0);\n\t\t\t\t\t\tconst inputs = this.getInputRefs();\n\t\t\t\t\t\t\n\t\t\t\t\t\tif (event.key === 'Backspace') {\n\t\t\t\t\t\t\tconst input = event.target;\n\t\t\t\t\t\t\t\n\t\t\t\t\t\t\tif (!input.value && index > 0) {\n\t\t\t\t\t\t\t\tinputs[index - 1].focus();\n\t\t\t\t\t\t\t\tinputs[index - 1].value = '';\n\t\t\t\t\t\t\t\tevent.preventDefault();\n\t\t\t\t\t\t\t}\n\t\t\t\t\t\t}\n\t\t\t\t\t\t\n\t\t\t\t\t\tif (event.key === 'ArrowLeft' && index > 0) {\n\t\t\t\t\t\t\tinputs[index - 1].focus();\n\t\t\t\t\t\t\tevent.preventDefault();\n\t\t\t\t\t\t}\n\t\t\t\t\t\t\n\t\t\t\t\t\tif (event.key === 'ArrowRight' && index < inputs.length - 1) {\n\t\t\t\t\t\t\tinputs[index + 1].focus();\n\t\t\t\t\t\t\tevent.preventDefault();\n\t\t\t\t\t\t}\n\t\t\t\t\t\t\n\t\t\t\t\t\tthis.updateValue();\n\t\t\t\t\t},\n\t\t\t\t\t\n\t\t\t\t\thandlePaste(event) {\n\t\t\t\t\t\tevent.preventDefault();\n\t\t\t\t\t\tconst inputs = this.getInputRefs();\n\t\t\t\t\t\t\n\t\t\t\t\t\tconst pastedData = (event.clipboardData || window.clipboardData).getData('text');\n\t\t\t\t\t\tconst pastedChars = pastedData.trim().split('');\n\t\t\t\t\t\t\n\t\t\t\t\t\tfor (let i = 0; i < inputs.length && i < pastedChars.length; i++) {\n\t\t\t\t\t\t\tinputs[i].value = pastedChars[i];\n\t\t\t\t\t\t}\n\t\t\t\t\t\t\n\t\t\t\t\t\tthis.focusFirstEmptyInput();\n\t\t\t\t\t\tthis.updateValue();\n\t\t\t\t\t}\n\t\t\t\t}));\n\t\t\t});\n\t\t</script>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return nil
		})
		templ_7745c5c3_Err = handle.Once().Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func InputOTP(props InputOTPProps) templ.Component {
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
		templ_7745c5c3_Var4 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var4 == nil {
			templ_7745c5c3_Var4 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		if props.ID == "" {
			props.ID = utils.RandomID()
		}
		if props.Length <= 0 {
			props.Length = 6
		}
		if props.Type == "" {
			props.Type = "text"
		}
		if props.Placeholder == "" {
			props.Placeholder = ""
		}
		templ_7745c5c3_Err = InputOTPScript().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 = []any{utils.TwMerge(
			"flex flex-row gap-2 w-fit",
			props.Class,
		)}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var5...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "<div id=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var6 string
		templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(props.ID + "-container")
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/input-otp.templ`, Line: 161, Col: 30}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "\" class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var7 string
		templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var5).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/input-otp.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "\" x-data=\"inputOTP\" data-length=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var8 string
		templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(props.Length))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/input-otp.templ`, Line: 167, Col: 42}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if props.Value != "" {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 7, " data-value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var9 string
			templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(props.Value)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/input-otp.templ`, Line: 169, Col: 27}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 8, "\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 9, " x-on:paste.prevent=\"handlePaste\"><input type=\"hidden\" id=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var10 string
		templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(props.ID)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/input-otp.templ`, Line: 175, Col: 16}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 10, "\" name=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var11 string
		templ_7745c5c3_Var11, templ_7745c5c3_Err = templ.JoinStringErrs(props.Name)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/input-otp.templ`, Line: 176, Col: 20}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var11))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 11, "\" x-model=\"otpValue\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if props.Required {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 12, " required")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 13, "> ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for i := 0; i < props.Length; i++ {
			var templ_7745c5c3_Var12 = []any{utils.TwMerge(
				"w-10 h-12 text-center",
				"rounded-md border border-input bg-background text-sm",
				"file:border-0 file:bg-transparent file:text-sm file:font-medium",
				"placeholder:text-muted-foreground",
				"focus-visible:outline-hidden focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2",
				"disabled:cursor-not-allowed disabled:opacity-50",
				utils.TwIf("border-destructive ring-destructive", props.HasError),
			)}
			templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var12...)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 14, "<input x-ref=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var13 string
			templ_7745c5c3_Var13, templ_7745c5c3_Err = templ.JoinStringErrs("otpDigit" + strconv.Itoa(i))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/input-otp.templ`, Line: 182, Col: 40}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var13))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 15, "\" type=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var14 string
			templ_7745c5c3_Var14, templ_7745c5c3_Err = templ.JoinStringErrs(props.Type)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/input-otp.templ`, Line: 183, Col: 21}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var14))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 16, "\" inputmode=\"numeric\" placeholder=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var15 string
			templ_7745c5c3_Var15, templ_7745c5c3_Err = templ.JoinStringErrs(props.Placeholder)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/input-otp.templ`, Line: 185, Col: 35}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var15))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 17, "\" maxlength=\"1\" class=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var16 string
			templ_7745c5c3_Var16, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var12).String())
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/input-otp.templ`, Line: 1, Col: 0}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var16))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 18, "\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if props.Disabled {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 19, " disabled")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 20, " data-input-index=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var17 string
			templ_7745c5c3_Var17, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(i))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/input-otp.templ`, Line: 197, Col: 38}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var17))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 21, "\" x-on:input=\"handleInput\" x-on:keydown=\"handleKeydown\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 22, "</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
