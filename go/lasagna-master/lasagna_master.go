package lasagna

// PreparationTime is the time required for preparing the lasagna.
func PreparationTime(layers []string, time int) int {
	if time == 0 {
		time = 2
	}
	return len(layers) * time
}

// Quantities determines the quantities of noodles and sauce needed.
func Quantities(layers []string) (noodles int, sauce float64) {
	noodleGramsPerLayer := 50
	sauceLitresPerLayer := 0.2

	noodlesLayers := 0
	for _, v := range layers {
		if v == "noodles" {
			noodlesLayers++
		}
	}

	sauceLayers := 0
	for _, v := range layers {
		if v == "sauce" {
			sauceLayers++
		}
	}

	return noodlesLayers * noodleGramsPerLayer, float64(sauceLayers) * sauceLitresPerLayer
}

// AddSecretIngredient adds the secret ingredient from your friend's list.
func AddSecretIngredient(f, m []string) {
	m[len(m)-1] = f[len(f)-1]
}

// ScaleRecipe scales the recipe.
func ScaleRecipe(quantities []float64, portions int) (scaledRecipe []float64) {
	scaledRecipe = make([]float64, len(quantities))

	for i, v := range quantities {
		scaledRecipe[i] = v / 2 * float64(portions)
	}

	return scaledRecipe
}
