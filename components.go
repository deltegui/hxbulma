package bulma

import (
	"strconv"
	"strings"

	"github.com/deltegui/hx"
)

// Crumb is an item of a Breadcrumb. The last one is usually Current.
type Crumb struct {
	Label, Href string
	Current     bool
}

// Breadcrumb is the trail "Home / Products / Edit". Chain Centered, Right or a
// separator (HasArrowSeparator...).
func Breadcrumb(crumbs ...Crumb) hx.INode {
	items := make([]hx.INode, len(crumbs))
	for i, c := range crumbs {
		a := hx.A().Href(c.Href).Text(c.Label)
		if c.Current {
			a.Attribute("aria-current", "page")
		}
		items[i] = hx.Li().ClassIf(Active, c.Current).Body(a)
	}
	return hx.Nav().Class("breadcrumb").Attribute("aria-label", "breadcrumbs").Body(hx.Ul().BodyList(items))
}

// Separators of a Breadcrumb.
const (
	HasArrowSeparator    = "has-arrow-separator"
	HasBulletSeparator   = "has-bullet-separator"
	HasDotSeparator      = "has-dot-separator"
	HasSucceedsSeparator = "has-succeeds-separator"
)

// Card is a box with parts: put CardHeader, CardImage, CardContent and
// CardFooter in it.
func Card(children ...hx.INode) hx.INode { return hx.Div().Class("card").Body(children...) }

// CardHeader is the title bar of a card. Chain Centered on it to center the title.
func CardHeader(title string, extra ...hx.INode) hx.INode {
	return hx.El("header").Class("card-header").Body(
		append([]hx.INode{hx.P().Class("card-header-title").Text(title)}, extra...)...)
}

// CardHeaderIcon is a button at the end of the header (the arrow that folds a card).
func CardHeaderIcon(icon string, onClick func()) hx.INode {
	b := hx.Button().Type("button").Class("card-header-icon").Attribute("aria-label", "more options").Body(Icon(icon))
	if onClick != nil {
		b.OnClick(func(hx.EventContext) { onClick() })
	}
	return b
}

func CardImage(image hx.INode) hx.INode { return hx.Div().Class("card-image").Body(image) }

func CardContent(children ...hx.INode) hx.INode {
	return hx.Div().Class("card-content").Body(children...)
}

func CardFooter(children ...hx.INode) hx.INode {
	return hx.El("footer").Class("card-footer").Body(children...)
}

// CardFooterItem is one of the actions of the footer. If onClick is nil it is
// plain text; if href is empty it is a button.
func CardFooterItem(label, href string, onClick func()) hx.INode {
	if href != "" {
		return hx.A().Href(href).Class("card-footer-item").Text(label)
	}
	if onClick == nil {
		return hx.P().Class("card-footer-item").Text(label)
	}
	return hx.A().Attribute("role", "button").Class("card-footer-item").Text(label).
		OnClick(func(hx.EventContext) { onClick() })
}

// Dropdown is a button that opens a menu. open says if it is open and the button
// toggles it; items go in the menu (DropdownItem, DropdownDivider). Choosing an
// item does not close it: do open.Set(false) in the handler of the item.
func Dropdown(open *hx.SignalT[bool], label string, items ...hx.INode) hx.INode {
	button := hx.Button().Type("button").Class("button").Attribute("aria-haspopup", "true").Body(
		hx.Span().Text(label),
		hx.Span().Class("icon", "is-small").Body(hx.I().Class("dropdown-arrow").Text("▾")))
	button.OnClick(func(hx.EventContext) { open.Set(!open.Get()) })
	button.Effect(func() { button.Attribute("aria-expanded", strconv.FormatBool(open.Get())) })
	drop := hx.Div().Class("dropdown").Body(
		hx.Div().Class("dropdown-trigger").Body(button),
		hx.Div().Class("dropdown-menu").Attribute("role", "menu").Body(
			hx.Div().Class("dropdown-content").Body(items...)))
	return BindClass(drop, Active, open)
}

// DropdownItem is a link of a Dropdown; with href empty it is a button.
func DropdownItem(label, href string, onClick func()) hx.INode {
	a := hx.A().Class("dropdown-item").Text(label)
	if href != "" {
		a.Href(href)
	}
	if onClick != nil {
		a.OnClick(func(hx.EventContext) { onClick() })
	}
	return a
}

func DropdownDivider() hx.INode { return hx.Hr().Class("dropdown-divider") }

// Menu is a vertical menu for the side of a page. Put MenuLabel and MenuList in it.
func Menu(children ...hx.INode) hx.INode { return hx.Aside().Class("menu").Body(children...) }

func MenuLabel(text string) hx.INode { return hx.P().Class("menu-label").Text(text) }

func MenuList(children ...hx.INode) hx.INode { return hx.Ul().Class("menu-list").Body(children...) }

