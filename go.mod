module grgd

go 1.15

replace github.com/gregod-com/grgd => ./

require (
	github.com/common-nighthawk/go-figure v0.0.0-20200609044655-c4b36f998cf2
	github.com/golang/mock v1.4.4
	github.com/gregod-com/grgd v0.0.0-00010101000000-000000000000
	github.com/gregod-com/grgdplugincontracts v0.3.0
	github.com/sirupsen/logrus v1.6.0
	github.com/urfave/cli/v2 v2.2.0
	golang.org/x/mod v0.3.0
	golang.org/x/sys v0.0.0-20190422165155-953cdadca894
	gopkg.in/yaml.v2 v2.3.0
	gorm.io/driver/sqlite v1.1.3
	gorm.io/gorm v1.20.1
	gotest.tools/v3 v3.0.2
)
