package bulma_test

import (
	"strings"
	"testing"

	"github.com/deltegui/hx"
	"github.com/deltegui/hxbulma"
	"github.com/deltegui/hx/ssr"
)

func render(build func() hx.INode) string { return ssr.Render(build) }

func TestComponents(t *testing.T) {
	text := hx.Signal("hi")
	flag := hx.Signal(true)
	off := hx.Signal(false)
	num := hx.Signal(3)
	tab := hx.Signal("b")
	choice := hx.Signal("x")
	nop := func() {}
	node := func() hx.INode { return hx.P().Text("x") }

	cases := []struct {
		name  string
		build func() hx.INode
		want  []string
	}{
		{"button", func() hx.INode { return bulma.Button("Go").Class(bulma.Primary, bulma.Large) }, []string{`<button`, `class="button is-large is-primary"`, `type="button"`, `>Go</button>`}},
		{"block box content", func() hx.INode { return bulma.Block(bulma.Box(bulma.Content(node()))) }, []string{`class="block"`, `class="box"`, `class="content"`}},
		{"buttons", func() hx.INode { return bulma.Buttons(bulma.SubmitButton("a"), bulma.ButtonLink("/x", "b")) }, []string{`class="buttons"`, `type="submit"`, `href="/x"`}},
		{"delete", func() hx.INode { return bulma.DeleteButton(nop) }, []string{`class="delete"`}},
		{"icon", func() hx.INode { return bulma.IconText("fas fa-home", "Home") }, []string{`icon-text`, `class="icon"`, `fa-home fas`}},
		{"image", func() hx.INode { return bulma.Image("/a.png", "A").Class(bulma.Ratio16x9) }, []string{`<figure class="image is-16by9"`, `alt="A"`}},
		{"notification", func() hx.INode { return bulma.DismissibleNotification(nop, node()).Class(bulma.Success) }, []string{`is-success notification`, `class="delete"`}},
		{"progress", func() hx.INode { return bulma.ProgressOf(num, 10) }, []string{`<progress`, `value="3"`, `max="10"`}},
		{"skeleton", func() hx.INode { return bulma.Block(bulma.SkeletonBlock(), bulma.SkeletonLines(3)) }, []string{`skeleton-block`, `skeleton-lines`, `<div></div><div></div><div></div>`}},
		{"table", func() hx.INode {
			return bulma.TableContainer(bulma.SimpleTable([]string{"A", "B"}, hx.Tr())).Class("x")
		}, []string{`table-container`, `class="table"`, `<th scope="col">A</th>`}},
		{"tags", func() hx.INode { return bulma.Tags(bulma.Tag("a").Class(bulma.Info), bulma.DeletableTag("b", nop)) }, []string{`class="tags"`, `is-info tag`, `class="delete is-small"`}},
		{"title", func() hx.INode { return bulma.Block(bulma.Title(2, "T"), bulma.Subtitle(4, "S")) }, []string{`<h2 class="is-2 title">T</h2>`, `<h4 class="is-4 subtitle">S</h4>`}},
		{"field", func() hx.INode {
			return bulma.FormField("Name", bulma.Input("text", text), hx.Signal([]string{"bad", "worse"}))
		}, []string{`class="field"`, `class="label"`, `class="control"`, `class="input"`, `help is-danger`, `bad · worse`}},
		{"form controls", func() hx.INode {
			return hx.Div().Body(bulma.TextArea(text), bulma.Select(text, "—", bulma.Plain("a", "b")), bulma.Checkbox("ok", off), bulma.Radios("r", choice, bulma.Plain("x", "y")), bulma.File("Pick", text, nil))
		}, []string{`class="textarea"`, `class="select"`, `<option>—</option>`, `class="checkbox"`, `type="radio"`, `class="file-input"`, `file-cta`, `class="file-name"`}},
		{"addons", func() hx.INode {
			return bulma.Addons(bulma.Control(bulma.Input("text", text)), bulma.Control(bulma.Button("Go")))
		}, []string{`field has-addons`}},
		{"horizontal", func() hx.INode { return bulma.HorizontalField("A", bulma.Field()) }, []string{`field is-horizontal`, `field-label is-normal`, `field-body`}},
		{"layout", func() hx.INode {
			return bulma.Container(bulma.Section(bulma.Columns(bulma.Column(node()).Class(bulma.Span(4), bulma.Offset(2)))))
		}, []string{`class="container"`, `class="section"`, `class="columns"`, `column is-4 is-offset-2`}},
		{"hero", func() hx.INode { return bulma.SimpleHero("T", "S").Class(bulma.Primary) }, []string{`hero is-primary`, `hero-body`, `is-1 title`}},
		{"level media tile", func() hx.INode {
			return bulma.Block(bulma.Level(bulma.LevelLeft(bulma.LevelItem(node()))), bulma.Media(bulma.Image("/a", "a"), node(), nil), bulma.Tile(node()).Class(bulma.Ancestor))
		}, []string{`class="level"`, `level-left`, `level-item`, `media-left`, `media-content`, `is-ancestor tile`}},
		{"grid", func() hx.INode { return bulma.FixedGrid(bulma.Grid(bulma.Cell(node()))).Class(bulma.HasCols(3)) }, []string{`fixed-grid has-3-cols`, `class="grid"`, `class="cell"`}},
		{"footer", func() hx.INode { return bulma.Footer(node()) }, []string{`<footer class="footer">`}},
		{"breadcrumb", func() hx.INode {
			return bulma.Breadcrumb(bulma.Crumb{Label: "Home", Href: "/"}, bulma.Crumb{Label: "Now", Href: "/n", Current: true})
		}, []string{`class="breadcrumb"`, `<li class="is-active"><a`, `aria-current="page"`}},
		{"card", func() hx.INode {
			return bulma.Card(bulma.CardHeader("Title"), bulma.CardImage(bulma.Image("/a", "a")), bulma.CardContent(node()),
				bulma.CardFooter(bulma.CardFooterItem("Save", "", nop), bulma.CardFooterItem("Go", "/g", nil)))
		}, []string{`class="card"`, `card-header-title">Title`, `card-image`, `card-content`, `card-footer-item`}},
		{"dropdown open", func() hx.INode {
			return bulma.Dropdown(flag, "Menu", bulma.DropdownItem("One", "/1", nil), bulma.DropdownDivider())
		}, []string{`dropdown is-active`, `dropdown-trigger`, `dropdown-content`, `dropdown-item`, `dropdown-divider`, `aria-expanded="true"`}},
		{"dropdown closed", func() hx.INode { return bulma.Dropdown(off, "Menu") }, []string{`class="dropdown"`, `aria-expanded="false"`}},
		{"menu", func() hx.INode {
			return bulma.Menu(bulma.MenuLabel("L"), bulma.MenuList(bulma.MenuLink("A", "/a", true, bulma.MenuList(bulma.MenuLink("B", "/b", false)))))
		}, []string{`class="menu"`, `menu-label`, `menu-list`, `<a class="is-active" href="/a">A</a><ul class="menu-list">`}},
		{"message", func() hx.INode { return bulma.Message("Hey", nop, node()).Class(bulma.Danger) }, []string{`is-danger message`, `message-header`, `class="delete"`, `message-body`}},
		{"message plain", func() hx.INode { return bulma.Message("", nil, node()) }, []string{`message-body`}},
		{"modal open", func() hx.INode { return bulma.Modal(flag, nop, node) }, []string{`is-active modal`, `modal-background`, `modal-content`, `modal-close`}},
		{"modal closed", func() hx.INode { return hx.Div().Body(bulma.Modal(off, nop, node)) }, []string{`<div></div>`}},
		{"modal card", func() hx.INode { return bulma.ModalCard(flag, "Sure?", nop, node, node) }, []string{`modal-card-head`, `modal-card-title">Sure?`, `modal-card-body`, `modal-card-foot`}},
		{"navbar", func() hx.INode {
			return bulma.Navbar(
				bulma.NavbarBrand(bulma.NavbarItem("Logo", "/"), bulma.NavbarBurger(flag)),
				bulma.NavbarMenu(flag, bulma.NavbarStart(bulma.NavbarItem("A", "/a").Class(bulma.Active), bulma.NavbarDropdown("More", bulma.NavbarItem("B", "/b"), bulma.NavbarDivider())),
					bulma.NavbarEnd(bulma.NavbarBox(bulma.Button("x")))))
		}, []string{`class="navbar"`, `navbar-brand`, `is-active navbar-burger`, `is-active navbar-menu`, `navbar-start`, `navbar-end`, `has-dropdown is-hoverable navbar-item`, `navbar-divider`}},
		{"pagination", func() hx.INode { return bulma.Pagination(hx.Signal(5), hx.Signal(10), func(int) {}) }, []string{`class="pagination"`, `pagination-previous`, `pagination-next`, `pagination-list`, `is-current`, `pagination-ellipsis`}},
		{"pagination first", func() hx.INode { return bulma.Pagination(hx.Signal(1), hx.Signal(1), func(int) {}) }, []string{`pagination-previous" disabled`}},
		{"panel", func() hx.INode {
			return bulma.Panel("Repos", bulma.PanelTabs(hx.A().Text("all")), bulma.PanelBlock(node()), bulma.PanelLink("fa fa-book", "bulma", "/b"))
		}, []string{`class="panel"`, `panel-heading">Repos`, `panel-tabs`, `panel-block`, `panel-icon`}},
		{"tabs", func() hx.INode {
			return hx.Div().Body(bulma.Tabs(tab, bulma.Tab{ID: "a", Label: "A"}, bulma.Tab{ID: "b", Label: "B", Icon: "fa fa-b"}).Class(bulma.Boxed),
				bulma.TabPanel(tab, "a", func() hx.INode { return hx.P().Text("PANEL-A") }),
				bulma.TabPanel(tab, "b", func() hx.INode { return hx.P().Text("PANEL-B") }))
		}, []string{`class="is-boxed tabs"`, `<li class="is-active" role="presentation"><a role="tab"`, `PANEL-B`}},
		{"helpers", func() hx.INode {
			return hx.Div().Class(bulma.TextColor(bulma.Primary), bulma.BackgroundColor(bulma.Light), bulma.Spacing("m", "t", 3), bulma.SpacingAuto("x"), bulma.On(bulma.Span(6), "tablet"))
		}, []string{`has-background-light has-text-primary is-6-tablet mt-3 mx-auto`}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			html := render(c.build)
			for _, w := range c.want {
				if !strings.Contains(html, w) {
					t.Errorf("missing %q in\n%s", w, html)
				}
			}
			if strings.Contains(html, "TabPanel") {
				t.Error("leak")
			}
		})
	}
	if html := render(func() hx.INode { return hx.Div().Body(bulma.TabPanel(tab, "a", node)) }); strings.Contains(html, "<p>") {
		t.Errorf("TabPanel shows an inactive tab: %s", html)
	}
}
