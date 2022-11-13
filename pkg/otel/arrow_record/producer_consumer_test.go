package arrow_record

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/pdata/plog/plogotlp"
	"go.opentelemetry.io/collector/pdata/pmetric/pmetricotlp"
	"go.opentelemetry.io/collector/pdata/ptrace/ptraceotlp"

	arrowpb "github.com/f5/otel-arrow-adapter/api/collector/arrow/v1"
	"github.com/f5/otel-arrow-adapter/pkg/datagen"
	"github.com/f5/otel-arrow-adapter/pkg/otel/assert"
)

func TestProducerConsumerTraces(t *testing.T) {
	dg := datagen.NewTracesGenerator(
		datagen.DefaultResourceAttributes(),
		datagen.DefaultInstrumentationScopes(),
	)
	traces := dg.Generate(10, time.Minute)

	producer := NewProducer()

	batch, err := producer.BatchArrowRecordsFromTraces(traces)
	require.NoError(t, err)
	require.Equal(t, arrowpb.OtlpArrowPayloadType_SPANS, batch.OtlpArrowPayloads[0].Type)

	consumer := NewConsumer()
	received, err := consumer.TracesFrom(batch)
	require.Equal(t, 1, len(received))

	assert.Equiv(
		t,
		[]json.Marshaler{ptraceotlp.NewExportRequestFromTraces(traces)},
		[]json.Marshaler{ptraceotlp.NewExportRequestFromTraces(received[0])},
	)
}

func TestProducerConsumerLogs(t *testing.T) {
	dg := datagen.NewLogsGenerator(
		datagen.DefaultResourceAttributes(),
		datagen.DefaultInstrumentationScopes(),
	)
	logs := dg.Generate(10, time.Minute)

	producer := NewProducer()

	batch, err := producer.BatchArrowRecordsFromLogs(logs)
	require.NoError(t, err)
	require.Equal(t, arrowpb.OtlpArrowPayloadType_LOGS, batch.OtlpArrowPayloads[0].Type)

	consumer := NewConsumer()
	received, err := consumer.LogsFrom(batch)
	require.Equal(t, 1, len(received))

	assert.Equiv(
		t,
		[]json.Marshaler{plogotlp.NewExportRequestFromLogs(logs)},
		[]json.Marshaler{plogotlp.NewExportRequestFromLogs(received[0])},
	)
}

func TestProducerConsumerMetrics(t *testing.T) {
	dg := datagen.NewMetricsGenerator(
		datagen.DefaultResourceAttributes(),
		datagen.DefaultInstrumentationScopes(),
	)
	metrics := dg.Generate(10, time.Minute)

	producer := NewProducer()

	batch, err := producer.BatchArrowRecordsFromMetrics(metrics)
	require.NoError(t, err)
	require.Equal(t, arrowpb.OtlpArrowPayloadType_METRICS, batch.OtlpArrowPayloads[0].Type)

	consumer := NewConsumer()
	received, err := consumer.MetricsFrom(batch)
	require.Equal(t, 1, len(received))

	assert.Equiv(
		t,
		[]json.Marshaler{pmetricotlp.NewExportRequestFromMetrics(metrics)},
		[]json.Marshaler{pmetricotlp.NewExportRequestFromMetrics(received[0])},
	)
}
