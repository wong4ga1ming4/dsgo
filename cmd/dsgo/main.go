package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"dsgo/internal/exitx"
)

func main() {
	os.Exit(execute(os.Args[1:], os.Stdout, os.Stderr))
}

func execute(args []string, stdout, stderr io.Writer) int {
	root := newRoot(stdout, stderr)
	root.SetArgs(args)
	root.SetOut(stdout)
	root.SetErr(stderr)
	err := root.Execute()
	if err == nil {
		return 0
	}
	var coded *exitx.Error
	if errors.As(err, &coded) {
		if coded.Msg != "" {
			fmt.Fprintln(stderr, coded.Msg)
		}
		return coded.Code
	}
	msg := err.Error()
	if strings.HasPrefix(msg, "unknown command") {
		fmt.Fprintf(stderr, "未知命令\n\n%s", root.UsageString())
		return 2
	}
	fmt.Fprintln(stderr, msg)
	return 1
}

func newRoot(stdout, stderr io.Writer) *cobra.Command {
	root := &cobra.Command{
		Use:           "dsgo",
		Short:         "调用 DeepSeek API",
		SilenceUsage:  true,
		SilenceErrors: true,
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}
	root.SetOut(stdout)
	root.SetErr(stderr)
	root.PersistentFlags().String("config", "", "配置文件路径")
	root.PersistentFlags().Bool("json", false, "以 JSON 输出")
	root.SetFlagErrorFunc(func(cmd *cobra.Command, err error) error {
		return exitx.Usage(err.Error())
	})
	root.AddCommand(
		newStatus(),
		newConfig(),
		newChat(),
		newSearch(),
		newFetch(),
		newOCR(),
	)
	return root
}

func wantJSON(cmd *cobra.Command) bool {
	value, _ := cmd.Flags().GetBool("json")
	return value
}

func configFlag(cmd *cobra.Command) string {
	value, _ := cmd.Flags().GetString("config")
	return value
}
