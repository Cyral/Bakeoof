package main

import (
	"fmt"
	"math/rand"
	"time"
)



func cookTime() int{
	rand.Seed(time.Now().Unix())
	return rand.Intn((30 - 20 ) + 30)
}

func coolTime() int{
	rand.Seed(time.Now().Unix())
	return rand.Intn((30 - 5 ) + 5)
}

func cookCake(ingridients [7]string ){
	var size [3]string
	var sizeNum int

	sizeNum  = rand.Intn(3)

	fmt.Printf("in a %s size bowl, add baking powder, flour, ",size[sizeNum])
	fmt.Printf("add %d eggs and whisk all together\n",getEggs())
	fmt.Printf("put contents of bowl into  cake pan. Put directly into oven %d minutes.\n",cookTime())
	fmt.Printf("Let the cake pan cool for %d minutes.\n",coolTime())
}