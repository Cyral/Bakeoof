package main

import (
	"fmt"
	"math/rand"
	"time"
)

func preheat() string{
	firstStep := fmt.Sprintf("Preheat oven to %d degrees \n", rand.Intn(350))
	return firstStep
}

func cookTime() int{
	rand.Seed(time.Now().Unix())
	return rand.Intn((30 - 20) + 30)
}

func coolTime() int{
	rand.Seed(time.Now().Unix())
	return rand.Intn((30 - 5 ) + 5)
}


func getEggs() int{
	return rand.Intn(4) + 1
}
func beatEggs() string{
	eggStep := fmt.Sprintf("Beat %d eggs in a bowl \n", getEggs())
	return eggStep
}

func numMuff() int{
	n := rand.Intn(13) + 1
	return n
}

func getMuffinTins() string{
	muffinStep := fmt.Sprintf("Grease %d cup muffin pan or use %d Muffin Liners \n", numMuff(), numMuff())
	return muffinStep
}

func pourMuffin() string{
	pourMuffin := fmt.Sprintf("Pour mixture into %d muffin tins \n", numMuff())
	return pourMuffin
}

func bowlSize() string{
	bowls := [4]string{"small", "medium", "large", "giant"}
	return bowls[rand.Intn(4) - 1]
}

func combineAll() string {
	combineStep := fmt.Sprintf("Mix all your ingridents in a %s size bowl", bowlSize())
	return combineStep
}

