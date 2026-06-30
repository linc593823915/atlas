# Week 18：MCP Server RFC

最后更新：2026-06-30

## 本周目标

这一周开始从内部接口走向外部协议面。

## 本周学完后你应该会什么

本周重点是定义 Atlas 为什么以及如何暴露成 MCP Surface。

## 本周学习重点

- MCP 能力边界
- tools/resources 划分
- 协议职责

## 本周 Cookbook 做法

### Recipe 1：写 MCP RFC

### Recipe 2：定义 capability inventory

### Recipe 3：明确 transport 预期

## 本周交付标准

- MCP RFC
- capability inventory
- 协议边界说明

## 本周能力焦点

- 协议思维（`Protocol Thinking`）
- 兼容性（`Compatibility`）
- 把文档当产品（`Documentation as Product`）

## 本周知识清单

- MCP surface 选择是 API product 决策：只暴露稳定、可消费的能力，而不是把内部接口直接倒出去。
- tools 更适合命令式动作，resources 更适合只读状态或上下文，两者混用会让 client 心智混乱。
- MCP RFC 要先定义 session lifecycle、transport 假设和 error shape，再谈实现。
- 第一版 server scope 必须刻意收窄，先学兼容性而不是一次性暴露全部能力。

## 本周 Atlas 观察路径

- `[Week 18 英文原文](../weeks/week-18.md)`：对照原始英文任务拆分，确认这一周的 deliverables、source reading 和 daily links。
- `[Atlas 主线 Issue（Month 05）](../atlas/issues/month-05/atlas-m05-01.md)`：把本周主题放回 Atlas 主线问题卡，观察范围、验收项和剩余风险。
- `[Week 18 周报](../reports/weekly/week-18-report.md)`：看本周结果、证据和后续动作是如何被记录的。
- `[docs/framework/google-public-sources.md](../framework/google-public-sources.md)`：对照 MCP / capability 设计常用公开资料来源。

## 本周 Atlas 落地检查

- 最小落地动作：围绕“写 MCP RFC”推进，让至少一项本周交付标准在仓库里可见。
- 必留证据：MCP RFC、capability inventory 和 surface scope note。
- 最大风险：把内部 helper 行为发布成长期协议 contract，后续无法演进。
- 审查不变量：任何暴露到 MCP 的能力都必须脱离 Atlas 内部细节也能被解释和支持。
- 本周最先要盯住的交付：MCP RFC

## 本周最小演示场景

- 背景：这一周你需要把“MCP Server RFC”从概念推进成可被别人看懂、验证和质疑的工程结果。
- 你至少要能演示：
  - MCP RFC、capability inventory、协议边界说明
  - “MCP 能力边界”为什么会直接影响 Atlas 的主路径，而不只是局部实现
  - 风险“把内部 helper 行为发布成长期协议 contract，后续无法演进。”是如何被控制、暴露或留证的
- reviewer 大概率会追问：
  - 为什么这周先做“写 MCP RFC”，而不是先做别的？
  - 不变量“任何暴露到 MCP 的能力都必须脱离 Atlas 内部细节也能被解释和支持。”如果被破坏，会先造成什么后果？
  - 你打算拿什么证据证明这周成果成立？提示：MCP RFC、capability inventory 和 surface scope note。

## 配套案例

- 本周最接近的案例：读 [专题案例包](casebook.md) 中的“案例 5：MCP 服务能启动，但根本不是可消费协议”。
- 本周最该对照的失败模式：读 [失败场景图谱](failure-patterns.md) 中的“模式一：边界缺失型故障”。
- 推荐顺序：先读本周知识清单，再读案例，再回本周 Atlas 观察路径和周报，检查自己能不能解释坏信号与最小证据。

## 本周工作顺序

- [`RFC`](artifact-playbook.md#artifact-rfc)：先把“写 MCP RFC”对应的问题范围、目标、约束和非目标写清楚。
- [`Interface`](artifact-playbook.md#artifact-interface)：围绕“MCP 能力边界”定义输入、输出、owner、错误语义和边界。
- [`Implementation`](artifact-playbook.md#artifact-implementation)：把最小主路径打通，至少让“MCP RFC、capability inventory、协议边界说明”里的一项开始可验证。
- [`Unit Test`](artifact-playbook.md#artifact-unit-test)：优先覆盖最容易击中风险“把内部 helper 行为发布成长期协议 contract，后续无法演进。”的失败路径、边界条件或 contract。
- [`Review`](artifact-playbook.md#artifact-review)：站在 reviewer 视角检查不变量“任何暴露到 MCP 的能力都必须脱离 Atlas 内部细节也能被解释和支持。”是否真的被系统保护，而不是只写在文档里。
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

1. 如果要把“MCP Server RFC”做成 Atlas 可交付能力，你会如何设计主路径和边界？
2. 本周最值得优先验证的失败模式是什么？你会如何收集证据？
3. 这周的 RFC、实现、验证或复盘里，哪一项最能体现你的工程判断？为什么？
4. 如果时间被砍掉一半，你会保留什么，推迟什么？
5. 本周产物如何自然承接到下周，而不是变成一次性文档或一次性代码？

## 每日入口

- [Day 120](days-semester-2/day-120.md): 理解问题和边界
- [Day 121](days-semester-2/day-121.md): 边界与契约设计
- [Day 122](days-semester-2/day-122.md): 最小主路径实现
- [Day 123](days-semester-2/day-123.md): 补全真实可用路径
- [Day 124](days-semester-2/day-124.md): 验证与实验
- [Day 125](days-semester-2/day-125.md): 文档与评审
- [Day 126](days-semester-2/day-126.md): 复盘与 memory

## 配套阅读

- [Week 18](../weeks/week-18.md)
