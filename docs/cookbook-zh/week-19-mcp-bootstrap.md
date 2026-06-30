# Week 19：MCP 启动与 Transport 模型

最后更新：2026-06-30

## 本周目标

本周把 MCP 从纸面设计推进到最小服务壳。

## 本周学完后你应该会什么

重点不是功能多，而是把服务壳、transport 和启动模型打通。

## 本周学习重点

- MCP bootstrap
- transport model
- 服务启动路径

## 本周 Cookbook 做法

### Recipe 1：建立 MCP server shell

### Recipe 2：打通 transport

### Recipe 3：准备第一个 capability

## 本周交付标准

- 可运行的 MCP 服务壳
- transport 路径
- 最小 capability 入口

## 本周能力焦点

- 协议思维（`Protocol Thinking`）
- 兼容性（`Compatibility`）
- 把文档当产品（`Documentation as Product`）

## 本周知识清单

- MCP server bootstrap 要定义 session 如何启动、关闭以及如何拿到共享依赖。
- transport 选择会改变运行时语义：stdio、socket、HTTP wrapper 的生命周期假设完全不同。
- request dispatch 需要在保留 protocol context 的同时安全接入 Atlas runtime。
- 启动路径要把 capability registration、健康信号和优雅关闭都讲清楚。

## 本周 Atlas 观察路径

- `[Week 19 英文原文](../weeks/week-19.md)`：对照原始英文任务拆分，确认这一周的 deliverables、source reading 和 daily links。
- `[Atlas 主线 Issue（Month 05）](../atlas/issues/month-05/atlas-m05-02.md)`：把本周主题放回 Atlas 主线问题卡，观察范围、验收项和剩余风险。
- `[Week 19 周报](../reports/weekly/week-19-report.md)`：看本周结果、证据和后续动作是如何被记录的。
- `[docs/operations/program-control-center.md](../operations/program-control-center.md)`：观察协议面能力如何进入项目控制视图。

## 本周 Atlas 落地检查

- 最小落地动作：围绕“建立 MCP server shell”推进，让至少一项本周交付标准在仓库里可见。
- 必留证据：bootstrap path、transport model 和 startup sequence。
- 最大风险：把 protocol bootstrap 当成 `main()` 的薄壳，结果 session lifecycle 和 runtime lifecycle 互相打架。
- 审查不变量：session 生命周期不能和依赖生命周期相互冲突。
- 本周最先要盯住的交付：可运行的 MCP 服务壳

## 本周最小演示场景

- 背景：这一周你需要把“MCP 启动与 Transport 模型”从概念推进成可被别人看懂、验证和质疑的工程结果。
- 你至少要能演示：
  - 可运行的 MCP 服务壳、transport 路径、最小 capability 入口
  - “MCP bootstrap”为什么会直接影响 Atlas 的主路径，而不只是局部实现
  - 风险“把 protocol bootstrap 当成 `main()` 的薄壳，结果 session lifecycle 和 runtime lifecycle 互相打架。”是如何被控制、暴露或留证的
- reviewer 大概率会追问：
  - 为什么这周先做“建立 MCP server shell”，而不是先做别的？
  - 不变量“session 生命周期不能和依赖生命周期相互冲突。”如果被破坏，会先造成什么后果？
  - 你打算拿什么证据证明这周成果成立？提示：bootstrap path、transport model 和 startup sequence。

## 配套案例

- 本周最接近的案例：读 [专题案例包](casebook.md) 中的“案例 5：MCP 服务能启动，但根本不是可消费协议”。
- 本周最该对照的失败模式：读 [失败场景图谱](failure-patterns.md) 中的“模式一：边界缺失型故障”。
- 推荐顺序：先读本周知识清单，再读案例，再回本周 Atlas 观察路径和周报，检查自己能不能解释坏信号与最小证据。

## 本周工作顺序

- [`RFC`](artifact-playbook.md#artifact-rfc)：先把“建立 MCP server shell”对应的问题范围、目标、约束和非目标写清楚。
- [`Interface`](artifact-playbook.md#artifact-interface)：围绕“MCP bootstrap”定义输入、输出、owner、错误语义和边界。
- [`Implementation`](artifact-playbook.md#artifact-implementation)：把最小主路径打通，至少让“可运行的 MCP 服务壳、transport 路径、最小 capability 入口”里的一项开始可验证。
- [`Unit Test`](artifact-playbook.md#artifact-unit-test)：优先覆盖最容易击中风险“把 protocol bootstrap 当成 `main()` 的薄壳，结果 session lifecycle 和 runtime lifecycle 互相打架。”的失败路径、边界条件或 contract。
- [`Review`](artifact-playbook.md#artifact-review)：站在 reviewer 视角检查不变量“session 生命周期不能和依赖生命周期相互冲突。”是否真的被系统保护，而不是只写在文档里。
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

1. 如果要把“MCP 启动与 Transport 模型”做成 Atlas 可交付能力，你会如何设计主路径和边界？
2. 本周最值得优先验证的失败模式是什么？你会如何收集证据？
3. 这周的 RFC、实现、验证或复盘里，哪一项最能体现你的工程判断？为什么？
4. 如果时间被砍掉一半，你会保留什么，推迟什么？
5. 本周产物如何自然承接到下周，而不是变成一次性文档或一次性代码？

## 每日入口

- [Day 127](days-semester-2/day-127.md): 理解问题和边界
- [Day 128](days-semester-2/day-128.md): 边界与契约设计
- [Day 129](days-semester-2/day-129.md): 最小主路径实现
- [Day 130](days-semester-2/day-130.md): 补全真实可用路径
- [Day 131](days-semester-2/day-131.md): 验证与实验
- [Day 132](days-semester-2/day-132.md): 文档与评审
- [Day 133](days-semester-2/day-133.md): 复盘与 memory

## 配套阅读

- [Week 19](../weeks/week-19.md)
