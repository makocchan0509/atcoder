package common

func genIntArr(def, cnt int) []int {
	arr := make([]int, cnt)

	for i := 0; i < cnt; i++ {
		arr[i] = def
	}
	return arr
}

func genBoolArr(def bool, cnt int) []bool {
	arr := make([]bool, cnt)

	for i := 0; i < cnt; i++ {
		arr[i] = def
	}
	return arr
}

func genIntTwo(d, n, w int) [][]int {
	ii := make([][]int, n)
	for i := 0; i < n; i++ {
		ii[i] = genIntArr(-1, w)
	}
	return ii
}
