package internal

import (
	"{{ProjectName}}/internal/processor"
	"sync"
)

type ExampleProcessor struct {
	processor.Processor
	exampleChannel chan string
}

var (
	exampleOnce sync.Once
	example     *ExampleProcessor
)

func GetExample() *ExampleProcessor {
	exampleOnce.Do(func() {
		example = NewExample()
	})
	return example
}

func NewExample() *ExampleProcessor {
	return &ExampleProcessor{
		Processor:      processor.New(),
		exampleChannel: make(chan string)}
}

func (g *ExampleProcessor) Initialize(size int) {
	g.RunPool(g.Processor, size)
}

func (g *ExampleProcessor) Tell(envelop string) {
	g.exampleChannel <- envelop
}

func (g *ExampleProcessor) Worker() {
	for {
		envelop := <-g.exampleChannel
		g.process(envelop)
	}
}

func (g *ExampleProcessor) process(envelop string) {
	//Write your login here
}
