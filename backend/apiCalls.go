package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func firstCall(IPAddress string) string {
	var compteur int = 3000
	var found bool = true
	var secretKey string

	for found && compteur < 4002 {

		port := strconv.Itoa(compteur)
		response, err := http.Get(IPAddress + port)
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
			compteur++
		} else {
			found = false
			data, _ := ioutil.ReadAll(response.Body)
			//fmt.Println(string(data))

			secretKeyTempo := strings.Split(string(data), ": ")
			secretKey = secretKeyTempo[1]
			fmt.Println("The secret Key is: " + secretKey + " on port: " + port)
		}
	}

	return secretKey
}

func secondCall(ipAddress string, keyValue string) {
	// Ici je fait un Json Vide
	values := map[string]string{}
	json_data, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(ipAddress+"?secretKey="+keyValue, "application/json",
		bytes.NewBuffer(json_data))

	//response, err := http.Get(SecondIp + "?secretKey=" + secretKey)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(data))
	}
}

func lastCall(ipAddress string, port string, keyName string, keyValue string) {

	values := map[string]string{}
	json_data, err := json.Marshal(values)

	resp2, err := http.Post(ipAddress+port+"?"+keyName+"="+keyValue, "application/json",
		bytes.NewBuffer(json_data))
	//response2, err := http.Get(SecondIp + "?secretKey=" + secretKey)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)

	} else {

		data, _ := ioutil.ReadAll(resp2.Body)
		fmt.Println(string(data))

	}
}
