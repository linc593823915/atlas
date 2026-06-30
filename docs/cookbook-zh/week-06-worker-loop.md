# Week 06：Worker 主循环与执行路径

最后更新：2026-06-30

## 本周目标

这一周开始让后台任务真正跑起来。

你要做的不只是“从队列里取任务”，而是设计一个可启动、可停止、可观察的 worker 主循环。

## 本周学习重点

- worker 生命周期
- graceful shutdown
- 入队与执行主路径

## 本周 Cookbook 做法

### Recipe 1：建立 worker loop

### Recipe 2：跑通 enqueue 与 execute

### Recipe 3：补 shutdown 行为

## 本周交付标准

你至少要能演示：

- 任务如何进入队列
- worker 如何消费
- worker 如何被优雅关闭

## 本周能力焦点

- 可信可靠（`Dependability`）
- 可靠性（`Reliability`）
- 运维思维（`Operational Thinking`）

## 本周知识清单

- worker 主循环本质上是 claim -> execute -> ack/release，每一步都要有失败语义。
- graceful shutdown 的关键是停拉新任务、排空在途任务并保存必要进度，而不是简单 kill 进程。
- 并发度和 backpressure 策略决定了队列积压会变成延迟问题还是雪崩问题。
- 如果任务执行时间可能超过一次 poll 周期，就要考虑 lease、heartbeat 或 owner 续期。

## 本周 Atlas 观察路径

- `[Week 06 英文原文](../weeks/week-06.md)`：对照原始英文任务拆分，确认这一周的 deliverables、source reading 和 daily links。
- `[Atlas 主线 Issue（Month 02）](../atlas/issues/month-02/atlas-m02-02.md)`：把本周主题放回 Atlas 主线问题卡，观察范围、验收项和剩余风险。
- `[Week 06 周报](../reports/weekly/week-06-report.md)`：看本周结果、证据和后续动作是如何被记录的。
- `[docs/operations/review-metrics.md](../operations/review-metrics.md)`：观察失败路径和可靠性指标在治理上如何落表。

## 本周 Atlas 落地检查

- 最小落地动作：围绕“建立 worker loop”推进，让至少一项本周交付标准在仓库里可见。
- 必留证据：worker loop 设计、shutdown drill 和主路径实现说明。
- 最大风险：worker 崩溃或退出时同时造成重复执行和任务丢失。
- 审查不变量：同一个 job 在没有显式 lease 转移的情况下不能被两个 worker 同时拥有。
- 本周最先要盯住的交付：任务如何进入队列

## 本周最小演示场景

- 背景：这一周你需要把“Worker 主循环与执行路径”从概念推进成可被别人看懂、验证和质疑的工程结果。
- 你至少要能演示：
  - 任务如何进入队列、worker 如何消费、worker 如何被优雅关闭
  - “worker 生命周期”为什么会直接影响 Atlas 的主路径，而不只是局部实现
  - 风险“worker 崩溃或退出时同时造成重复执行和任务丢失。”是如何被控制、暴露或留证的
- reviewer 大概率会追问：
  - 为什么这周先做“建立 worker loop”，而不是先做别的？
  - 不变量“同一个 job 在没有显式 lease 转移的情况下不能被两个 worker 同时拥有。”如果被破坏，会先造成什么后果？
  - 你打算拿什么证据证明这周成果成立？提示：worker loop 设计、shutdown drill 和主路径实现说明。

## 配套案例

- 本周最接近的案例：读 [专题案例包](casebook.md) 中的“案例 2：任务系统看起来能跑，线上却越跑越乱”。
- 本周最该对照的失败模式：读 [失败场景图谱](failure-patterns.md) 中的“模式二：失败语义缺失型故障”。
- 推荐顺序：先读本周知识清单，再读案例，再回本周 Atlas 观察路径和周报，检查自己能不能解释坏信号与最小证据。

## 本周工作顺序

- [`RFC`](artifact-playbook.md#artifact-rfc)：先把“建立 worker loop”对应的问题范围、目标、约束和非目标写清楚。
- [`Interface`](artifact-playbook.md#artifact-interface)：围绕“worker 生命周期”定义输入、输出、owner、错误语义和边界。
- [`Implementation`](artifact-playbook.md#artifact-implementation)：把最小主路径打通，至少让“任务如何进入队列、worker 如何消费、worker 如何被优雅关闭”里的一项开始可验证。
- [`Unit Test`](artifact-playbook.md#artifact-unit-test)：优先覆盖最容易击中风险“worker 崩溃或退出时同时造成重复执行和任务丢失。”的失败路径、边界条件或 contract。
- [`Review`](artifact-playbook.md#artifact-review)：站在 reviewer 视角检查不变量“同一个 job 在没有显式 lease 转移的情况下不能被两个 worker 同时拥有。”是否真的被系统保护，而不是只写在文档里。
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

1. 如果要把“Worker 主循环与执行路径”做成 Atlas 可交付能力，你会如何设计主路径和边界？
2. 本周最值得优先验证的失败模式是什么？你会如何收集证据？
3. 这周的 RFC、实现、验证或复盘里，哪一项最能体现你的工程判断？为什么？
4. 如果时间被砍掉一半，你会保留什么，推迟什么？
5. 本周产物如何自然承接到下周，而不是变成一次性文档或一次性代码？

## 每日入口

- [Day 036](days-semester-1/day-036.md): 理解问题和边界
- [Day 037](days-semester-1/day-037.md): 边界与契约设计
- [Day 038](days-semester-1/day-038.md): 最小主路径实现
- [Day 039](days-semester-1/day-039.md): 补全真实可用路径
- [Day 040](days-semester-1/day-040.md): 验证与实验
- [Day 041](days-semester-1/day-041.md): 文档与评审
- [Day 042](days-semester-1/day-042.md): 复盘与 memory

## 配套阅读

- [Week 06](../weeks/week-06.md)
- [Week 06 Report](../reports/weekly/week-06-report.md)
