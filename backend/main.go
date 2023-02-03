package main

var IPAddress string = "http://34.77.36.161:"
var SecondIp string = "http://34.77.36.161:3941"

func main() {

	secretKey := firstCall(IPAddress)

	secondCall(SecondIp, secretKey)

	apiData := readDataFile()

	lastCall(IPAddress, apiData[0], apiData[1], apiData[2])

}
