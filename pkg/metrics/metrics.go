// Package metrics provides security metrics and KPI tracking.
package metrics

import (
	"fmt"
	"time"
)

// MetricType represents a type of security metric.
type MetricType string

const (
	TypeVulnerability   MetricType = "vulnerability"
	TypeIncident        MetricType = "incident"
	TypeCompliance      MetricType = "compliance"
	TypeDetection       MetricType = "detection"
	TypeResponse        MetricType = "response"
	TypePrevention      MetricType = "prevention"
	TypeTraining        MetricType = "training"
	TypeRisk            MetricType = "risk"
)

// SecurityMetric represents a security metric.
type SecurityMetric struct {
	ID          string
	Name        string
	Type        MetricType
	Value       float64
	Unit        string
	Target      float64
	Status      string
	Timestamp   time.Time
	Description string
	Category    string
}

// KPIKey represents a key performance indicator key.
type KPIKey string

const (
	KPI_MTTR            KPIKey = "mttr"
	KPI_MTTC            KPIKey = "mttc"
	KPI_MTTD            KPIKey = "mttd"
	KPI_Coverage        KPIKey = "coverage"
	KPI_Compliance      KPIKey = "compliance"
	KPI_RemediationRate KPIKey = "remediation_rate"
	KPI_DetectionRate   KPIKey = "detection_rate"
	KPI_ResponseTime    KPIKey = "response_time"
)

// KPI represents a security KPI.
type KPI struct {
	Key           KPIKey
	Name          string
	Description   string
	Value         float64
	Target        float64
	Unit          string
	Status        string
	Trend         string
	LastUpdated   time.Time
	Category      string
}

// MetricsCollector collects security metrics.
type MetricsCollector struct {
	metrics  []SecurityMetric
	kpis     []KPI
	summary  *MetricsSummary
}

// MetricsSummary represents a metrics summary.
type MetricsSummary struct {
	TotalMetrics      int
	TotalKPIS         int
	ComplianceScore   float64
	RiskScore         float64
	OverallHealth     string
	LastUpdated       time.Time
}

// NewMetricsCollector creates a new metrics collector.
func NewMetricsCollector() *MetricsCollector {
	return &MetricsCollector{
		metrics: make([]SecurityMetric, 0),
		kpis:    make([]KPI, 0),
		summary: &MetricsSummary{},
	}
}

// AddMetric adds a security metric.
func (c *MetricsCollector) AddMetric(metric SecurityMetric) {
	metric.Timestamp = time.Now()
	c.metrics = append(c.metrics, metric)
	c.updateSummary()
}

// AddKPI adds a KPI.
func (c *MetricsCollector) AddKPI(kpi KPI) {
	kpi.LastUpdated = time.Now()
	c.kpis = append(c.kpis, kpi)
	c.updateSummary()
}

// GetMetrics returns all metrics.
func (c *MetricsCollector) GetMetrics() []SecurityMetric {
	return c.metrics
}

// GetKPIS returns all KPIs.
func (c *MetricsCollector) GetKPIS() []KPI {
	return c.kpis
}

// GetKPI retrieves a KPI by key.
func (c *MetricsCollector) GetKPI(key KPIKey) *KPI {
	for i := range c.kpis {
		if c.kpis[i].Key == key {
			return &c.kpis[i]
		}
	}
	return nil
}

// GetMetricByType returns metrics by type.
func (c *MetricsCollector) GetMetricByType(metricType MetricType) []SecurityMetric {
	var result []SecurityMetric
	for _, metric := range c.metrics {
		if metric.Type == metricType {
			result = append(result, metric)
		}
	}
	return result
}

// GetComplianceScore calculates compliance score.
func (c *MetricsCollector) GetComplianceScore() float64 {
	var total float64
	var weighted float64

	for _, metric := range c.GetMetricByType(TypeCompliance) {
		total += 1.0
		weighted += metric.Value / metric.Target * 100.0
	}

	if total == 0 {
		return 0.0
	}

	return weighted / total
}

// GetRiskScore calculates risk score.
func (c *MetricsCollector) GetRiskScore() float64 {
	var total float64
	var weighted float64

	for _, metric := range c.GetMetricByType(TypeRisk) {
		total += 1.0
		weighted += metric.Value
	}

	if total == 0 {
		return 0.0
	}

	return weighted / total
}

// CalculateMTTR calculates mean time to respond.
func CalculateMTTR(responseTimes []float64) float64 {
	if len(responseTimes) == 0 {
		return 0.0
	}

	var total float64
	for _, time := range responseTimes {
		total += time
	}

	return total / float64(len(responseTimes))
}

// CalculateMTTD calculates mean time to detect.
func CalculateMTTD(detectionTimes []float64) float64 {
	if len(detectionTimes) == 0 {
		return 0.0
	}

	var total float64
	for _, time := range detectionTimes {
		total += time
	}

	return total / float64(len(detectionTimes))
}

// CalculateMTTC calculates mean time to contain.
func CalculateMTTC(containmentTimes []float64) float64 {
	if len(containmentTimes) == 0 {
		return 0.0
	}

	var total float64
	for _, time := range containmentTimes {
		total += time
	}

	return total / float64(len(containmentTimes))
}

// CalculateCoverage calculates security coverage.
func CalculateCoverage(covered, total int) float64 {
	if total == 0 {
		return 0.0
	}
	return float64(covered) / float64(total) * 100.0
}

