package student

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
	Courses       []Course `json:"courses"`
}

type Course struct {
	CourseName  string `json:"course_name"`
	CourseNum   int    `json:"course_num"`
	CourseHours int    `json:"course_hours"`
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
