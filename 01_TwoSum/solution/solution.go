package solution

func TwoSum(nums []int, target int) []int {

	hashTable := make(map[int]int)
	for i, num := range nums {
		val, ok := hashTable[num]
		if !ok {
			hashTable[target-num] = i
		} else {
			return []int{val, i}
		}
	}

	return nil

}
