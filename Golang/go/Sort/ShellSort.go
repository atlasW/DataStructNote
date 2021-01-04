package sort

//步长收缩
// n log2 n
func ShellSort(array []int) {
	n := len(array)
	if n < 2 {
		return
	}
	key := n / 2
	for key > 0 {
		for i := key; i < n; i++ {
			j := i
			for j >= key && array[j] < array[j-key] {
				array[j], array[j-key] = array[j-key], array[j]
				j = j - key
			}
		}
		key = key / 2
	}
}
