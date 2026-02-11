package resource

import (
	domainInterfaces "{{.ModulePath}}/interfaces"
	domainHelloWorldRepo "{{.ModulePath}}/repositories/start/hello_world"
	"{{.ModulePath}}/usecases"
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
