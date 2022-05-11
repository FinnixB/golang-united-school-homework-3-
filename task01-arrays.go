package homework

func average(input [15]float32) (result float32) {
	result = 0.0

	for i := 0; i < len(input); i++ {
		result = input[i] + result
	}

	result = result / len(input)

	return
}
