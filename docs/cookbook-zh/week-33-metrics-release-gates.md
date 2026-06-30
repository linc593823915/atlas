# Week 33：指标、Regression Gate 与发布规则

最后更新：2026-06-30

## 本周目标

本周进入“什么时候算过关”的问题。

## 本周学完后你应该会什么

你要把评估结果转成真正的 release gate。

## 本周学习重点

- metrics
- thresholds
- release policy

## 本周 Cookbook 做法

### Recipe 1：定义阈值

### Recipe 2：定义 regression gate

### Recipe 3：整理指标口径

## 本周交付标准

- release gate 规则
- 指标阈值说明
- 回归判断逻辑

## 本周能力焦点

- 评估体系（`Evaluation`）
- 数据驱动决策（`Data-Driven Decisions`）
- 可观测性（`Observability`）

## 本周知识清单

- metrics 要和用户或 operator 感知到的失败挂钩，而不是收集一堆 vanity counter。
- threshold 和 regression gate 需要 owner、升级路径和误报处理策略。
- release policy 必须说明什么阻断发布、什么只告警、什么仅留证据。
- 门禁只有在可重复 workload 上运行才有意义，否则每次结果都不可比较。

## 本周 Atlas 观察路径

- `[Week 33 英文原文](../weeks/week-33.md)`：对照原始英文任务拆分，确认这一周的 deliverables、source reading 和 daily links。
- `[Atlas 主线 Issue（Month 08）](../atlas/issues/month-08/atlas-m08-03.md)`：把本周主题放回 Atlas 主线问题卡，观察范围、验收项和剩余风险。
- `[Week 33 周报](../reports/weekly/week-33-report.md)`：看本周结果、证据和后续动作是如何被记录的。
- `[docs/reports/monthly/month-08-scorecard.md](../reports/monthly/month-08-scorecard.md)`：看评估体系如何进入月度验收。

## 本周 Atlas 落地检查

- 最小落地动作：围绕“定义阈值”推进，让至少一项本周交付标准在仓库里可见。
- 必留证据：metric spec、threshold table 和 release policy。
- 最大风险：设了很多 gate，但没人信、没人懂，也没人据此做决定。
- 审查不变量：一次 gate failure 必须能对应到清楚的发布或调查动作。
- 本周最先要盯住的交付：release gate 规则

## 本周最小演示场景

- 背景：这一周你需要把“指标、Regression Gate 与发布规则”从概念推进成可被别人看懂、验证和质疑的工程结果。
- 你至少要能演示：
  - release gate 规则、指标阈值说明、回归判断逻辑
  - “metrics”为什么会直接影响 Atlas 的主路径，而不只是局部实现
  - 风险“设了很多 gate，但没人信、没人懂，也没人据此做决定。”是如何被控制、暴露或留证的
- reviewer 大概率会追问：
  - 为什么这周先做“定义阈值”，而不是先做别的？
  - 不变量“一次 gate failure 必须能对应到清楚的发布或调查动作。”如果被破坏，会先造成什么后果？
  - 你打算拿什么证据证明这周成果成立？提示：metric spec、threshold table 和 release policy。

## 配套案例

- 本周最接近的案例：读 [专题案例包](casebook.md) 中的“案例 8：有了 trace 和评估，但没人知道这些数据该怎么用”。
- 本周最该对照的失败模式：读 [失败场景图谱](failure-patterns.md) 中的“模式三：质量证据失真型故障”。
- 推荐顺序：先读本周知识清单，再读案例，再回本周 Atlas 观察路径和周报，检查自己能不能解释坏信号与最小证据。

## 本周工作顺序

- [`RFC`](artifact-playbook.md#artifact-rfc)：先把“定义阈值”对应的问题范围、目标、约束和非目标写清楚。
- [`Interface`](artifact-playbook.md#artifact-interface)：围绕“metrics”定义输入、输出、owner、错误语义和边界。
- [`Implementation`](artifact-playbook.md#artifact-implementation)：把最小主路径打通，至少让“release gate 规则、指标阈值说明、回归判断逻辑”里的一项开始可验证。
- [`Unit Test`](artifact-playbook.md#artifact-unit-test)：优先覆盖最容易击中风险“设了很多 gate，但没人信、没人懂，也没人据此做决定。”的失败路径、边界条件或 contract。
- [`Review`](artifact-playbook.md#artifact-review)：站在 reviewer 视角检查不变量“一次 gate failure 必须能对应到清楚的发布或调查动作。”是否真的被系统保护，而不是只写在文档里。
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

1. 如果要把“指标、Regression Gate 与发布规则”做成 Atlas 可交付能力，你会如何设计主路径和边界？
2. 本周最值得优先验证的失败模式是什么？你会如何收集证据？
3. 这周的 RFC、实现、验证或复盘里，哪一项最能体现你的工程判断？为什么？
4. 如果时间被砍掉一半，你会保留什么，推迟什么？
5. 本周产物如何自然承接到下周，而不是变成一次性文档或一次性代码？

## 每日入口

- [Day 225](days-semester-3/day-225.md): 理解问题和边界
- [Day 226](days-semester-3/day-226.md): 边界与契约设计
- [Day 227](days-semester-3/day-227.md): 最小主路径实现
- [Day 228](days-semester-3/day-228.md): 补全真实可用路径
- [Day 229](days-semester-3/day-229.md): 验证与实验
- [Day 230](days-semester-3/day-230.md): 文档与评审
- [Day 231](days-semester-3/day-231.md): 复盘与 memory

## 配套阅读

- [Week 33](../weeks/week-33.md)
