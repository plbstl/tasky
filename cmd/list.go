/*
Copyright © 2021 Paul Ebose <paulebose@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"text/tabwriter"

	"github.com/paulebose/tasky/todo"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "print out available todos",
	Long:    `Shows a list of saved todos. By default, it prints out only uncompleted todos.`,
	Run:     listRun,
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func listRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(dataFile)
	cobra.CheckErr(err)

	sort.Sort(todo.ByPriority(items))

	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)
	for _, i := range items {
		fmt.Fprintln(
			w,
			i.Label()+"\t"+i.PrintPriority()+"\t"+i.Text+"\t",
		)
	}
	fmt.Fprintln(w, "- -")
	fmt.Fprintln(w, strconv.Itoa(len(items))+" total")
	w.Flush()
}
