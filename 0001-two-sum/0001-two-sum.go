func twoSum(nums []int, target int) []int {
    result := []int{}
    for i := 0 ; i < len(nums) ; i++ {
        for j := i ; j < len(nums) ; j++ {
            if nums[i] + nums[j] == target && i != j {
                result = []int{i,j}
            } 
        }
    }
    return result
}