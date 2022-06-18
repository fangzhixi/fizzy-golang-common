package sort

/*
 * @Author       : zhixi.fang
 * @Date         : 2022-06-13 22:23:14
 * @LastEditors  : zhixi.fang
 * @LastEditTime : 2022-06-18 11:07:57
 */

// QuickSort 快速排序
func QuickSort[T int8 | int16 | int | int32 | int64 | float32 | float64](numbers ...T) []T {
	quickSort(numbers, 0, len(numbers)-1)
	return numbers
}

// QuickSortByArray 快速排序
func QuickSortByArray[T int8 | int16 | int | int32 | int64 | float32 | float64](numbers []T) []T {
	quickSort(numbers, 0, len(numbers)-1)
	return numbers
}

func quickSort[T int8 | int16 | int | int32 | int64 | float32 | float64](nums []T, left, right int) {
	for left < right {
		mid := partition(nums, left, right)
		quickSort(nums, left, mid-1)
		left = mid + 1
	}
}

// partition 分类排序
func partition[T int8 | int16 | int | int32 | int64 | float32 | float64](nums []T, left, right int) int {
	// choose a pivot

	// 改变的地方
	pivot := threeNumMedian(nums[left], nums[(left+right)/2], nums[right])
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

// 三个数值中间值 （Min <= Med =< Max）
func threeNumMedian[T int8 | int16 | int | int32 | int64 | float32 | float64](a, b, c T) T {

	if a < b {
		a, b = b, a
	}

	if c > a {
		return a
	}

	if c > b {
		return c
	}
	return b
}
