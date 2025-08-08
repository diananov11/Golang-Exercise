package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, minute int) int {
    if minute == 0 {
        return len(layers) * 2
    }
    return len(layers) * minute
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int,float64) {
    noodles, sauce := 0, 0.0
    for _,v := range layers {
        if v == "sauce" {
            sauce += 0.2 
        } else if v == "noodles" {
            noodles += 50
        }
    }
    return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList, myList []string) {
	myList[len(myList)-1] = friendList[len(friendList)-1] 
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portion int) []float64 {
    multiplesPortion := float64(portion)/2
    newQuantities := []float64{}
    for _, v := range quantities {
        newQuantities = append(newQuantities, v * float64(multiplesPortion))
    }
    return newQuantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
