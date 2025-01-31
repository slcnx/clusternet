/*
Copyright 2021 The Clusternet Authors.

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
	goflag "flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/spf13/pflag"
	"k8s.io/component-base/logs"

	"github.com/clusternet/clusternet/cmd/clusternet-hub/app"
	"github.com/clusternet/clusternet/pkg/utils"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	logs.InitLogs()
	defer logs.FlushLogs()

	ctx := utils.GracefulStopWithContext()
	command := app.NewClusternetHubCmd(ctx)
	pflag.CommandLine.SetNormalizeFunc(utils.WordSepNormalizeFunc)
	pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)

	if err := command.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
