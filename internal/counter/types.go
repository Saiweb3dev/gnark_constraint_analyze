package counter

import (
    "time"
    "github.com/consensys/gnark/frontend"
)

// ConstraintType represents different types of constraints
type ConstraintType int

const (
    ConstraintMul ConstraintType = iota
    ConstraintDiv
    ConstraintInverse
    ConstraintCmp
    ConstraintSelect
    ConstraintLookup
    ConstraintRange
    ConstraintHash
    ConstraintAssertIsEqual
    ConstraintAssertIsDifferent
    ConstraintAssertIsBoolean
    ConstraintAssertIsLessOrEqual
)

// ConstraintInfo holds detailed information about a constraint
type ConstraintInfo struct {
    Type        ConstraintType
    Description string
    LineNumber  int
    Operation   string
    Cost        int // Relative cost weight
}

// CounterStats aggregates constraint statistics
type CounterStats struct {
    TotalConstraints   int
    ConstraintsByType  map[ConstraintType]int
    ExecutionTime     time.Duration
    CircuitName       string
    Details           []ConstraintInfo
}

// CounterAPI wraps the gnark frontend.API to intercept and count operations
type CounterAPI struct {
    api    frontend.API
    stats  *CounterStats
    depth  int // Track call depth for analysis
}