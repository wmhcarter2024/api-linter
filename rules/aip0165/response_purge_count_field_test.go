// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package aip0165

import (
	"strings"
	"testing"

	"github.com/googleapis/api-linter/rules/internal/testutils"
	"github.com/jhump/protoreflect/desc"
)

func TestResponsePurgeCountField(t *testing.T) {
	for _, test := range []struct {
		name     string
		RPC      string
		Field    string
		problems testutils.Problems
	}{
		{"Valid", "PurgeBooks", "int32 purge_count = 1;", nil},
		{"Missing", "PurgeBooks", "", testutils.Problems{{Message: "no `purge_count`"}}},
		{"InvalidType", "PurgeBooks", "int64 purge_count = 1;", testutils.Problems{{Suggestion: "int32"}}},
		{"InvalidTypeRepeated", "PurgeBooks", "repeated int32 purge_count = 1;", testutils.Problems{{Suggestion: "int32"}}},
		{"IrrelevantMessage", "EnumerateBooks", "", nil},
	} {
		t.Run(test.name, func(t *testing.T) {
			f := testutils.ParseProto3Tmpl(t, `
				message {{.RPC}}Response {
					{{.Field}}
					repeated string names = 2;
				}
			`, test)
			var d desc.Descriptor = f.GetMessageTypes()[0]
			if strings.HasPrefix(test.name, "InvalidType") {
				d = f.GetMessageTypes()[0].GetFields()[0]
			}
			if diff := test.problems.SetDescriptor(d).Diff(responsePurgeCountField.Lint(f)); diff != "" {
				t.Error(diff)
			}
		})
	}
}
