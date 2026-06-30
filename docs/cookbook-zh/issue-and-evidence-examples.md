# Issue 与 Evidence 示例

最后更新：2026-06-30

## 这份文档解决什么问题

到这里，教材已经告诉你：

- 月 / 周 / 日分别学什么
- 代码和证据该去哪里看
- RFC、Review Note、Memory 该怎么写

但还有一类产物，学习者经常知道“名字”，不知道“长相”：

- Issue Card
- Execution Log
- Review Checkpoints
- Evidence Entry

这份文档就是把这些治理产物写成可以模仿的中文示例。

## 示例 1：Month 01 的 Issue Card 应该怎么读

先对照真实文件：

- [ATLAS-M01-01](../atlas/issues/month-01/atlas-m01-01.md)

### 这个卡片真正要表达什么

一张像样的 issue card，不只是“待办事项”。

它至少要同时表达：

1. 这个问题为什么存在
2. 它本月主线里的位置是什么
3. 什么算完成
4. 需要留下什么证据
5. 当前推进到了哪里

### 中文化阅读模板

```md
## Board Metadata

- status: planned
- owner: learner
- target week: Week 01
- dependencies: 无
- risk level: medium

## Title

建立 Atlas 的 bootstrap 外壳与 runtime 生命周期边界。

## Problem

如果第一月不先收住启动边界，后续配置、日志、db/cache 和 Agent 能力都会散落在入口代码里。

## Scope

- 冻结入口层、bootstrap 层、runtime 层职责
- 打通最小启动 / 停止主路径
- 留下后续扩展 db/cache/http 的挂点

## Out of Scope

- 不在这一张卡里解决所有基础设施接线
- 不在没有 contract 的前提下扩展新抽象

## Acceptance Checklist

- 有 RFC 或 interface note
- 主路径实现存在
- 至少一条验证证据存在
- 文档或 runbook 已更新

## Evidence Expectations

- RFC 草稿或设计说明
- 代码或伪代码产物
- smoke test / 单测结果
- memory note

## Execution Log

- planned approach: 先冻结边界，再打主路径
- current progress: 已完成 bootstrap 结构草稿
- blocker if any: runtime 责任仍偏空
- next action: 补 runtime 扩展说明和 smoke check

## Review Checkpoints

- RFC checkpoint: 边界划分已完成首轮评审
- implementation checkpoint: 启动主路径可运行
- validation checkpoint: smoke check 已通过
- documentation checkpoint: README 已更新启动说明
```

### 评审这种卡片时最该问什么

1. scope 是不是太大，导致一张卡里塞了多个月份的问题
2. acceptance checklist 是不是可验证，而不是口号
3. execution log 是不是能让别人接手，而不是只给作者自己看

## 示例 2：Execution Log 应该怎么写才有用

很多学习者写 execution log 时容易写成流水账：

- 今天看了这个
- 又改了那个
- 感觉差不多

这种写法几乎没有复用价值。

更好的写法应该围绕“当前状态”和“下一动作”。

### 不好的写法

```md
- planned approach: 先研究一下
- current progress: 写了一些东西
- blocker if any:
- next action: 继续写
```

### 更好的写法

```md
- planned approach: 先冻结 bootstrap / runtime 边界，再打最小启动路径
- current progress: bootstrap.New 已统一接 config、logger、runtime；serve 入口已能启动
- blocker if any: runtime 仍是空壳，后续依赖接入位置还需补说明
- next action: 在 review note 中明确 runtime 下一批扩展点，并补一条 smoke check
```

### 判断标准

好的 execution log 至少要让接手者能回答：

1. 已经做了什么
2. 卡在哪里
3. 下一步最小动作是什么

## 示例 3：Review Checkpoints 不是重复 checklist

很多人把 `Review Checkpoints` 写成 acceptance checklist 的重复版。

这会让它失去意义。

checkpoint 真正应该表达的是：

- 这个阶段有没有留下可 review 的东西
- 当前结论是不是已经站住

