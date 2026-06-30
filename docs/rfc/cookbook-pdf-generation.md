# 中文 Cookbook PDF 生成 RFC

最后更新：2026-06-30

## 背景

`docs/cookbook-zh` 已经形成一套中文 cookbook 教材，但目前只能以 Markdown 文件树阅读。学习者如果想离线阅读、归档、打印或统一评审，需要一份按教材顺序排版的 PDF。

现状问题：

- 目录下有根目录章节，也有 `days-*` 子目录，简单按文件名合并会破坏教材阅读顺序。
- 中文 PDF 不能依赖 LaTeX 环境，因为本机只确认存在 `pandoc` 和 Chrome。
- 生成能力需要可重复执行，并纳入仓库已有 `Makefile` 工作流。

## 目标

- 递归收集 `docs/cookbook-zh` 下全部 Markdown 文件，生成一本总 PDF。
- 默认输出 `docs/cookbook-zh/atlas-cookbook-zh.pdf`。
- 使用稳定排序规则，让总书符合 cookbook 阅读顺序。
- 保留重复文件，避免脚本自行改变“全部递归”的范围。
- 生成脚本具备可测试的收集、排序和参数解析逻辑。

## 非目标

- 不修改 cookbook 正文内容。
- 不生成分册 PDF。
- 不引入 Node、LaTeX、WeasyPrint 或 wkhtmltopdf 依赖。
- 不保证 Markdown 内部相对链接在 PDF 内全部可跳转。

## 方案

新增 `scripts/cookbookpdf` Go 脚本：

1. 解析参数并确定输入目录、输出 PDF 和临时工作目录。
2. 递归收集输入目录中的 `.md` 文件。
3. 按教材顺序排序：
   - `README.md`
   - 使用说明、总览、术语、源码阅读路径
   - 学期章节
   - 月度章节
   - 周度索引与周章节
   - 日度索引
   - 辅助材料
   - `days-*` 子目录日页
4. 生成临时合并 Markdown，并在每个文件前插入来源注释与分页标记。
5. 调用 `pandoc` 转换为 HTML。
6. 调用 Chrome headless 将 HTML 打印为 PDF。

脚本日志必须通过 `internal/logger` 输出；配置初始化使用 `internal/config`。

## 接口

命令：

```bash
go run ./scripts/cookbookpdf
```

可选参数：

- `--input`：Markdown 输入目录，默认 `docs/cookbook-zh`
- `--output`：PDF 输出路径，默认 `docs/cookbook-zh/atlas-cookbook-zh.pdf`
- `--workdir`：临时工作目录，默认自动创建
- `--keep-workdir`：保留临时工作目录，便于排查 HTML 或合并 Markdown

`Makefile` 新增：

```bash
make cookbook-pdf
```

## 验收

- `go test ./...` 通过。
- `make cookbook-pdf` 生成非空 PDF。
- 生成日志输出输入文件数量和输出路径。
- PDF 中文可读，目录和章节分页可扫描。

## 风险

- Chrome 路径可能因系统不同而变化。脚本需要按常见命令名和 macOS 应用路径查找，找不到时明确报错。
- 496 个 Markdown 文件合并后的 PDF 较大，生成耗时可能明显高于普通文档。
- Markdown 中部分相对链接指向仓库外层文档，PDF 中保留文本即可，不作为本次验收阻断项。
