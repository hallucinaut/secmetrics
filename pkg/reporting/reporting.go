// Package reporting provides security metrics reporting.
package reporting

import (
	"fmt"
	"time"
)

// ReportFormat represents a report format.
type ReportFormat string

const (
	FormatJSON    ReportFormat = "json"
	FormatYAML    ReportFormat = "yaml"
	FormatMarkdown ReportFormat = "markdown"
	FormatHTML    ReportFormat = "html"
	FormatCSV     ReportFormat = "csv"
)

// Report represents a security metrics report.
type Report struct {
	ID            string
	Title         string
	Description   string
	Format        ReportFormat
	CreatedAt     time.Time
	Metrics       []MetricData
	KPIS          []KPIData
	Executive     ExecutiveSummary
	Technical     TechnicalSummary
	Recommendations []string
}

// MetricData represents metric data for reporting.
type MetricData struct {
	Name     string
	Type     string
	Value    float64
	Target   float64
	Status   string
	Trend    string
	Timestamp time.Time
}

// KPIData represents KPI data for reporting.
type KPIData struct {
	Key        string
	Name       string
	Value      float64
	Target     float64
	Status     string
	Trend      string
	Unit       string
	Category   string
}

// ExecutiveSummary provides executive-level summary.
type ExecutiveSummary struct {
	OverallHealth      string
	ComplianceScore    float64
	RiskScore          float64
	TopConcerns        []string
	TopAchievements    []string
	Recommendations    []string
	ActionItems        []string
}

// TechnicalSummary provides technical-level summary.
type TechnicalSummary struct {
	MetricsCovered     int
	KPIsTracked        int
	AlertsActive       int
	IncidentsLastMonth int
	VulnerabilitiesOpen int
	ComplianceStatus   string
	DetectionRate      float64
	ResponseTime       float64
}

// ReportGenerator generates security metrics reports.
type ReportGenerator struct {
	reports []Report
}

// NewReportGenerator creates a new report generator.
func NewReportGenerator() *ReportGenerator {
	return &ReportGenerator{
		reports: make([]Report, 0),
	}
}

// GenerateReport generates a security metrics report.
func (g *ReportGenerator) GenerateReport(title, description string, format ReportFormat) *Report {
	report := &Report{
		ID:          "rpt-" + time.Now().Format("20060102150405"),
		Title:       title,
		Description: description,
		Format:      format,
		CreatedAt:   time.Now(),
		Metrics:     make([]MetricData, 0),
		KPIS:        make([]KPIData, 0),
		Executive:   ExecutiveSummary{},
		Technical:   TechnicalSummary{},
	}

	g.reports = append(g.reports, *report)
	return report
}

// AddMetric adds metric data to report.
func (g *ReportGenerator) AddMetric(reportID string, metric MetricData) {
	for i := range g.reports {
		if g.reports[i].ID == reportID {
			g.reports[i].Metrics = append(g.reports[i].Metrics, metric)
			break
		}
	}
}

// AddKPI adds KPI data to report.
func (g *ReportGenerator) AddKPI(reportID string, kpi KPIData) {
	for i := range g.reports {
		if g.reports[i].ID == reportID {
			g.reports[i].KPIS = append(g.reports[i].KPIS, kpi)
			break
		}
	}
}

// SetExecutiveSummary sets executive summary for report.
func (g *ReportGenerator) SetExecutiveSummary(reportID string, summary ExecutiveSummary) {
	for i := range g.reports {
		if g.reports[i].ID == reportID {
			g.reports[i].Executive = summary
			break
		}
	}
}

// SetTechnicalSummary sets technical summary for report.
func (g *ReportGenerator) SetTechnicalSummary(reportID string, summary TechnicalSummary) {
	for i := range g.reports {
		if g.reports[i].ID == reportID {
			g.reports[i].Technical = summary
			break
		}
	}
}

// GetReport retrieves a report by ID.
func (g *ReportGenerator) GetReport(reportID string) *Report {
	for i := range g.reports {
		if g.reports[i].ID == reportID {
			return &g.reports[i]
		}
	}
	return nil
}

// GetReports returns all reports.
func (g *ReportGenerator) GetReports() []Report {
	return g.reports
}

