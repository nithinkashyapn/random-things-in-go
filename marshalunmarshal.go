package main

import (
	"encoding/json"
	"fmt"
)

type Details struct {
	Name    string
	Email   string
	Website string
}

func main() {

	details := Details{"nithin", "1234567890", "example.com"}
	fmt.Println("Details structure is ", details)

	detailsAsBytes, _ := json.Marshal(details)

	fmt.Println("Details as Bytes is ", detailsAsBytes)

	var assetunmarshall Details
	err := json.Unmarshal(detailsAsBytes, &assetunmarshall)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("Details structure is ", assetunmarshall)

}
