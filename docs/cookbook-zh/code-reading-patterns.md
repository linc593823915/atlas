# 代码阅读 Patterns

最后更新：2026-06-30

## 这份文档怎么用

在 day 文档里，有一类阅读提示长这样：

- `net/http server startup flow and Cobra command initialization patterns`
- `existing worker examples and context cancellation patterns`
- `channel coordination patterns and worker shutdown logic`

这些词不是仓库里的真实文件名。

它们本质上是在提醒你：

“读代码时，别只看函数名，要看它背后的实现模式。”

这份文档就是把这些模式提示变成可直接阅读的中文说明。

## 启动与入口模式

<a id="net-http-server-startup-flow-and-cobra-command-initialization-patterns"></a>
### `net/http server startup flow and Cobra command initialization patterns`

- 适用主题：
  - CLI / HTTP 启动
  - 服务外壳
  - 第一月入口层设计
- 读法：
  - 先看 [cmd/atlas/main.go](../../cmd/atlas/main.go)
  - 再看 [internal/app/root.go](../../internal/app/root.go)
  - 再看 [internal/app/serve.go](../../internal/app/serve.go)
- 最该盯的点：
  - `main` 是否保持足够薄
  - Cobra command tree 是否只是入口组织，而不是依赖初始化场所
  - 真正的 start / stop 生命周期是否被收进 bootstrap / runtime
- 常见误判：
  - 把“命令能运行”误认为“入口层边界已经合理”
  - 把 CLI 逻辑和服务启动逻辑混在一起

## 并发与后台执行模式

<a id="existing-worker-examples-and-context-cancellation-patterns"></a>
### `existing worker examples and context cancellation patterns`

- 适用主题：
  - worker loop
  - queue / job 执行
  - graceful shutdown
- 最该盯的点：
  - 谁创建 context，谁负责取消
  - worker 停止时是在停拉新、排在途，还是直接退出
  - 重试、超时和 cancel 会不会让同一个 job 进入不一致状态
- 学习动作：
  - 把它和 [internal/bootstrap/bootstrap.go](../../internal/bootstrap/bootstrap.go)、[internal/runtime/runtime.go](../../internal/runtime/runtime.go) 一起看，思考未来 worker 真接进 runtime 时 owner 应该是谁

<a id="channel-coordination-patterns-and-worker-shutdown-logic"></a>
### `channel coordination patterns and worker shutdown logic`

- 适用主题：
  - goroutine 协调
  - worker pool
  - shutdown path
- 最该盯的点：
  - channel 是在表达 ownership，还是只在传值
  - shutdown 时 close 顺序和 drain 顺序有没有被定义
  - goroutine 是否可能泄漏或阻塞
- 常见误判：
  - “用了 channel”就等于并发设计合理
  - “有 context”就等于 shutdown 一定正确

<a id="worker-pool-implementation-notes"></a>
### `Worker pool implementation notes`

- 适用主题：
  - 多 worker 执行模型
- 最该盯的点：
  - 并发度是固定的、可调的，还是隐式长出来的
  - 失败 worker 如何被替换
  - 积压增长时系统是否有 backpressure 策略

## 重试、超时与失败面模式

<a id="idempotency-and-retry-design-notes"></a>
### `Idempotency and retry design notes`

- 适用主题：
  - Retry
  - Idempotency
  - Dead-letter
- 最该盯的点：
  - 重试是在保护暂时性错误，还是在放大不可恢复错误
  - 重复执行是否会放大副作用
  - dead-letter 到底是隔离区，还是垃圾桶

<a id="tool-calling-failure-patterns"></a>
### `Tool calling failure patterns`

- 适用主题：
  - 单 Agent / 多 Tool
  - Tool Gateway
  - MCP tools
- 最该盯的点：
  - tool failure 后是 retry、fallback 还是 fail-closed
  - schema mismatch 时谁负责兜住
  - operator 能不能从结果、日志或 trace 理解失败

<a id="context-deadline-handling-and-error-classification-flow"></a>
### `context deadline handling and error classification flow`

- 适用主题：
  - timeout / retry
  - request cancellation
  - tool failure handling
- 最该盯的点：
  - deadline 在哪一层被设定
  - cancel 后错误会不会被错误归类成业务失败
  - operator 能不能从错误面区分 timeout、cancel 和真实下游错误

<a id="database-sql-connection-management-and-slog-handler-internals"></a>
### `database/sql connection management and slog handler internals`

- 适用主题：
  - 配置、日志、数据库与缓存接线
- 最该盯的点：
  - db / cache 连接生命周期是不是被 runtime 正确拥有
  - slog handler 是否已经稳定表达 level、format 和 source
  - 配置变化是否会正确影响连接和日志行为

