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
	"fmt"
	"io"
	"sync"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"helm.sh/helm/v3/cmd/helm/require"
	"helm.sh/helm/v3/pkg/getter"
	"helm.sh/helm/v3/pkg/repo"
)

const validateDesc = `
Validate gets the latest information about charts from the respective chart repositories.
Details are enumerated for any charts that fail validation.
`

var errNoRepositoriesToValidate = errors.New("no repositories found. You must add one before validating")

type repoValidateOptions struct {
	validate  func([]*repo.ChartRepository, io.Writer)
	repoFile  string
	repoCache string
}

func newRepoValidateCmd(out io.Writer) *cobra.Command {
	o := &repoValidateOptions{validate: validateCharts}

	cmd := &cobra.Command{
		Use:               "validate",
		Aliases:           []string{},
		Short:             "validate information of available charts locally from chart repositories",
		Long:              validateDesc,
		Args:              require.NoArgs,
		ValidArgsFunction: noCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			o.repoFile = settings.RepositoryConfig
			o.repoCache = settings.RepositoryCache
			return o.run(out)
		},
	}
	return cmd
}

func (o *repoValidateOptions) run(out io.Writer) error {
	f, err := repo.LoadFile(o.repoFile)
	switch {
	case isNotExist(err):
		return errNoRepositoriesToValidate
	case err != nil:
		return errors.Wrapf(err, "failed loading file: %s", o.repoFile)
	case len(f.Repositories) == 0:
		return errNoRepositoriesToValidate
	}

	var repos []*repo.ChartRepository
	for _, cfg := range f.Repositories {
		r, err := repo.NewChartRepository(cfg, getter.All(settings))
		if err != nil {
			return err
		}
		if o.repoCache != "" {
			r.CachePath = o.repoCache
		}
		repos = append(repos, r)
	}

	o.validate(repos, out)
	return nil
}

func validateCharts(repos []*repo.ChartRepository, out io.Writer) {
	fmt.Fprintln(out, "Hang tight while we grab the latest from your chart repositories...")
	var wg sync.WaitGroup
	for _, re := range repos {
		wg.Add(1)
		go func(re *repo.ChartRepository) {
			defer wg.Done()
			if _, err := re.ValidateAndDownloadIndexFile(); err != nil {
				fmt.Fprintf(out, "...Unable to validate the %q chart repository (%s):\n\t%s\n", re.Config.Name, re.Config.URL, err)
			} else {
				fmt.Fprintf(out, "...Successfully validated the %q chart repository\n", re.Config.Name)
			}
		}(re)
	}
	wg.Wait()
	fmt.Fprintln(out, "Validation Complete. ⎈Happy Helming!⎈")
}
