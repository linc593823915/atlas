# Week 46：平台整合 RFC

最后更新：2026-06-30

## 本周目标

这一周开始准备毕业平台形态。你要先把最终平台的子系统边界说清楚。

## 本周学完后你应该会什么

不是先做整合，而是先回答平台到底由哪些部分组成。

## 本周学习重点

- platform boundary
- subsystem map
- integration order

## 本周 Cookbook 做法

### Recipe 1：写平台 RFC

### Recipe 2：画系统图

### Recipe 3：定义整合顺序

## 本周交付标准

- 平台 RFC
- 子系统图
- 整合顺序说明

## 本周能力焦点

- 性能工程（`Performance Engineering`）
- 成本意识（`Cost Awareness`）
- 容量规划（`Capacity Planning`）

## 本周知识清单

- platform RFC 要回答子系统是否能在一个统一 operating model 下真正组合起来。
- subsystem map 要显式画出依赖和整合顺序，避免集成期出现循环重构。
- 毕业级架构至少要为 config、tool、protocol、workflow、eval 和 governance 留出稳定边界。
- 这一周的工作是把零散能力压缩成可以 defend 的平台故事。

## 本周 Atlas 观察路径

- `[Week 46 英文原文](../weeks/week-46.md)`：对照原始英文任务拆分，确认这一周的 deliverables、source reading 和 daily links。
- `[Atlas 主线 Issue（Month 11）](../atlas/issues/month-11/atlas-m11-03.md)`：把本周主题放回 Atlas 主线问题卡，观察范围、验收项和剩余风险。
- `[Week 46 周报](../reports/weekly/week-46-report.md)`：看本周结果、证据和后续动作是如何被记录的。
- `[docs/reviews/archive/month-11-review.md](../reviews/archive/month-11-review.md)`：看 benchmark 复盘真正追问哪些 trade-off。

## 本周 Atlas 落地检查

- 最小落地动作：围绕“写平台 RFC”推进，让至少一项本周交付标准在仓库里可见。
- 必留证据：platform RFC、subsystem map 和 integration order。
- 最大风险：口头上说平台化，实际上跨子系统仍然靠手工胶水和作者记忆。
- 审查不变量：跨子系统整合不应该依赖隐藏的 manual glue。
- 本周最先要盯住的交付：平台 RFC

## 本周最小演示场景

- 背景：这一周你需要把“平台整合 RFC”从概念推进成可被别人看懂、验证和质疑的工程结果。
- 你至少要能演示：
  - 平台 RFC、子系统图、整合顺序说明
  - “platform boundary”为什么会直接影响 Atlas 的主路径，而不只是局部实现
  - 风险“口头上说平台化，实际上跨子系统仍然靠手工胶水和作者记忆。”是如何被控制、暴露或留证的
- reviewer 大概率会追问：
  - 为什么这周先做“写平台 RFC”，而不是先做别的？
  - 不变量“跨子系统整合不应该依赖隐藏的 manual glue。”如果被破坏，会先造成什么后果？
  - 你打算拿什么证据证明这周成果成立？提示：platform RFC、subsystem map 和 integration order。

## 配套案例

- 本周最接近的案例：读 [专题案例包](casebook.md) 中的“案例 11：你拿到一堆性能数字，但一个能信的结论都没有”。
- 本周最该对照的失败模式：读 [失败场景图谱](failure-patterns.md) 中的“模式三：质量证据失真型故障”。
- 推荐顺序：先读本周知识清单，再读案例，再回本周 Atlas 观察路径和周报，检查自己能不能解释坏信号与最小证据。

## 本周工作顺序

- [`RFC`](artifact-playbook.md#artifact-rfc)：先把“写平台 RFC”对应的问题范围、目标、约束和非目标写清楚。
- [`Interface`](artifact-playbook.md#artifact-interface)：围绕“platform boundary”定义输入、输出、owner、错误语义和边界。
- [`Implementation`](artifact-playbook.md#artifact-implementation)：把最小主路径打通，至少让“平台 RFC、子系统图、整合顺序说明”里的一项开始可验证。
- [`Unit Test`](artifact-playbook.md#artifact-unit-test)：优先覆盖最容易击中风险“口头上说平台化，实际上跨子系统仍然靠手工胶水和作者记忆。”的失败路径、边界条件或 contract。
- [`Review`](artifact-playbook.md#artifact-review)：站在 reviewer 视角检查不变量“跨子系统整合不应该依赖隐藏的 manual glue。”是否真的被系统保护，而不是只写在文档里。
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

1. 如果要把“平台整合 RFC”做成 Atlas 可交付能力，你会如何设计主路径和边界？
2. 本周最值得优先验证的失败模式是什么？你会如何收集证据？
3. 这周的 RFC、实现、验证或复盘里，哪一项最能体现你的工程判断？为什么？
4. 如果时间被砍掉一半，你会保留什么，推迟什么？
5. 本周产物如何自然承接到下周，而不是变成一次性文档或一次性代码？

## 每日入口

- [Day 316](days-semester-4/day-316.md): 理解问题和边界
- [Day 317](days-semester-4/day-317.md): 边界与契约设计
- [Day 318](days-semester-4/day-318.md): 最小主路径实现
- [Day 319](days-semester-4/day-319.md): 补全真实可用路径
- [Day 320](days-semester-4/day-320.md): 验证与实验
- [Day 321](days-semester-4/day-321.md): 文档与评审
- [Day 322](days-semester-4/day-322.md): 复盘与 memory

## 配套阅读

- [Week 46](../weeks/week-46.md)