// CalculateRemediationRate calculates remediation rate.
func CalculateRemediationRate(remediated, total int) float64 {
	if total == 0 {
		return 0.0
	}
	return float64(remediated) / float64(total) * 100.0
}

// CalculateDetectionRate calculates detection rate.
func CalculateDetectionRate(detected, total int) float64 {
	if total == 0 {
		return 0.0
	}
	return float64(detected) / float64(total) * 100.0
}

// updateSummary updates metrics summary.
func (c *MetricsCollector) updateSummary() {
	c.summary.TotalMetrics = len(c.metrics)
	c.summary.TotalKPIS = len(c.kpis)
	c.summary.ComplianceScore = c.GetComplianceScore()
	c.summary.RiskScore = c.GetRiskScore()
	c.summary.OverallHealth = determineHealth(c.summary.ComplianceScore, c.summary.RiskScore)
	c.summary.LastUpdated = time.Now()
}

// determineHealth determines overall health.
func determineHealth(compliance, risk float64) string {
	if compliance >= 90 && risk <= 30 {
		return "HEALTHY"
	} else if compliance >= 70 && risk <= 50 {
		return "GOOD"
	} else if compliance >= 50 && risk <= 70 {
		return "FAIR"
	}
	return "POOR"
}

// GetSummary returns metrics summary.
func (c *MetricsCollector) GetSummary() *MetricsSummary {
	return c.summary
}

// GenerateReport generates metrics report.
func (c *MetricsCollector) GenerateReport() string {
	var report string

	report += "=== Security Metrics Report ===\n\n"

	// Summary
	summary := c.GetSummary()
	report += "Overall Health: " + summary.OverallHealth + "\n"
	report += "Compliance Score: " + fmt.Sprintf("%.1f%%", summary.ComplianceScore) + "\n"
	report += "Risk Score: " + fmt.Sprintf("%.1f", summary.RiskScore) + "\n"
	report += "Total Metrics: " + fmt.Sprintf("%d", summary.TotalMetrics) + "\n"
	report += "Total KPIs: " + fmt.Sprintf("%d", summary.TotalKPIS) + "\n\n"

	// KPIs
	if len(c.kpis) > 0 {
		report += "Key Performance Indicators:\n"
		for i, kpi := range c.kpis {
			report += "  [" + fmt.Sprintf("%d", i+1) + "] " + kpi.Name + "\n"
			report += "      Value: " + fmt.Sprintf("%.1f", kpi.Value) + " " + kpi.Unit + "\n"
			report += "      Target: " + fmt.Sprintf("%.1f", kpi.Target) + " " + kpi.Unit + "\n"
			report += "      Status: " + kpi.Status + "\n"
			report += "      Trend: " + kpi.Trend + "\n\n"
		}
	}

	// Metrics by type
	report += "Metrics by Type:\n"
	for _, metricType := range []MetricType{TypeVulnerability, TypeIncident, TypeCompliance} {
		metrics := c.GetMetricByType(metricType)
		if len(metrics) > 0 {
			report += "  " + string(metricType) + ": " + fmt.Sprintf("%d", len(metrics)) + " metrics\n"
		}
	}

	return report
}

// GetCommonKPIs returns common security KPIs.
func GetCommonKPIs() []KPI {
	return []KPI{
		{
			Key:           KPI_MTTR,
			Name:          "Mean Time to Respond (MTTR)",
			Description:   "Average time to respond to security incidents",
			Value:         2.5,
			Target:        1.0,
			Unit:          "hours",
			Status:        "BELOW_TARGET",
			Trend:         "IMPROVING",
			Category:      "Response",
		},
		{
			Key:           KPI_MTTC,
			Name:          "Mean Time to Contain (MTTC)",
			Description:   "Average time to contain security incidents",
			Value:         4.0,
			Target:        2.0,
			Unit:          "hours",
			Status:        "BELOW_TARGET",
			Trend:         "STABLE",
			Category:      "Response",
		},
		{
			Key:           KPI_MTTD,
			Name:          "Mean Time to Detect (MTTD)",
			Description:   "Average time to detect security incidents",
			Value:         0.5,
			Target:        0.25,
			Unit:          "hours",
			Status:        "BELOW_TARGET",
			Trend:         "IMPROVING",
			Category:      "Detection",
		},
		{
			Key:           KPI_Coverage,
			Name:          "Security Coverage",
			Description:   "Percentage of assets with security controls",
			Value:         85.0,
			Target:        100.0,
			Unit:          "%",
			Status:        "BELOW_TARGET",
			Trend:         "IMPROVING",
			Category:      "Prevention",
		},
		{
			Key:           KPI_Compliance,
			Name:          "Compliance Score",
			Description:   "Overall compliance with security policies",
			Value:         92.0,
			Target:        100.0,
			Unit:          "%",
			Status:        "BELOW_TARGET",
			Trend:         "STABLE",
			Category:      "Compliance",
		},
		{
			Key:           KPI_RemediationRate,
			Name:          "Vulnerability Remediation Rate",
			Description:   "Percentage of vulnerabilities remediated within SLA",
			Value:         78.0,
			Target:        95.0,
			Unit:          "%",
			Status:        "BELOW_TARGET",
			Trend:         "IMPROVING",
			Category:      "Remediation",
		},
	}
}

// GetKPI returns KPI.
func GetKPI(collector *MetricsCollector, key KPIKey) *KPI {
	return collector.GetKPI(key)
}