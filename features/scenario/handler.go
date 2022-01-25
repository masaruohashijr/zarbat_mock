package scenario

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ensureCors(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Request-Headers", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Add("Access-Control-Allow-Methods", "DELETE, PUT, POST, GET, OPTIONS")
	w.Header().Add("Content-Type", "application/json,text/plain")
	return w
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET ALL")
	scenarios := GetScenarios()
	response, _ := json.Marshal(scenarios)
	fmt.Println(string(response))
	ensureCors(w)
	fmt.Fprintf(w, string(response))
}

func Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	println("id = ", id)
	scenario := GetScenario(id)
	response, _ := json.Marshal(scenario)
	fmt.Println(string(response))
	ensureCors(w)
	fmt.Fprintf(w, string(response))
}

func Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DELETE")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	scenario := DeleteScenario(id)
	response, _ := json.Marshal(scenario)
	ensureCors(w)
	fmt.Fprintf(w, string(response))
}

func Put(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PUT")
	var scenario Scenario
	_ = json.NewDecoder(r.Body).Decode(&scenario)
	updatedScenario := UpdateScenario(scenario)
	response, _ := json.Marshal(updatedScenario)
	fmt.Println(string(response))
	ensureCors(w)
	fmt.Fprintf(w, string(response))
}

func Post(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST")
	var scenario Scenario
	_ = json.NewDecoder(r.Body).Decode(&scenario)
	newScenario := AddScenario(scenario)
	response, _ := json.Marshal(newScenario)
	fmt.Println(string(response))
	ensureCors(w)
	fmt.Fprintf(w, string(response))
}

func sendErr(w http.ResponseWriter, code int, message string) {
	resp, _ := json.Marshal(map[string]string{"error": message})
	http.Error(w, string(resp), code)
}

func Teste(w http.ResponseWriter, r *http.Request) {
	println("BATEU")
}
