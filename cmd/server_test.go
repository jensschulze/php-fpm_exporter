package cmd

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/stretchr/testify/require"
)

func TestNewMetricsRegistry_SpecificMetricsPresence(t *testing.T) {
	// Test that specific well-known metrics from each collector are present/absent
	exporter := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "phpfpm_up",
		Help: "test PHP-FPM up metric",
	})
	exporter.Set(1)

	t.Run("with all collectors", func(t *testing.T) {
		registry := newMetricsRegistry(exporter, false)
		families, err := registry.Gather()
		require.NoError(t, err)

		metricNames := make(map[string]bool)
		for _, family := range families {
			metricNames[family.GetName()] = true
		}

		// Exporter metric
		require.True(t, metricNames["phpfpm_up"], "expected phpfpm_up metric")

		// Go collector metrics (sample)
		require.True(t, metricNames["go_goroutines"], "expected go_goroutines metric")
		require.True(t, metricNames["go_gc_duration_seconds"], "expected go_gc_duration_seconds metric")

		// Process collector metrics (sample)
		require.True(t, metricNames["process_cpu_seconds_total"], "expected process_cpu_seconds_total metric")
		require.True(t, metricNames["process_start_time_seconds"], "expected process_start_time_seconds metric")
	})

	t.Run("with phpfpm metrics only", func(t *testing.T) {
		registry := newMetricsRegistry(exporter, true)
		families, err := registry.Gather()
		require.NoError(t, err)

		metricNames := make(map[string]bool)
		for _, family := range families {
			metricNames[family.GetName()] = true
		}

		// Exporter metric should be present
		require.True(t, metricNames["phpfpm_up"], "expected phpfpm_up metric")

		// Go collector metrics should NOT be present
		require.False(t, metricNames["go_goroutines"], "expected go_goroutines metric to be absent")
		require.False(t, metricNames["go_gc_duration_seconds"], "expected go_gc_duration_seconds metric to be absent")

		// Process collector metrics should NOT be present
		require.False(t, metricNames["process_cpu_seconds_total"], "expected process_cpu_seconds_total metric to be absent")
		require.False(t, metricNames["process_start_time_seconds"], "expected process_start_time_seconds metric to be absent")
	})
}

func TestPromhttpMetricsHandler(t *testing.T) {
	tests := []struct {
		name           string
		phpfpmOnly     bool
		expectPromhttp bool
	}{
		{
			name:           "phpfpmMetricsOnly=false includes promhttp metrics",
			phpfpmOnly:     false,
			expectPromhttp: true,
		},
		{
			name:           "phpfpmMetricsOnly=true excludes promhttp metrics",
			phpfpmOnly:     true,
			expectPromhttp: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			exporter := prometheus.NewGauge(prometheus.GaugeOpts{
				Name: "phpfpm_test_metric",
				Help: "test PHP-FPM metric",
			})
			exporter.Set(1)

			registry := newMetricsRegistry(exporter, tc.phpfpmOnly)

			// Create the handler the same way as in serverCmd
			metricsHandler := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
			var handler http.Handler
			if tc.phpfpmOnly {
				handler = metricsHandler
			} else {
				handler = promhttp.InstrumentMetricHandler(registry, metricsHandler)
			}

			req := httptest.NewRequest(http.MethodGet, "/metrics", nil)

			// First request triggers promhttp counter registration.
			// The counter is incremented after the response is written,
			// so we need a second request to observe it.
			handler.ServeHTTP(httptest.NewRecorder(), req)

			// Second request should include the promhttp metrics
			rec := httptest.NewRecorder()
			handler.ServeHTTP(rec, req)

			require.Equal(t, http.StatusOK, rec.Code)

			body := rec.Body.String()
			hasPromhttp := strings.Contains(body, "promhttp_metric_handler_requests_total")

			require.Equal(t, tc.expectPromhttp, hasPromhttp,
				"promhttp_metric_handler_requests_total: expected=%v, got=%v", tc.expectPromhttp, hasPromhttp)
		})
	}
}
