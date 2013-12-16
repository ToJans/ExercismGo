package hamming

func Distance(A string, B string) int {
	distance := 0
	for i := 0; i < len(A) && i < len(B); i++ {
		if A[i] != B[i] {
			distance += 1
		}

	}
	return distance
}
