package kata

func FindOutlier(arr []int) int {
    evenCount, oddCount := 0, 0

    for i := 0; i < 3; i++ {
        if arr[i] % 2 == 0 {
            evenCount++
        } else {
            oddCount++
        }
    }

    if evenCount > oddCount {
        for _, num := range arr {
            if num % 2 != 0 {
                return num
            }
        }
    } else {
        for _, num := range arr {
            if num % 2 == 0 {
                return num
            }
        }
    }
    
    return 0
}
