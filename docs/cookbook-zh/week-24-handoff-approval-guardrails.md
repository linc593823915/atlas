# Week 24：Handoff、Approval 与 Guardrail

最后更新：2026-06-30

## 本周目标

这一周进入多 Agent 最核心也最容易失控的部分。

## 本周学完后你应该会什么

你要让 handoff 和 approval 变成显式系统行为，而不是隐式 Prompt 行为。

## 本周学习重点

- handoff 触发条件
- approval path
- guardrail 约束

## 本周 Cookbook 做法

### Recipe 1：加入 handoff 逻辑

### Recipe 2：加入 approval path

### Recipe 3：加入 guardrail 规则

## 本周交付标准

- 可解释的 handoff 流程
- approval 示例
- guardrail 规则初稿

## 本周能力焦点

- 编排能力（`Orchestration`）
- 以接口体现技术领导力（`Leadership Through Interfaces`）
- 审批纪律（`Approval Discipline`）

## 本周知识清单

- handoff 触发条件必须显式定义，而不是靠“感觉另一个 agent 也许更合适”。
- approval path 要讲清 requester、approver、scope、timeout 和批准后的审计证据。
- guardrail 是 runtime 规则：哪些动作禁止、哪些动作需确认、哪些输出必须收敛。
- 当策略信心不足时，系统应该可预期地降级到人工 review。

## 本周 Atlas 观察路径

- `[Week 24 英文原文](../weeks/week-24.md)`：对照原始英文任务拆分，确认这一周的 deliverables、source reading 和 daily links。
- `[Atlas 主线 Issue（Month 06）](../atlas/issues/month-06/atlas-m06-03.md)`：把本周主题放回 Atlas 主线问题卡，观察范围、验收项和剩余风险。
- `[Week 24 周报](../reports/weekly/week-24-report.md)`：看本周结果、证据和后续动作是如何被记录的。
- `[docs/governance/committee-decision-log.md](../governance/committee-decision-log.md)`：把 approval path 放回正式决策日志里理解。

## 本周 Atlas 落地检查

- 最小落地动作：围绕“加入 handoff 逻辑”推进，让至少一项本周交付标准在仓库里可见。
- 必留证据：handoff triggers、approval path 和 guardrail 规则。
- 最大风险：把高风险动作偷偷塞进 agent autonomy，既没有审批点也没有人工接管点。
- 审查不变量：敏感动作在执行前必须可以被阻断，而不是事后补日志。
- 本周最先要盯住的交付：可解释的 handoff 流程

## 本周最小演示场景

- 背景：这一周你需要把“Handoff、Approval 与 Guardrail”从概念推进成可被别人看懂、验证和质疑的工程结果。
- 你至少要能演示：
  - 可解释的 handoff 流程、approval 示例、guardrail 规则初稿
  - “handoff 触发条件”为什么会直接影响 Atlas 的主路径，而不只是局部实现
  - 风险“把高风险动作偷偷塞进 agent autonomy，既没有审批点也没有人工接管点。”是如何被控制、暴露或留证的
- reviewer 大概率会追问：
  - 为什么这周先做“加入 handoff 逻辑”，而不是先做别的？
  - 不变量“敏感动作在执行前必须可以被阻断，而不是事后补日志。”如果被破坏，会先造成什么后果？
  - 你打算拿什么证据证明这周成果成立？提示：handoff triggers、approval path 和 guardrail 规则。

## 配套案例

- 本周最接近的案例：读 [专题案例包](casebook.md) 中的“案例 6：多 Agent 看起来更强，实际上没人知道谁负责”。
- 本周最该对照的失败模式：读 [失败场景图谱](failure-patterns.md) 中的“模式四：治理失守型故障”。
- 推荐顺序：先读本周知识清单，再读案例，再回本周 Atlas 观察路径和周报，检查自己能不能解释坏信号与最小证据。

## 本周工作顺序

- [`RFC`](artifact-playbook.md#artifact-rfc)：先把“加入 handoff 逻辑”对应的问题范围、目标、约束和非目标写清楚。
- [`Interface`](artifact-playbook.md#artifact-interface)：围绕“handoff 触发条件”定义输入、输出、owner、错误语义和边界。
- [`Implementation`](artifact-playbook.md#artifact-implementation)：把最小主路径打通，至少让“可解释的 handoff 流程、approval 示例、guardrail 规则初稿”里的一项开始可验证。
- [`Unit Test`](artifact-playbook.md#artifact-unit-test)：优先覆盖最容易击中风险“把高风险动作偷偷塞进 agent autonomy，既没有审批点也没有人工接管点。”的失败路径、边界条件或 contract。
- [`Review`](artifact-playbook.md#artifact-review)：站在 reviewer 视角检查不变量“敏感动作在执行前必须可以被阻断，而不是事后补日志。”是否真的被系统保护，而不是只写在文档里。
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

1. 如果要把“Handoff、Approval 与 Guardrail”做成 Atlas 可交付能力，你会如何设计主路径和边界？
2. 本周最值得优先验证的失败模式是什么？你会如何收集证据？
3. 这周的 RFC、实现、验证或复盘里，哪一项最能体现你的工程判断？为什么？
4. 如果时间被砍掉一半，你会保留什么，推迟什么？
5. 本周产物如何自然承接到下周，而不是变成一次性文档或一次性代码？

## 每日入口

- [Day 162](days-semester-2/day-162.md): 理解问题和边界
- [Day 163](days-semester-2/day-163.md): 边界与契约设计
- [Day 164](days-semester-2/day-164.md): 最小主路径实现
- [Day 165](days-semester-2/day-165.md): 补全真实可用路径
- [Day 166](days-semester-2/day-166.md): 验证与实验
- [Day 167](days-semester-2/day-167.md): 文档与评审
- [Day 168](days-semester-2/day-168.md): 复盘与 memory

## 配套阅读

- [Week 24](../weeks/week-24.md)
