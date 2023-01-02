package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"strconv"
	"time"
)

type MonitoringMiddleware struct {
	requestCount    *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
	router          *gin.Engine
	listenAddress   string
	MetricsPath     string
}

func NewMonitoringMiddleware(systemName string, metricPath string) *MonitoringMiddleware {
	m := &MonitoringMiddleware{
		MetricsPath: metricPath,
	}

	m.registerMetrics(systemName)
	return m
}

func (m *MonitoringMiddleware) Use(e *gin.Engine) {
	e.GET(m.MetricsPath, gin.WrapH(promhttp.Handler()))
	e.Use(m.HandlerFunc())
}

func (m *MonitoringMiddleware) registerMetrics(subsystem string) {

	m.requestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: subsystem,
			Name:      "request_total",
			Help:      "Number of requests by url, method and status code",
		},
		[]string{"code", "method", "handler", "host", "url"},
	)
	if err := prometheus.Register(m.requestCount); err != nil {
		fmt.Print("request_total could not be registered in Prometheus")
	}
	m.requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Subsystem: subsystem,
			Name:      "request_duration_seconds",
			Help:      "Request lantencies in seconds",
		},
		[]string{"code", "method", "url"},
	)
	if err := prometheus.Register(m.requestDuration); err != nil {
		fmt.Print("request_total could not be registered in Prometheus")
	}
}

func (m *MonitoringMiddleware) HandlerFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == m.MetricsPath {
			c.Next()
			return
		}
		start := time.Now()
		c.Next()
		url := c.Request.URL.Path

		statusCode := strconv.Itoa(c.Writer.Status())
		elapsed := float64(time.Since(start)) / float64(time.Second)

		m.requestDuration.WithLabelValues(statusCode, c.Request.Method, url).Observe(elapsed)
		m.requestCount.WithLabelValues(statusCode, c.Request.Method, c.HandlerName(), c.Request.Host, url).Inc()
	}
}
