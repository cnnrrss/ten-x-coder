/**
 * @param {string[]} strs
 * @return {string}
 */
var longestCommonPrefix = function(strs) {
    if (strs.length == 0) {
        return '';
    }
    for (let i=0; i <= strs[0].length; i++) {
        for (let str of strs) {
            if (str[i] !== strs[0][i]) {
                return str.slice(0, i);    
            }
        }
    }
    return strs[0];
};

console.log(longestCommonPrefix(["flow", "flower", "flight"]));
console.log(longestCommonPrefix(["flow", "flight", "flower"]));
console.log(longestCommonPrefix(["abb", "abc"]));
console.log(longestCommonPrefix(["abb", "babc"]));
console.log(longestCommonPrefix(["abb", "abbc"]));
console.log(longestCommonPrefix(["abbcde", "abbc"]));