package handler

import (
	"gin_instrumentation/stat"
	"time"

	"github.com/gin-gonic/gin"
)

func MwPromethusHandler(c *gin.Context) {
	start := time.Now()
	method := c.Request.Method

	stat.GaugeVecApiMethod.WithLabelValues(method).Inc()
	c.Next()
	end := time.Now()
	d := end.Sub(start) / time.Millisecond
	stat.GaugeVecApiDuration.WithLabelValues(method).Set(float64(d))
}
