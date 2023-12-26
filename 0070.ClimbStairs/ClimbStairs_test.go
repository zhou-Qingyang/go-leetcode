package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem(t *testing.T) {
	fmt.Printf("------------------------Leetcode Problem 70------------------------\n")

	fmt.Printf("上100级台阶 有多少种走法:%d\n", climbStairs(100))
	fmt.Printf("上3级台阶 有多少种走法:%d\n", climbStairs(3))
	fmt.Printf("上5级台阶 有多少种走法:%d\n", climbStairs(5))
	fmt.Printf("上2级台阶 有多少种走法:%d\n", climbStairs(2))
}
