// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.833
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"fmt"
	"github.com/axzilla/templui/icons"
	"github.com/axzilla/templui/utils"
	"strconv"
	"time"
)

type TimePickerProps struct {
	ID          string
	Name        string
	Value       time.Time
	Use12Hours  bool
	AMLabel     string
	PMLabel     string
	Placeholder string
	Required    bool
	Disabled    bool
	HasError    bool
	Class       string
	Attributes  templ.Attributes
}

type SelectOption struct {
	Label string
	Value string
}

func generateHourOptions(use12Hours bool) []SelectOption {
	options := make([]SelectOption, 0)
	if use12Hours {
		options = append(options, SelectOption{
			Label: "12",
			Value: "12",
		})
		for i := 1; i <= 11; i++ {
			options = append(options, SelectOption{
				Label: fmt.Sprintf("%02d", i),
				Value: fmt.Sprintf("%d", i),
			})
		}
	} else {
		for i := 0; i <= 23; i++ {
			options = append(options, SelectOption{
				Label: fmt.Sprintf("%02d", i),
				Value: fmt.Sprintf("%d", i),
			})
		}
	}
	return options
}

func generateMinuteOptions() []SelectOption {
	options := make([]SelectOption, 60)
	for i := 0; i < 60; i++ {
		options[i] = SelectOption{
			Label: fmt.Sprintf("%02d", i),
			Value: fmt.Sprintf("%d", i),
		}
	}
	return options
}

