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
	eggStep := fmt.Sprintf("Beat %d eggs in a %s sized bowl \n", getEggs(), bowlSize())
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

func cookAll() string{
	cookStep := fmt.Sprintf("Bake for %d minutes", cookTime())
	return cookStep
}

func coolLocation() string{
	bowls := [2]string{"refrigerator", "freezer"}
	return bowls[rand.Intn(2) - 1]
}

func coolAll() string{
	coolStep := fmt.Sprintf("Cool for %d minutes in the %s", coolTime(), coolLocation())
	return coolStep
}



func CookMuffin() []string {
	steps := make([]string, 7)
	steps[0] = preheat()
	steps[1] = getMuffinTins()
	steps[2] = beatEggs()
	steps[3] = combineAll()
	steps[4] = pourMuffin()
	steps[5] = cookAll()
	steps[6] = coolAll()

	return steps
}

func main() {
	fmt.Print(CookMuffin())
}