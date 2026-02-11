package usecases

import (
	domainInterfaces "github.com/bimonugraraga/zinogre/templates/server/interfaces"
	mStart "github.com/bimonugraraga/zinogre/templates/server/repositories/start/models"
	mUsecase "github.com/bimonugraraga/zinogre/templates/server/usecases/models"
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
