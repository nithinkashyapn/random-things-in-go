package main

import (
	"fmt"
	"strings"
)

func main() {

	exampleString := "Failed to copy asset: Transaction processing for endorser [localhost:7051]: Chaincode status Code: (500) UNKNOWN. Description: No such record to query 100 from the ledger"
	a := strings.Split(exampleString, "Description: ")
	fmt.Println(a[len(a)-1])
}
