package day1

func Find2(nums []int) int {
	for _, x1 := range nums {
		for _, x2 := range nums {
			if x1+x2 == 2020 {
				return x1 * x2
			}
		}
	}
	return -1
}

func Find3(nums []int) int {
	for _, x1 := range nums {
		for _, x2 := range nums {
			for _, x3 := range nums {
				if x1+x2+x3 == 2020 {
					return x1 * x2 * x3
				}
			}
		}
	}
	return -1
}
