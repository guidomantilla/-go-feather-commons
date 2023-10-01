package resilience

import (
	"errors"
	"fmt"

	"github.com/eapache/go-resiliency/breaker"
)

func Run(failsafe FailSaver, target func() error) error {

	if err := failsafe.Run(target); err != nil {
		if errors.Is(err, breaker.ErrBreakerOpen) {
			return fmt.Errorf("target function wasn't run because the breaker was open %s", breaker.ErrBreakerOpen.Error())
		}

		return errors.New("target function failed to run: " + err.Error())
	}

	return nil
}
