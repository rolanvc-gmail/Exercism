package lasagna

// TODO: define the 'PreparationTime()' function

// TODO: define the 'Quantities()' function

// TODO: define the 'AddSecretIngredient()' function

// TODO: define the 'ScaleRecipe()' function

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
func PreparationTime(slice []string, aveTime int) int {
	if aveTime == 0 {
		aveTime = 2
	}

	return len(slice) * aveTime

}

func Quantities(slice []string) (noodles int, sauce float64) {
	noodles = 0
	sauce = 0
	for _, layer := range slice {
		if layer == "sauce" {
			sauce += 0.2

		} else if layer == "noodles" {
			noodles += 50
		}
	}
	return
}

func AddSecretIngredient(friendsList []string, myList []string) {
	friendListLength := len(friendsList)
	myListLength := len(myList)
	secretIngredient := friendsList[friendListLength-1]
	if myList[myListLength-1] == "?" {
		myList[myListLength-1] = secretIngredient
	}
}

func ScaleRecipe(amounts []float64, portions int) []float64 {
	scale := float64(float64(portions) / 2)
	result := make([]float64, len(amounts), len(amounts))
	for i, amt := range amounts {
		result[i] = amt * scale
	}

	return result
}
