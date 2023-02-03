package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var IPAddress string = "http://34.77.36.161:"
var SecondIp string = "http://34.77.36.161:3941"

func main() {

	var compteur int = 3368
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
			fmt.Println(string(data))

			secretKeyTempo := strings.Split(string(data), ": ")
			secretKey = secretKeyTempo[1]
			fmt.Println("The secret Key is: " + secretKey + " on port: " + port)
		}
	}
	// Ici je fait un Json Vide
	values := map[string]string{}
	json_data, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(SecondIp+"?secretKey="+secretKey, "application/json",
		bytes.NewBuffer(json_data))

	//response, err := http.Get(SecondIp + "?secretKey=" + secretKey)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		compteur++
	} else {
		found = false
		data, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(data))
	}

	// Je prends l'info qu'il y a dans le txt

	f, err := os.Open("finalResult.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var apiData []string

	for scanner.Scan() {

		fmt.Println(scanner.Text())
		apiData = append(apiData, scanner.Text())
	}
	fmt.Println("apiData")
	fmt.Println(apiData)
	fmt.Println(apiData[2])

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	resp2, err := http.Post(IPAddress+apiData[0]+"?"+apiData[1]+"="+apiData[2], "application/json",
		bytes.NewBuffer(json_data))
	//response2, err := http.Get(SecondIp + "?secretKey=" + secretKey)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		compteur++
	} else {
		found = false
		data, _ := ioutil.ReadAll(resp2.Body)
		fmt.Println(string(data))
	}

}
