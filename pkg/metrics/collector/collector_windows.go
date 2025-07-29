package collector

import (
	"github.com/rokukoo/win_perf_counters"
)

const defaultMaxBufferSize = 100 * 1024 * 1024

var PerformanceQuery = win_perf_counters.MustNewOpenPerformanceQuery(defaultMaxBufferSize)

func init() {
	PerformanceQuery.CollectData()
}
