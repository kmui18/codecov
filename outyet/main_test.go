/*
Copyright 2014 Google Inc.

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

package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// statusHandler is an http.Handler that writes an empty response using itself
// as the response status code.
type statusHandler int

func (h *statusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(int(*h))
}

func TestIsTagged(t *testing.T) {
	// Set up a fake "Google Code" web server reporting 404 not found.
	status := statusHandler(http.StatusNotFound)
	s := httptest.NewServer(&status)
	defer s.Close()

	if isTagged(s.URL) {
		t.Fatal("isTagged == true, want false")
	}

	// Change fake server status to 200 OK and try again.
	status = http.StatusOK

	if !isTagged(s.URL) {
		t.Fatal("isTagged == false, want true")
	}
}
