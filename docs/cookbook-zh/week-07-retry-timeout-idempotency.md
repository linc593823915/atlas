# Week 07：重试、超时、幂等与死信

最后更新：2026-06-30

## 本周目标

本周要把后台执行模型从“能跑”升级到“有失败治理”。

## 本周学习重点

- retry 的边界
- timeout 的责任
- 幂等的真正用途
- dead-letter 何时出现

## 本周 Cookbook 做法

### Recipe 1：定义 retry 策略

### Recipe 2：加入 timeout/cancel

### Recipe 3：补 idempotency 与 dead-letter

## 本周交付标准

你至少要能解释：

- 什么错误会重试
- 什么错误必须直接失败
- 重复任务为什么不会造成灾难

## 本周能力焦点

- 可信可靠（`Dependability`）
- 可靠性（`Reliability`）
- 运维思维（`Operational Thinking`）

## 本周知识清单

- retry 策略要区分暂时性依赖错误和确定性业务错误，否则系统只会放大噪音。
- timeout 的预算归属必须清楚：调用方、runtime 和下游客户端谁负责取消、谁负责补偿。
- 幂等保护的意义是在重试、重启和重复投递下仍然不放大副作用。
- dead-letter queue 不是垃圾桶，而是 operator 学习失败模式和改进系统的证据面。

## 本周 Atlas 观察路径

- `[Week 07 英文原文](../weeks/week-07.md)`：对照原始英文任务拆分，确认这一周的 deliverables、source reading 和 daily links。
- `[Atlas 主线 Issue（Month 02）](../atlas/issues/month-02/atlas-m02-03.md)`：把本周主题放回 Atlas 主线问题卡，观察范围、验收项和剩余风险。
- `[Week 07 周报](../reports/weekly/week-07-report.md)`：看本周结果、证据和后续动作是如何被记录的。
- `[docs/operations/evidence-registry.md](../operations/evidence-registry.md)`：把重试、超时、dead-letter 这类证据放回统一登记表理解。

## 本周 Atlas 落地检查

- 最小落地动作：围绕“定义 retry 策略”推进，让至少一项本周交付标准在仓库里可见。
- 必留证据：retry matrix、timeout policy、幂等策略和 dead-letter 处理说明。
- 最大风险：无限重试把真实 bug 伪装成“系统还在努力工作”。
- 审查不变量：同一个任务重试多次也不能放大不可逆副作用。
- 本周最先要盯住的交付：什么错误会重试

## 本周最小演示场景

- 背景：这一周你需要把“重试、超时、幂等与死信”从概念推进成可被别人看懂、验证和质疑的工程结果。
- 你至少要能演示：
  - 什么错误会重试、什么错误必须直接失败、重复任务为什么不会造成灾难
  - “retry 的边界”为什么会直接影响 Atlas 的主路径，而不只是局部实现
  - 风险“无限重试把真实 bug 伪装成“系统还在努力工作”。”是如何被控制、暴露或留证的
- reviewer 大概率会追问：
  - 为什么这周先做“定义 retry 策略”，而不是先做别的？
  - 不变量“同一个任务重试多次也不能放大不可逆副作用。”如果被破坏，会先造成什么后果？
  - 你打算拿什么证据证明这周成果成立？提示：retry matrix、timeout policy、幂等策略和 dead-letter 处理说明。

## 配套案例

- 本周最接近的案例：读 [专题案例包](casebook.md) 中的“案例 2：任务系统看起来能跑，线上却越跑越乱”。
- 本周最该对照的失败模式：读 [失败场景图谱](failure-patterns.md) 中的“模式二：失败语义缺失型故障”。
- 推荐顺序：先读本周知识清单，再读案例，再回本周 Atlas 观察路径和周报，检查自己能不能解释坏信号与最小证据。

## 本周工作顺序

- [`RFC`](artifact-playbook.md#artifact-rfc)：先把“定义 retry 策略”对应的问题范围、目标、约束和非目标写清楚。
- [`Interface`](artifact-playbook.md#artifact-interface)：围绕“retry 的边界”定义输入、输出、owner、错误语义和边界。
- [`Implementation`](artifact-playbook.md#artifact-implementation)：把最小主路径打通，至少让“什么错误会重试、什么错误必须直接失败、重复任务为什么不会造成灾难”里的一项开始可验证。
- [`Unit Test`](artifact-playbook.md#artifact-unit-test)：优先覆盖最容易击中风险“无限重试把真实 bug 伪装成“系统还在努力工作”。”的失败路径、边界条件或 contract。
- [`Review`](artifact-playbook.md#artifact-review)：站在 reviewer 视角检查不变量“同一个任务重试多次也不能放大不可逆副作用。”是否真的被系统保护，而不是只写在文档里。
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

1. 如果要把“重试、超时、幂等与死信”做成 Atlas 可交付能力，你会如何设计主路径和边界？
2. 本周最值得优先验证的失败模式是什么？你会如何收集证据？
3. 这周的 RFC、实现、验证或复盘里，哪一项最能体现你的工程判断？为什么？
4. 如果时间被砍掉一半，你会保留什么，推迟什么？
5. 本周产物如何自然承接到下周，而不是变成一次性文档或一次性代码？

## 每日入口

- [Day 043](days-semester-1/day-043.md): 理解问题和边界
- [Day 044](days-semester-1/day-044.md): 边界与契约设计
- [Day 045](days-semester-1/day-045.md): 最小主路径实现
- [Day 046](days-semester-1/day-046.md): 补全真实可用路径
- [Day 047](days-semester-1/day-047.md): 验证与实验
- [Day 048](days-semester-1/day-048.md): 文档与评审
- [Day 049](days-semester-1/day-049.md): 复盘与 memory

## 配套阅读

- [Week 07](../weeks/week-07.md)
- [Month 02 Issue Set](../atlas/issues/month-02/README.md)
