# Week 48：Evals、Approval 与 Audit 整合

最后更新：2026-06-30

## 本周目标

这一周把治理和评估也拉回平台主链路。

## 本周学完后你应该会什么

重点是让平台默认就是可治理、可评估的。

## 本周学习重点

- eval integration
- approval integration
- audit integration

## 本周 Cookbook 做法

### Recipe 1：接入 eval

### Recipe 2：接入 approval

### Recipe 3：接入 audit

## 本周交付标准

- 平台级评估链路
- 审批链路
- 审计链路

## 本周能力焦点

- 平台整合（`Platform Integration`）
- 架构判断（`Architecture Judgment`）
- 技术领导力（`Technical Leadership`）

## 本周知识清单

- eval、approval 和 audit integration 的目标是让高风险决策变成一条带证据的治理链。
- 同一个动作应该携带 evaluation context、approval state 和 audit history 一起流动。
- release governance 只有在质量信号和策略信号汇合时才真正成立。
- provenance 决定了一条记录是普通日志还是可审计决策轨迹。

## 本周 Atlas 观察路径

- `[Week 48 英文原文](../weeks/week-48.md)`：对照原始英文任务拆分，确认这一周的 deliverables、source reading 和 daily links。
- `[Atlas 主线 Issue（Month 12）](../atlas/issues/month-12/atlas-m12-02.md)`：把本周主题放回 Atlas 主线问题卡，观察范围、验收项和剩余风险。
- `[Week 48 周报](../reports/weekly/week-48-report.md)`：看本周结果、证据和后续动作是如何被记录的。
- `[docs/governance/steering-committee-pack.md](../governance/steering-committee-pack.md)`：对照跨子系统整合后的决策材料组织方式。

## 本周 Atlas 落地检查

- 最小落地动作：围绕“接入 eval”推进，让至少一项本周交付标准在仓库里可见。
- 必留证据：integrated flow、provenance model 和 governance path。
- 最大风险：质量门禁和审批门禁各自存在，但彼此没有共享决策语义。
- 审查不变量：一项敏感发布或动作必须能从一条 joined evidence chain 中被完整解释。
- 本周最先要盯住的交付：平台级评估链路

## 本周最小演示场景

- 背景：这一周你需要把“Evals、Approval 与 Audit 整合”从概念推进成可被别人看懂、验证和质疑的工程结果。
- 你至少要能演示：
  - 平台级评估链路、审批链路、审计链路
  - “eval integration”为什么会直接影响 Atlas 的主路径，而不只是局部实现
  - 风险“质量门禁和审批门禁各自存在，但彼此没有共享决策语义。”是如何被控制、暴露或留证的
- reviewer 大概率会追问：
  - 为什么这周先做“接入 eval”，而不是先做别的？
  - 不变量“一项敏感发布或动作必须能从一条 joined evidence chain 中被完整解释。”如果被破坏，会先造成什么后果？
  - 你打算拿什么证据证明这周成果成立？提示：integrated flow、provenance model 和 governance path。

## 配套案例

- 本周最接近的案例：读 [专题案例包](casebook.md) 中的“案例 12：功能都在，为什么还是说不清这是个平台”。
- 本周最该对照的失败模式：读 [失败场景图谱](failure-patterns.md) 中的“模式五：平台错觉型故障”。
- 推荐顺序：先读本周知识清单，再读案例，再回本周 Atlas 观察路径和周报，检查自己能不能解释坏信号与最小证据。

## 本周工作顺序

- [`RFC`](artifact-playbook.md#artifact-rfc)：先把“接入 eval”对应的问题范围、目标、约束和非目标写清楚。
- [`Interface`](artifact-playbook.md#artifact-interface)：围绕“eval integration”定义输入、输出、owner、错误语义和边界。
- [`Implementation`](artifact-playbook.md#artifact-implementation)：把最小主路径打通，至少让“平台级评估链路、审批链路、审计链路”里的一项开始可验证。
- [`Unit Test`](artifact-playbook.md#artifact-unit-test)：优先覆盖最容易击中风险“质量门禁和审批门禁各自存在，但彼此没有共享决策语义。”的失败路径、边界条件或 contract。
- [`Review`](artifact-playbook.md#artifact-review)：站在 reviewer 视角检查不变量“一项敏感发布或动作必须能从一条 joined evidence chain 中被完整解释。”是否真的被系统保护，而不是只写在文档里。
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

1. 如果要把“Evals、Approval 与 Audit 整合”做成 Atlas 可交付能力，你会如何设计主路径和边界？
2. 本周最值得优先验证的失败模式是什么？你会如何收集证据？
3. 这周的 RFC、实现、验证或复盘里，哪一项最能体现你的工程判断？为什么？
4. 如果时间被砍掉一半，你会保留什么，推迟什么？
5. 本周产物如何自然承接到下周，而不是变成一次性文档或一次性代码？

## 每日入口

- [Day 330](days-semester-4/day-330.md): 理解问题和边界
- [Day 331](days-semester-4/day-331.md): 边界与契约设计
- [Day 332](days-semester-4/day-332.md): 最小主路径实现
- [Day 333](days-semester-4/day-333.md): 补全真实可用路径
- [Day 334](days-semester-4/day-334.md): 验证与实验
- [Day 335](days-semester-4/day-335.md): 文档与评审
- [Day 336](days-semester-4/day-336.md): 复盘与 memory

## 配套阅读

- [Week 48](../weeks/week-48.md)
