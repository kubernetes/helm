/*
Copyright The Helm Authors.

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

func TestReleaseTesting(t *testing.T) {
	tests := []releaseCase{
		{
			name:      "basic test",
			args:      []string{"example-release"},
			flags:     []string{},
			responses: map[string]release.TestRun_Status{"PASSED: green lights everywhere": release.TestRun_SUCCESS},
			err:       false,
		},
		{
			name:      "test failure",
			args:      []string{"example-fail"},
			flags:     []string{},
			responses: map[string]release.TestRun_Status{"FAILURE: red lights everywhere": release.TestRun_FAILURE},
			err:       true,
		},
		{
			name:      "test unknown",
			args:      []string{"example-unknown"},
			flags:     []string{},
			responses: map[string]release.TestRun_Status{"UNKNOWN: yellow lights everywhere": release.TestRun_UNKNOWN},
			err:       false,
		},
		{
			name:      "test error",
			args:      []string{"example-error"},
			flags:     []string{},
			responses: map[string]release.TestRun_Status{"ERROR: yellow lights everywhere": release.TestRun_FAILURE},
			err:       true,
		},
		{
			name:      "test running",
			args:      []string{"example-running"},
			flags:     []string{},
			responses: map[string]release.TestRun_Status{"RUNNING: things are happening": release.TestRun_RUNNING},
			err:       false,
		},
		{
			name:  "multiple tests example",
			args:  []string{"example-suite"},
			flags: []string{},
			responses: map[string]release.TestRun_Status{
				"RUNNING: things are happening":           release.TestRun_RUNNING,
				"PASSED: party time":                          release.TestRun_SUCCESS,
				"RUNNING: things are happening again":         release.TestRun_RUNNING,
				"FAILURE: good thing u checked :)":            release.TestRun_FAILURE,
				"RUNNING: things are happening yet again": release.TestRun_RUNNING,
				"PASSED: feel free to party again":            release.TestRun_SUCCESS},
			err: true,
		},
	}

	runReleaseCases(t, tests, func(c *helm.FakeClient, out io.Writer) *cobra.Command {
		return newReleaseTestCmd(c, out)
	})
}
