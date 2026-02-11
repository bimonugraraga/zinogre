package main

import (
	resource "{{.ModulePath}}/dal/resources"
	"{{.ModulePath}}/routes"
)

func main() {
	resource.InitNewRepo()
	resource.InitNewUsecase()
	routes.InitFiber()
}
