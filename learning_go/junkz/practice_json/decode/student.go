package decode

import (
	"encoding/json"
	"os"
)

type Student struct {
	ID            int      `json:"id"`
	LastName      string   `json:"last_name"`
	MiddleInitial string   `json:"middle_initial"`
	FirstName     string   `json:"first_name"`
	Enrolled      bool     `json:"enrolled"`
	Classes       []Course `json:"classes"`
}

type Course struct {
	CourseName  string `json:"CourseName"`
	CourseNum   int    `json:"CourseNum"`
	CourseHours int    `json:"CourseHours"`
}

func ReadJsonFile(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func ParseJson(data []byte, student *Student) error {
	err := json.Unmarshal(data, student)
	if err != nil {
		return err
	}
	return nil
}
