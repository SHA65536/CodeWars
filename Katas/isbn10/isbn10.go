package isbn10

/*
ISBN-10 identifiers are ten digits long. The first nine characters are digits 0-9. The last digit can be 0-9 or X, to indicate a value of 10.
An ISBN-10 number is valid if the sum of the digits multiplied by their position modulo 11 equals zero.
*/

func ValidISBN10(isbn string) bool {
	var sum int
	if len(isbn) != 10 {
		return false
	}
	for i := range isbn {
		switch {
		case isbn[i] >= '0' && isbn[i] <= '9':
			sum += int(isbn[i]-'0') * (i + 1)
		case i == 9 && (isbn[i] == 'X' || isbn[i] == 'x'):
			sum += 10 * (i + 1)
		default:
			return false
		}
	}
	return sum%11 == 0
}
