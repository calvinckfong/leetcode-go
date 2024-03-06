// 1. Two Sum
func twoSum(nums []int, target int) []int {
    m := make(map[int]int)

    for idx, num := range nums {
        if pos, ok := m[target-num]; ok {
            return []int{pos, idx}
        }
        m[num] = idx
    }
    return []int{}
}
