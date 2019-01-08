package ClimbingStairs

func climbStairs(n int) int {
	if n <= 0 {
		return 1
	}
	twoSteps, oneStep := 1, 1
	for i := 2; i <= n; i++ {
		twoSteps, oneStep = oneStep, twoSteps+oneStep
	}
	return oneStep
}
