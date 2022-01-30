package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/yaml.v3"
)

type config struct {
	Key string `yaml: "keys"`
}

func (c *config) getKeys() *config {
	yamlFile, err := ioutil.ReadFile("../keys/keys.yml")
	if err != nil {
		log.Printf("Issue getting yaml file #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	return config
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
	fmt.Printf("Here is the key: %s", config.Key)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
