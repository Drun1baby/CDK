/*
Copyright 2022 The Authors of https://github.com/CDK-TEAM/CDK .

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

package plugin

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
)

type ExploitInterface interface {
	Desc() string
	Run() bool
	GetExploitType() string
}

type TaskInterface interface {
	Exec() bool
	Desc() string
}

var Exploits map[string]ExploitInterface
var Tasks map[string]TaskInterface

func init() {
	Exploits = make(map[string]ExploitInterface)
	Tasks = make(map[string]TaskInterface)
}

func ListAllExploit() {

	writer := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)

	type kv struct {
		Name        string
		ExploitType string
		Desc        string
	}

	sortedExploits := make([]kv, 0)

	for name, plugin := range Exploits {
		sortedExploits = append(sortedExploits, kv{name, plugin.GetExploitType(), plugin.Desc()})
	}

	sort.Slice(sortedExploits, func(i, j int) bool {
		return sortedExploits[i].ExploitType < sortedExploits[j].ExploitType
	})

	fmt.Fprintln(writer, "TYPE \t NAME \t DESC")

	for _, kv := range sortedExploits {
		str := fmt.Sprintf("%s \t %s \t %s", kv.ExploitType, kv.Name, kv.Desc)
		fmt.Fprintln(writer, str)
	}

	writer.Flush()
}

func RunSingleExploit(name string) {
	Exploits[name].Run()
}

func RegisterExploit(name string, exploit ExploitInterface) {
	Exploits[name] = exploit
}

func RunSingleTask(name string) {
	// fmt.Printf("[+] Running exploit: %s.\n", name)
	// fmt.Printf("[+] %s\n", Tasks[name].Desc())
	// Can not call cli.Args here, because it will cause "import cycle".
	// fmt.Printf("[+] Args: %v.\n", cli.Args["<args>"])
	Tasks[name].Exec()
}

func RegisterTask(name string, task TaskInterface) {
	Tasks[name] = task
}
