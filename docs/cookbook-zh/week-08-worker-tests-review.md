# Week 08：可靠性验证与第二月复盘

最后更新：2026-06-30

## 本周目标

用测试、验证和 drill 去证明后台执行模型不是“看起来能用”，而是真的可用。

## 本周学习重点

- 失败路径测试
- drill 文档
- 第二月复盘

## 本周 Cookbook 做法

### Recipe 1：补 worker tests

### Recipe 2：写 failure drill

### Recipe 3：做第二月复盘

## 本周交付标准

你至少要能展示：

- 一组失败路径测试
- 一份 drill 说明
- 第二月有哪些可靠性结论

## 本周能力焦点

- 可信可靠（`Dependability`）
- 可靠性（`Reliability`）
- 运维思维（`Operational Thinking`）

## 本周知识清单

- 可靠性证明依赖失败路径测试，而不是只看 worker 能不能跑通 happy path。
- drill 至少要覆盖 worker crash、timeout、重复投递、队列积压和人工恢复步骤。
- 可靠性证据不仅是测试通过，还包括 queue lag、retry count、dead-letter volume 这类运行信号。
- 月度复盘应该说明哪些失败模式已经系统化，哪些还依赖人工经验。

## 本周 Atlas 观察路径

- `[Week 08 英文原文](../weeks/week-08.md)`：对照原始英文任务拆分，确认这一周的 deliverables、source reading 和 daily links。
- `[Atlas 主线 Issue（Month 02）](../atlas/issues/month-02/atlas-m02-04.md)`：把本周主题放回 Atlas 主线问题卡，观察范围、验收项和剩余风险。
- `[Week 08 周报](../reports/weekly/week-08-report.md)`：看本周结果、证据和后续动作是如何被记录的。
- `[docs/operations/issue-status-board.md](../operations/issue-status-board.md)`：观察任务状态、阻塞和恢复动作如何被管理。

## 本周 Atlas 落地检查

- 最小落地动作：围绕“补 worker tests”推进，让至少一项本周交付标准在仓库里可见。
- 必留证据：失败路径测试、drill 文档和第二月复盘。
- 最大风险：测试全部绿色，但从未演练真正会在生产里出现的故障。
- 审查不变量：每条可靠性声明都要能映射到至少一个测试、drill 或 runbook。
- 本周最先要盯住的交付：一组失败路径测试

## 本周最小演示场景

- 背景：这一周你需要把“可靠性验证与第二月复盘”从概念推进成可被别人看懂、验证和质疑的工程结果。
- 你至少要能演示：
  - 一组失败路径测试、一份 drill 说明、第二月有哪些可靠性结论
  - “失败路径测试”为什么会直接影响 Atlas 的主路径，而不只是局部实现
  - 风险“测试全部绿色，但从未演练真正会在生产里出现的故障。”是如何被控制、暴露或留证的
- reviewer 大概率会追问：
  - 为什么这周先做“补 worker tests”，而不是先做别的？
  - 不变量“每条可靠性声明都要能映射到至少一个测试、drill 或 runbook。”如果被破坏，会先造成什么后果？
  - 你打算拿什么证据证明这周成果成立？提示：失败路径测试、drill 文档和第二月复盘。

## 配套案例

- 本周最接近的案例：读 [专题案例包](casebook.md) 中的“案例 2：任务系统看起来能跑，线上却越跑越乱”。
- 本周最该对照的失败模式：读 [失败场景图谱](failure-patterns.md) 中的“模式二：失败语义缺失型故障”。
- 推荐顺序：先读本周知识清单，再读案例，再回本周 Atlas 观察路径和周报，检查自己能不能解释坏信号与最小证据。

## 本周工作顺序

- [`RFC`](artifact-playbook.md#artifact-rfc)：先把“补 worker tests”对应的问题范围、目标、约束和非目标写清楚。
- [`Interface`](artifact-playbook.md#artifact-interface)：围绕“失败路径测试”定义输入、输出、owner、错误语义和边界。
- [`Implementation`](artifact-playbook.md#artifact-implementation)：把最小主路径打通，至少让“一组失败路径测试、一份 drill 说明、第二月有哪些可靠性结论”里的一项开始可验证。
- [`Unit Test`](artifact-playbook.md#artifact-unit-test)：优先覆盖最容易击中风险“测试全部绿色，但从未演练真正会在生产里出现的故障。”的失败路径、边界条件或 contract。
- [`Review`](artifact-playbook.md#artifact-review)：站在 reviewer 视角检查不变量“每条可靠性声明都要能映射到至少一个测试、drill 或 runbook。”是否真的被系统保护，而不是只写在文档里。
- [`Benchmark（按需）`](artifact-playbook.md#artifact-benchmark)：如果本周做了 failure drill，至少记录一次失败注入结果和恢复时间。
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

1. 如果要把“可靠性验证与第二月复盘”做成 Atlas 可交付能力，你会如何设计主路径和边界？
2. 本周最值得优先验证的失败模式是什么？你会如何收集证据？
3. 这周的 RFC、实现、验证或复盘里，哪一项最能体现你的工程判断？为什么？
4. 如果时间被砍掉一半，你会保留什么，推迟什么？
5. 本周产物如何自然承接到下周，而不是变成一次性文档或一次性代码？

## 每日入口

- [Day 050](days-semester-1/day-050.md): 理解问题和边界
- [Day 051](days-semester-1/day-051.md): 边界与契约设计
- [Day 052](days-semester-1/day-052.md): 最小主路径实现
- [Day 053](days-semester-1/day-053.md): 补全真实可用路径
- [Day 054](days-semester-1/day-054.md): 验证与实验
- [Day 055](days-semester-1/day-055.md): 文档与评审
- [Day 056](days-semester-1/day-056.md): 复盘与 memory

## 配套阅读

- [Week 08](../weeks/week-08.md)
- [Month 02 Scorecard](../reports/monthly/month-02-scorecard.md)
