var maxProfit = function (prices) {
    let currentMax = 0;
    let max = 0;
    for (let i = prices.length - 1; i >= 0; i--) {
        let price = prices[i];
        currentMax = Math.max(currentMax, price);
        let potential = currentMax - price;
        max = Math.max(max, potential);
    }
    return max
};

// 5, 5, 6
console.log(maxProfit([9, 11, 8, 5, 10]));
console.log(maxProfit([9, 4, 8, 9, 3]));
console.log(maxProfit([3, 4, 8, 9, 3]));