// Copyright 2022 Google LLC
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

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package recommendationengine_test

import (
	"context"

	recommendationengine "cloud.google.com/go/recommendationengine/apiv1beta1"
	"google.golang.org/api/iterator"
	recommendationenginepb "google.golang.org/genproto/googleapis/cloud/recommendationengine/v1beta1"
)

func ExampleNewPredictionClient() {
	ctx := context.Background()
	c, err := recommendationengine.NewPredictionClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExamplePredictionClient_Predict() {
	ctx := context.Background()
	c, err := recommendationengine.NewPredictionClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &recommendationenginepb.PredictRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/recommendationengine/v1beta1#PredictRequest.
	}
	it := c.Predict(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}
