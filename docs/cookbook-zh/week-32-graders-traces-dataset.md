# Week 32：Graders、Tracing 与数据集

最后更新：2026-06-30

## 本周目标

本周把评估和 tracing 真正接起来。

## 本周学完后你应该会什么

重点是让系统行为可观察，而不是靠日志猜。

## 本周学习重点

- graders
- trace capture
- benchmark dataset

## 本周 Cookbook 做法

### Recipe 1：接入 trace pipeline

### Recipe 2：准备 grader

### Recipe 3：构建 benchmark dataset

## 本周交付标准

- trace 样例
- grader 结构
- 数据集初稿

## 本周能力焦点

- 评估体系（`Evaluation`）
- 数据驱动决策（`Data-Driven Decisions`）
- 可观测性（`Observability`）

## 本周知识清单

- grader 只有在 trace 暴露了正确事件、输入、输出和决策点时才真正有用。
- trace capture 要足够便宜以便常开，也要足够有信息量，能定位 routing 和 policy 行为。
- benchmark dataset 需要覆盖正常、边界和 adversarial 样本，而不是只留最容易得分的题。
- 数据采集边界要兼顾隐私和 evaluability。

## 本周 Atlas 观察路径

- `[Week 32 英文原文](../weeks/week-32.md)`：对照原始英文任务拆分，确认这一周的 deliverables、source reading 和 daily links。
- `[Atlas 主线 Issue（Month 08）](../atlas/issues/month-08/atlas-m08-02.md)`：把本周主题放回 Atlas 主线问题卡，观察范围、验收项和剩余风险。
- `[Week 32 周报](../reports/weekly/week-32-report.md)`：看本周结果、证据和后续动作是如何被记录的。
- `[docs/operations/dashboard-export.md](../operations/dashboard-export.md)`：观察 trace / metric 输出如何变成运营面板材料。

## 本周 Atlas 落地检查

- 最小落地动作：围绕“接入 trace pipeline”推进，让至少一项本周交付标准在仓库里可见。
- 必留证据：trace model、grader inputs 和 benchmark dataset。
- 最大风险：收集了一堆 verbose trace，却没有任何真正能支持评分和调试的结构。
- 审查不变量：每个 score 都应该能从 dataset row 和 trace 中被解释出来。
- 本周最先要盯住的交付：trace 样例

## 本周最小演示场景

- 背景：这一周你需要把“Graders、Tracing 与数据集”从概念推进成可被别人看懂、验证和质疑的工程结果。
- 你至少要能演示：
  - trace 样例、grader 结构、数据集初稿
  - “graders”为什么会直接影响 Atlas 的主路径，而不只是局部实现
  - 风险“收集了一堆 verbose trace，却没有任何真正能支持评分和调试的结构。”是如何被控制、暴露或留证的
- reviewer 大概率会追问：
  - 为什么这周先做“接入 trace pipeline”，而不是先做别的？
  - 不变量“每个 score 都应该能从 dataset row 和 trace 中被解释出来。”如果被破坏，会先造成什么后果？
  - 你打算拿什么证据证明这周成果成立？提示：trace model、grader inputs 和 benchmark dataset。

## 配套案例

- 本周最接近的案例：读 [专题案例包](casebook.md) 中的“案例 8：有了 trace 和评估，但没人知道这些数据该怎么用”。
- 本周最该对照的失败模式：读 [失败场景图谱](failure-patterns.md) 中的“模式三：质量证据失真型故障”。
- 推荐顺序：先读本周知识清单，再读案例，再回本周 Atlas 观察路径和周报，检查自己能不能解释坏信号与最小证据。

## 本周工作顺序

- [`RFC`](artifact-playbook.md#artifact-rfc)：先把“接入 trace pipeline”对应的问题范围、目标、约束和非目标写清楚。
- [`Interface`](artifact-playbook.md#artifact-interface)：围绕“graders”定义输入、输出、owner、错误语义和边界。
- [`Implementation`](artifact-playbook.md#artifact-implementation)：把最小主路径打通，至少让“trace 样例、grader 结构、数据集初稿”里的一项开始可验证。
- [`Unit Test`](artifact-playbook.md#artifact-unit-test)：优先覆盖最容易击中风险“收集了一堆 verbose trace，却没有任何真正能支持评分和调试的结构。”的失败路径、边界条件或 contract。
- [`Review`](artifact-playbook.md#artifact-review)：站在 reviewer 视角检查不变量“每个 score 都应该能从 dataset row 和 trace 中被解释出来。”是否真的被系统保护，而不是只写在文档里。
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

1. 如果要把“Graders、Tracing 与数据集”做成 Atlas 可交付能力，你会如何设计主路径和边界？
2. 本周最值得优先验证的失败模式是什么？你会如何收集证据？
3. 这周的 RFC、实现、验证或复盘里，哪一项最能体现你的工程判断？为什么？
4. 如果时间被砍掉一半，你会保留什么，推迟什么？
5. 本周产物如何自然承接到下周，而不是变成一次性文档或一次性代码？

## 每日入口

- [Day 218](days-semester-3/day-218.md): 理解问题和边界
- [Day 219](days-semester-3/day-219.md): 边界与契约设计
- [Day 220](days-semester-3/day-220.md): 最小主路径实现
- [Day 221](days-semester-3/day-221.md): 补全真实可用路径
- [Day 222](days-semester-3/day-222.md): 验证与实验
- [Day 223](days-semester-3/day-223.md): 文档与评审
- [Day 224](days-semester-3/day-224.md): 复盘与 memory

## 配套阅读

- [Week 32](../weeks/week-32.md)
