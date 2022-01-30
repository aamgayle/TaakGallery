package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
)

type config struct {
	Key string `yaml:"twitter"`
}

func (c *config) getKeys() (*config, error) {
	yamlFile, err := ioutil.ReadFile("keys/keys.yml")

	if err != nil {
		log.Printf("There's an error here: %v", err)
	}

	err = yaml.Unmarshal(yamlFile, c)
	return c, err
}

func homePage(w http.ResponseWriter, r *http.Request) {
	twitterConfig := new(config)
	twitKey, err := twitterConfig.getKeys()
	var key = twitKey.Key
	if err != nil {
		log.Printf("Issue getting yaml file #%v ", err)
		return
	}

	fmt.Fprint(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
	fmt.Printf("Here is the key: %s", key)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
