// Package bulma wraps Bulma 1.x in Hix nodes: every component of the framework
// is a function that returns an hx.INode with the right tags and classes.
//
// It uses only Bulma's CSS. Bulma has no JavaScript, and what a page usually
// does with it by hand (open a modal, toggle the burger of the navbar, switch
// a tab) is here a Signal, so the same code runs on the server (hx/ssr), in the
// browser (hx/wasm) and when the browser hydrates what the server wrote. It
// imports only hx, so it works with both engines.
//
// A component takes what it needs to exist (its text, its children, the signal
// it reads) and returns a node. Modifiers are classes, so they are chained like
// any other class, and the package has a constant for each of them:
//
//	bulma.Button("Save").Class(bulma.Primary, bulma.Large)
//	bulma.Notification(hx.P().Text("Saved")).Class(bulma.Success, bulma.Light)
//
// Bulma itself is a stylesheet that the page loads:
//
//	<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bulma@1.0.4/css/bulma.min.css">
//
// (the constant CSS has that address).
package bulma

import (
	"strconv"
	"strings"

	"github.com/deltegui/hx"
)

// CSS is the stylesheet of the version of Bulma that this package is written for.
const CSS = "https://cdn.jsdelivr.net/npm/bulma@1.0.4/css/bulma.min.css"

// Modifiers. A modifier is the class that Bulma defines for it, so it goes in
// Class(). Colors.
const (
	White   = "is-white"
	Black   = "is-black"
	Light   = "is-light"
	Dark    = "is-dark"
	Primary = "is-primary"
	Link    = "is-link"
	Info    = "is-info"
	Success = "is-success"
	Warning = "is-warning"
	Danger  = "is-danger"
	Text    = "is-text"
	Ghost   = "is-ghost"
)

// Sizes.
const (
	Small  = "is-small"
	Normal = "is-normal"
	Medium = "is-medium"
	Large  = "is-large"
)

// State and shape.
const (
	Active     = "is-active"
	Hovered    = "is-hovered"
	Focused    = "is-focused"
	Loading    = "is-loading"
	Outlined   = "is-outlined"
	Inverted   = "is-inverted"
	Rounded    = "is-rounded"
	Fullwidth  = "is-fullwidth"
	Responsive = "is-responsive"
	Static     = "is-static"
	Selected   = "is-selected"
	Delete     = "delete"
)

// Layout modifiers (columns, tiles, levels, heroes, navbars, tabs, cards...).
const (
	Centered             = "is-centered"
	Right                = "is-right"
	Left                 = "is-left"
	Boxed                = "is-boxed"
	Toggle               = "is-toggle"
	Multiline            = "is-multiline"
	Mobile               = "is-mobile"
	Desktop              = "is-desktop"
	Gapless              = "is-gapless"
	VCentered            = "is-vcentered"
	Fixed                = "is-fixed-top"
	FixedBottom          = "is-fixed-bottom"
	Transparent          = "is-transparent"
	Spaced               = "is-spaced"
	Narrow               = "is-narrow"
	Fluid                = "is-fluid"
	Widescreen           = "is-widescreen"
	Fullhd               = "is-fullhd"
	Halfheight           = "is-halfheight"
	Fullheight           = "is-fullheight"
	FullheightWithNavbar = "is-fullheight-with-navbar"
	Bold                 = "is-bold"
	Hoverable            = "is-hoverable"
	Up                   = "is-up"
	Horizontal           = "is-horizontal"
	Vertical             = "is-vertical"
	Ancestor             = "is-ancestor"
	Parent               = "is-parent"
	Child                = "is-child"
	Striped              = "is-striped"
	Bordered             = "is-bordered"
)

// Column widths: Span(6), Offset(2), or one of these.
const (
	ThreeQuarters = "is-three-quarters"
	TwoThirds     = "is-two-thirds"
	Half          = "is-half"
	OneThird      = "is-one-third"
	OneQuarter    = "is-one-quarter"
	OneFifth      = "is-one-fifth"
	Full          = "is-full"
)

// Ratios of an Image.
const (
	Square    = "is-square"
	Ratio1x1  = "is-1by1"
	Ratio5x4  = "is-5by4"
	Ratio4x3  = "is-4by3"
	Ratio3x2  = "is-3by2"
	Ratio5x3  = "is-5by3"
	Ratio16x9 = "is-16by9"
	Ratio2x1  = "is-2by1"
	Ratio3x1  = "is-3by1"
	Ratio4x5  = "is-4by5"
	Ratio3x4  = "is-3by4"
	Ratio2x3  = "is-2by3"
	Ratio3x5  = "is-3by5"
	Ratio9x16 = "is-9by16"
	Ratio1x2  = "is-1by2"
	Ratio1x3  = "is-1by3"
)

// Span is the class of a column that takes n of the 12 parts: Span(6).
func Span(n int) string { return "is-" + strconv.Itoa(n) }

