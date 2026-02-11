package main

import (
	resource "github.com/bimonugraraga/zinogre/dal/resources"
	"github.com/bimonugraraga/zinogre/routes"
)

func main() {
	resource.InitNewRepo()
	resource.InitNewUsecase()
	routes.InitFiber()
}
