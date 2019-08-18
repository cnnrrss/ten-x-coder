# Datatypes
Bits, Bytes \
ASCII, Unicode, Hex \
Char, Char array \
Primitives: String, Int, Float, Double, Long, Boolean \
- Reverse Integer is like string (push and pop)
	//pop operation:
	pop = x % 10;
	x /= 10;

	//push operation:
	temp = rev * 10 + pop;
	rev = temp;

- Signed and Unsigned Integers
- Handling Overflow
	- Integer.MAX_VALUE/10 && pop > 7
	- Integer.MIN_VALUE/10 && pop < -8


Objects \
Go: rune, int32, uint32, etc...