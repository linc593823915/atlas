# Week 43：Benchmark RFC

最后更新：2026-06-30

## 本周目标

本周开始进入性能与成本视角。先把 benchmark 要测什么、为什么测清楚。

## 本周学完后你应该会什么

重点不是先跑压测，而是先定义测量契约。

## 本周学习重点

- benchmark contract
- latency metric
- cost metric

## 本周 Cookbook 做法

### Recipe 1：写 benchmark RFC

### Recipe 2：定义指标口径

### Recipe 3：定义工作负载模型

## 本周交付标准

- benchmark RFC
- 指标口径
- 工作负载定义

## 本周能力焦点

- 生产所有权（`Production Ownership`）
- SLO 思维（`SLO Thinking`）
- 回滚准备度（`Rollback Readiness`）

## 本周知识清单

- benchmark RFC 先冻结 workload、environment、metric 和 repeatability 假设，再收数字。
- latency、cost 和 throughput 只有挂在代表性场景上才有意义。
- benchmark contract 要说清楚模型时间、tool 时间、queue 时间和 infra 时间哪些在统计范围内。
- 没有 contract 的测量只会制造误导性的 before/after 对比。

## 本周 Atlas 观察路径

- `[Week 43 英文原文](../weeks/week-43.md)`：对照原始英文任务拆分，确认这一周的 deliverables、source reading 和 daily links。
- `[Atlas 主线 Issue（Month 10）](../atlas/issues/month-10/atlas-m10-04.md)`：把本周主题放回 Atlas 主线问题卡，观察范围、验收项和剩余风险。
- `[Week 43 周报](../reports/weekly/week-43-report.md)`：看本周结果、证据和后续动作是如何被记录的。
- `[docs/reviews/archive/month-10-review.md](../reviews/archive/month-10-review.md)`：复盘周重点看哪些运维假设被证据打脸。

## 本周 Atlas 落地检查

- 最小落地动作：围绕“写 benchmark RFC”推进，让至少一项本周交付标准在仓库里可见。
- 必留证据：benchmark RFC、metric 定义和 workload set。
- 最大风险：拿不同环境、不同 workload 的数字放在一起比较，还以为自己在做性能工程。
- 审查不变量：每一条 benchmark 结果都必须能从文档化 setup 中复现。
- 本周最先要盯住的交付：benchmark RFC

## 本周最小演示场景

- 背景：这一周你需要把“Benchmark RFC”从概念推进成可被别人看懂、验证和质疑的工程结果。
- 你至少要能演示：
  - benchmark RFC、指标口径、工作负载定义
  - “benchmark contract”为什么会直接影响 Atlas 的主路径，而不只是局部实现
  - 风险“拿不同环境、不同 workload 的数字放在一起比较，还以为自己在做性能工程。”是如何被控制、暴露或留证的
- reviewer 大概率会追问：
  - 为什么这周先做“写 benchmark RFC”，而不是先做别的？
  - 不变量“每一条 benchmark 结果都必须能从文档化 setup 中复现。”如果被破坏，会先造成什么后果？
  - 你打算拿什么证据证明这周成果成立？提示：benchmark RFC、metric 定义和 workload set。

## 配套案例

- 本周最接近的案例：读 [专题案例包](casebook.md) 中的“案例 10：Kubernetes 部署了，但一点都不像生产系统”。
- 本周最该对照的失败模式：读 [失败场景图谱](failure-patterns.md) 中的“模式五：平台错觉型故障”。
- 推荐顺序：先读本周知识清单，再读案例，再回本周 Atlas 观察路径和周报，检查自己能不能解释坏信号与最小证据。

## 本周工作顺序

- [`RFC`](artifact-playbook.md#artifact-rfc)：先把“写 benchmark RFC”对应的问题范围、目标、约束和非目标写清楚。
- [`Interface`](artifact-playbook.md#artifact-interface)：围绕“benchmark contract”定义输入、输出、owner、错误语义和边界。
- [`Implementation`](artifact-playbook.md#artifact-implementation)：把最小主路径打通，至少让“benchmark RFC、指标口径、工作负载定义”里的一项开始可验证。
- [`Unit Test`](artifact-playbook.md#artifact-unit-test)：优先覆盖最容易击中风险“拿不同环境、不同 workload 的数字放在一起比较，还以为自己在做性能工程。”的失败路径、边界条件或 contract。
- [`Review`](artifact-playbook.md#artifact-review)：站在 reviewer 视角检查不变量“每一条 benchmark 结果都必须能从文档化 setup 中复现。”是否真的被系统保护，而不是只写在文档里。
- [`Benchmark（按需）`](artifact-playbook.md#artifact-benchmark)：本周 benchmark 是主题核心，必须固定 workload、环境和指标口径。
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

1. 如果要把“Benchmark RFC”做成 Atlas 可交付能力，你会如何设计主路径和边界？
2. 本周最值得优先验证的失败模式是什么？你会如何收集证据？
3. 这周的 RFC、实现、验证或复盘里，哪一项最能体现你的工程判断？为什么？
4. 如果时间被砍掉一半，你会保留什么，推迟什么？
5. 本周产物如何自然承接到下周，而不是变成一次性文档或一次性代码？

## 每日入口

- [Day 295](days-semester-4/day-295.md): 理解问题和边界
- [Day 296](days-semester-4/day-296.md): 边界与契约设计
- [Day 297](days-semester-4/day-297.md): 最小主路径实现
- [Day 298](days-semester-4/day-298.md): 补全真实可用路径
- [Day 299](days-semester-4/day-299.md): 验证与实验
- [Day 300](days-semester-4/day-300.md): 文档与评审
- [Day 301](days-semester-4/day-301.md): 复盘与 memory

## 配套阅读

- [Week 43](../weeks/week-43.md)
