package resource

import (
	domainInterfaces "github.com/bimonugraraga/zinogre/templates/server/interfaces"
	domainHelloWorldRepo "github.com/bimonugraraga/zinogre/templates/server/repositories/start/hello_world"
	"github.com/bimonugraraga/zinogre/templates/server/usecases"
)

// All Repositories
var (
	HelloWorldRepository domainInterfaces.HelloWorldRepository
)

// All Usecase
var (
	Usecase domainInterfaces.Usecase
)

// Init All Repo
func InitNewRepo() {
	InitHelloWorldRepository()
}

func InitHelloWorldRepository() {
	HelloWorldRepository = domainHelloWorldRepo.NewRepository()
}

// Init All Usecase
func InitNewUsecase() {
	Usecase = usecases.NewUsecase(HelloWorldRepository)
}
