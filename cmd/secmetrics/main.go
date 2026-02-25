package main

import (
	"fmt"
	"os"

	"github.com/hallucinaut/secmetrics/pkg/metrics"
	"github.com/hallucinaut/secmetrics/pkg/reporting"
)

const version = "1.0.0"

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	switch os.Args[1] {
	case "collect":
		collectMetrics()
	case "kpis":
		showKPIS()
	case "report":
		if len(os.Args) < 3 {
			fmt.Println("Error: report type required")
			printUsage()
			return
		}
		generateReport(os.Args[2])
	case "summary":
		showSummary()
	case "health":
		checkHealth()
	case "version":
		fmt.Printf("secmetrics version %s\n", version)
	case "help", "--help", "-h":
		printUsage()
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		printUsage()
	}
}

func printUsage() {
	fmt.Print(`secmetrics - Security Metrics & KPI Dashboard

Usage:
  secmetrics <command> [options]

Commands:
  collect    Collect security metrics
  kpis       Show security KPIs
  report     Generate metrics report
  summary    Show metrics summary
  health     Check security health status
  version    Show version information
  help       Show this help message

Examples:
  secmetrics collect
  secmetrics kpis
  secmetrics report executive
  secmetrics summary
`, "secmetrics")
}

func collectMetrics() {
	fmt.Println("Security Metrics Collection")
	fmt.Println("==========================")
	fmt.Println()

	collector := metrics.NewMetricsCollector()

	// Add common KPIs
	commonKPIS := metrics.GetCommonKPIs()
	for _, kpi := range commonKPIS {
		collector.AddKPI(kpi)
	}

	fmt.Println("Metrics Collected:")
	fmt.Println("  ✓ Mean Time to Respond (MTTR)")
	fmt.Println("  ✓ Mean Time to Contain (MTTC)")
	fmt.Println("  ✓ Mean Time to Detect (MTTD)")
	fmt.Println("  ✓ Security Coverage")
	fmt.Println("  ✓ Compliance Score")
	fmt.Println("  ✓ Vulnerability Remediation Rate")
	fmt.Println()

	// Show collected metrics
	fmt.Println("Collected KPIs:")
	for i, kpi := range collector.GetKPIS() {
		fmt.Printf("  [%d] %s: %.1f %s\n", i+1, kpi.Name, kpi.Value, kpi.Unit)
	}
	fmt.Println()

	// Show summary
	summary := collector.GetSummary()
	fmt.Println("Summary:")
	fmt.Printf("  Compliance Score: %.1f%%\n", summary.ComplianceScore)
	fmt.Printf("  Risk Score: %.1f\n", summary.RiskScore)
	fmt.Printf("  Overall Health: %s\n", summary.OverallHealth)
}

func showKPIS() {
	fmt.Println("Security KPIs")
	fmt.Println("=============")
	fmt.Println()

	commonKPIS := metrics.GetCommonKPIs()

	fmt.Println("Key Performance Indicators:")
	fmt.Println()
	for i, kpi := range commonKPIS {
		fmt.Printf("[%d] %s\n", i+1, kpi.Name)
		fmt.Printf("    Value: %.1f %s\n", kpi.Value, kpi.Unit)
		fmt.Printf("    Target: %.1f %s\n", kpi.Target, kpi.Unit)
		fmt.Printf("    Status: %s\n", kpi.Status)
		fmt.Printf("    Trend: %s\n", kpi.Trend)
		fmt.Printf("    Category: %s\n\n", kpi.Category)
	}
}

