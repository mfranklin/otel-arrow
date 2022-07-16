package trace

import (
	"github.com/apache/arrow/go/v9/arrow"
	coltracepb "otel-arrow-adapter/api/go.opentelemetry.io/proto/otlp/collector/trace/v1"
	"otel-arrow-adapter/pkg/otel/common"
	"otel-arrow-adapter/pkg/otel/constants"
	"otel-arrow-adapter/pkg/rbb"
)

func OtlpTraceToArrowEvents(rbr *rbb.RecordBatchRepository, request *coltracepb.ExportTraceServiceRequest) ([]arrow.Record, error) {
	for _, resourceSpans := range request.ResourceSpans {
		for _, scopeSpans := range resourceSpans.ScopeSpans {
			for _, span := range scopeSpans.Spans {
				record := rbb.NewRecord()

				if span.StartTimeUnixNano > 0 {
					record.U64Field(constants.START_TIME_UNIX_NANO, span.StartTimeUnixNano)
				}
				if span.EndTimeUnixNano > 0 {
					record.U64Field(constants.END_TIME_UNIX_NANO, span.EndTimeUnixNano)
				}
				common.AddResource(record, resourceSpans.Resource)
				common.AddScope(record, constants.SCOPE_SPANS, scopeSpans.Scope)

				if span.TraceId != nil && len(span.TraceId) > 0 {
					record.BinaryField(constants.TRACE_ID, span.TraceId)
				}
				if span.SpanId != nil && len(span.SpanId) > 0 {
					record.BinaryField(constants.SPAN_ID, span.SpanId)
				}
				if len(span.TraceState) > 0 {
					record.StringField(constants.TRACE_STATE, span.TraceState)
				}
				if span.ParentSpanId != nil && len(span.ParentSpanId) > 0 {
					record.BinaryField(constants.PARENT_SPAN_ID, span.SpanId)
				}
				if len(span.Name) > 0 {
					record.StringField(constants.NAME, span.Name)
				}
				record.I32Field(constants.KIND, int32(span.Kind))
				attributes := common.NewAttributes(span.Attributes)
				if attributes != nil {
					record.AddField(attributes)
				}

				if span.DroppedAttributesCount > 0 {
					record.U32Field(constants.DROPPED_ATTRIBUTES_COUNT, uint32(span.DroppedAttributesCount))
				}

				// ToDo Events
				// ToDo DroppedEventsCount
				// ToDo Links
				// ToDo DroppedLinksCount

				if span.Status != nil {
					record.I32Field(constants.STATUS, int32(span.Status.Code))
					record.StringField(constants.STATUS_MESSAGE, span.Status.Message)
				}

				rbr.AddRecord(record)
			}
		}
	}

	logsRecords, err := rbr.Build()
	if err != nil {
		return nil, err
	}

	result := make([]arrow.Record, 0, len(logsRecords))
	for _, record := range logsRecords {
		result = append(result, record)
	}

	return result, nil
}