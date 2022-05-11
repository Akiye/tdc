package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"tdc/database"
	"tdc/entities"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	fmt.Println("Go MySQL Test")

	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/tasks")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	log.Println("Starting the HTTP server on port 8090")
	router := mux.NewRouter().StrictSlash(true)
	log.Fatal(http.ListenAndServe(":8085", router))

	//insert, err := db.Exec("INSERT INTO `survive` VALUES ('159-COM-TEST','Identify Combatant and Non-Combatant Personnel & Hybrid Threats','AN');")
	/*if err != nil {
		panic(err.Error())
	}
	defer insert.Close()*/

	/*results, err := db.Query("SELECT * FROM task_group")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var user user
		err = results.Scan(&user.Name)
		if err != nil {
			panic(err.Error())
		}
	}*/

	fmt.Println("Successfully Connected to MySQL database")
}
func createPerson(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var person entities.User
	json.Unmarshal(requestBody, &person)

	database.Connector.Create(person)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(person)
}

func getPersonByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var person entities.User
	database.Connector.First(&person, key)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)
}

func updatePersonByID(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var person entities.User
	json.Unmarshal(requestBody, &person)
	database.Connector.Save(&person)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(person)
}

func deletPersonByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var person entities.User
	id, _ := strconv.ParseInt(key, 10, 64)
	database.Connector.Where("id = ?", id).Delete(&person)
	w.WriteHeader(http.StatusNoContent)
}
