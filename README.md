# go-basic-example
Golang basics commands, functions, etc.

### Prerequisites

- [Visual Studio Code](https://code.visualstudio.com/) or other IDE (Integrated Development Environment);
- [Golang](https://golang.org/);

### Preparing the environment

&nbsp;&nbsp;&nbsp;&nbsp;The following technologies are critical for running/compiling application sources:

- Access the application directory from the terminal and execute:
  - Install the [gorilla/mux](https://github.com/gorilla/mux):
  > go get -u github.com/gorilla/mux


### Testing basic Go Commands

> go run .\basics-go\main.go


### Testing a simple example of REST web service

> go run .\rest\main.go

  - [GET] - http://localhost:8000/contato
  - [GET] - http://localhost:8000/contato/1
  - [POST] - http://localhost:8000/contato/3
  - [DELETE] - http://localhost:8000/contato/3
