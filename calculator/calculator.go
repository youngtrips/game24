package calculator

import (
	"fmt"
)

type Calculator struct {
	h    map[string]bool
	res  []string
	nums []int
	sum  int
}

func (c *Calculator) Calc() []string {
	exps := make([]string, 0)
	for _, i := range c.nums {
		exps = append(exps, fmt.Sprintf("%d", i))
	}
	c.calc(len(c.nums), exps)
	return c.res
}

/*
 * 1. C(4, 2) * 6
 * 2. C(3, 2) * 6
 * 3. C(2, 2) * 6
 */
func (c *Calculator) calc(n int, exps []string) {
	if n == 1 && c.nums[0] == c.sum {
		exp := exps[0]
		if _, present := c.h[exp]; !present {
			c.h[exp] = true
			c.res = append(c.res, exp)
		}
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for p := 1; p <= 6; p++ {
				a := c.nums[i]
				b := c.nums[j]
				expA := exps[i]
				expB := exps[j]
				switch p {
				case 1: // +
					c.nums[i] += c.nums[j]
					c.nums[j] = c.nums[n-1]
					if a < b {
						exps[i] = fmt.Sprintf("(%s+%s)", expA, expB)
					} else {
						exps[i] = fmt.Sprintf("(%s+%s)", expB, expA)
					}
					exps[j] = exps[n-1]
					c.calc(n-1, exps)
					break
				case 2: // -
					c.nums[i] -= c.nums[j]
					c.nums[j] = c.nums[n-1]
					exps[i] = fmt.Sprintf("(%s-%s)", expA, expB)
					exps[j] = exps[n-1]
					c.calc(n-1, exps)
					break
				case 3: // *
					c.nums[i] *= c.nums[j]
					c.nums[j] = c.nums[n-1]
					if a < b {
						exps[i] = fmt.Sprintf("%s*%s", expA, expB)
					} else {
						exps[i] = fmt.Sprintf("%s*%s", expB, expA)
					}
					exps[j] = exps[n-1]
					c.calc(n-1, exps)
					break
				case 4: // /
					if c.nums[j] != 0 && c.nums[i]%c.nums[j] == 0 {
						c.nums[i] /= c.nums[j]
						c.nums[j] = c.nums[n-1]
						exps[i] = fmt.Sprintf("(%s/%s)", expA, expB)
						exps[j] = exps[n-1]
						c.calc(n-1, exps)
					}
					break
				case 5: // rev -
					c.nums[i] = c.nums[j] - c.nums[i]
					c.nums[j] = c.nums[n-1]
					exps[i] = fmt.Sprintf("(%s-%s)", expB, expA)
					exps[j] = exps[n-1]
					c.calc(n-1, exps)
					break
				case 6: // rev /
					if c.nums[i] != 0 && c.nums[j]%c.nums[i] == 0 {
						c.nums[i] = c.nums[j] / c.nums[i]
						c.nums[j] = c.nums[n-1]
						exps[i] = fmt.Sprintf("(%s/%s)", expB, expA)
						exps[j] = exps[n-1]
						c.calc(n-1, exps)
					}
					break
				}
				c.nums[i] = a
				c.nums[j] = b
				exps[i] = expA
				exps[j] = expB
			}
		}
	}
}

func New(nums []int, sum int) *Calculator {
	return &Calculator{
		h:    make(map[string]bool),
		res:  make([]string, 0),
		nums: nums,
		sum:  sum,
	}
}
