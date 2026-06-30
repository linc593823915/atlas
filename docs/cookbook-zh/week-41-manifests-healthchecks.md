# Week 41：Manifests、健康检查与资源控制

最后更新：2026-06-30

## 本周目标

本周把部署设计落到实际部署单元。重点是 manifests、probes 和资源约束。

## 本周学完后你应该会什么

你要让 Atlas 从“能跑”进化到“能被平台托管”。

## 本周学习重点

- manifests
- probes
- resource limits

## 本周 Cookbook 做法

### Recipe 1：写 manifests / Helm

### Recipe 2：补 probes

### Recipe 3：补资源边界

## 本周交付标准

- 部署清单
- 健康检查配置
- 资源限制说明

## 本周能力焦点

- 生产所有权（`Production Ownership`）
- SLO 思维（`SLO Thinking`）
- 回滚准备度（`Rollback Readiness`）

## 本周知识清单

- manifest 或 Helm chart 是可执行架构，命名、模板和默认值本身就在暴露系统结构。
- readiness、liveness 和 startup probe 分别回答不同问题，不能混成一个“健康检查”。
- resource request/limit 本质上是在写性能假设和失败预算。
- 扩缩容信号要来自 workload 行为，而不是拍脑袋设 replica 数。

## 本周 Atlas 观察路径

- `[Week 41 英文原文](../weeks/week-41.md)`：对照原始英文任务拆分，确认这一周的 deliverables、source reading 和 daily links。
- `[Atlas 主线 Issue（Month 10）](../atlas/issues/month-10/atlas-m10-02.md)`：把本周主题放回 Atlas 主线问题卡，观察范围、验收项和剩余风险。
- `[Week 41 周报](../reports/weekly/week-41-report.md)`：看本周结果、证据和后续动作是如何被记录的。
- `[docs/operations/review-calendar.md](../operations/review-calendar.md)`：对照 rollout、alert、rollback 相关的节奏安排。

## 本周 Atlas 落地检查

- 最小落地动作：围绕“写 manifests / Helm”推进，让至少一项本周交付标准在仓库里可见。
- 必留证据：manifests、probe 设计和 resource-control 说明。
- 最大风险：复制一套 probe 和 resource 配置，却不知道它们真正会在什么故障下触发。
- 审查不变量：只有当声明依赖真正 ready 时，实例才能开始接流量。
- 本周最先要盯住的交付：部署清单

## 本周最小演示场景

- 背景：这一周你需要把“Manifests、健康检查与资源控制”从概念推进成可被别人看懂、验证和质疑的工程结果。
- 你至少要能演示：
  - 部署清单、健康检查配置、资源限制说明
  - “manifests”为什么会直接影响 Atlas 的主路径，而不只是局部实现
  - 风险“复制一套 probe 和 resource 配置，却不知道它们真正会在什么故障下触发。”是如何被控制、暴露或留证的
- reviewer 大概率会追问：
  - 为什么这周先做“写 manifests / Helm”，而不是先做别的？
  - 不变量“只有当声明依赖真正 ready 时，实例才能开始接流量。”如果被破坏，会先造成什么后果？
  - 你打算拿什么证据证明这周成果成立？提示：manifests、probe 设计和 resource-control 说明。

## 配套案例

- 本周最接近的案例：读 [专题案例包](casebook.md) 中的“案例 10：Kubernetes 部署了，但一点都不像生产系统”。
- 本周最该对照的失败模式：读 [失败场景图谱](failure-patterns.md) 中的“模式五：平台错觉型故障”。
- 推荐顺序：先读本周知识清单，再读案例，再回本周 Atlas 观察路径和周报，检查自己能不能解释坏信号与最小证据。

## 本周工作顺序

- [`RFC`](artifact-playbook.md#artifact-rfc)：先把“写 manifests / Helm”对应的问题范围、目标、约束和非目标写清楚。
- [`Interface`](artifact-playbook.md#artifact-interface)：围绕“manifests”定义输入、输出、owner、错误语义和边界。
- [`Implementation`](artifact-playbook.md#artifact-implementation)：把最小主路径打通，至少让“部署清单、健康检查配置、资源限制说明”里的一项开始可验证。
- [`Unit Test`](artifact-playbook.md#artifact-unit-test)：优先覆盖最容易击中风险“复制一套 probe 和 resource 配置，却不知道它们真正会在什么故障下触发。”的失败路径、边界条件或 contract。
- [`Review`](artifact-playbook.md#artifact-review)：站在 reviewer 视角检查不变量“只有当声明依赖真正 ready 时，实例才能开始接流量。”是否真的被系统保护，而不是只写在文档里。
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

1. 如果要把“Manifests、健康检查与资源控制”做成 Atlas 可交付能力，你会如何设计主路径和边界？
2. 本周最值得优先验证的失败模式是什么？你会如何收集证据？
3. 这周的 RFC、实现、验证或复盘里，哪一项最能体现你的工程判断？为什么？
4. 如果时间被砍掉一半，你会保留什么，推迟什么？
5. 本周产物如何自然承接到下周，而不是变成一次性文档或一次性代码？

## 每日入口

- [Day 281](days-semester-4/day-281.md): 理解问题和边界
- [Day 282](days-semester-4/day-282.md): 边界与契约设计
- [Day 283](days-semester-4/day-283.md): 最小主路径实现
- [Day 284](days-semester-4/day-284.md): 补全真实可用路径
- [Day 285](days-semester-4/day-285.md): 验证与实验
- [Day 286](days-semester-4/day-286.md): 文档与评审
- [Day 287](days-semester-4/day-287.md): 复盘与 memory

## 配套阅读

- [Week 41](../weeks/week-41.md)
