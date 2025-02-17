/*
Copyright The ORAS Authors.
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
	"runtime"
	"strings"

	"github.com/spf13/cobra"

	"oras.land/oras/internal/version"
)

func versionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Show the oras version information",
		Long: `Show the oras version information

Example - print version:
  oras version
`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runVersion()
		},
	}

	return cmd
}

func runVersion() error {
	items := [][]string{
		{"Version", version.GetVersion()},
		{"Go version", runtime.Version()},
	}
	if version.GitCommit != "" {
		items = append(items, []string{"Git commit", version.GitCommit})
	}
	if version.GitTreeState != "" {
		items = append(items, []string{"Git tree state", version.GitTreeState})
	}

	size := 0
	for _, item := range items {
		if length := len(item[0]); length > size {
			size = length
		}
	}
	for _, item := range items {
		fmt.Println(item[0] + ": " + strings.Repeat(" ", size-len(item[0])) + item[1])
	}

	return nil
}
