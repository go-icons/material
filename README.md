# material-icons

File-type icons from the [Material Icon Theme](https://github.com/material-extensions/vscode-material-icon-theme)
(by Material Extensions, MIT), as embedded SVG documents keyed by file name —
for pure-Go UIs that render their own icons.

```go
import materialicons "github.com/go-widgets/material-icons"

svg := materialicons.Icon("paper.tex") // the Material .tex glyph, as an SVG string
dir := materialicons.Folder()          // the folder glyph
```

`Icon(filename)` matches by exact base name first, then by extension, then falls
back to a generic document. It is a **data package**: it returns SVG strings and
draws nothing. A renderer such as
[go-widgets/toolkit](https://github.com/go-widgets/toolkit)'s `SVGIcon` turns the
SVG into a drawn glyph.

## Licence

The Go code is BSD-3-Clause (`LICENSE`). The embedded Material Icon Theme artwork
is MIT, © Material Extensions (`MATERIAL-LICENSE.md`) — redistributed unmodified.
