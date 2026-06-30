# Week 45：优化循环与容量规划

最后更新：2026-06-30

## 本周目标

这一周从“测量”进入“优化”。

## 本周学完后你应该会什么

重点是做一次有证据的 before/after 优化。

## 本周学习重点

- optimization loop
- capacity assumption
- bottleneck analysis

## 本周 Cookbook 做法

### Recipe 1：做一次优化

### Recipe 2：写 capacity 假设

### Recipe 3：解释瓶颈

## 本周交付标准

- before/after 数据
- capacity 说明
- 瓶颈分析

## 本周能力焦点

- 性能工程（`Performance Engineering`）
- 成本意识（`Cost Awareness`）
- 容量规划（`Capacity Planning`）

## 本周知识清单

- optimization loop 是提出瓶颈假设、改一件事、再测、再决定，而不是同时乱改十处。
- capacity planning 从 workload mix、并发度和饱和点假设开始。
- 成本和性能的取舍是架构选择，不只是财务报表。
- bottleneck 分析要区分时间花在 model call、tool、queue 还是 infrastructure 上。

## 本周 Atlas 观察路径

- `[Week 45 英文原文](../weeks/week-45.md)`：对照原始英文任务拆分，确认这一周的 deliverables、source reading 和 daily links。
- `[Atlas 主线 Issue（Month 11）](../atlas/issues/month-11/atlas-m11-02.md)`：把本周主题放回 Atlas 主线问题卡，观察范围、验收项和剩余风险。
- `[Week 45 周报](../reports/weekly/week-45-report.md)`：看本周结果、证据和后续动作是如何被记录的。
- `[docs/reports/monthly/month-11-scorecard.md](../reports/monthly/month-11-scorecard.md)`：对照 baseline、优化和容量判断的月度评分标准。

## 本周 Atlas 落地检查

- 最小落地动作：围绕“做一次优化”推进，让至少一项本周交付标准在仓库里可见。
- 必留证据：before/after data、bottleneck analysis 和 capacity assumptions。
- 最大风险：不知道主瓶颈在哪，就开始做一堆看起来很勤奋的小优化。
- 审查不变量：每一条优化收益都必须同时展示收益和代价。
- 本周最先要盯住的交付：before/after 数据

## 本周最小演示场景

- 背景：这一周你需要把“优化循环与容量规划”从概念推进成可被别人看懂、验证和质疑的工程结果。
- 你至少要能演示：
  - before/after 数据、capacity 说明、瓶颈分析
  - “optimization loop”为什么会直接影响 Atlas 的主路径，而不只是局部实现
  - 风险“不知道主瓶颈在哪，就开始做一堆看起来很勤奋的小优化。”是如何被控制、暴露或留证的
- reviewer 大概率会追问：
  - 为什么这周先做“做一次优化”，而不是先做别的？
  - 不变量“每一条优化收益都必须同时展示收益和代价。”如果被破坏，会先造成什么后果？
  - 你打算拿什么证据证明这周成果成立？提示：before/after data、bottleneck analysis 和 capacity assumptions。

## 配套案例

- 本周最接近的案例：读 [专题案例包](casebook.md) 中的“案例 11：你拿到一堆性能数字，但一个能信的结论都没有”。
- 本周最该对照的失败模式：读 [失败场景图谱](failure-patterns.md) 中的“模式三：质量证据失真型故障”。
- 推荐顺序：先读本周知识清单，再读案例，再回本周 Atlas 观察路径和周报，检查自己能不能解释坏信号与最小证据。

## 本周工作顺序

- [`RFC`](artifact-playbook.md#artifact-rfc)：先把“做一次优化”对应的问题范围、目标、约束和非目标写清楚。
- [`Interface`](artifact-playbook.md#artifact-interface)：围绕“optimization loop”定义输入、输出、owner、错误语义和边界。
- [`Implementation`](artifact-playbook.md#artifact-implementation)：把最小主路径打通，至少让“before/after 数据、capacity 说明、瓶颈分析”里的一项开始可验证。
- [`Unit Test`](artifact-playbook.md#artifact-unit-test)：优先覆盖最容易击中风险“不知道主瓶颈在哪，就开始做一堆看起来很勤奋的小优化。”的失败路径、边界条件或 contract。
- [`Review`](artifact-playbook.md#artifact-review)：站在 reviewer 视角检查不变量“每一条优化收益都必须同时展示收益和代价。”是否真的被系统保护，而不是只写在文档里。
- [`Benchmark（按需）`](artifact-playbook.md#artifact-benchmark)：本周 benchmark 是主题核心，必须保留 before / after 和瓶颈解释。
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

1. 如果要把“优化循环与容量规划”做成 Atlas 可交付能力，你会如何设计主路径和边界？
2. 本周最值得优先验证的失败模式是什么？你会如何收集证据？
3. 这周的 RFC、实现、验证或复盘里，哪一项最能体现你的工程判断？为什么？
4. 如果时间被砍掉一半，你会保留什么，推迟什么？
5. 本周产物如何自然承接到下周，而不是变成一次性文档或一次性代码？

## 每日入口

- [Day 309](days-semester-4/day-309.md): 理解问题和边界
- [Day 310](days-semester-4/day-310.md): 边界与契约设计
- [Day 311](days-semester-4/day-311.md): 最小主路径实现
- [Day 312](days-semester-4/day-312.md): 补全真实可用路径
- [Day 313](days-semester-4/day-313.md): 验证与实验
- [Day 314](days-semester-4/day-314.md): 文档与评审
- [Day 315](days-semester-4/day-315.md): 复盘与 memory

## 配套阅读

- [Week 45](../weeks/week-45.md)
