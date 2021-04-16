module github.com/gregod-com/grgd

go 1.16

replace github.com/gregod-com/grgd/interfaces => ../grgd/interfaces

require (
	github.com/common-nighthawk/go-figure v0.0.0-20200609044655-c4b36f998cf2
	github.com/golang/mock v1.4.4
	github.com/golang/protobuf v1.4.3
	github.com/rancher/k3d/v4 v4.4.1
	github.com/sirupsen/logrus v1.8.1
	github.com/tj/assert v0.0.3
	github.com/urfave/cli/v2 v2.3.0
	golang.org/x/mod v0.3.0
	golang.org/x/sys v0.0.0-20210309074719-68d13333faf2
	golang.org/x/term v0.0.0-20201210144234-2321bbc49cbf
	google.golang.org/protobuf v1.25.0
	gopkg.in/yaml.v2 v2.4.0
	gorm.io/driver/sqlite v1.1.3
	gorm.io/gorm v1.21.7
)
