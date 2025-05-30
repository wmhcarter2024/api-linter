// Copyright 2019 Google LLC
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

package aip0136

import (
	"testing"

	"github.com/googleapis/api-linter/rules/internal/testutils"
)

func TestVerbNoun(t *testing.T) {
	tests := []struct {
		MethodName string
		problems   testutils.Problems
	}{
		{"ArchiveBook", testutils.Problems{}},
		{"Archive", testutils.Problems{{Message: "verb followed by a noun"}}},
	}
	for _, test := range tests {
		t.Run(test.MethodName, func(t *testing.T) {
			file := testutils.ParseProto3Tmpl(t, `
				service Library {
					rpc {{.MethodName}}({{.MethodName}}Request) returns ({{.MethodName}}Response);
				}
				message {{.MethodName}}Request {}
				message {{.MethodName}}Response {}
			`, test)
			method := file.GetServices()[0].GetMethods()[0]
			got := verbNoun.Lint(file)
			if diff := test.problems.SetDescriptor(method).Diff(got); diff != "" {
				t.Error(diff)
			}
		})
	}
}
