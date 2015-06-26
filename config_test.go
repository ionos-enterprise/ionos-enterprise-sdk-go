package profitbricks

import (
		"testing"
		"fmt"
		"os"
		"strconv"
)

// bad_type is the return string for bad type errors
func bad_type(shouldbe, got string) string {
	return " Type is " + got + " should be " + shouldbe
}

// bad_status is the return string for bad status errors
func bad_status(wanted, got int) string {
	return " StatusCode is " + strconv.Itoa(got) + " wanted " + strconv.Itoa(wanted)
}

// Set Username and Password here for Testing.

var username = "jclouds@stackpointcloud.com"
var passwd = os.Getenv("PB_PASSWORD")

// Set Endpoint for testing
var endpoint = "https://api.profitbricks.com/rest"

func TestSetAuth(t *testing.T) {
	fmt.Println("Current Username ", Username)
	SetAuth(username, passwd)
	fmt.Println("Applied Username ", Username)
}

func TestSetEndpoint(t *testing.T) {
	SetEndpoint(endpoint)
	fmt.Println("Endpoint is ", Endpoint)

}

// Setup creds for single running tests
func setupCredentials() {
	SetAuth(username, passwd)
}
