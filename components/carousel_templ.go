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
)

type CarouselSlide struct {
	Content     templ.Component
	Description string
	Attributes  templ.Attributes
}

type CarouselProps struct {
	ID             string
	Slides         []CarouselSlide
	ShowControls   bool
	ShowIndicators bool
	Autoplay       bool
	Interval       int
	Loop           bool
	Class          string
	Attributes     templ.Attributes
}

func CarouselScript() templ.Component {
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
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/carousel.templ`, Line: 31, Col: 43}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "\">\n            function initCarousel(carouselEl) {\n                const state = {\n                    currentIndex: 0,\n                    slideCount: parseInt(carouselEl.dataset.slideCount || 0),\n                    autoplay: carouselEl.dataset.autoplay === 'true',\n                    interval: parseInt(carouselEl.dataset.interval || 5000),\n                    loop: carouselEl.dataset.loop === 'true',\n                    autoplayInterval: null,\n                    isHovering: false,\n                    touchStartX: 0\n                };\n                \n                const track = carouselEl.querySelector('.carousel-track');\n                const indicators = Array.from(carouselEl.querySelectorAll('.carousel-indicator'));\n                const prevBtn = carouselEl.querySelector('.carousel-prev');\n                const nextBtn = carouselEl.querySelector('.carousel-next');\n                \n                function updateTrackPosition() {\n                    if (track) {\n                        track.style.transform = `translateX(-${state.currentIndex * 100}%)`;\n                    }\n                }\n                \n                function updateIndicators() {\n                    indicators.forEach((indicator, i) => {\n                        if (i === state.currentIndex) {\n                            indicator.classList.add('bg-white');\n                            indicator.classList.remove('bg-white/50');\n                        } else {\n                            indicator.classList.remove('bg-white');\n                            indicator.classList.add('bg-white/50');\n                        }\n                    });\n                }\n                \n                function updateButtons() {\n                    if (prevBtn) {\n                        prevBtn.disabled = !state.loop && state.currentIndex === 0;\n                        if (prevBtn.disabled) {\n                            prevBtn.classList.add('opacity-50', 'cursor-not-allowed');\n                        } else {\n                            prevBtn.classList.remove('opacity-50', 'cursor-not-allowed');\n                        }\n                    }\n                    \n                    if (nextBtn) {\n                        nextBtn.disabled = !state.loop && state.currentIndex === state.slideCount - 1;\n                        if (nextBtn.disabled) {\n                            nextBtn.classList.add('opacity-50', 'cursor-not-allowed');\n                        } else {\n                            nextBtn.classList.remove('opacity-50', 'cursor-not-allowed');\n                        }\n                    }\n                }\n                \n                function startAutoplay() {\n                    if (state.autoplayInterval) {\n                        clearInterval(state.autoplayInterval);\n                    }\n                    \n                    state.autoplayInterval = setInterval(() => {\n                        if (!state.isHovering) {\n                            goToNext();\n                        }\n                    }, state.interval);\n                }\n                \n                function stopAutoplay() {\n                    if (state.autoplayInterval) {\n                        clearInterval(state.autoplayInterval);\n                        state.autoplayInterval = null;\n                    }\n                }\n                \n                function goToNext() {\n                    let nextIndex = state.currentIndex + 1;\n                    \n                    if (nextIndex >= state.slideCount) {\n                        if (state.loop) {\n                            nextIndex = 0;\n                        } else {\n                            nextIndex = state.slideCount - 1;\n                        }\n                    }\n                    \n                    goToSlide(nextIndex);\n                }\n                \n                function goToPrev() {\n                    let prevIndex = state.currentIndex - 1;\n                    \n                    if (prevIndex < 0) {\n                        if (state.loop) {\n                            prevIndex = state.slideCount - 1;\n                        } else {\n                            prevIndex = 0;\n                        }\n                    }\n                    \n                    goToSlide(prevIndex);\n                }\n                \n                function goToSlide(index) {\n                    if (index === state.currentIndex) return;\n                    \n                    state.currentIndex = index;\n                    updateTrackPosition();\n                    updateIndicators();\n                    updateButtons();\n                    \n                    if (state.autoplay) {\n                        stopAutoplay();\n                        if (!state.isHovering) {\n                            startAutoplay();\n                        }\n                    }\n                }\n                \n                function handleTouchStart(event) {\n                    state.touchStartX = event.touches[0].clientX;\n                }\n                \n                function handleTouchEnd(event) {\n                    const touchEndX = event.changedTouches[0].clientX;\n                    const diff = state.touchStartX - touchEndX;\n                    \n                    if (Math.abs(diff) > 50) {\n                        if (diff > 0) {\n                            goToNext();\n                        } else {\n                            goToPrev();\n                        }\n                    }\n                }\n                \n                if (track) {\n                    track.addEventListener('touchstart', handleTouchStart);\n                    track.addEventListener('touchend', handleTouchEnd);\n                }\n                \n                indicators.forEach((indicator, index) => {\n                    indicator.addEventListener('click', () => goToSlide(index));\n                });\n                \n                if (prevBtn) {\n                    prevBtn.addEventListener('click', goToPrev);\n                }\n                \n                if (nextBtn) {\n                    nextBtn.addEventListener('click', goToNext);\n                }\n                \n                carouselEl.addEventListener('mouseenter', () => {\n                    state.isHovering = true;\n                    if (state.autoplay) {\n                        stopAutoplay();\n                    }\n                });\n                \n                carouselEl.addEventListener('mouseleave', () => {\n                    state.isHovering = false;\n                    if (state.autoplay) {\n                        startAutoplay();\n                    }\n                });\n                \n                updateTrackPosition();\n                updateIndicators();\n                updateButtons();\n                \n                if (state.autoplay) {\n                    startAutoplay();\n                }\n            }\n            \n            document.addEventListener('DOMContentLoaded', () => {\n                document.querySelectorAll('.carousel-component').forEach(carousel => {\n                    initCarousel(carousel);\n                });\n            });\n        </script>")
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

func Carousel(props CarouselProps) templ.Component {
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
		templ_7745c5c3_Err = CarouselScript().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 = []any{utils.TwMerge(
			"carousel-component relative overflow-hidden w-full",
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
		templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(props.ID)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/carousel.templ`, Line: 222, Col: 15}
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
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/carousel.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "\" data-slide-count=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var8 string
		templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("%d", len(props.Slides)))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/carousel.templ`, Line: 227, Col: 57}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "\" data-autoplay=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var9 string
		templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.FormatBool(props.Autoplay))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/carousel.templ`, Line: 228, Col: 52}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 7, "\" data-interval=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var10 string
		templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("%d", func() int {
			if props.Interval == 0 {
				return 5000
			}
			return props.Interval
		}()))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/carousel.templ`, Line: 234, Col: 12}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 8, "\" data-loop=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var11 string
		templ_7745c5c3_Var11, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.FormatBool(props.Loop))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/carousel.templ`, Line: 235, Col: 44}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var11))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 9, "\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.RenderAttributes(ctx, templ_7745c5c3_Buffer, props.Attributes)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 10, "><div class=\"carousel-track flex h-full w-full transition-transform duration-500 ease-in-out\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, slide := range props.Slides {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 11, "<div class=\"carousel-slide flex-shrink-0 w-full h-full relative\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templ.RenderAttributes(ctx, templ_7745c5c3_Buffer, slide.Attributes)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 12, ">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = slide.Content.Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if slide.Description != "" {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 13, "<div class=\"absolute bottom-0 left-0 right-0 bg-black/50 text-white p-4\"><p>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var12 string
				templ_7745c5c3_Var12, templ_7745c5c3_Err = templ.JoinStringErrs(slide.Description)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/carousel.templ`, Line: 244, Col: 29}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var12))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 14, "</p></div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 15, "</div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 16, "</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if props.ShowControls {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 17, "<button class=\"carousel-prev absolute left-2 top-1/2 transform -translate-y-1/2 p-2 rounded-full bg-black/20 text-white hover:bg-black/40 focus:outline-none\" aria-label=\"Previous slide\" type=\"button\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = icons.ChevronLeft(icons.IconProps{}).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 18, "</button> <button class=\"carousel-next absolute right-2 top-1/2 transform -translate-y-1/2 p-2 rounded-full bg-black/20 text-white hover:bg-black/40 focus:outline-none\" aria-label=\"Next slide\" type=\"button\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = icons.ChevronRight(icons.IconProps{}).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 19, "</button> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if props.ShowIndicators {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 20, "<div class=\"absolute bottom-4 left-1/2 transform -translate-x-1/2 flex gap-2\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			for i := range props.Slides {
				var templ_7745c5c3_Var13 = []any{utils.TwMerge(
					"carousel-indicator w-3 h-3 rounded-full bg-white/50 hover:bg-white/80 focus:outline-none transition-colors",
					func() string {
						if i == 0 {
							return "bg-white"
						}
						return ""
					}(),
				)}
				templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var13...)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 21, "<button class=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var14 string
				templ_7745c5c3_Var14, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var13).String())
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/carousel.templ`, Line: 1, Col: 0}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var14))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 22, "\" aria-label=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var15 string
				templ_7745c5c3_Var15, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("Go to slide %d", i+1))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/carousel.templ`, Line: 279, Col: 53}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var15))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 23, "\" type=\"button\"></button>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 24, "</div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 25, "</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
