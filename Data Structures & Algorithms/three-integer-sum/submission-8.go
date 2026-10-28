import (
	"slices"
)

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	result := make([][]int, 0)

	for i, x := range nums {
		if i > 0 && x == nums[i-1] {
			continue
		}

		target := -x
		y, z := i + 1, len(nums) - 1
		for y < z {
			sum := nums[y] + nums[z]
			if sum == target {
				result = append(result, []int{x, nums[y], nums[z]})
				y++
				for y < z && nums[y] == nums[y-1] {
					y++
				}
				z--
			} else if sum > target {
				z--
			} else {
				y++
			}
		}
	}

	return result
}
