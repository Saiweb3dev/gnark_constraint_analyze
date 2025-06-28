package counter

import (
    "fmt"
    "time"
    "github.com/consensys/gnark/frontend"
    "github.com/consensys/gnark/frontend/cs/r1cs"
    "github.com/consensys/gnark-crypto/ecc"
)

// CountingEngine manages the constraint counting process
type CountingEngine struct {
    circuitName string
    verbose     bool
}

// NewCountingEngine creates a new constraint counting engine
func NewCountingEngine(circuitName string, verbose bool) *CountingEngine {
    return &CountingEngine{
        circuitName: circuitName,
        verbose:     verbose,
    }
}

// CountConstraints analyzes a circuit and returns constraint statistics
func (e *CountingEngine) CountConstraints(circuit frontend.Circuit) (*CounterStats, error) {
    startTime := time.Now()
    
    if e.verbose {
        fmt.Printf("Starting constraint analysis for circuit: %s\n", e.circuitName)
    }
    
    // Create R1CS compiler
    r1cs, err := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, circuit)
    if err != nil {
        return nil, fmt.Errorf("failed to compile circuit: %w", err)
    }
    
    // Extract constraint information from compiled circuit
    stats := &CounterStats{
        TotalConstraints:  r1cs.GetNbConstraints(),
        ConstraintsByType: make(map[ConstraintType]int),
        ExecutionTime:    time.Since(startTime),
        CircuitName:      e.circuitName,
        Details:          make([]ConstraintInfo, 0),
    }
    
    if e.verbose {
        fmt.Printf("Circuit compiled successfully in %v\n", stats.ExecutionTime)
        fmt.Printf("Total constraints: %d\n", stats.TotalConstraints)
    }
    
    return stats, nil
}

// Need to go through all the API provided and create their interfaces with all the function in separate files
// CountConstraintsWithWrapper analyzes a circuit using the API wrapper approach
// func (e *CountingEngine) CountConstraintsWithWrapper(circuitFactory func(api frontend.API) frontend.Circuit) (*CounterStats, error) {
//     startTime := time.Now()
    
//     // Create a mock API for analysis
//     mockAPI := &MockAPI{}
//     counterAPI := NewCounterAPI(mockAPI, e.circuitName)
    
//     // Execute circuit with counting wrapper
//     circuit := circuitFactory(counterAPI)
    
//     // Compile to get accurate constraint count
//     r1cs, err := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, circuit)
//     if err != nil {
//         return nil, fmt.Errorf("failed to compile circuit: %w", err)
//     }
    
//     stats := counterAPI.GetStats()
//     stats.ExecutionTime = time.Since(startTime)
    
//     // Update with actual constraint count from compilation
//     actualConstraints := r1cs.GetNbConstraints()
//     if actualConstraints != stats.TotalConstraints {
//         if e.verbose {
//             fmt.Printf("Note: Wrapper counted %d operations, compiler generated %d constraints\n", 
//                       stats.TotalConstraints, actualConstraints)
//         }
//         stats.TotalConstraints = actualConstraints
//     }
    
//     return stats, nil
// }

// // MockAPI provides a minimal implementation for analysis
// type MockAPI struct{}

// func (m *MockAPI) Add(i1, i2 frontend.Variable, inputs ...frontend.Variable) frontend.Variable { return 0 }
// func (m *MockAPI) Mul(i1, i2 frontend.Variable, inputs ...frontend.Variable) frontend.Variable { return 0 }
// func (m *MockAPI) Sub(i1, i2 frontend.Variable, inputs ...frontend.Variable) frontend.Variable { return 0 }
// func (m *MockAPI) Div(i1, i2 frontend.Variable) frontend.Variable { return 0 }
// func (m *MockAPI) Inverse(i1 frontend.Variable) frontend.Variable { return 0 }
// func (m *MockAPI) Neg(i1 frontend.Variable) frontend.Variable { return 0 }
// func (m *MockAPI) Select(b frontend.Variable, i1, i2 frontend.Variable) frontend.Variable { return 0 }
// func (m *MockAPI) Lookup2(t frontend.Variable, t0, t1 frontend.Variable) frontend.Variable { return 0 }
// func (m *MockAPI) AssertIsEqual(i1, i2 frontend.Variable) {}
// func (m *MockAPI) AssertIsDifferent(i1, i2 frontend.Variable) {}
// func (m *MockAPI) AssertIsBoolean(i1 frontend.Variable) {}
// func (m *MockAPI) AssertIsLessOrEqual(v frontend.Variable, bound frontend.Variable) {}
// func (m *MockAPI) AssertIsCrumb(i1 frontend.Variable) {}
// func (m *MockAPI) And(i1, i2 frontend.Variable) frontend.Variable { return 0 }