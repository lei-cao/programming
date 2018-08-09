package sort

type Sort interface {
	Sort(nums []int, id string)
}

func DoSort(nums []int, sorter Sort, id string) {
	tmp := make([]int, len(nums))
	copy(tmp, nums)
	sorter.Sort(tmp, id)
}
