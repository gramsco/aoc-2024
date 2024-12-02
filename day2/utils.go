package main

func removeLevelByindex(reports []int, index int) []int {
	var clone []int
	for i, val := range reports {
		if i != index {
			clone = append(clone, val)
		}
	}
	return clone
}
