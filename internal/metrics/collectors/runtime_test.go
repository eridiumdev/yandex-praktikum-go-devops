package collectors

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"eridiumdev/yandex-praktikum-go-devops/internal/metrics/domain"
)

func TestRuntimeCollect(t *testing.T) {
	col := NewRuntimeCollector("runtime")
	snapshot, err := col.Collect(context.Background())

	require.NoError(t, err)
	assert.Greater(t, len(snapshot), 0)

	for _, m := range snapshot {
		if m.Name == domain.Alloc {
			assert.Equal(t, domain.TypeGauge, m.Type)
			assert.NotNil(t, m.Gauge)
			break
		}
	}
}
