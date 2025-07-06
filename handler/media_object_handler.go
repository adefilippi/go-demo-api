package handler

import (
	"github.com/lunmy/go-api-core/event"
	coreHandler "github.com/lunmy/go-api-core/handler"
	"github.com/lunmy/go-api-core/repository"
	"github.com/lunmy/go-demo-api/entity"
)

type MediaObjectHandler struct {
	*coreHandler.GenericHandler
}

func NewMediaObjectHandler(repo repository.GenericRepository, dispatcher event.Dispatcher) *MediaObjectHandler {
	return &MediaObjectHandler{
		GenericHandler: coreHandler.NewGenericHandler(&entity.MediaObject{}, repo, dispatcher),
	}
}
