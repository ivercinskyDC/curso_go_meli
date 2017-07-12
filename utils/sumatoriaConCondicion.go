package utils

//Mint is the same as ints but with my methods
type Mint int

//IsDiv3Or5 returns true if the instance is divisible by 3 o 5
func (x Mint) IsDiv3Or5() bool {
	return x%3 == 0 || x%5 == 0
}

//AddWithCond adds all the Mints from 1 to t that satisfies the condition
func AddWithCond(c func(t Mint) bool, t int) int {
	res := 0
	for i := 1; i <= t; i++ {
		if c(Mint(i)) {
			res += i
		}
	}
	return res
}
