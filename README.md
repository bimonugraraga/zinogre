# zinogre

Hi guys, This is my small project. It is a Golang Backend server boilerplate based on clean architecture and hexagonal architecture by Uncle Bob (Robert C. Martin).

Spec:
1. This boiler plate using ```https://github.com/gofiber/fiber``` as framework (can chaged after)
2. This boiler plate need minimum ```go 1.22.0```


To set up this boilerplate:
1. Install  ```go install github.com/bimonugraraga/zinogre/cmd/zinogre@v0.0.3```
2. Run Command ```zinogre new {{project-name}}```, project name preferably repo name

To test run:
1. Go to project that just created by ```zinogre new {{project-name}}```
2. run ```go mod tidy```
3. run ```go run .\app\http\http.go``` to start server

Current stable
1. ``` github.com/bimonugraraga/zinogre/cmd/zinogre@v0.0.3```