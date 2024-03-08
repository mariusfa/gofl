package circuitbreaker

import (
	"errors"
	"sync"
	"time"
)

type State int

const (
	Closed State = iota
	Open
	HalfOpen
)

type CircuitBreaker struct {
	state                State
	mutex                sync.Mutex
	failureThreshold     int
	recoveryThreshold    int
	testRequestsAllowed  int
	failureCount         int
	halfOpenSuccessCount int
	testRequestCount     int
	resetTimeout         time.Duration
	failureResetInterval time.Duration
	intervalTimestamp    time.Time
}

type Options struct {
	FailureThreshold    int
	RecoveryThreshold   int
	TestRequestsAllowed int
	ResetTimeout        time.Duration
	//
	FailureResetInterval time.Duration
}

const (
	DefaultFailureThreshold     = 5
	DefaultRecoveryThreshold    = 3
	DefaultTestRequestsAllowed  = 3
	DefaultResetTimeout         = 5 * time.Second
	DefaultFailureResetInterval = 60 * time.Second
)

func NewCircuitBreaker(opts Options) *CircuitBreaker {
	if opts.FailureThreshold == 0 {
		opts.FailureThreshold = DefaultFailureThreshold
	}
	if opts.RecoveryThreshold == 0 {
		opts.RecoveryThreshold = DefaultRecoveryThreshold
	}
	if opts.TestRequestsAllowed == 0 {
		opts.TestRequestsAllowed = DefaultTestRequestsAllowed
	}
	if opts.ResetTimeout == 0 {
		opts.ResetTimeout = DefaultResetTimeout
	}
	if opts.FailureResetInterval == 0 {
		opts.FailureResetInterval = DefaultFailureResetInterval
	}

	return &CircuitBreaker{
		state:                Closed,
		failureThreshold:     opts.FailureThreshold,
		recoveryThreshold:    opts.RecoveryThreshold,
		testRequestsAllowed:  opts.TestRequestsAllowed,
		resetTimeout:         opts.ResetTimeout,
		failureResetInterval: opts.FailureResetInterval,
		intervalTimestamp:    time.Now(),
	}
}

func (cb *CircuitBreaker) setToOpen() {
	cb.state = Open
	// Reset counters
	cb.failureCount = 0
	cb.halfOpenSuccessCount = 0
	cb.testRequestCount = 0
	time.AfterFunc(cb.resetTimeout, func() {
		cb.mutex.Lock()
		defer cb.mutex.Unlock()
		cb.state = HalfOpen
	})
}

func (cb *CircuitBreaker) setToClosed() {
	cb.state = Closed
	// Reset counters
	cb.failureCount = 0
	cb.halfOpenSuccessCount = 0
	cb.testRequestCount = 0
}

func (cb *CircuitBreaker) Execute(action func() error) error {
	cb.mutex.Lock()

	if cb.state == Open {
		cb.mutex.Unlock()
		return errors.New("circuit breaker is open")
	}

	if cb.state == HalfOpen {
		if cb.testRequestCount >= cb.testRequestsAllowed {
			cb.mutex.Unlock()
			return errors.New("circuit breaker half open test requests exhausted")
		}
		cb.testRequestCount++
	}

	if time.Since(cb.intervalTimestamp) > cb.failureResetInterval && cb.failureCount > 0 {
		cb.failureCount = 0
		cb.intervalTimestamp = time.Now()
	}

	cb.mutex.Unlock()

	err := action()

	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	if err != nil && cb.state == HalfOpen {
		cb.intervalTimestamp = time.Now()
		cb.setToOpen()
		return err
	}

	if err != nil {
		cb.failureCount++
		println(cb.failureCount)
		if cb.failureCount >= cb.failureThreshold {
			cb.setToOpen()
		}
		return err
	}

	if cb.state == HalfOpen {
		cb.halfOpenSuccessCount++
		cb.testRequestCount--

		if cb.halfOpenSuccessCount >= cb.recoveryThreshold {
			cb.setToClosed()
		}
	}

	return nil
}