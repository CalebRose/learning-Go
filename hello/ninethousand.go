package hello

import (
	"fmt"
	"math/rand"
	"time"
)

func ninethousand() {

	// Changing the seed will change the way random numbers are generated
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("Hello, World.")
	powerLevel := rand.Intn(50000)
	if powerLevel > 9000 {
		fmt.Println("It's nine THOUSAAAAAAAAAAND!")
	} else {
		fmt.Println("My power level is:", powerLevel)
	}
	fmt.Println(addition(2, 2))
	a, b := swap("hello", "there")
	fmt.Println(a, b)
	fmt.Println(split(25))
	for i := 0; i < 50000; i++ {

	}
	fmt.Println("Done!")
}

func addition(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	var random = rand.Intn(6)
	x = sum * random / 9
	y = sum - x
	return
}
