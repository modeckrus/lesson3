package tdd

func Add(a, b int) int {
	summ := a + b
	if summ < 0 {
		return -summ
	}
	if summ == 0 {
		return 1
	}
	return summ
}

func Summattor(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

func AddSummattor(a, b int) int {
	return Summattor(a, b, Add)
}
