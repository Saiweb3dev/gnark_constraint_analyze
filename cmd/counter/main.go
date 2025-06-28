package main

import (
    "flag"
    "fmt"
    "log"
    
    "gnark_basic_setup/internal/counter"
    "gnark_basic_setup/internal/reporter"
    "gnark_basic_setup/test/circuits"
)

func main() {
    verbose := flag.Bool("verbose", false, "Enable verbose output")
    circuitType := flag.String("circuit", "simple", "Circuit type to analyze (simple, complex)")
    flag.Parse()
    
    fmt.Println("Gnark Constraint Counter MVP")
    fmt.Println("============================")
    
    engine := counter.NewCountingEngine(*circuitType, *verbose)
    rep := reporter.NewReporter(*verbose)
    
    var stats *counter.CounterStats
    var err error
    
    switch *circuitType {
    case "simple":
        circuit := &circuits.SimpleCircuit{}
        stats, err = engine.CountConstraints(circuit)
    case "complex":
        circuit := &circuits.ComplexCircuit{}
        stats, err = engine.CountConstraints(circuit)
    default:
        log.Fatalf("Unknown circuit type: %s", *circuitType)
    }
    
    if err != nil {
        log.Fatalf("Analysis failed: %v", err)
    }
    
    rep.Report(stats)
}