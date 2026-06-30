# Week 11：多 Tool 路由与失败处理

最后更新：2026-06-30

## 本周目标

把单 Tool 的 Agent 扩成多个 Tool，并开始处理真实失败。

## 本周学习重点

- 多 Tool 路由
- fallback 思维
- 工具错误如何不拖崩整条链

## 本周 Cookbook 做法

### Recipe 1：扩到至少三个 Tool

### Recipe 2：定义路由条件

### Recipe 3：加入 fallback / failure path

## 本周交付标准

你至少要能说明：

- 为什么路由逻辑合理
- 失败后系统会怎么退回
- 为什么这比“直接报错退出”更成熟

## 本周能力焦点

- 问题求解（`Problem Solving`）
- 面向用户的工具使用（`User-Facing Tool Use`）
- 度量基础（`Measurement Basics`）

## 本周知识清单

- 多 Tool routing 需要清楚的选择策略，而不是注册多个工具后随缘调用。
- fallback 设计要回答何时重试当前 tool、何时切换 tool、何时请求人工或直接失败。
- 一个 tool 失败不应该把整条链拖崩，runtime 必须定义可退化路径。
- routing 逻辑必须可解释，否则 operator 无法理解为什么模型做出某次工具选择。

## 本周 Atlas 观察路径

- `[Week 11 英文原文](../weeks/week-11.md)`：对照原始英文任务拆分，确认这一周的 deliverables、source reading 和 daily links。
- `[Atlas 主线 Issue（Month 03）](../atlas/issues/month-03/atlas-m03-03.md)`：把本周主题放回 Atlas 主线问题卡，观察范围、验收项和剩余风险。
- `[Week 11 周报](../reports/weekly/week-11-report.md)`：看本周结果、证据和后续动作是如何被记录的。
- `[docs/reviews/archive/month-03-review.md](../reviews/archive/month-03-review.md)`：看第三个月 review 最关心哪些问题。

## 本周 Atlas 落地检查

- 最小落地动作：围绕“扩到至少三个 Tool”推进，让至少一项本周交付标准在仓库里可见。
- 必留证据：routing path、fallback case 和 tool negative tests。
- 最大风险：工具越来越多，但没有任何可审计、可解释的选择逻辑。
- 审查不变量：每一次 routing 决策和 tool failure 都必须事后可追溯。
- 本周最先要盯住的交付：为什么路由逻辑合理

## 本周最小演示场景

- 背景：这一周你需要把“多 Tool 路由与失败处理”从概念推进成可被别人看懂、验证和质疑的工程结果。
- 你至少要能演示：
  - 为什么路由逻辑合理、失败后系统会怎么退回、为什么这比“直接报错退出”更成熟
  - “多 Tool 路由”为什么会直接影响 Atlas 的主路径，而不只是局部实现
  - 风险“工具越来越多，但没有任何可审计、可解释的选择逻辑。”是如何被控制、暴露或留证的
- reviewer 大概率会追问：
  - 为什么这周先做“扩到至少三个 Tool”，而不是先做别的？
  - 不变量“每一次 routing 决策和 tool failure 都必须事后可追溯。”如果被破坏，会先造成什么后果？
  - 你打算拿什么证据证明这周成果成立？提示：routing path、fallback case 和 tool negative tests。

## 配套案例

- 本周最接近的案例：读 [专题案例包](casebook.md) 中的“案例 3：做出了会聊天的 Agent，但系统完全不可依赖”。
- 本周最该对照的失败模式：读 [失败场景图谱](failure-patterns.md) 中的“模式三：质量证据失真型故障”。
- 推荐顺序：先读本周知识清单，再读案例，再回本周 Atlas 观察路径和周报，检查自己能不能解释坏信号与最小证据。

## 本周工作顺序

- [`RFC`](artifact-playbook.md#artifact-rfc)：先把“扩到至少三个 Tool”对应的问题范围、目标、约束和非目标写清楚。
- [`Interface`](artifact-playbook.md#artifact-interface)：围绕“多 Tool 路由”定义输入、输出、owner、错误语义和边界。
- [`Implementation`](artifact-playbook.md#artifact-implementation)：把最小主路径打通，至少让“为什么路由逻辑合理、失败后系统会怎么退回、为什么这比“直接报错退出”更成熟”里的一项开始可验证。
- [`Unit Test`](artifact-playbook.md#artifact-unit-test)：优先覆盖最容易击中风险“工具越来越多，但没有任何可审计、可解释的选择逻辑。”的失败路径、边界条件或 contract。
- [`Review`](artifact-playbook.md#artifact-review)：站在 reviewer 视角检查不变量“每一次 routing 决策和 tool failure 都必须事后可追溯。”是否真的被系统保护，而不是只写在文档里。
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

1. 如果要把“多 Tool 路由与失败处理”做成 Atlas 可交付能力，你会如何设计主路径和边界？
2. 本周最值得优先验证的失败模式是什么？你会如何收集证据？
3. 这周的 RFC、实现、验证或复盘里，哪一项最能体现你的工程判断？为什么？
4. 如果时间被砍掉一半，你会保留什么，推迟什么？
5. 本周产物如何自然承接到下周，而不是变成一次性文档或一次性代码？

## 每日入口

- [Day 071](days-semester-1/day-071.md): 理解问题和边界
- [Day 072](days-semester-1/day-072.md): 边界与契约设计
- [Day 073](days-semester-1/day-073.md): 最小主路径实现
- [Day 074](days-semester-1/day-074.md): 补全真实可用路径
- [Day 075](days-semester-1/day-075.md): 验证与实验
- [Day 076](days-semester-1/day-076.md): 文档与评审
- [Day 077](days-semester-1/day-077.md): 复盘与 memory

## 配套阅读

- [Week 11](../weeks/week-11.md)
- [Month 03 Issue Set](../atlas/issues/month-03/README.md)
