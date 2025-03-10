package shared

type SideLink struct {
	Text  string
	Href  string
	Icon  string
	Click string
}

type Section struct {
	Title string
	Links []SideLink
}

var Sections = []Section{
	{
		Title: "Getting Started",
		Links: []SideLink{
			{
				Text: "Introduction",
				Href: "/docs/introduction",
			},
			{
				Text: "How to Use",
				Href: "/docs/how-to-use",
			},
			{
				Text: "Themes",
				Href: "/docs/themes",
			},
		},
	},
	{
		Title: "Components",
		Links: []SideLink{
			{
				Text: "Accordion",
				Href: "/docs/components/accordion",
			},
			{
				Text: "Alert",
				Href: "/docs/components/alert",
			},
			{
				Text: "Aspect Ratio",
				Href: "/docs/components/aspect-ratio",
			},
			{
				Text: "Avatar",
				Href: "/docs/components/avatar",
			},
			{
				Text: "Badge",
				Href: "/docs/components/badge",
			},
			{
				Text: "Breadcrumb",
				Href: "/docs/components/breadcrumb",
			},
			{
				Text: "Button",
				Href: "/docs/components/button",
			},
			{
				Text: "Card",
				Href: "/docs/components/card",
			},
			{
				Text: "Code",
				Href: "/docs/components/code",
			},
			{
				Text: "Checkbox",
				Href: "/docs/components/checkbox",
			},
			{
				Text: "Checkbox Card",
				Href: "/docs/components/checkbox-card",
			},
			{
				Text: "Datepicker",
				Href: "/docs/components/datepicker",
			},
			{
				Text: "Dropdown Menu",
				Href: "/docs/components/dropdown-menu",
			},
			{
				Text: "Form",
				Href: "/docs/components/form",
			},
			{
				Text: "Icon",
				Href: "/docs/components/icon",
			},
			{
				Text: "Input",
				Href: "/docs/components/input",
			},
			{
				Text: "Label",
				Href: "/docs/components/label",
			},
			{
				Text: "Modal",
				Href: "/docs/components/modal",
			},
			{
				Text: "Radio",
				Href: "/docs/components/radio",
			},
			{
				Text: "RadioCard",
				Href: "/docs/components/radio-card",
			},
			{
				Text: "Select",
				Href: "/docs/components/select",
			},
			{
				Text: "Sheet",
				Href: "/docs/components/sheet",
			},
			{
				Text: "Slider",
				Href: "/docs/components/slider",
			},
			{
				Text: "Tabs",
				Href: "/docs/components/tabs",
			},
			{
				Text: "Textarea",
				Href: "/docs/components/textarea",
			},
			{
				Text: "Timepicker",
				Href: "/docs/components/timepicker",
			},
			{
				Text: "Toast",
				Href: "/docs/components/toast",
			},
			{
				Text: "Toggle",
				Href: "/docs/components/toggle",
			},
			{
				Text: "Tooltip",
				Href: "/docs/components/tooltip",
			},
		},
	},
}
