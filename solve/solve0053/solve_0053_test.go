package solve0053

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxSubArray(t *testing.T) {
	// nums := [...]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	nums := make([]int, 0)
	nums = append(nums, -2, 1, -3, 4, -1, 2, 1, -5, 4)

	assert.Equal(t, 6, maxSubArray(nums))
}
