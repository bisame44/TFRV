package stm

func minValIndex(arr []int) int {
	i := 0
	for j := range arr {
		if arr[j] < arr[i] {
			i = j
		}
	}
	return i
}

func maxValIndex(arr []int) int {
	i := 0
	for j := range arr {
		if arr[j] > arr[i] {
			i = j
		}
	}
	return i
}

type StrategyMatrix struct {
	c [][]int
	i int
	a []int
	b []int
	p []int
	q []int
}

func NewStrategyMatrix(c [][]int) *StrategyMatrix {
	return &StrategyMatrix{c, 0, make([]int, len(c)), make([]int, len(c)), make([]int, len(c)), make([]int, len(c))}
}

func (sm *StrategyMatrix) iterateB() {
	i := maxValIndex(sm.a)
	for j, val := range sm.c[i] {
		sm.b[j] += val
	}
	sm.q[i]++
}

func (sm *StrategyMatrix) iterateA() {
	i := minValIndex(sm.b)
	for row := range sm.c {
		sm.a[row] += sm.c[row][i]
	}
	sm.p[i]++
}

func (sm StrategyMatrix) VMin() float64 {
	min := sm.b[0]
	for _, val := range sm.b {
		if min > val {
			min = val
		}
	}
	return float64(min) / float64(sm.i)
}

func (sm StrategyMatrix) VMax() float64 {
	max := sm.a[0]
	for _, val := range sm.a {
		if max < val {
			max = val
		}
	}
	return float64(max) / float64(sm.i)
}

func (sm StrategyMatrix) VAvg() float64 {
	return (sm.VMax() + sm.VMin()) / 2
}

func (sm StrategyMatrix) Iterations() int {
	return sm.i
}

func (sm StrategyMatrix) P() []float64 {
	p := make([]float64, 0, len(sm.p))
	for _, pv := range sm.p {
		p = append(p, float64(pv)/float64(sm.i))
	}
	return p
}

func (sm StrategyMatrix) Q() []float64 {
	q := make([]float64, 0, len(sm.q))
	for _, qv := range sm.q {
		q = append(q, float64(qv)/float64(sm.i))
	}
	return q
}

func (sm *StrategyMatrix) Iterate() {
	sm.i++
	sm.iterateB()
	sm.iterateA()
}
