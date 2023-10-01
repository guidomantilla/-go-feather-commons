package resilience

type DefaultFailSaver struct {
	breaker FailSaver
	retrier FailSaver
}

type DefaultFailSaverOption func(*DefaultFailSaver)

func WithCircuitBreaker(breaker FailSaver) DefaultFailSaverOption {
	return func(defaultFailSaver *DefaultFailSaver) {
		defaultFailSaver.breaker = breaker
	}
}

func WithRetrier(retrier FailSaver) DefaultFailSaverOption {
	return func(defaultFailSaver *DefaultFailSaver) {
		defaultFailSaver.retrier = retrier
	}
}

func NewDefaultFailSaver(opts ...DefaultFailSaverOption) *DefaultFailSaver {

	failSafe := &DefaultFailSaver{
		breaker: NewNoFailSaver(),
		retrier: NewNoFailSaver(),
	}
	for _, opt := range opts {
		opt(failSafe)
	}

	return failSafe
}

func (failSafe *DefaultFailSaver) Run(target func() error) error {
	return Run(failSafe.retrier, func() error {
		return failSafe.breaker.Run(target)
	})
}
