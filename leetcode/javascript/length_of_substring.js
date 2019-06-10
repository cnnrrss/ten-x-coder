// var lengthOfLongestSubstring = function(s) {
// 	var n = s.length;
// 	var ans = 0;
// 	var index = 0;

// 	for(let i=0; i < n; i++) {
// 		j = Math.max(index[s.charAt(i)], i)
// 		ans = Math.max(ans, j - i + 1);
// 		index[s.charAt(j)] = j + 1;
// 	}
// 	return ans;
// };
// index.charAt()
// var lengthOfLongestSubstring = function(s) {
// 	let longest = 0;
// 	let curStart = 0;
// 	for (let i=0; i<s.length; i++) {
// 		for (let j=0; j<s.slice(curStart, i); j++) {
// 			if (s[i] == s.slice(curStart, i)[j]) {
// 				if (i - curStart > longest) {
// 					longest = i - curStart;
// 				}
// 				curStart += j + 1;
// 				break
// 			}
// 		}
// 	}
// 	if (s.length - curStart > longest) {
// 		longest = s.length - curStart
// 	}
// 	return longest;
// };

// var lengthOfLongestSubstring = function(s) {
// 	let stack = [];
// 	let z = 0;
//     for (let i=0; i < s.length; i++) {
//     	if (stack.length == 0) {
//     		stack.push(s[i]);
//     	} else if(s[i] !== stack[i-1]) {
//     		stack.push(s[i])
//     	} else if (stack.. > z) {
// 			z = stack.length;
// 		}
// 	}
// 	if (z >= stack.length) {
// 		console.log(z, stack);
// 		return z;
// 	}
// 	console.log(z, stack);
// 	return stack.length;
// 	// console.log(stack, z);
// };

var lengthOfLongestSubstring = function(s) {
	if (s.length <= 2) { return s.length; }

	let ans = 0;
	let left = 0;
	let map = new Map();

	for (let i=0; i < s.length; i++) {
		if (map.has(s[i]) && map.has(s[i]) > 0) {
			while (map.get(s[i]) > 0) {
				map.set(s[left], map.get(s[left]) - 1);
				left++;
			}
		}
		map.set(s[i], 1);
		console.log(map, i - left + 1, ans);
		ans = Math.max(ans, i - left + 1);
	}
	return ans;
};

a = lengthOfLongestSubstring("abccdefgghijklmno");
b = lengthOfLongestSubstring("pwwkew");
c = lengthOfLongestSubstring("bb");

console.log("\n\nAns:", a, "-", b, "-", c);

// Loop through each character
// Create list of unique characters
// If a duplicate is hit, record the index val
// and start a new list?
// 