package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"rest-api/entity"
	"rest-api/estrutura"
	"time"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func main() {
	addPerson()
	addConfigServer()
}

func addConfigServer() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	router := mux.NewRouter()
	router.HandleFunc("/contato", getPeople).Methods("GET")
	router.HandleFunc("/contato/{id}", getPerson).Methods("GET")
	router.HandleFunc("/contato/{id}", createPerson).Methods("POST")
	router.HandleFunc("/contato/{id}", deletePerson).Methods("DELETE")
	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":8000", router))

	srv := &http.Server{
		Addr:         "localhost:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	// Execute o nosso servidor em uma goroutine para que ele não bloqueie.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// Aceitaremos shutdowns graciosos quando sairmos via SIGINT (Ctrl + C)
	// SIGKILL, SIGQUIT ou SIGTERM (Ctrl + /) não serão capturados.
	signal.Notify(c, os.Interrupt)

	// Bloqueie até recebermos o nosso sinal.
	<-c

	// Crie um prazo para esperar.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Não bloqueia se não houver conexões, mas, caso contrário, esperará
	// até o prazo limite.
	srv.Shutdown(ctx)
	// Opcionalmente, você pode executar srv.Shutdown em uma goroutine e bloquear em
	// <-ctx.Done () se seu aplicativo deve esperar por outros serviços
	// para finalizar com base no cancelamento do contexto.
	log.Println("shutting down")
	os.Exit(0)
}

func addPerson() {
	people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
	people = append(people, Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})
	people = append(people, Person{ID: "3", Firstname: "Francis", Lastname: "Sunday"})
}

func getPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func getPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func createPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

func deletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(people)
	}
}
