package sort

/*
 * @Author       : zhixi.fang
 * @Date         : 2022-06-13 22:23:14
 * @LastEditors  : zhixi.fang
 * @LastEditTime : 2022-06-13 23:04:58
 */

// QuickSort 快速排序
func QuickSort[T int8 | int16 | int32 | int | int64 | float32 | float64](numbers ...T) []T {
	quickSort(numbers, 0, len(numbers)-1)
	return numbers
}

func quickSort[T int8 | int16 | int32 | int | int64 | float32 | float64](nums []T, left, right int) {
	for left < right {
		mid := partition(nums, left, right)
		quickSort(nums, left, mid-1)
		left = mid + 1
	}
}

func partition[T int8 | int16 | int32 | int | int64 | float32 | float64](nums []T, left, right int) int {
	// choose a pivot

	// 改变的地方
	pivot := threeSumMedian(nums[left], nums[(left+right)/2], nums[right])
	//pivot := nums[left]
	nums[left], pivot = pivot, nums[left]
	// 改变结束

	for left < right {
		for pivot <= nums[right] && left < right {
			right -= 1
		}
		nums[left] = nums[right]

		for nums[left] <= pivot && left < right {
			left += 1
		}
		nums[right] = nums[left]
	}

	nums[left] = pivot
	return left

}

// input 10 20 30 ---> return 20 ; input 10 10 11 --> return 10
func threeSumMedian[T int8 | int16 | int32 | int | int64 | float32 | float64](a, b, c T) T {

	if a < b {
		a, b = b, a
	}
	// a  > b

	if c > a {
		return a
	} else {
		if c > b {
			return c
		} else {
			return b
		}
	}
}
