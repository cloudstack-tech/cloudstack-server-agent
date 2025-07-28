package metrics

var (
	processorUtilityHandle = performanceQuery.MustAddCounterToQuery("\\Processor Information(_Total)\\% Processor Utility")
)

type CpuUsageTotalCollector struct {
}

func (c *CpuUsageTotalCollector) init() error {
	performanceQuery.CollectData()
	return nil
}

func NewCpuUsageTotalCollector() (*CpuUsageTotalCollector, error) {
	collector := &CpuUsageTotalCollector{}
	err := collector.init()
	if err != nil {
		return nil, err
	}
	return collector, nil
}

func (c *CpuUsageTotalCollector) GetName() string {
	return "cpu_usage_total"
}

func (c *CpuUsageTotalCollector) CollectMetrics() (float64, error) {
	performanceQuery.CollectData()

	usage, err := performanceQuery.GetFormattedCounterValueDouble(processorUtilityHandle)
	if err != nil {
		return 0, err
	}

	if usage > 100 {
		usage = 100
	}

	return usage, nil
}