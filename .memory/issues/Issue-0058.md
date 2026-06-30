RFC:
中文 cookbook 需要可重复生成一本总 PDF，并保持教材阅读顺序。

问题是 `docs/cookbook-zh` 已经有根目录章节、学期/月/周索引和多个 `days-*` 子目录，但缺少成书输出链路。简单按文件名合并会打乱 cookbook 顺序，依赖 LaTeX 又会增加本机环境风险。

目标：
- 递归包含 `docs/cookbook-zh` 下全部 Markdown。
- 生成一本总 PDF：`docs/cookbook-zh/atlas-cookbook-zh.pdf`。
- 通过脚本和 `Makefile` 固化生成入口。
- 保留重复日页，不在生成阶段自行去重。

非目标：
- 不修改 cookbook 正文内容。
- 不生成分册。
- 不引入 Node、LaTeX、WeasyPrint 或 wkhtmltopdf。

Interface:
新增命令入口：

```bash
make cookbook-pdf
```

等价于：

```bash
go run ./scripts/cookbookpdf
```

脚本参数：
- `--input`：默认 `docs/cookbook-zh`
- `--output`：默认 `docs/cookbook-zh/atlas-cookbook-zh.pdf`
- `--workdir`：指定临时目录
- `--keep-workdir`：保留中间 Markdown 和 HTML

Implementation:
新增 `scripts/cookbookpdf` Go 脚本。

实现要点：
- 使用 `internal/config` 初始化配置。
- 使用 `internal/logger` 输出脚本状态和外部命令信息。
- 递归收集 `.md` 文件。
- 按 cookbook 阅读顺序排序：入口、总览、学期、月、周、日索引、辅助材料、`days-*` 日页。
- 生成临时合并 Markdown 和 HTML。
- 使用 `pandoc` 生成 HTML。
- 使用 Chrome headless 打印 PDF。
- 截断 Chrome 噪声日志，避免输出过长。

Unit Test:
新增 `scripts/cookbookpdf` 测试，覆盖：
- 默认参数解析。
- 拒绝多余位置参数。
- 空目录错误。
- 递归收集 Markdown。
- cookbook 排序规则。
- 外部命令输出截断。

已运行：

```bash
make test
make cookbook-pdf
```

Review:
审查重点：
- 脚本没有直接访问 `os.Getenv`。
- 脚本日志没有绕过 `internal/logger`。
- 新函数满足 Rule of 30。
- 生成链路没有修改 cookbook 正文。
- Chrome headless 的 stderr 噪声不会导致生成失败，只截断记录。

Benchmark:
本 issue 不涉及服务性能路径，不需要 benchmark。

实际生成耗时约 55 秒，输入 Markdown 数量为 497 个，生成 PDF 大小约 28 MB。

Documentation:
新增：
- `docs/rfc/cookbook-pdf-generation.md`
- `docs/cookbook-zh/pdf-generation-interface.md`

并在 `docs/cookbook-zh/README.md` 中增加 PDF 生成说明。

Memory:
记住：中文 cookbook 的 PDF 生成链路采用 `pandoc -> HTML -> Chrome headless PDF`，避免依赖 LaTeX；脚本必须走 `internal/config` 与 `internal/logger`，并按 cookbook 语义排序而不是按文件名直接合并。
