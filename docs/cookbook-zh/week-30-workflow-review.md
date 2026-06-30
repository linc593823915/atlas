# Week 30：工作流测试与复盘

最后更新：2026-06-30

## 本周目标

这一周把工作流模块收口。重点是证明它真的可恢复、可解释、可交付。

## 本周学完后你应该会什么

你需要通过测试和文档把复杂度压住。

## 本周学习重点

- resume tests
- stateful testing
- workflow docs

## 本周 Cookbook 做法

### Recipe 1：补工作流测试

### Recipe 2：补恢复说明

### Recipe 3：做第七个月复盘

## 本周交付标准

- 工作流测试集
- 恢复文档
- 第七个月复盘结论

## 本周能力焦点

- 有状态系统（`Stateful Systems`）
- 恢复设计（`Recovery Design`）
- 运行韧性（`Operational Resilience`）

## 本周知识清单

- stateful tests 要模拟重启、重复 resume、过期 checkpoint 和分支失败。
- 长流程 drill 需要覆盖 operator takeover、卡死状态和 checkpoint 损坏后的恢复决策。
- workflow docs 必须教 operator 看懂状态、决定是 resume、replay 还是 abort。
- 月度复盘要说明 durable model 是真正可运行的，还是只在理论上可恢复。

## 本周 Atlas 观察路径

- `[Week 30 英文原文](../weeks/week-30.md)`：对照原始英文任务拆分，确认这一周的 deliverables、source reading 和 daily links。
- `[Atlas 主线 Issue（Month 07）](../atlas/issues/month-07/atlas-m07-04.md)`：把本周主题放回 Atlas 主线问题卡，观察范围、验收项和剩余风险。
- `[Week 30 周报](../reports/weekly/week-30-report.md)`：看本周结果、证据和后续动作是如何被记录的。
- `[docs/reviews/archive/month-07-review.md](../reviews/archive/month-07-review.md)`：看 durable workflow 复盘真正追问什么。

## 本周 Atlas 落地检查

- 最小落地动作：围绕“补工作流测试”推进，让至少一项本周交付标准在仓库里可见。
- 必留证据：stateful tests、recovery drills 和 workflow docs。
- 最大风险：只验证本地 happy-path resume，从未演练真实恢复流程。
- 审查不变量：每一条恢复路径都应该对应清楚的 operator decision 和测试证据。
- 本周最先要盯住的交付：工作流测试集

## 本周最小演示场景

- 背景：这一周你需要把“工作流测试与复盘”从概念推进成可被别人看懂、验证和质疑的工程结果。
- 你至少要能演示：
  - 工作流测试集、恢复文档、第七个月复盘结论
  - “resume tests”为什么会直接影响 Atlas 的主路径，而不只是局部实现
  - 风险“只验证本地 happy-path resume，从未演练真实恢复流程。”是如何被控制、暴露或留证的
- reviewer 大概率会追问：
  - 为什么这周先做“补工作流测试”，而不是先做别的？
  - 不变量“每一条恢复路径都应该对应清楚的 operator decision 和测试证据。”如果被破坏，会先造成什么后果？
  - 你打算拿什么证据证明这周成果成立？提示：stateful tests、recovery drills 和 workflow docs。

## 配套案例

- 本周最接近的案例：读 [专题案例包](casebook.md) 中的“案例 7：长流程跑得很长，但根本恢复不了”。
- 本周最该对照的失败模式：读 [失败场景图谱](failure-patterns.md) 中的“模式二：失败语义缺失型故障”。
- 推荐顺序：先读本周知识清单，再读案例，再回本周 Atlas 观察路径和周报，检查自己能不能解释坏信号与最小证据。

## 本周工作顺序

- [`RFC`](artifact-playbook.md#artifact-rfc)：先把“补工作流测试”对应的问题范围、目标、约束和非目标写清楚。
- [`Interface`](artifact-playbook.md#artifact-interface)：围绕“resume tests”定义输入、输出、owner、错误语义和边界。
- [`Implementation`](artifact-playbook.md#artifact-implementation)：把最小主路径打通，至少让“工作流测试集、恢复文档、第七个月复盘结论”里的一项开始可验证。
- [`Unit Test`](artifact-playbook.md#artifact-unit-test)：优先覆盖最容易击中风险“只验证本地 happy-path resume，从未演练真实恢复流程。”的失败路径、边界条件或 contract。
- [`Review`](artifact-playbook.md#artifact-review)：站在 reviewer 视角检查不变量“每一条恢复路径都应该对应清楚的 operator decision 和测试证据。”是否真的被系统保护，而不是只写在文档里。
- [`Benchmark（按需）`](artifact-playbook.md#artifact-benchmark)：本周不做性能 benchmark，但要保留一轮恢复 drill 或 stateful test 结果。
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

1. 如果要把“工作流测试与复盘”做成 Atlas 可交付能力，你会如何设计主路径和边界？
2. 本周最值得优先验证的失败模式是什么？你会如何收集证据？
3. 这周的 RFC、实现、验证或复盘里，哪一项最能体现你的工程判断？为什么？
4. 如果时间被砍掉一半，你会保留什么，推迟什么？
5. 本周产物如何自然承接到下周，而不是变成一次性文档或一次性代码？

## 每日入口

- [Day 204](days-semester-3/day-204.md): 理解问题和边界
- [Day 205](days-semester-3/day-205.md): 边界与契约设计
- [Day 206](days-semester-3/day-206.md): 最小主路径实现
- [Day 207](days-semester-3/day-207.md): 补全真实可用路径
- [Day 208](days-semester-3/day-208.md): 验证与实验
- [Day 209](days-semester-3/day-209.md): 文档与评审
- [Day 210](days-semester-3/day-210.md): 复盘与 memory

## 配套阅读

- [Week 30](../weeks/week-30.md)
