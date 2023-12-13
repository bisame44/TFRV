package dsp_test

import (
	"testing"

	"https://github.com/bisame44/TFRV/tree/main/TFRV/lab2/dsp"
	"https://github.com/bisame44/TFRV/tree/main/TFRV/lab2/tasks"
)

func TestFFDH(t *testing.T) {
	records := []tasks.Task{{3, 6}, {5, 3}, {4, 2}, {6, 2}, {2, 2}, {7, 1}}
	packed, err := dsp.FFDH(records, 10)
	if err != nil {
		t.Fatal(err.Error())
	}
	tt := dsp.TotalTime(packed)
	if tt != 9 {
		t.Errorf("ffdh pass through test case %v should give %d total time, not %d", records, 9, tt)
	}
}
