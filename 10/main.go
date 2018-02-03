// Interface used to reuse the same function written for 2 different types
// Interface check for the function with same types of parameter & return value
// Interface can't assign values in tat type

package main

import (
	"fmt"
)

type infy interface {
	noOfBuildings() int //
}

type puneDc struct {
	buildings int
}
type blrDc struct {
	buildings int
}

func main() {

	pd := puneDc{
		buildings: 50,
	}
	bd := blrDc{
		buildings: 60,
	}

	// Without interface
	printPdNoOfBuildings(pd)
	printBdNoOfBuildings(bd)

	// with interface
	printNoOfBuildings(pd)
	printNoOfBuildings(bd)

}

func printNoOfBuildings(dc infy) {
	fmt.Println("No of buildings ", dc.noOfBuildings()) // this is defined in interface
}

// Function logic is same but we cannot reuse
// passing diffrent types of arguments

func printPdNoOfBuildings(pd puneDc) {
	fmt.Println("No of buildings ", pd.noOfBuildings())
}

func printBdNoOfBuildings(bd blrDc) {
	fmt.Println("No of buildings ", bd.noOfBuildings())
}

func (pd puneDc) noOfBuildings() int {
	return pd.buildings
}

func (bd blrDc) noOfBuildings() int {
	return bd.buildings
}
