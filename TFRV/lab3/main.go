package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Belstowe/distrib-cs-1-autumn/lab3/stm"
)

var (
	n                  int
	c                  [3]int
	sm                 *stm.StrategyMatrix
	toOutputMatrix     bool
	toOutputStrategies bool
	toOutputGameCost   bool
	toOutputIterations bool
	toOutputTimeSpent  bool
	epsilon            float64
	measurements       map[string]interface{}
)

func main() {
	log.SetFlags(0)
	parseArgs()
	validateArgs()
	generateMatrix()
	iterateMatrix()
}

func parseArgs() {
	flag.BoolVar(&toOutputMatrix, "output-matrix", false, "Print C matrix")
	flag.BoolVar(&toOutputStrategies, "output-strategies", false, "Print the optimal strategies")
	flag.BoolVar(&toOutputGameCost, "output-gc", false, "Print the game cost result")
	flag.BoolVar(&toOutputIterations, "output-iterations", false, "Print num of algo iterations spent")
	flag.BoolVar(&toOutputTimeSpent, "output-time", false, "Print time spent (in seconds, accuracy up to microseconds)")
	flag.Float64Var(&epsilon, "eps", 0.01, "The desired precision of optimal game strategies")

	flag.Parse()

	progName, args := os.Args[0], flag.Args()
	if len(args) != 4 {
		log.Fatalf("usage:\n\t$ %s [-eps {e}] [-output-matrix] [-output-strategies] [-output-gc] [-output-iterations] [-output-time] {n} {c_1} {c_2} {c_3}\n", progName)
	}

	n = int(assert(strconv.ParseInt(args[0], 0, 32)))
	c[0] = int(assert(strconv.ParseInt(args[1], 0, 32)))
	c[1] = int(assert(strconv.ParseInt(args[2], 0, 32)))
	c[2] = int(assert(strconv.ParseInt(args[3], 0, 32)))
}

func validateArgs() {
	defects := make([]string, 0)
	if n <= 0 {
		defects = append(defects, fmt.Sprintf("  - n must be positive {got %d instead}", n))
	}
	if c[0] < 1 || c[0] > 3 {
		defects = append(defects, fmt.Sprintf("  - c_1 must be in range {1, 2, 3} (got %d instead)", c[0]))
	}
	if c[1] < 4 || c[1] > 6 {
		defects = append(defects, fmt.Sprintf("  - c_2 must be in range {4, 5, 6} (got %d instead)", c[1]))
	}
	if c[2] < 4 || c[2] > 6 {
		defects = append(defects, fmt.Sprintf("  - c_3 must be in range {4, 5, 6} (got %d instead)", c[2]))
	}
	if epsilon <= 0 || epsilon >= 1 {
		defects = append(defects, fmt.Sprintf("  - eps must be in range (0; 1) (got %f instead)", epsilon))
	}
	if len(defects) != 0 {
		log.Println("following defects in arguments were found:")
		log.Fatalln(strings.Join(defects, "\n"))
	}
}

func generateMatrix() {
	C := make([][]int, n)
	var max int = 0
	for i := range C {
		C[i] = make([]int, n)
		j := 0
		for ; j <= i; j++ {
			C[i][j] = j*c[0] + (i-j)*c[1]
			if C[i][j] > max {
				max = C[i][j]
			}
		}
		for ; j < n; j++ {
			C[i][j] = i*c[1] + (j-i)*c[2]
			if C[i][j] > max {
				max = C[i][j]
			}
		}
	}
	if toOutputMatrix {
		colWidth := getDigitNum(max)
		for _, v := range C {
			for _, n := range v {
				fmt.Printf("%*d ", colWidth, n)
			}
			fmt.Println()
		}
	}
	sm = stm.NewStrategyMatrix(C)
}

func iterateMatrix() {
	start := time.Now()
	sm.Iterate()
	for sm.VMax()-sm.VMin() > epsilon {
		sm.Iterate()
	}
	end := float64(time.Since(start).Microseconds()) / 1_000_000
	measurements = make(map[string]interface{})
	if toOutputStrategies {
		measurements["DispatcherStrategies"] = sm.P()
		measurements["DcsStrategies"] = sm.Q()
	}
	if toOutputGameCost {
		measurements["GameCostMin"] = sm.VMin()
		measurements["GameCostAvg"] = sm.VAvg()
		measurements["GameCostMax"] = sm.VMax()
	}
	if toOutputIterations {
		measurements["IterationsSpent"] = sm.Iterations()
	}
	if toOutputTimeSpent {
		measurements["TimeSpentSeconds"] = end
	}
	if len(measurements) != 0 {
		fmt.Println(string(assert(json.Marshal(measurements))))
	}
}

func getDigitNum(n int) int {
	digits := 0
	for ; n != 0; digits++ {
		n /= 10
	}
	return digits
}

func assert[T any](res T, err error) T {
	if err != nil {
		log.Fatalln(err.Error())
	}
	return res
}
