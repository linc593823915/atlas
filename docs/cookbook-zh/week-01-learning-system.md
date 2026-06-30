# Week 01：学习系统搭建与服务骨架 RFC

最后更新：2026-06-30

## 本周目标

这一周不是写很多代码，而是建立“正确的学习和工程起点”。

你要做两件事：

1. 建立全年学习的工作方式
2. 为 Atlas 写出第一份服务骨架 RFC

## 本周学完后你应该会什么

你应该能解释：

- 为什么架构学习一开始就要写 RFC
- 为什么学习系统本身也需要结构
- 为什么服务骨架不是“以后再整理”的问题

## 本周学习重点

- 学会把问题先写清楚，再开工
- 明确 Atlas 的最小模块边界
- 建立学习日志、Issue 清单、架构笔记

## 本周 Cookbook 做法

### Recipe 1：建立学习日志

产物：

- 每日记录方式
- 每周复盘入口

### Recipe 2：写服务骨架 RFC

产物：

- 包结构草稿
- 生命周期责任划分

## 本周交付标准

到本周结束时，你至少要有：

- 一份 RFC 草稿
- 一份学习日志入口
- 一份包结构初稿

## 本周能力焦点

- 结果交付（`Deliver Results`）
- 结构与清晰度（`Structure and Clarity`）
- 代码健康（`Code Health`）

## 本周知识清单

- RFC 的价值不是写文档，而是先冻结服务骨架的责任边界，避免 `main`、CLI、runtime 和业务初始化互相渗透。
- 学习系统本身也要有证据链：日志、Issue、架构笔记和复盘记录，后续每个月都要依赖这些资产。
- 最小模块边界至少要分清入口层、bootstrap 层、runtime 层和基础设施层，否则后续 Agent 能力会直接长在入口代码上。
- 第一周真正的产物不是代码量，而是后续 11 个月都能复用的工作方式和问题定义模板。

## 本周 Atlas 观察路径

- `[Week 01 英文原文](../weeks/week-01.md)`：对照原始英文任务拆分，确认这一周的 deliverables、source reading 和 daily links。
- `[Atlas 主线 Issue（Month 01）](../atlas/issues/month-01/atlas-m01-01.md)`：把本周主题放回 Atlas 主线问题卡，观察范围、验收项和剩余风险。
- `[Week 01 周报](../reports/weekly/week-01-report.md)`：看本周结果、证据和后续动作是如何被记录的。
- `[cmd/atlas/main.go](../../cmd/atlas/main.go)` 与 `[internal/bootstrap/bootstrap.go](../../internal/bootstrap/bootstrap.go)`：观察进程入口和 bootstrap 如何分工。

## 本周 Atlas 落地检查

- 最小落地动作：围绕“建立学习日志”推进，让至少一项本周交付标准在仓库里可见。
- 必留证据：至少留下 RFC 草稿、包结构图和学习记录入口。
- 最大风险：把服务骨架理解成目录命名练习，忽略生命周期和依赖 ownership。
- 审查不变量：任何新入口都应该复用同一套 bootstrap / shutdown 流程。
- 本周最先要盯住的交付：一份 RFC 草稿

## 本周最小演示场景

- 背景：这一周你需要把“学习系统搭建与服务骨架 RFC”从概念推进成可被别人看懂、验证和质疑的工程结果。
- 你至少要能演示：
  - 一份 RFC 草稿、一份学习日志入口、一份包结构初稿
  - “学会把问题先写清楚，再开工”为什么会直接影响 Atlas 的主路径，而不只是局部实现
  - 风险“把服务骨架理解成目录命名练习，忽略生命周期和依赖 ownership。”是如何被控制、暴露或留证的
- reviewer 大概率会追问：
  - 为什么这周先做“建立学习日志”，而不是先做别的？
  - 不变量“任何新入口都应该复用同一套 bootstrap / shutdown 流程。”如果被破坏，会先造成什么后果？
  - 你打算拿什么证据证明这周成果成立？提示：至少留下 RFC 草稿、包结构图和学习记录入口。

## 配套案例

- 本周最接近的案例：读 [专题案例包](casebook.md) 中的“案例 1：服务能启动，但谁都不敢改”。
- 本周最该对照的失败模式：读 [失败场景图谱](failure-patterns.md) 中的“模式一：边界缺失型故障”。
- 推荐顺序：先读本周知识清单，再读案例，再回本周 Atlas 观察路径和周报，检查自己能不能解释坏信号与最小证据。

## 本周工作顺序

- [`RFC`](artifact-playbook.md#artifact-rfc)：先把“建立学习日志”对应的问题范围、目标、约束和非目标写清楚。
- [`Interface`](artifact-playbook.md#artifact-interface)：围绕“学会把问题先写清楚，再开工”定义输入、输出、owner、错误语义和边界。
- [`Implementation`](artifact-playbook.md#artifact-implementation)：把最小主路径打通，至少让“一份 RFC 草稿、一份学习日志入口、一份包结构初稿”里的一项开始可验证。
- [`Unit Test`](artifact-playbook.md#artifact-unit-test)：优先覆盖最容易击中风险“把服务骨架理解成目录命名练习，忽略生命周期和依赖 ownership。”的失败路径、边界条件或 contract。
- [`Review`](artifact-playbook.md#artifact-review)：站在 reviewer 视角检查不变量“任何新入口都应该复用同一套 bootstrap / shutdown 流程。”是否真的被系统保护，而不是只写在文档里。
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

1. 如果要把“学习系统搭建与服务骨架 RFC”做成 Atlas 可交付能力，你会如何设计主路径和边界？
2. 本周最值得优先验证的失败模式是什么？你会如何收集证据？
3. 这周的 RFC、实现、验证或复盘里，哪一项最能体现你的工程判断？为什么？
4. 如果时间被砍掉一半，你会保留什么，推迟什么？
5. 本周产物如何自然承接到下周，而不是变成一次性文档或一次性代码？

## 每日入口

- [Day 001](days-month-01/day-001.md): 理解问题和边界
- [Day 002](days-month-01/day-002.md): 边界与契约设计
- [Day 003](days-month-01/day-003.md): 最小主路径实现
- [Day 004](days-month-01/day-004.md): 补全真实可用路径
- [Day 005](days-month-01/day-005.md): 验证与实验
- [Day 006](days-month-01/day-006.md): 文档与评审
- [Day 007](days-month-01/day-007.md): 复盘与 memory

## 配套阅读

- [Week 01](../weeks/week-01.md)
- [Day 001](../days/day-001.md)
