package interfaces

import (
	mStart "{{.ModulePath}}/repositories/start/models"
	mUsecase "{{.ModulePath}}/usecases/models"
)

type Usecase interface {
	GetHelloWorld(params mUsecase.HelloWorldUsecaseRequest) mUsecase.HelloWorldUsecaseResponse
}

type HelloWorldRepository interface {
	GetHelloWorld(params mStart.HelloWorldRequest) mStart.HelloWorldResponse
}
