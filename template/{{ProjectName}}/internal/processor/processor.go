package processor

type Processor interface {
	Worker()
	RunPool(p Processor, size int)
}

type processor struct{}

func New() Processor {
	return processor{}
}

func (processor) RunPool(p Processor, size int) {
	for i := 0; i < size; i++ {
		go p.Worker()
	}
}
func (processor) Worker() {
}
