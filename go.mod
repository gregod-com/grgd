module github.com/gregod-com/grgd

go 1.14

replace (
	github.com/gregod-com/grgdplugins/shared => ../grgdplugins/shared
	github.com/gregod-com/implementations => ../implementations
	github.com/gregod-com/interfaces => ../interfaces
)

require (
	github.com/buger/goterm v0.0.0-20200322175922-2f3e71b85129
	github.com/gregod-com/grgdplugincontracts v0.0.3
	github.com/gregod-com/implementations v0.0.2
	github.com/gregod-com/interfaces v0.0.22
	github.com/kr/pretty v0.1.0 // indirect
	github.com/urfave/cli/v2 v2.2.0
	golang.org/x/mod v0.3.0
	golang.org/x/sys v0.0.0-20200625212154-ddb9806d33ae // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	gopkg.in/yaml.v2 v2.3.0
)