func TimePickerScript() templ.Component {
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
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/time_picker.templ`, Line: 69, Col: 43}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "\">\n\t\t\tdocument.addEventListener('alpine:init', () => {\n\t\t\t\tAlpine.data('timePicker', () => ({\n\t\t\t\t\topen: false,\n\t\t\t\t\tformValue: null,\n\t\t\t\t\tdisplayValue: null,\n\t\t\t\t\tselectedHour: 0,\n\t\t\t\t\tselectedMinute: 0,\n\t\t\t\t\tisPM: false,\n\t\t\t\t\thours: [],\n\t\t\t\t\tminutes: [],\n\t\t\t\t\tuse12Hours: false,\n\t\t\t\t\tposition: 'bottom',\n\n\t\t\t\t\tinit() {\n\t\t\t\t\t\tthis.use12Hours = this.$el.dataset.use12hours === 'true';\n\t\t\t\t\t\tif (this.$el.dataset?.value) {\n\t\t\t\t\t\t\tconst [hours, minutes] = this.$el.dataset.value.split(':').map(Number);\n\t\t\t\t\t\t\tthis.selectedMinute = minutes;\n\t\t\t\t\t\t\t\n\t\t\t\t\t\t\tif (this.use12Hours) {\n\t\t\t\t\t\t\t\tthis.isPM = hours >= 12;\n\t\t\t\t\t\t\t\tthis.selectedHour = hours > 12 ? hours - 12 : (hours === 0 ? 12 : hours);\n\t\t\t\t\t\t\t} else {\n\t\t\t\t\t\t\t\tthis.selectedHour = hours;\n\t\t\t\t\t\t\t}\n\t\t\t\t\t\t\tthis.updateValues();\n\t\t\t\t\t\t}\n\t\t\t\t\t\t\n\t\t\t\t\t\t// Set up listeners for the selects\n\t\t\t\t\t\tthis.$nextTick(() => {\n                            // We handle selection by listening to click events on the select items\n                            // This is handled by the SelectScript but we need to sync the values\n                            \n                            const hourSelectId = 'hour-select-' + this.$el.dataset.inputId;\n                            const minuteSelectId = 'minute-select-' + this.$el.dataset.inputId;\n                            \n                            // Listen for changes on the hidden inputs to sync our values\n                            const hourInput = document.querySelector(`[data-select-id=\"${hourSelectId}\"] input[type=\"hidden\"]`);\n                            const minuteInput = document.querySelector(`[data-select-id=\"${minuteSelectId}\"] input[type=\"hidden\"]`);\n                            \n                            if (hourInput) {\n                                hourInput.addEventListener('change', () => {\n                                    this.selectedHour = parseInt(hourInput.value);\n                                    this.updateValues();\n                                });\n                            }\n                            \n                            if (minuteInput) {\n                                minuteInput.addEventListener('change', () => {\n                                    this.selectedMinute = parseInt(minuteInput.value);\n                                    this.updateValues();\n                                });\n                            }\n                        });\n\t\t\t\t\t},\n\n\t\t\t\t\ttoggleTimePicker() {\n\t\t\t\t\t\tthis.open = !this.open;\n\t\t\t\t\t\tif (this.open) {\n\t\t\t\t\t\t\tthis.$nextTick(() => this.updatePosition());\n\t\t\t\t\t\t}\n\t\t\t\t\t},\n\n\t\t\t\t\tupdatePosition() {\n\t\t\t\t\t\tconst inputId = this.$root.dataset.inputId;\n\t\t\t\t\t\tconst trigger = document.getElementById(inputId);\n\t\t\t\t\t\tconst popup = this.$refs.timePickerPopup;\n\t\t\t\t\t\t\n\t\t\t\t\t\tif (!trigger || !popup) return;\n\t\t\t\t\t\t\n\t\t\t\t\t\tconst rect = trigger.getBoundingClientRect();\n\t\t\t\t\t\tconst popupRect = popup.getBoundingClientRect();\n\t\t\t\t\t\tconst viewportHeight = window.innerHeight;\n\t\t\t\t\t\t\n\t\t\t\t\t\tif (rect.bottom + popupRect.height > viewportHeight && rect.top > popupRect.height) {\n\t\t\t\t\t\t\tthis.position = 'top';\n\t\t\t\t\t\t} else {\n\t\t\t\t\t\t\tthis.position = 'bottom';\n\t\t\t\t\t\t}\n\t\t\t\t\t},\n\n\t\t\t\t\tpositionClass() {\n\t\t\t\t\t\treturn this.position === 'bottom' ? 'top-full mt-1' : 'bottom-full mb-1';\n\t\t\t\t\t},\n\n\t\t\t\t\tcloseTimePicker() {\n\t\t\t\t\t\tthis.open = false;\n\t\t\t\t\t},\n\n\t\t\t\t\tupdateValues() {\n\t\t\t\t\t\tlet hour24 = this.selectedHour;\n\t\t\t\t\t\tif (this.use12Hours) {\n\t\t\t\t\t\t\tif (this.isPM && hour24 !== 12) hour24 += 12;\n\t\t\t\t\t\t\tif (!this.isPM && hour24 === 12) hour24 = 0;\n\t\t\t\t\t\t}\n\t\t\t\t\t\tthis.formValue = `${String(hour24).padStart(2, '0')}:${String(this.selectedMinute).padStart(2, '0')}`;\n\t\t\t\t\t\tif (this.use12Hours) {\n\t\t\t\t\t\t\tthis.displayValue = `${String(this.selectedHour).padStart(2, '0')}:${String(this.selectedMinute).padStart(2, '0')} ${this.isPM ? 'PM' : 'AM'}`;\n\t\t\t\t\t\t} else {\n\t\t\t\t\t\t\tthis.displayValue = this.formValue;\n\t\t\t\t\t\t}\n                        \n                        // Update the display element with the new value\n                        const displayEl = this.$el.querySelector('.timepicker-display');\n                        if (displayEl) {\n                            displayEl.textContent = this.displayValue;\n                            displayEl.classList.remove('text-muted-foreground');\n                        }\n\t\t\t\t\t},\n\n\t\t\t\t\tselectPeriod() {\n\t\t\t\t\t\tconst period = this.$el.getAttribute('data-period');\n\t\t\t\t\t\tthis.isPM = period === 'PM';\n\t\t\t\t\t\tthis.updateValues();\n\t\t\t\t\t}\n\t\t\t\t}));\n\t\t\t});\n\t\t</script>")
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

func TimePicker(props TimePickerProps) templ.Component {
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
		if props.Placeholder == "" {
			props.Placeholder = "Select time"
		}
		if props.AMLabel == "" {
			props.AMLabel = "AM"
		}
		if props.PMLabel == "" {
			props.PMLabel = "PM"
		}
		var templ_7745c5c3_Var5 = []any{utils.TwMerge("relative", props.Class)}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var5...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "<div class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var6 string
		templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var5).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/time_picker.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if props.Value != (time.Time{}) {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, " data-value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var7 string
			templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(props.Value.Format("15:04"))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/time_picker.templ`, Line: 207, Col: 43}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 7, " data-use12hours=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var8 string
		templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("%t", props.Use12Hours))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/time_picker.templ`, Line: 209, Col: 55}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 8, "\" x-data=\"timePicker\" data-input-id=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var9 string
		templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(props.ID)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/time_picker.templ`, Line: 211, Col: 26}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 9, "\"><div class=\"relative\"><input type=\"hidden\" name=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var10 string
		templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(props.Name)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/time_picker.templ`, Line: 216, Col: 21}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 10, "\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if props.Required {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 11, " required")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 12, ">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var11 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
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
			var templ_7745c5c3_Var12 = []any{utils.TwMerge(
				"timepicker-display truncate",
				utils.TwIf("text-muted-foreground", props.Value == (time.Time{})),
			)}
			templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var12...)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 13, "<span class=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var13 string
			templ_7745c5c3_Var13, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var12).String())
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/time_picker.templ`, Line: 1, Col: 0}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var13))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 14, "\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if props.Value != (time.Time{}) {
				if props.Use12Hours {
					var templ_7745c5c3_Var14 string
					templ_7745c5c3_Var14, templ_7745c5c3_Err = templ.JoinStringErrs(props.Value.Format("03:04 PM"))
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/time_picker.templ`, Line: 243, Col: 39}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var14))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				} else {
					var templ_7745c5c3_Var15 string
					templ_7745c5c3_Var15, templ_7745c5c3_Err = templ.JoinStringErrs(props.Value.Format("15:04"))
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/time_picker.templ`, Line: 245, Col: 36}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var15))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				}
			} else {
				var templ_7745c5c3_Var16 string
				templ_7745c5c3_Var16, templ_7745c5c3_Err = templ.JoinStringErrs(props.Placeholder)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/time_picker.templ`, Line: 248, Col: 25}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var16))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 15, "</span>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = icons.Clock(icons.IconProps{
				Class: "ml-2 h-4 w-4 text-muted-foreground",
			}).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return nil
		})
		templ_7745c5c3_Err = Button(ButtonProps{
			Type:    "button",
			Variant: ButtonVariantOutline,
			Class: utils.TwMerge(
				"w-full flex justify-between items-center",
				utils.TwIf("border-destructive ring-destructive", props.HasError),
			),
			Disabled: props.Disabled,
			Attributes: utils.MergeAttributes(
				templ.Attributes{
					"@click": "toggleTimePicker",
					"id":     props.ID,
				},
				props.Attributes,
			),
		}).Render(templ.WithChildren(ctx, templ_7745c5c3_Var11), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 16, "</div><div x-show=\"open\" x-ref=\"timePickerPopup\" @click.away=\"closeTimePicker\" x-transition.opacity class=\"absolute left-0 z-50 w-72 p-4 rounded-lg border bg-popover shadow-md\" :class=\"positionClass\" @resize.window=\"updatePosition\" style=\"display: none;\"><div class=\"grid grid-cols-2 gap-2 flex-1\"><div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var17 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
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
			templ_7745c5c3_Var18 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
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
				templ_7745c5c3_Var19 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
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
					if props.Value != (time.Time{}) {
						if props.Use12Hours {
							templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 17, " ")
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
							hour := props.Value.Hour() % 12
							if hour == 0 {
								hour = 12
							}
							var templ_7745c5c3_Var20 string
							templ_7745c5c3_Var20, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("%02d", hour))
							if templ_7745c5c3_Err != nil {
								return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/time_picker.templ`, Line: 286, Col: 37}
							}
							_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var20))
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
						} else {
							var templ_7745c5c3_Var21 string
							templ_7745c5c3_Var21, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("%02d", props.Value.Hour()))
							if templ_7745c5c3_Err != nil {
								return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/time_picker.templ`, Line: 288, Col: 51}
							}
							_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var21))
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
						}
					} else {
						templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 18, "Hour")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
					}
					return nil
				})
				templ_7745c5c3_Err = SelectValue(SelectValueProps{
					Placeholder: "Hour",
				}).Render(templ.WithChildren(ctx, templ_7745c5c3_Var19), templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				return nil
			})
			templ_7745c5c3_Err = SelectTrigger(SelectTriggerProps{
				Name: "hour",
			}).Render(templ.WithChildren(ctx, templ_7745c5c3_Var18), templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 19, " ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Var22 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
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
				templ_7745c5c3_Var23 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
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
					for _, option := range generateHourOptions(props.Use12Hours) {
						templ_7745c5c3_Var24 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
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
							var templ_7745c5c3_Var25 string
							templ_7745c5c3_Var25, templ_7745c5c3_Err = templ.JoinStringErrs(option.Label)
							if templ_7745c5c3_Err != nil {
								return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/time_picker.templ`, Line: 306, Col: 24}
							}
							_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var25))
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
							return nil
						})
						templ_7745c5c3_Err = SelectItem(SelectItemProps{
							Value: option.Value,
							Selected: props.Value != (time.Time{}) &&
								((props.Use12Hours &&
									((props.Value.Hour()%12 == 0 && option.Value == "12") ||
										(props.Value.Hour()%12 == parseInt(option.Value)))) ||
									(!props.Use12Hours && props.Value.Hour() == parseInt(option.Value))),
						}).Render(templ.WithChildren(ctx, templ_7745c5c3_Var24), templ_7745c5c3_Buffer)
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
					}
					return nil
				})
				templ_7745c5c3_Err = SelectGroup(SelectGroupProps{}).Render(templ.WithChildren(ctx, templ_7745c5c3_Var23), templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				return nil
			})
			templ_7745c5c3_Err = SelectContent(SelectContentProps{}).Render(templ.WithChildren(ctx, templ_7745c5c3_Var22), templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return nil
		})
		templ_7745c5c3_Err = Select(SelectProps{
			ID:    "hour-select-" + props.ID,
			Class: "w-full",
		}).Render(templ.WithChildren(ctx, templ_7745c5c3_Var17), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 20, "</div><div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var26 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
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
			templ_7745c5c3_Var27 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
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
				templ_7745c5c3_Var28 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
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
					if props.Value != (time.Time{}) {
						var templ_7745c5c3_Var29 string
						templ_7745c5c3_Var29, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("%02d", props.Value.Minute()))
						if templ_7745c5c3_Err != nil {
							return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/time_picker.templ`, Line: 325, Col: 52}
						}
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var29))
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
					} else {
						templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 21, "Minute")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
					}
					return nil
				})
				templ_7745c5c3_Err = SelectValue(SelectValueProps{
					Placeholder: "Minute",
				}).Render(templ.WithChildren(ctx, templ_7745c5c3_Var28), templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				return nil
			})
			templ_7745c5c3_Err = SelectTrigger(SelectTriggerProps{
				Name: "minute",
			}).Render(templ.WithChildren(ctx, templ_7745c5c3_Var27), templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 22, " ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Var30 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
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
				templ_7745c5c3_Var31 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
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
					for _, option := range generateMinuteOptions() {
						templ_7745c5c3_Var32 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
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
							var templ_7745c5c3_Var33 string
							templ_7745c5c3_Var33, templ_7745c5c3_Err = templ.JoinStringErrs(option.Label)
							if templ_7745c5c3_Err != nil {
								return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/time_picker.templ`, Line: 338, Col: 24}
							}
							_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var33))
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
							return nil
						})
						templ_7745c5c3_Err = SelectItem(SelectItemProps{
							Value:    option.Value,
							Selected: props.Value != (time.Time{}) && props.Value.Minute() == parseInt(option.Value),
						}).Render(templ.WithChildren(ctx, templ_7745c5c3_Var32), templ_7745c5c3_Buffer)
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
					}
					return nil
				})
				templ_7745c5c3_Err = SelectGroup(SelectGroupProps{}).Render(templ.WithChildren(ctx, templ_7745c5c3_Var31), templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				return nil
			})
			templ_7745c5c3_Err = SelectContent(SelectContentProps{}).Render(templ.WithChildren(ctx, templ_7745c5c3_Var30), templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return nil
		})
		templ_7745c5c3_Err = Select(SelectProps{
			ID:    "minute-select-" + props.ID,
			Class: "w-full",
		}).Render(templ.WithChildren(ctx, templ_7745c5c3_Var26), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 23, "</div></div><div class=\"flex justify-between mt-4 gap-2\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if props.Use12Hours {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 24, "<div class=\"flex justify-center gap-2\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var34 = []any{utils.TwMerge(
				"px-3 py-1 text-sm rounded-md hover:bg-muted",
				utils.TwIf("bg-primary text-primary-foreground", props.Value != (time.Time{}) && props.Value.Hour() < 12),
			)}
			templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var34...)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 25, "<button type=\"button\" data-period=\"AM\" @click=\"selectPeriod\" class=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var35 string
			templ_7745c5c3_Var35, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var34).String())
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/time_picker.templ`, Line: 1, Col: 0}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var35))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 26, "\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var36 string
			templ_7745c5c3_Var36, templ_7745c5c3_Err = templ.JoinStringErrs(props.AMLabel)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/time_picker.templ`, Line: 358, Col: 22}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var36))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 27, "</button> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var37 = []any{utils.TwMerge(
				"px-3 py-1 text-sm rounded-md hover:bg-muted",
				utils.TwIf("bg-primary text-primary-foreground", props.Value != (time.Time{}) && props.Value.Hour() >= 12),
			)}
			templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var37...)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 28, "<button type=\"button\" data-period=\"PM\" @click=\"selectPeriod\" class=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var38 string
			templ_7745c5c3_Var38, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var37).String())
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/time_picker.templ`, Line: 1, Col: 0}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var38))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 29, "\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var39 string
			templ_7745c5c3_Var39, templ_7745c5c3_Err = templ.JoinStringErrs(props.PMLabel)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/time_picker.templ`, Line: 369, Col: 22}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var39))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 30, "</button></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 31, "<div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Var40 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
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
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 32, "Done")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return nil
		})
		templ_7745c5c3_Err = Button(ButtonProps{
			Type:    "button",
			Variant: ButtonVariantSecondary,
			Attributes: templ.Attributes{
				"@click": "closeTimePicker",
			},
		}).Render(templ.WithChildren(ctx, templ_7745c5c3_Var40), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 33, "</div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

// Helper function to safely parse integers
func parseInt(str string) int {
	val, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return val
}

var _ = templruntime.GeneratedTemplate
