/*
Copyright 2016 The Kubernetes Authors All rights reserved.

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
	"io"
	"testing"

	"github.com/spf13/cobra"

	"k8s.io/helm/pkg/helm"
	"k8s.io/helm/pkg/proto/hapi/release"
)

func TestGetValuesCmd(t *testing.T) {
	tests := []releaseCase{
		{
			name:     "get values with a release",
			resp:     helm.ReleaseMock(&helm.MockReleaseOptions{Name: "thomas-guide"}),
			args:     []string{"thomas-guide"},
			expected: "name: \"value\"",
			rels:     []*release.Release{helm.ReleaseMock(&helm.MockReleaseOptions{Name: "thomas-guide"})},
		},
		{
			name: "get values requires release name arg",
			err:  true,
		},
	}
	cmd := func(c *helm.FakeClient, out io.Writer) *cobra.Command {
		return newGetValuesCmd(c, out)
	}
	runReleaseCases(t, tests, cmd)
}

func TestGetValuesCmdJson(t *testing.T) {
	tests := []releaseCase{
		{
			name:     "get values with a release in JSON",
			resp:     helm.ReleaseMock(&helm.MockReleaseOptions{Name: "thomas-guide"}),
			flags:    []string{"--output", "json"},
			args:     []string{"thomas-guide"},
			expected: "{\n    \"name\": \"value\"\n}\n",
			rels:     []*release.Release{helm.ReleaseMock(&helm.MockReleaseOptions{Name: "thomas-guide"})},
		},
		{
			name: "get values requires release name arg",
			err:  true,
		},
	}
	cmd := func(c *helm.FakeClient, out io.Writer) *cobra.Command {
		return newGetValuesCmd(c, out)
	}
	runReleaseCases(t, tests, cmd)
}
