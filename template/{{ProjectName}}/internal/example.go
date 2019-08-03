package internal

import (
	"{{ProjectName}}/internal/processor"
	"{{ProjectName}}/internal/repositories"
	"{{ProjectName}}/pkg"
	"{{ProjectName}}/pkg/pubsub"
)

type ExampleProcessor struct {
	log    *pkg.Logger
	db     repositories.Database
	pubsub pubsub.PubSub
	processor.Processor
	exampleChannel chan string
}

func NewExample(log *pkg.Logger, db repositories.Database, pubsub pubsub.PubSub) *ExampleProcessor {
	return &ExampleProcessor{
		log:            log,
		db:             db,
		pubsub:         pubsub,
		Processor:      processor.New(),
		exampleChannel: make(chan string)}
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
	//Write your login here
	g.log.Info("process")
}
