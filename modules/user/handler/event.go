package handler

import (
	"modular-monolith-libraryApp/modules/user/usecase"
)

type UserEventHandler struct {
	repo usecase.CreateUserInterface
}

func NewUserEventHandler(uc usecase.CreateUserInterface) *UserEventHandler {
	return &UserEventHandler{uc}
}

func (e *UserEventHandler) OnClientRegister(event any) {

}
