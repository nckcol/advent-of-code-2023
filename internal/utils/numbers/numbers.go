package numbers

func Gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return Gcd(b, a%b)
}

func Lcm(a int, b int) int {
	return a * b / Gcd(a, b)
}

func LcmSlice(numbers []int) int {
	result := 1
	for _, n := range numbers {
		result = Lcm(result, n)
	}
	return result
}
