/**
 * @param {number} x
 * @return {number}
 */
var reverse = function(x) {
	let i = 0;
	let rev = 0;
	let rem = 0;
	while (i < x.length) {
		console.log(i);
		rev = x / 10;
		rem = x % 10;
		i++
	}
};

reverse(123);