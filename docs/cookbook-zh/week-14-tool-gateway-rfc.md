# Week 14：Tool Gateway RFC

最后更新：2026-06-30

## 本周目标

这一周开始真正进入“平台边界设计”。重点是先把 Tool Gateway 的职责、边界和价值定义清楚。

## 本周学完后你应该会什么

本周不是先写很多实现，而是先回答：工具为什么必须统一入口？

## 本周学习重点

- Tool Gateway 的职责边界
- 统一工具入口的价值
- 为什么 schema、timeout、audit 需要一起考虑

## 本周 Cookbook 做法

### Recipe 1：写 Tool Gateway RFC

### Recipe 2：画出 registry 与 execution path

### Recipe 3：定义最小边界与责任分工

## 本周交付标准

- Tool Gateway RFC 草稿
- 工具网关职责图
- 接口边界说明

## 本周能力焦点

- 接口设计（`Interface Design`）
- 代码评审质量（`Code Review Quality`）
- 小步变更纪律（`Small Change Discipline`）

## 本周知识清单

- Tool Gateway 的价值在于把 request envelope、validation、timeout、auth、audit 和 result normalization 放到同一入口。
- gateway boundary 要区分用户意图、工具 contract 和工具实现，不能把所有细节都暴露给调用方。
- 统一入口的意义是让工具行为可治理、可观测、可测试，而不是只少写几行接线代码。
- schema、timeout 和 audit 必须一起考虑，因为它们共同定义了一次 tool call 的运行契约。

## 本周 Atlas 观察路径

- `[Week 14 英文原文](../weeks/week-14.md)`：对照原始英文任务拆分，确认这一周的 deliverables、source reading 和 daily links。
- `[Atlas 主线 Issue（Month 04）](../atlas/issues/month-04/atlas-m04-01.md)`：把本周主题放回 Atlas 主线问题卡，观察范围、验收项和剩余风险。
- `[Week 14 周报](../reports/weekly/week-14-report.md)`：看本周结果、证据和后续动作是如何被记录的。
- `[docs/framework/google-style-capability-model.md](../framework/google-style-capability-model.md)`：观察 capability 是如何被拆成可评审模型的。

## 本周 Atlas 落地检查

- 最小落地动作：围绕“写 Tool Gateway RFC”推进，让至少一项本周交付标准在仓库里可见。
- 必留证据：Tool Gateway RFC、request/result envelope 和 boundary 说明。
- 最大风险：把 gateway 写成一个薄路由器，实际上没有收回任何控制权。
- 审查不变量：每一次 tool invocation 都必须走过同一条可检查、可拦截的 gateway path。
- 本周最先要盯住的交付：Tool Gateway RFC 草稿

## 本周最小演示场景

- 背景：这一周你需要把“Tool Gateway RFC”从概念推进成可被别人看懂、验证和质疑的工程结果。
- 你至少要能演示：
  - Tool Gateway RFC 草稿、工具网关职责图、接口边界说明
  - “Tool Gateway 的职责边界”为什么会直接影响 Atlas 的主路径，而不只是局部实现
  - 风险“把 gateway 写成一个薄路由器，实际上没有收回任何控制权。”是如何被控制、暴露或留证的
- reviewer 大概率会追问：
  - 为什么这周先做“写 Tool Gateway RFC”，而不是先做别的？
  - 不变量“每一次 tool invocation 都必须走过同一条可检查、可拦截的 gateway path。”如果被破坏，会先造成什么后果？
  - 你打算拿什么证据证明这周成果成立？提示：Tool Gateway RFC、request/result envelope 和 boundary 说明。

## 配套案例

- 本周最接近的案例：读 [专题案例包](casebook.md) 中的“案例 4：工具越来越多，但系统越来越难治理”。
- 本周最该对照的失败模式：读 [失败场景图谱](failure-patterns.md) 中的“模式一：边界缺失型故障”。
- 推荐顺序：先读本周知识清单，再读案例，再回本周 Atlas 观察路径和周报，检查自己能不能解释坏信号与最小证据。

## 本周工作顺序

- [`RFC`](artifact-playbook.md#artifact-rfc)：先把“写 Tool Gateway RFC”对应的问题范围、目标、约束和非目标写清楚。
- [`Interface`](artifact-playbook.md#artifact-interface)：围绕“Tool Gateway 的职责边界”定义输入、输出、owner、错误语义和边界。
- [`Implementation`](artifact-playbook.md#artifact-implementation)：把最小主路径打通，至少让“Tool Gateway RFC 草稿、工具网关职责图、接口边界说明”里的一项开始可验证。
- [`Unit Test`](artifact-playbook.md#artifact-unit-test)：优先覆盖最容易击中风险“把 gateway 写成一个薄路由器，实际上没有收回任何控制权。”的失败路径、边界条件或 contract。
- [`Review`](artifact-playbook.md#artifact-review)：站在 reviewer 视角检查不变量“每一次 tool invocation 都必须走过同一条可检查、可拦截的 gateway path。”是否真的被系统保护，而不是只写在文档里。
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

1. 如果要把“Tool Gateway RFC”做成 Atlas 可交付能力，你会如何设计主路径和边界？
2. 本周最值得优先验证的失败模式是什么？你会如何收集证据？
3. 这周的 RFC、实现、验证或复盘里，哪一项最能体现你的工程判断？为什么？
4. 如果时间被砍掉一半，你会保留什么，推迟什么？
5. 本周产物如何自然承接到下周，而不是变成一次性文档或一次性代码？

## 每日入口

- [Day 092](days-semester-2/day-092.md): 理解问题和边界
- [Day 093](days-semester-2/day-093.md): 边界与契约设计
- [Day 094](days-semester-2/day-094.md): 最小主路径实现
- [Day 095](days-semester-2/day-095.md): 补全真实可用路径
- [Day 096](days-semester-2/day-096.md): 验证与实验
- [Day 097](days-semester-2/day-097.md): 文档与评审
- [Day 098](days-semester-2/day-098.md): 复盘与 memory

## 配套阅读

- [Week 14](../weeks/week-14.md)
