# Week 49：端到端测试与失败演练

最后更新：2026-06-30

## 本周目标

本周验证整个平台不是“结构上完成”，而是“运行上成立”。

## 本周学完后你应该会什么

重点是做 end-to-end 和 failure drills。

## 本周学习重点

- e2e tests
- failure drills
- recovery readiness

## 本周 Cookbook 做法

### Recipe 1：补端到端测试

### Recipe 2：补失败演练

### Recipe 3：记录恢复逻辑

## 本周交付标准

- E2E 测试集
- drill 记录
- 恢复说明

## 本周能力焦点

- 平台整合（`Platform Integration`）
- 架构判断（`Architecture Judgment`）
- 技术领导力（`Technical Leadership`）

## 本周知识清单

- 端到端演练要覆盖真实跨子系统旅程，而不是把模块级 happy path 拼成幻觉。
- failure drill 至少要覆盖依赖丢失、policy deny、timeout、stale config 和 recovery step。
- recovery readiness 证明的是操作行为可重复，而不是大家都相信“应该能恢复”。
- 环境一致性很重要，toy 环境里的 drill 很容易漏掉部署现实。

## 本周 Atlas 观察路径

- `[Week 49 英文原文](../weeks/week-49.md)`：对照原始英文任务拆分，确认这一周的 deliverables、source reading 和 daily links。
- `[Atlas 主线 Issue（Month 12）](../atlas/issues/month-12/atlas-m12-03.md)`：把本周主题放回 Atlas 主线问题卡，观察范围、验收项和剩余风险。
- `[Week 49 周报](../reports/weekly/week-49-report.md)`：看本周结果、证据和后续动作是如何被记录的。
- `[docs/governance/annual-review-memo.md](../governance/annual-review-memo.md)`：看年终复盘应该如何用证据讲系统故事。

## 本周 Atlas 落地检查

- 最小落地动作：围绕“补端到端测试”推进，让至少一项本周交付标准在仓库里可见。
- 必留证据：e2e 场景、failure drills 和 recovery checklist。
- 最大风险：模块级检查全是绿色，但整个平台一串起来就掉链子。
- 审查不变量：每条关键用户旅程都必须有清楚的 degraded path 或 recovery path。
- 本周最先要盯住的交付：E2E 测试集

## 本周最小演示场景

- 背景：这一周你需要把“端到端测试与失败演练”从概念推进成可被别人看懂、验证和质疑的工程结果。
- 你至少要能演示：
  - E2E 测试集、drill 记录、恢复说明
  - “e2e tests”为什么会直接影响 Atlas 的主路径，而不只是局部实现
  - 风险“模块级检查全是绿色，但整个平台一串起来就掉链子。”是如何被控制、暴露或留证的
- reviewer 大概率会追问：
  - 为什么这周先做“补端到端测试”，而不是先做别的？
  - 不变量“每条关键用户旅程都必须有清楚的 degraded path 或 recovery path。”如果被破坏，会先造成什么后果？
  - 你打算拿什么证据证明这周成果成立？提示：e2e 场景、failure drills 和 recovery checklist。

## 配套案例

- 本周最接近的案例：读 [专题案例包](casebook.md) 中的“案例 12：功能都在，为什么还是说不清这是个平台”。
- 本周最该对照的失败模式：读 [失败场景图谱](failure-patterns.md) 中的“模式五：平台错觉型故障”。
- 推荐顺序：先读本周知识清单，再读案例，再回本周 Atlas 观察路径和周报，检查自己能不能解释坏信号与最小证据。

## 本周工作顺序

- [`RFC`](artifact-playbook.md#artifact-rfc)：先把“补端到端测试”对应的问题范围、目标、约束和非目标写清楚。
- [`Interface`](artifact-playbook.md#artifact-interface)：围绕“e2e tests”定义输入、输出、owner、错误语义和边界。
- [`Implementation`](artifact-playbook.md#artifact-implementation)：把最小主路径打通，至少让“E2E 测试集、drill 记录、恢复说明”里的一项开始可验证。
- [`Unit Test`](artifact-playbook.md#artifact-unit-test)：优先覆盖最容易击中风险“模块级检查全是绿色，但整个平台一串起来就掉链子。”的失败路径、边界条件或 contract。
- [`Review`](artifact-playbook.md#artifact-review)：站在 reviewer 视角检查不变量“每条关键用户旅程都必须有清楚的 degraded path 或 recovery path。”是否真的被系统保护，而不是只写在文档里。
- [`Benchmark（按需）`](artifact-playbook.md#artifact-benchmark)：本周不做性能 benchmark，但要保留端到端 drills 和恢复准备度证据。
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

1. 如果要把“端到端测试与失败演练”做成 Atlas 可交付能力，你会如何设计主路径和边界？
2. 本周最值得优先验证的失败模式是什么？你会如何收集证据？
3. 这周的 RFC、实现、验证或复盘里，哪一项最能体现你的工程判断？为什么？
4. 如果时间被砍掉一半，你会保留什么，推迟什么？
5. 本周产物如何自然承接到下周，而不是变成一次性文档或一次性代码？

## 每日入口

- [Day 337](days-semester-4/day-337.md): 理解问题和边界
- [Day 338](days-semester-4/day-338.md): 边界与契约设计
- [Day 339](days-semester-4/day-339.md): 最小主路径实现
- [Day 340](days-semester-4/day-340.md): 补全真实可用路径
- [Day 341](days-semester-4/day-341.md): 验证与实验
- [Day 342](days-semester-4/day-342.md): 文档与评审
- [Day 343](days-semester-4/day-343.md): 复盘与 memory

## 配套阅读

- [Week 49](../weeks/week-49.md)
