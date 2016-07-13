package profitbricks

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)


// bad_status is the return string for bad status errors
func bad_status(wanted, got int) string {
	return " StatusCode is " + strconv.Itoa(got) + " wanted " + strconv.Itoa(wanted)
}

// Set Username and Password here for Testing.
/*var username = "jclouds@stackpointcloud.com"
var passwd = os.Getenv("PB_PASSWORD")*/
var username = os.Getenv("PROFITBRICKS_USERNAME")
var passwd = os.Getenv("PROFITBRICKS_PASSWORD")

// Set Endpoint for testing
var endpoint = "https://api.profitbricks.com/rest/v2"

func TestSetAuth(t *testing.T) {
	fmt.Println("Current Username ", Username)
	SetAuth(username, passwd)
	fmt.Println("Applied Username ", Username)
}

func TestSetEndpoint(t *testing.T) {
	SetEndpoint(endpoint)
	fmt.Println("Endpoint is ", Endpoint)
}

func TestMain(m *testing.M) {
	r := m.Run()
	SetAuth(os.Getenv("PROFITBRICKS_USERNAME"), os.Getenv("PROFITBRICKS_PASSWORD"))
	os.Exit(r)
}

// Setup creds for single running tests
func setupCredentials() {
	SetEndpoint("https://api.profitbricks.com/rest/v2")
	SetAuth(os.Getenv("PROFITBRICKS_USERNAME"), os.Getenv("PROFITBRICKS_PASSWORD"))
}
