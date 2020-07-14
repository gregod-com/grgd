module github.com/gregod-com/grgd

go 1.14

replace (
	github.com/gregod-com/grgdplugins/shared => ../grgdplugins/shared
	github.com/gregod-com/implementations => ../implementations
)
require (
	github.com/buger/goterm v0.0.0-20200322175922-2f3e71b85129
	github.com/gregod-com/grgdplugins/shared v0.0.0-00010101000000-000000000000
	github.com/gregod-com/implementations v0.0.1
	github.com/gregod-com/interfaces v0.0.17
	github.com/imdario/mergo v0.3.9 // indirect
	github.com/urfave/cli/v2 v2.2.0
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d // indirect
	golang.org/x/sys v0.0.0-20200625212154-ddb9806d33ae
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e // indirect
	gopkg.in/yaml.v2 v2.3.0
	k8s.io/apimachinery v0.16.11
	k8s.io/client-go v0.16.11
	k8s.io/utils v0.0.0-20200619165400-6e3d28b6ed19 // indirect
)
