package main

// hasTwoComplements indicates that complements contains two values which sums equals to target value
func hasTwoComplements(target int, complements []int) bool {
	set := map[int]struct{}{}
	for _, v := range complements {
		complement := target - v
		if _, ok := set[complement]; ok == true {
			return true
		}
		set[v] = struct{}{}
	}
	return false
}
