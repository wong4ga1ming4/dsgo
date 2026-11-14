package main

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/spf13/cobra"

	"dsgo/internal/api"
	"dsgo/internal/client"
	"dsgo/internal/config"
	"dsgo/internal/exitx"
	"dsgo/internal/output"
)

func newStatus() *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "查看密钥、模型和余额",
		Args:  noArgs("用法: dsgo status"),
		RunE: func(cmd *cobra.Command, _ []string) error {
			settings, err := config.Load(configFlag(cmd))
			if err != nil {
				return exitx.Fail(err.Error())
			}
			if settings.APIKey == "" {
				return exitx.Fail("未设置 DEEPSEEK_API_KEY")
			}
			status, err := api.LoadStatus(cmd.Context(), client.New(settings.BaseURL, settings.APIKey))
			if err != nil {
				return exitx.Fail(err.Error())
			}
			var text strings.Builder
			if status.OK {
				text.WriteString("密钥可用\n")
			} else {
				text.WriteString("密钥不可用\n")
			}
			for _, model := range status.Models {
				if model.Name != "" && model.Name != model.ID {
					fmt.Fprintf(&text, "%s    %s\n", model.ID, model.Name)
					continue
				}
				fmt.Fprintln(&text, model.ID)
			}
			for _, balance := range status.Balances {
				fmt.Fprintf(&text, "%s  总额 %s  赠送 %s  充值 %s\n",
					balance.Currency, balance.Total, balance.Granted, balance.ToppedUp)
			}
			return writeOut(cmd, text.String(), status)
		},
	}
}

func newConfig() *cobra.Command {
	var apiKey, model string
	cmd := &cobra.Command{
		Use:   "config",
		Short: "写入配置文件",
		Args:  noArgs("用法: dsgo config --apikey <密钥> --model <模型 id>"),
		RunE: func(cmd *cobra.Command, _ []string) error {
			if !cmd.Flags().Changed("apikey") && !cmd.Flags().Changed("model") {
				return exitx.Usage("用法: dsgo config --apikey <密钥> --model <模型 id>")
			}
			var keyPtr, modelPtr *string
			if cmd.Flags().Changed("apikey") {
				keyPtr = &apiKey
			}
			if cmd.Flags().Changed("model") {
				modelPtr = &model
			}
			path, err := config.Update(configFlag(cmd), keyPtr, modelPtr)
			if err != nil {
				return exitx.Fail(err.Error())
			}
			text := "已写入 " + path + "\n"
			return writeOut(cmd, text, map[string]string{"path": path})
		},
	}
	cmd.Flags().StringVar(&apiKey, "apikey", "", "写入 API 密钥")
	cmd.Flags().StringVar(&model, "model", "", "写入默认模型")
	return cmd
}

func newChat() *cobra.Command {
	var model string
	var files []string
	cmd := &cobra.Command{
		Use:   "chat [问题]",
		Short: "提问一次",
		RunE: func(cmd *cobra.Command, args []string) error {
			prompt := strings.TrimSpace(strings.Join(args, " "))
			if prompt == "" {
				return exitx.Usage("用法: dsgo chat [--model 模型] [--file 路径] \"问题\"")
			}
			settings, err := config.Load(configFlag(cmd))
			if err != nil {
				return exitx.Fail(err.Error())
			}
			if settings.APIKey == "" {
				return exitx.Fail("未设置 DEEPSEEK_API_KEY")
			}
			chosen := settings.Model
			if cmd.Flags().Changed("model") {
				chosen = model
			}
			result, err := api.Chat(cmd.Context(), client.New(settings.BaseURL, settings.APIKey), chosen, prompt, files)
			if err != nil {
				return exitx.Fail(err.Error())
			}
			return writeOut(cmd, result.Content+"\n", result)
		},
	}
	cmd.Flags().StringVar(&model, "model", "", "模型 id")
	cmd.Flags().StringArrayVar(&files, "file", nil, "随问题上传的本地文件")
	return cmd
}

func newSearch() *cobra.Command {
	return &cobra.Command{
		Use:   "search [查询]",
		Short: "搜索并返回正文",
		RunE: func(cmd *cobra.Command, args []string) error {
			query := strings.TrimSpace(strings.Join(args, " "))
			if n := utf8.RuneCountInString(query); n < 1 || n > 128 {
				return exitx.Usage("用法: dsgo search \"查询\"")
			}
			_, apiClient, err := readyClient(cmd)
			if err != nil {
				return err
			}
			result, err := api.Search(cmd.Context(), apiClient, query)
			if err != nil {
				return exitx.Fail(err.Error())
			}
			return writeOut(cmd, pagesText(result.Pages), result)
		},
	}
}

func newFetch() *cobra.Command {
	return &cobra.Command{
		Use:   "fetch <url>",
		Short: "抓取网页正文",
		Args:  exact(1, "用法: dsgo fetch <url>"),
		RunE: func(cmd *cobra.Command, args []string) error {
			rawURL := strings.TrimSpace(args[0])
			if rawURL == "" {
				return exitx.Usage("用法: dsgo fetch <url>")
			}
			_, apiClient, err := readyClient(cmd)
			if err != nil {
				return err
			}
			page, err := api.Fetch(cmd.Context(), apiClient, rawURL)
			if err != nil {
				return exitx.Fail(err.Error())
			}
			text := page.Title + "\n" + strings.TrimRight(page.Content, "\n") + "\n"
			return writeOut(cmd, text, page)
		},
	}
}

func newOCR() *cobra.Command {
	return &cobra.Command{
		Use:   "ocr <file>",
		Short: "识别文档或图片中的文字",
		Args:  exact(1, "用法: dsgo ocr <file>"),
		RunE: func(cmd *cobra.Command, args []string) error {
			_, apiClient, err := readyClient(cmd)
			if err != nil {
				return err
			}
			result, err := api.OCR(cmd.Context(), apiClient, args[0])
			if err != nil {
				return exitx.Fail(err.Error())
			}
			parts := make([]string, 0, len(result.Pages))
			for _, page := range result.Pages {
				parts = append(parts, page.Text)
			}
			return writeOut(cmd, output.JoinPages(parts), result)
		},
	}
}

func readyClient(cmd *cobra.Command) (config.Settings, *client.Client, error) {
	settings, err := config.Load(configFlag(cmd))
	if err != nil {
		return config.Settings{}, nil, exitx.Fail(err.Error())
	}
	if settings.APIKey == "" {
		return config.Settings{}, nil, exitx.Fail("未设置 DEEPSEEK_API_KEY")
	}
	return settings, client.New(settings.BaseURL, settings.APIKey), nil
}

func writeOut(cmd *cobra.Command, text string, value any) error {
	if err := output.Write(cmd.OutOrStdout(), wantJSON(cmd), text, value); err != nil {
		return exitx.Fail(err.Error())
	}
	return nil
}

func pagesText(pages []api.Page) string {
	parts := make([]string, 0, len(pages))
	for _, page := range pages {
		parts = append(parts, page.Title+"\n"+strings.TrimRight(page.Content, "\n"))
	}
	if len(parts) == 0 {
		return "\n"
	}
	return strings.Join(parts, "\n\n") + "\n"
}

func noArgs(usage string) cobra.PositionalArgs {
	return func(_ *cobra.Command, args []string) error {
		if len(args) != 0 {
			return exitx.Usage(usage)
		}
		return nil
	}
}

func exact(n int, usage string) cobra.PositionalArgs {
	return func(_ *cobra.Command, args []string) error {
		if len(args) != n {
			return exitx.Usage(usage)
		}
		return nil
	}
}
