package twoballs

import "math"

func twoballs(breaks []bool) int {
    sr := math.Sqrt(float64(len(breaks)))
    isPerfect := int(sr * sr) == len(breaks)
    if isPerfect {
        breaks = append(breaks, true)
    }
    
    step := int(math.Floor(math.Sqrt(float64(len(breaks)))))

    firstBreak := step
    for i := firstBreak; i < len(breaks); i += step {
        if breaks[i] {
            firstBreak = i
            break
        }
    }

    start := firstBreak - step
    end := firstBreak
    var res int

    for i := start; i <= end; i++ {
        if breaks[i] {
            res = i
            break
        }
    }

    return res + 1
}