// MenuLink is an item of a MenuList. A nested MenuList goes in children.
func MenuLink(label, href string, active bool, children ...hx.INode) hx.INode {
	return hx.Li().Body(append([]hx.INode{hx.A().Href(href).ClassIf(Active, active).Text(label)}, children...)...)
}

// Message is a colored box with a header. A title makes the header and onDelete
// adds its delete button; both can be empty. Chain a color.
func Message(title string, onDelete func(), body ...hx.INode) hx.INode {
	m := hx.El("article").Class("message")
	content := hx.Div().Class("message-body").Body(body...)
	if title == "" && onDelete == nil {
		return m.Body(content)
	}
	head := []hx.INode{hx.P().Text(title)}
	if onDelete != nil {
		head = append(head, DeleteButton(onDelete))
	}
	return m.Body(hx.Div().Class("message-header").BodyList(head), content)
}

// Modal is in the tree while open is true; build makes its content each time it
// opens, so a closed modal costs nothing. The content is whatever you build
// (a Box, a Card, an Image). onClose runs when the user clicks the background
// or the corner button.
func Modal(open hx.Gettable[bool], onClose func(), build func() hx.INode) hx.INode {
	return hx.Dynamic(func() hx.INode {
		if !open.Get() {
			return nil
		}
		var modal hx.INode
		hx.Untrack(func() {
			modal = hx.Div().Class("modal", Active).Body(
				hx.Div().Class("modal-background").OnClick(func(hx.EventContext) { onClose() }),
				hx.Div().Class("modal-content").Body(build()),
				hx.Button().Type("button").Class("modal-close", Large).Attribute("aria-label", "close").
					OnClick(func(hx.EventContext) { onClose() }))
		})
		return modal
	})
}

// ModalCard is a Modal with a title bar, a body and a footer with the buttons.
func ModalCard(open hx.Gettable[bool], title string, onClose func(), body, foot func() hx.INode) hx.INode {
	return hx.Dynamic(func() hx.INode {
		if !open.Get() {
			return nil
		}
		var modal hx.INode
		hx.Untrack(func() {
			modal = hx.Div().Class("modal", Active).Attribute("role", "dialog").Body(
				hx.Div().Class("modal-background").OnClick(func(hx.EventContext) { onClose() }),
				hx.Div().Class("modal-card").Body(
					hx.El("header").Class("modal-card-head").Body(
						hx.P().Class("modal-card-title").Text(title),
						hx.Button().Type("button").Class("delete").Attribute("aria-label", "close").
							OnClick(func(hx.EventContext) { onClose() })),
					hx.El("section").Class("modal-card-body").Body(body()),
					hx.El("footer").Class("modal-card-foot").Body(foot())))
		})
		return modal
	})
}

// Navbar is the bar at the top of the page. Put a NavbarBrand and a NavbarMenu in it.
func Navbar(children ...hx.INode) hx.INode {
	return hx.Nav().Class("navbar").Attribute("role", "navigation").Attribute("aria-label", "main navigation").Body(children...)
}

// NavbarBrand is the left part that is always visible: the logo and the burger.
func NavbarBrand(children ...hx.INode) hx.INode {
	return hx.Div().Class("navbar-brand").Body(children...)
}

// NavbarBurger is the button that opens the menu on a phone. It toggles open,
// which NavbarMenu reads.
func NavbarBurger(open *hx.SignalT[bool]) hx.INode {
	burger := hx.A().Attribute("role", "button").Class("navbar-burger").Attribute("aria-label", "menu").Body(
		hx.Span().Attribute("aria-hidden", "true"),
		hx.Span().Attribute("aria-hidden", "true"),
		hx.Span().Attribute("aria-hidden", "true"),
		hx.Span().Attribute("aria-hidden", "true"))
	burger.OnClick(func(hx.EventContext) { open.Set(!open.Get()) })
	burger.Effect(func() { burger.Attribute("aria-expanded", strconv.FormatBool(open.Get())) })
	return BindClass(burger, Active, open)
}

// NavbarMenu is the part that folds on a phone: a NavbarStart and a NavbarEnd.
func NavbarMenu(open hx.Gettable[bool], children ...hx.INode) hx.INode {
	return BindClass(hx.Div().Class("navbar-menu").Body(children...), Active, open)
}

func NavbarStart(children ...hx.INode) hx.INode {
	return hx.Div().Class("navbar-start").Body(children...)
}

func NavbarEnd(children ...hx.INode) hx.INode { return hx.Div().Class("navbar-end").Body(children...) }

// NavbarItem is a link of the navbar. Chain Active for the current page.
func NavbarItem(label, href string) hx.INode {
	return hx.A().Href(href).Class("navbar-item").Text(label)
}

// NavbarBox is an item that is not a link (buttons, a field).
func NavbarBox(children ...hx.INode) hx.INode {
	return hx.Div().Class("navbar-item").Body(children...)
}

