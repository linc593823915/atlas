# Week 09：第一个 Agent RFC 与 Tool 清单

最后更新：2026-06-30

## 本周目标

这周第一次真正站到 Agent 系统设计门口。

你要先写清楚：

- 这个 Agent 做什么
- 它有哪些 Tool
- 输出契约是什么

## 本周学习重点

- Agent 角色定义
- Tool inventory
- Output contract

## 本周 Cookbook 做法

### Recipe 1：写 Agent RFC

### Recipe 2：列出 Tool 清单

### Recipe 3：定义输出契约

## 本周交付标准

你至少要完成：

- Agent RFC
- Tool 列表
- 输出契约草稿

## 本周能力焦点

- 问题求解（`Problem Solving`）
- 面向用户的工具使用（`User-Facing Tool Use`）
- 度量基础（`Measurement Basics`）

## 本周知识清单

- Agent role 要回答模型负责什么、runtime 负责什么、tool 负责什么，而不是只写一段 prompt。
- tool inventory 应该按外部能力组织，而不是把 repo 里的每个 helper 都变成 tool。
- output contract 让下游系统依赖的是结构和语义，而不是“模型大概会这么说”。
- 第一个 Agent RFC 应该同时描述 role、tool schema 和 runtime orchestration 的分工。

## 本周 Atlas 观察路径

- `[Week 09 英文原文](../weeks/week-09.md)`：对照原始英文任务拆分，确认这一周的 deliverables、source reading 和 daily links。
- `[Atlas 主线 Issue（Month 03）](../atlas/issues/month-03/atlas-m03-01.md)`：把本周主题放回 Atlas 主线问题卡，观察范围、验收项和剩余风险。
- `[Week 09 周报](../reports/weekly/week-09-report.md)`：看本周结果、证据和后续动作是如何被记录的。
- `[docs/atlas/atlas.md](../atlas/atlas.md)`：把第一个 Agent 放回 Atlas 整体叙事里理解。

## 本周 Atlas 落地检查

- 最小落地动作：围绕“写 Agent RFC”推进，让至少一项本周交付标准在仓库里可见。
- 必留证据：Agent RFC、tool inventory 和 output contract。
- 最大风险：做出一个能聊的 Demo，却没有任何可复用的 contract。
- 审查不变量：凡是要流入下游逻辑的模型输出，都必须挂在显式 contract 之下。
- 本周最先要盯住的交付：Agent RFC

## 本周最小演示场景

- 背景：这一周你需要把“第一个 Agent RFC 与 Tool 清单”从概念推进成可被别人看懂、验证和质疑的工程结果。
- 你至少要能演示：
  - Agent RFC、Tool 列表、输出契约草稿
  - “Agent 角色定义”为什么会直接影响 Atlas 的主路径，而不只是局部实现
  - 风险“做出一个能聊的 Demo，却没有任何可复用的 contract。”是如何被控制、暴露或留证的
- reviewer 大概率会追问：
  - 为什么这周先做“写 Agent RFC”，而不是先做别的？
  - 不变量“凡是要流入下游逻辑的模型输出，都必须挂在显式 contract 之下。”如果被破坏，会先造成什么后果？
  - 你打算拿什么证据证明这周成果成立？提示：Agent RFC、tool inventory 和 output contract。

## 配套案例

- 本周最接近的案例：读 [专题案例包](casebook.md) 中的“案例 3：做出了会聊天的 Agent，但系统完全不可依赖”。
- 本周最该对照的失败模式：读 [失败场景图谱](failure-patterns.md) 中的“模式三：质量证据失真型故障”。
- 推荐顺序：先读本周知识清单，再读案例，再回本周 Atlas 观察路径和周报，检查自己能不能解释坏信号与最小证据。

## 本周工作顺序

- [`RFC`](artifact-playbook.md#artifact-rfc)：先把“写 Agent RFC”对应的问题范围、目标、约束和非目标写清楚。
- [`Interface`](artifact-playbook.md#artifact-interface)：围绕“Agent 角色定义”定义输入、输出、owner、错误语义和边界。
- [`Implementation`](artifact-playbook.md#artifact-implementation)：把最小主路径打通，至少让“Agent RFC、Tool 列表、输出契约草稿”里的一项开始可验证。
- [`Unit Test`](artifact-playbook.md#artifact-unit-test)：优先覆盖最容易击中风险“做出一个能聊的 Demo，却没有任何可复用的 contract。”的失败路径、边界条件或 contract。
- [`Review`](artifact-playbook.md#artifact-review)：站在 reviewer 视角检查不变量“凡是要流入下游逻辑的模型输出，都必须挂在显式 contract 之下。”是否真的被系统保护，而不是只写在文档里。
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

1. 如果要把“第一个 Agent RFC 与 Tool 清单”做成 Atlas 可交付能力，你会如何设计主路径和边界？
2. 本周最值得优先验证的失败模式是什么？你会如何收集证据？
3. 这周的 RFC、实现、验证或复盘里，哪一项最能体现你的工程判断？为什么？
4. 如果时间被砍掉一半，你会保留什么，推迟什么？
5. 本周产物如何自然承接到下周，而不是变成一次性文档或一次性代码？

## 每日入口

- [Day 057](days-semester-1/day-057.md): 理解问题和边界
- [Day 058](days-semester-1/day-058.md): 边界与契约设计
- [Day 059](days-semester-1/day-059.md): 最小主路径实现
- [Day 060](days-semester-1/day-060.md): 补全真实可用路径
- [Day 061](days-semester-1/day-061.md): 验证与实验
- [Day 062](days-semester-1/day-062.md): 文档与评审
- [Day 063](days-semester-1/day-063.md): 复盘与 memory

## 配套阅读

- [Week 09](../weeks/week-09.md)
- [Month 03 Backlog](../atlas/backlog/month-03.md)
