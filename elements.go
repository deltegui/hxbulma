package bulma

import (
	"strconv"
	"strings"

	"github.com/deltegui/hx"
)

// Block is a div with the spacing of Bulma between blocks.
func Block(children ...hx.INode) hx.INode { return hx.Div().Class("block").Body(children...) }

// Box is a white box with a shadow and a padding.
func Box(children ...hx.INode) hx.INode { return hx.Div().Class("box").Body(children...) }

// Content is the container that styles plain HTML (headings, lists, tables,
// quotes) inside it, for text that has no classes.
func Content(children ...hx.INode) hx.INode { return hx.Div().Class("content").Body(children...) }

// Button is a button that does not submit. Chain its modifiers: .Class(Primary).
func Button(label string) hx.INode {
	return hx.Button().Type("button").Class("button").Text(label)
}

// SubmitButton submits the form it is in.
func SubmitButton(label string) hx.INode {
	return hx.Button().Type("submit").Class("button").Text(label)
}

// ButtonLink is an anchor that looks like a button.
func ButtonLink(href, label string) hx.INode {
	return hx.A().Href(href).Class("button").Text(label)
}

// Buttons groups buttons. With the class HasAddons they are stuck together.
func Buttons(children ...hx.INode) hx.INode { return hx.Div().Class("buttons").Body(children...) }

// DeleteButton is the small round "x" of notifications, messages and tags.
func DeleteButton(onClick func()) hx.INode {
	b := hx.Button().Type("button").Class("delete").Attribute("aria-label", "delete")
	if onClick != nil {
		b.OnClick(func(hx.EventContext) { onClick() })
	}
	return b
}

// Icon is a box for an icon of the font you use (Bulma has none). class is the
// class of that font: Icon("fas fa-home"), Icon("mdi mdi-home").
func Icon(class string) hx.INode {
	return hx.Span().Class("icon").Body(hx.I().Class(strings.Fields(class)...))
}

// IconText puts an icon before a text, on one line.
func IconText(iconClass, text string) hx.INode {
	return hx.Span().Class("icon-text").Body(Icon(iconClass), hx.Span().Text(text))
}

// Image is a figure with an image. Chain a size (.Class("is-64x64")) or a ratio
// (Ratio16x9).
func Image(src, alt string) hx.INode {
	return hx.El("figure").Class("image").Body(hx.Img().Src(src).Attribute("alt", alt))
}

// Notification is a colored box that stands out. Chain a color.
func Notification(children ...hx.INode) hx.INode {
	return hx.Div().Class("notification").Body(children...)
}

// DismissibleNotification is a Notification with a delete button in the corner.
func DismissibleNotification(onDelete func(), children ...hx.INode) hx.INode {
	return hx.Div().Class("notification").Body(append([]hx.INode{DeleteButton(onDelete)}, children...)...)
}

// Progress is a bar that is value of max full.
func Progress(value, max int) hx.INode {
	return hx.El("progress").Class("progress").
		Attribute("value", strconv.Itoa(value)).Attribute("max", strconv.Itoa(max)).
		Text(strconv.Itoa(value))
}

// ProgressOf is a Progress that follows a signal.
func ProgressOf(value hx.Gettable[int], max int) hx.INode {
	p := hx.El("progress").Class("progress").Attribute("max", strconv.Itoa(max))
	return p.Effect(func() {
		v := strconv.Itoa(value.Get())
		p.Attribute("value", v).Text(v)
	})
}

// Indeterminate is a Progress that is not full of anything: it only moves.
func Indeterminate() hx.INode {
	return hx.El("progress").Class("progress").Attribute("max", "100")
}

// SkeletonBlock is a gray block that moves, to stand in for content that loads.
func SkeletonBlock() hx.INode { return hx.Div().Class("skeleton-block") }

// SkeletonLines stands in for a paragraph of n lines.
func SkeletonLines(n int) hx.INode {
	lines := make([]hx.INode, n)
	for i := range lines {
		lines[i] = hx.Div()
	}
	return hx.Div().Class("skeleton-lines").BodyList(lines)
}

// Table is a table. Put THead, TBody and TFoot in it, or use SimpleTable.
// Chain Bordered, Striped, Narrow, Hoverable, Fullwidth.
func Table(children ...hx.INode) hx.INode { return hx.Table().Class("table").Body(children...) }

// SimpleTable is a table with a header row of headers and body inside a TBody.
func SimpleTable(headers []string, rows ...hx.INode) hx.INode {
	cells := make([]hx.INode, len(headers))
	for i, h := range headers {
		cells[i] = hx.Th().Attribute("scope", "col").Text(h)
	}
	return Table(hx.THead().Body(hx.Tr().BodyList(cells)), hx.TBody().Body(rows...))
}

// TableContainer lets a wide table scroll sideways instead of breaking the page.
func TableContainer(table hx.INode) hx.INode { return hx.Div().Class("table-container").Body(table) }

// Tag is a small label. Chain a color or a size.
func Tag(text string) hx.INode { return hx.Span().Class("tag").Text(text) }

// DeletableTag is a Tag with a delete button after the text.
func DeletableTag(text string, onDelete func()) hx.INode {
	return hx.Span().Class("tag").Body(hx.Span().Text(text), DeleteButton(onDelete).Class(Small))
}

// Tags groups tags. With HasAddons the tags are stuck together.
func Tags(children ...hx.INode) hx.INode { return hx.Div().Class("tags").Body(children...) }

// Title is a heading of level 1 to 6 (<h1>..<h6>) with the size of that level.
func Title(level int, text string) hx.INode {
	return heading(level).Class("title", "is-"+strconv.Itoa(level)).Text(text)
}

// Subtitle is a lighter heading of level 1 to 6, to go under a Title.
func Subtitle(level int, text string) hx.INode {
	return heading(level).Class("subtitle", "is-"+strconv.Itoa(level)).Text(text)
}

func heading(level int) hx.INode {
	switch level {
	case 1:
		return hx.H1()
	case 2:
		return hx.H2()
	case 3:
		return hx.H3()
	case 4:
		return hx.H4()
	case 5:
		return hx.H5()
	}
	return hx.H6()
}

// HasAddons sticks the buttons of Buttons, the tags of Tags and the controls of a
// field together.
const HasAddons = "has-addons"
