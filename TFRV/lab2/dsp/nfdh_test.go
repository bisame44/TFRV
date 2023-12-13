package dsp_test

import (
	"testing"

	"github.com/Belstowe/distrib-cs-1-autumn/lab2/dsp"
	"github.com/Belstowe/distrib-cs-1-autumn/lab2/tasks"
)

func TestNFDH(t *testing.T) {
	records := []tasks.Task{{3, 6}, {5, 3}, {4, 2}, {6, 2}, {2, 2}, {7, 1}}
	packed, err := dsp.NFDH(records, 10)
	if err != nil {
		t.Fatal(err.Error())
	}
	tt := dsp.TotalTime(packed)
	if tt != 10 {
		t.Errorf("nfdh pass through test case %v should give %d total time, not %d", records, 10, tt)
	}
}
