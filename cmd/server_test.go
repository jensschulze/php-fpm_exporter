package cmd

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/require"
)

func TestParseConstLabels(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected prometheus.Labels
		err      string
	}{
		{
			name:     "empty input",
			input:    nil,
			expected: nil,
		},
		{
			name:  "single label",
			input: []string{"env=prod"},
			expected: prometheus.Labels{
				"env": "prod",
			},
		},
		{
			name:  "multiple labels with commas",
			input: []string{"team=core,service=api", "env=prod"},
			expected: prometheus.Labels{
				"team":    "core",
				"service": "api",
				"env":     "prod",
			},
		},
		{
			name:  "duplicate label key",
			input: []string{"env=prod", "env=staging"},
			err:   `duplicate label key "env"`,
		},
		{
			name:  "missing key",
			input: []string{"=prod"},
			err:   "label key cannot be empty",
		},
		{
			name:  "missing equals",
			input: []string{"env"},
			err:   "expected key=value",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			labels, err := parseConstLabels(tc.input)
			if tc.err != "" {
				require.EqualError(t, err, tc.err)
				require.Nil(t, labels)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tc.expected, labels)
		})
	}
}
