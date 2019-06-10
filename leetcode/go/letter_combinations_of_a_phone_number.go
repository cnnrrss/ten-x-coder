var numberMap = map[rune]string {
    '2': "abc",
    '3': "def",
    '4': "ghi",
    '5': "jkl",
    '6': "mno",
    '7': "pqrs",
    '8': "tuv",
    '9': "wxyz",
}

func letterCombinations(digits string) []string {
    var combinations = []string{}
    for _, d := range digits {
        letters := numberMap[d]
        
        var temp []string
        for _, l := range letters {
            if len(combinations) == 0 {
                temp = append(temp, string(l))
            } else {
                for _, c := range combinations {
                    temp = append(temp, c + string(l))
                }
            }
        }
        combinations = temp
    }
    return combinations
}