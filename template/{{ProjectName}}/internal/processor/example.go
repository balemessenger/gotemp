package processor

import (
	"{{ProjectName}}/internal/repositories"
	"{{ProjectName}}/pkg"
	processor2 "{{ProjectName}}/pkg/processor"
	"{{ProjectName}}/pkg/pubsub"
)

type ExampleProcessor struct {
	log    *pkg.Logger
	db     repositories.Database
	pubsub pubsub.PubSub
	processor2.Processor
	exampleChannel chan string
}

func NewExample(log *pkg.Logger, db repositories.Database, pubsub pubsub.PubSub) *ExampleProcessor {
	return &ExampleProcessor{
		log:            log,
		db:             db,
		pubsub:         pubsub,
		Processor:      processor2.New(),
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