### 不好的写法

```md
- RFC checkpoint: 已写 RFC
- implementation checkpoint: 已实现
- validation checkpoint: 已测试
- documentation checkpoint: 已文档
```

### 更好的写法

```md
- RFC checkpoint: role / boundary / failure mode 已完成首轮评审，剩余争议是 response validation owner
- implementation checkpoint: 单一主路径已可演示，但 fallback 仍未纳入统一 gateway
- validation checkpoint: happy path 与 negative path 都有证据，timeout 路径还需补一条 case
- documentation checkpoint: operator note 已更新，但未补 error semantics 说明
```

### 判断标准

checkpoint 要能帮助 reviewer 快速定位：

1. 哪些阶段已经稳定
2. 哪些阶段仍然脆弱
3. 下一次 review 应该先看哪里

## 示例 4：Evidence Entry 应该怎么写

先对照真实材料：

- [Evidence Registry](../operations/evidence-registry.md)
- [Sample Evidence Entries](../governance/sample-evidence-entries.md)

### 最小结构

```md
- artifact id:
- linked issue or period:
- artifact type:
- why it matters:
- reviewer comment:
```

### 单月示例 1：设计证据

```md
- artifact id: EVD-M04-GATEWAY-RFC
- linked issue or period: ATLAS-M04-01 / Week 14
- artifact type: RFC
- why it matters: 冻结 tool gateway 的 request/result envelope、timeout、audit 责任边界
- reviewer comment: gateway control 面已经明确，但 response validation 还要补 owner
```

### 单月示例 2：验证证据

```md
- artifact id: EVD-M08-GATE-NEGATIVE
- linked issue or period: ATLAS-M08-03 / Week 33
- artifact type: validation suite
- why it matters: 证明 regression gate 在阈值退化时能正确阻断发布
- reviewer comment: gate fail path 成立，但误报处理流程还要补说明
```

### 单月示例 3：文档证据

```md
- artifact id: EVD-M12-OPS-GUIDE
- linked issue or period: Month 12
- artifact type: operator guide
- why it matters: 让接手者能运行平台级 drill、理解 rollback 和排障路径
- reviewer comment: 平台 story 已完整，但还要再压缩成更强的答辩证据组
```

### 写 evidence entry 时最容易犯的错

1. `why it matters` 写成“我做了什么”，而不是“为什么它重要”
2. `reviewer comment` 只写鼓励，没有指出下一步风险
3. linked issue or period 太宽，导致以后追不回上下文

## 示例 5：Month Review Packet 里的一句话应该怎么落证据

先对照真实文件：

- [Month 03 Review Packet](../reviews/archive/month-03-review.md)
- [Month 09 Review Packet](../reviews/archive/month-09-review.md)

### 一个不够好的句子

```md
- strongest result: Agent system works better now
```

问题在于：

- 没有说明“更好”是什么意思
- 没有指向任何 artifact

### 一个更好的句子

```md
- strongest result: 结构化输出与最小 eval baseline 已形成闭环，证据可回溯到 EVD-M03-RUNBOOK、Week 12 Report 和 Month 03 Scorecard
```

这类写法更强，因为它同时提供了：

1. 结论
2. 范围
3. 证据入口

## 一份像样的 Issue / Evidence 组合最少要满足什么

如果你只想记住最低标准，就记住下面这组检查：

1. issue card 能说明问题、范围、验收和证据期待
2. execution log 能说明当前状态和下一动作
3. review checkpoints 能说明哪里稳、哪里不稳
4. evidence entry 能说明为什么这个 artifact 值得被保留

## 最后一句

真正成熟的学习，不是“我写了很多文档”。

而是：

你写下的 card、review、evidence 和 memory，能让未来的你或别人，在没有你口头解释的前提下，仍然重建出这次工程判断。

如果你下一步要把这些材料继续压成月度 / 季度判断，再看：

- [Scorecard 与 Review Packet 示例](scorecard-and-review-examples.md)
