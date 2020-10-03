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
	"testing"

	"github.com/stretchr/testify/assert"

	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/api/global/internal"
	"go.opentelemetry.io/otel/api/trace/tracetest"
)

func TestTraceWithSDK(t *testing.T) {
	internal.ResetForTest()

	ctx := context.Background()
	gtp := global.TracerProvider()
	tracer1 := gtp.Tracer("pre")
	// This is started before an SDK was registered and should be dropped.
	_, span1 := tracer1.Start(ctx, "span1")

	sr := new(tracetest.StandardSpanRecorder)
	tp := tracetest.NewTracerProvider(tracetest.WithSpanRecorder(sr))
	global.SetTracerProvider(tp)

	// This span was started before initialization, it is expected to be dropped.
	span1.End()

	// The existing Tracer should have been configured to now use the configured SDK.
	_, span2 := tracer1.Start(ctx, "span2")
	span2.End()

	// The global TracerProvider should now create Tracers that also use the newly configured SDK.
	tracer2 := gtp.Tracer("post")
	_, span3 := tracer2.Start(ctx, "span3")
	span3.End()

	filterNames := func(spans []*tracetest.Span) []string {
		names := make([]string, len(spans))
		for i := range spans {
			names[i] = spans[i].Name()
		}
		return names
	}
	expected := []string{"span2", "span3"}
	assert.ElementsMatch(t, expected, filterNames(sr.Started()))
	assert.ElementsMatch(t, expected, filterNames(sr.Completed()))
}
