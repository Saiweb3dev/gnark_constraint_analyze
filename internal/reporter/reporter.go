package reporter

import (
    "fmt"
    "sort"
    "strings"
    "gnark_basic_setup/internal/counter"
)

// Reporter handles formatting and outputting constraint analysis results
type Reporter struct {
    verbose bool
}

// NewReporter creates a new analysis reporter
func NewReporter(verbose bool) *Reporter {
    return &Reporter{verbose: verbose}
}

// Report outputs a formatted analysis report
func (r *Reporter) Report(stats *counter.CounterStats) {
    fmt.Printf("\n" + strings.Repeat("=", 60) + "\n")
    fmt.Printf("GNARK CONSTRAINT ANALYSIS REPORT\n")
    fmt.Printf(strings.Repeat("=", 60) + "\n")
    
    fmt.Printf("Circuit Name: %s\n", stats.CircuitName)
    fmt.Printf("Total Constraints: %d\n", stats.TotalConstraints)
    fmt.Printf("Analysis Time: %v\n", stats.ExecutionTime)
    fmt.Printf("\n")
    
    if len(stats.ConstraintsByType) > 0 {
        r.reportByType(stats)
    }
    
    if r.verbose && len(stats.Details) > 0 {
        r.reportDetails(stats)
    }
    
    r.reportSummary(stats)
}

func (r *Reporter) reportByType(stats *counter.CounterStats) {
    fmt.Printf("CONSTRAINTS BY TYPE:\n")
    fmt.Printf(strings.Repeat("-", 40) + "\n")
    
    // Sort by count (descending)
    type typeCount struct {
        constraintType counter.ConstraintType
        count         int
    }
    
    var sorted []typeCount
    for t, count := range stats.ConstraintsByType {
        sorted = append(sorted, typeCount{t, count})
    }
    
    sort.Slice(sorted, func(i, j int) bool {
        return sorted[i].count > sorted[j].count
    })
    
    for _, tc := range sorted {
        typeName := r.getConstraintTypeName(tc.constraintType)
        percentage := float64(tc.count) / float64(stats.TotalConstraints) * 100
        fmt.Printf("%-20s: %6d (%5.1f%%)\n", typeName, tc.count, percentage)
    }
    fmt.Printf("\n")
}

func (r *Reporter) reportDetails(stats *counter.CounterStats) {
    fmt.Printf("DETAILED OPERATIONS:\n")
    fmt.Printf(strings.Repeat("-", 40) + "\n")
    
    for i, detail := range stats.Details {
        fmt.Printf("%3d. %-15s (cost: %d)\n", i+1, detail.Operation, detail.Cost)
    }
    fmt.Printf("\n")
}

func (r *Reporter) reportSummary(stats *counter.CounterStats) {
    fmt.Printf("PERFORMANCE INSIGHTS:\n")
    fmt.Printf(strings.Repeat("-", 40) + "\n")
    
    if stats.TotalConstraints == 0 {
        fmt.Printf("• No constraints detected - circuit may use only linear operations\n")
    } else if stats.TotalConstraints < 100 {
        fmt.Printf("• Low constraint count - efficient circuit\n")
    } else if stats.TotalConstraints < 1000 {
        fmt.Printf("• Moderate constraint count - reasonable complexity\n")
    } else {
        fmt.Printf("• High constraint count - consider optimization\n")
    }
		// Identify most expensive operations
    mulCount := stats.ConstraintsByType[counter.ConstraintMul]
    if mulCount > stats.TotalConstraints/2 {
        fmt.Printf("• Multiplication-heavy circuit - consider reducing Mul operations\n")
    }
    
    fmt.Printf(strings.Repeat("=", 60) + "\n\n")
}

func (r *Reporter) getConstraintTypeName(t counter.ConstraintType) string {
    switch t {
    case counter.ConstraintMul:
        return "Multiplication"
    case counter.ConstraintDiv:
        return "Division"
    case counter.ConstraintInverse:
        return "Inverse"
    case counter.ConstraintCmp:
        return "Comparison"
    case counter.ConstraintSelect:
        return "Select"
    case counter.ConstraintLookup:
        return "Lookup"
    case counter.ConstraintRange:
        return "Range Check"
    case counter.ConstraintHash:
        return "Hash"
    case counter.ConstraintAssertIsEqual:
        return "Assert Equal"
    case counter.ConstraintAssertIsDifferent:
        return "Assert Different"
    case counter.ConstraintAssertIsBoolean:
        return "Assert Boolean"
    case counter.ConstraintAssertIsLessOrEqual:
        return "Assert LessOrEqual"
    default:
        return "Unknown"
    }
}