func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)
	for _, num := range nums { // construct freq map
		freqMap[num]++
	}

	buckets := make([][]int, len(nums))
	for num, freq := range freqMap { // put into buckets
		buckets[freq - 1] = append(buckets[freq - 1], num)
	}

	result := make([]int, 0, k)
	for i := len(buckets) - 1; i >= 0; i-- { // select top k
		for _, num := range buckets[i] {
			result = append(result, num)

			if len(result) == k {
				return result
			}
		}
	}

	return result
}
