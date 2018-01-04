package main

import (
	"net/http"
	"fmt"
	"os"
	"log"
	"io/ioutil"
	//"strconv"
)


func main() {
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(string(responseData))

	Data := string(responseData)

	IsPresent := true
	if Data == ""{
		IsPresent = false
	}
	if IsPresent {
		os.Exit(1)
	}

}
