package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

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
