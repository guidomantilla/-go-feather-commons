package resilience

type FailSaver interface {
	Run(work func() error) error
}
