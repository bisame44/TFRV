package dsp

import (
	"fmt"

	"https://github.com/bisame44/TFRV/tree/main/TFRV/lab2/tasks"
)

func NFDH(records []tasks.Task, n int) ([][]tasks.Task, error) {
	oversize := make([]tasks.Task, 0)
	for _, record := range records {
		if record.R > n {
			oversize = append(oversize, record)
		}
	}
	if len(oversize) != 0 {
		return nil, fmt.Errorf("only having %d processes, while tasks %v require more", n, oversize)
	}

	tasksLength := len(records)
	taskLevels := make([][]tasks.Task, 0)
	if tasksLength == 0 {
		return taskLevels, nil
	}

	j := 0
outerLoop:
	for i := 0; ; i++ {
		takenProc := 0
		taskLevels = append(taskLevels, make([]tasks.Task, 0))
		for {
			takenProc += records[j].R
			if takenProc > n {
				continue outerLoop
			}
			taskLevels[i] = append(taskLevels[i], records[j])
			j++
			if j >= tasksLength {
				break outerLoop
			}
		}
	}

	return taskLevels, nil
}
