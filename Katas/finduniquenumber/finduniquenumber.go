package finduniquenumber

/*
There is an array with some numbers. All numbers are equal except for one. Try to find it!
findUniq([ 1, 1, 1, 2, 1, 1 ]) === 2
findUniq([ 0, 0, 0.55, 0, 0 ]) === 0.55
It’s guaranteed that array contains at least 3 numbers.
The tests contain some very huge arrays, so think about performance.
*/

func FindUniq(arr []float32) float32 {
	var same float32
	if arr[0] == arr[1] || arr[0] == arr[2] {
		same = arr[0]
	} else {
		same = arr[2]
	}
	for _, num := range arr {
		if num != same {
			return num
		}
	}
	return arr[0]
}
