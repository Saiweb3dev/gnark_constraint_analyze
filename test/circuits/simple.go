package circuits

import (
    "github.com/consensys/gnark/frontend"
)

// SimpleCircuit demonstrates basic operations
type SimpleCircuit struct {
    A frontend.Variable `gnark:",public"`
    B frontend.Variable `gnark:",public"`
    C frontend.Variable `gnark:",public"`
}

// Define implements the gnark Circuit interface
func (circuit *SimpleCircuit) Define(api frontend.API) error {
    // This creates 1 constraint (multiplication)
    temp := api.Mul(circuit.A, circuit.B)
    
    // This creates 0 constraints (addition is linear)
    result := api.Add(temp, circuit.C)
    
    // This creates 0 constraints (assertion with linear combination)
    api.AssertIsEqual(result, 15)
    
    return nil
}

// ComplexCircuit demonstrates more operations
type ComplexCircuit struct {
    X frontend.Variable `gnark:",public"`
    Y frontend.Variable `gnark:",public"`
    Z frontend.Variable `gnark:",public"`
}

func (circuit *ComplexCircuit) Define(api frontend.API) error {
    // Multiple multiplications (each creates 1 constraint)
    x2 := api.Mul(circuit.X, circuit.X)  // x^2
    y2 := api.Mul(circuit.Y, circuit.Y)  // y^2
    
    // Addition (no constraints)
    sum := api.Add(x2, y2)
    
    // Division (creates constraint)
    ratio := api.Div(sum, circuit.Z)
    
    // Boolean assertion (creates constraint)
    isValid := api.Sub(ratio, 1)
    api.AssertIsBoolean(isValid)
    
    return nil
}