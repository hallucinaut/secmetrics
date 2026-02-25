# secmetrics - Security Metrics & KPI Dashboard

[![Go](https://img.shields.io/badge/Go-1.21-blue)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-green)](LICENSE)

**Track and measure security program effectiveness with comprehensive metrics and KPIs.**

Monitor security performance, measure control effectiveness, and track security program health.

## 🚀 Features

- **Security KPIs**: Track key security performance indicators
- **Metrics Collection**: Collect and aggregate security metrics
- **Health Monitoring**: Monitor overall security health status
- **Report Generation**: Generate executive and technical reports
- **Compliance Tracking**: Track compliance scores and status
- **Trend Analysis**: Monitor security metrics trends over time

## 📦 Installation

### Build from Source

```bash
git clone https://github.com/hallucinaut/secmetrics.git
cd secmetrics
go build -o secmetrics ./cmd/secmetrics
sudo mv secmetrics /usr/local/bin/
```

### Install via Go

```bash
go install github.com/hallucinaut/secmetrics/cmd/secmetrics@latest
```

## 🎯 Usage

### Collect Metrics

```bash
# Collect security metrics
secmetrics collect
```

### Show KPIs

```bash
# Display security KPIs
secmetrics kpis
```

### Generate Report

```bash
# Generate executive report
secmetrics report executive

# Generate technical report
secmetrics report technical

# Generate markdown report
secmetrics report markdown
```

### Show Summary

```bash
# Show metrics summary
secmetrics summary
```

### Check Health

```bash
# Check security health status
secmetrics health
```

### Programmatic Usage

```go
package main

import (
    "fmt"
    "github.com/hallucinaut/secmetrics/pkg/metrics"
    "github.com/hallucinaut/secmetrics/pkg/reporting"
)

func main() {
    // Create metrics collector
    collector := metrics.NewMetricsCollector()
    
    // Add common KPIs
    commonKPIS := metrics.GetCommonKPIs()
    for _, kpi := range commonKPIS {
        collector.AddKPI(kpi)
    }
    
    // Get KPIs
    kpis := collector.GetKPIS()
    fmt.Printf("KPIs: %d\n", len(kpis))
    
    // Get compliance score
    complianceScore := collector.GetComplianceScore()
    fmt.Printf("Compliance Score: %.1f%%\n", complianceScore)
    
    // Get risk score
    riskScore := collector.GetRiskScore()
    fmt.Printf("Risk Score: %.1f\n", riskScore)
    
    // Generate report
    generator := reporting.NewReportGenerator()
    report := generator.GenerateReport("Security Report", "Metrics report", reporting.FormatMarkdown)
    
    // Set executive summary
    report.Executive = reporting.ExecutiveSummary{
        OverallHealth: "HEALTHY",
        ComplianceScore: complianceScore,
        RiskScore: riskScore,
    }
    
    fmt.Println(reporting.GenerateExecutiveReport(report))
}
```

## 📊 Key Performance Indicators

### Response Metrics

| KPI | Description | Target | Status |
|-----|-------------|--------|--------|
| MTTR | Mean Time to Respond | 1.0 hours | Monitoring |
| MTTC | Mean Time to Contain | 2.0 hours | Monitoring |
| MTTD | Mean Time to Detect | 0.25 hours | Monitoring |

### Prevention Metrics

| KPI | Description | Target | Status |
|-----|-------------|--------|--------|
| Coverage | Security Coverage | 100% | Monitoring |
| Patching | Security Patches Applied | 100% | Monitoring |

### Compliance Metrics

| KPI | Description | Target | Status |
|-----|-------------|--------|--------|
| Compliance | Compliance Score | 100% | Monitoring |

### Remediation Metrics

| KPI | Description | Target | Status |
|-----|-------------|--------|--------|
| Remediation | Vulnerability Remediation Rate | 95% | Monitoring |

## 🧪 Calculated Metrics

### MTTR (Mean Time to Respond)
Average time to respond to security incidents.

```go
responseTimes := []float64{2.0, 3.0, 1.5, 2.5}
mttr := metrics.CalculateMTTR(responseTimes)
fmt.Printf("MTTR: %.1f hours\n", mttr)
```

### MTTD (Mean Time to Detect)
Average time to detect security incidents.

```go
detectionTimes := []float64{0.5, 0.3, 0.4, 0.6}
mttd := metrics.CalculateMTTD(detectionTimes)
fmt.Printf("MTTD: %.1f hours\n", mttd)
```

### MTTC (Mean Time to Contain)
Average time to contain security incidents.

```go
containmentTimes := []float64{3.0, 4.0, 2.5, 3.5}
mttc := metrics.CalculateMTTC(containmentTimes)
fmt.Printf("MTTC: %.1f hours\n", mttc)
```

### Coverage
Percentage of assets with security controls.

```go
covered := 85
total := 100
coverage := metrics.CalculateCoverage(covered, total)
fmt.Printf("Coverage: %.1f%%\n", coverage)
```

### Remediation Rate
Percentage of vulnerabilities remediated within SLA.

```go
remediated := 95
total := 100
rate := metrics.CalculateRemediationRate(remediated, total)
fmt.Printf("Remediation Rate: %.1f%%\n", rate)
```

## 📊 Report Types

### Executive Report
High-level summary for leadership and board members.

- Overall security health
- Top concerns and achievements
- Strategic recommendations
- Key metrics and trends

### Technical Report
Detailed metrics and technical analysis.

- Security metrics details
- KPI performance analysis
- Technical recommendations
- Trend analysis

### Markdown Report
Formatted report in Markdown format.

- Clean formatting
- Tables and charts
- Easy to share
- Version control friendly

## 🏥 Health Status

| Health | Compliance | Risk | Action |
|--------|------------|------|--------|
| HEALTHY | ≥90% | ≤30% | Maintain posture |
| GOOD | ≥70% | ≤50% | Address concerns |
| FAIR | ≥50% | ≤70% | Improve security |
| POOR | <50% | >70% | Immediate action |

## 🧪 Testing

```bash
# Run all tests
go test ./...

# Run with coverage
go test -cover ./...

# Run specific test
go test -v ./pkg/metrics -run TestCalculateMTTR
```

## 📋 Example Output

```
$ secmetrics kpis

Security KPIs
=============

Key Performance Indicators:

[1] Mean Time to Respond (MTTR)
    Value: 2.5 hours
    Target: 1.0 hours
    Status: BELOW_TARGET
    Trend: IMPROVING
    Category: Response

[2] Security Coverage
    Value: 85.0%
    Target: 100.0%
    Status: BELOW_TARGET
    Trend: IMPROVING
    Category: Prevention

[3] Compliance Score
    Value: 92.0%
    Target: 100.0%
    Status: BELOW_TARGET
    Trend: STABLE
    Category: Compliance

Summary:
  Compliance Score: 92.0%
  Risk Score: 45.5
  Overall Health: GOOD
```

## 🏗️ Architecture

```
secmetrics/
├── cmd/
│   └── secmetrics/
│       └── main.go          # CLI entry point
├── pkg/
│   ├── metrics/
│   │   ├── metrics.go      # Metrics collection
│   │   └── metrics_test.go # Unit tests
│   └── reporting/
│       ├── reporting.go    # Report generation
│       └── reporting_test.go # Unit tests
└── README.md
```

## 🔒 Security Use Cases

- **Security Program Management**: Track and measure security program effectiveness
- **Board Reporting**: Generate executive reports for leadership
- **Compliance Audits**: Document compliance metrics and status
- **Security Operations**: Monitor security operations performance
- **Risk Management**: Track security risk metrics
- **Continuous Improvement**: Identify areas for security improvement

## 🛡️ Best Practices

1. **Set realistic targets** based on industry benchmarks
2. **Track metrics consistently** over time
3. **Review KPIs regularly** in security meetings
4. **Share reports** with stakeholders
5. **Use metrics** to drive security improvements
6. **Benchmark against peers** and industry standards

## 📄 License

MIT License

## 🙏 Acknowledgments

- Security metrics community
- CISO forums
- Industry working groups

## 🔗 Resources

- [NIST SP 800-53](https://csrc.nist.gov/publications/detail/sp/800-53/rev-5/final)
- [CIS Controls](https://www.cisecurity.org/controls)
- [OWASP Security Metrics](https://owasp.org/www-project-web-security-testing-guide/latest/4-Web_Application_Security_Testing/00-Web_Application_Security_Testing_Introduction/04-Security_Metrics)
- [SANS Security Metrics](https://www.sans.org/top20-security-controls/)

---

**Built with GPU by [hallucinaut](https://github.com/hallucinaut)**