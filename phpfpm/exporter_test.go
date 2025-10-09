package phpfpm

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
)

func TestNewExporterWithConstLabels(t *testing.T) {
	labels := prometheus.Labels{"env": "prod"}

	exporter := NewExporter(PoolManager{}, labels)

	assert.Contains(t, exporter.up.String(), `constLabels: {env="prod"}`)
}
