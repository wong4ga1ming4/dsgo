# dsgo

调用 DeepSeek API 的命令行。一次一条命令，然后退出。默认纯文本；全局 `--json` 输出 JSON，`--config` 指定配置文件。

## 命令

- `status`：密钥是否可用、默认模型、模型列表和余额。
- `config`：把 `--apikey` 或 `--model` 写入配置文件。不访问网络。
- `chat`：提问一次。
- `search`：检索并返回正文。
- `fetch`：抓取一个网页的正文。
- `ocr`：识别文档或图片中的文字。

`chat` 可以加上这些开关，也可以同时用：

```text
dsgo chat [--model 模型] [--search] [--deepthink] [--file 路径 ...] "问题"
```

- `--search`：让模型在服务端联网搜索，模型固定为 `deepseek-flash`。正文后面附引用 URL。
- `--deepthink`：打开思考。和 `--search` 一起就是深度搜索。
- `--file`：先上传本地文件，再放进同一次对话。
- 都不带：普通对话。

## 配置

配置文件是 `~/.config/dsgo/config.yaml`：

```yaml
api_key: ""
base_url: https://api.deepseek.com
model: deepseek-flash
```

优先级：环境变量 `DEEPSEEK_API_KEY`、`DEEPSEEK_BASE_URL`、`DEEPSEEK_MODEL`，然后是命令上的显式参数（例如 `chat --model`），然后是配置文件，最后是内置默认值。密钥不要写进仓库或镜像。

## 安装

二进制从 GitHub Release `v0.1.0` 下载。换版本时设置 `DSGO_VERSION`。密钥用环境变量或配置文件，不要写进命令。

macOS 和 Linux：

```bash
curl -fsSL https://raw.githubusercontent.com/wong4ga1ming4/dsgo/main/install.sh | sh
```

Windows PowerShell：

```powershell
irm https://raw.githubusercontent.com/wong4ga1ming4/dsgo/main/install.ps1 | iex
```

## 构建

```bash
go build -o dsgo ./cmd/dsgo
```

## Docker

镜像里只有 `dsgo` 和系统根证书，以非 root 用户运行。密钥从运行时的环境变量传入。

```bash
docker build -t dsgo .
docker run --rm -e DEEPSEEK_API_KEY dsgo status
```
