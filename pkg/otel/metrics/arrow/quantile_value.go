package arrow

import (
	"fmt"

	"github.com/apache/arrow/go/v11/arrow"
	"github.com/apache/arrow/go/v11/arrow/array"
	"github.com/apache/arrow/go/v11/arrow/memory"
	"go.opentelemetry.io/collector/pdata/pmetric"

	"github.com/f5/otel-arrow-adapter/pkg/otel/constants"
)

// QuantileValueDT is the Arrow Data Type describing a quantile value.
var (
	QuantileValueDT = arrow.StructOf(
		arrow.Field{Name: constants.SUMMARY_QUANTILE, Type: arrow.PrimitiveTypes.Float64},
		arrow.Field{Name: constants.SUMMARY_VALUE, Type: arrow.PrimitiveTypes.Float64},
	)
)

// QuantileValueBuilder is a builder for a quantile value.
type QuantileValueBuilder struct {
	released bool

	builder *array.StructBuilder

	sqb *array.Float64Builder // summary quantile builder
	svb *array.Float64Builder // summary quantile value builder
}

// NewQuantileValueBuilder creates a new QuantileValueBuilder with a given memory allocator.
func NewQuantileValueBuilder(pool memory.Allocator) *QuantileValueBuilder {
	return QuantileValueBuilderFrom(array.NewStructBuilder(pool, QuantileValueDT))
}

// QuantileValueBuilderFrom creates a new QuantileValueBuilder from an existing StructBuilder.
func QuantileValueBuilderFrom(ndpb *array.StructBuilder) *QuantileValueBuilder {
	return &QuantileValueBuilder{
		released: false,
		builder:  ndpb,

		sqb: ndpb.FieldBuilder(0).(*array.Float64Builder),
		svb: ndpb.FieldBuilder(1).(*array.Float64Builder),
	}
}

// Build builds the underlying array.
//
// Once the array is no longer needed, Release() should be called to free the memory.
func (b *QuantileValueBuilder) Build() (*array.Struct, error) {
	if b.released {
		return nil, fmt.Errorf("QuantileValueBuilder: Build() called after Release()")
	}

	defer b.Release()
	return b.builder.NewStructArray(), nil
}

// Release releases the underlying memory.
func (b *QuantileValueBuilder) Release() {
	if b.released {
		return
	}

	b.released = true
	b.builder.Release()
}

// Append appends a new quantile value to the builder.
func (b *QuantileValueBuilder) Append(sdp pmetric.SummaryDataPointValueAtQuantile) error {
	if b.released {
		return fmt.Errorf("QuantileValueBuilder: Append() called after Release()")
	}

	b.builder.Append(true)
	b.sqb.Append(sdp.Quantile())
	b.svb.Append(sdp.Value())
	return nil
}