package cli

import "fmt"

var HelpTemplate = `Usage:
  {{.Usage}}
Description:
  {{.Description}}
Options:
  {.Option}
`
var VersionTemplate = `{{.Name}} (go-coreutils) {{.Version}}`

