// Copyright 2018 The Kura Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package blobserver

import (
	"fmt"

	"github.com/kurafs/kura/pkg/cli"
)

var BlobServerCmd = &cli.Command{
	Run:       blobServerCmdRun,
	UsageLine: "blob-server [-f] [-a arg]",
	Short:     "blob-server command overview",
	Long: `
Blob server detailed overview.
    `,
}

func blobServerCmdRun(cmd *cli.Command, args []string) error {
	f := cmd.FlagSet.Bool("f", false, "Flag usage")
	a := cmd.FlagSet.String("a", "", "Argument parameter usage")
	if err := cmd.FlagSet.Parse(args); err != nil {
		return cli.CmdParseError(err)
	}

	fmt.Println(fmt.Sprintf("%s: parsed successfully: %t, %s", cmd.Name(), *f, *a))
	return nil
}