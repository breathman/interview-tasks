package arrays

import (
	"fmt"
	"github.com/influxdata/influxdb/pkg/testing/assert"
	"testing"
)

func TestMerge(t *testing.T) {
	ar1, ar2 := []int{1, 3, 5, 8}, []int{1, 2, 4, 9}
	ar3, ar4 := []int{1, 8}, []int{1, 2, 4, 5, 9}
	fmt.Println(ar1)
	fmt.Println(ar2)

	res := Merge(ar1, ar2)
	res2 := Merge(ar3, ar4)
	assert.Equal(t, res, []int{1, 1, 2, 3, 4, 5, 8, 9})
	assert.Equal(t, res2, []int{1, 1, 2, 4, 5, 8, 9})
}
