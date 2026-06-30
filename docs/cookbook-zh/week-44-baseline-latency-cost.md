# Week 44：基线延迟、吞吐与成本测量

最后更新：2026-06-30

## 本周目标

本周拿到第一版性能和成本基线。

## 本周学完后你应该会什么

你要开始用数字理解 Atlas，而不是只靠主观感觉。

## 本周学习重点

- latency baseline
- throughput baseline
- token cost

## 本周 Cookbook 做法

### Recipe 1：跑 baseline

### Recipe 2：记录成本

### Recipe 3：整理 dashboard 数据

## 本周交付标准

- baseline 数据
- 成本记录
- 初版分析结论

## 本周能力焦点

- 性能工程（`Performance Engineering`）
- 成本意识（`Cost Awareness`）
- 容量规划（`Capacity Planning`）

## 本周知识清单

- baseline 要区分 cold start、warm path、正常负载和压力边界，否则平均值会掩盖真实问题。
- throughput 和 token cost 只能在代表性请求组合上测，不能靠一个 synthetic case 推断整体。
- measurement hygiene 包括固定环境、样本数量、方差意识和结果记录方式。
- baseline 的第一目标是可信，而不是好看。

## 本周 Atlas 观察路径

- `[Week 44 英文原文](../weeks/week-44.md)`：对照原始英文任务拆分，确认这一周的 deliverables、source reading 和 daily links。
- `[Atlas 主线 Issue（Month 11）](../atlas/issues/month-11/atlas-m11-01.md)`：把本周主题放回 Atlas 主线问题卡，观察范围、验收项和剩余风险。
- `[Week 44 周报](../reports/weekly/week-44-report.md)`：看本周结果、证据和后续动作是如何被记录的。
- `[docs/reports/snapshots/month-11-snapshot.md](../reports/snapshots/month-11-snapshot.md)`：先看性能、成本和容量结果如何被摘要。

## 本周 Atlas 落地检查

- 最小落地动作：围绕“跑 baseline”推进，让至少一项本周交付标准在仓库里可见。
- 必留证据：baseline runbook、measurement data 和初始结果。
- 最大风险：基线本身不稳定，后面所有优化都只是在优化噪音。
- 审查不变量：所有 before/after 对比都必须跑在同一份 benchmark contract 上。
- 本周最先要盯住的交付：baseline 数据

## 本周最小演示场景

- 背景：这一周你需要把“基线延迟、吞吐与成本测量”从概念推进成可被别人看懂、验证和质疑的工程结果。
- 你至少要能演示：
  - baseline 数据、成本记录、初版分析结论
  - “latency baseline”为什么会直接影响 Atlas 的主路径，而不只是局部实现
  - 风险“基线本身不稳定，后面所有优化都只是在优化噪音。”是如何被控制、暴露或留证的
- reviewer 大概率会追问：
  - 为什么这周先做“跑 baseline”，而不是先做别的？
  - 不变量“所有 before/after 对比都必须跑在同一份 benchmark contract 上。”如果被破坏，会先造成什么后果？
  - 你打算拿什么证据证明这周成果成立？提示：baseline runbook、measurement data 和初始结果。

## 配套案例

- 本周最接近的案例：读 [专题案例包](casebook.md) 中的“案例 11：你拿到一堆性能数字，但一个能信的结论都没有”。
- 本周最该对照的失败模式：读 [失败场景图谱](failure-patterns.md) 中的“模式三：质量证据失真型故障”。
- 推荐顺序：先读本周知识清单，再读案例，再回本周 Atlas 观察路径和周报，检查自己能不能解释坏信号与最小证据。

## 本周工作顺序

- [`RFC`](artifact-playbook.md#artifact-rfc)：先把“跑 baseline”对应的问题范围、目标、约束和非目标写清楚。
- [`Interface`](artifact-playbook.md#artifact-interface)：围绕“latency baseline”定义输入、输出、owner、错误语义和边界。
- [`Implementation`](artifact-playbook.md#artifact-implementation)：把最小主路径打通，至少让“baseline 数据、成本记录、初版分析结论”里的一项开始可验证。
- [`Unit Test`](artifact-playbook.md#artifact-unit-test)：优先覆盖最容易击中风险“基线本身不稳定，后面所有优化都只是在优化噪音。”的失败路径、边界条件或 contract。
- [`Review`](artifact-playbook.md#artifact-review)：站在 reviewer 视角检查不变量“所有 before/after 对比都必须跑在同一份 benchmark contract 上。”是否真的被系统保护，而不是只写在文档里。
- [`Benchmark（按需）`](artifact-playbook.md#artifact-benchmark)：本周 benchmark 是主题核心，必须保留 baseline 原始数据和统计口径。
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

1. 如果要把“基线延迟、吞吐与成本测量”做成 Atlas 可交付能力，你会如何设计主路径和边界？
2. 本周最值得优先验证的失败模式是什么？你会如何收集证据？
3. 这周的 RFC、实现、验证或复盘里，哪一项最能体现你的工程判断？为什么？
4. 如果时间被砍掉一半，你会保留什么，推迟什么？
5. 本周产物如何自然承接到下周，而不是变成一次性文档或一次性代码？

## 每日入口

- [Day 302](days-semester-4/day-302.md): 理解问题和边界
- [Day 303](days-semester-4/day-303.md): 边界与契约设计
- [Day 304](days-semester-4/day-304.md): 最小主路径实现
- [Day 305](days-semester-4/day-305.md): 补全真实可用路径
- [Day 306](days-semester-4/day-306.md): 验证与实验
- [Day 307](days-semester-4/day-307.md): 文档与评审
- [Day 308](days-semester-4/day-308.md): 复盘与 memory

## 配套阅读

- [Week 44](../weeks/week-44.md)
