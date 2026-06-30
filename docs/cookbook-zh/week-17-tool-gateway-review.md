# Week 17：Tool Gateway 测试与复盘

最后更新：2026-06-30

## 本周目标

这一周要证明 Tool Gateway 不是“看起来合理”，而是真的可交付。

## 本周学完后你应该会什么

你需要通过测试、review 和复盘，把它收口成平台资产。

## 本周学习重点

- Contract tests
- Negative tests
- Review criteria

## 本周 Cookbook 做法

### Recipe 1：补 Tool Gateway 测试

### Recipe 2：整理接入文档

### Recipe 3：做第四个月复盘

## 本周交付标准

- Tool Gateway 测试集
- 接入文档
- 第四个月复盘结论

## 本周能力焦点

- 接口设计（`Interface Design`）
- 代码评审质量（`Code Review Quality`）
- 小步变更纪律（`Small Change Discipline`）

## 本周知识清单

- contract tests 要证明 gateway 对 schema、auth、timeout 和 audit 的承诺都成立。
- negative tests 才是 gateway 真正发挥价值的地方，因为错误输入和拒绝路径比 happy path 更容易出事故。
- review criteria 应该覆盖 debuggability、policy surface 和 extension cost，而不是只看代码风格。
- 月度复盘要说明 gateway 是否真的简化了工具接入，而不是只新增了一层抽象。

## 本周 Atlas 观察路径

- `[Week 17 英文原文](../weeks/week-17.md)`：对照原始英文任务拆分，确认这一周的 deliverables、source reading 和 daily links。
- `[Atlas 主线 Issue（Month 04）](../atlas/issues/month-04/atlas-m04-04.md)`：把本周主题放回 Atlas 主线问题卡，观察范围、验收项和剩余风险。
- `[Week 17 周报](../reports/weekly/week-17-report.md)`：看本周结果、证据和后续动作是如何被记录的。
- `[docs/reviews/archive/month-04-review.md](../reviews/archive/month-04-review.md)`：复盘周重点看哪些 contract 最脆弱。

## 本周 Atlas 落地检查

- 最小落地动作：围绕“补 Tool Gateway 测试”推进，让至少一项本周交付标准在仓库里可见。
- 必留证据：contract tests、negative tests 和 gateway review。
- 最大风险：只测试工具能否成功调用，却从不验证 contract 被破坏时系统怎么表现。
- 审查不变量：每一条 gateway promise 都应该至少有一个正例和一个负例。
- 本周最先要盯住的交付：Tool Gateway 测试集

## 本周最小演示场景

- 背景：这一周你需要把“Tool Gateway 测试与复盘”从概念推进成可被别人看懂、验证和质疑的工程结果。
- 你至少要能演示：
  - Tool Gateway 测试集、接入文档、第四个月复盘结论
  - “Contract tests”为什么会直接影响 Atlas 的主路径，而不只是局部实现
  - 风险“只测试工具能否成功调用，却从不验证 contract 被破坏时系统怎么表现。”是如何被控制、暴露或留证的
- reviewer 大概率会追问：
  - 为什么这周先做“补 Tool Gateway 测试”，而不是先做别的？
  - 不变量“每一条 gateway promise 都应该至少有一个正例和一个负例。”如果被破坏，会先造成什么后果？
  - 你打算拿什么证据证明这周成果成立？提示：contract tests、negative tests 和 gateway review。

## 配套案例

- 本周最接近的案例：读 [专题案例包](casebook.md) 中的“案例 4：工具越来越多，但系统越来越难治理”。
- 本周最该对照的失败模式：读 [失败场景图谱](failure-patterns.md) 中的“模式一：边界缺失型故障”。
- 推荐顺序：先读本周知识清单，再读案例，再回本周 Atlas 观察路径和周报，检查自己能不能解释坏信号与最小证据。

## 本周工作顺序

- [`RFC`](artifact-playbook.md#artifact-rfc)：先把“补 Tool Gateway 测试”对应的问题范围、目标、约束和非目标写清楚。
- [`Interface`](artifact-playbook.md#artifact-interface)：围绕“Contract tests”定义输入、输出、owner、错误语义和边界。
- [`Implementation`](artifact-playbook.md#artifact-implementation)：把最小主路径打通，至少让“Tool Gateway 测试集、接入文档、第四个月复盘结论”里的一项开始可验证。
- [`Unit Test`](artifact-playbook.md#artifact-unit-test)：优先覆盖最容易击中风险“只测试工具能否成功调用，却从不验证 contract 被破坏时系统怎么表现。”的失败路径、边界条件或 contract。
- [`Review`](artifact-playbook.md#artifact-review)：站在 reviewer 视角检查不变量“每一条 gateway promise 都应该至少有一个正例和一个负例。”是否真的被系统保护，而不是只写在文档里。
- [`Benchmark（按需）`](artifact-playbook.md#artifact-benchmark)：本周不做性能 benchmark，但要留下一轮 contract / negative test 结果。
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

1. 如果要把“Tool Gateway 测试与复盘”做成 Atlas 可交付能力，你会如何设计主路径和边界？
2. 本周最值得优先验证的失败模式是什么？你会如何收集证据？
3. 这周的 RFC、实现、验证或复盘里，哪一项最能体现你的工程判断？为什么？
4. 如果时间被砍掉一半，你会保留什么，推迟什么？
5. 本周产物如何自然承接到下周，而不是变成一次性文档或一次性代码？

## 每日入口

- [Day 113](days-semester-2/day-113.md): 理解问题和边界
- [Day 114](days-semester-2/day-114.md): 边界与契约设计
- [Day 115](days-semester-2/day-115.md): 最小主路径实现
- [Day 116](days-semester-2/day-116.md): 补全真实可用路径
- [Day 117](days-semester-2/day-117.md): 验证与实验
- [Day 118](days-semester-2/day-118.md): 文档与评审
- [Day 119](days-semester-2/day-119.md): 复盘与 memory

## 配套阅读

- [Week 17](../weeks/week-17.md)