<a id="timeout-middleware-examples"></a>
### `Timeout middleware examples`

- 适用主题：
  - timeout policy
  - auth / audit hook
  - request context propagation
- 最该盯的点：
  - timeout budget 的 owner 是谁
  - cancel 信号是否真的能传到下游
  - timeout 后系统返回的是“中断了”，还是“悄悄坏了”

<a id="request-context-propagation-and-logging-enrichment-flow"></a>
### `request context propagation and logging enrichment flow`

- 适用主题：
  - request-scoped logging
  - trace / audit 字段传播
- 最该盯的点：
  - context 中到底放哪些字段
  - logger enrichment 是统一做，还是各层随手补
  - trace id / actor / approval state 有没有可能传播丢失

## 状态与恢复模式

<a id="state-persistence-patterns"></a>
### `State persistence patterns`

- 适用主题：
  - checkpoint
  - workflow state
  - state store
- 最该盯的点：
  - 哪些状态必须持久化
  - 部分写入之后系统还能不能恢复
  - 状态版本变更时兼容策略是什么

<a id="resume-semantics-for-long-running-jobs"></a>
### `Resume semantics for long-running jobs`

- 适用主题：
  - pause / resume
  - replay
- 最该盯的点：
  - resume cursor 是否唯一
  - 恢复后控制流是否和预期一致
  - human intervention 之后 provenance 是否保留

<a id="workflow-engine-state-transitions-and-storage-adapters"></a>
### `workflow engine state transitions and storage adapters`

- 适用主题：
  - state machine
  - storage adapter
- 最该盯的点：
  - 状态迁移是否合法
  - adapter 是否只是存储细节，还是正在偷偷改变语义
  - replay / resume 是否依赖某个 adapter 的偶然行为

<a id="review-notes-for-workflow-engines"></a>
### `Review notes for workflow engines`

- 适用主题：
  - 工作流 review
- 最该盯的点：
  - 方案是否真的能解释恢复路径
  - 状态和副作用边界是否清楚
  - operator 是否能在不找作者的前提下做判断

## 评估、治理与平台模式

<a id="operational-dashboard-examples"></a>
### `Operational dashboard examples`

- 适用主题：
  - eval / trace 结果解读
  - operator dashboard
- 最该盯的点：
  - dashboard 是否真的能支持决策，而不是只堆指标
  - 一个异常信号能不能被追到具体 trace / sample / gate

<a id="atlas-eval-scripts-trace-outputs-and-release-checklist"></a>
### `Atlas eval scripts, trace outputs, and release checklist`

- 适用主题：
  - gate
  - release policy
  - 最终评审
- 最该盯的点：
  - 评估脚本、trace 结果和 gate 规则是否能连成一条因果链
  - fail 之后到底是回滚、修数据还是调阈值

<a id="openai-sdk-response-parsing-and-schema-validation-code"></a>
### `OpenAI SDK response parsing and schema validation code`

- 适用主题：
  - structured output
  - validator
- 最该盯的点：
  - 解析失败时系统怎样从 SDK 返回值进入本地 contract
  - schema validation 是在进入后续逻辑前发生，还是事后补救
  - 哪些字段还是模型输出，哪些字段已经是系统结构化结果

<a id="integration-architecture-notes"></a>
### `Integration architecture notes`

- 适用主题：
  - config / tool / MCP integration
  - platform boundary
- 最该盯的点：
  - 共享 contract 是否只有一份 truth
  - 子系统整合是不是靠手工 glue 在维持

<a id="system-test-design-notes"></a>
### `System test design notes`

- 适用主题：
  - end-to-end drills
  - failure recovery
- 最该盯的点：
  - 测试旅程是否覆盖真实用户路径
  - 失败演练是否包含真正的 operator decision

<a id="atlas-bottleneck-traces-and-queue-metrics"></a>
### `Atlas bottleneck traces and queue metrics`

- 适用主题：
  - optimization loop
  - capacity planning
- 最该盯的点：
  - 瓶颈到底在 model、tool、queue 还是 infra
  - 指标变化是否能解释优化收益

<a id="final-atlas-performance-traces-eval-outputs-and-ops-notes"></a>
### `final Atlas performance traces, eval outputs, and ops notes`

- 适用主题：
  - final review
  - graduation defense
- 最该盯的点：
  - 最终证据是否已经从“很多材料”压成“几个最强证据”
  - 性能、质量和 operator 视角是否能讲成一个故事

## 一句话提醒

看到 `... patterns` 这类词时，不要问：

- “这里的正确答案是什么？”

要问：

- “这个模式在保护哪条边界？”
