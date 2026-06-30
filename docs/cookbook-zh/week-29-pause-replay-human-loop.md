# Week 29：Pause、Replay 与人工介入

最后更新：2026-06-30

## 本周目标

本周进入工作流最难的运行时问题：什么时候停、怎么回放、人工怎么接入。

## 本周学完后你应该会什么

你要把“人工介入”设计成系统能力，而不是例外处理。

## 本周学习重点

- pause/resume
- replay
- human intervention

## 本周 Cookbook 做法

### Recipe 1：补 pause/resume 行为

### Recipe 2：补 replay 方案

### Recipe 3：定义人工介入点

## 本周交付标准

- 可暂停恢复流程
- replay 说明
- 人工介入路径

## 本周能力焦点

- 有状态系统（`Stateful Systems`）
- 恢复设计（`Recovery Design`）
- 运行韧性（`Operational Resilience`）

## 本周知识清单

- pause 是显式状态迁移，不是把 worker 杀掉以后祈祷下次还能接着跑。
- replay 要说清哪些步骤重新执行，哪些步骤只从证据回放，哪些历史绝不修改。
- human intervention point 需要 schema、audit trail 和重新并回工作流的语义。
- 人工修改过的状态如果没有 provenance，后面的自动化就无法诚实地继续运行。

## 本周 Atlas 观察路径

- `[Week 29 英文原文](../weeks/week-29.md)`：对照原始英文任务拆分，确认这一周的 deliverables、source reading 和 daily links。
- `[Atlas 主线 Issue（Month 07）](../atlas/issues/month-07/atlas-m07-03.md)`：把本周主题放回 Atlas 主线问题卡，观察范围、验收项和剩余风险。
- `[Week 29 周报](../reports/weekly/week-29-report.md)`：看本周结果、证据和后续动作是如何被记录的。
- `[docs/governance/sample-evidence-entries.md](../governance/sample-evidence-entries.md)`：对照人工介入和恢复流程该留下哪些证据。

## 本周 Atlas 落地检查

- 最小落地动作：围绕“补 pause/resume 行为”推进，让至少一项本周交付标准在仓库里可见。
- 必留证据：pause/resume design、replay 规则和 human-loop contract。
- 最大风险：把人工备注混进状态里，却没有任何来源、审批或合并规则。
- 审查不变量：任何人工介入都必须让 workflow 保持合法且可审计。
- 本周最先要盯住的交付：可暂停恢复流程

## 本周最小演示场景

- 背景：这一周你需要把“Pause、Replay 与人工介入”从概念推进成可被别人看懂、验证和质疑的工程结果。
- 你至少要能演示：
  - 可暂停恢复流程、replay 说明、人工介入路径
  - “pause/resume”为什么会直接影响 Atlas 的主路径，而不只是局部实现
  - 风险“把人工备注混进状态里，却没有任何来源、审批或合并规则。”是如何被控制、暴露或留证的
- reviewer 大概率会追问：
  - 为什么这周先做“补 pause/resume 行为”，而不是先做别的？
  - 不变量“任何人工介入都必须让 workflow 保持合法且可审计。”如果被破坏，会先造成什么后果？
  - 你打算拿什么证据证明这周成果成立？提示：pause/resume design、replay 规则和 human-loop contract。

## 配套案例

- 本周最接近的案例：读 [专题案例包](casebook.md) 中的“案例 7：长流程跑得很长，但根本恢复不了”。
- 本周最该对照的失败模式：读 [失败场景图谱](failure-patterns.md) 中的“模式二：失败语义缺失型故障”。
- 推荐顺序：先读本周知识清单，再读案例，再回本周 Atlas 观察路径和周报，检查自己能不能解释坏信号与最小证据。

## 本周工作顺序

- [`RFC`](artifact-playbook.md#artifact-rfc)：先把“补 pause/resume 行为”对应的问题范围、目标、约束和非目标写清楚。
- [`Interface`](artifact-playbook.md#artifact-interface)：围绕“pause/resume”定义输入、输出、owner、错误语义和边界。
- [`Implementation`](artifact-playbook.md#artifact-implementation)：把最小主路径打通，至少让“可暂停恢复流程、replay 说明、人工介入路径”里的一项开始可验证。
- [`Unit Test`](artifact-playbook.md#artifact-unit-test)：优先覆盖最容易击中风险“把人工备注混进状态里，却没有任何来源、审批或合并规则。”的失败路径、边界条件或 contract。
- [`Review`](artifact-playbook.md#artifact-review)：站在 reviewer 视角检查不变量“任何人工介入都必须让 workflow 保持合法且可审计。”是否真的被系统保护，而不是只写在文档里。
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

1. 如果要把“Pause、Replay 与人工介入”做成 Atlas 可交付能力，你会如何设计主路径和边界？
2. 本周最值得优先验证的失败模式是什么？你会如何收集证据？
3. 这周的 RFC、实现、验证或复盘里，哪一项最能体现你的工程判断？为什么？
4. 如果时间被砍掉一半，你会保留什么，推迟什么？
5. 本周产物如何自然承接到下周，而不是变成一次性文档或一次性代码？

## 每日入口

- [Day 197](days-semester-3/day-197.md): 理解问题和边界
- [Day 198](days-semester-3/day-198.md): 边界与契约设计
- [Day 199](days-semester-3/day-199.md): 最小主路径实现
- [Day 200](days-semester-3/day-200.md): 补全真实可用路径
- [Day 201](days-semester-3/day-201.md): 验证与实验
- [Day 202](days-semester-3/day-202.md): 文档与评审
- [Day 203](days-semester-3/day-203.md): 复盘与 memory

## 配套阅读

- [Week 29](../weeks/week-29.md)
