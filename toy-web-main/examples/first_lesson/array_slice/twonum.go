package main

func main() {

}

func twonum(nums []int, sum1 int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if idx, found := numMap[complement]; found {
			return []int{idx, i}
		}
		numMap[num] = i
	}
	return nil // 如果没有找到符合条件的索引
}
