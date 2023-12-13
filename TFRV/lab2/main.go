package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/Belstowe/distrib-cs-1-autumn/lab2/dsp"
	"github.com/Belstowe/distrib-cs-1-autumn/lab2/output"
	"github.com/Belstowe/distrib-cs-1-autumn/lab2/tasks"
)

const (
	None int = iota
	NFDH
	FFDH
)

func main() {
	filepath, algo, rn, prettyPrint := flagParse()
	records := readRecords(filepath)
	records = sortRecords(records)

	jsonFormat := json.Marshal
	if prettyPrint {
		jsonFormat = func(v any) ([]byte, error) {
			return json.MarshalIndent(v, "", "  ")
		}
	}

	switch algo {
	case None:
		fmt.Println(records)
	case NFDH:
		fmt.Println(string(assert(jsonFormat(outputAlgoEfficiency(records, rn, dsp.NFDH)))))
	case FFDH:
		fmt.Println(string(assert(jsonFormat(outputAlgoEfficiency(records, rn, dsp.FFDH)))))
	}
}

func flagParse() (string, int, int, bool) {
	filepath := flag.String("f", "", "path to file with tasks (.csv format, split with spaces, machines & time columns)")
	nfdhFlag := flag.Bool("nfdh", false, "use NFDH (Next Fit Decreasing Height) algorithm")
	ffdhFlag := flag.Bool("ffdh", false, "use FFDH (First Fit Decreasing Height) algorithm")
	rn := flag.Int("n", 0, "num of elementary machines")
	prettyPrint := flag.Bool("p", false, "pretty print algorithm result")

	flag.Parse()

	if *filepath == "" {
		invalidUsage("path should be set")
	}

	algo := None
	if *nfdhFlag {
		algo = NFDH
	} else if *ffdhFlag {
		algo = FFDH
	} else {
		invalidUsage("one of algorithm flag set required")
	}

	if *rn <= 0 {
		invalidUsage("num of elementary machines either not set or is not positive")
	}

	return *filepath, algo, *rn, *prettyPrint
}

func readRecords(filepath string) []tasks.Task {
	f := assert(os.Open(filepath))

	r := csv.NewReader(f)
	r.Comma = ' '
	r.Comment = '#'

	records := make([]tasks.Task, 0)
	rawRecords := assert(r.ReadAll())
	for _, rawRecord := range rawRecords {
		records = append(records, tasks.Task{
			R: assert(strconv.Atoi(rawRecord[0])),
			T: assert(strconv.Atoi(rawRecord[1])),
		})
	}
	return records
}

func sortRecords(records []tasks.Task) []tasks.Task {
	min, max := records[0].T, records[1].T
	for _, record := range records {
		if record.T < min {
			min = record.T
		} else if record.T > max {
			max = record.T
		}
	}

	recordBuckets := make([][]tasks.Task, max-min+1)
	for i := 0; i < max-min+1; i++ {
		recordBuckets[i] = make([]tasks.Task, 0)
	}

	for _, record := range records {
		recordBuckets[record.T-min] = append(recordBuckets[record.T-min], record)
	}

	sortedRecords := make([]tasks.Task, 0, len(records))
	for i := len(recordBuckets) - 1; i >= 0; i-- {
		sortedRecords = append(sortedRecords, recordBuckets[i]...)
	}
	return sortedRecords
}

func outputAlgoEfficiency(records []tasks.Task, rn int, algo func([]tasks.Task, int) ([][]tasks.Task, error)) output.AlgoEfficiency {
	start := time.Now()
	taskLevels := assert(algo(records, rn))
	return output.AlgoEfficiency{
		PerformanceTime:   float64(time.Since(start).Microseconds()) / 1_000_000,
		ScheduleTotalTime: dsp.TotalTime(taskLevels),
		ScheduleDeviation: dsp.FnDeviation(taskLevels, rn),
		Schedule:          taskLevels,
	}
}

func assert[T any](res T, err error) T {
	if err != nil {
		log.Fatalln(err.Error())
	}
	return res
}

func invalidUsage(format string, v ...any) {
	log.Printf(format, v...)
	flag.Usage()
	os.Exit(1)
}
