package homework

func average(input [15]float32) (result float32) {
	result = 0.0
	var ilen = (float32)(len(input))

	for i := 0; i < len(input); i++ {
		result = input[i]/ilen + result
	}

	return
}
