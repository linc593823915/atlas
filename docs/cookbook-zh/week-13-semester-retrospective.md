# Week 13：第一学期复盘与第二学期准备

最后更新：2026-06-30

## 本周目标

这一周不是做新功能，而是做第一学期的系统性结算。

## 本周学习重点

- 哪些能力已经建立
- 哪些问题会影响第二学期
- 第二学期为什么必须转向接口和运行时边界

## 本周 Cookbook 做法

### Recipe 1：盘点第一学期成果

### Recipe 2：列出技术债与盲点

### Recipe 3：写第二学期准备 RFC

## 本周交付标准

你至少要产出：

- 第一学期复盘
- 技术债清单
- 第二学期准备说明

## 本周能力焦点

- 问题求解（`Problem Solving`）
- 面向用户的工具使用（`User-Facing Tool Use`）
- 度量基础（`Measurement Basics`）

## 本周知识清单

- 第一学期复盘要把“做过什么”压缩成“已经具备哪些能力”：骨架、任务执行、单 Agent 主路径。
- 第二学期准备的关键是找出哪些 contract 仍然隐含在实现里，一旦走向接口和协议就会破。
- 技术债要按对接口质量、runtime 清晰度和可靠性的影响排序，而不是按个人好恶排序。
- 证据整理要服务后续学习，保证第二学期开工时还能读懂第一学期的设计判断。

## 本周 Atlas 观察路径

- `[Week 13 英文原文](../weeks/week-13.md)`：对照原始英文任务拆分，确认这一周的 deliverables、source reading 和 daily links。
- `[Atlas 主线 Issue（Month 03）](../atlas/issues/month-03/atlas-m03-05.md)`：把本周主题放回 Atlas 主线问题卡，观察范围、验收项和剩余风险。
- `[Week 13 周报](../reports/weekly/week-13-report.md)`：看本周结果、证据和后续动作是如何被记录的。
- `[docs/governance/quarter-1-end-to-end-sample.md](../governance/quarter-1-end-to-end-sample.md)`：学期复盘周用它对照季度级证据组织。

## 本周 Atlas 落地检查

- 最小落地动作：围绕“盘点第一学期成果”推进，让至少一项本周交付标准在仓库里可见。
- 必留证据：第一学期复盘、技术债清单和第二学期准备说明。
- 最大风险：周周都完成了，但学期级能力和剩余缺口根本没有被抽象出来。
- 审查不变量：任何宣称已经掌握的能力都应该指向代码、文档和验证证据。
- 本周最先要盯住的交付：第一学期复盘

## 本周最小演示场景

- 背景：这一周你需要把“第一学期复盘与第二学期准备”从概念推进成可被别人看懂、验证和质疑的工程结果。
- 你至少要能演示：
  - 第一学期复盘、技术债清单、第二学期准备说明
  - “哪些能力已经建立”为什么会直接影响 Atlas 的主路径，而不只是局部实现
  - 风险“周周都完成了，但学期级能力和剩余缺口根本没有被抽象出来。”是如何被控制、暴露或留证的
- reviewer 大概率会追问：
  - 为什么这周先做“盘点第一学期成果”，而不是先做别的？
  - 不变量“任何宣称已经掌握的能力都应该指向代码、文档和验证证据。”如果被破坏，会先造成什么后果？
  - 你打算拿什么证据证明这周成果成立？提示：第一学期复盘、技术债清单和第二学期准备说明。

## 配套案例

- 本周最接近的案例：读 [专题案例包](casebook.md) 中的“案例 3：做出了会聊天的 Agent，但系统完全不可依赖”。
- 本周最该对照的失败模式：读 [失败场景图谱](failure-patterns.md) 中的“模式三：质量证据失真型故障”。
- 推荐顺序：先读本周知识清单，再读案例，再回本周 Atlas 观察路径和周报，检查自己能不能解释坏信号与最小证据。

## 本周工作顺序

- [`RFC`](artifact-playbook.md#artifact-rfc)：先把“盘点第一学期成果”对应的问题范围、目标、约束和非目标写清楚。
- [`Interface`](artifact-playbook.md#artifact-interface)：围绕“哪些能力已经建立”定义输入、输出、owner、错误语义和边界。
- [`Implementation`](artifact-playbook.md#artifact-implementation)：把最小主路径打通，至少让“第一学期复盘、技术债清单、第二学期准备说明”里的一项开始可验证。
- [`Unit Test`](artifact-playbook.md#artifact-unit-test)：优先覆盖最容易击中风险“周周都完成了，但学期级能力和剩余缺口根本没有被抽象出来。”的失败路径、边界条件或 contract。
- [`Review`](artifact-playbook.md#artifact-review)：站在 reviewer 视角检查不变量“任何宣称已经掌握的能力都应该指向代码、文档和验证证据。”是否真的被系统保护，而不是只写在文档里。
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

1. 如果要把“第一学期复盘与第二学期准备”做成 Atlas 可交付能力，你会如何设计主路径和边界？
2. 本周最值得优先验证的失败模式是什么？你会如何收集证据？
3. 这周的 RFC、实现、验证或复盘里，哪一项最能体现你的工程判断？为什么？
4. 如果时间被砍掉一半，你会保留什么，推迟什么？
5. 本周产物如何自然承接到下周，而不是变成一次性文档或一次性代码？

## 每日入口

- [Day 085](days-semester-1/day-085.md): 理解问题和边界
- [Day 086](days-semester-1/day-086.md): 边界与契约设计
- [Day 087](days-semester-1/day-087.md): 最小主路径实现
- [Day 088](days-semester-1/day-088.md): 补全真实可用路径
- [Day 089](days-semester-1/day-089.md): 验证与实验
- [Day 090](days-semester-1/day-090.md): 文档与评审
- [Day 091](days-semester-1/day-091.md): 复盘与 memory

## 配套阅读

- [Week 13](../weeks/week-13.md)
- [Quarter 1 Replay Pack](../reviews/archive/quarter-1-replay.md)
