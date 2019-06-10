func judgeCircle(moves string) bool {
    x, y := 0,0 
    for _, dir := range moves {
        fmt.Println(dir)
        switch {
            case dir == 85:
                y++
            case dir == 68:
                y--
            case dir == 76:
                x--
            default:
                x++
        }   
    }
    
    return x == 0 && y == 0
}