# 原子概念补充

最后更新：2026-06-30

## 这份文档解决什么问题

随着教材越来越细，day 文档里还会残留一批很短、但反复出现的词：

- `main`
- `config.Manager`
- `threat model`
- `review note`

这些词看起来很短，却经常是学习卡点。

这份文档就是给这些“太短以至于前面那些大索引不方便收”的词条补一个稳定落点。

## 代码与结构类

<a id="concept-main"></a>
### `main`

- 在这套教材里，它不是“程序入口”这么简单。
- 你真正要关心的是：
  - `main` 是否足够薄
  - 是否只负责退出码与入口转发
  - 有没有把 bootstrap / runtime / config 初始化重新吸回自己身上

<a id="concept-main-function"></a>
### `main()`

- 它是 `main` 的函数形态提示，重点仍然不是语法。
- 在 Atlas 里看到 `main()` 时，优先检查：
  - 是否只负责调用 command/root 入口
  - 是否把错误转成退出码
  - 是否避免直接初始化 config、logger、runtime 或业务依赖

如果 `main()` 开始承担业务初始化，它通常意味着入口边界正在退化。

<a id="concept-config-manager"></a>
### `config.Manager`

- 它不是普通配置对象。
- 在当前仓库里，它承担的是：
  - 默认值
  - 文件读取
  - 环境覆盖
  - 对外稳定读取接口

如果业务代码绕过它自己去找环境变量，边界就开始泄漏。

<a id="concept-os-getenv"></a>
### `os.Getenv`

- 在这套教材里，它经常作为“配置边界泄漏”的反例出现。
- 你真正要关心的是：
  - 业务代码是否绕过 `config.Manager`
  - 环境变量读取是否还能被测试和验证
  - 默认值、配置文件、环境覆盖和校验是否仍然有统一顺序

项目规则已经明确：不要直接访问 `os.Getenv`，用 `internal/config`。

## 设计与治理类

<a id="concept-threat-model"></a>
### `threat model`

- 它不是一张安全清单。
- 最低要求是能回答：
  - 资产是什么
  - 入口是什么
  - 攻击路径是什么
  - 控制点在哪里

如果一份 threat model 不能映射到真实高风险动作，它就还只是“听起来很安全”的文档。

<a id="concept-review-note"></a>
### `review note`

- 它不是心得体会。
- 最低要有：
  - 结论
  - 主要发现
  - 证据
  - 下一步

具体写法直接看：

- [Review Note 示例](review-note-examples.md)

## 最后一句

这些词越短，越容易被误以为“我懂了”。

真正的标准不是你认得它，而是你能不能把它放回当前的系统边界里解释。
