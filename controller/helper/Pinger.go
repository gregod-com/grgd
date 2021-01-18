package helper

import (
	"log"
	"net/http"

	"github.com/gregod-com/grgd/interfaces"
)

// Connection ...
type Connection struct {
	Endpoint string
	TimeOut  int
	Success  bool
}

// Pinger ...
type Pinger struct {
	logger interfaces.ILogger
}

// ProvidePinger ...
func ProvidePinger(logger interfaces.ILogger) interfaces.IPinger {
	pinger := new(Pinger)
	pinger.logger = logger
	return pinger
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
