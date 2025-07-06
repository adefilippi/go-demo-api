package handler

import (
	"github.com/lunmy/go-api-core/event"
	coreHandler "github.com/lunmy/go-api-core/handler"
	"github.com/lunmy/go-api-core/repository"
	"github.com/lunmy/go-demo-api/entity"
)

type ModelHandler struct {
	*coreHandler.GenericHandler
}

func NewModelHandler(repo repository.GenericRepository, dispatcher event.Dispatcher) *ModelHandler {
	return &ModelHandler{
		GenericHandler: coreHandler.NewGenericHandler(&entity.Model{}, repo, dispatcher),
	}
}
