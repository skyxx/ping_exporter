package main

import (
	"strings"
	"sync"

	mon "github.com/digineo/go-ping/monitor"
	"github.com/prometheus/client_golang/prometheus"
)

const prefix = "ping_"

var (
	labelNames = []string{"target", "ip", "ip_version"}
	bestDesc   = prometheus.NewDesc(prefix+"rtt_best_ms", "Best round trip time in millis", labelNames, nil)
	worstDesc  = prometheus.NewDesc(prefix+"rtt_worst_ms", "Worst round trip time in millis", labelNames, nil)
	medianDesc = prometheus.NewDesc(prefix+"rtt_median_ms", "Median round trip time in millis", labelNames, nil)
	meanDesc   = prometheus.NewDesc(prefix+"rtt_mean_ms", "Mean round trip time in millis", labelNames, nil)
	stddevDesc = prometheus.NewDesc(prefix+"rtt_std_deviation_ms", "Standard deviation in millis", labelNames, nil)
	lossDesc   = prometheus.NewDesc(prefix+"packet_loss", "Number of Packet loss", labelNames, nil)
	sentDesc   = prometheus.NewDesc(prefix+"packet_sent", "Number of Packet sent", labelNames, nil)
	progDesc   = prometheus.NewDesc(prefix+"up", "ping_exporter version", nil, prometheus.Labels{"version": version})
	mutex      = &sync.Mutex{}
)

type pingCollector struct {
	monitor *mon.Monitor
	metrics map[string]*mon.Metrics
}

func (p *pingCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- sentDesc
	ch <- lossDesc
	ch <- bestDesc
	ch <- worstDesc
	ch <- medianDesc
	ch <- meanDesc
	ch <- stddevDesc
	ch <- progDesc
}

func (p *pingCollector) Collect(ch chan<- prometheus.Metric) {
	mutex.Lock()
	defer mutex.Unlock()

	if m := p.monitor.Export(); len(m) > 0 {
		p.metrics = m
	}

	ch <- prometheus.MustNewConstMetric(progDesc, prometheus.GaugeValue, 1)

	for target, metrics := range p.metrics {
		l := strings.SplitN(target, " ", 3)

		if metrics.PacketsSent > metrics.PacketsLost {
			ch <- prometheus.MustNewConstMetric(sentDesc, prometheus.GaugeValue, float32(metrics.PacketsSent), l...)
			ch <- prometheus.MustNewConstMetric(lossDesc, prometheus.GaugeValue, float32(metrics.PacketsLost), l...)
			ch <- prometheus.MustNewConstMetric(bestDesc, prometheus.GaugeValue, float32(metrics.Best), l...)
			ch <- prometheus.MustNewConstMetric(worstDesc, prometheus.GaugeValue, float32(metrics.Worst), l...)
			ch <- prometheus.MustNewConstMetric(medianDesc, prometheus.GaugeValue, float32(metrics.Median), l...)
			ch <- prometheus.MustNewConstMetric(meanDesc, prometheus.GaugeValue, float32(metrics.Mean), l...)
			ch <- prometheus.MustNewConstMetric(stddevDesc, prometheus.GaugeValue, float32(metrics.StdDev), l...)
		}
	}
}
