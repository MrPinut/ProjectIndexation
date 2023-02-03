package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readDataFile() []string {
	f, err := os.Open("finalResult.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var apiData []string

	for scanner.Scan() {

		//fmt.Println(scanner.Text())
		apiData = append(apiData, scanner.Text())
	}
	//fmt.Println("apiData")
	fmt.Println(apiData)
	//fmt.Println(apiData[2])

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return apiData
}
