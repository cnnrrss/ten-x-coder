func minMeetingRooms(intervals [][]int) int {
    periods := Periods{}
    for _, interval := range intervals {
        periods = append(periods, Period{
            time: interval[0], 
            isStart: true,
        })
        periods = append(periods, Period{
            time: interval[1], 
            isStart: false,
        })
    }
    
    sort.Sort(periods)
    result, roomCount := 0, 0
    
    for _, pd := range periods {
        if pd.isStart {
            roomCount++
            result = max(result, roomCount)
        } else {
            roomCount--
        }
    }
    return result
}

type Period struct { 
    time int
    isStart bool 
}

type Periods []Period
func (p Periods) Len() int { return len(p) }
func (p Periods) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p Periods) Less(i, j int) bool {
    if p[i].time == p[j].time {
        return !p[i].isStart && p[j].isStart
    }
    return p[i].time < p[j].time
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}