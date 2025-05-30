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

package aip0164

import (
	"testing"

	"github.com/googleapis/api-linter/rules/internal/testutils"
	"github.com/jhump/protoreflect/desc/builder"
)

func TestRequestUnknownFields(t *testing.T) {
	// Set up the testing permutations.
	tests := []struct {
		testName    string
		messageName string
		fieldName   string
		fieldType   *builder.FieldType
		problems    testutils.Problems
	}{
		{"Etag", "UndeleteBookRequest", "etag", builder.FieldTypeString(), testutils.Problems{}},
		{"RequestId", "UndeleteBookRequest", "request_id", builder.FieldTypeString(), testutils.Problems{}},
		{"ValidateOnly", "UndeleteBookRequest", "validate_only", builder.FieldTypeBool(), testutils.Problems{}},
		{"Invalid", "UndeleteBookRequest", "application_id", builder.FieldTypeString(), testutils.Problems{{
			Message: "Unexpected field",
		}}},
		{"Irrelevant", "RemoveBookRequest", "application_id", builder.FieldTypeString(), testutils.Problems{}},
	}

	// Run each test individually.
	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			// Create an appropriate message descriptor.
			message, err := builder.NewMessage(test.messageName).AddField(
				builder.NewField("name", builder.FieldTypeString()),
			).AddField(
				builder.NewField(test.fieldName, test.fieldType),
			).Build()
			if err != nil {
				t.Fatalf("Could not build UndeleteBookRequest message.")
			}

			// Run the lint rule, and establish that it returns the correct problems.
			wantProblems := test.problems.SetDescriptor(message.FindFieldByName(test.fieldName))
			gotProblems := requestUnknownFields.Lint(message.GetFile())
			if diff := wantProblems.Diff(gotProblems); diff != "" {
				t.Error(diff)
			}
		})
	}
}
