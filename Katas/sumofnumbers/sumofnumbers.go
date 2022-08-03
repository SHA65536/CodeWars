package sumofnumbers

/*
Given two integers a and b, which can be positive or negative, find the sum of all the integers between and including them and return it.
If the two numbers are equal return a or b.
Note: a and b are not ordered!
*/

func GetSum(a, b int) int {
	if a > b {
		a, b = b, a
	}
	n := (b - a + 1)
	return n * (2*a + (n - 1)) / 2
}
