# Week 47：配置中心、Tool Gateway 与 MCP 整合

最后更新：2026-06-30

## 本周目标

本周开始做真正的平台整合。

## 本周学完后你应该会什么

重点是让系统不再是松散模块，而是一个可贯通平台。

## 本周学习重点

- config integration
- tool gateway integration
- MCP integration

## 本周 Cookbook 做法

### Recipe 1：打通 config / tool / MCP

### Recipe 2：整理依赖顺序

### Recipe 3：验证主链路

## 本周交付标准

- 整合后的主链路
- 依赖顺序说明
- 平台骨架

## 本周能力焦点

- 平台整合（`Platform Integration`）
- 架构判断（`Architecture Judgment`）
- 技术领导力（`Technical Leadership`）

## 本周知识清单

- config center、tool gateway 和 MCP integration 应该复用同一套 capability contract，而不是各自为政。
- 配置传播要保证内部 runtime、外部 surface 和 tool 执行路径看到的是同一套 truth。
- MCP 暴露出的能力应与 gateway 强制执行的 schema、timeout 和 policy 保持一致。
- 集成失败最常见的表现是命名不一致、schema drift 和 timeout 语义不一致。

## 本周 Atlas 观察路径

- `[Week 47 英文原文](../weeks/week-47.md)`：对照原始英文任务拆分，确认这一周的 deliverables、source reading 和 daily links。
- `[Atlas 主线 Issue（Month 12）](../atlas/issues/month-12/atlas-m12-01.md)`：把本周主题放回 Atlas 主线问题卡，观察范围、验收项和剩余风险。
- `[Week 47 周报](../reports/weekly/week-47-report.md)`：看本周结果、证据和后续动作是如何被记录的。
- `[docs/governance/exec-dashboard-pack.md](../governance/exec-dashboard-pack.md)`：观察平台级状态如何被组织给高层/评审看。

## 本周 Atlas 落地检查

- 最小落地动作：围绕“打通 config / tool / MCP”推进，让至少一项本周交付标准在仓库里可见。
- 必留证据：integration design、shared contract map 和 consistency check。
- 最大风险：内部有一套 tool truth，外部 MCP 又有一套 tool truth，长期必然分叉。
- 审查不变量：同一个 capability 在 config、gateway 和 MCP 上只能有一份 canonical contract。
- 本周最先要盯住的交付：整合后的主链路

## 本周最小演示场景

- 背景：这一周你需要把“配置中心、Tool Gateway 与 MCP 整合”从概念推进成可被别人看懂、验证和质疑的工程结果。
- 你至少要能演示：
  - 整合后的主链路、依赖顺序说明、平台骨架
  - “config integration”为什么会直接影响 Atlas 的主路径，而不只是局部实现
  - 风险“内部有一套 tool truth，外部 MCP 又有一套 tool truth，长期必然分叉。”是如何被控制、暴露或留证的
- reviewer 大概率会追问：
  - 为什么这周先做“打通 config / tool / MCP”，而不是先做别的？
  - 不变量“同一个 capability 在 config、gateway 和 MCP 上只能有一份 canonical contract。”如果被破坏，会先造成什么后果？
  - 你打算拿什么证据证明这周成果成立？提示：integration design、shared contract map 和 consistency check。

## 配套案例

- 本周最接近的案例：读 [专题案例包](casebook.md) 中的“案例 12：功能都在，为什么还是说不清这是个平台”。
- 本周最该对照的失败模式：读 [失败场景图谱](failure-patterns.md) 中的“模式五：平台错觉型故障”。
- 推荐顺序：先读本周知识清单，再读案例，再回本周 Atlas 观察路径和周报，检查自己能不能解释坏信号与最小证据。

## 本周工作顺序

- [`RFC`](artifact-playbook.md#artifact-rfc)：先把“打通 config / tool / MCP”对应的问题范围、目标、约束和非目标写清楚。
- [`Interface`](artifact-playbook.md#artifact-interface)：围绕“config integration”定义输入、输出、owner、错误语义和边界。
- [`Implementation`](artifact-playbook.md#artifact-implementation)：把最小主路径打通，至少让“整合后的主链路、依赖顺序说明、平台骨架”里的一项开始可验证。
- [`Unit Test`](artifact-playbook.md#artifact-unit-test)：优先覆盖最容易击中风险“内部有一套 tool truth，外部 MCP 又有一套 tool truth，长期必然分叉。”的失败路径、边界条件或 contract。
- [`Review`](artifact-playbook.md#artifact-review)：站在 reviewer 视角检查不变量“同一个 capability 在 config、gateway 和 MCP 上只能有一份 canonical contract。”是否真的被系统保护，而不是只写在文档里。
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

1. 如果要把“配置中心、Tool Gateway 与 MCP 整合”做成 Atlas 可交付能力，你会如何设计主路径和边界？
2. 本周最值得优先验证的失败模式是什么？你会如何收集证据？
3. 这周的 RFC、实现、验证或复盘里，哪一项最能体现你的工程判断？为什么？
4. 如果时间被砍掉一半，你会保留什么，推迟什么？
5. 本周产物如何自然承接到下周，而不是变成一次性文档或一次性代码？

## 每日入口

- [Day 323](days-semester-4/day-323.md): 理解问题和边界
- [Day 324](days-semester-4/day-324.md): 边界与契约设计
- [Day 325](days-semester-4/day-325.md): 最小主路径实现
- [Day 326](days-semester-4/day-326.md): 补全真实可用路径
- [Day 327](days-semester-4/day-327.md): 验证与实验
- [Day 328](days-semester-4/day-328.md): 文档与评审
- [Day 329](days-semester-4/day-329.md): 复盘与 memory

## 配套阅读

- [Week 47](../weeks/week-47.md)
