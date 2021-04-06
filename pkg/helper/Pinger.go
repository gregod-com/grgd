package helper

import (
	"log"
	"net/http"
	"reflect"

	"github.com/gregod-com/grgd/interfaces"
)

// Connection ...
type Connection struct {
	Endpoint string
	TimeOut  int
	Success  bool
}

// ProvidePinger ...
func ProvidePinger(logger interfaces.ILogger) interfaces.IPinger {
	pinger := new(Pinger)
	pinger.pkg = reflect.TypeOf(Pinger{}).PkgPath()
	pinger.logger = logger
	pinger.logger.Tracef("provide %T", pinger)
	return pinger
}

// Pinger ...
type Pinger struct {
	logger interfaces.ILogger
	pkg    string
}

// CheckConnections ...
func (p *Pinger) CheckConnections(conns map[string]interface{}) {
	for k := range conns {
		if conn, ok := conns[k].(Connection); ok {
			_, err := http.Get(conn.Endpoint)
			if err != nil {
				log.Fatal(err)
			}
			conn.Success = true
		}
	}
	return
}
