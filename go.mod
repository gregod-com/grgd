module github.com/gregod-com/grgd

go 1.19

replace github.com/gregod-com/grgd/interfaces => ../grgd/interfaces

require (
	github.com/golang/mock v1.6.0
	github.com/sirupsen/logrus v1.9.0
	github.com/stretchr/testify v1.8.0
	github.com/urfave/cli/v2 v2.20.3
	golang.org/x/mod v0.6.0
	golang.org/x/sys v0.1.0
	golang.org/x/term v0.1.0
	gopkg.in/yaml.v3 v3.0.1
	gorm.io/driver/sqlite v1.4.3
	gorm.io/gorm v1.24.0
)

require (
	github.com/cpuguy83/go-md2man/v2 v2.0.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/mattn/go-sqlite3 v1.14.15 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/xrash/smetrics v0.0.0-20201216005158-039620a65673 // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
)
