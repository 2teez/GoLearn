package main

import (
	"encoding/json"
	"fmt"
)

type greeting struct {
	Message string
}

type student struct {
	StudentId     int      `json:"id"`
	LastName      string   `json:"lname"`
	MiddleInitial string   `json:"-"`
	FirstName     string   `json:"fname"`
	IsEnrolled    bool     `json:"enrolled,omitempty"`
	Courses       []course `json:"classes"`
}

type course struct {
	Name   string `json:"coursename"`
	Number int    `json:"coursenumber"`
	Hours  int    `json:"coursehours"`
}

func unmarshalJson(data []byte, v any) {
	err := json.Unmarshal(data, &v)
	if err != nil {
		fmt.Println(err)
	}
}

func marshalJson(s any) {
	jsData, err := json.MarshalIndent(s, "", "	")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("\n", string(jsData))
}
