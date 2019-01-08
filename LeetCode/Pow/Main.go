package Pow

func myPow(x float64, n int) float64 {
	if n < 0 {
		n = -n
		x = 1 / x
	}
	return getResult(x, n)
}

func getResult(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	result := getResult(x, n/2)
	result *= result
	if n%2 != 0 {
		result *= x
	}
	return result
}

func anotherPow(x float64, n int) float64 {
	if n < 0 {
		n = -n
		x = 1 / x
	}
	var result float64 = 1
	for n != 0 {
		if n&1 == 1 {
			result *= x
		}
		n >>= 1
		x *= x
	}
	return result
}
