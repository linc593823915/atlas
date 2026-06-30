# Week 31：Eval Loop RFC

最后更新：2026-06-30

## 本周目标

这一周进入“系统如何被度量”这个问题。重点是先定义评估回路。

## 本周学完后你应该会什么

你要从“感觉系统更好”升级到“知道系统为什么更好或更差”。

## 本周学习重点

- eval loop
- grader strategy
- dataset structure

## 本周 Cookbook 做法

### Recipe 1：写 eval RFC

### Recipe 2：定义 grader 和 dataset

### Recipe 3：定义 release gate 起点

## 本周交付标准

- Eval Loop RFC
- grader 设计草稿
- dataset 结构草稿

## 本周能力焦点

- 评估体系（`Evaluation`）
- 数据驱动决策（`Data-Driven Decisions`）
- 可观测性（`Observability`）

## 本周知识清单

- eval loop 先定义 evaluation unit：输入、上下文、tool trace、期望行为和评分语义。
- grader 可以是 heuristic、model-based 或 hybrid，但每种都要有失败模式和校准规则。
- dataset 结构需要支持 versioning、labeling 和可复现 rerun。
- RFC 还要回答谁拥有 score interpretation，以及 score 如何进入发布决策。

## 本周 Atlas 观察路径

- `[Week 31 英文原文](../weeks/week-31.md)`：对照原始英文任务拆分，确认这一周的 deliverables、source reading 和 daily links。
- `[Atlas 主线 Issue（Month 08）](../atlas/issues/month-08/atlas-m08-01.md)`：把本周主题放回 Atlas 主线问题卡，观察范围、验收项和剩余风险。
- `[Week 31 周报](../reports/weekly/week-31-report.md)`：看本周结果、证据和后续动作是如何被记录的。
- `[docs/operations/review-metrics.md](../operations/review-metrics.md)`：对照 eval 指标和 release gate 的治理表述。

## 本周 Atlas 落地检查

- 最小落地动作：围绕“写 eval RFC”推进，让至少一项本周交付标准在仓库里可见。
- 必留证据：eval loop RFC、grader strategy 和 dataset schema。
- 最大风险：把 eval 当作零散样例集合，没有任何稳定 loop 语义。
- 审查不变量：同一条数据在同一版本代码上应该产生可比较的评估证据。
- 本周最先要盯住的交付：Eval Loop RFC

## 本周最小演示场景

- 背景：这一周你需要把“Eval Loop RFC”从概念推进成可被别人看懂、验证和质疑的工程结果。
- 你至少要能演示：
  - Eval Loop RFC、grader 设计草稿、dataset 结构草稿
  - “eval loop”为什么会直接影响 Atlas 的主路径，而不只是局部实现
  - 风险“把 eval 当作零散样例集合，没有任何稳定 loop 语义。”是如何被控制、暴露或留证的
- reviewer 大概率会追问：
  - 为什么这周先做“写 eval RFC”，而不是先做别的？
  - 不变量“同一条数据在同一版本代码上应该产生可比较的评估证据。”如果被破坏，会先造成什么后果？
  - 你打算拿什么证据证明这周成果成立？提示：eval loop RFC、grader strategy 和 dataset schema。

## 配套案例

- 本周最接近的案例：读 [专题案例包](casebook.md) 中的“案例 8：有了 trace 和评估，但没人知道这些数据该怎么用”。
- 本周最该对照的失败模式：读 [失败场景图谱](failure-patterns.md) 中的“模式三：质量证据失真型故障”。
- 推荐顺序：先读本周知识清单，再读案例，再回本周 Atlas 观察路径和周报，检查自己能不能解释坏信号与最小证据。

## 本周工作顺序

- [`RFC`](artifact-playbook.md#artifact-rfc)：先把“写 eval RFC”对应的问题范围、目标、约束和非目标写清楚。
- [`Interface`](artifact-playbook.md#artifact-interface)：围绕“eval loop”定义输入、输出、owner、错误语义和边界。
- [`Implementation`](artifact-playbook.md#artifact-implementation)：把最小主路径打通，至少让“Eval Loop RFC、grader 设计草稿、dataset 结构草稿”里的一项开始可验证。
- [`Unit Test`](artifact-playbook.md#artifact-unit-test)：优先覆盖最容易击中风险“把 eval 当作零散样例集合，没有任何稳定 loop 语义。”的失败路径、边界条件或 contract。
- [`Review`](artifact-playbook.md#artifact-review)：站在 reviewer 视角检查不变量“同一条数据在同一版本代码上应该产生可比较的评估证据。”是否真的被系统保护，而不是只写在文档里。
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

1. 如果要把“Eval Loop RFC”做成 Atlas 可交付能力，你会如何设计主路径和边界？
2. 本周最值得优先验证的失败模式是什么？你会如何收集证据？
3. 这周的 RFC、实现、验证或复盘里，哪一项最能体现你的工程判断？为什么？
4. 如果时间被砍掉一半，你会保留什么，推迟什么？
5. 本周产物如何自然承接到下周，而不是变成一次性文档或一次性代码？

## 每日入口

- [Day 211](days-semester-3/day-211.md): 理解问题和边界
- [Day 212](days-semester-3/day-212.md): 边界与契约设计
- [Day 213](days-semester-3/day-213.md): 最小主路径实现
- [Day 214](days-semester-3/day-214.md): 补全真实可用路径
- [Day 215](days-semester-3/day-215.md): 验证与实验
- [Day 216](days-semester-3/day-216.md): 文档与评审
- [Day 217](days-semester-3/day-217.md): 复盘与 memory

## 配套阅读

- [Week 31](../weeks/week-31.md)
