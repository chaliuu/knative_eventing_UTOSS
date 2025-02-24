/*
Copyright 2020 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	"testing"
)

func TestRequestReplyGetStatus(t *testing.T) {
	r := &RequestReply{
		Status: RequestReplyStatus{},
	}
	if got, want := r.GetStatus(), &r.Status.Status; got != want {
		t.Errorf("GetStatus=%v, want=%v", got, want)
	}
}

func TestRequestReply_GetGroupVersionKind(t *testing.T) {
	src := RequestReply{}
	gvk := src.GetGroupVersionKind()

	if gvk.Kind != "RequestReply" {
		t.Errorf("Should be RequestReply, got %s.", gvk.Kind)
	}
}
