package metrics

type MetricsCollector[T any] interface {
	init() error
	GetName() string
	CollectMetrics() (T, error)
}
