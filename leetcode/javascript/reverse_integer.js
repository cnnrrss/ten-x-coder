/**
 * @param {number} x
 * @return {number}
 */
var reverse = function(x) {
	let rev = 0; 
	while(x > 0) {
	    pop = x % 10;
	    x /= 10;
	    temp = rev * 10 + pop;
	    rev = temp;
	}
    return rev;
};

console.log(reverse(103));