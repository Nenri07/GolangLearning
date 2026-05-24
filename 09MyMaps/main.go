package main

import "fmt"

func main() {
	fmt.Println("this is the class of Maps")

	var languages = make(map[string]string)

	languages["NJ"] = " new Jersy"
	languages["Ny"] = "new York"

	fmt.Println("this is languages", languages)

	delete(languages, "Ny")
	fmt.Println("Updated languages", languages)

	languages["JP"] = "NihonJin"
	languages["Pt"] = "Brazilian"

	for key, value := range languages {
		fmt.Printf("this is key %v\n and this is value %v\n", key, value)
	}

	//comma ok syntax

	for _, value := range languages {
		fmt.Printf("this is value %v\n", value)
	}
}
