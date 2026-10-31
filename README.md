# hxbulma

Every component of [Bulma](https://bulma.io) 1.x as a Hix node. Only the CSS of
Bulma is used (it has no JavaScript), and the parts that move (modal, dropdown,
navbar burger, tabs) are driven by a `Signal`, so it works with `hx/ssr`, with
`hx/wasm` and when the browser hydrates what the server wrote. It imports only `hx`.

```go
import "github.com/deltegui/hxbulma"

// in <head>: <link rel="stylesheet" href="bulma.CSS">

open := hx.Signal(false)
bulma.Section(bulma.Container(
	bulma.Title(1, "Hello"),
	bulma.Button("Open").Class(bulma.Primary, bulma.Large).OnClick(func(hx.EventContext) { open.Set(true) }),
	bulma.Modal(open, func() { open.Set(false) }, func() hx.INode {
		return bulma.Box(hx.P().Text("A modal is in the tree only while it is open."))
	}),
))
```

A component takes what it needs (text, children, the signal it reads) and
returns an `hx.INode`; **modifiers are classes** and are chained: `.Class(bulma.Info, bulma.Light)`.

| File | Components |
| --- | --- |
| `elements.go` | Block, Box, Content, Button(s), DeleteButton, Icon, Image, Notification, Progress, Skeleton, Table, Tag(s), Title, Subtitle |
| `form.go` | Field, Control, Label, Help, FormField, Addons, Grouped, HorizontalField, Input, TextArea, Select, Checkbox, Radio(s), File |
| `layout.go` | Container, Section, Footer, Hero, Columns/Column, Level, Media, Tile, Grid/FixedGrid/Cell |
| `components.go` | Breadcrumb, Card, Dropdown, Menu, Message, Modal, ModalCard, Navbar, Pagination, Panel, Tabs |
| `bulma.go` | the modifier constants (colors, sizes, states, helpers), `Span`, `Offset`, `On`, `Spacing`, `BindClass` |

`BindClass(node, class, signal)` toggles a class with a signal without rebuilding
the node. Bulma does not ship icons: `Icon("fas fa-home")` takes the class of the font you use.

`OWL_DEMO/feature/bulma` shows all of them in `/bulma`; `make check-bulma` there drives it in Chromium.
