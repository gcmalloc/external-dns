/*
Copyright 2017 The Kubernetes Authors.

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

package source

import "github.com/kubernetes-incubator/external-dns/endpoint"

// mockSource returns mock endpoints.
type mockSource struct {
	store []*endpoint.Endpoint
}

// NewMockSource creates a new mockSource returning the given endpoints.
func NewMockSource(endpoints []*endpoint.Endpoint) Source {
	return &mockSource{store: endpoints}
}

// Endpoints returns the desired mock endpoints.
func (s *mockSource) Endpoints() ([]*endpoint.Endpoint, error) {
	return s.store, nil
}
