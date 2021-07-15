func lengthOfLIS(nums []int) int {
    dp := make([]int,len(nums))
    finalLength := 0
    for i :=0 ; i < len(dp); i++ {
            dp[i] = 1
    }
    
    for i := 1 ; i < len(nums); i++ {
        for j := 0; j < i ; j++ {
            if nums[i] > nums[j]{
                if dp[i] < dp[j]+1 {
                        dp[i] = dp[j]+1
                }
            }
        }
    }
    fmt.Println(dp)
    for i := 0; i < len(dp); i++ {
        if finalLength < dp[i] {
                finalLength = dp[i]
        }
    }
    return finalLength
}


//Input = {10, 22, 9, 33, 21, 50, 41, 60, 80}
//output = 6
//Explanation 	{10, 22, 33, 50, 60, 80}