Decision:
继续把高频 day 阅读提示收口成真实落点，优先处理配置 / 日志、超时 / 错误分类、结构化输出解析和 `docs/ROADMAP.md` 这类高频项。

Reason:
在吃掉大批抽象提示之后，仍有一些高频阅读点长期重复出现：
- database/sql connection management and slog handler internals
- context deadline handling and error classification flow
- OpenAI SDK response parsing and schema validation code

这些提示已经足够高频，值得单独收口，而不应继续留在 day 文档里做纯文本提醒。

Future:
继续吃掉剩余高频提示
把更多 raw 提示替换成稳定链接
进一步减少 day 文档中的纯英文抽象项
