// Provides basic bulding blocks for advanced console UI
package view

// GetHelpTemplate so if i start typing something here
func GetHelpTemplate() string {
	return `


USAGE:
	{{- if .UsageText}}
		{{.UsageText}}
	{{- else}}
		{{.HelpName}} 
		{{- if .VisibleFlags}}	[global options]
		{{- end}}
		{{- if .Commands}} command [command options]
		{{- end}} 
		{{- if .ArgsUsage}} {{.ArgsUsage}}
		{{- else}} [arguments...]
		{{- end}}
	{{end}}
COMMANDS:
	{{- if .VisibleCommands}}
	{{- range .VisibleCategories}}
	 	{{- if .Name}}
	{{ .Name}}
		{{- end}}
		{{- range .VisibleCommands}}
			{{$names:= join .Names ", " }}
			{{- printf "%-20v" $names }} {{"->  "}}{{.Usage}}
		{{- end}}
		{{end}}
	{{- end}}
	{{if .VisibleFlags}}
	GLOBAL OPTIONS:
		{{- range $index, $option := .VisibleFlags}}
			{{if $index}}
		{{- end}}
		{{- printf "%-20v" $option }}
	{{- end}}
{{end}}`
}
