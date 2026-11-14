# Chat Completions API

根据输入的上下文，来让模型补全对话内容。

## 接口

```http
POST /chat/completions
```

operationId: `createChatCompletion`

### 服务地址

- `https://api.deepseek.com`

## 身份验证

### 验证方式 1

**bearerAuth**

`Authorization: Bearer <API_KEY>`

```yaml
type: http
scheme: bearer
```

## 请求体

**必填**

### application/json

- 引用: [CreateChatCompletionRequest](#schema-create-chat-completion-request) (`#/components/schemas/CreateChatCompletionRequest`)

## 响应

### 200 (No streaming)

OK, 返回一个 `chat completion` 对象。

#### application/json

- 引用: [CreateChatCompletionResponse](#schema-create-chat-completion-response) (`#/components/schemas/CreateChatCompletionResponse`)

**示例**

```json
{
  "id": "88beaa90-17fd-45fc-939c-2ef75e2d913c",
  "object": "chat.completion",
  "created": 1793084842,
  "model": "deepseek-flash",
  "choices": [
    {
      "index": 0,
      "message": {
        "role": "assistant",
        "content": "Hello! 👋 How can I help you today?"
      },
      "logprobs": null,
      "finish_reason": "stop"
    }
  ],
  "usage": {
    "prompt_tokens": 7,
    "completion_tokens": 11,
    "total_tokens": 18,
    "prompt_tokens_details": {
      "cached_tokens": 0
    },
    "prompt_cache_hit_tokens": 0,
    "prompt_cache_miss_tokens": 7
  },
  "system_fingerprint": "e72d3c7f663d3e1383b00b3884cb4501"
}
```

### 200 (Streaming)

OK, 返回包含一系列 `chat completion chunk` 对象的流式输出。

#### text/event-stream

- `type`: `array`

- **数组元素**

  - 引用: [CreateChatCompletionStreamResponse](#schema-create-chat-completion-stream-response) (`#/components/schemas/CreateChatCompletionStreamResponse`)

**示例**

```
data: {"id":"071f229d-daa0-4557-a580-106a22e8df73","object":"chat.completion.chunk","created":1793084907,"model":"deepseek-flash","system_fingerprint":"e72d3c7f663d3e1383b00b3884cb4501","choices":[{"index":0,"delta":{"role":"assistant","content":""},"logprobs":null,"finish_reason":null}],"usage":null}

data: {"id":"071f229d-daa0-4557-a580-106a22e8df73","object":"chat.completion.chunk","created":1793084907,"model":"deepseek-flash","system_fingerprint":"e72d3c7f663d3e1383b00b3884cb4501","choices":[{"index":0,"delta":{"content":"Hello"},"logprobs":null,"finish_reason":null}],"usage":null}

data: {"id":"071f229d-daa0-4557-a580-106a22e8df73","object":"chat.completion.chunk","created":1793084907,"model":"deepseek-flash","system_fingerprint":"e72d3c7f663d3e1383b00b3884cb4501","choices":[{"index":0,"delta":{"content":"!"},"logprobs":null,"finish_reason":null}],"usage":null}

data: {"id":"071f229d-daa0-4557-a580-106a22e8df73","object":"chat.completion.chunk","created":1793084907,"model":"deepseek-flash","system_fingerprint":"e72d3c7f663d3e1383b00b3884cb4501","choices":[{"index":0,"delta":{"content":" 👋"},"logprobs":null,"finish_reason":null}],"usage":null}

data: {"id":"071f229d-daa0-4557-a580-106a22e8df73","object":"chat.completion.chunk","created":1793084907,"model":"deepseek-flash","system_fingerprint":"e72d3c7f663d3e1383b00b3884cb4501","choices":[{"index":0,"delta":{"content":" How"},"logprobs":null,"finish_reason":null}],"usage":null}

data: {"id":"071f229d-daa0-4557-a580-106a22e8df73","object":"chat.completion.chunk","created":1793084907,"model":"deepseek-flash","system_fingerprint":"e72d3c7f663d3e1383b00b3884cb4501","choices":[{"index":0,"delta":{"content":" can"},"logprobs":null,"finish_reason":null}],"usage":null}

data: {"id":"071f229d-daa0-4557-a580-106a22e8df73","object":"chat.completion.chunk","created":1793084907,"model":"deepseek-flash","system_fingerprint":"e72d3c7f663d3e1383b00b3884cb4501","choices":[{"index":0,"delta":{"content":" I"},"logprobs":null,"finish_reason":null}],"usage":null}

data: {"id":"071f229d-daa0-4557-a580-106a22e8df73","object":"chat.completion.chunk","created":1793084907,"model":"deepseek-flash","system_fingerprint":"e72d3c7f663d3e1383b00b3884cb4501","choices":[{"index":0,"delta":{"content":" help"},"logprobs":null,"finish_reason":null}],"usage":null}

data: {"id":"071f229d-daa0-4557-a580-106a22e8df73","object":"chat.completion.chunk","created":1793084907,"model":"deepseek-flash","system_fingerprint":"e72d3c7f663d3e1383b00b3884cb4501","choices":[{"index":0,"delta":{"content":" you"},"logprobs":null,"finish_reason":null}],"usage":null}

data: {"id":"071f229d-daa0-4557-a580-106a22e8df73","object":"chat.completion.chunk","created":1793084907,"model":"deepseek-flash","system_fingerprint":"e72d3c7f663d3e1383b00b3884cb4501","choices":[{"index":0,"delta":{"content":" today"},"logprobs":null,"finish_reason":null}],"usage":null}

data: {"id":"071f229d-daa0-4557-a580-106a22e8df73","object":"chat.completion.chunk","created":1793084907,"model":"deepseek-flash","system_fingerprint":"e72d3c7f663d3e1383b00b3884cb4501","choices":[{"index":0,"delta":{"content":"?"},"logprobs":null,"finish_reason":null}],"usage":null}

data: {"id":"071f229d-daa0-4557-a580-106a22e8df73","object":"chat.completion.chunk","created":1793084907,"model":"deepseek-flash","system_fingerprint":"e72d3c7f663d3e1383b00b3884cb4501","choices":[{"index":0,"delta":{"content":""},"logprobs":null,"finish_reason":"stop"}],"usage":{"prompt_tokens":7,"completion_tokens":11,"total_tokens":18,"prompt_tokens_details":{"cached_tokens":0},"prompt_cache_hit_tokens":0,"prompt_cache_miss_tokens":7}}

data: [DONE]
```

## 代码示例

### curl

```bash
# Set DEEPSEEK_API_KEY to your API key before running this example.
curl --silent --show-error --max-time 60 "https://api.deepseek.com/chat/completions" \
  --header "Authorization: Bearer ${DEEPSEEK_API_KEY}" \
  --header "Content-Type: application/json" \
  --data '{
  "model": "deepseek-flash",
  "messages": [
    {
      "role": "user",
      "content": "Say hello."
    }
  ],
  "max_tokens": 32,
  "temperature": 0,
  "thinking": {
    "type": "disabled"
  },
  "stream": false
}'
```

### Python — OpenAI SDK

```python
# Set DEEPSEEK_API_KEY to your API key before running this example.
import os
from openai import OpenAI

client = OpenAI(
    api_key=os.environ["DEEPSEEK_API_KEY"],
    base_url="https://api.deepseek.com",
    timeout=60.0,
)

response = client.chat.completions.create(
    model="deepseek-flash",
    messages=[{"role": "user", "content": "Say hello."}],
    max_tokens=32,
    temperature=0,
    stream=False,
    extra_body={"thinking": {"type": "disabled"}},
)
print(response.model_dump_json(indent=2))
```

## 引用的数据结构

### schema-create-chat-completion-request

**CreateChatCompletionRequest** — `#/components/schemas/CreateChatCompletionRequest`

- `type`: `object`

- `required`: `["model","messages"]`

- **字段**

  - **`messages`** (必填)

    对话的消息列表。

    - `type`: `array`

    - `minItems`: `1`

    - **数组元素**

      - 引用: [ChatCompletionRequestMessage](#schema-chat-completion-request-message) (`#/components/schemas/ChatCompletionRequestMessage`)

    **示例**

    ```json
    [
      {
        "content": "You are a helpful assistant",
        "role": "system"
      },
      {
        "content": "Hi",
        "role": "user"
      }
    ]
    ```

  - **`model`** (必填)

    使用的模型的 ID。请使用 `deepseek-flash` 或 `deepseek-v4-pro`。

    - `type`: `string`

    - `enum`: `["deepseek-flash","deepseek-v4-pro"]`

    **示例**

    ```json
    "deepseek-flash"
    ```

  - **`thinking`** (可选)

    控制思考模式与非思考模式的转换。省略 `thinking` 或传 `null` 时，模式由 `reasoning_effort` 决定；如果 `reasoning_effort` 也未提供，则默认开启思考模式。提供此对象时，`type` 必填，其模式设置优先于 `reasoning_effort`。

    - `type`: `object`

    - `nullable`: `true`

    - `required`: `["type"]`

    - **字段**

      - **`type`** (必填)

        如果设为 `enabled`，则使用思考模式。如果设为 `disabled`，则使用非思考模式

        - `type`: `string`

        - `enum`: `["enabled","disabled"]`

  - **`reasoning_effort`** (可选)

    控制思考模式开关与思考强度。`none` 关闭思考模式；`low` / `high` / `max` 开启思考模式。默认强度为 `high`。出于兼容考虑，`minimal` 映射为 `low`，`medium` / `xhigh` 映射为 `high`。

    - `type`: `string`

    - `enum`: `["none","low","high","max"]`

  - **`max_tokens`** (可选)

    限制一次请求中模型生成 completion 的最大 token 数。取值范围为 1 到 384K（393216）。未设置时，非思考模式默认 8K，思考模式默认 64K（`reasoning_effort` 为 `max` 时为 128K）。输入 token 和输出 token 的总长度受模型的上下文长度的限制。详见[模型 & 价格](https://api-docs.deepseek.com/zh-cn/quick_start/pricing.md)。

    - `exclusiveMinimum`: `true`

    - `type`: `integer`

    - `nullable`: `true`

    **示例**

    ```json
    4096
    ```

  - **`response_format`** (可选)

    一个 object，指定模型必须输出的格式。省略或传 `null` 时使用普通文本；提供此对象时，`type` 必填。

    设置为 { "type": "json\_object" } 以启用 JSON 模式，该模式保证模型生成的消息是有效的 JSON。

    **注意:** 使用 JSON 模式时，你还必须通过系统或用户消息指示模型生成 JSON。如果模型的输出无法被识别为 JSON，API 会将其中的非 JSON 内容替换为空格（`U+0020`），因此返回的消息内容可能仅包含空格（看起来为空，此时 finish\_reason 仍可能是 stop）。此外，如果 finish\_reason="length"，这表示生成超过了 max\_tokens 或对话超过了最大上下文长度，消息内容可能会被部分截断。

    - `type`: `object`

    - `nullable`: `true`

    - `required`: `["type"]`

    - **字段**

      - **`type`** (必填)

        必须是 `text` 或 `json_object`。

        - `type`: `string`

        - `enum`: `["text","json_object"]`

        **示例**

        ```json
        "text"
        ```

  - **`stop`** (可选)

    一个 string 或最多包含 16 个 string 的 list，在遇到这些词时，API 将停止生成更多的 token。

    - `nullable`: `true`

    - `default`: `null`

    - **oneOf**

      - oneOf 1

        - title: `Single stop sequence`

        - `type`: `string`

        - `nullable`: `true`

        **示例**

        ```json
        null
        ```

      - oneOf 2

        - title: `Stop sequence list`

        - `type`: `array`

        - `minItems`: `1`

        - `maxItems`: `16`

        - **数组元素**

          - `type`: `string`

  - **`stream`** (可选)

    如果设置为 `true`，将会以 SSE（server-sent events）的形式以流式发送消息增量。消息流以 `data: [DONE]` 结尾。

    - `type`: `boolean`

    - `nullable`: `true`

    - `default`: `false`

  - **`stream_options`** (可选)

    - 引用: [ChatCompletionStreamOptions](#schema-chat-completion-stream-options) (`#/components/schemas/ChatCompletionStreamOptions`)

  - **`temperature`** (可选)

    采样温度，介于 0 和 2 之间。更高的值，如 0.8，会使输出更随机，而更低的值，如 0.2，会使其更加集中和确定。 我们通常建议可以更改这个值或者更改 `top_p`，但不建议同时对两者进行修改。思考模式下不生效。

    - `type`: `number`

    - `minimum`: `0`

    - `maximum`: `2`

    - `default`: `1`

    - `nullable`: `true`

    **示例**

    ```json
    1
    ```

  - **`top_p`** (可选)

    作为调节采样温度的替代方案，模型会考虑前 `top_p` 概率的 token 的结果。所以 0.1 就意味着只有包括在最高 10% 概率中的 token 会被考虑。取值必须大于 0 且不超过 1。我们通常建议修改这个值或者更改 `temperature`，但不建议同时对两者进行修改。实际生效值取决于模型和模式：

    - `deepseek-flash`：思考模式下默认值为 1.0，低于 0.95 的取值会按 0.95 处理；非思考模式下恒为 1.0，传入的值会被忽略。

    - `deepseek-v4-pro`：思考模式下，启用工具调用时恒为 0.95，否则恒为 1.0，传入的值会被忽略。即使提供了 `tools`，将 `tool_choice` 设为 `none` 也会关闭工具调用。非思考模式下使用传入的值，默认值为 1.0。

    - `type`: `number`

    - `minimum`: `0`

    - `exclusiveMinimum`: `true`

    - `maximum`: `1`

    - `nullable`: `true`

    **示例**

    ```json
    1
    ```

  - **`tools`** (可选)

    模型可能会调用的 tool 的列表：function，以及联网工具 `deepseek_web_search` 与 `deepseek_read_url`（见 [联网工具指南](https://api-docs.deepseek.com/zh-cn/guides/web_tools.md#chat-completions-api)）。使用 function 来提供以 JSON 作为输入参数的 function 列表。tool 名称必须唯一。

    - `type`: `array`

    - `nullable`: `true`

    - **数组元素**

      - **oneOf**

        - oneOf 1

          - 引用: [ChatCompletionTool](#schema-chat-completion-tool) (`#/components/schemas/ChatCompletionTool`)

        - oneOf 2

          - 引用: [ChatCompletionDeepSeekWebSearchTool](#schema-chat-completion-deep-seek-web-search-tool) (`#/components/schemas/ChatCompletionDeepSeekWebSearchTool`)

        - oneOf 3

          - 引用: [ChatCompletionDeepSeekReadUrlTool](#schema-chat-completion-deep-seek-read-url-tool) (`#/components/schemas/ChatCompletionDeepSeekReadUrlTool`)

    **示例**

    ```json
    null
    ```

  - **`tool_choice`** (可选)

    - 引用: [ChatCompletionToolChoiceOption](#schema-chat-completion-tool-choice-option) (`#/components/schemas/ChatCompletionToolChoiceOption`)

  - **`safety_check`** (可选)

    DeepSeek 扩展字段。`level` 设为 `high` 时开启基础风控，审核模型的输入与输出。详见 [模型内容安全指南](https://api-docs.deepseek.com/zh-cn/guides/safety.md)。

    - `type`: `object`

    - `nullable`: `true`

    - `required`: `["level"]`

    - **字段**

      - **`level`** (必填)

        `high`：开启基础风控。`low`：默认检查，与不传 `safety_check` 相同。

        - `type`: `string`

        - `enum`: `["low","high"]`

  - **`title`** (可选)

    DeepSeek 扩展字段。传入 `{"enabled": true}`，响应会带上简短的会话标题。详见 [会话标题指南](https://api-docs.deepseek.com/zh-cn/guides/conversation_title.md)。

    - `type`: `object`

    - `nullable`: `true`

    - `required`: `["enabled"]`

    - **字段**

      - **`enabled`** (必填)

        是否返回会话标题。

        - `type`: `boolean`

  - **`logprobs`** (可选)

    是否返回所输出 token 的对数概率。如果为 `true`，则在 `message` 的 `content` 中返回每个输出 token 的对数概率。

    - `type`: `boolean`

    - `default`: `false`

    - `nullable`: `true`

  - **`top_logprobs`** (可选)

    一个介于 0 到 20 之间的整数 N，指定每个输出位置返回输出概率 top N 的 token，且返回这些 token 的对数概率。指定此参数时，logprobs 必须为 `true`。

    - `type`: `integer`

    - `minimum`: `0`

    - `maximum`: `20`

    - `nullable`: `true`

    - `default`: `null`

    **示例**

    ```json
    null
    ```

  - **`user_id`** (可选)

    您自定义的 user\_id。建议仅使用英文字母、数字、下划线（`_`）和连字符（`-`）。UTF-8 编码后不得超过 512 字节。请不要在 user\_id 中包含用户隐私信息。

    - user\_id 可用于区分您业务侧的用户身份，以帮助我们进行内容安全处理。

    - user\_id 可用于 KVCache 缓存隔离，以进行隐私管理。

    - user\_id 可用于我们对您业务侧用户进行调度隔离。

    - 关于 user\_id 参数更详细的描述，请参考[限速与隔离](https://api-docs.deepseek.com/zh-cn/quick_start/rate_limit.md)

    * `nullable`: `true`

  - **`frequency_penalty`** (可选)

    该参数已不再支持。传入该参数将不会产生任何效果。

    - `deprecated`: `true`

  - **`presence_penalty`** (可选)

    该参数已不再支持。传入该参数将不会产生任何效果。

    - `deprecated`: `true`

### schema-create-chat-completion-response

**CreateChatCompletionResponse** — `#/components/schemas/CreateChatCompletionResponse`

Represents a chat completion response returned by model, based on the provided input.

- `type`: `object`

- `required`: `["choices","created","id","model","system_fingerprint","object"]`

- **字段**

  - **`id`** (必填)

    该对话的唯一标识符。

    - `type`: `string`

  - **`choices`** (必填)

    模型生成的 completion 的选择列表。

    - `type`: `array`

    - **数组元素**

      - `type`: `object`

      - `required`: `["finish_reason","index","message","logprobs"]`

      - **字段**

        - **`finish_reason`** (必填)

          模型停止生成 token 的原因。

          `stop`：模型自然停止生成，或遇到 `stop` 序列中列出的字符串。

          `length` ：输出长度达到了模型上下文长度限制，或达到了 `max_tokens` 的限制。

          `content_filter`：输出内容因触发过滤策略而被过滤。

          `tool_calls`：模型进行了工具调用。

          `insufficient_system_resource`：系统推理资源不足，生成被打断。

          `aborted`：生成过程被中断。

          - `type`: `string`

          - `enum`: `["stop","length","content_filter","tool_calls","insufficient_system_resource","aborted"]`

        - **`index`** (必填)

          该 completion 在模型生成的 completion 的选择列表中的索引。

          - `type`: `integer`

        - **`message`** (必填)

          - 引用: [ChatCompletionResponseMessage](#schema-chat-completion-response-message) (`#/components/schemas/ChatCompletionResponseMessage`)

        - **`logprobs`** (必填)

          该 choice 的对数概率信息。

          - `type`: `object`

          - `nullable`: `true`

          - `required`: `["content"]`

          - **字段**

            - **`content`** (必填)

              一个包含输出 token 对数概率信息的列表。

              - `type`: `array`

              - `nullable`: `true`

              - **数组元素**

                - 引用: [ChatCompletionTokenLogprob](#schema-chat-completion-token-logprob) (`#/components/schemas/ChatCompletionTokenLogprob`)

            - **`reasoning_content`** (可选)

              一个包含输出 token 对数概率信息的列表。

              - `type`: `array`

              - `nullable`: `true`

              - **数组元素**

                - 引用: [ChatCompletionTokenLogprob](#schema-chat-completion-token-logprob) (`#/components/schemas/ChatCompletionTokenLogprob`)

  - **`created`** (必填)

    创建聊天完成时的 Unix 时间戳（以秒为单位）。

    - `type`: `integer`

  - **`model`** (必填)

    生成该 completion 的模型名。

    - `type`: `string`

  - **`system_fingerprint`** (必填)

    This fingerprint represents the backend configuration that the model runs with.

    - `type`: `string`

  - **`object`** (必填)

    对象的类型, 其值为 `chat.completion`。

    - `type`: `string`

    - `enum`: `["chat.completion"]`

  - **`usage`** (可选)

    - 引用: [CompletionUsage](#schema-completion-usage) (`#/components/schemas/CompletionUsage`)

### schema-create-chat-completion-stream-response

**CreateChatCompletionStreamResponse** — `#/components/schemas/CreateChatCompletionStreamResponse`

基于提供的输入，由模型返回的对话补全响应的流式输出。

- `type`: `object`

- `required`: `["choices","created","id","model","system_fingerprint","object"]`

- **字段**

  - **`id`** (必填)

    该对话的唯一标识符。

    - `type`: `string`

  - **`choices`** (必填)

    模型生成的 completion 的选择列表。

    - `type`: `array`

    - **数组元素**

      - `type`: `object`

      - `required`: `["delta","finish_reason","index"]`

      - **字段**

        - **`delta`** (必填)

          - 引用: [ChatCompletionStreamResponseDelta](#schema-chat-completion-stream-response-delta) (`#/components/schemas/ChatCompletionStreamResponseDelta`)

        - **`server_tool_event`** (可选)

          DeepSeek 扩展字段，使用联网工具时返回。服务端工具调用的进度，每个 chunk 最多携带一个事件。详见 [联网工具指南](https://api-docs.deepseek.com/zh-cn/guides/web_tools.md#chat-completions-api)。

          - `type`: `object`

        - **`logprobs`** (可选)

          该 choice 的对数概率信息。

          - `type`: `object`

          - `nullable`: `true`

          - `required`: `["content"]`

          - **字段**

            - **`content`** (必填)

              一个包含输出 token 对数概率信息的列表。

              - `type`: `array`

              - `nullable`: `true`

              - **数组元素**

                - 引用: [ChatCompletionTokenLogprob](#schema-chat-completion-token-logprob) (`#/components/schemas/ChatCompletionTokenLogprob`)

            - **`reasoning_content`** (可选)

              一个包含输出 token 对数概率信息的列表。

              - `type`: `array`

              - `nullable`: `true`

              - **数组元素**

                - 引用: [ChatCompletionTokenLogprob](#schema-chat-completion-token-logprob) (`#/components/schemas/ChatCompletionTokenLogprob`)

        - **`finish_reason`** (必填)

          模型停止生成 token 的原因。

          `stop`：模型自然停止生成，或遇到 `stop` 序列中列出的字符串。

          `length` ：输出长度达到了模型上下文长度限制，或达到了 `max_tokens` 的限制。

          `content_filter`：输出内容因触发过滤策略而被过滤。

          `tool_calls`：模型进行了工具调用。

          `insufficient_system_resource`: 由于后端推理资源受限，请求被打断。

          `aborted`：生成过程被中断。

          - `type`: `string`

          - `enum`: `["stop","length","content_filter","tool_calls","insufficient_system_resource","aborted"]`

          - `nullable`: `true`

        - **`index`** (必填)

          该 completion 在模型生成的 completion 的选择列表中的索引。

          - `type`: `integer`

  - **`created`** (必填)

    创建聊天完成时的 Unix 时间戳（以秒为单位）。流式响应的每个 chunk 的时间戳相同。

    - `type`: `integer`

  - **`model`** (必填)

    生成该 completion 的模型名。

    - `type`: `string`

  - **`system_fingerprint`** (必填)

    This fingerprint represents the backend configuration that the model runs with.

    - `type`: `string`

  - **`object`** (必填)

    对象的类型, 其值为 `chat.completion.chunk`。

    - `type`: `string`

    - `enum`: `["chat.completion.chunk"]`

### schema-chat-completion-request-message

**ChatCompletionRequestMessage** — `#/components/schemas/ChatCompletionRequestMessage`

- **oneOf**

  - oneOf 1

    - 引用: [ChatCompletionRequestSystemMessage](#schema-chat-completion-request-system-message) (`#/components/schemas/ChatCompletionRequestSystemMessage`)

  - oneOf 2

    - 引用: [ChatCompletionRequestUserMessage](#schema-chat-completion-request-user-message) (`#/components/schemas/ChatCompletionRequestUserMessage`)

  - oneOf 3

    - 引用: [ChatCompletionRequestAssistantMessage](#schema-chat-completion-request-assistant-message) (`#/components/schemas/ChatCompletionRequestAssistantMessage`)

  - oneOf 4

    - 引用: [ChatCompletionRequestToolMessage](#schema-chat-completion-request-tool-message) (`#/components/schemas/ChatCompletionRequestToolMessage`)

  - oneOf 5

    - 引用: [ChatCompletionRequestLatestReminderMessage](#schema-chat-completion-request-latest-reminder-message) (`#/components/schemas/ChatCompletionRequestLatestReminderMessage`)

### schema-chat-completion-stream-options

**ChatCompletionStreamOptions** — `#/components/schemas/ChatCompletionStreamOptions`

流式输出相关选项。必须与 `stream: true` 一起使用；如果 `stream` 未设置为 `true`，API 会返回 `400` 错误。

- `type`: `object`

- `nullable`: `true`

- `default`: `null`

- **字段**

  - **`include_usage`** (可选)

    如果设置为 `true`，流式返回的所有块都会包含 `usage` 字段，其中除最后一个块外，该字段的值均为 `null`。如果不设置或设置为 `false`，则除最后一个块外，其余块都不含 `usage` 字段。

    无论是否设置，`data: [DONE]` 之前的最后一个块都会在其 `usage` 字段中给出整个请求的 token 使用统计信息。请注意，这里不会单独下发一个只含 usage 的块：统计信息附加在最后一个内容块上，该块的 `choices` 数组始终只包含一个元素，其中不含新增内容且 `finish_reason` 非 null。

    - `type`: `boolean`

**示例**

```json
null
```

### schema-chat-completion-tool

**ChatCompletionTool** — `#/components/schemas/ChatCompletionTool`

- title: `Function 工具`

- `type`: `object`

- `required`: `["type","function"]`

- **字段**

  - **`type`** (必填)

    tool 的类型，此场景下为 `function`。

    - `type`: `string`

    - `enum`: `["function"]`

  - **`function`** (必填)

    - 引用: [FunctionObject](#schema-function-object) (`#/components/schemas/FunctionObject`)

### schema-chat-completion-deep-seek-web-search-tool

**ChatCompletionDeepSeekWebSearchTool** — `#/components/schemas/ChatCompletionDeepSeekWebSearchTool`

- title: `DeepSeek 联网搜索工具`

DeepSeek 扩展。让模型联网搜索；开启思考时进行深度思考搜索。详见 [联网工具指南](https://api-docs.deepseek.com/zh-cn/guides/web_tools.md#chat-completions-api)。

- `type`: `object`

- `required`: `["type"]`

- **字段**

  - **`type`** (必填)

    tool 的类型，此场景下为 `deepseek_web_search`。

    - `type`: `string`

    - `enum`: `["deepseek_web_search"]`

### schema-chat-completion-deep-seek-read-url-tool

**ChatCompletionDeepSeekReadUrlTool** — `#/components/schemas/ChatCompletionDeepSeekReadUrlTool`

- title: `DeepSeek 阅读链接工具`

DeepSeek 扩展。让模型阅读最后一条用户消息中的链接。详见 [联网工具指南](https://api-docs.deepseek.com/zh-cn/guides/web_tools.md#read-links)。

- `type`: `object`

- `required`: `["type"]`

- **字段**

  - **`type`** (必填)

    tool 的类型，此场景下为 `deepseek_read_url`。

    - `type`: `string`

    - `enum`: `["deepseek_read_url"]`

### schema-chat-completion-tool-choice-option

**ChatCompletionToolChoiceOption** — `#/components/schemas/ChatCompletionToolChoiceOption`

控制模型调用 tool 的行为。

`none` 意味着模型不会调用任何 tool，而是生成一条消息。

`auto` 意味着模型可以选择生成一条消息或调用一个或多个 tool。

`required` 意味着模型必须调用一个或多个 tool。

通过 `{"type": "function", "function": {"name": "my_function"}}` 指定特定 tool，会强制模型调用该 tool。

当没有 tool 时，默认值为 `none`。如果有 tool 存在，默认值为 `auto`。

思考模式下不支持 `required` 和指定具体 tool 的用法，API 会返回 `400` 错误。请先关闭思考模式。声明了联网工具时，`required` 按 `auto` 处理。

- `type`: `["string","object"]`

- `nullable`: `true`

- **oneOf**

  - oneOf 1

    - 引用: [ChatCompletionToolChoice](#schema-chat-completion-tool-choice) (`#/components/schemas/ChatCompletionToolChoice`)

  - oneOf 2

    - 引用: [ChatCompletionNamedToolChoice](#schema-chat-completion-named-tool-choice) (`#/components/schemas/ChatCompletionNamedToolChoice`)

### schema-chat-completion-response-message

**ChatCompletionResponseMessage** — `#/components/schemas/ChatCompletionResponseMessage`

模型生成的 completion 消息。

- `type`: `object`

- `required`: `["role","content"]`

- **字段**

  - **`content`** (必填)

    该 completion 的内容。

    - `type`: `string`

    - `nullable`: `true`

  - **`reasoning_content`** (可选)

    仅适用于思考模式。内容为 assistant 消息中在最终答案之前的推理内容。

    - `type`: `string`

    - `nullable`: `true`

  - **`tool_calls`** (可选)

    - 引用: [ChatCompletionMessageToolCalls](#schema-chat-completion-message-tool-calls) (`#/components/schemas/ChatCompletionMessageToolCalls`)

  - **`role`** (必填)

    生成这条消息的角色。

    - `type`: `string`

    - `enum`: `["assistant"]`

  - **`annotations`** (可选)

    使用联网工具时返回。`content` 中的引用，见 [联网工具 - 引用](https://api-docs.deepseek.com/zh-cn/guides/web_tools.md#citations)。

    - `type`: `array`

    - **数组元素**

      - `type`: `object`

  - **`server_tool_calls`** (可选)

    DeepSeek 扩展字段，使用联网工具时返回。本次响应的服务端工具调用，每个元素是一步；下一轮请求中原样回传。见 [联网工具指南](https://api-docs.deepseek.com/zh-cn/guides/web_tools.md#chat-completions-api)。

    - `type`: `array`

    - **数组元素**

      - `type`: `object`

  - **`title`** (可选)

    DeepSeek 扩展字段。会话标题，`title.enabled` 为 `true` 时返回。见 [会话标题指南](https://api-docs.deepseek.com/zh-cn/guides/conversation_title.md)。

    - `type`: `string`

  - **`refusal`** (可选)

    请求中的图片未通过模型内容安全审核时，指出被拦截的图片。见 [模型内容安全指南](https://api-docs.deepseek.com/zh-cn/guides/safety.md#handling-refusals)。

    - `type`: `string`

### schema-chat-completion-token-logprob

**ChatCompletionTokenLogprob** — `#/components/schemas/ChatCompletionTokenLogprob`

- `type`: `object`

- `required`: `["token","logprob","bytes","top_logprobs"]`

- **字段**

  - **`token`** (必填)

    输出的 token。

    - `type`: `string`

  - **`logprob`** (必填)

    该 token 的对数概率。`-9999.0` 代表该 token 的输出概率极小，不在 top 20 最可能输出的 token 中。

    - `type`: `number`

  - **`bytes`** (必填)

    一个包含该 token UTF-8 字节表示的整数列表。一般在一个 UTF-8 字符被拆分成多个 token 来表示时有用。如果 token 没有对应的字节表示，则该值为 `null`。

    - `type`: `array`

    - `nullable`: `true`

    - **数组元素**

      - `type`: `integer`

  - **`top_logprobs`** (必填)

    一个包含在该输出位置上，输出概率 top N 的 token 的列表，以及它们的对数概率。在罕见情况下，返回的 token 数量可能少于请求参数中指定的 `top_logprobs` 值。

    - `type`: `array`

    - **数组元素**

      - `type`: `object`

      - `required`: `["token","logprob","bytes"]`

      - **字段**

        - **`token`** (必填)

          输出的 token。

          - `type`: `string`

        - **`logprob`** (必填)

          该 token 的对数概率。`-9999.0` 代表该 token 的输出概率极小，不在 top 20 最可能输出的 token 中。

          - `type`: `number`

        - **`bytes`** (必填)

          一个包含该 token UTF-8 字节表示的整数列表。一般在一个 UTF-8 字符被拆分成多个 token 来表示时有用。如果 token 没有对应的字节表示，则该值为 `null`。

          - `type`: `array`

          - `nullable`: `true`

          - **数组元素**

            - `type`: `integer`

### schema-completion-usage

**CompletionUsage** — `#/components/schemas/CompletionUsage`

该对话补全请求的用量信息。

- `type`: `object`

- `required`: `["prompt_tokens","prompt_tokens_details","prompt_cache_hit_tokens","prompt_cache_miss_tokens","completion_tokens","total_tokens"]`

- **字段**

  - **`completion_tokens`** (必填)

    模型 completion 产生的 token 数。

    - `type`: `integer`

  - **`prompt_tokens`** (必填)

    用户 prompt 所包含的 token 数。该值等于 `prompt_cache_hit_tokens + prompt_cache_miss_tokens`

    - `type`: `integer`

  - **`prompt_tokens_details`** (必填)

    prompt tokens 的详细信息。

    - `type`: `object`

    - **字段**

      - **`cached_tokens`** (可选)

        用户 prompt 中，命中上下文缓存的 token 数。与 `prompt_cache_hit_tokens` 相同。

        - `type`: `integer`

  - **`prompt_cache_hit_tokens`** (必填)

    用户 prompt 中，命中上下文缓存的 token 数。

    - `type`: `integer`

  - **`prompt_cache_miss_tokens`** (必填)

    用户 prompt 中，未命中上下文缓存的 token 数。

    - `type`: `integer`

  - **`total_tokens`** (必填)

    该请求中，所有 token 的数量（prompt + completion）。

    - `type`: `integer`

  - **`completion_tokens_details`** (可选)

    completion tokens 的详细信息。

    - `type`: `object`

    - **字段**

      - **`reasoning_tokens`** (可选)

        推理模型所产生的思维链 token 数量

        - `type`: `integer`

      - **`safety_tokens`** (可选)

        DeepSeek 扩展字段。安全审核产生的 token 数，已包含在 `completion_tokens` 中。仅在 `safety_check` 为 `high` 时返回。

        - `type`: `integer`

        - `minimum`: `0`

  - **`server_tool_use`** (可选)

    使用联网工具时返回。计费的搜索与抓取次数；各项计数均为零时省略。详见 [联网工具指南](https://api-docs.deepseek.com/zh-cn/guides/web_tools.md#chat-completions-api)。

    - `type`: `object`

    - **字段**

      - **`tool_search`** (可选)

        计费的搜索次数，不含图片搜索。

        - `type`: `integer`

        - `minimum`: `0`

      - **`tool_fetch`** (可选)

        计费的网页抓取次数。

        - `type`: `integer`

        - `minimum`: `0`

      - **`tool_image_search`** (可选)

        计费的图片搜索次数。

        - `type`: `integer`

        - `minimum`: `0`

### schema-chat-completion-stream-response-delta

**ChatCompletionStreamResponseDelta** — `#/components/schemas/ChatCompletionStreamResponseDelta`

流式返回的一个 completion 增量。

- `type`: `object`

- **字段**

  - **`content`** (可选)

    completion 增量的内容。

    - `type`: `string`

    - `nullable`: `true`

  - **`reasoning_content`** (可选)

    仅适用于思考模式。内容为 assistant 消息中在最终答案之前的推理内容。

    - `type`: `string`

    - `nullable`: `true`

  - **`role`** (可选)

    产生这条消息的角色。

    - `type`: `string`

    - `enum`: `["assistant"]`

  - **`tool_calls`** (可选)

    模型生成的 tool 调用，例如 function 调用。每个 tool 调用的第一个 chunk 携带 `id`、`type` 和 `function` 字段，后续 chunk 只携带 function 参数。

    - `type`: `array`

    - **数组元素**

      - 引用: [ChatCompletionMessageToolCallChunk](#schema-chat-completion-message-tool-call-chunk) (`#/components/schemas/ChatCompletionMessageToolCallChunk`)

  - **`annotations`** (可选)

    使用联网工具时返回。`content` 中的引用，见 [联网工具 - 引用](https://api-docs.deepseek.com/zh-cn/guides/web_tools.md#citations)。流式输出时只在最后一个 chunk 中出现一次。

    - `type`: `array`

    - **数组元素**

      - `type`: `object`

  - **`server_tool_calls`** (可选)

    DeepSeek 扩展字段，使用联网工具时返回。本次响应的服务端工具调用，每个元素是一步；下一轮请求中原样回传。见 [联网工具指南](https://api-docs.deepseek.com/zh-cn/guides/web_tools.md#chat-completions-api)。流式输出时只在最后一个 chunk 中出现一次。

    - `type`: `array`

    - **数组元素**

      - `type`: `object`

  - **`title`** (可选)

    DeepSeek 扩展字段。会话标题，`title.enabled` 为 `true` 时返回。见 [会话标题指南](https://api-docs.deepseek.com/zh-cn/guides/conversation_title.md)。流式输出时只在最后一个 chunk 中出现一次。

    - `type`: `string`

  - **`refusal`** (可选)

    请求中的图片未通过模型内容安全审核时，指出被拦截的图片。见 [模型内容安全指南](https://api-docs.deepseek.com/zh-cn/guides/safety.md#handling-refusals)。流式输出时只在最后一个 chunk 中出现一次。

    - `type`: `string`

### schema-chat-completion-request-system-message

**ChatCompletionRequestSystemMessage** — `#/components/schemas/ChatCompletionRequestSystemMessage`

- title: `System message`

- `type`: `object`

- `required`: `["content","role"]`

- **字段**

  - **`content`** (必填)

    system 消息的内容。

    - `type`: `string`

  - **`role`** (必填)

    该消息的发起角色，其值为 `system`。

    - `type`: `string`

    - `enum`: `["system"]`

  - **`name`** (可选)

    可以选填的参与者的名称，为模型提供信息以区分相同角色的参与者。

    - `type`: `string`

### schema-chat-completion-request-user-message

**ChatCompletionRequestUserMessage** — `#/components/schemas/ChatCompletionRequestUserMessage`

- title: `User message`

- `type`: `object`

- `required`: `["content","role"]`

- **字段**

  - **`content`** (必填)

    user 消息的内容。可以是字符串，也可以是内容块数组（可携带图片）。详见[图像理解指南](https://api-docs.deepseek.com/zh-cn/guides/vision.md)。

    - **oneOf**

      - oneOf 1

        - title: `Text content`

        - `type`: `string`

      - oneOf 2

        - title: `Array of content parts`

        - `type`: `array`

        - `minItems`: `1`

        - **数组元素**

          - 引用: [ChatCompletionRequestContentPart](#schema-chat-completion-request-content-part) (`#/components/schemas/ChatCompletionRequestContentPart`)

  - **`role`** (必填)

    该消息的发起角色，其值为 `user`。

    - `type`: `string`

    - `enum`: `["user"]`

  - **`name`** (可选)

    可以选填的参与者的名称，为模型提供信息以区分相同角色的参与者。

    - `type`: `string`

### schema-chat-completion-request-assistant-message

**ChatCompletionRequestAssistantMessage** — `#/components/schemas/ChatCompletionRequestAssistantMessage`

- title: `Assistant message`

- `type`: `object`

- `required`: `["role","content"]`

- **字段**

  - **`content`** (必填)

    assistant 消息的内容。

    - `nullable`: `true`

    - `type`: `string`

  - **`role`** (必填)

    该消息的发起角色，其值为 `assistant`。

    - `type`: `string`

    - `enum`: `["assistant"]`

  - **`name`** (可选)

    可以选填的参与者的名称，为模型提供信息以区分相同角色的参与者。

    - `type`: `string`

  - **`prefix`** (可选)

    (Beta) 设置此参数为 `true`，来强制模型在其回答中以此 `assistant` 消息中提供的前缀内容开始。

    您必须设置 `base_url="https://api.deepseek.com/beta"` 来使用此功能。

    - `type`: `bool`

  - **`reasoning_content`** (可选)

    (Beta) 用于思考模式下在[对话前缀续写](https://api-docs.deepseek.com/zh-cn/guides/chat_prefix_completion.md)功能下，作为最后一条 assistant 思维链内容的输入。使用此功能时，`prefix` 参数必须设置为 `true`。

    - `nullable`: `true`

    - `type`: `string`

  - **`server_tool_calls`** (可选)

    DeepSeek 扩展字段。此前响应中的 `server_tool_calls`，原样回传。见 [联网工具指南](https://api-docs.deepseek.com/zh-cn/guides/web_tools.md#chat-completions-api)。

    - `type`: `array`

    - **数组元素**

      - `type`: `object`

### schema-chat-completion-request-tool-message

**ChatCompletionRequestToolMessage** — `#/components/schemas/ChatCompletionRequestToolMessage`

- title: `Tool message`

- `type`: `object`

- `required`: `["role","content","tool_call_id"]`

- **字段**

  - **`role`** (必填)

    该消息的发起角色，其值为 `tool`。

    - `type`: `string`

    - `enum`: `["tool"]`

  - **`content`** (必填)

    - title: `Text content`

    tool 消息的内容。可以是字符串，也可以是内容块数组（可以携带图片）。详见 [Vision 指南](https://api-docs.deepseek.com/zh-cn/guides/vision.md)。

    - **oneOf**

      - oneOf 1

        - title: `Text content`

        - `type`: `string`

      - oneOf 2

        - title: `Array of content parts`

        - `type`: `array`

        - `minItems`: `1`

        - **数组元素**

          - 引用: [ChatCompletionRequestContentPart](#schema-chat-completion-request-content-part) (`#/components/schemas/ChatCompletionRequestContentPart`)

  - **`tool_call_id`** (必填)

    此消息所响应的 tool call 的 ID。

    - `type`: `string`

### schema-chat-completion-request-latest-reminder-message

**ChatCompletionRequestLatestReminderMessage** — `#/components/schemas/ChatCompletionRequestLatestReminderMessage`

- title: `Latest reminder message`

DeepSeek 扩展。与 DeepSeek 网页端一样，向模型提供当前日期、地点、平台与语言。请紧挨着放在它所属的 `user` 消息之前，放在其他位置返回 `400`。详见 [用户元信息标记指南](https://api-docs.deepseek.com/zh-cn/guides/latest_reminder.md)。

- `type`: `object`

- `required`: `["role","content"]`

- **字段**

  - **`role`** (必填)

    该消息的发起角色，其值为 `latest_reminder`。

    - `type`: `string`

    - `enum`: `["latest_reminder"]`

  - **`content`** (必填)

    `<date>,<weekday>,<location>,<platform>,<language>`，例如 `2026-11-11,星期三,新加坡,Web,中文`。

    - `type`: `string`

### schema-function-object

**FunctionObject** — `#/components/schemas/FunctionObject`

- `type`: `object`

- `required`: `["name"]`

- **字段**

  - **`description`** (可选)

    function 的功能描述，供模型理解何时以及如何调用该 function。

    - `type`: `string`

  - **`name`** (必填)

    要调用的 function 名称。必须由 a-z、A-Z、0-9 字符组成，或包含下划线和连字符，最大长度为 128 个字符。

    - `type`: `string`

  - **`parameters`** (可选)

    - 引用: [FunctionParameters](#schema-function-parameters) (`#/components/schemas/FunctionParameters`)

  - **`strict`** (可选)

    如果设置为 `true`，API 将在函数调用中使用 strict 模式，以确保输出始终符合函数的 JSON schema 定义。该功能为 Beta 功能，详细使用方式请参阅[Tool Calls 指南](https://api-docs.deepseek.com/zh-cn/guides/tool_calls.md)

    - `type`: `boolean`

    - `default`: `false`

### schema-chat-completion-tool-choice

**ChatCompletionToolChoice** — `#/components/schemas/ChatCompletionToolChoice`

- `type`: `string`

- `enum`: `["none","auto","required"]`

### schema-chat-completion-named-tool-choice

**ChatCompletionNamedToolChoice** — `#/components/schemas/ChatCompletionNamedToolChoice`

指定模型应使用的 tool。用于强制模型调用特定 function。

- `type`: `object`

- `required`: `["type","function"]`

- **字段**

  - **`type`** (必填)

    tool 的类型。目前，仅支持 `function`。

    - `type`: `string`

    - `enum`: `["function"]`

  - **`function`** (必填)

    - `type`: `object`

    - `required`: `["name"]`

    - **字段**

      - **`name`** (必填)

        要调用的函数名称。

        - `type`: `string`

### schema-chat-completion-message-tool-calls

**ChatCompletionMessageToolCalls** — `#/components/schemas/ChatCompletionMessageToolCalls`

模型生成的 tool 调用，例如 function 调用。

- `type`: `array`

- **数组元素**

  - 引用: [ChatCompletionMessageToolCall](#schema-chat-completion-message-tool-call) (`#/components/schemas/ChatCompletionMessageToolCall`)

### schema-chat-completion-message-tool-call-chunk

**ChatCompletionMessageToolCallChunk** — `#/components/schemas/ChatCompletionMessageToolCallChunk`

- `type`: `object`

- `required`: `["index"]`

- **字段**

  - **`index`** (必填)

    - `type`: `integer`

  - **`id`** (可选)

    tool 的 ID。

    - `type`: `string`

  - **`type`** (可选)

    tool 的类型。目前仅支持 `function`。

    - `type`: `string`

    - `enum`: `["function"]`

  - **`function`** (可选)

    - `type`: `object`

    - **字段**

      - **`name`** (可选)

        要调用的 function 名。

        - `type`: `string`

      - **`arguments`** (可选)

        要调用的 function 的参数，由模型生成，格式为 JSON。请注意，模型并不总是生成有效的 JSON，并且可能会臆造出你的函数模式中未定义的参数。在调用函数之前，请在代码中验证这些参数。

        - `type`: `string`

### schema-chat-completion-request-content-part

**ChatCompletionRequestContentPart** — `#/components/schemas/ChatCompletionRequestContentPart`

- **oneOf**

  - oneOf 1

    - 引用: [ChatCompletionRequestContentPartText](#schema-chat-completion-request-content-part-text) (`#/components/schemas/ChatCompletionRequestContentPartText`)

  - oneOf 2

    - 引用: [ChatCompletionRequestContentPartImage](#schema-chat-completion-request-content-part-image) (`#/components/schemas/ChatCompletionRequestContentPartImage`)

  - oneOf 3

    - 引用: [ChatCompletionRequestContentPartFile](#schema-chat-completion-request-content-part-file) (`#/components/schemas/ChatCompletionRequestContentPartFile`)

### schema-function-parameters

**FunctionParameters** — `#/components/schemas/FunctionParameters`

function 的输入参数，以 JSON Schema 对象描述。请参阅[Tool Calls 指南](https://api-docs.deepseek.com/zh-cn/guides/tool_calls.md)获取示例，并参阅[JSON Schema 参考](https://json-schema.org/understanding-json-schema/)了解有关格式的文档。省略 `parameters` 会定义一个参数列表为空的 function。

- `type`: `object`

- **additionalProperties**

  - 数据结构: `true`

### schema-chat-completion-message-tool-call

**ChatCompletionMessageToolCall** — `#/components/schemas/ChatCompletionMessageToolCall`

- `type`: `object`

- `required`: `["id","type","function"]`

- **字段**

  - **`id`** (必填)

    tool 调用的 ID。

    - `type`: `string`

  - **`type`** (必填)

    tool 的类型。目前仅支持 `function`。

    - `type`: `string`

    - `enum`: `["function"]`

  - **`function`** (必填)

    模型调用的 function。

    - `type`: `object`

    - `required`: `["name","arguments"]`

    - **字段**

      - **`name`** (必填)

        模型调用的 function 名。

        - `type`: `string`

      - **`arguments`** (必填)

        要调用的 function 的参数，由模型生成，格式为 JSON。请注意，模型并不总是生成有效的 JSON，并且可能会臆造出你函数模式中未定义的参数。在调用函数之前，请在代码中验证这些参数。

        - `type`: `string`

### schema-chat-completion-request-content-part-text

**ChatCompletionRequestContentPartText** — `#/components/schemas/ChatCompletionRequestContentPartText`

- title: `Text content part`

- `type`: `object`

- `required`: `["type","text"]`

- **字段**

  - **`type`** (必填)

    内容块的类型，此场景下为 `text`。

    - `type`: `string`

    - `enum`: `["text"]`

  - **`text`** (必填)

    文本内容。

    - `type`: `string`

### schema-chat-completion-request-content-part-image

**ChatCompletionRequestContentPartImage** — `#/components/schemas/ChatCompletionRequestContentPartImage`

- title: `Image content part`

- `type`: `object`

- `required`: `["type","image_url"]`

- **字段**

  - **`type`** (必填)

    内容块的类型，此场景下为 `image_url`。

    - `type`: `string`

    - `enum`: `["image_url"]`

  - **`image_url`** (必填)

    - `type`: `object`

    - `required`: `["url"]`

    - **字段**

      - **`url`** (必填)

        图片的 `http(s)` URL（最多 8192 个字符）或 base64 编码的 data URL（`data:image/jpeg;base64,...`）。支持的格式：JPEG、PNG、GIF、WebP。

        - `type`: `string`

      - **`detail`** (可选)

        控制图片的处理方式。`low` 将图片缩小到 512x512（更快、更省 token）；`high`、`original` 与 `auto` 保留原图。

        - `type`: `string`

        - `enum`: `["low","high","original","auto"]`

### schema-chat-completion-request-content-part-file

**ChatCompletionRequestContentPartFile** — `#/components/schemas/ChatCompletionRequestContentPartFile`

- title: `File content part`

`file_id` 与 `file_data` 必须二选一：两者都不传返回 `400` 错误（"file must have a file\_id or file\_data"）；两者都传返回 `400` 错误（"file cannot have both file\_id and file\_data"）。

- `type`: `object`

- `required`: `["type"]`

- **字段**

  - **`type`** (必填)

    内容块的类型，此场景下为 `file`。

    - `type`: `string`

    - `enum`: `["file"]`

  - **`file_id`** (可选)

    通过 [Files API](https://api-docs.deepseek.com/zh-cn/guides/files_api.md) 上传的文件 ID，形如 `file-api-...`。与 `file_data` 互斥。

    - `type`: `string`

  - **`file_data`** (可选)

    图片的 base64 编码 data URL（`data:image/jpeg;base64,...`）。与 `file_id` 互斥。

    - `type`: `string`

  - **`filename`** (可选)

    可选的文件名，仅在配合 `file_data` 时有效。

    - `type`: `string`



# Responses API

以 OpenAI Responses API 格式创建模型响应。

该 API 是**无状态**的：服务端不存储响应与会话。多轮对话需要客户端在每次请求的 `input` 中回传完整对话历史。详细说明（含完整的参数兼容性表）请参考 [Responses API 指南](https://api-docs.deepseek.com/zh-cn/guides/responses_api.md)。

## 接口

```http
POST /responses
```

operationId: `createResponse`

### 服务地址

- `https://api.deepseek.com`

## 身份验证

### 验证方式 1

**bearerAuth**

`Authorization: Bearer <API_KEY>`

```yaml
type: http
scheme: bearer
```

## 请求体

**必填**

### application/json

- 引用: [CreateResponseRequest](#schema-create-response-request) (`#/components/schemas/CreateResponseRequest`)

## 响应

### 200 (No streaming)

OK, 返回一个 `response` 对象。

#### application/json

- 引用: [Response](#schema-response) (`#/components/schemas/Response`)

**示例**

```json
{
  "id": "387b53c1-7ad2-4259-9298-61b139f5bd20",
  "object": "response",
  "created_at": 1793084842,
  "status": "completed",
  "background": false,
  "completed_at": 1793084843,
  "content_filters": null,
  "error": null,
  "frequency_penalty": 0.0,
  "incomplete_details": null,
  "instructions": null,
  "max_output_tokens": 32,
  "max_tool_calls": null,
  "model": "deepseek-flash",
  "moderation": null,
  "output": [
    {
      "type": "message",
      "id": "4e706f35-6ea2-468a-bd19-a9c1adc54726",
      "status": "completed",
      "content": [
        {
          "type": "output_text",
          "annotations": [],
          "logprobs": [],
          "text": "Hello! 👋 How can I help you today?"
        }
      ],
      "phase": "final_answer",
      "role": "assistant"
    }
  ],
  "parallel_tool_calls": true,
  "presence_penalty": 0.0,
  "previous_response_id": null,
  "prompt_cache_key": null,
  "prompt_cache_retention": null,
  "reasoning": {
    "effort": "none",
    "summary": null
  },
  "safety_identifier": null,
  "service_tier": "default",
  "store": false,
  "temperature": 1.0,
  "text": {
    "format": {
      "type": "text"
    },
    "verbosity": null
  },
  "tool_choice": "auto",
  "tools": [],
  "top_logprobs": 0,
  "top_p": 1.0,
  "truncation": "disabled",
  "usage": {
    "input_tokens": 7,
    "input_tokens_details": {
      "cached_tokens": 0
    },
    "output_tokens": 11,
    "output_tokens_details": {
      "reasoning_tokens": 0
    },
    "total_tokens": 18
  },
  "user": null,
  "metadata": {}
}
```

### 200 (Streaming)

OK, 返回语义化的流式 SSE 事件序列。每个事件带有表示事件类型的 `event` 字段和递增的 `sequence_number`。最后一个事件是 `response.completed` / `response.incomplete` / `response.failed`（没有 `data: [DONE]` 消息）。完整事件列表请参考 [Responses API 指南](https://api-docs.deepseek.com/zh-cn/guides/responses_api.md#streaming)。

#### text/event-stream

- `type`: `array`

- **数组元素**

  - `type`: `object`

**示例**

```
event: response.created
data: {"type":"response.created","response":{"id":"19460ef2-58da-46c5-9510-9e7b04c47e4a","object":"response","created_at":1793084908,"status":"in_progress","background":false,"completed_at":null,"content_filters":null,"error":null,"frequency_penalty":0.0,"incomplete_details":null,"instructions":null,"max_output_tokens":32,"max_tool_calls":null,"model":"deepseek-flash","moderation":null,"output":[],"parallel_tool_calls":true,"presence_penalty":0.0,"previous_response_id":null,"prompt_cache_key":null,"prompt_cache_retention":null,"reasoning":{"effort":"none","summary":null},"safety_identifier":null,"service_tier":"default","store":false,"temperature":1.0,"text":{"format":{"type":"text"},"verbosity":null},"tool_choice":"auto","tools":[],"top_logprobs":0,"top_p":1.0,"truncation":"disabled","usage":null,"user":null,"metadata":{}},"sequence_number":0}

event: response.in_progress
data: {"type":"response.in_progress","response":{"id":"19460ef2-58da-46c5-9510-9e7b04c47e4a","object":"response","created_at":1793084908,"status":"in_progress","background":false,"completed_at":null,"content_filters":null,"error":null,"frequency_penalty":0.0,"incomplete_details":null,"instructions":null,"max_output_tokens":32,"max_tool_calls":null,"model":"deepseek-flash","moderation":null,"output":[],"parallel_tool_calls":true,"presence_penalty":0.0,"previous_response_id":null,"prompt_cache_key":null,"prompt_cache_retention":null,"reasoning":{"effort":"none","summary":null},"safety_identifier":null,"service_tier":"default","store":false,"temperature":1.0,"text":{"format":{"type":"text"},"verbosity":null},"tool_choice":"auto","tools":[],"top_logprobs":0,"top_p":1.0,"truncation":"disabled","usage":null,"user":null,"metadata":{}},"sequence_number":1}

event: response.output_item.added
data: {"type":"response.output_item.added","item":{"type":"message","id":"7b653eb5-6012-4257-ba1a-4e9dce20bd25","status":"in_progress","content":[],"phase":"final_answer","role":"assistant"},"output_index":0,"sequence_number":2}

event: response.content_part.added
data: {"type":"response.content_part.added","content_index":0,"item_id":"7b653eb5-6012-4257-ba1a-4e9dce20bd25","output_index":0,"part":{"type":"output_text","annotations":[],"logprobs":[],"text":""},"sequence_number":3}

event: response.output_text.delta
data: {"type":"response.output_text.delta","content_index":0,"delta":"Hello","item_id":"7b653eb5-6012-4257-ba1a-4e9dce20bd25","logprobs":[],"output_index":0,"sequence_number":4}

event: response.output_text.delta
data: {"type":"response.output_text.delta","content_index":0,"delta":"!","item_id":"7b653eb5-6012-4257-ba1a-4e9dce20bd25","logprobs":[],"output_index":0,"sequence_number":5}

event: response.output_text.delta
data: {"type":"response.output_text.delta","content_index":0,"delta":" 👋","item_id":"7b653eb5-6012-4257-ba1a-4e9dce20bd25","logprobs":[],"output_index":0,"sequence_number":6}

event: response.output_text.delta
data: {"type":"response.output_text.delta","content_index":0,"delta":" How","item_id":"7b653eb5-6012-4257-ba1a-4e9dce20bd25","logprobs":[],"output_index":0,"sequence_number":7}

event: response.output_text.delta
data: {"type":"response.output_text.delta","content_index":0,"delta":" can","item_id":"7b653eb5-6012-4257-ba1a-4e9dce20bd25","logprobs":[],"output_index":0,"sequence_number":8}

event: response.output_text.delta
data: {"type":"response.output_text.delta","content_index":0,"delta":" I","item_id":"7b653eb5-6012-4257-ba1a-4e9dce20bd25","logprobs":[],"output_index":0,"sequence_number":9}

event: response.output_text.delta
data: {"type":"response.output_text.delta","content_index":0,"delta":" help","item_id":"7b653eb5-6012-4257-ba1a-4e9dce20bd25","logprobs":[],"output_index":0,"sequence_number":10}

event: response.output_text.delta
data: {"type":"response.output_text.delta","content_index":0,"delta":" you","item_id":"7b653eb5-6012-4257-ba1a-4e9dce20bd25","logprobs":[],"output_index":0,"sequence_number":11}

event: response.output_text.delta
data: {"type":"response.output_text.delta","content_index":0,"delta":" today","item_id":"7b653eb5-6012-4257-ba1a-4e9dce20bd25","logprobs":[],"output_index":0,"sequence_number":12}

event: response.output_text.delta
data: {"type":"response.output_text.delta","content_index":0,"delta":"?","item_id":"7b653eb5-6012-4257-ba1a-4e9dce20bd25","logprobs":[],"output_index":0,"sequence_number":13}

event: response.output_text.done
data: {"type":"response.output_text.done","content_index":0,"item_id":"7b653eb5-6012-4257-ba1a-4e9dce20bd25","logprobs":[],"output_index":0,"sequence_number":14,"text":"Hello! 👋 How can I help you today?"}

event: response.content_part.done
data: {"type":"response.content_part.done","content_index":0,"item_id":"7b653eb5-6012-4257-ba1a-4e9dce20bd25","output_index":0,"part":{"type":"output_text","annotations":[],"logprobs":[],"text":"Hello! 👋 How can I help you today?"},"sequence_number":15}

event: response.output_item.done
data: {"type":"response.output_item.done","item":{"type":"message","id":"7b653eb5-6012-4257-ba1a-4e9dce20bd25","status":"completed","content":[{"type":"output_text","annotations":[],"logprobs":[],"text":"Hello! 👋 How can I help you today?"}],"phase":"final_answer","role":"assistant"},"output_index":0,"sequence_number":16}

event: response.completed
data: {"type":"response.completed","response":{"id":"19460ef2-58da-46c5-9510-9e7b04c47e4a","object":"response","created_at":1793084908,"status":"completed","background":false,"completed_at":1793084909,"content_filters":null,"error":null,"frequency_penalty":0.0,"incomplete_details":null,"instructions":null,"max_output_tokens":32,"max_tool_calls":null,"model":"deepseek-flash","moderation":null,"output":[{"type":"message","id":"7b653eb5-6012-4257-ba1a-4e9dce20bd25","status":"completed","content":[{"type":"output_text","annotations":[],"logprobs":[],"text":"Hello! 👋 How can I help you today?"}],"phase":"final_answer","role":"assistant"}],"parallel_tool_calls":true,"presence_penalty":0.0,"previous_response_id":null,"prompt_cache_key":null,"prompt_cache_retention":null,"reasoning":{"effort":"none","summary":null},"safety_identifier":null,"service_tier":"default","store":false,"temperature":1.0,"text":{"format":{"type":"text"},"verbosity":null},"tool_choice":"auto","tools":[],"top_logprobs":0,"top_p":1.0,"truncation":"disabled","usage":{"input_tokens":7,"input_tokens_details":{"cached_tokens":0},"output_tokens":11,"output_tokens_details":{"reasoning_tokens":0},"total_tokens":18},"user":null,"metadata":{}},"sequence_number":17}
```

## 代码示例

### curl

```bash
# Set DEEPSEEK_API_KEY to your API key before running this example.
curl --silent --show-error --max-time 60 "https://api.deepseek.com/responses" \
  --header "Authorization: Bearer ${DEEPSEEK_API_KEY}" \
  --header "Content-Type: application/json" \
  --data '{
  "model": "deepseek-flash",
  "input": "Say hello.",
  "max_output_tokens": 32,
  "reasoning": {
    "effort": "none"
  },
  "stream": false
}'
```

### Python — OpenAI SDK

```python
# Set DEEPSEEK_API_KEY to your API key before running this example.
import os
from openai import OpenAI

client = OpenAI(
    api_key=os.environ["DEEPSEEK_API_KEY"],
    base_url="https://api.deepseek.com",
    timeout=60.0,
)

response = client.responses.create(
    model="deepseek-flash",
    input="Say hello.",
    max_output_tokens=32,
    reasoning={"effort": "none"},
    stream=False,
)
print(response.model_dump_json(indent=2))
```

## 引用的数据结构

### schema-create-response-request

**CreateResponseRequest** — `#/components/schemas/CreateResponseRequest`

- `type`: `object`

- `required`: `["model"]`

- **字段**

  - **`model`** (必填)

    使用的模型的 ID。请使用 `deepseek-flash` 或 `deepseek-v4-pro`。

    - `type`: `string`

    - `enum`: `["deepseek-flash","deepseek-v4-pro"]`

    **示例**

    ```json
    "deepseek-flash"
    ```

  - **`input`** (可选)

    模型的输入。既可以传纯字符串（视作一条 `user` 消息），也可以传输入 item 列表。

    支持的输入 item 类型为 `message` / `function_call` / `function_call_output` / `custom_tool_call` / `custom_tool_call_output` / `reasoning` / `web_search_call`，其他类型会被忽略。消息角色支持 `user` / `assistant` / `system` / `developer`（`developer` 在 `deepseek-flash` 上视同 `system`，在 `deepseek-v4-pro` 上视同 `user`）。使用 `deepseek-flash` 模型时，`user` / `system` / `developer` 消息 item 以及 `function_call_output` / `custom_tool_call_output` item 的 `output` 中支持 `input_image` 内容块；`assistant` 消息中的图片将返回 `400` 错误。`user` 消息中支持用 `input_file` 的 `file_id` 引用已上传的文件，见[文件解析](https://api-docs.deepseek.com/zh-cn/guides/document_parsing.md)；`file_url` 不支持。

    `input` 与 `instructions` 至少传一个。

    使用联网工具或 `safety_check` 为 `high` 时，部分输入 item 不被接受。详见 [联网工具指南](https://api-docs.deepseek.com/zh-cn/guides/web_tools.md#responses-api)。

    - `nullable`: `true`

    - **oneOf**

      - oneOf 1

        - title: `Text input`

        - `type`: `string`

      - oneOf 2

        - title: `Input item list`

        - `type`: `array`

        - **数组元素**

          - 引用: [ResponseInputItem](#schema-response-input-item) (`#/components/schemas/ResponseInputItem`)

  - **`instructions`** (可选)

    系统级指令，作为模型上下文中的第一条 system 消息。

    - `type`: `string`

    - `nullable`: `true`

  - **`reasoning`** (可选)

    思考模式配置。

    - `type`: `object`

    - `nullable`: `true`

    - **字段**

      - **`effort`** (可选)

        控制思考模式开关与思考强度。`none` 关闭思考模式；`low` / `high` / `max` 开启思考模式。不传时使用模型默认的思考行为（默认开启）。出于兼容考虑，`minimal` 映射为 `low`，`medium` / `xhigh` 映射为 `high`。

        - `type`: `string`

        - `enum`: `["none","low","high","max"]`

  - **`max_output_tokens`** (可选)

    响应可生成的 token 数上限，包含可见的输出 token 与思维链 token。

    - `type`: `integer`

    - `nullable`: `true`

    **示例**

    ```json
    4096
    ```

  - **`stream`** (可选)

    如果设置为 `true`，响应将以语义化的流式 SSE 事件返回。最后一个事件是 `response.completed` / `response.incomplete` / `response.failed`（没有 `data: [DONE]` 消息）。完整事件列表请参考 [Responses API 指南](https://api-docs.deepseek.com/zh-cn/guides/responses_api.md#streaming)。

    - `type`: `boolean`

    - `nullable`: `true`

    - `default`: `false`

  - **`temperature`** (可选)

    采样温度，介于 0 和 2 之间。更高的值（如 0.8）会使输出更随机，而更低的值（如 0.2）会使其更加集中和确定。思考模式下不生效。

    - `type`: `number`

    - `minimum`: `0`

    - `maximum`: `2`

    - `default`: `1`

    - `nullable`: `true`

    **示例**

    ```json
    1
    ```

  - **`top_p`** (可选)

    用于调节输出的随机性，可替代 `temperature`。实际生效值取决于模型和模式：

    - `deepseek-flash`：思考模式下默认值为 1.0，低于 0.95 的取值会按 0.95 处理；非思考模式下恒为 1.0，传入的值会被忽略。

    - `deepseek-v4-pro`：思考模式下，启用工具调用时恒为 0.95，否则恒为 1.0，传入的值会被忽略。即使提供了 `tools`，将 `tool_choice` 设为 `none` 也会关闭工具调用。非思考模式下使用传入的值，默认值为 1.0。

    - `type`: `number`

    - `minimum`: `0`

    - `maximum`: `1`

    - `nullable`: `true`

    **示例**

    ```json
    1
    ```

  - **`text`** (可选)

    文本输出配置。

    - `type`: `object`

    - `nullable`: `true`

    - **字段**

      - **`format`** (可选)

        输出格式。`{"type": "text"}`（默认）为纯文本输出；`{"type": "json_object"}` 为 JSON 模式；`{"type": "json_schema", "name": ..., "schema": ...}` 为结构化输出，输出符合给定的 JSON Schema。

        省略 `text`、将 `text` 设为 `null`，或省略 `format` 时，使用纯文本输出。提供 `format` 对象时，`type` 必填。

        - `type`: `object`

        - `required`: `["type"]`

        - **字段**

          - **`type`** (必填)

            - `type`: `string`

            - `enum`: `["text","json_object","json_schema"]`

          - **`name`** (可选)

            schema 的名称。`type` 为 `json_schema` 时必填。

            - `type`: `string`

          - **`schema`** (可选)

            输出必须符合的 JSON Schema。`type` 为 `json_schema` 时必填。

            - `type`: `object`

  - **`tools`** (可选)

    模型可能会调用的工具的列表。函数名必须非空、不超过 128 个字符、匹配 `^[a-zA-Z0-9_-]+$`，且所有工具的名称必须唯一。支持 `function` 工具；`web_search` 在 `deepseek-v4-pro` 下支持，其余内置工具类型会被忽略。详情请参考 [Responses API 指南](https://api-docs.deepseek.com/zh-cn/guides/responses_api.md)。

    - `type`: `array`

    - `nullable`: `true`

    - **数组元素**

      - 引用: [ResponseTool](#schema-response-tool) (`#/components/schemas/ResponseTool`)

    **示例**

    ```json
    null
    ```

  - **`tool_choice`** (可选)

    控制模型调用工具的行为。

    `none` 意味着模型不会调用任何工具，而是生成一条消息。

    `auto`（默认）意味着模型可以选择生成一条消息或调用一个或多个工具。

    `required` 意味着模型必须调用一个或多个工具。

    通过 `{"type": "function", "name": "my_function"}` 指定特定工具，会强制模型调用该工具。

    - `nullable`: `true`

    - **oneOf**

      - oneOf 1

        - title: `Tool choice mode`

        - `type`: `string`

        - `enum`: `["none","auto","required"]`

      - oneOf 2

        - title: `Named tool choice`

        Forces the model to call the specified tool. `name` is required when `type` is `function`.

        - `type`: `object`

        - `required`: `["type"]`

        - **字段**

          - **`type`** (必填)

            - `type`: `string`

            - `enum`: `["function"]`

          - **`name`** (可选)

            The name of the function to call. Required when `type` is `function`.

            - `type`: `string`

  - **`top_logprobs`** (可选)

    一个介于 0 到 20 之间的整数 N，指定每个输出位置返回输出概率 top N 的 token，且返回这些 token 的对数概率。

    - `type`: `integer`

    - `minimum`: `0`

    - `maximum`: `20`

    - `nullable`: `true`

    - `default`: `null`

    **示例**

    ```json
    null
    ```

  - **`user`** (可选)

    自定义终端用户标识。建议仅使用英文字母、数字、下划线（`_`）和连字符（`-`）。UTF-8 编码后不得超过 512 字节。请勿在其中包含用户隐私信息。

    - 可用于区分您业务侧的用户身份，以帮助我们进行内容安全审核；也可用于 KVCache 隔离与调度隔离。详情请参考[限速与用户隔离](https://api-docs.deepseek.com/zh-cn/quick_start/rate_limit.md)

    * `type`: `string`

    * `nullable`: `true`

  - **`safety_check`** (可选)

    DeepSeek 扩展字段。`level` 设为 `high` 时开启基础风控，审核模型的输入与输出。详见 [模型内容安全指南](https://api-docs.deepseek.com/zh-cn/guides/safety.md)。

    - `type`: `object`

    - `nullable`: `true`

    - `required`: `["level"]`

    - **字段**

      - **`level`** (必填)

        `high`：开启基础风控。`low`：默认检查，与不传 `safety_check` 相同。

        - `type`: `string`

        - `enum`: `["low","high"]`

  - **`title`** (可选)

    DeepSeek 扩展字段。传入 `{"enabled": true}`，响应会带上简短的会话标题。详见 [会话标题指南](https://api-docs.deepseek.com/zh-cn/guides/conversation_title.md)。

    - `type`: `object`

    - `nullable`: `true`

    - `required`: `["enabled"]`

    - **字段**

      - **`enabled`** (必填)

        是否返回会话标题。

        - `type`: `boolean`

    - **additionalProperties**

      - 数据结构: `false`

### schema-response

**Response** — `#/components/schemas/Response`

OpenAI Responses API 格式的模型响应对象。

- `type`: `object`

- `required`: `["id","object","created_at","status","model","output"]`

- **字段**

  - **`id`** (必填)

    该响应的唯一标识符。

    - `type`: `string`

  - **`object`** (必填)

    object 的类型，其值恒为 `response`。

    - `type`: `string`

    - `enum`: `["response"]`

  - **`created_at`** (必填)

    标志响应创建时间的 Unix 时间戳（以秒为单位）。

    - `type`: `integer`

  - **`status`** (必填)

    响应的状态。

    - `type`: `string`

    - `enum`: `["in_progress","completed","incomplete","failed"]`

  - **`error`** (可选)

    响应失败时的错误对象，包含 `code` 和 `message` 字段。

    - `type`: `object`

    - `nullable`: `true`

  - **`incomplete_details`** (可选)

    响应不完整的原因详情。`reason` 字段可能为 `max_output_tokens` 或 `content_filter`。

    - `type`: `object`

    - `nullable`: `true`

    - **字段**

      - **`reason`** (可选)

        - `type`: `string`

        - `enum`: `["max_output_tokens","content_filter"]`

  - **`model`** (必填)

    生成该响应的模型。

    - `type`: `string`

  - **`output`** (必填)

    模型生成的输出 item 列表。思考模式下，思维链以 `reasoning` item 的形式在 `message` item 之前返回。函数调用以 `function_call` item 返回。

    - `type`: `array`

    - **数组元素**

      - 引用: [ResponseOutputItem](#schema-response-output-item) (`#/components/schemas/ResponseOutputItem`)

  - **`title`** (可选)

    DeepSeek 扩展字段。会话标题，只在请求开启了 `title` 且生成了标题时出现。流式响应请从结束事件的 `response.title` 中读取。

    - `type`: `string`

  - **`usage`** (可选)

    - 引用: [ResponseUsage](#schema-response-usage) (`#/components/schemas/ResponseUsage`)

### schema-response-input-item

**ResponseInputItem** — `#/components/schemas/ResponseInputItem`

输入 item。支持的类型为 `message` / `function_call` / `function_call_output` / `custom_tool_call` / `custom_tool_call_output` / `reasoning` / `web_search_call`，其他类型会被忽略。`reasoning` 和 `function_call` item 会由服务端自动归并到相邻的 assistant 消息中。上一次响应返回的 `web_search_call` item 可以回传，以将其搜索结果还原到上下文中。多轮对话中的回传规则请参考 [Responses API 指南](https://api-docs.deepseek.com/zh-cn/guides/responses_api.md#multi-turn-conversation)。

- `type`: `object`

- **字段**

  - **`type`** (可选)

    输入 item 的类型。对于 `message` item，如果传了 `role`，此字段可省略。`custom_tool_call` / `custom_tool_call_output` item 配合 `apply_patch` custom 工具使用。`web_search_call` item 用于还原上一次服务端搜索的结果。

    - `type`: `string`

    - `enum`: `["message","function_call","function_call_output","custom_tool_call","custom_tool_call_output","reasoning","web_search_call"]`

  - **`role`** (可选)

    用于 `message` item。消息作者的角色。`developer` 在 `deepseek-flash` 上视同 `system`，在 `deepseek-v4-pro` 上视同 `user`。

    - `type`: `string`

    - `enum`: `["user","assistant","system","developer"]`

  - **`content`** (可选)

    用于 `message` item 时为消息内容，可以是纯字符串或 `input_text` / `output_text` / `input_image` 内容块列表。用于 `reasoning` item 时为 `reasoning_text` 内容块列表。

    - **oneOf**

      - oneOf 1

        - title: `Text content`

        - `type`: `string`

      - oneOf 2

        - title: `Array of content parts`

        - `type`: `array`

        - **数组元素**

          - 引用: [ResponseInputContentPart](#schema-response-input-content-part) (`#/components/schemas/ResponseInputContentPart`)

  - **`call_id`** (可选)

    用于 `function_call` / `function_call_output` item。将函数调用与其结果配对的 ID。必须非空且唯一，且每个 `function_call` 必须有对应的 `function_call_output`。

    - `type`: `string`

  - **`name`** (可选)

    用于 `function_call` item。要调用的函数的名称。

    - `type`: `string`

  - **`arguments`** (可选)

    用于 `function_call` item。调用函数的入参，格式为 JSON。

    - `type`: `string`

  - **`output`** (可选)

    用于 `function_call_output` / `custom_tool_call_output` item。工具调用的结果，可以是纯字符串或 `input_text` / `input_image` 内容块列表。

    - **oneOf**

      - oneOf 1

        - title: `Text output`

        - `type`: `string`

      - oneOf 2

        - title: `Array of content parts`

        - `type`: `array`

        - **数组元素**

          - 引用: [ResponseInputContentPart](#schema-response-input-content-part) (`#/components/schemas/ResponseInputContentPart`)

### schema-response-tool

**ResponseTool** — `#/components/schemas/ResponseTool`

模型可能会调用的工具。`function` 定义一个函数工具——注意在 Responses API 格式中，`name` / `parameters` 等函数字段位于工具对象的顶层。`custom` 定义 custom 工具，仅支持 `{"type": "custom", "name": "apply_patch"}`，其他名称会返回 `400` 错误。`web_search` 启用服务端网页搜索，在 `deepseek-v4-pro` 下支持。`deepseek_web_search` 与 `deepseek_read_url` 是联网工具，不能与 `web_search` 同时声明，详见 [联网工具指南](https://api-docs.deepseek.com/zh-cn/guides/web_tools.md#responses-api)。

- `type`: `object`

- `required`: `["type"]`

- **字段**

  - **`type`** (必填)

    工具类型：`function`、`custom`（仅 `apply_patch`）、`web_search`（仅 `deepseek-v4-pro`）、`deepseek_web_search` 或 `deepseek_read_url`。

    - `type`: `string`

    - `enum`: `["function","custom","web_search","deepseek_web_search","deepseek_read_url"]`

  - **`name`** (可选)

    用于 `function` / `custom` 工具。工具的名称。对于函数工具，必须非空、不超过 128 个字符、匹配 `^[a-zA-Z0-9_-]+$`，且所有工具的名称必须唯一。对于 custom 工具，必须为 `apply_patch`。

    - `type`: `string`

  - **`description`** (可选)

    用于 `function` 工具。函数功能的描述，供模型理解何时以及如何调用该函数。

    - `type`: `string`

  - **`parameters`** (可选)

    - 引用: [FunctionParameters](#schema-function-parameters) (`#/components/schemas/FunctionParameters`)

### schema-response-output-item

**ResponseOutputItem** — `#/components/schemas/ResponseOutputItem`

模型生成的输出 item。`message` item 承载模型回复文本，`reasoning` item 承载思维链，`function_call` item 表示函数调用，`custom_tool_call` item 表示 custom 工具调用（如 `apply_patch`），`web_search_call` item 表示服务端网页搜索（仅 `deepseek-v4-pro`），`tool_search_call` item 表示工具搜索。使用联网工具时，每次工具调用都是一个 `web_search_call` item，并带扩展字段 `kind`、`page`、`signature` 与 `error`。

- `type`: `object`

- **字段**

  - **`type`** (可选)

    输出 item 的类型：`message` / `reasoning` / `function_call` / `custom_tool_call` / `web_search_call` / `tool_search_call`。

    - `type`: `string`

    - `enum`: `["message","reasoning","function_call","custom_tool_call","web_search_call","tool_search_call"]`

  - **`id`** (可选)

    输出 item 的唯一 ID。

    - `type`: `string`

  - **`status`** (可选)

    输出 item 的状态。

    - `type`: `string`

    - `enum`: `["in_progress","completed","incomplete"]`

  - **`role`** (可选)

    用于 `message` item。其值恒为 `assistant`。

    - `type`: `string`

    - `enum`: `["assistant"]`

  - **`content`** (可选)

    用于 `message` item 时为 `output_text` 内容块列表。用于 `reasoning` item 时为 `reasoning_text` 内容块列表，以明文承载思维链内容。

    - `type`: `array`

    - **数组元素**

      - `type`: `object`

      - **字段**

        - **`type`** (可选)

          - `type`: `string`

          - `enum`: `["output_text","reasoning_text"]`

        - **`text`** (可选)

          - `type`: `string`

  - **`call_id`** (可选)

    用于 `function_call` / `custom_tool_call` / `tool_search_call` item。将工具结果回传给 API 时使用的标识符。

    - `type`: `string`

  - **`name`** (可选)

    用于 `function_call` / `custom_tool_call` item。要调用的函数或 custom 工具的名称。

    - `type`: `string`

  - **`arguments`** (可选)

    用于 `function_call` item。模型生成的调用函数的入参，格式为 JSON。请注意，模型并不总是生成有效的 JSON，且可能会虚构出您的函数模式中未定义的参数。在调用函数之前，请在您的代码中验证这些入参是否有效。

    - `type`: `string`

  - **`kind`** (可选)

    - 引用: [AnthropicDeepSeekWebToolsKind](#schema-anthropic-deep-seek-web-tools-kind) (`#/components/schemas/AnthropicDeepSeekWebToolsKind`)

  - **`action`** (可选)

    用于 `web_search_call` item。`type` 为 `search`（带 `queries` 与 `sources`）、`open_page`（带 `url`）或 `find_in_page`（带 `url` 与 `pattern`）。使用联网工具时，`sources` 的每个元素还带有用于展示的 `title`、`snippet`、`site_name`、`site_icon` 与 `published_at`。

    - `type`: `object`

  - **`page`** (可选)

    DeepSeek 扩展字段，用于打开或阅读了网页的 `web_search_call` item：`url`、`title` 与 `published_at`。不返回网页正文。

    - `type`: `object`

  - **`signature`** (可选)

    DeepSeek 扩展字段，用于成功的 `read_link` 调用，是加密后的网页。请原样回传，网页才会拼回上下文。

    - `type`: `string`

  - **`error`** (可选)

    DeepSeek 扩展字段，用于失败的 `web_search_call` item（`status` 为 `failed`）。`code` 为 `unavailable`、`url_not_accessible` 或 `url_not_allowed`。

    - `type`: `object`

### schema-response-usage

**ResponseUsage** — `#/components/schemas/ResponseUsage`

该响应的 token 和服务端工具用量统计信息。流式响应的初始事件中为 null，请读取最终响应的 usage。

- `type`: `object`

- `nullable`: `true`

- `required`: `["input_tokens","output_tokens","total_tokens"]`

- **字段**

  - **`input_tokens`** (必填)

    输入 token 数。

    - `type`: `integer`

  - **`input_tokens_details`** (可选)

    输入 token 的细分信息。

    - `type`: `object`

    - **字段**

      - **`cached_tokens`** (可选)

        命中上下文缓存的输入 token 数。参考[上下文硬盘缓存](https://api-docs.deepseek.com/zh-cn/guides/kv_cache.md)。

        - `type`: `integer`

  - **`output_tokens`** (必填)

    输出 token 数。

    - `type`: `integer`

  - **`output_tokens_details`** (可选)

    输出 token 的细分信息。

    - `type`: `object`

    - **字段**

      - **`reasoning_tokens`** (可选)

        模型生成的思维链 token 数。

        - `type`: `integer`

      - **`safety_tokens`** (可选)

        DeepSeek 扩展字段。安全审核产生的 token 数，已包含在 `output_tokens` 中。仅在 `safety_check` 为 `high` 时返回。

        - `type`: `integer`

        - `minimum`: `0`

  - **`total_tokens`** (必填)

    该请求使用的 token 总数（输入 + 输出）。

    - `type`: `integer`

  - **`server_tool_use`** (可选)

    - title: `服务端工具用量`

    成功的服务端网页搜索和抓取次数，与 token 用量分别统计。仅在其中一项计数大于零时返回，否则省略整个字段。返回时各个子字段均存在。流式响应中，请从终止事件 response.completed、response.incomplete 或 response.failed 的 response.usage 读取。这些计数是累计值，不是增量。详见[服务端工具用量](https://api-docs.deepseek.com/zh-cn/guides/responses_api.md#server-tool-usage)。

    - `type`: `object`

    - `required`: `["tool_search","tool_fetch"]`

    - **字段**

      - **`tool_search`** (必填)

        成功的网页搜索查询次数。一次搜索含多个查询时可能累计多次；失败或结果为空的查询不计入。此字段可作为计费依据。

        - `type`: `integer`

        - `minimum`: `0`

      - **`tool_fetch`** (必填)

        通过 open\_page 或 find\_in\_page 成功抓取网页的次数。失败的尝试不计入。此字段可作为计费依据。使用联网工具时，页内查找不计入。

        - `type`: `integer`

        - `minimum`: `0`

      - **`tool_image_search`** (可选)

        使用联网工具时返回。成功的图片搜索次数。此字段可作为计费依据。

        - `type`: `integer`

        - `minimum`: `0`

    **示例**

    ```json
    {
      "tool_search": 1,
      "tool_fetch": 0
    }
    ```

### schema-response-input-content-part

**ResponseInputContentPart** — `#/components/schemas/ResponseInputContentPart`

- **oneOf**

  - oneOf 1

    - 引用: [ResponseInputTextContentPart](#schema-response-input-text-content-part) (`#/components/schemas/ResponseInputTextContentPart`)

  - oneOf 2

    - 引用: [ResponseInputImageContentPart](#schema-response-input-image-content-part) (`#/components/schemas/ResponseInputImageContentPart`)

  - oneOf 3

    - 引用: [ResponseInputFileContentPart](#schema-response-input-file-content-part) (`#/components/schemas/ResponseInputFileContentPart`)

  - oneOf 4

    - 引用: [ResponseInputReasoningTextContentPart](#schema-response-input-reasoning-text-content-part) (`#/components/schemas/ResponseInputReasoningTextContentPart`)

### schema-function-parameters

**FunctionParameters** — `#/components/schemas/FunctionParameters`

function 的输入参数，以 JSON Schema 对象描述。请参阅[Tool Calls 指南](https://api-docs.deepseek.com/zh-cn/guides/tool_calls.md)获取示例，并参阅[JSON Schema 参考](https://json-schema.org/understanding-json-schema/)了解有关格式的文档。省略 `parameters` 会定义一个参数列表为空的 function。

- `type`: `object`

- **additionalProperties**

  - 数据结构: `true`

### schema-anthropic-deep-seek-web-tools-kind

**AnthropicDeepSeekWebToolsKind** — `#/components/schemas/AnthropicDeepSeekWebToolsKind`

- title: `联网工具能力`

DeepSeek 扩展字段，出现在联网工具产生的块上：`search` 表示联网搜索，`deep_search` 表示深度思考搜索，`read_link` 表示阅读用户消息中的链接，`image_search` 与 `deep_image_search`（深度思考搜索中）表示图片搜索。回传时请原样保留。

- `type`: `string`

- `enum`: `["search","deep_search","read_link","image_search","deep_image_search"]`

### schema-response-input-text-content-part

**ResponseInputTextContentPart** — `#/components/schemas/ResponseInputTextContentPart`

- title: `文本内容块`

- `type`: `object`

- `required`: `["type","text"]`

- **字段**

  - **`type`** (必填)

    内容块的类型。

    - `type`: `string`

    - `enum`: `["input_text","output_text"]`

  - **`text`** (必填)

    文本内容。

    - `type`: `string`

### schema-response-input-image-content-part

**ResponseInputImageContentPart** — `#/components/schemas/ResponseInputImageContentPart`

- title: `图片内容块`

图片内容块。会作为真实图片处理。可以出现在 `user` / `system` / `developer` 消息 item 以及 `function_call_output` / `custom_tool_call_output` item 的 `output` 中；`assistant` 消息中的图片将返回 `400` 错误。

- `type`: `object`

- `required`: `["type"]`

- **字段**

  - **`type`** (必填)

    内容块的类型，此场景下为 `input_image`。

    - `type`: `string`

    - `enum`: `["input_image"]`

  - **`image_url`** (可选)

    图片来源，可以是图片的 `http(s)` URL（最多 8192 个字符）或 base64 编码的 data URL（`data:image/jpeg;base64,...`）。支持的格式：JPEG、PNG、GIF、WebP。与 `file_id` 互斥：两者都不传返回 `400` 错误（"input\_image must have image\_url or file\_id"），两者都传返回 `400` 错误（"input\_image cannot have both image\_url and file\_id"）。

    - `type`: `string`

  - **`detail`** (可选)

    控制图片的处理方式。`low` 将图片缩小到 512x512（更快、更省 token）；`high`、`original` 与 `auto` 保留原图。设置 `file_id` 时该字段被忽略。

    - `type`: `string`

    - `enum`: `["low","high","original","auto"]`

  - **`file_id`** (可选)

    通过 [Files API](https://api-docs.deepseek.com/zh-cn/guides/files_api.md) 上传的图片文件 ID，形如 `file-api-...`。与 `image_url` 互斥；设置 `file_id` 时 `detail` 被忽略。

    - `type`: `string`

### schema-response-input-file-content-part

**ResponseInputFileContentPart** — `#/components/schemas/ResponseInputFileContentPart`

- title: `文件内容块`

引用通过 Files API 上传的文件。可以出现在 `user` 消息 item 中。详见 [文件解析指南](https://api-docs.deepseek.com/zh-cn/guides/document_parsing.md#document-qa)。

- `type`: `object`

- `required`: `["type","file_id"]`

- **字段**

  - **`type`** (必填)

    内容块的类型，此场景下为 `input_file`。

    - `type`: `string`

    - `enum`: `["input_file"]`

  - **`file_id`** (必填)

    已上传文件的 ID，形如 `file-api-...`。

    - `type`: `string`

### schema-response-input-reasoning-text-content-part

**ResponseInputReasoningTextContentPart** — `#/components/schemas/ResponseInputReasoningTextContentPart`

- title: `推理文本内容块`

- `type`: `object`

- `required`: `["type","text"]`

- **字段**

  - **`type`** (必填)

    内容块的类型，此场景下为 `reasoning_text`。

    - `type`: `string`

    - `enum`: `["reasoning_text"]`

  - **`text`** (必填)

    思维链文本内容。

    - `type`: `string`
