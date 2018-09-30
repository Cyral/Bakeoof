package main

import (
	"fmt"
	"math/rand"
	"time"
)

func degrees() int {
	return rand.Intn((450 - 300) + 400)
}
func preheat() string {
	firstStep := fmt.Sprintf("Preheat oven to %d degrees \n", degrees())
	return firstStep
}
func cookTime() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn((30 - 20) + 30)
}

func coolTime() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn((30 - 5) + 5)
}

func getEggs() int {
	return rand.Intn(4) + 1
}
func beatEggs() string {
	eggStep := fmt.Sprintf("Beat %d eggs in a %s sized bowl \n", getEggs(), bowlSize())
	return eggStep
}

func numMuff() int {
	n := rand.Intn(13) + 1
	return n
}

func getMuffinTins() string {
	muffinStep := fmt.Sprintf("Grease %d cup muffin pan or use %d Muffin Liners \n", numMuff(), numMuff())
	return muffinStep
}

func pourMuffin() string {
	pourMuffin := fmt.Sprintf("Pour mixture into %d muffin tins \n", numMuff())
	return pourMuffin
}

func bowlSize() string {
	bowls := [4]string{"small", "medium", "large", "giant"}
	return bowls[rand.Intn(4)]
}

func combineAll() string {
	combineStep := fmt.Sprintf("Mix all your ingredients in a %s size bowl\n", bowlSize())
	return combineStep
}

func combineMain() string {
	combineStep := fmt.Sprintf("Mix your flour, baking powder, and eggs in a %s sized bowl", bowlSize())
	return combineStep
}

func fillMuff() string {
	fillStep := fmt.Sprintf("Whisk together all your remaining ingredients. Pour mixture into %s squeeze bottle and fill an equal amount into each muffin\n", bowlSize())
	return fillStep
}

func muffLayer(muffArray []string) string {
	layerStep := fmt.Sprintf("Chop the %s into %s pieces. Roast it on the stovetop, Set aside for later", muffArray[2], bowlSize())
	return layerStep
}

func layerMuff() string {
	layeringStep := fmt.Sprintf("Using the ingredients you chopped earlier sprinkle them onto cooked muffins")
	return layeringStep
}

func combineRemain() string {
	combineStep := fmt.Sprintf("Mix your remaining ingredients in a %s size bowl\n", bowlSize())
	return combineStep
}

func cookMuff() string {
	cookStep := fmt.Sprintf("Bake for %d minutes\n", cookTime())
	return cookStep
}

func coolLocation() string {
	bowls := [2]string{"refrigerator", "freezer"}
	return bowls[rand.Intn(2)]
}

func coolAll() string {
	coolStep := fmt.Sprintf("Cool for %d minutes in the %s\n", coolTime(), coolLocation())
	return coolStep
}

func cakeContents() string {
	combineStep := fmt.Sprintf("put contents of bowl into  cake pan. Put directly into oven %d minutes.\n", cookTime())
	return combineStep
}

func CookCake() []string {
	steps := make([]string, 5)
	steps[0] = preheat()
	steps[1] = beatEggs()
	steps[2] = combineAll()
	steps[3] = cakeContents()
	steps[4] = coolAll()
	return steps
}

func CookMuffin(muffArray []string) []string {
	rand.Seed(time.Now().Unix())
	randSelect := rand.Intn(3)
	steps := make([]string, 9)

	// 0 Generic Muffins
	// 1 Filling Muffins
	// 2 Crumb Muffins
	if randSelect == 0 {
		steps[0] = preheat()
		steps[1] = getMuffinTins()
		steps[2] = beatEggs()
		steps[3] = combineAll()
		steps[4] = pourMuffin()
		steps[5] = cookMuff()
		steps[6] = coolAll()
		return steps[0:6]
	} else if randSelect == 1 {
		steps[0] = preheat()
		steps[1] = getMuffinTins()
		steps[2] = beatEggs()
		steps[3] = combineMain()
		steps[4] = pourMuffin()
		steps[5] = cookMuff()
		steps[6] = fillMuff()
		steps[7] = coolAll()
		return steps[0:7]
	} else if randSelect == 2 {
		steps[0] = preheat()
		steps[1] = getMuffinTins()
		steps[2] = beatEggs()
		steps[3] = muffLayer(muffArray)
		steps[4] = combineRemain()
		steps[5] = pourMuffin()
		steps[6] = cookMuff()
		steps[7] = layerMuff()
		steps[8] = coolAll()
		return steps[0:8]
	}

	return nil
}

func cookieContents() string {
	combineStep := fmt.Sprintf("Mix eggs with melted butter, flour,baking soda, and salt\n")
	return combineStep
}
func cookieAddIngridients() string {
	combineStep := fmt.Sprintf(" add finely chopped -ingridients here-  into the bowl. Mix thourghly\n")
	return combineStep
}

func cookieCookInstruc() string {
	combineStep := fmt.Sprintf(" For each cookie, drop 1/4 a cup onto the baking tray\n")
	return combineStep
}

func CookCookies() []string {
	steps := make([]string, 6)
	steps[0] = preheat()
	steps[1] = beatEggs()
	steps[2] = cookieContents()
	steps[3] = cookieAddIngridients()
	steps[4] = cookieCookInstruc()
	steps[5] = coolAll()
	return steps
}

func GetSteps(recipe Recipe) []string {
	bakeType := recipe.BakeType

	if bakeType == "Muffins" {
		return CookMuffin(recipe.Ingredients)
	} else if bakeType == "Cake" {
		return CookCake()
	} else if bakeType == "Cookies" {
		return CookCake()
	}
	return nil
}
