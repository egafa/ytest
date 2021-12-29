package model

type Metric interface {
	SaveGaugeVal(nameMetric string, value float64)
	GetGaugeVal(nameMetric string) float64
	SaveCounterVal(nameMetric string, value int64)
	GetCounterVal(nameMetric string, num int64) int64
}

type MapMetric struct {
	GaugeData   map[string]float64
	CounterData map[string][]int64
}

func (m MapMetric) SaveGaugeVal(nameMetric string, value float64) {
	m.GaugeData[nameMetric] = value
}

func (m MapMetric) GetGaugeVal(nameMetric string) float64 {
	return m.GaugeData[nameMetric]
}

func (m MapMetric) SaveCounterVal(nameMetric string, value int64) {
	var v []int64

	v, ok := m.CounterData[nameMetric]
	if ok {
		m.CounterData[nameMetric] = append(m.CounterData[nameMetric], value)
	}

	v = append(v, value)
	m.CounterData[nameMetric] = v
}

func (m MapMetric) GetCounterVal(nameMetric string, num int64) int64 {
	return 0
}
