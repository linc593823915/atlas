# Week 35：治理 RFC 与威胁模型

最后更新：2026-06-30

## 本周目标

这一周开始正式处理 Agent 治理。先定义 threat model 和治理边界。

## 本周学完后你应该会什么

不要急着写 policy engine，先回答威胁来自哪里。

## 本周学习重点

- threat model
- risk taxonomy
- governance boundary

## 本周 Cookbook 做法

### Recipe 1：写治理 RFC

### Recipe 2：写 threat model

### Recipe 3：定义高风险动作

## 本周交付标准

- 治理 RFC
- threat model
- 高风险动作清单

## 本周能力焦点

- 安全判断（`Security Judgment`）
- 策略控制（`Policy Controls`）
- 可审计性（`Auditability`）

## 本周知识清单

- threat model 必须识别资产、角色、入口和攻击路径，而且要针对 Agent 系统的真实行为。
- governance boundary 决定哪些动作需要 policy、approval、sanitization 或直接 deny。
- risk taxonomy 要按 blast radius 和 exploitability 排，而不是按恐惧程度排。
- RFC 的价值在于把 threat 映射到 controls 和 tests，而不是写一份泛安全 memo。

## 本周 Atlas 观察路径

- `[Week 35 英文原文](../weeks/week-35.md)`：对照原始英文任务拆分，确认这一周的 deliverables、source reading 和 daily links。
- `[Atlas 主线 Issue（Month 09）](../atlas/issues/month-09/atlas-m09-01.md)`：把本周主题放回 Atlas 主线问题卡，观察范围、验收项和剩余风险。
- `[Week 35 周报](../reports/weekly/week-35-report.md)`：看本周结果、证据和后续动作是如何被记录的。
- `[docs/governance/committee-decision-examples.md](../governance/committee-decision-examples.md)`：对照 threat / control 如何变成可执行治理决定。

## 本周 Atlas 落地检查

- 最小落地动作：围绕“写治理 RFC”推进，让至少一项本周交付标准在仓库里可见。
- 必留证据：governance RFC、threat model 和 risk taxonomy。
- 最大风险：写了一个听起来很安全的文档，但没有把任何高风险动作真正拉进模型。
- 审查不变量：每一个高风险动作都必须在 threat model 里出现，并且至少挂一个 control。
- 本周最先要盯住的交付：治理 RFC

## 本周最小演示场景

- 背景：这一周你需要把“治理 RFC 与威胁模型”从概念推进成可被别人看懂、验证和质疑的工程结果。
- 你至少要能演示：
  - 治理 RFC、threat model、高风险动作清单
  - “threat model”为什么会直接影响 Atlas 的主路径，而不只是局部实现
  - 风险“写了一个听起来很安全的文档，但没有把任何高风险动作真正拉进模型。”是如何被控制、暴露或留证的
- reviewer 大概率会追问：
  - 为什么这周先做“写治理 RFC”，而不是先做别的？
  - 不变量“每一个高风险动作都必须在 threat model 里出现，并且至少挂一个 control。”如果被破坏，会先造成什么后果？
  - 你打算拿什么证据证明这周成果成立？提示：governance RFC、threat model 和 risk taxonomy。

## 配套案例

- 本周最接近的案例：读 [专题案例包](casebook.md) 中的“案例 9：治理规则很多，但真正出事时一条都兜不住”。
- 本周最该对照的失败模式：读 [失败场景图谱](failure-patterns.md) 中的“模式四：治理失守型故障”。
- 推荐顺序：先读本周知识清单，再读案例，再回本周 Atlas 观察路径和周报，检查自己能不能解释坏信号与最小证据。

## 本周工作顺序

- [`RFC`](artifact-playbook.md#artifact-rfc)：先把“写治理 RFC”对应的问题范围、目标、约束和非目标写清楚。
- [`Interface`](artifact-playbook.md#artifact-interface)：围绕“threat model”定义输入、输出、owner、错误语义和边界。
- [`Implementation`](artifact-playbook.md#artifact-implementation)：把最小主路径打通，至少让“治理 RFC、threat model、高风险动作清单”里的一项开始可验证。
- [`Unit Test`](artifact-playbook.md#artifact-unit-test)：优先覆盖最容易击中风险“写了一个听起来很安全的文档，但没有把任何高风险动作真正拉进模型。”的失败路径、边界条件或 contract。
- [`Review`](artifact-playbook.md#artifact-review)：站在 reviewer 视角检查不变量“每一个高风险动作都必须在 threat model 里出现，并且至少挂一个 control。”是否真的被系统保护，而不是只写在文档里。
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

1. 如果要把“治理 RFC 与威胁模型”做成 Atlas 可交付能力，你会如何设计主路径和边界？
2. 本周最值得优先验证的失败模式是什么？你会如何收集证据？
3. 这周的 RFC、实现、验证或复盘里，哪一项最能体现你的工程判断？为什么？
4. 如果时间被砍掉一半，你会保留什么，推迟什么？
5. 本周产物如何自然承接到下周，而不是变成一次性文档或一次性代码？

## 每日入口

- [Day 239](days-semester-3/day-239.md): 理解问题和边界
- [Day 240](days-semester-3/day-240.md): 边界与契约设计
- [Day 241](days-semester-3/day-241.md): 最小主路径实现
- [Day 242](days-semester-3/day-242.md): 补全真实可用路径
- [Day 243](days-semester-3/day-243.md): 验证与实验
- [Day 244](days-semester-3/day-244.md): 文档与评审
- [Day 245](days-semester-3/day-245.md): 复盘与 memory

## 配套阅读

- [Week 35](../weeks/week-35.md)
