// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package internal_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/api/global/internal"
	"go.opentelemetry.io/otel/api/metric"
	metrictest "go.opentelemetry.io/otel/api/metric/metrictest"
	"go.opentelemetry.io/otel/label"
)

var Must = metric.Must

var asInt = metric.NewInt64Number
var asFloat = metric.NewFloat64Number

func TestDirect(t *testing.T) {
	internal.ResetForTest()

	ctx := context.Background()
	meter1 := global.Meter("test1", metric.WithInstrumentationVersion("semver:v1.0.0"))
	meter2 := global.Meter("test2")
	labels1 := []label.KeyValue{label.String("A", "B")}
	labels2 := []label.KeyValue{label.String("C", "D")}
	labels3 := []label.KeyValue{label.String("E", "F")}

	counter := Must(meter1).NewInt64Counter("test.counter")
	counter.Add(ctx, 1, labels1...)
	counter.Add(ctx, 1, labels1...)

	valuerecorder := Must(meter1).NewFloat64ValueRecorder("test.valuerecorder")
	valuerecorder.Record(ctx, 1, labels1...)
	valuerecorder.Record(ctx, 2, labels1...)

	_ = Must(meter1).NewFloat64ValueObserver("test.valueobserver.float", func(_ context.Context, result metric.Float64ObserverResult) {
		result.Observe(1., labels1...)
		result.Observe(2., labels2...)
	})

	_ = Must(meter1).NewInt64ValueObserver("test.valueobserver.int", func(_ context.Context, result metric.Int64ObserverResult) {
		result.Observe(1, labels1...)
		result.Observe(2, labels2...)
	})

	second := Must(meter2).NewFloat64ValueRecorder("test.second")
	second.Record(ctx, 1, labels3...)
	second.Record(ctx, 2, labels3...)

	mock, provider := metrictest.NewMeterProvider()
	global.SetMeterProvider(provider)

	counter.Add(ctx, 1, labels1...)
	valuerecorder.Record(ctx, 3, labels1...)
	second.Record(ctx, 3, labels3...)

	mock.RunAsyncInstruments()

	measurements := metrictest.AsStructs(mock.MeasurementBatches)

	require.EqualValues(t,
		[]metrictest.Measured{
			{
				Name:                   "test.counter",
				InstrumentationName:    "test1",
				InstrumentationVersion: "semver:v1.0.0",
				Labels:                 metrictest.LabelsToMap(labels1...),
				Number:                 asInt(1),
			},
			{
				Name:                   "test.valuerecorder",
				InstrumentationName:    "test1",
				InstrumentationVersion: "semver:v1.0.0",
				Labels:                 metrictest.LabelsToMap(labels1...),
				Number:                 asFloat(3),
			},
			{
				Name:                "test.second",
				InstrumentationName: "test2",
				Labels:              metrictest.LabelsToMap(labels3...),
				Number:              asFloat(3),
			},
			{
				Name:                   "test.valueobserver.float",
				InstrumentationName:    "test1",
				InstrumentationVersion: "semver:v1.0.0",
				Labels:                 metrictest.LabelsToMap(labels1...),
				Number:                 asFloat(1),
			},
			{
				Name:                   "test.valueobserver.float",
				InstrumentationName:    "test1",
				InstrumentationVersion: "semver:v1.0.0",
				Labels:                 metrictest.LabelsToMap(labels2...),
				Number:                 asFloat(2),
			},
			{
				Name:                   "test.valueobserver.int",
				InstrumentationName:    "test1",
				InstrumentationVersion: "semver:v1.0.0",
				Labels:                 metrictest.LabelsToMap(labels1...),
				Number:                 asInt(1),
			},
			{
				Name:                   "test.valueobserver.int",
				InstrumentationName:    "test1",
				InstrumentationVersion: "semver:v1.0.0",
				Labels:                 metrictest.LabelsToMap(labels2...),
				Number:                 asInt(2),
			},
		},
		measurements,
	)
}

func TestBound(t *testing.T) {
	internal.ResetForTest()

	// Note: this test uses opposite Float64/Int64 number kinds
	// vs. the above, to cover all the instruments.
	ctx := context.Background()
	glob := global.Meter("test")
	labels1 := []label.KeyValue{label.String("A", "B")}

	counter := Must(glob).NewFloat64Counter("test.counter")
	boundC := counter.Bind(labels1...)
	boundC.Add(ctx, 1)
	boundC.Add(ctx, 1)

	valuerecorder := Must(glob).NewInt64ValueRecorder("test.valuerecorder")
	boundM := valuerecorder.Bind(labels1...)
	boundM.Record(ctx, 1)
	boundM.Record(ctx, 2)

	mock, provider := metrictest.NewMeterProvider()
	global.SetMeterProvider(provider)

	boundC.Add(ctx, 1)
	boundM.Record(ctx, 3)

	require.EqualValues(t,
		[]metrictest.Measured{
			{
				Name:                "test.counter",
				InstrumentationName: "test",
				Labels:              metrictest.LabelsToMap(labels1...),
				Number:              asFloat(1),
			},
			{
				Name:                "test.valuerecorder",
				InstrumentationName: "test",
				Labels:              metrictest.LabelsToMap(labels1...),
				Number:              asInt(3),
			},
		},
		metrictest.AsStructs(mock.MeasurementBatches))

	boundC.Unbind()
	boundM.Unbind()
}

