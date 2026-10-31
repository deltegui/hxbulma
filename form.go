package bulma

import (
	"strings"

	"github.com/deltegui/hx"
)

// Field is a row of a form: a Label, a Control and a Help.
func Field(children ...hx.INode) hx.INode { return hx.Div().Class("field").Body(children...) }

// Control wraps one input so Bulma can place icons and buttons around it.
// Chain HasIconsLeft or HasIconsRight if it has icons.
func Control(children ...hx.INode) hx.INode { return hx.Div().Class("control").Body(children...) }

// Label is the label of a field.
func Label(text string) hx.INode { return hx.Label().Class("label").Text(text) }

// Help is the small text under a field. Chain a color to make it an error.
func Help(text string) hx.INode { return hx.P().Class("help").Text(text) }

const (
	HasIconsLeft  = "has-icons-left"
	HasIconsRight = "has-icons-right"
)

// FormField is the common case: a label, a control and the errors of that field
// as a red help (nothing if errs is nil or empty). The error text follows the
// signal.
func FormField(label string, control hx.INode, errs hx.Gettable[[]string]) hx.INode {
	f := Field()
	if label != "" {
		f.Body(Label(label), Control(control), errorHelp(errs))
		return f
	}
	return f.Body(Control(control), errorHelp(errs))
}

func errorHelp(errs hx.Gettable[[]string]) hx.INode {
	help := hx.P().Class("help", Danger)
	if errs != nil {
		help.Effect(func() { help.Text(strings.Join(errs.Get(), " · ")) })
	}
	return help
}

// Addons is a field whose controls are stuck together (an input with a button).
func Addons(children ...hx.INode) hx.INode { return Field(children...).Class(HasAddons) }

// Grouped is a field whose controls are side by side with a gap.
func Grouped(children ...hx.INode) hx.INode { return Field(children...).Class("is-grouped") }

// HorizontalField is a label next to its fields, for wide forms: the label
// on the left, the Fields on the right.
func HorizontalField(label string, fields ...hx.INode) hx.INode {
	return hx.Div().Class("field", "is-horizontal").Body(
		hx.Div().Class("field-label", "is-normal").Body(Label(label)),
		hx.Div().Class("field-body").Body(fields...))
}

// Input is an input of the given type ("text", "email", "password", "number",
// "date", "tel", "url", "search"...), bound to the signal.
func Input(kind string, value *hx.SignalT[string]) hx.INode {
	return hx.Input().Type(kind).Class("input").BindValue(value).BindOnInput(value)
}

// TextArea is a multi line text.
func TextArea(value *hx.SignalT[string]) hx.INode {
	return hx.TextArea().Class("textarea").BindValue(value).BindOnInput(value)
}

// Select is a dropdown list. It has an empty first option with the text of
// blank (none if blank is empty).
func Select(value *hx.SignalT[string], blank string, options []Option) hx.INode {
	sel := hx.Select()
	var nodes []hx.INode
	if blank != "" {
		nodes = append(nodes, hx.Option().Value("").Text(blank))
	}
	for _, o := range options {
		nodes = append(nodes, hx.Option().Value(o.Value).Text(o.Label))
	}
	sel.BodyList(nodes).BindValue(value).BindOnChange(value)
	return hx.Div().Class("select").Body(sel)
}

// Checkbox is a checkbox with its label, bound to the signal.
func Checkbox(label string, value *hx.SignalT[bool]) hx.INode {
	return hx.Label().Class("checkbox").Body(
		hx.Input().Type("checkbox").BindChecked(value).BindOnChecked(value),
		hx.Span().Text(" "+label))
}

// Radio is one radio button: it is checked while the signal has its value, and
// it sets the signal when chosen. Radios of a group share name and signal.
func Radio(name, value, label string, selected *hx.SignalT[string]) hx.INode {
	input := hx.Input().Type("radio").Attribute("name", name).Value(value)
	input.Effect(func() { input.Checked(selected.Get() == value) })
	input.OnChange(func(hx.EventContext) { selected.Set(value) })
	return hx.Label().Class("radio").Body(input, hx.Span().Text(" "+label))
}

// Radios is a Control with a Radio for each option.
func Radios(name string, selected *hx.SignalT[string], options []Option) hx.INode {
	nodes := make([]hx.INode, len(options))
	for i, o := range options {
		nodes[i] = Radio(name, o.Value, o.Label, selected)
	}
	return Control(nodes...).Class("radios")
}

// File is a file picker: a button with label, and the name of the chosen file
// when filename is not nil. onChange runs when the user chooses; read the files
// with hxhttp.FilesOf(ctx.Target) in the browser.
func File(label string, filename hx.Gettable[string], onChange func(hx.EventContext)) hx.INode {
	input := hx.Input().Type("file").Class("file-input").Attribute("name", "file")
	if onChange != nil {
		input.OnChange(onChange)
	}
	parts := []hx.INode{input, hx.Span().Class("file-cta").Body(
		hx.Span().Class("file-label").Text(label))}
	if filename != nil {
		name := hx.Span().Class("file-name")
		name.Effect(func() { name.Text(filename.Get()) })
		parts = append(parts, name)
	}
	return hx.Div().Class("file").Body(hx.Label().Class("file-label").BodyList(parts))
}
