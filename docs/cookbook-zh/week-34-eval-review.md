# Week 34：评估体系复盘与运维文档

最后更新：2026-06-30

## 本周目标

这一周把第八个月的评估体系收口成运维资产。

## 本周学完后你应该会什么

重点是让别人也能跑、能看、能解释。

## 本周学习重点

- eval ops
- trace interpretation
- review notes

## 本周 Cookbook 做法

### Recipe 1：补运维说明

### Recipe 2：整理 review note

### Recipe 3：做第八个月复盘

## 本周交付标准

- 评估运维说明
- review note
- 第八个月复盘结论

## 本周能力焦点

- 评估体系（`Evaluation`）
- 数据驱动决策（`Data-Driven Decisions`）
- 可观测性（`Observability`）

## 本周知识清单

- eval operations docs 要说明怎么跑 eval、怎么读失败、怎么安全更新 dataset。
- trace interpretation 要能区分模型失败、工具失败、策略失败和数据质量问题。
- review note 应该记录 calibration blind spot，而不是只列一个总分。
- 月度复盘要能说明 measurement 是否真的改变了工程决策。

## 本周 Atlas 观察路径

- `[Week 34 英文原文](../weeks/week-34.md)`：对照原始英文任务拆分，确认这一周的 deliverables、source reading 和 daily links。
- `[Atlas 主线 Issue（Month 08）](../atlas/issues/month-08/atlas-m08-04.md)`：把本周主题放回 Atlas 主线问题卡，观察范围、验收项和剩余风险。
- `[Week 34 周报](../reports/weekly/week-34-report.md)`：看本周结果、证据和后续动作是如何被记录的。
- `[docs/reviews/archive/month-08-review.md](../reviews/archive/month-08-review.md)`：复盘周重点看哪些评估证据真正改变了判断。

## 本周 Atlas 落地检查

- 最小落地动作：围绕“补运维说明”推进，让至少一项本周交付标准在仓库里可见。
- 必留证据：eval ops 文档、trace interpretation note 和 review packet。
- 最大风险：做了 dashboard 和 trace，但没有任何操作手册告诉别人怎么用。
- 审查不变量：所有周期性 eval job 都应该有面向人的运行说明。
- 本周最先要盯住的交付：评估运维说明

## 本周最小演示场景

- 背景：这一周你需要把“评估体系复盘与运维文档”从概念推进成可被别人看懂、验证和质疑的工程结果。
- 你至少要能演示：
  - 评估运维说明、review note、第八个月复盘结论
  - “eval ops”为什么会直接影响 Atlas 的主路径，而不只是局部实现
  - 风险“做了 dashboard 和 trace，但没有任何操作手册告诉别人怎么用。”是如何被控制、暴露或留证的
- reviewer 大概率会追问：
  - 为什么这周先做“补运维说明”，而不是先做别的？
  - 不变量“所有周期性 eval job 都应该有面向人的运行说明。”如果被破坏，会先造成什么后果？
  - 你打算拿什么证据证明这周成果成立？提示：eval ops 文档、trace interpretation note 和 review packet。

## 配套案例

- 本周最接近的案例：读 [专题案例包](casebook.md) 中的“案例 8：有了 trace 和评估，但没人知道这些数据该怎么用”。
- 本周最该对照的失败模式：读 [失败场景图谱](failure-patterns.md) 中的“模式三：质量证据失真型故障”。
- 推荐顺序：先读本周知识清单，再读案例，再回本周 Atlas 观察路径和周报，检查自己能不能解释坏信号与最小证据。

## 本周工作顺序

- [`RFC`](artifact-playbook.md#artifact-rfc)：先把“补运维说明”对应的问题范围、目标、约束和非目标写清楚。
- [`Interface`](artifact-playbook.md#artifact-interface)：围绕“eval ops”定义输入、输出、owner、错误语义和边界。
- [`Implementation`](artifact-playbook.md#artifact-implementation)：把最小主路径打通，至少让“评估运维说明、review note、第八个月复盘结论”里的一项开始可验证。
- [`Unit Test`](artifact-playbook.md#artifact-unit-test)：优先覆盖最容易击中风险“做了 dashboard 和 trace，但没有任何操作手册告诉别人怎么用。”的失败路径、边界条件或 contract。
- [`Review`](artifact-playbook.md#artifact-review)：站在 reviewer 视角检查不变量“所有周期性 eval job 都应该有面向人的运行说明。”是否真的被系统保护，而不是只写在文档里。
- [`Benchmark（按需）`](artifact-playbook.md#artifact-benchmark)：如果本周调整了指标或 gate，至少保留一次前后对照结果。
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

1. 如果要把“评估体系复盘与运维文档”做成 Atlas 可交付能力，你会如何设计主路径和边界？
2. 本周最值得优先验证的失败模式是什么？你会如何收集证据？
3. 这周的 RFC、实现、验证或复盘里，哪一项最能体现你的工程判断？为什么？
4. 如果时间被砍掉一半，你会保留什么，推迟什么？
5. 本周产物如何自然承接到下周，而不是变成一次性文档或一次性代码？

## 每日入口

- [Day 232](days-semester-3/day-232.md): 理解问题和边界
- [Day 233](days-semester-3/day-233.md): 边界与契约设计
- [Day 234](days-semester-3/day-234.md): 最小主路径实现
- [Day 235](days-semester-3/day-235.md): 补全真实可用路径
- [Day 236](days-semester-3/day-236.md): 验证与实验
- [Day 237](days-semester-3/day-237.md): 文档与评审
- [Day 238](days-semester-3/day-238.md): 复盘与 memory

## 配套阅读

- [Week 34](../weeks/week-34.md)
