package main

import (
	"fmt"
	"slices"
)

func main() {

	// Search & Compare
	a := []int{1, 2, 3, 4, 5}

	slices.Contains(a, 3) // true
	slices.Contains(a, 9) // false

	slices.Index(a, 3)                                        // 2
	slices.Index(a, 9)                                        // -1
	slices.IndexFunc(a, func(v int) bool { return v%2 == 0 }) // 1 (first even)

	slices.Equal([]int{1, 2}, []int{1, 2}) // true
	slices.Equal([]int{1, 2}, []int{2, 1}) // false

	slices.EqualFunc([]int{1, 2}, []int{1, 2}, func(a, b int) bool { return a == b }) // true

	// Sorting & Reversing
	nums := []int{5, 2, 4, 1, 3}

	slices.Sort(nums) // [1 2 3 4 5]
	slices.SortFunc(nums, func(a, b int) int {
		return a - b
	})
	slices.IsSorted(nums) // true
	slices.IsSortedFunc(nums, func(a, b int) int {
		return a - b
	}) // true

	slices.Reverse(nums) // [5 4 3 2 1]

	// Cutting, Cloning & Appending
	s := []int{1, 2, 3, 4, 5}

	slices.Clone(s)               // makes a copy
	slices.Concat(s, []int{6, 7}) // [1 2 3 4 5 6 7]

	slices.Insert(s, 2, 99, 100) // [1 2 99 100 3 4 5]
	slices.Delete(s, 1, 3)       // remove elements at [1:3] → [1 4 5]

	slices.Grow(s, 10) // grows capacity by at least 10 (like make)

	// Compacting & Trimming
	s1 := []int{1, 1, 2, 2, 2, 3, 3, 4}

	slices.Compact(s1)                                            // [1 2 3 4] (removes consecutive duplicates)
	slices.CompactFunc(s1, func(a, b int) bool { return a == b }) // same but custom

	// Binary Search
	nums1 := []int{1, 3, 5, 7, 9}

	i, found := slices.BinarySearch(nums1, 5)
	// i=2, found=true

	i, found = slices.BinarySearch(nums1, 6)
	// i=3, found=false (where 6 would be inserted)

	// i := slices.BinarySearchFunc(nums1, 8, func(a, b int) int { return a - b })
	// i=4
	fmt.Println(i, found)

	// Replace, Compare & Min/Max
	s11 := []int{1, 2, 3}
	s22 := []int{1, 2, 4}

	slices.Compare(s11, s22) // -1 (s1 < s2)
	slices.Compare(s11, s11) // 0  (equal)
	slices.Compare(s22, s11) // 1  (s2 > s1)

	slices.CompareFunc(s11, s22, func(a, b int) int {
		return a - b
	}) // -1

	slices.Min([]int{3, 1, 4, 2}) // 1
	slices.Max([]int{3, 1, 4, 2}) // 4

	// Misc Utilities
	// Replace elements in a slice
	s3 := []int{1, 2, 3, 4}
	s4 := slices.Replace(s3, 1, 3, 9, 9, 9) // [1 9 9 9 4]

	// Clip capacity to current length
	s3 = s3[:2]
	s3 = slices.Clip(s3) // len=2, cap=2
	fmt.Println(s3, s4)

	// Special Helpers
	// DeleteFunc: remove elements by condition
	nums4 := []int{1, 2, 3, 4, 5}
	nums4 = slices.DeleteFunc(nums4, func(v int) bool {
		return v%2 == 0
	})
	fmt.Println(nums4) // [1 3 5]

	// IndexFunc example
	idx := slices.IndexFunc(nums4, func(v int) bool { return v > 3 }) // 2
	fmt.Println(idx)
}
