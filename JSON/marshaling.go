/**
-	struct ----> marshal ----> json in byte form ----> convert to string to print it
-	json ----> unmarshal ----> original struct in string form
**/
package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Name string
	Body string
	Time int64
	mode string
}

func main() {

	m := Message{"Mr.Robot", "Hello Hacker", 1294706395881547000, "chat"}

	fmt.Println(m)

	// MARSHALING a JSON
	b, err := json.Marshal(m)

	if err != nil {
		fmt.Println("Error occured while marshaling json")
	}

	// After marshaling b contains data in bytes
	fmt.Println(b)

	// Convert data stored in b to string
	// json package can only pick the exported fields in the struct
	fmt.Println(string(b))

	// UNMARSHALING a JSON

	var n Message

	// Following preference is followed while unmarshaling - 1) Checks for an exported field with tag <json key> 2) Checks for an exported field with name <json key> 3) Checks for an exported field named <case INSENSTIVE json key>
	error := json.Unmarshal(b, &n)

	if error != nil {
		fmt.Println("Error occured while unmarshaling json")
	}

	// Prints back the original json
	fmt.Println(n)

}
