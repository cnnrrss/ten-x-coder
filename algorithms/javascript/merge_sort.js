const mergeSort = (list) => {
    const len = list.length;
    if (len == 1) {
        return list
    };
    // get middle item 
    const mid = Math.ceil(len / 2);

    let left = list.slice(0, mid);
    let right = list.slice(mid, len);

    left = mergeSort(left);
    right = mergeSort(right);

    return merge(left, right);
}

const merge = (left, right) => {
    const sorted = [];
    while (left.length > 0 && right.length > 0) {
        const lItem = left[0];
        const rItem = right[0];
        if (lItem > rItem) {
            sorted.push(rItem);
            right.shift();
        } else {
            sorted.push(lItem);
            left.shift();
        }
    }
    // if arr's not same length, add the rest
    let pushRemaining = (arr) => {
        while (arr.length !== 0) {
            sorted.push(arr.shift());
        }
    };

    pushRemaining(right);
    pushRemaining(left);

    return sorted
}


let items = [3, 2, 4, 1, 5, 7, 8, 6];
console.log("before sort", items);
items = mergeSort(items);
console.log("after sort", items);

// Merge sort uses a divide-and-conquer approach to sort elements in
// an array. Basically, what this means is that instead of dealing with
// the array as a whole, it continually splits it in half until both
// halves are sorted, then the halves are merged into one solved list.