// NavbarDropdown is an item that opens a list when the pointer is over it (and
// on the phone it is always open). Put NavbarItem and NavbarDivider in it.
func NavbarDropdown(label string, items ...hx.INode) hx.INode {
	return hx.Div().Class("navbar-item", "has-dropdown", Hoverable).Body(
		hx.A().Class("navbar-link").Text(label),
		hx.Div().Class("navbar-dropdown").Body(items...))
}

func NavbarDivider() hx.INode { return hx.Hr().Class("navbar-divider") }

// Pagination shows previous, next, and a window of pages around the current one
// with the first and the last. page and pages count from 1.
func Pagination(page, pages hx.Gettable[int], goTo func(int)) hx.INode {
	link := func(label string, target int, enabled, current bool) hx.INode {
		b := hx.Button().Type("button").Class("pagination-link").Text(label)
		if current {
			b.Class("is-current").Attribute("aria-current", "page")
		}
		if enabled && !current {
			b.OnClick(func(hx.EventContext) { goTo(target) })
		}
		return hx.Li().Body(b)
	}
	step := func(class, label string, target int, enabled bool) hx.INode {
		b := hx.Button().Type("button").Class(class).Text(label)
		if enabled {
			b.OnClick(func(hx.EventContext) { goTo(target) })
		} else {
			b.Attribute("disabled", "")
		}
		return b
	}
	return hx.Nav().Class("pagination").Attribute("role", "navigation").Attribute("aria-label", "pagination").Body(
		hx.Dynamic(func() hx.INode {
			current, last := page.Get(), pages.Get()
			items := []hx.INode{}
			from, to := max(1, current-2), min(last, current+2)
			if from > 1 {
				items = append(items, link("1", 1, true, current == 1))
				if from > 2 {
					items = append(items, hx.Li().Body(hx.Span().Class("pagination-ellipsis").Text("…")))
				}
			}
			for p := from; p <= to; p++ {
				items = append(items, link(strconv.Itoa(p), p, true, p == current))
			}
			if to < last {
				if to < last-1 {
					items = append(items, hx.Li().Body(hx.Span().Class("pagination-ellipsis").Text("…")))
				}
				items = append(items, link(strconv.Itoa(last), last, true, current == last))
			}
			return hx.Noop().Body(
				step("pagination-previous", "«", current-1, current > 1),
				step("pagination-next", "»", current+1, current < last),
				hx.Ul().Class("pagination-list").BodyList(items))
		}))
}

// Panel is a box with a heading, tabs and blocks (a list of filters, files...).
func Panel(heading string, children ...hx.INode) hx.INode {
	return hx.Nav().Class("panel").Body(
		append([]hx.INode{hx.P().Class("panel-heading").Text(heading)}, children...)...)
}

// PanelTabs is the row of filters under the heading of a Panel.
func PanelTabs(children ...hx.INode) hx.INode { return hx.P().Class("panel-tabs").Body(children...) }

// PanelBlock is a row of a Panel that is not a link (a search field).
func PanelBlock(children ...hx.INode) hx.INode {
	return hx.Div().Class("panel-block").Body(children...)
}

// PanelLink is a row of a Panel that is a link, with an icon. Chain Active.
func PanelLink(icon, label, href string) hx.INode {
	a := hx.A().Class("panel-block").Href(href)
	if icon != "" {
		a.Body(hx.Span().Class("panel-icon").Body(hx.I().Class(strings.Fields(icon)...)), hx.Span().Text(label))
		return a
	}
	return a.Text(label)
}

// Tab is an entry of Tabs.
type Tab struct{ ID, Label, Icon string }

// Tabs is a row of tabs. active holds the ID of the chosen one and clicking a tab
// sets it; show the content with TabPanel. Chain Boxed, Toggle, Centered,
// Fullwidth or a size.
func Tabs(active *hx.SignalT[string], tabs ...Tab) hx.INode {
	items := make([]hx.INode, len(tabs))
	for i, t := range tabs {
		a := hx.A().Attribute("role", "tab")
		if t.Icon != "" {
			a.Body(hx.Span().Class("icon", Small).Body(hx.I().Class(strings.Fields(t.Icon)...)), hx.Span().Text(t.Label))
		} else {
			a.Text(t.Label)
		}
		a.OnClick(func(hx.EventContext) { active.Set(t.ID) })
		li := hx.Li().Attribute("role", "presentation").Body(a)
		items[i] = li.Effect(func() { li.ClassIf(Active, active.Get() == t.ID) })
	}
	return hx.Div().Class("tabs").Body(hx.Ul().BodyList(items))
}

// TabPanel shows build() while the Tabs have id chosen.
func TabPanel(active hx.Gettable[string], id string, build func() hx.INode) hx.INode {
	return hx.Dynamic(func() hx.INode {
		if active.Get() != id {
			return nil
		}
		var panel hx.INode
		hx.Untrack(func() { panel = build() })
		return panel
	})
}
