package algo

// Euclid (algo book p.912)
func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

// The least common multiple (lcm) of a and b is their product divided by their
// greatest common divisor (gcd)
// lcm(a, b) = ab/gcd(a,b)).
func LCM(a, b int) int {
	return a * b / GCD(a, b)
}

func LCMSlice(nums []int) int {
	lcm := LCM(nums[0], nums[1])
	for _, n := range nums[2:] {
		lcm = LCM(lcm, n)
	}
	return lcm
}