func TestUnbind(t *testing.T) {
	// Tests Unbind with SDK never installed.
	internal.ResetForTest()

	glob := global.Meter("test")
	labels1 := []label.KeyValue{label.String("A", "B")}

	counter := Must(glob).NewFloat64Counter("test.counter")
	boundC := counter.Bind(labels1...)

	valuerecorder := Must(glob).NewInt64ValueRecorder("test.valuerecorder")
	boundM := valuerecorder.Bind(labels1...)

	boundC.Unbind()
	boundM.Unbind()
}

func TestUnbindThenRecordOne(t *testing.T) {
	internal.ResetForTest()

	ctx := context.Background()
	mock, provider := metrictest.NewMeterProvider()

	meter := global.Meter("test")
	counter := Must(meter).NewInt64Counter("test.counter")
	boundC := counter.Bind()
	global.SetMeterProvider(provider)
	boundC.Unbind()

	require.NotPanics(t, func() {
		boundC.Add(ctx, 1)
	})
	require.Equal(t, 0, len(mock.MeasurementBatches))
}

type meterProviderWithConstructorError struct {
	metric.MeterProvider
}

type meterWithConstructorError struct {
	metric.MeterImpl
}

func (m *meterProviderWithConstructorError) Meter(iName string, opts ...metric.MeterOption) metric.Meter {
	return metric.WrapMeterImpl(&meterWithConstructorError{m.MeterProvider.Meter(iName, opts...).MeterImpl()}, iName, opts...)
}

func (m *meterWithConstructorError) NewSyncInstrument(_ metric.Descriptor) (metric.SyncImpl, error) {
	return metric.NoopSync{}, errors.New("constructor error")
}

func TestErrorInDeferredConstructor(t *testing.T) {
	internal.ResetForTest()

	ctx := context.Background()
	meter := global.MeterProvider().Meter("builtin")

	c1 := Must(meter).NewInt64Counter("test")
	c2 := Must(meter).NewInt64Counter("test")

	_, provider := metrictest.NewMeterProvider()
	sdk := &meterProviderWithConstructorError{provider}

	require.Panics(t, func() {
		global.SetMeterProvider(sdk)
	})

	c1.Add(ctx, 1)
	c2.Add(ctx, 2)
}

func TestImplementationIndirection(t *testing.T) {
	internal.ResetForTest()

	// Test that Implementation() does the proper indirection, i.e.,
	// returns the implementation interface not the global, after
	// registered.

	meter1 := global.Meter("test1")

	// Sync: no SDK yet
	counter := Must(meter1).NewInt64Counter("interface.counter")

	ival := counter.Measurement(1).SyncImpl().Implementation()
	require.NotNil(t, ival)

	_, ok := ival.(*metrictest.Sync)
	require.False(t, ok)

	// Async: no SDK yet
	valueobserver := Must(meter1).NewFloat64ValueObserver(
		"interface.valueobserver",
		func(_ context.Context, result metric.Float64ObserverResult) {},
	)

	ival = valueobserver.AsyncImpl().Implementation()
	require.NotNil(t, ival)

	_, ok = ival.(*metrictest.Async)
	require.False(t, ok)

	// Register the SDK
	_, provider := metrictest.NewMeterProvider()
	global.SetMeterProvider(provider)

	// Repeat the above tests

	// Sync
	ival = counter.Measurement(1).SyncImpl().Implementation()
	require.NotNil(t, ival)

	_, ok = ival.(*metrictest.Sync)
	require.True(t, ok)

	// Async
	ival = valueobserver.AsyncImpl().Implementation()
	require.NotNil(t, ival)

	_, ok = ival.(*metrictest.Async)
	require.True(t, ok)
}

func TestRecordBatchMock(t *testing.T) {
	internal.ResetForTest()

	meter := global.MeterProvider().Meter("builtin")

	counter := Must(meter).NewInt64Counter("test.counter")

	meter.RecordBatch(context.Background(), nil, counter.Measurement(1))

	mock, provider := metrictest.NewMeterProvider()
	global.SetMeterProvider(provider)

	meter.RecordBatch(context.Background(), nil, counter.Measurement(1))

	require.EqualValues(t,
		[]metrictest.Measured{
			{
				Name:                "test.counter",
				InstrumentationName: "builtin",
				Labels:              metrictest.LabelsToMap(),
				Number:              asInt(1),
			},
		},
		metrictest.AsStructs(mock.MeasurementBatches))
}
