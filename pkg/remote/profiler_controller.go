/*
 * Copyright 2022 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package remote

import (
	"context"
	"runtime/pprof"

	"github.com/cloudwego/kitex/pkg/profiler"
	"github.com/cloudwego/kitex/pkg/remote/transmeta"
)

const (
	typeTrace    = "trace"
	typeProfiler = "profiler"
)

type ProfilerController interface {
	Run(ctx context.Context) (err error)
	Stop()
	Tag(ctx context.Context, msg Message)
	Untag(ctx context.Context, msg Message)
}

var _ ProfilerController = (*profilerController)(nil)

func NewProfilerController(p profiler.Profiler) ProfilerController {
	return &profilerController{p: p}
}

type profilerController struct {
	p profiler.Profiler
}

func (c *profilerController) Run(ctx context.Context) (err error) {
	// tagging current goroutine
	// we could filter tag from pprof data to figure out that how much cost our profiler cause
	ctx = pprof.WithLabels(ctx, pprof.Labels("type", typeProfiler))
	return c.p.Run(ctx)
}

func (c *profilerController) Stop() {
	c.p.Stop()
}

func (c *profilerController) Tag(ctx context.Context, msg Message) {
	ti := msg.TransInfo()
	if ti == nil {
		return
	}
	c.p.Tag(ctx, []string{
		"type", typeTrace,
		"from", ti.TransIntInfo()[transmeta.FromService],
	})
}

func (c *profilerController) Untag(ctx context.Context, msg Message) {
	c.p.Untag(ctx)
}
