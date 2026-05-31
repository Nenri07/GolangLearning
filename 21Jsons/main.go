package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("this is json class")
	// coursesJson()
	decodeJson()
}

type course struct {
	Name  string   `json:"coursename"`
	Price int      `json:"coursePrice"`
	Tags  []string `json:"tagsofcourse,omitempty"`
	Pass  string   `json:"-"`
	Tutor string   `json:"tutorName"`
}

func coursesJson() {
	courses := []course{
		{"golang", 24, []string{"gin", "go", "multi"}, "22Arid4340", "nenrei"},
		{"nextjs", 24, []string{"tsx", "js", "web server side rendering"}, "Arid4340", "nenrei007"},
		{"Japnease", 24, []string{"jlpt", "nihongo"}, "4340", "nenreisensei"},
		{"postgre", 22, nil, "22Arid4340", "nenrei"},
	}
	// interfaces are actually just a struct simply
	jsonData, err := json.MarshalIndent(courses, "", "\t")

	CheckNillError(err)

	fmt.Println("this is data from the jsonconverter marshal", string(jsonData))
	fmt.Printf("%s\n This is data f pf", jsonData)
}
func CheckNillError(err error) {
	if err != nil {
		panic(err)
	}
}

func decodeJson() {
	bytejson := []byte(`
		{
                "coursename": "Japnease",
                "coursePrice": 24,
                "tagsofcourse": ["jlpt","nihongo"],
                "tutorName": "nenreisensei"
        }
	`)

	var courseJson course

	isValid := json.Valid(bytejson)
	if isValid {
		fmt.Println("this is valid json")
		json.Unmarshal(bytejson, &courseJson)
		fmt.Printf("this is json used%#v\n", courseJson)
	} else {
		fmt.Println("Json Invalid")
	}

	var keyJson map[string]any
	json.Unmarshal(bytejson, &keyJson)
	fmt.Println("this is valid map json")
	fmt.Printf("this is %#v\n", keyJson)

	for key, value := range keyJson {
		fmt.Printf("This is key %v and this is value %v and type is %T\n", key, value, value)
	}

}

type Abser interface {
	Abs() float64
}
