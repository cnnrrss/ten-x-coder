func fizzBuzz(n int) []string {
    fizzList := []string{}
    for i:=1; i<=n; i++ {
        if i % 3 == 0 && i % 5 == 0 {
            fizzList = append(fizzList, "FizzBuzz")
        } else if i % 3 == 0 {
            fizzList = append(fizzList, "Fizz")
        } else if i % 5 == 0 {
            fizzList = append(fizzList, "Buzz")
        } else {
            fizzList = append(fizzList, fmt.Sprintf("%v", i))
        }
    }
    return fizzList
}