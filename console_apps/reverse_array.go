package console_apps

func reverseArray(arr []int) []int {
	x := 0
	y := 0
	c := len(arr) - 1
	for i := 0; i < c; i++ {
		x = arr[i]
		y = arr[c]
		arr[i] = y
		arr[c] = x
		c--
	}
	return arr
}
