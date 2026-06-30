# Week 15：Schema 校验与 Registry 模型

最后更新：2026-06-30

## 本周目标

本周要把 Tool Gateway 从概念推进到可执行结构，重点是 registry 和 schema validation。

## 本周学完后你应该会什么

你要开始把“工具调用”变成“可验证接口”。

## 本周学习重点

- Registry 模型
- Schema 校验
- 输入输出契约

## 本周 Cookbook 做法

### Recipe 1：定义 tool registration 流程

### Recipe 2：接入 schema validation

### Recipe 3：整理错误模型

## 本周交付标准

- 可注册的 Tool 结构
- schema 校验路径
- 错误模型草稿

## 本周能力焦点

- 接口设计（`Interface Design`）
- 代码评审质量（`Code Review Quality`）
- 小步变更纪律（`Small Change Discipline`）

## 本周知识清单

- Registry 模型决定谁拥有 tool schema、version 和 discoverability metadata。
- validation 既要发生在执行前，也要考虑执行后返回结构是否仍满足 contract。
- 输入输出 contract 需要版本规则，否则 tool 演进会悄悄打碎调用方。
- registry 还要让 unsupported、deprecated 和 experimental 状态显式可见。

## 本周 Atlas 观察路径

- `[Week 15 英文原文](../weeks/week-15.md)`：对照原始英文任务拆分，确认这一周的 deliverables、source reading 和 daily links。
- `[Atlas 主线 Issue（Month 04）](../atlas/issues/month-04/atlas-m04-02.md)`：把本周主题放回 Atlas 主线问题卡，观察范围、验收项和剩余风险。
- `[Week 15 周报](../reports/weekly/week-15-report.md)`：看本周结果、证据和后续动作是如何被记录的。
- `[docs/atlas/issues/month-04/README.md](../atlas/issues/month-04/README.md)`：对照 Tool Gateway 相关 issue 集合。

## 本周 Atlas 落地检查

- 最小落地动作：围绕“定义 tool registration 流程”推进，让至少一项本周交付标准在仓库里可见。
- 必留证据：registry model、schema validation 和 I/O contract 示例。
- 最大风险：schema 分散在调用方和实现方两边，时间一长一定漂移。
- 审查不变量：没有 registry-backed contract 的 tool 不应该进入执行路径。
- 本周最先要盯住的交付：可注册的 Tool 结构

## 本周最小演示场景

- 背景：这一周你需要把“Schema 校验与 Registry 模型”从概念推进成可被别人看懂、验证和质疑的工程结果。
- 你至少要能演示：
  - 可注册的 Tool 结构、schema 校验路径、错误模型草稿
  - “Registry 模型”为什么会直接影响 Atlas 的主路径，而不只是局部实现
  - 风险“schema 分散在调用方和实现方两边，时间一长一定漂移。”是如何被控制、暴露或留证的
- reviewer 大概率会追问：
  - 为什么这周先做“定义 tool registration 流程”，而不是先做别的？
  - 不变量“没有 registry-backed contract 的 tool 不应该进入执行路径。”如果被破坏，会先造成什么后果？
  - 你打算拿什么证据证明这周成果成立？提示：registry model、schema validation 和 I/O contract 示例。

## 配套案例

- 本周最接近的案例：读 [专题案例包](casebook.md) 中的“案例 4：工具越来越多，但系统越来越难治理”。
- 本周最该对照的失败模式：读 [失败场景图谱](failure-patterns.md) 中的“模式一：边界缺失型故障”。
- 推荐顺序：先读本周知识清单，再读案例，再回本周 Atlas 观察路径和周报，检查自己能不能解释坏信号与最小证据。

## 本周工作顺序

- [`RFC`](artifact-playbook.md#artifact-rfc)：先把“定义 tool registration 流程”对应的问题范围、目标、约束和非目标写清楚。
- [`Interface`](artifact-playbook.md#artifact-interface)：围绕“Registry 模型”定义输入、输出、owner、错误语义和边界。
- [`Implementation`](artifact-playbook.md#artifact-implementation)：把最小主路径打通，至少让“可注册的 Tool 结构、schema 校验路径、错误模型草稿”里的一项开始可验证。
- [`Unit Test`](artifact-playbook.md#artifact-unit-test)：优先覆盖最容易击中风险“schema 分散在调用方和实现方两边，时间一长一定漂移。”的失败路径、边界条件或 contract。
- [`Review`](artifact-playbook.md#artifact-review)：站在 reviewer 视角检查不变量“没有 registry-backed contract 的 tool 不应该进入执行路径。”是否真的被系统保护，而不是只写在文档里。
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

1. 如果要把“Schema 校验与 Registry 模型”做成 Atlas 可交付能力，你会如何设计主路径和边界？
2. 本周最值得优先验证的失败模式是什么？你会如何收集证据？
3. 这周的 RFC、实现、验证或复盘里，哪一项最能体现你的工程判断？为什么？
4. 如果时间被砍掉一半，你会保留什么，推迟什么？
5. 本周产物如何自然承接到下周，而不是变成一次性文档或一次性代码？

## 每日入口

- [Day 099](days-semester-2/day-099.md): 理解问题和边界
- [Day 100](days-semester-2/day-100.md): 边界与契约设计
- [Day 101](days-semester-2/day-101.md): 最小主路径实现
- [Day 102](days-semester-2/day-102.md): 补全真实可用路径
- [Day 103](days-semester-2/day-103.md): 验证与实验
- [Day 104](days-semester-2/day-104.md): 文档与评审
- [Day 105](days-semester-2/day-105.md): 复盘与 memory

## 配套阅读

- [Week 15](../weeks/week-15.md)
