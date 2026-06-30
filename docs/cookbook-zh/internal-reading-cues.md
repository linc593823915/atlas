# 仓库内阅读线索

最后更新：2026-06-30

## 这份文档怎么用

day 文档里还有一类提示，不是外部资料，也不是仓库里的真实文件名。

它们更像是“老师在旁边提醒你该关注什么”：

- `Atlas bootstrap and entrypoint code plus Go project layout examples`
- `current single-agent flow and possible handoff points`
- `Atlas workflow, deployment, and recovery paths end to end`

这类提示真正想表达的是：

“去仓库里找这一组线索，一起看，而不是单看一个文件。”

这份文档就是把这些线索拆开说清楚。

## 启动与入口线索

<a id="atlas-bootstrap-and-entrypoint-code-plus-go-project-layout-examples"></a>
### `Atlas bootstrap and entrypoint code plus Go project layout examples`

- 适用阶段：
  - 第一个月
  - 服务骨架
- 建议一起看：
  - [cmd/atlas/main.go](../../cmd/atlas/main.go)
  - [internal/app/root.go](../../internal/app/root.go)
  - [internal/app/serve.go](../../internal/app/serve.go)
  - [internal/bootstrap/bootstrap.go](../../internal/bootstrap/bootstrap.go)
- 你真正要回答的问题：
  - Atlas 的入口层、command 层、bootstrap 层和 runtime 层是怎么被切开的
  - 哪些逻辑未来不应该回流到 `main` 和 `serve`

<a id="current-atlas-cli-flow-and-openai-sdk-request-models"></a>
### `current Atlas CLI flow and OpenAI SDK request models`

- 适用阶段：
  - 第三个月第一个 Agent
- 你真正要回答的问题：
  - 现在的 CLI flow 如果以后接模型请求，会卡在哪些边界
  - 请求构造、tool 调用、输出解析分别更适合落在哪层
- 阅读姿势：
  - 不是去找“完整 Agent 代码”，而是去看当前入口壳是否已经为请求主路径留好位置

<a id="openai-sdk-run-loop-and-atlas-runtime-bootstrap"></a>
### `OpenAI SDK run loop and Atlas runtime bootstrap`

- 适用阶段：
  - agent definition
  - runtime bootstrap
- 你真正要回答的问题：
  - SDK run loop 和 Atlas 现有 runtime 壳之间，谁该拥有请求生命周期
  - 哪些能力该留在 SDK 边，哪些该收回 Atlas runtime

<a id="current-atlas-tests-makefile-and-container-startup-path"></a>
### `current Atlas tests, Makefile, and container startup path`

- 适用阶段：
  - Docker / CI / 第一月复盘
- 你真正要回答的问题：
  - 本地运行、容器运行和测试运行是不是同一条可解释主路径
  - Makefile、测试和容器启动是否在讲同一个系统，而不是三套平行世界

<a id="atlas-config-and-logger-modules"></a>
### `Atlas config and logger modules`

- 适用阶段：
  - 配置、日志、基础设施边界
- 建议一起看：
  - [internal/config/manager.go](../../internal/config/manager.go)
  - [internal/config/loader.go](../../internal/config/loader.go)
  - [internal/logger/logger.go](../../internal/logger/logger.go)
  - [internal/logger/options.go](../../internal/logger/options.go)
- 你真正要回答的问题：
  - 配置和日志是否都已经有单一入口
  - 未来 trace / audit 接入时，当前 logger 抽象是否足够承接

<a id="atlas-module-boundaries-created-in-semester-1"></a>
### `Atlas module boundaries created in Semester 1`

- 适用阶段：
  - 第一学期复盘
- 你真正要回答的问题：
  - 第一学期真正留下了哪些稳定边界
  - 哪些边界只是薄壳，哪些边界已经被测试或文档保护

<a id="your-semester-1-notes"></a>
### `Your Semester 1 notes`

- 适用阶段：
  - 第一学期复盘
  - 第二学期准备
- 你真正要回答的问题：
  - 哪些能力已经从“练习产物”变成后续可以复用的系统边界
  - 哪些问题会影响接口、协议和运行时设计
- 阅读姿势：
  - 把第一学期的 week review、月度总结和 issue 证据放在一起看，不要只重读最后一页复盘

<a id="your-semester-2-notes"></a>
### `Your Semester 2 notes`

- 适用阶段：
  - 第二学期复盘
  - 第三学期 durable workflow 准备
