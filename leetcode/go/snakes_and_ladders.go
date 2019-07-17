func snakesAndLadders(board [][]int) int {
    n, steps := len(board), 0
    dist := map[int]int{}
    queue := []int{1}
    dist[1] = 0
    
    for len(queue) > 0 {
        steps++
        size := len(queue)      
        for j := 0; j < size; j++ {
            s := queue[j]
            for i := 1; i <= 6; i++ {
                s2 := s + i
                r, c := getPosition(s2, n)
                if board[r][c] != -1 {
                    dist[s2] = dist[s] + 1
                    s2 = board[r][c]
                }
                if s2 == n*n {
                    return steps
                }
                if _, ok := dist[s2]; !ok {
                    dist[s2] = dist[s] + 1
                    queue = append(queue, s2)
                }
            }
        }
        queue = queue[size:]
    }

    return -1
}
