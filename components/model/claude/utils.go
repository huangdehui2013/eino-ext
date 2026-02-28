/*
 * Copyright 2025 CloudWeGo Authors
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

package claude

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

var util utilHelper

type utilHelper struct{}

func (utilHelper) jsonStr(v any) string {
	b, err := json.Marshal(v)
	if err != nil {
		return fmt.Sprintf("%+v", v)
	}

	return string(b)
}

func (utilHelper) traceID(ctx context.Context) string {
	if ctx != nil {
		traceIDKeys := []string{"trace_id", "traceId", "traceID", "request_id", "requestId", "x-trace-id", "x-request-id"}
		for _, k := range traceIDKeys {
			if v := ctx.Value(k); v != nil {
				return fmt.Sprintf("%v", v)
			}
		}
	}

	return fmt.Sprintf("local-%d", time.Now().UnixNano())
}

func of[T any](v T) *T {
	return &v
}

func from[T any](v *T) T {
	if v == nil {
		var t T
		return t
	}

	return *v
}

func fromOrDefault[T any](v *T, d T) T {
	if v == nil {
		return d
	}

	return *v
}
