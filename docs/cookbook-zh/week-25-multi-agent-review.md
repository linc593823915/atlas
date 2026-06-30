# Week 25：多 Agent 测试与复盘

最后更新：2026-06-30

## 本周目标

本周要把多 Agent 运行时收口成可评估系统。

## 本周学完后你应该会什么

重点是补测试、补文档、补复盘。

## 本周学习重点

- Orchestration tests
- fallback tests
- runtime docs

## 本周 Cookbook 做法

### Recipe 1：补多 Agent 测试

### Recipe 2：补运行时文档

### Recipe 3：做第六个月复盘

## 本周交付标准

- 多 Agent 测试集
- 运行时文档
- 第六个月复盘结论

## 本周能力焦点

- 编排能力（`Orchestration`）
- 以接口体现技术领导力（`Leadership Through Interfaces`）
- 审批纪律（`Approval Discipline`）

## 本周知识清单

- orchestration tests 要覆盖多步委托、上下文传播和结果聚合，而不是只测单个 agent。
- fallback tests 应该模拟 child agent 卡住、失败或返回不可用结果时的退化行为。
- runtime docs 需要说明如何检查一条 agent 链、如何理解审批流和回退路径。
- 月度复盘要证明多 Agent 复杂度换来了真实能力，而不是只换来更多 moving parts。

## 本周 Atlas 观察路径

- `[Week 25 英文原文](../weeks/week-25.md)`：对照原始英文任务拆分，确认这一周的 deliverables、source reading 和 daily links。
- `[Atlas 主线 Issue（Month 06）](../atlas/issues/month-06/atlas-m06-04.md)`：把本周主题放回 Atlas 主线问题卡，观察范围、验收项和剩余风险。
- `[Week 25 周报](../reports/weekly/week-25-report.md)`：看本周结果、证据和后续动作是如何被记录的。
- `[docs/reviews/archive/month-06-review.md](../reviews/archive/month-06-review.md)`：复盘周重点看 orchestration 的真实弱点。

## 本周 Atlas 落地检查

- 最小落地动作：围绕“补多 Agent 测试”推进，让至少一项本周交付标准在仓库里可见。
- 必留证据：orchestration tests、fallback tests 和 runtime docs。
- 最大风险：demo 跑通一次，但没有任何可检查、可恢复、可追责的 runtime 证据。
- 审查不变量：每一条 orchestrated run 都应该留下完整的决策链。
- 本周最先要盯住的交付：多 Agent 测试集

## 本周最小演示场景

- 背景：这一周你需要把“多 Agent 测试与复盘”从概念推进成可被别人看懂、验证和质疑的工程结果。
- 你至少要能演示：
  - 多 Agent 测试集、运行时文档、第六个月复盘结论
  - “Orchestration tests”为什么会直接影响 Atlas 的主路径，而不只是局部实现
  - 风险“demo 跑通一次，但没有任何可检查、可恢复、可追责的 runtime 证据。”是如何被控制、暴露或留证的
- reviewer 大概率会追问：
  - 为什么这周先做“补多 Agent 测试”，而不是先做别的？
  - 不变量“每一条 orchestrated run 都应该留下完整的决策链。”如果被破坏，会先造成什么后果？
  - 你打算拿什么证据证明这周成果成立？提示：orchestration tests、fallback tests 和 runtime docs。

## 配套案例

- 本周最接近的案例：读 [专题案例包](casebook.md) 中的“案例 6：多 Agent 看起来更强，实际上没人知道谁负责”。
- 本周最该对照的失败模式：读 [失败场景图谱](failure-patterns.md) 中的“模式四：治理失守型故障”。
- 推荐顺序：先读本周知识清单，再读案例，再回本周 Atlas 观察路径和周报，检查自己能不能解释坏信号与最小证据。

## 本周工作顺序

- [`RFC`](artifact-playbook.md#artifact-rfc)：先把“补多 Agent 测试”对应的问题范围、目标、约束和非目标写清楚。
- [`Interface`](artifact-playbook.md#artifact-interface)：围绕“Orchestration tests”定义输入、输出、owner、错误语义和边界。
- [`Implementation`](artifact-playbook.md#artifact-implementation)：把最小主路径打通，至少让“多 Agent 测试集、运行时文档、第六个月复盘结论”里的一项开始可验证。
- [`Unit Test`](artifact-playbook.md#artifact-unit-test)：优先覆盖最容易击中风险“demo 跑通一次，但没有任何可检查、可恢复、可追责的 runtime 证据。”的失败路径、边界条件或 contract。
- [`Review`](artifact-playbook.md#artifact-review)：站在 reviewer 视角检查不变量“每一条 orchestrated run 都应该留下完整的决策链。”是否真的被系统保护，而不是只写在文档里。
- [`Benchmark（按需）`](artifact-playbook.md#artifact-benchmark)：本周不做性能 benchmark，但要保留一轮多 Agent fallback / orchestration 验证。
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

1. 如果要把“多 Agent 测试与复盘”做成 Atlas 可交付能力，你会如何设计主路径和边界？
2. 本周最值得优先验证的失败模式是什么？你会如何收集证据？
3. 这周的 RFC、实现、验证或复盘里，哪一项最能体现你的工程判断？为什么？
4. 如果时间被砍掉一半，你会保留什么，推迟什么？
5. 本周产物如何自然承接到下周，而不是变成一次性文档或一次性代码？

## 每日入口

- [Day 169](days-semester-2/day-169.md): 理解问题和边界
- [Day 170](days-semester-2/day-170.md): 边界与契约设计
- [Day 171](days-semester-2/day-171.md): 最小主路径实现
- [Day 172](days-semester-2/day-172.md): 补全真实可用路径
- [Day 173](days-semester-2/day-173.md): 验证与实验
- [Day 174](days-semester-2/day-174.md): 文档与评审
- [Day 175](days-semester-2/day-175.md): 复盘与 memory

## 配套阅读

- [Week 25](../weeks/week-25.md)
