# Week 23：Agent 定义与共享上下文

最后更新：2026-06-30

## 本周目标

本周开始把角色设计落到真实运行时结构。

## 本周学完后你应该会什么

重点是 agent definitions 和 shared context。

## 本周学习重点

- Agent definition
- shared context
- trace id

## 本周 Cookbook 做法

### Recipe 1：定义主 Agent / specialist

### Recipe 2：建立 shared context

### Recipe 3：准备 trace path

## 本周交付标准

- 可运行的 agent definitions
- shared context 设计
- trace path 草图

## 本周能力焦点

- 编排能力（`Orchestration`）
- 以接口体现技术领导力（`Leadership Through Interfaces`）
- 审批纪律（`Approval Discipline`）

## 本周知识清单

- agent definition 至少要描述 role、输入、允许的工具、输出 contract 和 escalation rule。
- shared context 需要筛选，不应把所有历史和状态一次性灌给每个 agent。
- trace id、thread id 或 span id 是多 agent run 事后可重建的最小前提。
- agent specialization 只有在约束真实存在时才有意义，否则只是换个 prompt 名字。

## 本周 Atlas 观察路径

- `[Week 23 英文原文](../weeks/week-23.md)`：对照原始英文任务拆分，确认这一周的 deliverables、source reading 和 daily links。
- `[Atlas 主线 Issue（Month 06）](../atlas/issues/month-06/atlas-m06-02.md)`：把本周主题放回 Atlas 主线问题卡，观察范围、验收项和剩余风险。
- `[Week 23 周报](../reports/weekly/week-23-report.md)`：看本周结果、证据和后续动作是如何被记录的。
- `[docs/governance/sample-evidence-entries.md](../governance/sample-evidence-entries.md)`：观察 handoff、approval、trace 证据应该怎么写。

## 本周 Atlas 落地检查

- 最小落地动作：围绕“定义主 Agent / specialist”推进，让至少一项本周交付标准在仓库里可见。
- 必留证据：agent definition、shared context model 和 trace propagation。
- 最大风险：所有 agent 共享一个巨大的 mutable context blob，后面没人知道信息从哪来。
- 审查不变量：每个 agent 都应该知道自己接收了哪些上下文，以及这些上下文为什么会到它这里。
- 本周最先要盯住的交付：可运行的 agent definitions

## 本周最小演示场景

- 背景：这一周你需要把“Agent 定义与共享上下文”从概念推进成可被别人看懂、验证和质疑的工程结果。
- 你至少要能演示：
  - 可运行的 agent definitions、shared context 设计、trace path 草图
  - “Agent definition”为什么会直接影响 Atlas 的主路径，而不只是局部实现
  - 风险“所有 agent 共享一个巨大的 mutable context blob，后面没人知道信息从哪来。”是如何被控制、暴露或留证的
- reviewer 大概率会追问：
  - 为什么这周先做“定义主 Agent / specialist”，而不是先做别的？
  - 不变量“每个 agent 都应该知道自己接收了哪些上下文，以及这些上下文为什么会到它这里。”如果被破坏，会先造成什么后果？
  - 你打算拿什么证据证明这周成果成立？提示：agent definition、shared context model 和 trace propagation。

## 配套案例

- 本周最接近的案例：读 [专题案例包](casebook.md) 中的“案例 6：多 Agent 看起来更强，实际上没人知道谁负责”。
- 本周最该对照的失败模式：读 [失败场景图谱](failure-patterns.md) 中的“模式四：治理失守型故障”。
- 推荐顺序：先读本周知识清单，再读案例，再回本周 Atlas 观察路径和周报，检查自己能不能解释坏信号与最小证据。

## 本周工作顺序

- [`RFC`](artifact-playbook.md#artifact-rfc)：先把“定义主 Agent / specialist”对应的问题范围、目标、约束和非目标写清楚。
- [`Interface`](artifact-playbook.md#artifact-interface)：围绕“Agent definition”定义输入、输出、owner、错误语义和边界。
- [`Implementation`](artifact-playbook.md#artifact-implementation)：把最小主路径打通，至少让“可运行的 agent definitions、shared context 设计、trace path 草图”里的一项开始可验证。
- [`Unit Test`](artifact-playbook.md#artifact-unit-test)：优先覆盖最容易击中风险“所有 agent 共享一个巨大的 mutable context blob，后面没人知道信息从哪来。”的失败路径、边界条件或 contract。
- [`Review`](artifact-playbook.md#artifact-review)：站在 reviewer 视角检查不变量“每个 agent 都应该知道自己接收了哪些上下文，以及这些上下文为什么会到它这里。”是否真的被系统保护，而不是只写在文档里。
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

1. 如果要把“Agent 定义与共享上下文”做成 Atlas 可交付能力，你会如何设计主路径和边界？
2. 本周最值得优先验证的失败模式是什么？你会如何收集证据？
3. 这周的 RFC、实现、验证或复盘里，哪一项最能体现你的工程判断？为什么？
4. 如果时间被砍掉一半，你会保留什么，推迟什么？
5. 本周产物如何自然承接到下周，而不是变成一次性文档或一次性代码？

## 每日入口

- [Day 155](days-semester-2/day-155.md): 理解问题和边界
- [Day 156](days-semester-2/day-156.md): 边界与契约设计
- [Day 157](days-semester-2/day-157.md): 最小主路径实现
- [Day 158](days-semester-2/day-158.md): 补全真实可用路径
- [Day 159](days-semester-2/day-159.md): 验证与实验
- [Day 160](days-semester-2/day-160.md): 文档与评审
- [Day 161](days-semester-2/day-161.md): 复盘与 memory

## 配套阅读

- [Week 23](../weeks/week-23.md)
