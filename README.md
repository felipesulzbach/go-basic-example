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


### Running the application

Basic Go Commands
> go run .\basics-go\main.go

REST web service

> go run .\rest\main.go

  - Find all contact: [GET] - http://localhost:8080/contact
  - Find contact by ID: [GET] - http://localhost:8080/contact/3
  - Create a new contact: [POST] - http://localhost:8080/contact
    
    ```json
    Body Request:
    {
       "ID": 5,
       "Firstname": "Felipe",
       "Lastname": "Sulzbach",
       "Address": {
           "City": "Seattle",
    	   "State": "Washington"
       }
    }
    ```
  - Delete a contact: [DELETE] - http://localhost:8080/contact/3