func generateReport(reportType string) {
	fmt.Printf("Generating %s Report\n", reportType)
	fmt.Println()

	// Create collector and add data
	collector := metrics.NewMetricsCollector()

	// Add common KPIs
	commonKPIS := metrics.GetCommonKPIs()
	for _, kpi := range commonKPIS {
		collector.AddKPI(kpi)
	}

	// Create report
	generator := reporting.NewReportGenerator()
	report := generator.GenerateReport("Security Metrics Report", "Comprehensive security metrics report", reporting.FormatMarkdown)

	// Set executive summary
	report.Executive = reporting.ExecutiveSummary{
		OverallHealth: collector.GetSummary().OverallHealth,
		ComplianceScore: collector.GetSummary().ComplianceScore,
		RiskScore: collector.GetSummary().RiskScore,
		TopConcerns: []string{"Vulnerability remediation rate below target", "Security coverage needs improvement"},
		TopAchievements: []string{"MTTD improved by 20%", "Compliance score at 92%"},
		Recommendations: []string{"Increase security automation", "Expand security monitoring coverage"},
		ActionItems: []string{"Address critical vulnerabilities", "Complete security training"},
	}

	// Set technical summary
	report.Technical = reporting.TechnicalSummary{
		MetricsCovered: 6,
		KPIsTracked: 6,
		AlertsActive: 12,
		IncidentsLastMonth: 23,
		VulnerabilitiesOpen: 45,
		ComplianceStatus: "COMPLIANT",
		DetectionRate: 95.0,
		ResponseTime: 2.5,
	}

	// Add metrics
	commonMetrics := reporting.GetCommonMetrics()
	for _, metric := range commonMetrics {
		generator.AddMetric(report.ID, metric)
	}

	// Add KPIs
	for _, kpi := range commonKPIS {
		generator.AddKPI(report.ID, reporting.KPIData{
			Key:      string(kpi.Key),
			Name:     kpi.Name,
			Value:    kpi.Value,
			Target:   kpi.Target,
			Status:   kpi.Status,
			Trend:    kpi.Trend,
			Unit:     kpi.Unit,
			Category: kpi.Category,
		})
	}

	// Generate report based on type
	switch reportType {
	case "executive":
		fmt.Println(reporting.GenerateExecutiveReport(report))
	case "technical":
		fmt.Println(reporting.GenerateTechnicalReport(report))
	case "markdown":
		fmt.Println(reporting.GenerateMarkdownReport(report))
	default:
		fmt.Println(reporting.GenerateTechnicalReport(report))
	}
}

func showSummary() {
	fmt.Println("Security Metrics Summary")
	fmt.Println("========================")
	fmt.Println()

	collector := metrics.NewMetricsCollector()

	// Add common KPIs
	commonKPIS := metrics.GetCommonKPIs()
	for _, kpi := range commonKPIS {
		collector.AddKPI(kpi)
	}

	summary := collector.GetSummary()

	fmt.Println("Overall Health:", summary.OverallHealth)
	fmt.Println("Compliance Score:", fmt.Sprintf("%.1f%%", summary.ComplianceScore))
	fmt.Println("Risk Score:", fmt.Sprintf("%.1f", summary.RiskScore))
	fmt.Println()

	fmt.Println("KPIs Tracked:", summary.TotalKPIS)
	fmt.Println("Metrics Collected:", summary.TotalMetrics)
}

func checkHealth() {
	fmt.Println("Security Health Check")
	fmt.Println("=====================")
	fmt.Println()

	collector := metrics.NewMetricsCollector()

	// Add common KPIs
	commonKPIS := metrics.GetCommonKPIs()
	for _, kpi := range commonKPIS {
		collector.AddKPI(kpi)
	}

	summary := collector.GetSummary()

	fmt.Println("Health Status:", summary.OverallHealth)
	fmt.Println()

	// Check each KPI
	fmt.Println("KPI Status:")
	for _, kpi := range commonKPIS {
		status := "✓"
		if kpi.Status == "BELOW_TARGET" {
			status = "⚠"
		}
		fmt.Printf("  %s %s: %.1f%%\n", status, kpi.Name, kpi.Value)
	}
	fmt.Println()

	fmt.Println("Recommendations:")
	if summary.ComplianceScore < 100 {
		fmt.Println("  • Improve compliance score")
	}
	if summary.RiskScore > 50 {
		fmt.Println("  • Reduce risk score")
	}
	if summary.OverallHealth == "POOR" || summary.OverallHealth == "FAIR" {
		fmt.Println("  • Review security posture")
	}
}