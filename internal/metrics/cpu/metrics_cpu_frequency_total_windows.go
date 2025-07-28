//go:build windows

package metrics

var (
	processorFrequencyHandle = performanceQuery.MustAddCounterToQuery("\\Processor Information(_Total)\\Actual Frequency")
)

type CpuFrequencyCollector struct {
}

func (c *CpuFrequencyCollector) init() error {
	performanceQuery.CollectData()
	return nil
}

func NewCpuFrequencyCollector() (*CpuFrequencyCollector, error) {
	collector := &CpuFrequencyCollector{}
	err := collector.init()
	if err != nil {
		return nil, err
	}
	return collector, nil
}

func (c *CpuFrequencyCollector) GetName() string {
	return "cpu_frequency_total"
}

func (c *CpuFrequencyCollector) CollectMetrics() (float64, error) {
	performanceQuery.CollectData()

	usage, err := performanceQuery.GetFormattedCounterValueDouble(processorFrequencyHandle)
	if err != nil {
		return 0, err
	}

	return usage, nil
}