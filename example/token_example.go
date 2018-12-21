package main

import (
	"fmt"
	"github.com/profitbricks/profitbricks-sdk-go"
	"os"
)

func main() {

	fmt.Println("Token: " + os.Getenv("PROFITBRICKS_TOKEN"))
	//Sets token
	//A token can be retrieved by:
	//curl --user '<username:password>' https://api.profitbricks.com/auth/v1/tokens/generate
	client := profitbricks.NewClientbyToken(os.Getenv("PROFITBRICKS_TOKEN"))

	datacenters, err := client.ListDatacenters()
	var datacenterids []string

	for _, dc := range datacenters.Items {
		datacenterids = append(datacenterids, dc.ID)
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println(datacenterids)
		os.Exit(0)
	}
}
