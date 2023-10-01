package resilience

type NoFailSaver struct {
}

func NewNoFailSaver() *NoFailSaver {
	return &NoFailSaver{}
}

func (failSafe *NoFailSaver) Run(target func() error) error {
	return target()
}
