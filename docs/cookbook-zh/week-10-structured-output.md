# Week 10：结构化输出与第一个 Tool

最后更新：2026-06-30

## 本周目标

这一周要打通第一个真实 Tool，并让 Agent 输出可被验证。

## 本周学习重点

- structured output 的工程价值
- schema 验证
- 第一条 end-to-end tool path

## 本周 Cookbook 做法

### Recipe 1：接入第一个 Tool

### Recipe 2：解析结构化输出

### Recipe 3：让输出经过校验

## 本周交付标准

你至少要能演示：

- Agent 怎么调 Tool
- 输出怎么被解析
- 为什么这个输出是“系统可用”的

## 本周能力焦点

- 问题求解（`Problem Solving`）
- 面向用户的工具使用（`User-Facing Tool Use`）
- 度量基础（`Measurement Basics`）

## 本周知识清单

- structured output 是 LLM 文本和确定性代码之间的桥梁，没有 schema 就没有稳定接口。
- schema 校验不只要判断对错，还要定义 reject、repair、retry 或 fail-closed 的后续动作。
- 第一条 tool path 应该把输入、tool result 和最终输出结构串成一个完整闭环。
- tool 返回值需要统一规范化，否则后面的 routing、eval 和 audit 都会变脆。

## 本周 Atlas 观察路径

- `[Week 10 英文原文](../weeks/week-10.md)`：对照原始英文任务拆分，确认这一周的 deliverables、source reading 和 daily links。
- `[Atlas 主线 Issue（Month 03）](../atlas/issues/month-03/atlas-m03-02.md)`：把本周主题放回 Atlas 主线问题卡，观察范围、验收项和剩余风险。
- `[Week 10 周报](../reports/weekly/week-10-report.md)`：看本周结果、证据和后续动作是如何被记录的。
- `[docs/reports/monthly/month-03-scorecard.md](../reports/monthly/month-03-scorecard.md)`：观察单 Agent 能力如何被月度验收。

## 本周 Atlas 落地检查

- 最小落地动作：围绕“接入第一个 Tool”推进，让至少一项本周交付标准在仓库里可见。
- 必留证据：schema、validator path 和第一条 end-to-end tool call。
- 最大风险：把“差不多符合预期”的自然语言当成可靠数据直接使用。
- 审查不变量：任何进入后续逻辑的结构化输出都不能绕过 schema 验证。
- 本周最先要盯住的交付：Agent 怎么调 Tool

## 本周最小演示场景

- 背景：这一周你需要把“结构化输出与第一个 Tool”从概念推进成可被别人看懂、验证和质疑的工程结果。
- 你至少要能演示：
  - Agent 怎么调 Tool、输出怎么被解析、为什么这个输出是“系统可用”的
  - “structured output 的工程价值”为什么会直接影响 Atlas 的主路径，而不只是局部实现
  - 风险“把“差不多符合预期”的自然语言当成可靠数据直接使用。”是如何被控制、暴露或留证的
- reviewer 大概率会追问：
  - 为什么这周先做“接入第一个 Tool”，而不是先做别的？
  - 不变量“任何进入后续逻辑的结构化输出都不能绕过 schema 验证。”如果被破坏，会先造成什么后果？
  - 你打算拿什么证据证明这周成果成立？提示：schema、validator path 和第一条 end-to-end tool call。

## 配套案例

- 本周最接近的案例：读 [专题案例包](casebook.md) 中的“案例 3：做出了会聊天的 Agent，但系统完全不可依赖”。
- 本周最该对照的失败模式：读 [失败场景图谱](failure-patterns.md) 中的“模式三：质量证据失真型故障”。
- 推荐顺序：先读本周知识清单，再读案例，再回本周 Atlas 观察路径和周报，检查自己能不能解释坏信号与最小证据。

## 本周工作顺序

- [`RFC`](artifact-playbook.md#artifact-rfc)：先把“接入第一个 Tool”对应的问题范围、目标、约束和非目标写清楚。
- [`Interface`](artifact-playbook.md#artifact-interface)：围绕“structured output 的工程价值”定义输入、输出、owner、错误语义和边界。
- [`Implementation`](artifact-playbook.md#artifact-implementation)：把最小主路径打通，至少让“Agent 怎么调 Tool、输出怎么被解析、为什么这个输出是“系统可用”的”里的一项开始可验证。
- [`Unit Test`](artifact-playbook.md#artifact-unit-test)：优先覆盖最容易击中风险“把“差不多符合预期”的自然语言当成可靠数据直接使用。”的失败路径、边界条件或 contract。
- [`Review`](artifact-playbook.md#artifact-review)：站在 reviewer 视角检查不变量“任何进入后续逻辑的结构化输出都不能绕过 schema 验证。”是否真的被系统保护，而不是只写在文档里。
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

1. 如果要把“结构化输出与第一个 Tool”做成 Atlas 可交付能力，你会如何设计主路径和边界？
2. 本周最值得优先验证的失败模式是什么？你会如何收集证据？
3. 这周的 RFC、实现、验证或复盘里，哪一项最能体现你的工程判断？为什么？
4. 如果时间被砍掉一半，你会保留什么，推迟什么？
5. 本周产物如何自然承接到下周，而不是变成一次性文档或一次性代码？

## 每日入口

- [Day 064](days-semester-1/day-064.md): 理解问题和边界
- [Day 065](days-semester-1/day-065.md): 边界与契约设计
- [Day 066](days-semester-1/day-066.md): 最小主路径实现
- [Day 067](days-semester-1/day-067.md): 补全真实可用路径
- [Day 068](days-semester-1/day-068.md): 验证与实验
- [Day 069](days-semester-1/day-069.md): 文档与评审
- [Day 070](days-semester-1/day-070.md): 复盘与 memory

## 配套阅读

- [Week 10](../weeks/week-10.md)
- [Week 10 Report](../reports/weekly/week-10-report.md)
