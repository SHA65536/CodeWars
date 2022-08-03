package int32toipv4

import "fmt"

/*
Take the following IPv4 address: 128.32.10.1
This address has 4 octets where each octet is a single byte (or 8 bits).
1st octet 128 has the binary representation: 10000000
2nd octet 32 has the binary representation: 00100000
3rd octet 10 has the binary representation: 00001010
4th octet 1 has the binary representation: 00000001
So 128.32.10.1 == 10000000.00100000.00001010.00000001
Because the above IP address has 32 bits, we can represent it as the unsigned 32 bit number: 2149583361
Complete the function that takes an unsigned 32 bit number and returns a string representation of its IPv4 address.
*/

func Int32ToIp(n uint32) string {
	var a, b, c, d uint32
	a = n & 255
	b = (n & 65280) >> 8
	c = (n & 16711680) >> 16
	d = (n & 4278190080) >> 24
	return fmt.Sprintf("%d.%d.%d.%d", d, c, b, a)
}
