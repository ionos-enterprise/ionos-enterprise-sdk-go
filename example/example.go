package main

import (
	"fmt"

	"github.com/profitbricks/profitbricks-sdk-go"
)

func main() {

	//Sets username and password
	profitbricks.SetAuth("vendors@stackpointcloud.com", "d:;,%28*4$EZW98/")
	//Sets depth.
	profitbricks.SetDepth("5")
	dcs := profitbricks.ListDatacenters()

	fmt.Println(dcs.Items[0].Id)
}