- 你真正要回答的问题：
  - Tool Gateway、MCP、Agent runtime 之间是否已经形成统一边界语言
  - 哪些协议或 contract 还没有被测试保护
- 阅读姿势：
  - 重点看接口、schema、错误模型和兼容性证据，而不是只看实现进度

<a id="your-semester-3-notes"></a>
### `Your Semester 3 notes`

- 适用阶段：
  - 第三学期复盘
  - 第四学期平台化准备
- 你真正要回答的问题：
  - workflow、eval、governance 是否已经能互相解释
  - 哪些恢复、审批或审计路径仍然只是文档承诺
- 阅读姿势：
  - 把恢复 drill、eval 结果、governance review 和 trace 证据放在同一张判断表里

## Tool / Agent / MCP 线索

<a id="current-single-agent-flow-and-possible-handoff-points"></a>
### `current single-agent flow and possible handoff points`

- 适用阶段：
  - 多 Agent runtime
- 你真正要回答的问题：
  - 现有单 Agent 链路如果升级成多 Agent，handoff 会发生在什么节点
  - 哪些上下文必须显式带走，哪些上下文不该扩散

<a id="atlas-tool-interfaces-and-existing-command-dispatch"></a>
### `Atlas tool interfaces and existing command dispatch`

- 适用阶段：
  - tool inventory
  - gateway 入口
- 你真正要回答的问题：
  - 当前 command / tool dispatch 里，哪些是未来统一 gateway 应接管的行为
  - 哪些接口今天还是隐式约定

<a id="atlas-agent-tests-and-response-validation-logic"></a>
### `Atlas agent tests and response validation logic`

- 适用阶段：
  - 单 Agent review
  - structured output / validator
- 你真正要回答的问题：
  - 测试有没有真的保护 response contract
  - validator 失败时系统是怎么收口的

<a id="atlas-tool-mcp-and-agent-boundaries-as-a-whole"></a>
### `Atlas tool, MCP, and agent boundaries as a whole`

- 适用阶段：
  - 第二学期复盘
- 你真正要回答的问题：
  - tool、协议面和 agent runtime 之间是不是已经形成一致的边界语言
  - 哪个边界最容易在下一学期被打穿

<a id="atlas-tool-wrappers-and-sdk-error-surfaces"></a>
### `Atlas tool wrappers and SDK error surfaces`

- 适用阶段：
  - 多 Tool routing
  - tool failure handling
- 你真正要回答的问题：
  - tool wrapper 是否把下游错误转成了统一语义
  - 失败后 operator 能不能从 error surface 看出责任边界

<a id="mcp-message-shapes-and-atlas-surface-selection-logic"></a>
### `MCP message shapes and Atlas surface selection logic`

- 适用阶段：
  - MCP capability / compatibility
- 你真正要回答的问题：
  - message shape 是否已经稳定到可以被外部 client 消费
  - Atlas 当前 surface 选择逻辑是不是还在泄漏内部实现

<a id="atlas-mcp-handlers-and-serialization-boundaries"></a>
### `Atlas MCP handlers and serialization boundaries`

- 适用阶段：
  - MCP capability
  - protocol review
- 你真正要回答的问题：
  - handler 层是否只做协议转换
  - 序列化边界是否把内部实现细节泄漏给 client

<a id="atlas-tool-test-layout-and-timeout-behavior"></a>
### `Atlas tool test layout and timeout behavior`

- 适用阶段：
  - gateway 测试
  - tool timeout / negative path
- 你真正要回答的问题：
  - 测试布局是不是在保护 contract，而不是只保护实现
  - timeout 之后系统是如何 fail 的

<a id="schema-timeout-audit-together"></a>
### `为什么 schema、timeout、audit 需要一起考虑`

- 适用阶段：
  - Tool Gateway
  - MCP capability
  - tool auth
- 你真正要回答的问题：
  - schema 保护输入输出 contract
  - timeout 保护资源和调用方预算
  - audit 保护事后解释和治理
  - 三者是否在同一条 tool invocation path 中对齐
- 阅读姿势：
  - 从一次 tool call 出发，标出校验、执行、失败、审计分别在哪里发生

<a id="unified-tool-entry-value"></a>
### `统一工具入口的价值`

- 适用阶段：
  - Tool Gateway RFC
  - 多 Tool routing
- 你真正要回答的问题：
  - 统一入口是否减少 schema drift
  - timeout、auth、audit 是否只有一套语义
  - Agent 和 MCP 是否复用同一个 tool contract
