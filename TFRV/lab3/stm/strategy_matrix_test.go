package stm_test

import (
	"testing"

	"github.com/Belstowe/distrib-cs-1-autumn/lab3/stm"
)

func TestStrategyMatrix(t *testing.T) {
	sm := stm.NewStrategyMatrix([][]int{
		{0, 6, 12, 18, 24},
		{4, 2, 10, 16, 22},
		{8, 6, 4, 14, 22},
		{12, 10, 8, 6, 18},
		{16, 14, 12, 10, 8},
	})
	sm.Iterate()
	sm.Iterate()
	sm.Iterate()
}
