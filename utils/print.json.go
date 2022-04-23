package utils

import (
	"encoding/json"
	"fmt"
)

// Print Pretty Json or slice Json
func PrintJSON(str string, data interface{}) {
	var p []byte
	p, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s = %s \n", str, p)
}
