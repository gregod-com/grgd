module github.com/gregod-com/grgd

go 1.13

replace (
	github.com/gregod-com/animaterm => ../animaterm
	github.com/gregod-com/grgd/actions => ./actions
	github.com/gregod-com/grgd/pluginindex => ./pluginindex
	github.com/gregod-com/grgd/templates => ./templates
	github.com/gregod-com/grgd/ui => ./ui

	github.com/gregod-com/implementations => ../implementations
	github.com/gregod-com/interfaces => ../interfaces
)

require (
	github.com/a8m/envsubst v1.1.0 // indirect
	github.com/buger/goterm v0.0.0-20200322175922-2f3e71b85129
	github.com/common-nighthawk/go-figure v0.0.0-20190529165535-67e0ed34491a
	github.com/coreos/go-semver v0.3.0
	github.com/docker/docker v0.7.3-0.20190327010347-be7ac8be2ae0
	github.com/gregod-com/animaterm v0.0.0-20200519172119-d659eadb52f4
	github.com/gregod-com/implementations v0.0.0-20200519205411-e73c133b2d7b
	github.com/gregod-com/interfaces v0.0.0-20200519181200-1f0c4d25c791
	github.com/urfave/cli/v2 v2.2.0
	golang.org/x/mobile v0.0.0-20200329125638-4c31acba0007 // indirect
	gopkg.in/yaml.v2 v2.3.0
)
