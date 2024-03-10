// @leet start
func twoSum(nums []int, target int) []int {
  // empty map to store the indices
  m := make(map[int]int)
  for i, v := range nums {
    // target - current value = complement
    if j, ok := m[target - v]; ok {
        return []int{j, i}
    }
    m[v] = i
  }
  return []int{}
}
// @leet end
