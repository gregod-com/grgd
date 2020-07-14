package templates

import (
	"bytes"
	"html/template"
)

// Description ...
func Description(app interface{}, selector string) string {
	var tpl bytes.Buffer
	t, err := template.New("foo").Parse(descriptions[selector])
	if err != nil {
		return "Problem rendering the template"
	}

	err = t.Execute(&tpl, app)
	if err != nil {
		return err.Error()
	}

	return tpl.String()
}

var descriptions = map[string]string{
	"init": `
	Welcome to the {{.Name}} {{.Version}} 🍄 🍄 🍄
	
	It looks like you just unlocked your first command. 🤗 🎉 🎉 🎉
	Sadly you are not going to use this one as often as the other ones.
	But still it is an important one.

	When ever you are ready, lets start setting up the {{.Name}} by defining
	the base path for your projects:
	`,
	"up": `
	The 'up' command allows you to start a single workload in your stack or even the whole stack at once. All services that are currently active can be started with the command:
	iam up [workload name] (i.e iam up database)
	This does not nessesarily mean that the workload is actually started as a container but that 'a' workload is made available to your stack via it's DNS name. You can for example start a workload like a database and just wire up the connection to
	an external database hosted as a process on your local machine, on a nearby dev server, or even with tunneling or ingress routing on a remote kubernetes cluster. If you omitt the workload name the cli assumes you want to start all defined workloads. 
	If some or all of them are already running, the command is ignored for them. If you need to restart a service have a look at iam restart.
	`,
	"down": `
	placeholder desription
	`,
	"restart": `
	placeholder desription
	`,
	"logs": `
	placeholder desription
	`,
	"config": `
	The Config is here yo
	`,
	"config-yaml": `
	The yaml subcommand in config
	`,
}
