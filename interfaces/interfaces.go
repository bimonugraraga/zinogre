package interfaces

import (
	mStart "github.com/bimonugraraga/zinogre/repositories/start/models"
	mUsecase "github.com/bimonugraraga/zinogre/usecases/models"
)

type Usecase interface {
	GetHelloWorld(params mUsecase.HelloWorldUsecaseRequest) mUsecase.HelloWorldUsecaseResponse
}

type HelloWorldRepository interface {
	GetHelloWorld(params mStart.HelloWorldRequest) mStart.HelloWorldResponse
}
