package homework

func reverse(input []int64) (result []int64) {
	result = make([]int64, len(input))

	var ilen = len(input)

	for i := 0; i < ilen; i++ {
		result[ilen-i-1] = input[i]
	}

	return
}
