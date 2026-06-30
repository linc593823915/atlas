# Week 26：第二学期复盘与第三学期准备

最后更新：2026-06-30

## 本周目标

这一周要总结第二学期从“功能”到“边界”的升级。

## 本周学完后你应该会什么

并为第三学期的状态、评估、治理复杂度做准备。

## 本周学习重点

- 第二学期成果
- 遗留风险
- 第三学期准备点

## 本周 Cookbook 做法

### Recipe 1：做第二学期复盘

### Recipe 2：列技术债

### Recipe 3：写第三学期准备说明

## 本周交付标准

- 第二学期复盘
- 风险清单
- 第三学期准备草稿

## 本周能力焦点

- 编排能力（`Orchestration`）
- 以接口体现技术领导力（`Leadership Through Interfaces`）
- 审批纪律（`Approval Discipline`）

## 本周知识清单

- 第二学期复盘要回答接口有没有更清楚，而不是数量是不是更多。
- 剩余风险通常藏在 shared schema、auth 语义和 runtime ownership 这些不易暴露的地方。
- 第三学期准备要明确哪些 state、eval 和 governance 问题还没有真正被设计。
- 这一学期的 memory 要保留下协议、接口和 runtime 方面可复用的结论。

## 本周 Atlas 观察路径

- `[Week 26 英文原文](../weeks/week-26.md)`：对照原始英文任务拆分，确认这一周的 deliverables、source reading 和 daily links。
- `[Atlas 主线 Issue（Month 06）](../atlas/issues/month-06/atlas-m06-05.md)`：把本周主题放回 Atlas 主线问题卡，观察范围、验收项和剩余风险。
- `[Week 26 周报](../reports/weekly/week-26-report.md)`：看本周结果、证据和后续动作是如何被记录的。
- `[docs/reports/monthly/month-06-scorecard.md](../reports/monthly/month-06-scorecard.md)`：观察多 Agent 月度验收口径。

## 本周 Atlas 落地检查

- 最小落地动作：围绕“做第二学期复盘”推进，让至少一项本周交付标准在仓库里可见。
- 必留证据：第二学期复盘、风险列表和第三学期准备草稿。
- 最大风险：把接口债带进第三学期，后面状态和治理层一叠上来就全面失控。
- 审查不变量：每个对外 contract 都应该有明确 owner 和 test story。
- 本周最先要盯住的交付：第二学期复盘

## 本周最小演示场景

- 背景：这一周你需要把“第二学期复盘与第三学期准备”从概念推进成可被别人看懂、验证和质疑的工程结果。
- 你至少要能演示：
  - 第二学期复盘、风险清单、第三学期准备草稿
  - “第二学期成果”为什么会直接影响 Atlas 的主路径，而不只是局部实现
  - 风险“把接口债带进第三学期，后面状态和治理层一叠上来就全面失控。”是如何被控制、暴露或留证的
- reviewer 大概率会追问：
  - 为什么这周先做“做第二学期复盘”，而不是先做别的？
  - 不变量“每个对外 contract 都应该有明确 owner 和 test story。”如果被破坏，会先造成什么后果？
  - 你打算拿什么证据证明这周成果成立？提示：第二学期复盘、风险列表和第三学期准备草稿。

## 配套案例

- 本周最接近的案例：读 [专题案例包](casebook.md) 中的“案例 6：多 Agent 看起来更强，实际上没人知道谁负责”。
- 本周最该对照的失败模式：读 [失败场景图谱](failure-patterns.md) 中的“模式四：治理失守型故障”。
- 推荐顺序：先读本周知识清单，再读案例，再回本周 Atlas 观察路径和周报，检查自己能不能解释坏信号与最小证据。

## 本周工作顺序

- [`RFC`](artifact-playbook.md#artifact-rfc)：先把“做第二学期复盘”对应的问题范围、目标、约束和非目标写清楚。
- [`Interface`](artifact-playbook.md#artifact-interface)：围绕“第二学期成果”定义输入、输出、owner、错误语义和边界。
- [`Implementation`](artifact-playbook.md#artifact-implementation)：把最小主路径打通，至少让“第二学期复盘、风险清单、第三学期准备草稿”里的一项开始可验证。
- [`Unit Test`](artifact-playbook.md#artifact-unit-test)：优先覆盖最容易击中风险“把接口债带进第三学期，后面状态和治理层一叠上来就全面失控。”的失败路径、边界条件或 contract。
- [`Review`](artifact-playbook.md#artifact-review)：站在 reviewer 视角检查不变量“每个对外 contract 都应该有明确 owner 和 test story。”是否真的被系统保护，而不是只写在文档里。
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

1. 如果要把“第二学期复盘与第三学期准备”做成 Atlas 可交付能力，你会如何设计主路径和边界？
2. 本周最值得优先验证的失败模式是什么？你会如何收集证据？
3. 这周的 RFC、实现、验证或复盘里，哪一项最能体现你的工程判断？为什么？
4. 如果时间被砍掉一半，你会保留什么，推迟什么？
5. 本周产物如何自然承接到下周，而不是变成一次性文档或一次性代码？

## 每日入口

- [Day 176](days-semester-2/day-176.md): 理解问题和边界
- [Day 177](days-semester-2/day-177.md): 边界与契约设计
- [Day 178](days-semester-2/day-178.md): 最小主路径实现
- [Day 179](days-semester-2/day-179.md): 补全真实可用路径
- [Day 180](days-semester-2/day-180.md): 验证与实验
- [Day 181](days-semester-2/day-181.md): 文档与评审
- [Day 182](days-semester-2/day-182.md): 复盘与 memory

## 配套阅读

- [Week 26](../weeks/week-26.md)
