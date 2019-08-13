function heapsAlgorithm(str) {
	var regex = /(.)\1+/g;
	var arr = str.split('');
	var permutations = [];
	var tmp;

	// return 0 if str contains same character
  	if (str.match(regex) !== null && str.match(regex)[0] === str) {
  		console.log("The same because the first match in regex is === to the total string.")
  		return 0;
  	}

	// swapWithTemp swaps the two values using a temp variable
	function swapWithTemp(x, y) {
		tmp = arr[x]
		arr[x] = arr[y];
		arr[y] = tmp;
	}

	// generate arrays of permutations 
	function generate(int) {
		if (int === 1) {
			permutations.push(arr.join(''));
		} else {
			for (var i = 0; i != int; i++) {
				generate(int - 1);
				swapWithTemp(int % 2 ? 0 : i, int-1);
			}
		}
	}

	generate(arr.length);

	var filtered = permutations.filter(function(string) {
		return !string.match(regex);
	});
	
	console.log(filtered);
	return filtered.length;
}

var ans = heapsAlgorithm("aaaBB");
console.log(ans);