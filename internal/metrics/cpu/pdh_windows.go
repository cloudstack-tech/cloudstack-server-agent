//go:build windows

package metrics

import "github.com/rokukoo/win_perf_counters"

const defaultMaxBufferSize = 100 * 1024 * 1024

var performanceQuery = win_perf_counters.MustNewOpenPerformanceQuery(defaultMaxBufferSize)

func init() {
	performanceQuery.CollectData()
}