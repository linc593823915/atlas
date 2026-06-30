# 评审清单与模式提示

最后更新：2026-06-30

## 这份文档怎么用

这不是另一本教材。

它的作用只有一个：

当你在 day 文档里看到下面这类词时，不要把它们当“随便再搜一搜”的提示，而是当成你应该立即带回当前问题的评审框架：

- `... review checklist`
- `... testing patterns`
- `... decision criteria`

使用顺序建议如下：

1. 先读当前 day 文档，确认今天的目标和最小交付
2. 再打开这里对应的小节，对照自己的设计、实现或证据
3. 用小节里的问题挑自己一次
4. 再回到 day 文档，把发现的问题写进 RFC、测试、review note 或 memory

## 架构与接口类

<a id="architecture-review-checklist"></a>
### `Architecture review checklist`

- 适用场景：
  - 第一学期复盘
  - 服务骨架、runtime、模块边界 review
- 重点追问：
  - 边界是不是通过代码和文档共同表达，而不是靠作者解释
  - 哪一层拥有配置、日志、依赖和生命周期
  - 如果新增入口或新增模块，现有边界会不会立刻被打穿
- 最低验收标准：
  - 模块职责能说清
  - 入口和依赖生命周期有单一 owner
  - 至少有一条测试或 smoke check 保护核心边界

<a id="api-review-checklist"></a>
### `API review checklist`

- 适用场景：
  - Tool Gateway
  - Tool schema / registry
  - MCP capability / interface review
- 重点追问：
  - 输入、输出、错误语义和 owner 是否写清楚
  - contract 是不是只在一处定义，而不是多处漂移
  - 这个 API 失败时调用方会看到什么、怎么恢复
- 最低验收标准：
  - 有显式 contract
  - 有负路径说明
  - 有版本或演进假设

<a id="architecture-review-checklist-for-orchestrators"></a>
### `Architecture review checklist for orchestrators`

- 适用场景：
  - 多 Agent runtime
  - handoff / approval / orchestration 设计
- 重点追问：
  - 谁拥有目标、谁拥有上下文、谁对最终结果负责
  - 一个 child agent 失败时，父流程怎么退化
  - shared context 有没有来源、边界和最小必要性
- 最低验收标准：
  - role map 清楚
  - handoff contract 清楚
  - 至少有一条 fallback 或人工接管路径

<a id="configuration-and-contract-governance-patterns"></a>
### `Configuration and contract governance patterns`

- 适用场景：
  - 配置中心
  - 共享 contract
  - 多子系统集成
- 重点追问：
  - 同一个配置或 contract 在系统里有没有 canonical source
  - default、覆盖、弃用和迁移规则是否明确
  - 新增能力时，是接入现有 contract，还是偷偷绕开
- 最低验收标准：
  - contract 有唯一来源
  - 配置优先级可解释
  - 变更路径可审查

## 测试与失败路径类

<a id="negative-path-testing-patterns"></a>
### `Negative-path testing patterns`

- 适用场景：
  - Tool Gateway
  - schema validation
  - MCP / API contract
- 重点追问：
  - 非法输入、拒绝路径、超时路径和部分失败路径有没有被覆盖
  - 错误输出是否仍然满足系统预期的错误 contract
  - 测试是不是只在验证 happy path
- 最低验收标准：
  - 至少有一个正例
  - 至少有一个负例
  - 失败后行为可解释

<a id="workflow-testing-patterns"></a>
### `Workflow testing patterns`

- 适用场景：
  - durable workflow
  - pause / resume / replay
  - 多 Agent orchestration
- 重点追问：
  - 中断、恢复、重复执行和人工介入有没有被系统化验证
  - 失败后是 replay、resume 还是 abort，路径是否清楚
  - 状态损坏或不完整时系统是否有诚实反应
- 最低验收标准：
  - 至少一条 stateful / orchestration 路径被验证
  - 至少一条故障恢复路径可复现
  - operator 可以照 runbook 操作

<a id="stateful-system-testing-patterns"></a>
### `Stateful system testing patterns`

- 适用场景：
  - checkpoint
  - resume cursor
  - state machine
- 重点追问：
  - 同一个状态能否被唯一解释
  - 重启后是否会进入不同控制流
  - 部分写入、部分副作用和版本变更是否被覆盖
- 最低验收标准：
  - 有状态一致性验证
  - 有损坏/过期状态的处理策略
  - 有 replay 或 resume 行为证据

<a id="failure-drill-patterns-for-distributed-platforms"></a>
### `Failure drill patterns for distributed platforms`

- 适用场景：
  - 部署、告警、回滚
  - 端到端 drill
  - 平台整合演练
- 重点追问：
  - drill 是不是覆盖真实用户旅程，而不是玩具路径
  - operator 在收到告警后到底执行哪些动作
  - 恢复路径是不是可以复现，而不是靠口头经验
