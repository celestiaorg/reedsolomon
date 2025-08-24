package reedsolomon

// GF16Init initializes the GF(2^16) lookup tables.
// This must be called before using GF16Mul.
func GF16Init() {
	initConstants()
}

// GF16Mul multiplies two GF(2^16) elements
// Uses GF(2^16) with polynomial 0x1002D
// Requires GF16Init() to be called first
func GF16Mul(a, b uint16) uint16 {
	if a == 0 || b == 0 {
		return 0
	}
	
	// Convert to ffe type for internal operations
	affe := ffe(a)
	bffe := ffe(b)
	
	// Use logarithm tables for multiplication
	// result = exp(log(a) + log(b))
	logSum := addMod(logLUT[affe], logLUT[bffe])
	if logSum >= modulus {
		logSum -= modulus
	}
	return uint16(expLUT[logSum])
}