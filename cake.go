package main

import (
	"fmt"
	"math/rand"
)

func displaySteps(args []string ) {

	fmt.Printf("In a %s size bowl, add baking powder, bla ","yo")

}

func main() {
	var size [3]string
	var sizeNum int

	sizeNum  = rand.Intn(3)

	fmt.Printf("in a %s size bowl, add baking powder, flour, (ingrendients go here) \n",size[sizeNum])
	fmt.Printf("add %d eggs and whisk all together\n",3)
	fmt.Printf("put contents of bowl into  cake pan. Put directly into oven %d minutes.\n",30)
	fmt.Printf("Let the cake pan cool for %d minutes.\n",20)

}
