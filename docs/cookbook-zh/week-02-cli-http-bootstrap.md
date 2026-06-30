# Week 02：CLI、HTTP 启动与服务外壳

最后更新：2026-06-30

## 本周目标

把 Atlas 的最小运行外壳搭起来。

这意味着：

- 有 CLI 入口
- 有 HTTP 健康接口
- 有明确的启动/停止路径

## 本周学习重点

- `main` 不应该承担太多责任
- CLI 与运行时入口要清楚分离
- 健康检查是最小系统可运行性的证明

## 本周 Cookbook 做法

### Recipe 1：建立 CLI 主入口

### Recipe 2：建立 HTTP health endpoint

### Recipe 3：把启动和停止收口到统一外壳

## 本周交付标准

到本周结束时，你至少要能演示：

- 命令如何启动 Atlas
- 健康接口如何返回
- 服务如何退出

## 本周能力焦点

- 结果交付（`Deliver Results`）
- 结构与清晰度（`Structure and Clarity`）
- 代码健康（`Code Health`）

## 本周知识清单

- `main` 只负责参数与进程退出码，真正的启动/停止逻辑必须下沉到可测试的 runtime 或 bootstrap 层。
- CLI 命令和 HTTP server 虽然是两种入口，但生命周期、配置加载和依赖初始化应该共用同一套骨架。
- 健康检查的意义不是返回 200，而是明确系统对外承诺了哪些最小可运行性信号。
- 服务外壳阶段就要想清楚 start、serve、shutdown 三段式控制流，后面才能安全加任务和 Agent 能力。

## 本周 Atlas 观察路径

- `[Week 02 英文原文](../weeks/week-02.md)`：对照原始英文任务拆分，确认这一周的 deliverables、source reading 和 daily links。
- `[Atlas 主线 Issue（Month 01）](../atlas/issues/month-01/atlas-m01-02.md)`：把本周主题放回 Atlas 主线问题卡，观察范围、验收项和剩余风险。
- `[Week 02 周报](../reports/weekly/week-02-report.md)`：看本周结果、证据和后续动作是如何被记录的。
- `[internal/app/root.go](../../internal/app/root.go)` 与 `[internal/app/serve.go](../../internal/app/serve.go)`：观察 CLI / HTTP 入口如何共用运行时壳。

## 本周 Atlas 落地检查

- 最小落地动作：围绕“建立 CLI 主入口”推进，让至少一项本周交付标准在仓库里可见。
- 必留证据：可运行 CLI、health endpoint 和生命周期说明。
- 最大风险：把 CLI、HTTP 和 bootstrap 各写一套初始化逻辑，后续每加能力都重复接线。
- 审查不变量：无论从哪个入口启动 Atlas，依赖初始化顺序和关闭顺序都必须一致。
- 本周最先要盯住的交付：命令如何启动 Atlas

## 本周最小演示场景

- 背景：这一周你需要把“CLI、HTTP 启动与服务外壳”从概念推进成可被别人看懂、验证和质疑的工程结果。
- 你至少要能演示：
  - 命令如何启动 Atlas、健康接口如何返回、服务如何退出
  - “`main` 不应该承担太多责任”为什么会直接影响 Atlas 的主路径，而不只是局部实现
  - 风险“把 CLI、HTTP 和 bootstrap 各写一套初始化逻辑，后续每加能力都重复接线。”是如何被控制、暴露或留证的
- reviewer 大概率会追问：
  - 为什么这周先做“建立 CLI 主入口”，而不是先做别的？
  - 不变量“无论从哪个入口启动 Atlas，依赖初始化顺序和关闭顺序都必须一致。”如果被破坏，会先造成什么后果？
  - 你打算拿什么证据证明这周成果成立？提示：可运行 CLI、health endpoint 和生命周期说明。

## 配套案例

- 本周最接近的案例：读 [专题案例包](casebook.md) 中的“案例 1：服务能启动，但谁都不敢改”。
- 本周最该对照的失败模式：读 [失败场景图谱](failure-patterns.md) 中的“模式一：边界缺失型故障”。
- 推荐顺序：先读本周知识清单，再读案例，再回本周 Atlas 观察路径和周报，检查自己能不能解释坏信号与最小证据。

## 本周工作顺序

- [`RFC`](artifact-playbook.md#artifact-rfc)：先把“建立 CLI 主入口”对应的问题范围、目标、约束和非目标写清楚。
- [`Interface`](artifact-playbook.md#artifact-interface)：围绕“`main` 不应该承担太多责任”定义输入、输出、owner、错误语义和边界。
- [`Implementation`](artifact-playbook.md#artifact-implementation)：把最小主路径打通，至少让“命令如何启动 Atlas、健康接口如何返回、服务如何退出”里的一项开始可验证。
- [`Unit Test`](artifact-playbook.md#artifact-unit-test)：优先覆盖最容易击中风险“把 CLI、HTTP 和 bootstrap 各写一套初始化逻辑，后续每加能力都重复接线。”的失败路径、边界条件或 contract。
- [`Review`](artifact-playbook.md#artifact-review)：站在 reviewer 视角检查不变量“无论从哪个入口启动 Atlas，依赖初始化顺序和关闭顺序都必须一致。”是否真的被系统保护，而不是只写在文档里。
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

1. 如果要把“CLI、HTTP 启动与服务外壳”做成 Atlas 可交付能力，你会如何设计主路径和边界？
2. 本周最值得优先验证的失败模式是什么？你会如何收集证据？
3. 这周的 RFC、实现、验证或复盘里，哪一项最能体现你的工程判断？为什么？
4. 如果时间被砍掉一半，你会保留什么，推迟什么？
5. 本周产物如何自然承接到下周，而不是变成一次性文档或一次性代码？

## 每日入口

- [Day 008](days-month-01/day-008.md): 理解问题和边界
- [Day 009](days-month-01/day-009.md): 边界与契约设计
- [Day 010](days-month-01/day-010.md): 最小主路径实现
- [Day 011](days-month-01/day-011.md): 补全真实可用路径
- [Day 012](days-month-01/day-012.md): 验证与实验
- [Day 013](days-month-01/day-013.md): 文档与评审
- [Day 014](days-month-01/day-014.md): 复盘与 memory

## 配套阅读

- [Week 02](../weeks/week-02.md)
- [Week 02 Report](../reports/weekly/week-02-report.md)
