# Week 03：配置、日志、数据库与缓存接线

最后更新：2026-06-30

## 本周目标

这一周要把服务骨架真正接上核心基础设施。

重点不是“连上数据库就行”，而是：

- 依赖边界清晰
- 配置入口统一
- 日志入口统一

## 本周学习重点

- `config.Manager` 为什么是唯一入口
- 日志为什么必须统一抽象
- 数据库和缓存为什么不能散在业务代码里

## 本周 Cookbook 做法

### Recipe 1：接入配置中心

### Recipe 2：接入统一日志入口

### Recipe 3：接入 Postgres / Redis 依赖边界

## 本周交付标准

到本周结束时，你至少要能解释：

- 配置从哪里进来
- 日志如何配置
- 依赖在什么时候初始化

## 本周能力焦点

- 结果交付（`Deliver Results`）
- 结构与清晰度（`Structure and Clarity`）
- 代码健康（`Code Health`）

## 本周知识清单

- `config.Manager` 要统一默认值、配置文件和环境覆盖顺序，业务代码不能直接读环境变量。
- 日志抽象的价值在于保留 caller source、结构化字段和未来 trace 对接能力，而不是简单包一层输出函数。
- 数据库和缓存客户端应该作为 runtime 依赖注入，而不是散落在业务代码里即时创建。
- 连接管理、超时、健康信号和配置约束都属于基础设施基线的一部分。

## 本周 Atlas 观察路径

- `[Week 03 英文原文](../weeks/week-03.md)`：对照原始英文任务拆分，确认这一周的 deliverables、source reading 和 daily links。
- `[Atlas 主线 Issue（Month 01）](../atlas/issues/month-01/atlas-m01-03.md)`：把本周主题放回 Atlas 主线问题卡，观察范围、验收项和剩余风险。
- `[Week 03 周报](../reports/weekly/week-03-report.md)`：看本周结果、证据和后续动作是如何被记录的。
- `[internal/config/manager.go](../../internal/config/manager.go)`、`[internal/logger/logger.go](../../internal/logger/logger.go)`：对照统一配置与统一日志抽象。

## 本周 Atlas 落地检查

- 最小落地动作：围绕“接入配置中心”推进，让至少一项本周交付标准在仓库里可见。
- 必留证据：配置优先级验证、统一日志行为和依赖接线 smoke test。
- 最大风险：直接在业务代码里读 `os.Getenv`、自己 new logger 或自己连数据库，导致边界失控。
- 审查不变量：业务代码只能消费注入后的配置、日志和依赖，不能绕过统一入口。
- 本周最先要盯住的交付：配置从哪里进来

## 本周最小演示场景

- 背景：这一周你需要把“配置、日志、数据库与缓存接线”从概念推进成可被别人看懂、验证和质疑的工程结果。
- 你至少要能演示：
  - 配置从哪里进来、日志如何配置、依赖在什么时候初始化
  - “`config.Manager` 为什么是唯一入口”为什么会直接影响 Atlas 的主路径，而不只是局部实现
  - 风险“直接在业务代码里读 `os.Getenv`、自己 new logger 或自己连数据库，导致边界失控。”是如何被控制、暴露或留证的
- reviewer 大概率会追问：
  - 为什么这周先做“接入配置中心”，而不是先做别的？
  - 不变量“业务代码只能消费注入后的配置、日志和依赖，不能绕过统一入口。”如果被破坏，会先造成什么后果？
  - 你打算拿什么证据证明这周成果成立？提示：配置优先级验证、统一日志行为和依赖接线 smoke test。

## 配套案例

- 本周最接近的案例：读 [专题案例包](casebook.md) 中的“案例 1：服务能启动，但谁都不敢改”。
- 本周最该对照的失败模式：读 [失败场景图谱](failure-patterns.md) 中的“模式一：边界缺失型故障”。
- 推荐顺序：先读本周知识清单，再读案例，再回本周 Atlas 观察路径和周报，检查自己能不能解释坏信号与最小证据。

## 本周工作顺序

- [`RFC`](artifact-playbook.md#artifact-rfc)：先把“接入配置中心”对应的问题范围、目标、约束和非目标写清楚。
- [`Interface`](artifact-playbook.md#artifact-interface)：围绕“`config.Manager` 为什么是唯一入口”定义输入、输出、owner、错误语义和边界。
- [`Implementation`](artifact-playbook.md#artifact-implementation)：把最小主路径打通，至少让“配置从哪里进来、日志如何配置、依赖在什么时候初始化”里的一项开始可验证。
- [`Unit Test`](artifact-playbook.md#artifact-unit-test)：优先覆盖最容易击中风险“直接在业务代码里读 `os.Getenv`、自己 new logger 或自己连数据库，导致边界失控。”的失败路径、边界条件或 contract。
- [`Review`](artifact-playbook.md#artifact-review)：站在 reviewer 视角检查不变量“业务代码只能消费注入后的配置、日志和依赖，不能绕过统一入口。”是否真的被系统保护，而不是只写在文档里。
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

1. 如果要把“配置、日志、数据库与缓存接线”做成 Atlas 可交付能力，你会如何设计主路径和边界？
2. 本周最值得优先验证的失败模式是什么？你会如何收集证据？
3. 这周的 RFC、实现、验证或复盘里，哪一项最能体现你的工程判断？为什么？
4. 如果时间被砍掉一半，你会保留什么，推迟什么？
5. 本周产物如何自然承接到下周，而不是变成一次性文档或一次性代码？

## 每日入口

- [Day 015](days-month-01/day-015.md): 理解问题和边界
- [Day 016](days-month-01/day-016.md): 边界与契约设计
- [Day 017](days-month-01/day-017.md): 最小主路径实现
- [Day 018](days-month-01/day-018.md): 补全真实可用路径
- [Day 019](days-month-01/day-019.md): 验证与实验
- [Day 020](days-month-01/day-020.md): 文档与评审
- [Day 021](days-month-01/day-021.md): 复盘与 memory

## 配套阅读

- [Week 03](../weeks/week-03.md)
- [Month 01 Backlog](../atlas/backlog/month-01.md)
