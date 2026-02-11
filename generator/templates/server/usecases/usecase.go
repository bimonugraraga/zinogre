package usecases

import (
	domainInterfaces "{{.ModulePath}}/interfaces"
	mStart "{{.ModulePath}}/repositories/start/models"
	mUsecase "{{.ModulePath}}/usecases/models"
)

type Usecase struct {
	HelloWorldRepo domainInterfaces.HelloWorldRepository
}

func NewUsecase(
	HelloWorldRepo domainInterfaces.HelloWorldRepository,
) *Usecase {
	return &Usecase{
		HelloWorldRepo: HelloWorldRepo,
	}
}

// Usecase only for Business Logic. Business Logic is forbidden in handler and repository
func (u *Usecase) GetHelloWorld(params mUsecase.HelloWorldUsecaseRequest) mUsecase.HelloWorldUsecaseResponse {
	r := u.HelloWorldRepo.GetHelloWorld(mStart.HelloWorldRequest{
		Ctx: params.Ctx,
	})
	return mUsecase.HelloWorldUsecaseResponse{
		Message: r.Message,
	}
}