// GenerateExecutiveReport generates executive summary report.
func GenerateExecutiveReport(report *Report) string {
	var reportStr string

	reportStr += "=== Executive Security Metrics Report ===\n\n"
	reportStr += "Report ID: " + report.ID + "\n"
	reportStr += "Title: " + report.Title + "\n"
	reportStr += "Created: " + report.CreatedAt.Format("2006-01-02 15:04:05") + "\n\n"

	// Executive Summary
	reportStr += "Executive Summary\n"
	reportStr += "=================\n\n"
	reportStr += "Overall Health: " + report.Executive.OverallHealth + "\n"
	reportStr += "Compliance Score: " + fmt.Sprintf("%.1f%%", report.Executive.ComplianceScore) + "\n"
	reportStr += "Risk Score: " + fmt.Sprintf("%.1f", report.Executive.RiskScore) + "\n\n"

	if len(report.Executive.TopConcerns) > 0 {
		reportStr += "Top Concerns:\n"
		for i, concern := range report.Executive.TopConcerns {
			reportStr += "  [" + fmt.Sprintf("%d", i+1) + "] " + concern + "\n"
		}
		reportStr += "\n"
	}

	if len(report.Executive.TopAchievements) > 0 {
		reportStr += "Top Achievements:\n"
		for i, achievement := range report.Executive.TopAchievements {
			reportStr += "  [" + fmt.Sprintf("%d", i+1) + "] " + achievement + "\n"
		}
		reportStr += "\n"
	}

	if len(report.Executive.Recommendations) > 0 {
		reportStr += "Recommendations:\n"
		for i, rec := range report.Executive.Recommendations {
			reportStr += "  [" + fmt.Sprintf("%d", i+1) + "] " + rec + "\n"
		}
		reportStr += "\n"
	}

	if len(report.Executive.ActionItems) > 0 {
		reportStr += "Action Items:\n"
		for i, action := range report.Executive.ActionItems {
			reportStr += "  [" + fmt.Sprintf("%d", i+1) + "] " + action + "\n"
		}
	}

	return reportStr
}

// GenerateTechnicalReport generates technical detail report.
func GenerateTechnicalReport(report *Report) string {
	var reportStr string

	reportStr += "=== Technical Security Metrics Report ===\n\n"
	reportStr += "Report ID: " + report.ID + "\n\n"

	// Technical Summary
	reportStr += "Technical Summary\n"
	reportStr += "=================\n\n"
	reportStr += "Metrics Covered: " + fmt.Sprintf("%d", report.Technical.MetricsCovered) + "\n"
	reportStr += "KPIs Tracked: " + fmt.Sprintf("%d", report.Technical.KPIsTracked) + "\n"
	reportStr += "Active Alerts: " + fmt.Sprintf("%d", report.Technical.AlertsActive) + "\n"
	reportStr += "Incidents (Last Month): " + fmt.Sprintf("%d", report.Technical.IncidentsLastMonth) + "\n"
	reportStr += "Open Vulnerabilities: " + fmt.Sprintf("%d", report.Technical.VulnerabilitiesOpen) + "\n"
	reportStr += "Compliance Status: " + report.Technical.ComplianceStatus + "\n"
	reportStr += "Detection Rate: " + fmt.Sprintf("%.1f%%", report.Technical.DetectionRate) + "\n"
	reportStr += "Response Time: " + fmt.Sprintf("%.1f hours", report.Technical.ResponseTime) + "\n\n"

	// Metrics
	if len(report.Metrics) > 0 {
		reportStr += "Security Metrics:\n"
		for i, metric := range report.Metrics {
			reportStr += "  [" + fmt.Sprintf("%d", i+1) + "] " + metric.Name + "\n"
			reportStr += "      Value: " + fmt.Sprintf("%.1f", metric.Value) + " " + metric.Type + "\n"
			reportStr += "      Target: " + fmt.Sprintf("%.1f", metric.Target) + " " + metric.Type + "\n"
			reportStr += "      Status: " + metric.Status + "\n"
			reportStr += "      Trend: " + metric.Trend + "\n\n"
		}
	}

	// KPIs
	if len(report.KPIS) > 0 {
		reportStr += "Key Performance Indicators:\n"
		for i, kpi := range report.KPIS {
			reportStr += "  [" + fmt.Sprintf("%d", i+1) + "] " + kpi.Name + "\n"
			reportStr += "      Value: " + fmt.Sprintf("%.1f", kpi.Value) + " " + kpi.Unit + "\n"
			reportStr += "      Target: " + fmt.Sprintf("%.1f", kpi.Target) + " " + kpi.Unit + "\n"
			reportStr += "      Status: " + kpi.Status + "\n"
			reportStr += "      Trend: " + kpi.Trend + "\n"
			reportStr += "      Category: " + kpi.Category + "\n\n"
		}
	}

	return reportStr
}