- 阅读姿势：
  - 对比“直接调用 tool”和“通过 gateway 调用 tool”时，哪些风险被集中处理

<a id="tool-error-isolation"></a>
### `工具错误如何不拖崩整条链`

- 适用阶段：
  - tool failure handling
  - fallback review
- 你真正要回答的问题：
  - tool error 如何分类
  - 哪些错误可以 fallback
  - 哪些错误必须 fail closed
  - 原始错误是否保留给 operator
- 阅读姿势：
  - 不要只检查 error != nil；要追错误如何进入输出、日志和 trace

<a id="multi-tool-routing-reading"></a>
### `多 Tool 路由`

- 适用阶段：
  - 多 Tool routing
  - Agent review
- 你真正要回答的问题：
  - routing 输入是什么
  - 候选 tool 如何筛选
  - schema mismatch 如何处理
  - fallback 和 deny path 是否清楚
- 阅读姿势：
  - 把 routing 当成系统决策点，而不是 prompt 技巧

<a id="fallback-thinking"></a>
### `fallback 思维`

- 适用阶段：
  - Agent reliability
  - tool failure handling
- 你真正要回答的问题：
  - fallback 后输出是否仍满足 contract
  - fallback 是否隐藏了真正错误
  - operator 能否知道发生过 fallback
- 阅读姿势：
  - 每个 fallback 都要同时写触发条件、输出形态和证据字段

<a id="structured-output-engineering-value"></a>
### `structured output 的工程价值`

- 适用阶段：
  - structured output
  - schema validation
- 你真正要回答的问题：
  - structured output 降低了哪类解析不确定性
  - schema 失败时系统如何处理
  - 调用方是否可以不读自然语言也能消费结果
- 阅读姿势：
  - 把结构化输出和 output contract、validator、negative tests 一起看

<a id="first-end-to-end-tool-path"></a>
### `第一条 end-to-end tool path`

- 适用阶段：
  - 单 Agent tool use
  - Tool Gateway 初版
- 你真正要回答的问题：
  - 从请求进入到 tool 执行再到输出解析的主路径是否完整
  - 每个边界是否有最小测试
  - 失败路径是否至少有一个可复现样本
- 阅读姿势：
  - 按时间顺序写出 request -> model/agent -> tool -> result -> validator -> response

<a id="atlas-orchestration-tests-and-approval-path-coverage"></a>
### `Atlas orchestration tests and approval path coverage`

- 适用阶段：
  - 多 Agent 测试与复盘
- 你真正要回答的问题：
  - 编排测试是不是已经覆盖 handoff、fallback 和 approval deny path
  - 哪条责任链仍然只能靠人工解释

<a id="current-approval-flow-code-and-trace-annotations"></a>
### `current approval flow code and trace annotations`

- 适用阶段：
  - handoff / approval / audit
- 你真正要回答的问题：
  - 哪些动作需要审批
  - 审批决定如何被留在 trace / audit 里
  - 拒绝路径是不是和允许路径一样清楚

<a id="existing-protocol-server-bootstraps-and-request-dispatch-flow"></a>
### `existing protocol server bootstraps and request dispatch flow`

- 适用阶段：
  - MCP bootstrap
- 你真正要回答的问题：
  - session lifecycle 和 runtime lifecycle 如何对齐
  - capability dispatch 是不是还在依赖内部实现细节

<a id="atlas-config-tool-and-mcp-modules-together"></a>
### `Atlas config, tool, and MCP modules together`

- 适用阶段：
  - 配置中心、Tool Gateway 与 MCP 整合
- 你真正要回答的问题：
  - 同一个 capability / config / contract 是否只有一份 truth
  - 外部 surface 是否和内部 gateway 使用了不同语义

<a id="validation-library-internals-and-current-atlas-config-loading"></a>
### `validation library internals and current Atlas config loading`

- 适用阶段：
  - schema registry
  - config / validation review
- 你真正要回答的问题：
  - validation 是在正确位置发生，还是只是结果最后才补一层
  - 当前 Atlas config loading 路径是否已经能稳定承载 schema 扩展

<a id="atlas-policy-engine-docs-and-test-evidence"></a>
### `Atlas policy engine docs and test evidence`

- 适用阶段：
  - governance review
- 你真正要回答的问题：
  - 文档里的 policy 是不是已经变成可执行控制
  - adversarial tests 是否真的能证明这些控制成立

