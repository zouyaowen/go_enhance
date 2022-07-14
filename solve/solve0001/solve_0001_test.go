package solve0001

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSum(t *testing.T) {
	assert.Equal(t, []int{0, 1}, twoSum([]int{2, 7, 11, 15}, 9))
	assert.Equal(t, []int{}, twoSum([]int{2, 7, 11, 15}, 4))
}
