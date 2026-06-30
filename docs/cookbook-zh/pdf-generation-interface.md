# 中文 Cookbook PDF 生成接口

最后更新：2026-06-30

## 命令

在仓库根目录执行：

```bash
make cookbook-pdf
```

等价于：

```bash
go run ./scripts/cookbookpdf
```

默认输入目录是 `docs/cookbook-zh`，默认输出文件是：

```text
docs/cookbook-zh/atlas-cookbook-zh.pdf
```

## 参数

脚本支持以下参数：

- `--input`：Markdown 输入目录。
- `--output`：PDF 输出路径。
- `--workdir`：临时工作目录。
- `--keep-workdir`：保留临时工作目录，便于检查合并 Markdown 和 HTML。

示例：

```bash
go run ./scripts/cookbookpdf \
  --input docs/cookbook-zh \
  --output docs/cookbook-zh/atlas-cookbook-zh.pdf \
  --keep-workdir
```

## 依赖

脚本依赖本机已有工具：

- `pandoc`
- Google Chrome 或 Chromium

本仓库当前采用 `pandoc` 生成 HTML，再用 Chrome headless 打印 PDF。这样可以避免额外依赖 LaTeX，同时保持中文字体渲染稳定。

## 排序规则

PDF 会递归包含 `docs/cookbook-zh` 下所有 Markdown 文件。文件按教材阅读顺序合并：

1. 总入口与使用说明
2. 学期、月度、周度、日度索引
3. 主题辅助材料
4. `days-*` 子目录中的日级页面

脚本不会去重。即使不同目录下存在内容相同的日页，也会按文件路径全部进入 PDF。
