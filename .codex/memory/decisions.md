# DECISIONS

No decision entries yet.

## 中文 Cookbook PDF 生成链路

- 时间：2026-06-30T08:34:02Z
- 内容：docs/cookbook-zh PDF 生成采用 pandoc -> HTML -> Chrome headless PDF，避免依赖 LaTeX。脚本入口是 scripts/cookbookpdf 和 make cookbook-pdf，排序必须按 cookbook 阅读语义处理，不能简单按文件名合并。
- 上下文：docs/cookbook-zh; scripts/cookbookpdf; Makefile
- 触发：完成中文 cookbook PDF 生成脚本与总 PDF 交付
