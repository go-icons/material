// Copyright (c) 2026 the go-widgets authors. All rights reserved.
// Use of this source code is governed by a BSD-3-Clause license that can be
// found in the LICENSE file at the root of this repository.

// Package materialicons serves file-type icons from the Material Icon Theme (by
// Material Extensions, MIT — see MATERIAL-LICENSE.md), as SVG documents keyed by
// file name.
//
// It is a data package: a curated subset of the icons is embedded, and [Icon]
// maps a file name to the best-matching SVG (falling back to a generic
// document), [Folder] returns the folder glyph. A renderer such as
// go-widgets/toolkit's SVGIcon turns the returned SVG into a drawn glyph; this
// package draws nothing itself.
package material

import (
	"embed"
	"path"
	"strings"
)

// Name is the human label a picker shows for this pack.
const Name = "Material"

//go:embed svg/*.svg
var files embed.FS

// byName maps a lower-cased base file name to an icon.
var byName = map[string]string{
	"license":        "document",
	"license.md":     "document",
	"license.txt":    "document",
	"copying":        "document",
	".gitignore":     "git",
	".gitattributes": "git",
	".gitmodules":    "git",
	"go.mod":         "go",
	"go.sum":         "go",
	"package.json":   "json",
	"dockerfile":     "settings",
	"makefile":       "settings",
}

// byExt maps a lower-cased extension (with the dot) to an icon.
var byExt = map[string]string{
	".tex": "tex", ".bib": "tex", ".dtx": "tex", ".ins": "tex",
	".sty": "settings", ".cls": "settings", // LaTeX packages/classes read apart from a .tex document
	".md": "markdown", ".markdown": "markdown",
	".json": "json",
	".yml":  "yaml", ".yaml": "yaml",
	".xml":  "xml",
	".svg":  "image",
	".html": "html", ".htm": "html",
	".css": "css",
	".js":  "javascript", ".mjs": "javascript", ".cjs": "javascript", ".jsx": "javascript",
	".ts": "typescript", ".tsx": "typescript",
	".py": "python",
	".go": "go",
	".rs": "rust",
	".c":  "c", ".h": "c",
	".cpp": "cpp", ".cc": "cpp", ".cxx": "cpp", ".hpp": "cpp", ".hh": "cpp",
	".cs": "csharp",
	".sh": "console", ".bash": "console", ".zsh": "console",
	".pdf": "pdf",
	".png": "image", ".jpg": "image", ".jpeg": "image", ".gif": "image", ".webp": "image", ".bmp": "image", ".eps": "image",
	".lua":  "lua",
	".rb":   "ruby",
	".java": "java",
	".php":  "php",
	".lock": "lock",
	".toml": "settings", ".ini": "settings", ".cfg": "settings", ".conf": "settings",
}

// Icon returns the SVG document for the file named filename (any path — only the
// base name matters), matching by exact name first, then by extension, then a
// generic document.
func Icon(filename string) string {
	base := strings.ToLower(path.Base(filename))
	if ic, ok := byName[base]; ok {
		return read(ic)
	}
	if ic, ok := byExt[strings.ToLower(path.Ext(base))]; ok {
		return read(ic)
	}
	return read("document")
}

// Folder returns the SVG document for a directory.
func Folder() string { return read("folder") }

// read returns the embedded SVG for an icon name, or "" when absent.
func read(name string) string {
	b, err := files.ReadFile("svg/" + name + ".svg")
	if err != nil {
		return ""
	}
	return string(b)
}
