package stat

import (
	"log"
	"purpleschool-go/advanced/pkg/event"
)

type StatServiceDeps struct {
	StatRepo *StatRepository
	EventBus *event.EventBus
}

type StatService struct {
	StatRepo *StatRepository
	EventBus *event.EventBus
}

func NewStatService(deps *StatServiceDeps) *StatService {
	return &StatService{
		StatRepo: deps.StatRepo,
		EventBus: deps.EventBus,
	}
}

func (service *StatService) AddClick() {
	for msg := range service.EventBus.Subscribe() {
		if msg.Type == event.EventLinkVisited {
			linkId, ok := msg.Data.(uint)

			if !ok {
				log.Fatalln("Invalid EventLinkVisited data", msg.Data)
				continue
			}

			service.StatRepo.AddClick(linkId)
		}
	}
}
