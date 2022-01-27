package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

func main() {
	yfile, err := ioutil.ReadFile("items.yaml")
	if err != nil {
		log.Fatal(err)
	}
	data := make(map[interface{}]interface{})
	err2 := yaml.Unmarshal(yfile, &data)
	if err2 != nil {
		log.Fatal(err2)
	}
	for k, v := range data {
		fmt.Printf("%s -> %v\n", k, v)
	}
}
