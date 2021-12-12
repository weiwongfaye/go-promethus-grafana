package stat

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	GaugeVecApiDuration = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "apiDuration",
		Help: "api latency in ms",
	}, []string{"WsorAPT"})

	GaugeVecApiMethod = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "apiCount",
		Help: "api count",
	}, []string{"method"})

	GaugeVecApiError = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "apiErrorCount",
		Help: "api error count",
	}, []string{"type"})
)

func init() {
	prometheus.MustRegister(GaugeVecApiDuration, GaugeVecApiError, GaugeVecApiMethod)
}