<a id="atlas-eval-pipeline-approval-hooks-and-audit-event-shapes"></a>
### `Atlas eval pipeline, approval hooks, and audit event shapes`

- 适用阶段：
  - eval / approval / audit integration
- 你真正要回答的问题：
  - 一条高风险决策链能不能同时看到质量信号、审批状态和审计事件
  - 这些 shape 是否已经可以互相对齐

## 工作流与治理线索

<a id="atlas-workflow-tests-and-persistence-adapters"></a>
### `Atlas workflow tests and persistence adapters`

- 适用阶段：
  - durable workflow review
- 你真正要回答的问题：
  - 状态测试在保护什么不变量
  - persistence adapter 是在保存状态，还是在偷偷改变恢复语义

<a id="atlas-worker-tests-and-logging-around-retries-and-dead-letters"></a>
### `Atlas worker tests and logging around retries and dead letters`

- 适用阶段：
  - worker 复盘
  - retry / dead-letter review
- 你真正要回答的问题：
  - 测试是不是只证明“能跑”，还是也证明了重复执行、超时和 dead-letter 语义
  - 日志能不能支持 operator 排查失败路径

<a id="worker-consume-and-graceful-shutdown-path"></a>
### `worker 如何消费` / `worker 如何被优雅关闭`

- 适用阶段：
  - worker 主循环
  - shutdown review
- 你真正要回答的问题：
  - worker claim 任务后，owner 是否清楚
  - shutdown 时是停止拉新、等待在途、释放 lease，还是直接丢弃上下文
  - 日志是否能解释“为什么这个任务完成、重试或被释放”
- 阅读姿势：
  - 把 queue contract、runtime lifecycle 和 shutdown drill 一起看，单看 worker loop 不够

<a id="task-queue-entry-path"></a>
### `任务如何进入队列`

- 适用阶段：
  - Job Runner
  - queue abstraction
- 你真正要回答的问题：
  - enqueue 入口是否验证输入和幂等 key
  - job 初始状态是否唯一
  - enqueue 成功后调用方能拿到什么可追踪证据
- 阅读姿势：
  - 从 API/command 入口一路追到 queue store，记录每次状态变化的 owner

<a id="retryable-and-terminal-error-path"></a>
### `什么错误会重试` / `什么错误必须直接失败`

- 适用阶段：
  - retry policy
  - dead-letter
- 你真正要回答的问题：
  - retryable 与 terminal 的判断是否由错误类型、状态还是 operator 决策驱动
  - 重试是否可能放大副作用
  - terminal failure 是否进入可排查的位置
- 阅读姿势：
  - 先写错误分类表，再回到 worker 和 queue 实现找证据

<a id="background-execution-boundary-first"></a>
### `为什么后台执行必须先定义边界再写实现`

- 适用阶段：
  - Job Runner RFC
  - worker loop
- 你真正要回答的问题：
  - job owner 是谁
  - 状态迁移在哪里发生
  - retry、timeout、dead-letter 由哪一层负责
  - operator 能否从状态和日志判断下一步
- 阅读姿势：
  - 先画状态图和 queue contract，再看实现；不要从 goroutine 写起

<a id="job-state-and-recoverability-questions"></a>
### `任务有哪些状态` / `哪些状态可恢复`

- 适用阶段：
  - job state model
  - recovery semantics
- 你真正要回答的问题：
  - queued、claimed、running、succeeded、failed、dead-letter 等状态是否互斥
  - 哪些状态能安全重试
  - 哪些状态需要 operator 决策
- 阅读姿势：
  - 把状态图、错误模型和 dead-letter 行为放在同一个表里

<a id="enqueue-and-execution-main-path"></a>
### `入队与执行主路径`

- 适用阶段：
  - worker 主路径
  - queue abstraction
- 你真正要回答的问题：
  - 从 enqueue 到 execute 再到 ack/release 的责任链是否连续
  - 每一步失败时状态如何变化
  - 哪些日志或 trace 能证明主路径走通
- 阅读姿势：
  - 逐步标出 caller、queue、worker、runtime 的 owner，不要只看函数调用

<a id="retry-boundary-and-idempotency-purpose"></a>
### `retry 的边界` / `timeout 的责任` / `幂等的真正用途`

- 适用阶段：
  - reliability review
  - worker failure model
