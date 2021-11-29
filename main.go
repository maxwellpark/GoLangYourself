package main

import (
	"encoding/json"
	"fmt"
	"gly/auth"
	"gly/data"
	"gly/req"
	"log"
)

func main() {
	auth.Run()

	content := req.GetUser()
	var data data.Data

	// Parse JSON - pass bytes and pointer to the target obj (data)
	err := json.Unmarshal([]byte(content), &data) 
	if err != nil {
		log.Fatal(err)
		return
	}
	user := data.User

	// %+v to see the field name and value
	fmt.Printf("%+v", user)
}

func getMsg(name string) string {
	var msg string = fmt.Sprintf("Welcome, %v.", name)
	return msg
}

func GetMsg(name string) string {
	return getMsg(name)
}
