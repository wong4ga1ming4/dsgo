---
name: dsgo-cli
description: >-
  用 dsgo 命令行调用 DeepSeek：问答、联网搜索、抓取网页、识别文档或图片。
  用户要咨询问题、联网搜索问题、打开某个 URL、文件OCR时使用。
---

## 命令

```text
dsgo status
dsgo chat "问题"
dsgo chat --search "问题"
dsgo chat --search --deepthink "问题"
dsgo chat --file <file_path> "问题"
dsgo search "查询"
dsgo fetch https://example.com
dsgo ocr 文件路径
```

- `status`：密钥是否可用、默认模型、模型列表、余额。
- `chat`：一次性对话。选项 `--search` 开启联网搜索。`--deepthink` 打开思考。`--file` 上传文件。
- `search`：直接检索。
- `fetch`：只抓一个 URL 的标题和正文。
- `ocr`：文档（PDF、Word、PowerPoint）或图片（PNG、JPEG、WebP、GIF、BMP、TIFF）。

## 帮助
使用 `dsgo --help` 获取信息

每次使用都要先`dsgo status` 确认状态
