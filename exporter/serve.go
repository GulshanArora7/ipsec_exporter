package exporter

import (
	"net/http"

	"github.com/GulshanArora7/ipsec_exporter/ipsec"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
)

// IPSecConfigFile Variable
var IPSecConfigFile string

// WebListenAddress Variable
var WebListenAddress string

var ipSecConfiguration *ipsec.Configuration

// Serve Function
func Serve() {
	var err error
	ipSecConfiguration, err = ipsec.NewConfiguration(IPSecConfigFile)
	if err != nil {
		log.Fatal(err)
		return
	}
	if !ipSecConfiguration.HasTunnels() {
		log.Warn("Found no configured connections in " + IPSecConfigFile)
	}

	collector := ipsec.NewCollector(ipSecConfiguration)
	prometheus.MustRegister(collector)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`<html>
             <head><title>IStrongswan IPSec Metrics Exporter</title></head>
             <body>
             <h1>Strongswan IPSec Metrics Exporter</h1>
             <p><a href='/metrics'>Metrics</a></p>
             </body>
             </html>`))
	})
	http.Handle("/metrics", promhttp.Handler())

	log.Infoln("Listening on", WebListenAddress)
	err = http.ListenAndServe(WebListenAddress, nil)
	if err != nil {
		log.Fatal(err)
	}
}
