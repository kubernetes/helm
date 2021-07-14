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

package chartutil

import (
	"log"
	"strings"

	"helm.sh/helm/v3/pkg/chart"
)

// processDependencyConditions disables charts based on condition path value in values
func processDependencyConditions(reqs []*chart.Dependency, cvals Values) {
	if reqs == nil {
		return
	}
	for _, r := range reqs {
		for _, c := range strings.Split(strings.TrimSpace(r.Condition), ",") {
			if len(c) > 0 {
				// retrieve value
				vv, err := cvals.PathValue(c)
				if err == nil {
					// if not bool, warn
					if bv, ok := vv.(bool); ok {
						r.Enabled = bv
						break
					} else {
						log.Printf("Warning: Condition path '%s' for chart %s returned non-bool value", c, r.Name)
					}
				} else if _, ok := err.(ErrNoValue); !ok {
					// this is a real error
					log.Printf("Warning: PathValue returned error %v", err)
				}
			}
		}
	}
}

// processDependencyTags disables charts based on tags in values
func processDependencyTags(reqs []*chart.Dependency, tags map[string]interface{}) {
	if reqs == nil || tags == nil {
		return
	}
	for _, r := range reqs {
		var hasTrue, hasFalse bool
		for _, k := range r.Tags {
			if b, ok := tags[k]; ok {
				// if not bool, warn
				if bv, ok := b.(bool); ok {
					if bv {
						hasTrue = true
					} else {
						hasFalse = true
					}
				} else {
					log.Printf("Warning: Tag '%s' for chart %s returned non-bool value", k, r.Name)
				}
			}
		}
		if !hasTrue && hasFalse {
			r.Enabled = false
		} else if hasTrue || !hasTrue && !hasFalse {
			r.Enabled = true
		}
	}
}

func GetTags(cvals Values) map[string]interface{} {
	vt, err := cvals.Table("tags")
	if err != nil {
		return nil
	}
	return vt
}

func getAliasDependency(charts []*chart.Chart, dep *chart.Dependency) *chart.Chart {
	for _, c := range charts {
		if c == nil {
			continue
		}
		if c.Name() != dep.Name {
			continue
		}
		if !IsCompatibleRange(dep.Version, c.Metadata.Version) {
			continue
		}

		out := *c
		md := *c.Metadata
		out.Metadata = &md

		if dep.Alias != "" {
			md.Name = dep.Alias
		}
		return &out
	}
	return nil
}

// ProcessDependencyEnabled removes disabled charts from dependencies
func ProcessDependencyEnabled(c *chart.Chart, v map[string]interface{}, tags map[string]interface{}) error {
	if c.Metadata.Dependencies == nil {
		return nil
	}

	var chartDependencies []*chart.Chart
	// If any dependency is not a part of Chart.yaml
	// then this should be added to chartDependencies.
	// However, if the dependency is already specified in Chart.yaml
	// we should not add it, as it would be anyways processed from Chart.yaml

Loop:
	for _, existing := range c.Dependencies() {
		for _, req := range c.Metadata.Dependencies {
			if existing.Name() == req.Name && IsCompatibleRange(req.Version, existing.Metadata.Version) {
				continue Loop
			}
		}
		chartDependencies = append(chartDependencies, existing)
	}

	for _, req := range c.Metadata.Dependencies {
		if chartDependency := getAliasDependency(c.Dependencies(), req); chartDependency != nil {
			chartDependencies = append(chartDependencies, chartDependency)
		}
		if req.Alias != "" {
			req.Name = req.Alias
		}
	}
	c.SetDependencies(chartDependencies...)

	// set all to true
	for _, lr := range c.Metadata.Dependencies {
		lr.Enabled = true
	}
	// flag dependencies as enabled/disabled
	processDependencyTags(c.Metadata.Dependencies, tags)
	processDependencyConditions(c.Metadata.Dependencies, v)
	// make a map of charts to remove
	rm := map[string]struct{}{}
	for _, r := range c.Metadata.Dependencies {
		if !r.Enabled {
			// remove disabled chart
			rm[r.Name] = struct{}{}
		}
	}
	// don't keep disabled charts in new slice
	cd := []*chart.Chart{}
	copy(cd, c.Dependencies()[:0])
	for _, n := range c.Dependencies() {
		if _, ok := rm[n.Metadata.Name]; !ok {
			cd = append(cd, n)
		}
	}
	// don't keep disabled charts in metadata
	cdMetadata := []*chart.Dependency{}
	copy(cdMetadata, c.Metadata.Dependencies[:0])
	for _, n := range c.Metadata.Dependencies {
		if _, ok := rm[n.Name]; !ok {
			cdMetadata = append(cdMetadata, n)
		}
	}
	// set the correct dependencies in metadata
	c.Metadata.Dependencies = nil
	c.Metadata.Dependencies = append(c.Metadata.Dependencies, cdMetadata...)
	c.SetDependencies(cd...)

	return nil
}

// pathToMap creates a nested map given a YAML path in dot notation.
func pathToMap(path string, data map[string]interface{}) map[string]interface{} {
	if path == "." {
		return data
	}
	return set(parsePath(path), data)
}

func set(path []string, data map[string]interface{}) map[string]interface{} {
	if len(path) == 0 {
		return nil
	}
	cur := data
	for i := len(path) - 1; i >= 0; i-- {
		cur = map[string]interface{}{path[i]: cur}
	}
	return cur
}

// processImportValues merges values from child to parent based on the chart's dependencies' ImportValues field.
func processImportValues(c *chart.Chart, cvals Values) error {
	if c.Metadata.Dependencies == nil {
		return nil
	}

	// import values from each dependency if specified in import-values
	for _, r := range c.Metadata.Dependencies {
		for _, riv := range r.ImportValues {
			var child, parent string
			switch iv := riv.(type) {
			case map[string]interface{}:
				child = iv["child"].(string)
				parent = iv["parent"].(string)
			case string:
				child = "exports." + iv
				parent = "."
			}
			// get child table
			vv, err := cvals.Table(r.Name + "." + child)
			if err != nil {
				log.Printf("Warning: ImportValues missing table %s from chart %s: %v", child, r.Name, err)
				continue
			}
			// create value map from child to be merged into parent
			CoalesceTables(cvals, pathToMap(parent, vv))
		}
	}

	return nil
}

// ProcessDependencyImportValues imports specified chart values from child to parent.
//
// v is expected to have existing path for every sub chart
func ProcessDependencyImportValues(c *chart.Chart, v map[string]interface{}) error {
	for _, d := range c.Dependencies() {
		// recurse
		dv := v[d.Name()].(map[string]interface{})
		if err := ProcessDependencyImportValues(d, dv); err != nil {
			return err
		}
	}
	return processImportValues(c, v)
}
