package encoder

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

func RunEncodeDefaultTypes() {
	num := 42
	str := "hello"
	boolVal := true
	arrays := [4]any{"hello", 34, true, 3.142}

	var output strings.Builder
	encoder := json.NewEncoder(&output)

	encoder.Encode(num)
	encoder.Encode(str)
	encoder.Encode(boolVal)
	encoder.Encode(arrays)

	fmt.Println(output.String())
}

func RunEncodeCustomTypes() {
	type Person struct {
		Name string    `json:"name"`
		Age  time.Time `json:"age"`
	}

	makeNewPerson := func(name string, age int) *Person {
		ageTime := time.Date(time.Now().Year()-age, 1, 1, 0, 0, 0, 0, time.UTC)
		return &Person{Name: name, Age: ageTime}
	}
	p := makeNewPerson("John", 36)

	/*calculatedAge := func(p *Person) int {
		return time.Now().Year() - p.Age.Year()
	}

	printablePerson := func(p *Person) string {
	return fmt.Sprintf("%s is %d years old", p.Name, calculatedAge(p))
	}*/
	var output strings.Builder
	encoder := json.NewEncoder(&output)
	encoder.Encode(p)

	fmt.Println(output.String())
}
