package calculator

import (
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/novalagung/golpal"
)

func randInts(min int32, max int32, cnt int) []int {
	rand.Seed(time.Now().UTC().UnixNano())

	nums := make([]int, cnt)
	for i := 0; i < cnt; i++ {
		nums[i] = int(min + rand.Int31n(max-min))
	}
	return nums
}

func doTest(t *testing.T, nums []int) {
	c := New(nums, 24)
	exps := c.Calc()
	for _, exp := range exps {
		//t.Log("exp: ", exp)
		if res, err := golpal.New().ExecuteSimple(exp); err != nil {
			t.Error("eval expr failed: ", exp, err)
			t.FailNow()
		} else {
			if n, err := strconv.Atoi(res); err != nil {
				t.Error("eval expr failed: ", exp, err)
				t.FailNow()
			} else if n != 24 {
				t.Error("invalid expr : ", exp)
				t.FailNow()
			}
		}
	}
}

func TestCalculatorSimple(t *testing.T) {
	nums := []int{6, 6, 6, 6}
	//t.Log("nums: ", nums)
	doTest(t, nums)
}

func TestCalculatorRandom(t *testing.T) {
	nums := randInts(1, 10, 4)
	//t.Log("nums: ", nums)
	doTest(t, nums)
}
