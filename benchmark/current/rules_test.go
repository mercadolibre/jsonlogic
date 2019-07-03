package main

import "testing"

var result Users

func BenchmarkRulesExecution10(b *testing.B) {
	// attribution to avoid the compiler from eliminating the function call
	result = ExecuteRule(10)
}

func BenchmarkRulesExecution100(b *testing.B) {
	// attribution to avoid the compiler from eliminating the function call
	result = ExecuteRule(100)
}

func BenchmarkRulesExecution1000(b *testing.B) {
	// attribution to avoid the compiler from eliminating the function call
	result = ExecuteRule(1000)
}

func BenchmarkRulesExecution10000(b *testing.B) {
	// attribution to avoid the compiler from eliminating the function call
	result = ExecuteRule(10000)
}
