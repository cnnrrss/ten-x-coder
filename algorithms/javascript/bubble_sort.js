// Time O(n2)
// The algorithm works by comparing each item in the list 
// with the item next to it and swapping them if required.

const bubbleSort = (arr) => {
    for (let i = arr.length-1; i >= 0; i--) {
        for (let j = 1; j <= i; j++) {
            if (arr[j-1] > arr[j]) {
                let temp = arr[j-1];
                arr[j-1] = arr[j];
                arr[j] = temp;
            }
        }
    }
    return arr
};

let items = [3, 2, 4, 1, 5, 7, 8, 6];
console.log("before sort", items);
items = bubbleSort(items);
console.log("after sort", items);

// One step of the algorithm
// 7, 5, 2, 4, 3, 9
// 5, 7, 2, 4, 3, 9
// 5, 2, 7, 4, 3, 9
// 5, 2, 4, 7, 3, 9
// 5, 2, 4, 3, 7, 9
// 5, 2, 4, 3, 7, 9