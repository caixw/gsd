// SPDX-FileCopyrightText: 2018-2026 caixw
//
// SPDX-License-Identifier: MIT

package main

const tpl = `# {{.Title}}

{{if .Desc}}{{.Desc}}{{end}}

{{range .Objects}}
{{if .Title}}## {{.Title}}{{end}}

{{if .Desc}}{{.Desc}}{{end}}

{{if .Items}}
| JSON | YAML | XML | TOML | {{$.TypeLocale}} | {{$.DescLocale}} |
|------|------|-----|------|------------------|------------------|
{{range .Items -}}
| {{.JSON}} | {{.YAML}} | {{.XML}} | {{.TOML}} | {{.Type}} | {{.Desc}} |
{{end}}

{{end}}

{{- end}}
`
