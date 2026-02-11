package main

import (
	resource "github.com/bimonugraraga/zinogre/generator/templates/server/dal/resources"
	"github.com/bimonugraraga/zinogre/generator/templates/server/routes"
)

func main() {
	resource.InitNewRepo()
	resource.InitNewUsecase()
	routes.InitFiber()
}
