// file: max.go
package branchless

func MaxIf(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MaxBranchless(a, b int) int {
	return a*(boolToInt(a >= b)) + b*(boolToInt(b > a))
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