// Offset is the class of a column that skips n of the 12 parts.
func Offset(n int) string { return "is-offset-" + strconv.Itoa(n) }

// On adds the modifier to a class only on the given breakpoint: On(Span(6),
// "tablet") is "is-6-tablet". The breakpoints are mobile, tablet, desktop,
// widescreen and fullhd.
func On(modifier, breakpoint string) string { return modifier + "-" + breakpoint }

// BindClass adds the class to the node while the signal is true, and removes it
// otherwise. It is how a component reacts to a Signal without being rebuilt:
//
//	bulma.BindClass(bulma.Button("Save"), bulma.Loading, saving)
func BindClass(n hx.INode, class string, on hx.Gettable[bool]) hx.INode {
	return n.Effect(func() { n.ClassIf(class, on.Get()) })
}

// Has is a "has-" helper: Has("text-primary"), Has("background-light").
func Has(helper string) string { return "has-" + helper }

// Helpers of color. TextColor(Primary) is "has-text-primary" and
// BackgroundColor(Light) is "has-background-light".
func TextColor(color string) string { return "has-text-" + strings.TrimPrefix(color, "is-") }
func BackgroundColor(color string) string {
	return "has-background-" + strings.TrimPrefix(color, "is-")
}

// Helpers of typography and alignment.
const (
	TextCentered       = "has-text-centered"
	TextJustified      = "has-text-justified"
	TextLeft           = "has-text-left"
	TextRight          = "has-text-right"
	TextWeightBold     = "has-text-weight-bold"
	TextWeightNormal   = "has-text-weight-normal"
	TextWeightLight    = "has-text-weight-light"
	TextWeightSemibold = "has-text-weight-semibold"
	Italic             = "is-italic"
	Underlined         = "is-underlined"
	Uppercase          = "is-uppercase"
	Lowercase          = "is-lowercase"
	Capitalized        = "is-capitalized"
	FamilyMonospace    = "is-family-monospace"
	FamilySansSerif    = "is-family-sans-serif"
	FamilySerif        = "is-family-serif"
	FamilyCode         = "is-family-code"
	Clearfix           = "is-clearfix"
	PulledLeft         = "is-pulled-left"
	PulledRight        = "is-pulled-right"
	Overlay            = "is-overlay"
	Clipped            = "is-clipped"
	Radiusless         = "is-radiusless"
	Shadowless         = "is-shadowless"
	Unselectable       = "is-unselectable"
	Clickable          = "is-clickable"
	Relative           = "is-relative"
	Marginless         = "is-marginless"
	Paddingless        = "is-paddingless"
	Hidden             = "is-hidden"
	SrOnly             = "is-sr-only"
	Invisible          = "is-invisible"
	Flex               = "is-flex"
	InlineBlock        = "is-inline-block"
	InlineFlex         = "is-inline-flex"
	Inline             = "is-inline"
)

// Flexbox helpers.
const (
	FlexDirectionRow           = "is-flex-direction-row"
	FlexDirectionColumn        = "is-flex-direction-column"
	FlexWrapWrap               = "is-flex-wrap-wrap"
	FlexWrapNowrap             = "is-flex-wrap-nowrap"
	JustifyContentCenter       = "is-justify-content-center"
	JustifyContentFlexStart    = "is-justify-content-flex-start"
	JustifyContentFlexEnd      = "is-justify-content-flex-end"
	JustifyContentSpaceBetween = "is-justify-content-space-between"
	JustifyContentSpaceAround  = "is-justify-content-space-around"
	AlignItemsCenter           = "is-align-items-center"
	AlignItemsFlexStart        = "is-align-items-flex-start"
	AlignItemsFlexEnd          = "is-align-items-flex-end"
	AlignItemsStretch          = "is-align-items-stretch"
	FlexGrow1                  = "is-flex-grow-1"
	FlexShrink0                = "is-flex-shrink-0"
)

// Spacing is a margin or padding helper: Spacing("m", "t", 3) is "mt-3",
// Spacing("p", "", 4) is "p-4". kind is "m" or "p", side is "", "t", "b", "l",
// "r", "x" or "y", and n goes from 0 to 6 (or "auto" for a margin: use
// SpacingAuto).
func Spacing(kind, side string, n int) string { return kind + side + "-" + strconv.Itoa(n) }

// SpacingAuto is a margin helper with auto: SpacingAuto("x") is "mx-auto".
func SpacingAuto(side string) string { return "m" + side + "-auto" }

// Option is an item of a Select or a group of Radios: what it sends and what
// the user reads.
type Option struct{ Value, Label string }

// Plain makes options whose label is their value.
func Plain(values ...string) []Option {
	options := make([]Option, len(values))
	for i, v := range values {
		options[i] = Option{v, v}
	}
	return options
}
