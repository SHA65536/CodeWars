package nextperfectsquare

/*
You might know some pretty large perfect squares. But what about the NEXT one?
Complete the findNextSquare method that finds the next integral perfect square after the one passed as a parameter.
Recall that an integral perfect square is an integer n such that sqrt(n) is also an integer.
If the parameter is itself not a perfect square then -1 should be returned. You may assume the parameter is non-negative.
*/
import "math"

func FindNextSquare(sq int64) int64 {
	root := math.Sqrt(float64(sq))
	if root != float64(int64(root)) {
		return -1
	}
	root++
	return int64(root * root)
}