- 你真正要回答的问题：
  - retry 解决的是 transient failure，还是在重复不可恢复错误
  - timeout 是调用方预算、worker 预算还是下游预算
  - 幂等保护的是哪种重复副作用
- 阅读姿势：
  - 把 retry、timeout、idempotency 三者放在同一条失败路径里解释

<a id="idempotency-risk-note"></a>
### `重复任务为什么不会造成灾难`

- 适用阶段：
  - idempotency
  - worker reliability review
- 你真正要回答的问题：
  - 重复执行的最坏副作用是什么
  - 幂等 key 或状态检查在哪里生效
  - 重复执行是否会被 trace 或日志发现
- 阅读姿势：
  - 不要只说“幂等”。要指出哪一步让重复执行变成可控结果

<a id="current-atlas-runtime-state-handling-and-failure-recovery-gaps"></a>
### `current Atlas runtime state handling and failure recovery gaps`

- 适用阶段：
  - durable workflow
  - 平台 recovery
- 你真正要回答的问题：
  - 当前 runtime 对失败和恢复的表达缺口在哪里
  - 哪些 gap 必须在恢复 drill 前先补上

<a id="current-resume-code-and-decision-checkpoints"></a>
### `current resume code and decision checkpoints`

- 适用阶段：
  - resume / replay review
- 你真正要回答的问题：
  - 恢复后的下一步是不是唯一
  - decision checkpoint 是否真的能支持 replay

<a id="minimal-recoverable-flow-and-state-recovery-path"></a>
### `最小可恢复流程` / `状态恢复路径`

- 适用阶段：
  - durable workflow
  - recovery drill
- 你真正要回答的问题：
  - 最小恢复场景是什么
  - checkpoint 保存了哪些状态
  - 恢复后下一步是否唯一
  - 哪些副作用明确不能重放
- 阅读姿势：
  - 把 workflow state、checkpoint store、resume tests 和 recovery 文档放在一起看

<a id="workflow-test-and-recovery-doc-package"></a>
### `工作流测试集` / `恢复文档`

- 适用阶段：
  - workflow review
  - 第七个月复盘
- 你真正要回答的问题：
  - 测试集是否覆盖 pause、resume、replay、human intervention 和失败恢复
  - 恢复文档是否能让 operator 在没有作者上下文时执行
- 阅读姿势：
  - 先看测试命名，再看文档是否解释同一组 failure mode

<a id="pause-resume-semantics"></a>
### `pause/resume` / `pause/resume 语义`

- 适用阶段：
  - durable workflow
  - human intervention
- 你真正要回答的问题：
  - pause 停在哪里
  - resume 用什么 cursor
  - 人工输入如何合并
  - 哪些副作用不能因为 resume 重复发生
- 阅读姿势：
  - 从暂停点倒推 checkpoint 内容，而不是先设计存储表

<a id="human-intervention-reading"></a>
### `human intervention`

- 适用阶段：
  - approval
  - recovery
- 你真正要回答的问题：
  - 人在 workflow 中提供的是审批、补充输入还是纠错
  - 人工动作是否有 provenance
  - merge 后系统是否仍能 replay 或 audit
- 阅读姿势：
  - 把人的动作当作系统事件，不要当作系统外备注

<a id="atlas-workflow-eval-and-governance-boundaries-together"></a>
### `Atlas workflow, eval, and governance boundaries together`

- 适用阶段：
  - 第三学期复盘
- 你真正要回答的问题：
  - state、quality、governance 这三套边界是否真的能讲成一个系统故事
  - 哪个 failure mode 横跨了三者

<a id="atlas-audit-outputs-and-rejection-paths"></a>
### `Atlas audit outputs and rejection paths`

- 适用阶段：
  - policy / audit / misuse review
- 你真正要回答的问题：
  - 审计记录是不是只在记录结果，还是也在记录决策原因
  - deny path 是否和 happy path 一样可解释

<a id="atlas-high-risk-paths-tool-boundaries-and-approval-actions"></a>
### `Atlas high-risk paths, tool boundaries, and approval actions`

- 适用阶段：
  - threat model
  - governance RFC
- 你真正要回答的问题：
  - 高风险动作到底发生在哪些路径
  - 哪些路径应该被 policy、approval、audit 同时覆盖

<a id="atlas-action-execution-path-and-event-logging-fields"></a>
### `Atlas action execution path and event logging fields`

- 适用阶段：
  - policy engine
  - 审计事件
