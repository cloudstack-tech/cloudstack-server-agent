package metrics

var (
	processorUtilityHandle = performanceQuery.MustAddCounterToQuery("\\Processor Information(_Total)\\% Processor Utility")
)

type CpuUsageCollector struct {
}

func (c *CpuUsageCollector) init() error {
	performanceQuery.CollectData()
	return nil
}

func NewCpuUsageCollector() *CpuUsageCollector {
	collector := &CpuUsageCollector{}
	collector.init()
	return collector
}

func (c *CpuUsageCollector) GetName() string {
	return "cpu_usage"
}

func (c *CpuUsageCollector) CollectMetrics() (float64, error) {
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