- 最低验收标准：
  - drill 场景清楚
  - failure signal 清楚
  - recovery readiness 有证据

<a id="tool-calling-failure-patterns"></a>
### `Tool calling failure patterns`

- 适用场景：
  - 单 Agent / 多 Tool
  - Tool Gateway
  - MCP tool execution
- 重点追问：
  - tool timeout、schema mismatch、tool crash 时系统怎么表现
  - tool failure 会不会把整条链直接拖崩
  - fallback 和 retry 的责任是否清楚
- 最低验收标准：
  - 至少一条 tool failure path 被验证
  - fallback 或 fail-closed 行为被说明
  - operator 能从日志/trace 理解失败原因

## 治理与风险类

<a id="risk-review-checklist"></a>
### `Risk review checklist`

- 适用场景：
  - 学期复盘
  - 复杂系统切换阶段
  - roadmap 准备
- 重点追问：
  - 当前最大的技术债和架构风险是什么
  - 哪些问题已经有证据证明，哪些还停留在猜测
  - 如果进入下一阶段，最先会被放大的缺口是什么
- 最低验收标准：
  - 风险被排序
  - 风险和证据被关联
  - 下一步动作不是空话

<a id="security-review-checklist"></a>
### `Security review checklist`

- 适用场景：
  - threat model
  - policy engine
  - adversarial test
- 重点追问：
  - 高风险动作是否被完整纳入模型
  - 控制是在执行前、执行中还是执行后生效
  - 拦截与误伤是否都有证据
- 最低验收标准：
  - 有 threat model
  - 有 enforceable control
  - 有 adversarial evidence

<a id="authorization-hook-patterns"></a>
### `Authorization hook patterns`

- 适用场景：
  - Tool auth
  - Approval path
  - 敏感动作控制
- 重点追问：
  - 哪些动作需要显式授权
  - 授权失败时系统返回什么
  - 授权决策能否被审计和复盘
- 最低验收标准：
  - actor / action / resource 被识别
  - deny path 明确
  - 审计证据完整

<a id="compliance-style-documentation-structure"></a>
### `Compliance-style documentation structure`

- 适用场景：
  - governance docs
  - audit package
  - 年度答辩材料
- 重点追问：
  - 文档有没有同时包含目标、控制、证据和残余风险
  - reader 是 operator、reviewer 还是 committee
  - 文档是否能支持事后 reconstruction
- 最低验收标准：
  - 结构稳定
  - 证据可追
  - 风险不被粉饰

## 性能与最终验收类

<a id="benchmark-decision-criteria"></a>
### `Benchmark decision criteria`

- 适用场景：
  - baseline
  - optimization
  - capacity plan
- 重点追问：
  - workload、环境和指标口径是否一致
  - before / after 是否真的可比
  - 数字变化是否对应可解释的瓶颈变化
- 最低验收标准：
  - 有 benchmark contract
  - 有 baseline
  - 有解释而不只是数字

<a id="final-review-checklist"></a>
### `Final review checklist`

- 适用场景：
  - 最终 benchmark 周
  - 年度答辩周
  - 平台整合最终评审
- 重点追问：
  - 这套系统故事是否同时覆盖目标、设计、证据和限制
  - 哪些 claim 有证据，哪些只是意向
  - 如果删掉一半材料，最值得保留的是哪几组证据
- 最低验收标准：
  - 关键主张都有证据
  - trade-off 和限制被诚实说明
  - 下一步路线不是空泛口号

<a id="production-readiness-checklist"></a>
### `Production-readiness checklist`

- 适用场景：
  - 部署前检查
  - production-like 环境验证
- 重点追问：
  - 启动、探针、配置、告警、回滚和 runbook 是否齐备
  - 出现故障时有没有明确 responder action
  - 配置和依赖变更是否能被安全发布
- 最低验收标准：
  - 有 runbook
  - 有 alert / rollback 规则
  - 有 smoke / drill 证据

## 实现观察模式

<a id="go-concurrency-patterns"></a>
### `Go concurrency patterns`

- 适用场景：
  - worker loop
  - goroutine / channel 协调
- 重点追问：
  - goroutine 生命周期谁负责结束
  - 是否存在泄漏、阻塞或重复消费风险
  - shutdown 与 context cancellation 是否配合

<a id="state-persistence-patterns"></a>
### `State persistence patterns`

- 适用场景：
  - checkpoint store
  - workflow state 持久化
- 重点追问：
  - 哪些状态必须持久化
  - 状态版本变化时如何兼容
  - 部分写入如何恢复

<a id="review-criteria"></a>
### `Review criteria`

- 适用场景：
  - 做文档评审、接口评审或代码评审前
- 最少要问：
  - 这项工作真正降低了什么复杂度
  - 增加了哪些未来维护成本
  - 证据是否足以支持“已经完成”的说法

## 最后一句

如果你读完一个 checklist 之后，还是只会复述条目而不会拿它挑自己的设计，那说明你还没有真正用起来。
