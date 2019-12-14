package templates


// GetHelpTemplate ...
func GetHelpTemplate() string {
	return `
 USAGE:
	{{if .UsageText}}{{.UsageText}}{{else}}{{.HelpName}} {{if .VisibleFlags}}[global options]{{end}}{{if .Commands}} command [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{end}}
	{{if .VisibleCommands}}
 COMMANDS:{{range .VisibleCategories}}{{if .Name}}
 
	{{.Name}}:{{end}}{{range .VisibleCommands}}
	  {{join .Names ", "}}{{"\t\t\t-> "}}{{.Usage}}{{end}}{{end}}{{end}}{{if .VisibleFlags}}
 
 GLOBAL OPTIONS:
	{{range $index, $option := .VisibleFlags}}{{if $index}}
	{{end}}{{$option}}{{end}}{{end}}
 `
}