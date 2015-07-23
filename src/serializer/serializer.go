package main

import "fmt"
import "encoding/json"

type Circle struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Radius int `json:"radius"`
}

func main() {
	shapes := []Circle{
		Circle{X: 1, Y: 2, Radius: 3},
		Circle{X: 2, Y: 3, Radius: 4},
	}
	fmt.Println(shapes)
	
	jsonBytes, _ := json.Marshal(shapes)
	fmt.Println(string(jsonBytes))
	
	var circles []Circle
	json.Unmarshal(jsonBytes, &circles);
	fmt.Println(circles)
}