// GenerateReport generates report in specified format.
func GenerateReport(report *Report, format ReportFormat) string {
	switch format {
	case FormatMarkdown:
		return GenerateMarkdownReport(report)
	case FormatHTML:
		return GenerateHTMLReport(report)
	case FormatCSV:
		return GenerateCSVReport(report)
	default:
		return GenerateTechnicalReport(report)
	}
}

// GenerateMarkdownReport generates Markdown format report.
func GenerateMarkdownReport(report *Report) string {
	var reportStr string

	reportStr += "# Security Metrics Report\n\n"
	reportStr += "**Report ID:** " + report.ID + "\n\n"
	reportStr += "**Title:** " + report.Title + "\n"
	reportStr += "**Created:** " + report.CreatedAt.Format("2006-01-02 15:04:05") + "\n\n"

	reportStr += "## Executive Summary\n\n"
	reportStr += "| Metric | Value |\n"
	reportStr += "|--------|-------|\n"
	reportStr += "| Overall Health | " + report.Executive.OverallHealth + " |\n"
	reportStr += "| Compliance Score | " + fmt.Sprintf("%.1f%%", report.Executive.ComplianceScore) + " |\n"
	reportStr += "| Risk Score | " + fmt.Sprintf("%.1f", report.Executive.RiskScore) + " |\n\n"

	return reportStr
}

// GenerateHTMLReport generates HTML format report.
func GenerateHTMLReport(report *Report) string {
	var reportStr string

	reportStr = "<!DOCTYPE html>\n<html>\n<head>\n"
	reportStr += "<title>Security Metrics Report - " + report.Title + "</title>\n"
	reportStr += "</head>\n<body>\n"
	reportStr += "<h1>Security Metrics Report</h1>\n"
	reportStr += "<h2>" + report.Title + "</h2>\n"
	reportStr += "<p><strong>Report ID:</strong> " + report.ID + "</p>\n"
	reportStr += "<p><strong>Created:</strong> " + report.CreatedAt.Format("2006-01-02 15:04:05") + "</p>\n"
	reportStr += "</body>\n</html>\n"

	return reportStr
}

// GenerateCSVReport generates CSV format report.
func GenerateCSVReport(report *Report) string {
	var reportStr string

	reportStr += "Metric Name,Value,Target,Status,Trend\n"
	for _, metric := range report.Metrics {
		reportStr += metric.Name + "," + fmt.Sprintf("%.1f", metric.Value) + "," + fmt.Sprintf("%.1f", metric.Target) + "," + metric.Status + "," + metric.Trend + "\n"
	}

	return reportStr
}

// GetCommonMetrics returns common security metrics.
func GetCommonMetrics() []MetricData {
	return []MetricData{
		{
			Name:    "Vulnerabilities Open",
			Type:    "count",
			Value:   45.0,
			Target:  20.0,
			Status:  "ABOVE_TARGET",
			Trend:   "IMPROVING",
		},
		{
			Name:    "Critical Vulnerabilities",
			Type:    "count",
			Value:   3.0,
			Target:  0.0,
			Status:  "ABOVE_TARGET",
			Trend:   "STABLE",
		},
		{
			Name:    "Security Patches Applied",
			Type:    "percentage",
			Value:   92.0,
			Target:  100.0,
			Status:  "BELOW_TARGET",
			Trend:   "IMPROVING",
		},
		{
			Name:    "Security Training Completion",
			Type:    "percentage",
			Value:   87.0,
			Target:  100.0,
			Status:  "BELOW_TARGET",
			Trend:   "IMPROVING",
		},
	}
}

// GetReport returns report.
func GetReport(generator *ReportGenerator, reportID string) *Report {
	return generator.GetReport(reportID)
}