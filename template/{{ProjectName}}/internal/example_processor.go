package internal

import (
	"{{ProjectName}}/pkg"
)

type ProcessorRequest struct {
	ID      int32
	Request interface{}
}

type ProcessorResponse struct {
	Response interface{}
	Err      error
}

type ExampleProcessor struct {
	pkg.Processor
	// Inject dependencies here
	exampleChannel chan string
}

func NewExampleProcessor() *ExampleProcessor {
	return &ExampleProcessor{
		exampleChannel: make(chan string),
	}
}

func (g *ExampleProcessor) Start(size int) {
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
	// Write your logic here
}
