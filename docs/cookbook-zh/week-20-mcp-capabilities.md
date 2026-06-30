# Week 20：MCP Capability 与兼容性

最后更新：2026-06-30

## 本周目标

本周要让 MCP 不只是能启动，还要开始具备真正可消费的能力。

## 本周学完后你应该会什么

同时，你必须第一次认真考虑兼容性。

## 本周学习重点

- tools/resources capability
- 兼容性边界
- 协议演进

## 本周 Cookbook 做法

### Recipe 1：暴露至少两类 capability

### Recipe 2：记录兼容策略

### Recipe 3：整理 error path

## 本周交付标准

- 至少两类 MCP capability
- 兼容性说明
- 协议错误路径

## 本周能力焦点

- 协议思维（`Protocol Thinking`）
- 兼容性（`Compatibility`）
- 把文档当产品（`Documentation as Product`）

## 本周知识清单

- capability descriptor 是客户端发现 contract 的方式，不只是内部标签。
- 兼容性边界要回答新增、修改、弃用和删除 capability 时的规则。
- 协议演进应优先走 additive change 和显式 version signal，避免静默 break。
- client 在 capability 不可用或版本过老时也需要可预测的错误与 fallback 语义。

## 本周 Atlas 观察路径

- `[Week 20 英文原文](../weeks/week-20.md)`：对照原始英文任务拆分，确认这一周的 deliverables、source reading 和 daily links。
- `[Atlas 主线 Issue（Month 05）](../atlas/issues/month-05/atlas-m05-03.md)`：把本周主题放回 Atlas 主线问题卡，观察范围、验收项和剩余风险。
- `[Week 20 周报](../reports/weekly/week-20-report.md)`：看本周结果、证据和后续动作是如何被记录的。
- `[docs/reports/monthly/month-05-scorecard.md](../reports/monthly/month-05-scorecard.md)`：看协议兼容和文档质量如何进入月度验收。

## 本周 Atlas 落地检查

- 最小落地动作：围绕“暴露至少两类 capability”推进，让至少一项本周交付标准在仓库里可见。
- 必留证据：capability schema、compatibility rules 和 evolution note。
- 最大风险：服务端一改 capability，客户端却直到线上失败才知道 contract 变了。
- 审查不变量：client 在调用前应该能知道 server 支持什么、不支持什么。
- 本周最先要盯住的交付：至少两类 MCP capability

## 本周最小演示场景

- 背景：这一周你需要把“MCP Capability 与兼容性”从概念推进成可被别人看懂、验证和质疑的工程结果。
- 你至少要能演示：
  - 至少两类 MCP capability、兼容性说明、协议错误路径
  - “tools/resources capability”为什么会直接影响 Atlas 的主路径，而不只是局部实现
  - 风险“服务端一改 capability，客户端却直到线上失败才知道 contract 变了。”是如何被控制、暴露或留证的
- reviewer 大概率会追问：
  - 为什么这周先做“暴露至少两类 capability”，而不是先做别的？
  - 不变量“client 在调用前应该能知道 server 支持什么、不支持什么。”如果被破坏，会先造成什么后果？
  - 你打算拿什么证据证明这周成果成立？提示：capability schema、compatibility rules 和 evolution note。

## 配套案例

- 本周最接近的案例：读 [专题案例包](casebook.md) 中的“案例 5：MCP 服务能启动，但根本不是可消费协议”。
- 本周最该对照的失败模式：读 [失败场景图谱](failure-patterns.md) 中的“模式一：边界缺失型故障”。
- 推荐顺序：先读本周知识清单，再读案例，再回本周 Atlas 观察路径和周报，检查自己能不能解释坏信号与最小证据。

## 本周工作顺序

- [`RFC`](artifact-playbook.md#artifact-rfc)：先把“暴露至少两类 capability”对应的问题范围、目标、约束和非目标写清楚。
- [`Interface`](artifact-playbook.md#artifact-interface)：围绕“tools/resources capability”定义输入、输出、owner、错误语义和边界。
- [`Implementation`](artifact-playbook.md#artifact-implementation)：把最小主路径打通，至少让“至少两类 MCP capability、兼容性说明、协议错误路径”里的一项开始可验证。
- [`Unit Test`](artifact-playbook.md#artifact-unit-test)：优先覆盖最容易击中风险“服务端一改 capability，客户端却直到线上失败才知道 contract 变了。”的失败路径、边界条件或 contract。
- [`Review`](artifact-playbook.md#artifact-review)：站在 reviewer 视角检查不变量“client 在调用前应该能知道 server 支持什么、不支持什么。”是否真的被系统保护，而不是只写在文档里。
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

1. 如果要把“MCP Capability 与兼容性”做成 Atlas 可交付能力，你会如何设计主路径和边界？
2. 本周最值得优先验证的失败模式是什么？你会如何收集证据？
3. 这周的 RFC、实现、验证或复盘里，哪一项最能体现你的工程判断？为什么？
4. 如果时间被砍掉一半，你会保留什么，推迟什么？
5. 本周产物如何自然承接到下周，而不是变成一次性文档或一次性代码？

## 每日入口

- [Day 134](days-semester-2/day-134.md): 理解问题和边界
- [Day 135](days-semester-2/day-135.md): 边界与契约设计
- [Day 136](days-semester-2/day-136.md): 最小主路径实现
- [Day 137](days-semester-2/day-137.md): 补全真实可用路径
- [Day 138](days-semester-2/day-138.md): 验证与实验
- [Day 139](days-semester-2/day-139.md): 文档与评审
- [Day 140](days-semester-2/day-140.md): 复盘与 memory

## 配套阅读

- [Week 20](../weeks/week-20.md)
