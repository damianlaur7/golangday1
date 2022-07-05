package main

import (
	"ex3/entity"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

var person entity.Person

func main() {
	router := chi.NewRouter()
	router.Route("/v1", func(r chi.Router) {
		router.Post("/person", createPerson)
		router.Get("/person", getPerson)
	})
	router.Route("/v2", func(r chi.Router) {
		router.Post("/person", createPerson2)
		router.Get("/person", getPerson2)
	})
	http.ListenAndServe(":8081", router)
}

func getPerson(w http.ResponseWriter, r *http.Request) {
	fmt.Print(r.Method)
	value := fmt.Sprintf("Person is %v", person)
	fmt.Fprint(w, value)
}

func createPerson(w http.ResponseWriter, r *http.Request) {
	fmt.Print(r.Method)
	person = entity.Person{
		Name:   "Laur",
		ID:     "12344",
		Gender: "Male",
	}
	fmt.Fprint(w, person)
}

func getPerson2(w http.ResponseWriter, r *http.Request) {
	fmt.Print(r.Method)
	value := fmt.Sprintf("Person is %v", person)
	fmt.Fprint(w, value)
}

func createPerson2(w http.ResponseWriter, r *http.Request) {
	fmt.Print(r.Method)
	person = entity.Person{
		Name:   "Ovidiu",
		ID:     "44321",
		Gender: "Male",
	}
	fmt.Fprint(w, person)
}