- 你真正要回答的问题：
  - 动作执行路径里哪些字段必须被日志 / 审计保留
  - 事件字段是否足以支撑事后 reconstruction

<a id="data-leakage-and-prompt-injection-paths"></a>
### `data leakage` / `prompt injection`

- 适用阶段：
  - threat model
  - misuse tests
- 你真正要回答的问题：
  - 敏感数据可能从输入、tool、trace、memory、日志中的哪一处泄漏
  - prompt injection 是否能影响 tool 调用或审批决策
  - 当前 guardrail 是阻断、降级还是只记录
- 阅读姿势：
  - 把攻击路径写成可复现测试，而不是只列风险名称

## 评估、平台与性能线索

<a id="atlas-eval-scripts-trace-outputs-and-release-checklist"></a>
### `Atlas eval scripts, trace outputs, and release checklist`

- 适用阶段：
  - eval review
  - gate 评审
- 你真正要回答的问题：
  - score、trace 和 gate 是否已经连成一条决策链
  - release checklist 是不是只在列项，没有真正约束动作

<a id="platform-level-eval-chain"></a>
### `平台级评估链路`

- 适用阶段：
  - release gate
  - platform review
- 你真正要回答的问题：
  - eval 样本、grader、trace、gate 和 release decision 是否能连成一条证据链
  - 哪些质量退化会阻断发布
- 阅读姿势：
  - 不要分开看 eval 和 release。把一次失败样本追到最终是否允许发布

<a id="semester-outcomes-and-next-prep"></a>
### `第二学期成果` / `第三学期成果` / `第三学期准备点` / `第四学期准备点`

- 适用阶段：
  - 学期复盘
  - 下一阶段计划
- 你真正要回答的问题：
  - 本阶段留下了哪些可复用 contract
  - 哪些证据足以支持进入下一阶段
  - 下一阶段第一条主线为什么由当前风险推出
- 阅读姿势：
  - 先列证据，再写结论；不要按时间线重述做过什么

<a id="semester-two-interface-runtime-shift"></a>
### `第二学期为什么必须转向接口和运行时边界`

- 适用阶段：
  - 第一学期末复盘
  - 第二学期准备
- 你真正要回答的问题：
  - 第一学期的服务骨架和 worker 练习留下了哪些边界问题
  - 为什么下一阶段不能继续只做功能
  - 哪些接口和 runtime contract 必须先冻结
- 阅读姿势：
  - 用第一学期的失败模式来解释第二学期主题，而不是只看课程安排

<a id="capabilities-and-semester-two-risk"></a>
### `哪些能力已经建立` / `哪些问题会影响第二学期`

- 适用阶段：
  - 第一学期复盘
- 你真正要回答的问题：
  - 哪些能力已经被文档、测试或 review 证明
  - 哪些问题只是暂时没爆
  - 第二学期最先会被这些问题影响的是哪个接口或 runtime 边界
- 阅读姿势：
  - 把能力和风险都绑定到证据，不要写成自我评价

<a id="remaining-risk-and-legacy-risk"></a>
### `剩余风险` / `遗留风险`

- 适用阶段：
  - 复盘
  - review packet
- 你真正要回答的问题：
  - 风险为什么还没消失
  - 它影响下一阶段哪条主线
  - 需要什么证据才能关闭
- 阅读姿势：
  - 剩余风险要带下一步动作，否则只是免责声明

<a id="final-benchmark-package"></a>
### `最终 benchmark 包`

- 适用阶段：
  - 最终 benchmark
  - 毕业答辩
- 你真正要回答的问题：
  - workload、环境、指标、baseline、before/after 和解释是否齐全
  - benchmark 结论是否能指导容量或成本决策
- 阅读姿势：
  - 先检查 benchmark contract，再看数字；没有 contract 的数字不能当最终证据

<a id="current-atlas-tests-and-identify-missing-regression-signals"></a>
### `current Atlas tests and identify missing regression signals`

- 适用阶段：
  - eval / benchmark / 最终 review
- 你真正要回答的问题：
  - 当前测试覆盖的是功能正确性，还是也覆盖了质量退化信号
  - 哪些 regression signal 现在还缺

<a id="atlas-agent-workflow-hot-paths-and-cache-opportunities"></a>
### `Atlas agent workflow hot paths and cache opportunities`

- 适用阶段：
  - 优化与容量规划
- 你真正要回答的问题：
  - 热路径到底在哪
  - cache 是否真的在缩短关键路径，而不是只在增加复杂度

