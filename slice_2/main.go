package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"alice", "bob", "cindy"}
	n, found := slices.BinarySearch(names, "bob")
	fmt.Println(n, found)
	nums1 := []int{1, 2, 3}
	nums2 := []int{4, 5, 6}
	nums3 := []int{7, 8, 9}
	nums := slices.Concat(nums1, nums2, nums3)
	fmt.Println(nums)
	names = slices.Delete(names, 0, 2)
	fmt.Println(names)
	fmt.Println(slices.Equal(nums1, nums2))
	idx := slices.Index(names, "cindy")
	fmt.Println(idx)
	fmt.Println(slices.Max(nums))
}
