# Week 04：Docker、CI、测试与第一月复盘

最后更新：2026-06-30

## 本周目标

这一周的重点是把“骨架代码”变成“可交付系统”。

你需要补上：

- Docker
- CI
- 基础测试
- 第一月复盘

## 本周学习重点

- 工程卫生为什么从第一月开始就要建立
- 测试不是最后补，而是交付的一部分
- 复盘不是总结心情，而是校准方向

## 本周 Cookbook 做法

### Recipe 1：补 Docker 与运行方式

### Recipe 2：补 CI 与基础检查

### Recipe 3：补测试与第一月复盘

## 本周交付标准

到本周结束时，你至少要能展示：

- Docker 如何跑
- CI 如何过
- 第一月有什么完成了，什么没完成

## 本周能力焦点

- 结果交付（`Deliver Results`）
- 结构与清晰度（`Structure and Clarity`）
- 代码健康（`Code Health`）

## 本周知识清单

- Dockerfile 的作用是证明本地运行路径和交付运行路径一致，而不是为了“有个容器文件”。
- CI 应该在第一月就承担 build、test、lint 或最小验证门禁，越晚补越难建立纪律。
- 单元测试要围绕前几周建立的边界和不变量写，而不是随机挑函数凑覆盖率。
- 月度复盘要比较原始设计、实际实现和剩余债务，给第二个月提供明确输入。

## 本周 Atlas 观察路径

- `[Week 04 英文原文](../weeks/week-04.md)`：对照原始英文任务拆分，确认这一周的 deliverables、source reading 和 daily links。
- `[Atlas 主线 Issue（Month 01）](../atlas/issues/month-01/atlas-m01-04.md)`：把本周主题放回 Atlas 主线问题卡，观察范围、验收项和剩余风险。
- `[Week 04 周报](../reports/weekly/week-04-report.md)`：看本周结果、证据和后续动作是如何被记录的。
- `[internal/bootstrap/bootstrap_test.go](../../internal/bootstrap/bootstrap_test.go)` 与 `[internal/logger/logger_test.go](../../internal/logger/logger_test.go)`：看第一月边界如何被测试保护。

## 本周 Atlas 落地检查

- 最小落地动作：围绕“补 Docker 与运行方式”推进，让至少一项本周交付标准在仓库里可见。
- 必留证据：Docker 运行路径、CI 记录、测试结果和第一月复盘结论。
- 最大风险：把工程卫生当附属品，导致每次演示都依赖作者本人的机器状态。
- 审查不变量：任何被视为“完成”的能力都必须可以在容器或 CI 环境里复现。
- 本周最先要盯住的交付：Docker 如何跑

## 本周最小演示场景

- 背景：这一周你需要把“Docker、CI、测试与第一月复盘”从概念推进成可被别人看懂、验证和质疑的工程结果。
- 你至少要能演示：
  - Docker 如何跑、CI 如何过、第一月有什么完成了，什么没完成
  - “工程卫生为什么从第一月开始就要建立”为什么会直接影响 Atlas 的主路径，而不只是局部实现
  - 风险“把工程卫生当附属品，导致每次演示都依赖作者本人的机器状态。”是如何被控制、暴露或留证的
- reviewer 大概率会追问：
  - 为什么这周先做“补 Docker 与运行方式”，而不是先做别的？
  - 不变量“任何被视为“完成”的能力都必须可以在容器或 CI 环境里复现。”如果被破坏，会先造成什么后果？
  - 你打算拿什么证据证明这周成果成立？提示：Docker 运行路径、CI 记录、测试结果和第一月复盘结论。

## 配套案例

- 本周最接近的案例：读 [专题案例包](casebook.md) 中的“案例 1：服务能启动，但谁都不敢改”。
- 本周最该对照的失败模式：读 [失败场景图谱](failure-patterns.md) 中的“模式一：边界缺失型故障”。
- 推荐顺序：先读本周知识清单，再读案例，再回本周 Atlas 观察路径和周报，检查自己能不能解释坏信号与最小证据。

## 本周工作顺序

- [`RFC`](artifact-playbook.md#artifact-rfc)：先把“补 Docker 与运行方式”对应的问题范围、目标、约束和非目标写清楚。
- [`Interface`](artifact-playbook.md#artifact-interface)：围绕“工程卫生为什么从第一月开始就要建立”定义输入、输出、owner、错误语义和边界。
- [`Implementation`](artifact-playbook.md#artifact-implementation)：把最小主路径打通，至少让“Docker 如何跑、CI 如何过、第一月有什么完成了，什么没完成”里的一项开始可验证。
- [`Unit Test`](artifact-playbook.md#artifact-unit-test)：优先覆盖最容易击中风险“把工程卫生当附属品，导致每次演示都依赖作者本人的机器状态。”的失败路径、边界条件或 contract。
- [`Review`](artifact-playbook.md#artifact-review)：站在 reviewer 视角检查不变量“任何被视为“完成”的能力都必须可以在容器或 CI 环境里复现。”是否真的被系统保护，而不是只写在文档里。
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

1. 如果要把“Docker、CI、测试与第一月复盘”做成 Atlas 可交付能力，你会如何设计主路径和边界？
2. 本周最值得优先验证的失败模式是什么？你会如何收集证据？
3. 这周的 RFC、实现、验证或复盘里，哪一项最能体现你的工程判断？为什么？
4. 如果时间被砍掉一半，你会保留什么，推迟什么？
5. 本周产物如何自然承接到下周，而不是变成一次性文档或一次性代码？

## 每日入口

- [Day 022](days-month-01/day-022.md): 理解问题和边界
- [Day 023](days-month-01/day-023.md): 边界与契约设计
- [Day 024](days-month-01/day-024.md): 最小主路径实现
- [Day 025](days-month-01/day-025.md): 补全真实可用路径
- [Day 026](days-month-01/day-026.md): 验证与实验
- [Day 027](days-month-01/day-027.md): 文档与评审
- [Day 028](days-month-01/day-028.md): 复盘与 memory

## 配套阅读

- [Week 04](../weeks/week-04.md)
- [Month 01 Scorecard](../reports/monthly/month-01-scorecard.md)
