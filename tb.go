package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	winsSwitch := 0
	winsKeep := 0
	repetitons := 9999
	for i := 0; i < repetitons; i++ {
		b := createBoxes()
		// choose a random box
		choice := rand.Int32N(3)
		var chosenBox bool
		switch choice {
		case 0:
			chosenBox = b.FirstBox
		case 1:
			chosenBox = b.SecondBox
		case 2:
			chosenBox = b.ThirdBox
		}

		if chosenBox {
			winsKeep++
		} else {
			winsSwitch++
		}

	}

	fmt.Println("Repetitons of the test:", repetitons)
	fmt.Println("Wins with Switching strategy:", winsSwitch)
	fmt.Println("Wins with Keeping strategy:", winsKeep)
	fmt.Println()
	fmt.Println("Switching strategy win rate:", float64(winsSwitch)*100/float64(repetitons), "%")
	fmt.Println("Keeping strategy win rate:", float64(winsKeep)*100/float64(repetitons), "%")
}

type Boxes struct {
	FirstBox  bool
	SecondBox bool
	ThirdBox  bool
}

// create 3 boxes, in one of them a ball(true), the rest are empty(false)
func createBoxes() Boxes {
	// create an array of random permutation of 0, 1 and 2
	ar := rand.Perm(3)

	var boxes Boxes
	// depending on where is 2, assign the ball
	switch {
	case ar[0] == 2:
		boxes.FirstBox = true
		boxes.SecondBox = false
		boxes.ThirdBox = false
	case ar[1] == 2:
		boxes.FirstBox = false
		boxes.SecondBox = true
		boxes.ThirdBox = false
	case ar[2] == 2:
		boxes.FirstBox = false
		boxes.SecondBox = false
		boxes.ThirdBox = true
	}
	return boxes
}
