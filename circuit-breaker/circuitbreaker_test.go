package circuitbreaker

import (
	"errors"
	"testing"
	"time"
)

// Happy path test
func TestPerformAction(t *testing.T) {
	actionPerformed := false
	action := func() error {
		actionPerformed = true
		return nil
	}

	cb := NewCircuitBreaker(Options{})
	cb.Execute(action)

	if !actionPerformed {
		t.Error("action should have been performed")
	}

	if cb.state != Closed {
		t.Error("state should be closed")
	}

	if cb.failureCount != 0 {
		t.Error("failure count should be 0")
	}
}

// Test that the circuit breaker opens after the failure threshold is reached
func TestOpenCircuit(t *testing.T) {
	action := func() error {
		return errors.New("error")
	}

	cb := NewCircuitBreaker(Options{
		FailureThreshold: 2,
	})

	cb.Execute(action)
	cb.Execute(action)
	cb.Execute(action)

	if cb.state != Open {
		t.Error("state should be open")
	}
}

// Test that failure count is reset after the failure reset interval
func TestFailureCountReset(t *testing.T) {
	action := func() error {
		return errors.New("error")
	}

	cb := NewCircuitBreaker(Options{
		FailureThreshold:     2,
		FailureResetInterval: 1 * time.Microsecond,
	})

	cb.Execute(action)
	cb.Execute(action)
	time.Sleep(2 * time.Microsecond) // waits for rest of failure count
	cb.Execute(action)

	if cb.state != Closed {
		t.Error("state should be closed")
	}

	if cb.failureCount != 1 {
		t.Error("failure count should be 1")
	}
}
