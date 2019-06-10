// It’s very common in questions of this type to work 
// through the optimal solutions in this manner. 
// First find a brute-force approach, 
// then see if sorting would improve things, and 
// finally always consider a greedy approach which uses 
// either an array or a hash map to store values. 
// Chances are the greedy approach is the best one.

// Brute force naive approach ---- O(n^2) time
const twoSum = (arr, target) => {
    let results = [];
    for (let i=0; i<arr.length; i++) {
		for (let j=i+1; j<arr.length; j++) {
			if (arr[j] === target - arr[i]) {
				results.push([arr[i], arr[j]])
			}
		}
    }
    return results;
}

// O(n log n)
const binarySearch = (sortedArr, target) => {
	let min = 0,
	   max = sortedArr.length - 1,
	   guess;
   while (min <= max) {
	   guess = Math.floor((min + max) / 2);
	   if (sortedArr[guess] === target) {
		   return guess;
	   } else {
		   if (sortedArr[guess] < target) {
			   min = guess + 1;
		   } else {
			   max = guess - 1;
		   }
	   }
   }
   return false;   
}

const twoSum = (arr, target) => {
	let sortedArr = arr.sort(),
		results = [];

	for (let i=0; i<sortedArr.length; i++) {
    	let diff = target - sortedArr[i],
			binaryIndex = binarySearch(sortedArr, diff);
		if (binaryIndex) {
        	results.push([sortedArr[i], sortedArr[binaryIndex]]);
		}
	}
	return results;
}

// Best Solution is a `Greedy Approach` O(n)!
const twoSum = (arr, target) => {
	let map = {},
		results = [];
	for (let i=0; i<arr.length; i++) {
		if (map[arr[i]] !== undefined) {
			results.push([map[arr[i]], arr[i]])
		} else {
			map[target - arr[i]] = arr[i];
		}
	}
	return results;
}

var twoSum = function(nums, target) {
	let m = new Map();
	for (let i=0; i<nums.length; i++) {
		if (m.has(target - nums[i])) {
			return [m.get(target - nums[i]), i];
		} else {
            m.set(nums[i], i);
		}
	}
	return results;
};