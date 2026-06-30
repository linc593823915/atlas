# Week 28：状态图、Checkpoint 与恢复模型

最后更新：2026-06-30

## 本周目标

本周把持久化工作流从设计推进到结构。重点是状态持久化与恢复模型。

## 本周学完后你应该会什么

你要开始让长流程“真的能恢复”，而不是逻辑上看起来能恢复。

## 本周学习重点

- graph state
- checkpoint store
- resume model

## 本周 Cookbook 做法

### Recipe 1：建立状态结构

### Recipe 2：接入 checkpoint

### Recipe 3：定义恢复路径

## 本周交付标准

- checkpoint 存储模型
- 状态恢复路径
- 最小可恢复流程

## 本周能力焦点

- 有状态系统（`Stateful Systems`）
- 恢复设计（`Recovery Design`）
- 运行韧性（`Operational Resilience`）

## 本周知识清单

- graph state 表达要能显式表示并行分支、依赖关系和 waiting state。
- checkpoint store 设计要处理原子性、版本和部分写入后的恢复问题。
- resume model 需要一个明确 cursor 或 decision point，告诉 runtime 应该从哪里继续。
- 部分副作用完成后的恢复策略必须先想清楚，不然 replay/continue 都会撒谎。

## 本周 Atlas 观察路径

- `[Week 28 英文原文](../weeks/week-28.md)`：对照原始英文任务拆分，确认这一周的 deliverables、source reading 和 daily links。
- `[Atlas 主线 Issue（Month 07）](../atlas/issues/month-07/atlas-m07-02.md)`：把本周主题放回 Atlas 主线问题卡，观察范围、验收项和剩余风险。
- `[Week 28 周报](../reports/weekly/week-28-report.md)`：看本周结果、证据和后续动作是如何被记录的。
- `[docs/reports/monthly/month-07-scorecard.md](../reports/monthly/month-07-scorecard.md)`：把 checkpoint / resume 结果放回评分口径里理解。

## 本周 Atlas 落地检查

- 最小落地动作：围绕“建立状态结构”推进，让至少一项本周交付标准在仓库里可见。
- 必留证据：graph model、checkpoint store 和 resume path。
- 最大风险：保存了 snapshot 却无法从中恢复出唯一的控制流决策。
- 审查不变量：加载一个 checkpoint 后，系统必须只能得到一个明确的下一步。
- 本周最先要盯住的交付：checkpoint 存储模型

## 本周最小演示场景

- 背景：这一周你需要把“状态图、Checkpoint 与恢复模型”从概念推进成可被别人看懂、验证和质疑的工程结果。
- 你至少要能演示：
  - checkpoint 存储模型、状态恢复路径、最小可恢复流程
  - “graph state”为什么会直接影响 Atlas 的主路径，而不只是局部实现
  - 风险“保存了 snapshot 却无法从中恢复出唯一的控制流决策。”是如何被控制、暴露或留证的
- reviewer 大概率会追问：
  - 为什么这周先做“建立状态结构”，而不是先做别的？
  - 不变量“加载一个 checkpoint 后，系统必须只能得到一个明确的下一步。”如果被破坏，会先造成什么后果？
  - 你打算拿什么证据证明这周成果成立？提示：graph model、checkpoint store 和 resume path。

## 配套案例

- 本周最接近的案例：读 [专题案例包](casebook.md) 中的“案例 7：长流程跑得很长，但根本恢复不了”。
- 本周最该对照的失败模式：读 [失败场景图谱](failure-patterns.md) 中的“模式二：失败语义缺失型故障”。
- 推荐顺序：先读本周知识清单，再读案例，再回本周 Atlas 观察路径和周报，检查自己能不能解释坏信号与最小证据。

## 本周工作顺序

- [`RFC`](artifact-playbook.md#artifact-rfc)：先把“建立状态结构”对应的问题范围、目标、约束和非目标写清楚。
- [`Interface`](artifact-playbook.md#artifact-interface)：围绕“graph state”定义输入、输出、owner、错误语义和边界。
- [`Implementation`](artifact-playbook.md#artifact-implementation)：把最小主路径打通，至少让“checkpoint 存储模型、状态恢复路径、最小可恢复流程”里的一项开始可验证。
- [`Unit Test`](artifact-playbook.md#artifact-unit-test)：优先覆盖最容易击中风险“保存了 snapshot 却无法从中恢复出唯一的控制流决策。”的失败路径、边界条件或 contract。
- [`Review`](artifact-playbook.md#artifact-review)：站在 reviewer 视角检查不变量“加载一个 checkpoint 后，系统必须只能得到一个明确的下一步。”是否真的被系统保护，而不是只写在文档里。
- [`Benchmark（按需）`](artifact-playbook.md#artifact-benchmark)：本周如果没有性能主题，不强做 benchmark；重点是留下能证明交付成立的测试、drill 或评估证据。
- [`Documentation`](artifact-playbook.md#artifact-documentation)：补运行说明、review note 或 operator 指引，让别人能复用本周结论。
- [`Memory`](artifact-playbook.md#artifact-memory)：沉淀一条可复用结论，说明为什么本周的边界、证据或失败判断值得保留。

## 本周评分标准

- 满分 `100` 分。
- 交付物完成度：25 分。
- 技术清晰度与取舍解释：20 分。
- 代码健康、测试或验证深度：15 分。
- 可靠性与边界条件思考：15 分。
- 文档与评审质量：10 分。
- 复盘、memory 和下一步质量：15 分。
- 及格线：`75+`，且本周计划产物全部存在。
- 优秀线：`90+`，并能证明本周工作真实改善了 Atlas，而不只是完成表面任务。

## 本周面试追问

1. 如果要把“状态图、Checkpoint 与恢复模型”做成 Atlas 可交付能力，你会如何设计主路径和边界？
2. 本周最值得优先验证的失败模式是什么？你会如何收集证据？
3. 这周的 RFC、实现、验证或复盘里，哪一项最能体现你的工程判断？为什么？
4. 如果时间被砍掉一半，你会保留什么，推迟什么？
5. 本周产物如何自然承接到下周，而不是变成一次性文档或一次性代码？

## 每日入口

- [Day 190](days-semester-3/day-190.md): 理解问题和边界
- [Day 191](days-semester-3/day-191.md): 边界与契约设计
- [Day 192](days-semester-3/day-192.md): 最小主路径实现
- [Day 193](days-semester-3/day-193.md): 补全真实可用路径
- [Day 194](days-semester-3/day-194.md): 验证与实验
- [Day 195](days-semester-3/day-195.md): 文档与评审
- [Day 196](days-semester-3/day-196.md): 复盘与 memory

## 配套阅读

- [Week 28](../weeks/week-28.md)
