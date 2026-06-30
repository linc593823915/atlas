# Week 27：持久化工作流 RFC

最后更新：2026-06-30

## 本周目标

这一周开始进入真正的长流程复杂度。重点是先把工作流状态、checkpoint 与恢复逻辑定义清楚。

## 本周学完后你应该会什么

你要先回答“长流程状态属于谁”，而不是先写一堆流程代码。

## 本周学习重点

- workflow state
- checkpoint 边界
- pause/resume 语义

## 本周 Cookbook 做法

### Recipe 1：写 Durable Workflow RFC

### Recipe 2：画状态图

### Recipe 3：定义 checkpoint contract

## 本周交付标准

- Durable Workflow RFC
- 状态图
- checkpoint 设计草稿

## 本周能力焦点

- 有状态系统（`Stateful Systems`）
- 恢复设计（`Recovery Design`）
- 运行韧性（`Operational Resilience`）

## 本周知识清单

- durable workflow 先是 state model 问题：run、step、checkpoint 和 terminal outcome 要先被命名。
- checkpoint boundary 决定哪些东西必须在进程丢失后仍然存在，哪些东西可以重算。
- pause/resume 语义要有确定性，否则 resume 以后就变成了另一条流程。
- memory boundary 要回答哪些事实属于 workflow state，哪些事实属于外部长期记忆。

## 本周 Atlas 观察路径

- `[Week 27 英文原文](../weeks/week-27.md)`：对照原始英文任务拆分，确认这一周的 deliverables、source reading 和 daily links。
- `[Atlas 主线 Issue（Month 07）](../atlas/issues/month-07/atlas-m07-01.md)`：把本周主题放回 Atlas 主线问题卡，观察范围、验收项和剩余风险。
- `[Week 27 周报](../reports/weekly/week-27-report.md)`：看本周结果、证据和后续动作是如何被记录的。
- `[docs/reports/snapshots/month-07-snapshot.md](../reports/snapshots/month-07-snapshot.md)`：观察长流程能力在月度快照里如何被描述。

## 本周 Atlas 落地检查

- 最小落地动作：围绕“写 Durable Workflow RFC”推进，让至少一项本周交付标准在仓库里可见。
- 必留证据：workflow RFC、state model 和 checkpoint boundary。
- 最大风险：把长流程当成一个会 while 循环的普通请求，只在出问题时临时持久化。
- 审查不变量：一次 resume 必须能够从保存状态和确定性代码中被完整解释。
- 本周最先要盯住的交付：Durable Workflow RFC

## 本周最小演示场景

- 背景：这一周你需要把“持久化工作流 RFC”从概念推进成可被别人看懂、验证和质疑的工程结果。
- 你至少要能演示：
  - Durable Workflow RFC、状态图、checkpoint 设计草稿
  - “workflow state”为什么会直接影响 Atlas 的主路径，而不只是局部实现
  - 风险“把长流程当成一个会 while 循环的普通请求，只在出问题时临时持久化。”是如何被控制、暴露或留证的
- reviewer 大概率会追问：
  - 为什么这周先做“写 Durable Workflow RFC”，而不是先做别的？
  - 不变量“一次 resume 必须能够从保存状态和确定性代码中被完整解释。”如果被破坏，会先造成什么后果？
  - 你打算拿什么证据证明这周成果成立？提示：workflow RFC、state model 和 checkpoint boundary。

## 配套案例

- 本周最接近的案例：读 [专题案例包](casebook.md) 中的“案例 7：长流程跑得很长，但根本恢复不了”。
- 本周最该对照的失败模式：读 [失败场景图谱](failure-patterns.md) 中的“模式二：失败语义缺失型故障”。
- 推荐顺序：先读本周知识清单，再读案例，再回本周 Atlas 观察路径和周报，检查自己能不能解释坏信号与最小证据。

## 本周工作顺序

- [`RFC`](artifact-playbook.md#artifact-rfc)：先把“写 Durable Workflow RFC”对应的问题范围、目标、约束和非目标写清楚。
- [`Interface`](artifact-playbook.md#artifact-interface)：围绕“workflow state”定义输入、输出、owner、错误语义和边界。
- [`Implementation`](artifact-playbook.md#artifact-implementation)：把最小主路径打通，至少让“Durable Workflow RFC、状态图、checkpoint 设计草稿”里的一项开始可验证。
- [`Unit Test`](artifact-playbook.md#artifact-unit-test)：优先覆盖最容易击中风险“把长流程当成一个会 while 循环的普通请求，只在出问题时临时持久化。”的失败路径、边界条件或 contract。
- [`Review`](artifact-playbook.md#artifact-review)：站在 reviewer 视角检查不变量“一次 resume 必须能够从保存状态和确定性代码中被完整解释。”是否真的被系统保护，而不是只写在文档里。
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

1. 如果要把“持久化工作流 RFC”做成 Atlas 可交付能力，你会如何设计主路径和边界？
2. 本周最值得优先验证的失败模式是什么？你会如何收集证据？
3. 这周的 RFC、实现、验证或复盘里，哪一项最能体现你的工程判断？为什么？
4. 如果时间被砍掉一半，你会保留什么，推迟什么？
5. 本周产物如何自然承接到下周，而不是变成一次性文档或一次性代码？

## 每日入口

- [Day 183](days-semester-3/day-183.md): 理解问题和边界
- [Day 184](days-semester-3/day-184.md): 边界与契约设计
- [Day 185](days-semester-3/day-185.md): 最小主路径实现
- [Day 186](days-semester-3/day-186.md): 补全真实可用路径
- [Day 187](days-semester-3/day-187.md): 验证与实验
- [Day 188](days-semester-3/day-188.md): 文档与评审
- [Day 189](days-semester-3/day-189.md): 复盘与 memory

## 配套阅读

- [Week 27](../weeks/week-27.md)
