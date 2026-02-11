package interfaces

import (
	mStart "github.com/bimonugraraga/zinogre/generator/templates/server/repositories/start/models"
	mUsecase "github.com/bimonugraraga/zinogre/generator/templates/server/usecases/models"
)

type Usecase interface {
	GetHelloWorld(params mUsecase.HelloWorldUsecaseRequest) mUsecase.HelloWorldUsecaseResponse
}

type HelloWorldRepository interface {
	GetHelloWorld(params mStart.HelloWorldRequest) mStart.HelloWorldResponse
}