<a id="atlas-bottleneck-traces-and-queue-metrics"></a>
### `Atlas bottleneck traces and queue metrics`

- 适用阶段：
  - 性能优化
- 你真正要回答的问题：
  - 瓶颈在 model、tool、queue 还是 infra
  - 当前指标是不是能支撑容量判断

<a id="atlas-architecture-docs-and-review-history-as-one-body-of-work"></a>
### `Atlas architecture, docs, and review history as one body of work`

- 适用阶段：
  - 年度复盘
  - graduation defense
- 你真正要回答的问题：
  - 这一年的架构、文档和 review 是否已经能组成一个统一叙事
  - 哪些结论只是周作业，哪些已经成为平台级判断

<a id="atlas-runtime-spans-and-result-classification-points"></a>
### `Atlas runtime spans and result classification points`

- 适用阶段：
  - eval / trace / final review
- 你真正要回答的问题：
  - trace span 是否已经覆盖关键 runtime decision
  - 结果分类点是在模型后、tool 后还是整体 orchestration 后发生

<a id="atlas-hot-path-spans-and-tool-call-latency-breakdown"></a>
<a id="atlas-hot-path-spans-and-tool-call-latency-breakdowns"></a>
### `Atlas hot-path spans and tool-call latency breakdowns`

- 适用阶段：
  - baseline / optimization
- 你真正要回答的问题：
  - 热路径的时间到底花在工具、模型还是编排
  - 哪一段值得先优化

<a id="atlas-workflow-deployment-and-recovery-paths-end-to-end"></a>
### `Atlas workflow, deployment, and recovery paths end to end`

- 适用阶段：
  - 端到端 drill
- 你真正要回答的问题：
  - 一条真实旅程里，workflow、deployment、recovery 是怎样串起来的
  - 当前 drill 是不是只验证模块，不验证整条链

<a id="atlas-startup-and-shutdown-code-for-probe-design"></a>
### `Atlas startup and shutdown code for probe design`

- 适用阶段：
  - probes
  - readiness / liveness
- 你真正要回答的问题：
  - 探针应该对应哪条真实启动 / 关闭语义
  - 当前代码里哪些状态适合拿来表达 readiness，哪些适合表达 liveness

<a id="atlas-process-model-and-identify-deployment-seams"></a>
### `Atlas process model and identify deployment seams`

- 适用阶段：
  - deployment topology
- 你真正要回答的问题：
  - 哪些职责未来应该拆成独立部署单元
  - 当前进程模型已经在哪些位置暴露出 seam

<a id="every-atlas-subsystem-boundary-created-so-far"></a>
### `every Atlas subsystem boundary created so far`

- 适用阶段：
  - 平台整合 RFC
  - 年度答辩
- 你真正要回答的问题：
  - 这一年到底留下了哪些稳定边界
  - 哪些边界是平台级边界，哪些只是模块级边界

<a id="every-major-atlas-subsystem-and-document-real-responsibilities"></a>
### `every major Atlas subsystem and document real responsibilities`

- 适用阶段：
  - 架构文档
  - operator 文档
- 你真正要回答的问题：
  - 每个子系统的真实责任是什么
  - 文档写的是愿景，还是当前真实能力

<a id="atlas-metrics-surfaces-and-current-ci-checks"></a>
### `Atlas metrics surfaces and current CI checks`

- 适用阶段：
  - release gate
  - metrics review
- 你真正要回答的问题：
  - 当前 CI 检查的是构建正确性，还是也在检查质量 / 运营信号
  - metrics surface 是否已经可被机器消费

<a id="atlas-metrics-logs-and-deployment-scripts"></a>
### `Atlas metrics, logs, and deployment scripts`

- 适用阶段：
  - platform final review
- 你真正要回答的问题：
  - 性能、运行和部署证据是不是已经可以合成统一答辩故事

<a id="final-atlas-performance-traces-eval-outputs-and-ops-notes"></a>
### `final Atlas performance traces, eval outputs, and ops notes`

- 适用阶段：
  - 最终 benchmark
  - 毕业答辩
- 你真正要回答的问题：
  - 最终证据是不是已经能压缩成少数几组最强证明
  - 性能、质量和 operator 视角是否在讲同一个平台故事

## 一个简单用法

当你在 day 文档里看到这类提示时，不要问：

- “这具体是哪份文件？”

先问：

- “它想让我把哪几层东西放在一起看？”
