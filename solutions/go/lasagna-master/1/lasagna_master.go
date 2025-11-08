package lasagna

// TODO: define the 'PreparationTime()' function

func PreparationTime(layers []string, minutes int) int {
       if minutes == 0 {
        minutes = 2
    }
    return len(layers) * minutes 
   
    
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
    noodleLayer := 0
    sauceLayer := 0
    for _, v := range layers{
        if v == "noodles"{
            noodleLayer++
        }else if v == "sauce"{
            sauceLayer++
        }
    }
    noodle := noodleLayer * 50
    sauce := float64(sauceLayer) * 0.2
    return noodle, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
   myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, numPortions int) []float64 {
	scaled := make([]float64, len(quantities)) // create a new slice of the same length

	for i, v := range quantities {
		scaled[i] = v * float64(numPortions) / 2 // scale based on 2-portion recipe
	}

	return scaled
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
