package counter

import (
    "github.com/consensys/gnark/frontend"
)

// NewCounterAPI creates a new wrapper around gnark's API
func NewCounterAPI(api frontend.API, circuitName string) *CounterAPI {
    return &CounterAPI{
        api: api,
        stats: &CounterStats{
            ConstraintsByType: make(map[ConstraintType]int),
            CircuitName:      circuitName,
            Details:          make([]ConstraintInfo, 0),
        },
        depth: 0,
    }
}

// GetStats returns the current constraint statistics
func (c *CounterAPI) GetStats() *CounterStats {
    return c.stats
}

// recordConstraint logs a constraint and updates statistics
func (c *CounterAPI) recordConstraint(constraintType ConstraintType, operation string, cost int) {
    c.stats.TotalConstraints++
    c.stats.ConstraintsByType[constraintType]++
    
    info := ConstraintInfo{
        Type:        constraintType,
        Description: operation,
        Operation:   operation,
        Cost:        cost,
    }
    c.stats.Details = append(c.stats.Details, info)
}

// Implement frontend.API interface with counting

// Mul wraps multiplication with constraint counting
func (c *CounterAPI) Mul(i1, i2 frontend.Variable, inputs ...frontend.Variable) frontend.Variable {
    c.recordConstraint(ConstraintMul, "Mul", 1)
    return c.api.Mul(i1, i2, inputs...)
}

// Div wraps division with constraint counting
func (c *CounterAPI) Div(i1, i2 frontend.Variable) frontend.Variable {
    c.recordConstraint(ConstraintDiv, "Div", 1)
    return c.api.Div(i1, i2)
}

// Inverse wraps inverse with constraint counting
func (c *CounterAPI) Inverse(i1 frontend.Variable) frontend.Variable {
    c.recordConstraint(ConstraintInverse, "Inverse", 1)
    return c.api.Inverse(i1)
}

// Select wraps conditional select with constraint counting
func (c *CounterAPI) Select(b frontend.Variable, i1, i2 frontend.Variable) frontend.Variable {
    c.recordConstraint(ConstraintSelect, "Select", 1)
    return c.api.Select(b, i1, i2)
}

// Lookup wraps lookup operations
func (c *CounterAPI) Lookup2(t frontend.Variable, t0, t1 frontend.Variable) frontend.Variable {
    c.recordConstraint(ConstraintLookup, "Lookup2", 1)
    return c.api.Lookup2(t, t0, t1,t,t0,t1)
}

// AssertIsEqual wraps equality assertions
func (c *CounterAPI) AssertIsEqual(i1, i2 frontend.Variable) {
    c.recordConstraint(ConstraintAssertIsEqual, "AssertIsEqual", 0)
    c.api.AssertIsEqual(i1, i2)
}

// AssertIsDifferent wraps inequality assertions
func (c *CounterAPI) AssertIsDifferent(i1, i2 frontend.Variable) {
    c.recordConstraint(ConstraintAssertIsDifferent, "AssertIsDifferent", 1)
    c.api.AssertIsDifferent(i1, i2)
}

// AssertIsBoolean wraps boolean assertions
func (c *CounterAPI) AssertIsBoolean(i1 frontend.Variable) {
    c.recordConstraint(ConstraintAssertIsBoolean, "AssertIsBoolean", 1)
    c.api.AssertIsBoolean(i1)
}

// AssertIsLessOrEqual wraps comparison assertions
func (c *CounterAPI) AssertIsLessOrEqual(v frontend.Variable, bound frontend.Variable) {
    c.recordConstraint(ConstraintAssertIsLessOrEqual, "AssertIsLessOrEqual", 1)
    c.api.AssertIsLessOrEqual(v, bound)
}


// Delegated methods (no constraints generated)

// Add delegates addition (no constraints)
func (c *CounterAPI) Add(i1, i2 frontend.Variable, inputs ...frontend.Variable) frontend.Variable {
    return c.api.Add(i1, i2, inputs...)
}

// Sub delegates subtraction (no constraints)
func (c *CounterAPI) Sub(i1, i2 frontend.Variable, inputs ...frontend.Variable) frontend.Variable {
    return c.api.Sub(i1, i2, inputs...)
}

// Neg delegates negation (no constraints)
func (c *CounterAPI) Neg(i1 frontend.Variable) frontend.Variable {
    return c.api.Neg(i1)
}