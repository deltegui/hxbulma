package bulma

import (
	"strconv"

	"github.com/deltegui/hx"
)

// Container centers the content with a maximum width. Chain Fluid, Widescreen
// or Fullhd.
func Container(children ...hx.INode) hx.INode { return hx.Div().Class("container").Body(children...) }

// Section is a block of the page with a vertical padding.
func Section(children ...hx.INode) hx.INode { return hx.Section().Class("section").Body(children...) }

// Footer is the footer of the page.
func Footer(children ...hx.INode) hx.INode { return hx.Footer().Class("footer").Body(children...) }

// Hero is a big banner. Put HeroHead, HeroBody and HeroFoot in it, and chain a
// color and a height (Medium, Large, Halfheight, Fullheight).
func Hero(children ...hx.INode) hx.INode { return hx.Section().Class("hero").Body(children...) }

func HeroHead(children ...hx.INode) hx.INode { return hx.Div().Class("hero-head").Body(children...) }
func HeroBody(children ...hx.INode) hx.INode { return hx.Div().Class("hero-body").Body(children...) }
func HeroFoot(children ...hx.INode) hx.INode { return hx.Div().Class("hero-foot").Body(children...) }

// SimpleHero is a Hero with a title and a subtitle in its body.
func SimpleHero(title, subtitle string) hx.INode {
	return Hero(HeroBody(Title(1, title), Subtitle(2, subtitle)))
}

// Columns is a row of Column. Chain Mobile to keep them side by side on a phone,
// Multiline to wrap, Gapless, Centered or VCentered.
func Columns(children ...hx.INode) hx.INode { return hx.Div().Class("columns").Body(children...) }

// Column is a column of Columns. Chain its width: Span(4), Half, Narrow,
// Offset(2), On(Span(6), "tablet").
func Column(children ...hx.INode) hx.INode { return hx.Div().Class("column").Body(children...) }

// Level is a horizontal bar with items on the left and the right (toolbars,
// headings with actions).
func Level(children ...hx.INode) hx.INode { return hx.Div().Class("level").Body(children...) }

func LevelLeft(children ...hx.INode) hx.INode { return hx.Div().Class("level-left").Body(children...) }
func LevelRight(children ...hx.INode) hx.INode {
	return hx.Div().Class("level-right").Body(children...)
}
func LevelItem(children ...hx.INode) hx.INode { return hx.Div().Class("level-item").Body(children...) }

// Media is the object of the "image on the side, text next to it" pattern
// (comments, lists of people). left and right can be nil.
func Media(left, content, right hx.INode) hx.INode {
	var parts []hx.INode
	if left != nil {
		parts = append(parts, hx.El("figure").Class("media-left").Body(left))
	}
	parts = append(parts, hx.Div().Class("media-content").Body(content))
	if right != nil {
		parts = append(parts, hx.Div().Class("media-right").Body(right))
	}
	return hx.El("article").Class("media").BodyList(parts)
}

// Tile is an element of the Metro-like layout. The first one has Ancestor, its
// children have Parent, and their children Child. Chain Vertical and a size.
func Tile(children ...hx.INode) hx.INode { return hx.Div().Class("tile").Body(children...) }

// FixedGrid is the container of a Grid with a fixed number of columns: chain
// "has-3-cols" (HasCols(3)).
func FixedGrid(grid hx.INode) hx.INode { return hx.Div().Class("fixed-grid").Body(grid) }

// HasCols is the class of a FixedGrid with n columns, from 1 to 12.
func HasCols(n int) string { return "has-" + strconv.Itoa(n) + "-cols" }

// Grid is the CSS grid of Bulma 1. Without FixedGrid it makes as many columns as
// fit (SmartGrid); with it, the columns of the FixedGrid.
func Grid(children ...hx.INode) hx.INode { return hx.Div().Class("grid").Body(children...) }

// SmartGrid is a Grid that makes as many columns as fit.
func SmartGrid(children ...hx.INode) hx.INode { return Grid(children...).Class("is-col-min-8") }

// Cell is one cell of a Grid.
func Cell(children ...hx.INode) hx.INode { return hx.Div().Class("cell").Body(children...) }
