package handler

import (
	"modular-monolith-libraryApp/modules/user/usecase"
)

type UserEventHandler struct {
	repo usecase.CreateUser
}

func NewUserEventHandler(uc usecase.CreateUser) *UserEventHandler {
	return &UserEventHandler{uc}
}

func (e *UserEventHandler) OnClientRegister(event any) {

}
