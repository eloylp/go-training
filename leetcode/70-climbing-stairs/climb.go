package climb

func climbStairs(n int) int {
	if n <= 3 {
		return n
	}
	return climbStairs(n-1) + climbStairs(n-2)
}

func cl(n int) int {
	f := make([]int, n+1, n+2)
	if n < 3 {
		f = f[0:3]
	}
	f[0] = 0
	f[1] = 1
	f[2] = 2
	for i := 3; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}

// 1 2 3 5
