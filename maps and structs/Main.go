package main

import (
	"fmt"
	// reflect package used to get to tags of fields
	"reflect"
)

// type uppercase as convention if not exporting
type House struct {
// uppercase exports structs fields
	number int
	rooms []string
	m2 map[string]int
}

type Animal struct {
	Name string `required:"true" max:"100"`
	Origin string
}

type Bird struct {
	// composition is not inheritance
	// Bird is not Animal but can call fields from Animal
	Animal
	speedKPH float32
	canFly bool
}

func main() {
	// declarations 
	countryPopulations := make(map[string]int)
	// can use make function is key/ values are not known
	// else as below with :=
	countryPopulations = map[string]int{
		// key and value pair declared above
		"Italy": 3446334,
		"France": 455424,
		"UK": 6235232,
		"Germany": 823432,
		"Spain": 9234423,
	}
	fmt.Print(countryPopulations)

	// can create map like this because array is a
	// valid keytype but slice is not.
	m := map[[3]int]string{}
	fmt.Print(m);
	fmt.Println("/n")

	// access values of keys
	fmt.Println(countryPopulations["Italy"])
	// modify value
	countryPopulations["Italy"] = 1
	fmt.Println(countryPopulations["Italy"])
	// doing this can modify order of keys
	fmt.Println(countryPopulations)
	// delete key and value
	delete(countryPopulations, "Spain")
	fmt.Println(countryPopulations)
	// deleted keys return value 0
	fmt.Println(countryPopulations["Spain"])

	// ok can be used to check if key exists
	spain, ok := countryPopulations["Spain"]
	fmt.Println(spain, ok)
	italy ,ok := countryPopulations["Italy"]
	fmt.Println(italy, ok)

	cp := countryPopulations
	// this deletes key in both cp and original map
	delete(cp, "UK")
	fmt.Println(cp)
	fmt.Println(countryPopulations)

	// struct are useful for mixing data types
	// the above is declared above main func
	myHouse := House{
	// removing the keys still works but not recommended
		number: 62,
		rooms: []string {
			"Kitchen",
			"Bedroom",
			"Living Room",
		},
		m2: map[string]int{
			"inside": 52,
			"outside": 10,
		},
	}
	fmt.Println(myHouse)
	// . is used to drill into struct
	fmt.Println(myHouse.rooms)
	// [] for index of slice/array
	fmt.Println(myHouse.rooms[1])

	// anonymous struct
	// used to organise data for single use
	person := struct{name string}{name: "Hamza Zbili"}
	fmt.Println(person)
	someoneElse := person
	// using the & pointer modifies person as arrays
	thirdPerson := &someoneElse
	thirdPerson.name = "Bugs Bunny"
	fmt.Println(person)
	fmt.Println(someoneElse)
	fmt.Println(thirdPerson)

	// call Animal struct to create Bird
	 sparrow := Bird{}
	 sparrow.Name = "latin name"
	 sparrow.Origin = "some park"
	 sparrow.speedKPH = 50
	 sparrow.canFly = true

	 fmt.Println(sparrow)
	 fmt.Println("/n")

	// add tag to add information
	t := reflect.TypeOf((Animal{}))
	field, _ := t.FieldByName("Name")
	fmt.Println(field) //{Name  string required max:"100" 0 [0] false}
}
