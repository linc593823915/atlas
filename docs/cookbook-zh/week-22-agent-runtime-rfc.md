# Week 22：多 Agent Runtime RFC

最后更新：2026-06-30

## 本周目标

这一周从协议面进入运行时。

## 本周学完后你应该会什么

你要先定义多 Agent 的角色、交接和控制边界。

## 本周学习重点

- Agent role map
- Runtime responsibility
- handoff 边界

## 本周 Cookbook 做法

### Recipe 1：写多 Agent RFC

### Recipe 2：定义角色图

### Recipe 3：识别审批点

## 本周交付标准

- 多 Agent RFC
- 角色图
- 审批边界草稿

## 本周能力焦点

- 编排能力（`Orchestration`）
- 以接口体现技术领导力（`Leadership Through Interfaces`）
- 审批纪律（`Approval Discipline`）

## 本周知识清单

- 多 Agent runtime 首先是 role map：谁负责规划、谁负责执行、谁负责 review、谁对最终结果负责。
- runtime responsibility 要把协调逻辑从单个 agent 行为里分出来，否则 orchestration 无法测试。
- handoff boundary 要定义上下文、权限和证据在 agent 之间如何转移。
- RFC 还要回答什么时候根本不该再起一个 agent。

## 本周 Atlas 观察路径

- `[Week 22 英文原文](../weeks/week-22.md)`：对照原始英文任务拆分，确认这一周的 deliverables、source reading 和 daily links。
- `[Atlas 主线 Issue（Month 06）](../atlas/issues/month-06/atlas-m06-01.md)`：把本周主题放回 Atlas 主线问题卡，观察范围、验收项和剩余风险。
- `[Week 22 周报](../reports/weekly/week-22-report.md)`：看本周结果、证据和后续动作是如何被记录的。
- `[docs/governance/role-gap-analysis.md](../governance/role-gap-analysis.md)`：对照多 Agent role map 和责任缺口分析。

## 本周 Atlas 落地检查

- 最小落地动作：围绕“写多 Agent RFC”推进，让至少一项本周交付标准在仓库里可见。
- 必留证据：multi-agent runtime RFC、role map 和 handoff boundary。
- 最大风险：把多 Agent 设计成失控的层层委托，最终没人真正负责结果。
- 审查不变量：每一次 handoff 都必须保留目标 ownership、上下文来源和结果责任。
- 本周最先要盯住的交付：多 Agent RFC

## 本周最小演示场景

- 背景：这一周你需要把“多 Agent Runtime RFC”从概念推进成可被别人看懂、验证和质疑的工程结果。
- 你至少要能演示：
  - 多 Agent RFC、角色图、审批边界草稿
  - “Agent role map”为什么会直接影响 Atlas 的主路径，而不只是局部实现
  - 风险“把多 Agent 设计成失控的层层委托，最终没人真正负责结果。”是如何被控制、暴露或留证的
- reviewer 大概率会追问：
  - 为什么这周先做“写多 Agent RFC”，而不是先做别的？
  - 不变量“每一次 handoff 都必须保留目标 ownership、上下文来源和结果责任。”如果被破坏，会先造成什么后果？
  - 你打算拿什么证据证明这周成果成立？提示：multi-agent runtime RFC、role map 和 handoff boundary。

## 配套案例

- 本周最接近的案例：读 [专题案例包](casebook.md) 中的“案例 6：多 Agent 看起来更强，实际上没人知道谁负责”。
- 本周最该对照的失败模式：读 [失败场景图谱](failure-patterns.md) 中的“模式四：治理失守型故障”。
- 推荐顺序：先读本周知识清单，再读案例，再回本周 Atlas 观察路径和周报，检查自己能不能解释坏信号与最小证据。

## 本周工作顺序

- [`RFC`](artifact-playbook.md#artifact-rfc)：先把“写多 Agent RFC”对应的问题范围、目标、约束和非目标写清楚。
- [`Interface`](artifact-playbook.md#artifact-interface)：围绕“Agent role map”定义输入、输出、owner、错误语义和边界。
- [`Implementation`](artifact-playbook.md#artifact-implementation)：把最小主路径打通，至少让“多 Agent RFC、角色图、审批边界草稿”里的一项开始可验证。
- [`Unit Test`](artifact-playbook.md#artifact-unit-test)：优先覆盖最容易击中风险“把多 Agent 设计成失控的层层委托，最终没人真正负责结果。”的失败路径、边界条件或 contract。
- [`Review`](artifact-playbook.md#artifact-review)：站在 reviewer 视角检查不变量“每一次 handoff 都必须保留目标 ownership、上下文来源和结果责任。”是否真的被系统保护，而不是只写在文档里。
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

1. 如果要把“多 Agent Runtime RFC”做成 Atlas 可交付能力，你会如何设计主路径和边界？
2. 本周最值得优先验证的失败模式是什么？你会如何收集证据？
3. 这周的 RFC、实现、验证或复盘里，哪一项最能体现你的工程判断？为什么？
4. 如果时间被砍掉一半，你会保留什么，推迟什么？
5. 本周产物如何自然承接到下周，而不是变成一次性文档或一次性代码？

## 每日入口

- [Day 148](days-semester-2/day-148.md): 理解问题和边界
- [Day 149](days-semester-2/day-149.md): 边界与契约设计
- [Day 150](days-semester-2/day-150.md): 最小主路径实现
- [Day 151](days-semester-2/day-151.md): 补全真实可用路径
- [Day 152](days-semester-2/day-152.md): 验证与实验
- [Day 153](days-semester-2/day-153.md): 文档与评审
- [Day 154](days-semester-2/day-154.md): 复盘与 memory

## 配套阅读

- [Week 22](../weeks/week-22.md)
