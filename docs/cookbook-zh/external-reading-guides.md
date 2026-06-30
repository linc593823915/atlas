# 外部阅读提示解读

最后更新：2026-06-30

## 这份文档怎么用

day 文档里有很多阅读提示不是仓库内文件，而是“你该去补什么外部视角”的提示词。

例如：

- `OpenAI production best practices`
- `Kubernetes overview and OpenTelemetry overview`
- `NIST AI RMF GenAI profile`

真正重要的不是“把那篇材料看完”。

而是：

看完之后，你应该从里面提炼什么判断，再带回 Atlas。

这份文档就是把这些提示词翻译成“阅读目标”。

## 服务与基础设施类

<a id="kubernetes-overview-and-opentelemetry-overview"></a>
### `Kubernetes overview and OpenTelemetry overview`

- 适用阶段：
  - 第一学期开篇
  - 第十个月生产部署
- 阅读目标：
  - 理解“运行中的系统”为什么总要处理生命周期、健康、观测和依赖边界
  - 提前建立平台视角，而不是把 Atlas 只当一个本地 CLI
- 回到仓库后重点问：
  - 当前 bootstrap / runtime 有没有为未来部署和观测留出位置
  - 日志和健康检查是不是已经被当成 contract

<a id="go-net-http-documentation"></a>
### `Go net/http documentation`

- 适用阶段：
  - Week 02 CLI / HTTP 服务外壳
- 阅读目标：
  - 关注 server 生命周期、handler 边界和 shutdown 行为
- 回到仓库后重点问：
  - `serve.go` 是否只是入口 glue
  - HTTP 生命周期未来应该接到哪一层

<a id="cobra-command-structure-guidance"></a>
### `Cobra command structure guidance`

- 适用阶段：
  - CLI 入口
  - command tree 扩展
- 阅读目标：
  - 理解 command tree 只是组织入口，不应该承担依赖初始化
- 回到仓库后重点问：
  - `root.go` 和 `serve.go` 有没有开始承担过多业务责任

<a id="ci-workflow-examples-for-go-services"></a>
### `CI workflow examples for Go services`

- 适用阶段：
  - 第一月交付卫生
- 阅读目标：
  - 理解 build / test / lint / smoke 的最小门禁应该长什么样
- 回到仓库后重点问：
  - 当前 CI 期望验证的是可运行性、代码健康，还是只是“命令不报错”

<a id="databasesql-and-redis-client-documentation"></a>
### `database/sql and Redis client documentation`

- 适用阶段：
  - 配置、日志、数据库与缓存接线
- 阅读目标：
  - 理解连接管理、超时、资源释放和 client 生命周期
- 回到仓库后重点问：
  - db/cache 应该接在哪层
  - 哪些资源边界需要 runtime owner

<a id="queue-semantics-for-your-chosen-broker"></a>
### `Queue semantics for your chosen broker`

- 适用阶段：
  - Job Runner
  - queue abstraction
- 阅读目标：
  - 理解 enqueue、claim、ack、release、retry 这些动作在真实 broker 里分别意味着什么
- 回到仓库后重点问：
  - 你的 queue 抽象是不是漏掉了关键状态迁移或 lease 语义

<a id="timeout-and-cancellation-semantics"></a>
### `Timeout and cancellation semantics`

- 适用阶段：
  - timeout / retry
  - tool calling
  - worker execution
- 阅读目标：
  - 理解 deadline、cancel 和 error propagation 的基本 contract
- 回到仓库后重点问：
  - 当前 timeout 行为是在中断工作，还是只在制造另一类隐藏错误

<a id="retry-and-dead-letter-operational-guidance"></a>
### `Retry and dead-letter operational guidance`

- 适用阶段：
  - retry / dead-letter
- 阅读目标：
  - 理解 operator 视角下，什么错误应重试，什么错误应隔离
- 回到仓库后重点问：
  - 当前 retry 策略是不是在放大噪音，dead-letter 是否真的能被消费

<a id="failure-drill-examples-for-background-jobs"></a>
### `Failure drill examples for background jobs`

- 适用阶段：
  - worker review
  - 月度可靠性复盘
- 阅读目标：
  - 学会把超时、重复消费、worker crash 这些问题写成可演练场景
- 回到仓库后重点问：
  - 当前 drill 是否真的覆盖了 operator 最怕的故障

<a id="testing-asynchronous-systems"></a>
### `Testing asynchronous systems`

- 适用阶段：
  - worker loop
  - retry / timeout
  - durable workflow
- 阅读目标：
  - 理解异步系统测试为什么要覆盖顺序、重试、取消、超时和最终状态
  - 学会把“偶尔会发生”的 race 或恢复问题转成可重复场景
- 回到仓库后重点问：
  - 当前测试是否只证明函数返回，还是也证明了状态最终一致
  - 是否有测试能覆盖 worker 退出、重复消费和 timeout 后的状态

<a id="contract-testing-examples"></a>
### `Contract testing examples`

- 适用阶段：
  - Tool Gateway
  - MCP capability
  - schema registry
- 阅读目标：
  - 理解 contract test 保护的是调用方可见行为，而不是某个实现细节
  - 学会把输入、输出、错误语义和兼容性写成测试样本
- 回到仓库后重点问：
  - 现有测试是否能发现 schema drift
  - unsupported、deprecated 和错误响应是否都有调用方视角的验证

## Agent 与模型能力类

<a id="openai-responses-api-guide"></a>
### `OpenAI Responses API guide`

- 适用阶段：
  - 第三个月单 Agent
- 阅读目标：
  - 理解请求、tool 调用、输出收口和结果结构化之间的主路径
- 回到仓库后重点问：
  - 当前 Agent 主路径里，哪一层应该拥有模型请求与结果解析责任

<a id="openai-tools-guide"></a>
### `OpenAI tools guide`

- 适用阶段：
  - tool inventory
  - 多 Tool routing
- 阅读目标：
  - 理解工具定义、调用契约和失败面
- 回到仓库后重点问：
  - tool schema 是否已经足够稳定，失败路径是否明确

<a id="openai-orchestration-and-handoffs-guide"></a>
### `OpenAI orchestration and handoffs guide`

- 适用阶段：
  - multi-agent runtime
  - handoff / specialist 设计
- 阅读目标：
  - 理解 orchestration、handoff 和 role ownership 的常见组织方式
- 回到仓库后重点问：
  - 你的 handoff 是真实 contract，还是 prompt 层顺手拼接

<a id="openai-agent-evals-guide"></a>
### `OpenAI agent evals guide`

- 适用阶段：
  - eval baseline
  - agent regression
- 阅读目标：
  - 理解 agent quality 应该怎样通过数据和 grader 被追踪
- 回到仓库后重点问：
  - 当前评估是不是仍然只在测模型，不在测系统行为

<a id="agent-regression-design-examples"></a>
### `Agent regression design examples`

- 适用阶段：
  - 单 Agent review
  - eval baseline
  - regression gate
- 阅读目标：
  - 理解 agent regression 不只是“回答变差”，也包括 tool 调错、格式漂移、fallback 失效和治理绕过
  - 学会把 regression signal 映射到数据集、grader、trace 和 release gate
- 回到仓库后重点问：
  - 当前 regression 判断是否覆盖输出契约、tool path 和失败路径
  - 哪些退化今天还只能靠人工 spot check

<a id="openai-agents-sdk-guide"></a>
### `OpenAI Agents SDK guide`

- 适用阶段：
  - Agent definitions
  - multi-agent runtime
- 阅读目标：
  - 看清 agent run loop、handoff、shared context 等抽象常见是怎么组织的
- 回到仓库后重点问：
  - Atlas 需要借鉴哪些抽象，哪些又必须保留为本地 runtime 责任

<a id="openai-responses-api-guide"></a>
### `OpenAI Responses API guide`

- 适用阶段：
  - 第三个月单 Agent
  - tool 调用主路径
- 阅读目标：
  - 看清请求输入、模型输出、tool invocation 和解析结果之间的基本链路
- 回到仓库后重点问：
  - Atlas 的请求主路径里，哪些责任应该由 SDK 承担，哪些责任必须收回本地 runtime

<a id="agent-definitions-guide"></a>
### `Agent definitions guide`

- 适用阶段：
  - Agent role / specialist 定义
- 阅读目标：
  - 理解一个 agent definition 至少要描述什么信息
- 回到仓库后重点问：
  - 当前 agent role map 是否已经具备最小输入 / 输出 / owner 说明

<a id="running-agents-guide"></a>
### `Running agents guide`

- 适用阶段：
  - Agent definition
  - multi-agent runtime
- 阅读目标：
  - 理解 agent run loop、角色边界和最小执行语义
- 回到仓库后重点问：
  - 现有单 Agent 流程升级成多 Agent 时，哪一段 run loop 最先变化

<a id="context-and-tracing-guidance"></a>
### `Context and tracing guidance`

- 适用阶段：
  - shared context
  - trace propagation
- 阅读目标：
  - 理解上下文传播与 trace 传播为什么必须同时考虑
- 回到仓库后重点问：
  - 哪些上下文字段必须跨 agent 保留，哪些字段会制造噪音

<a id="openai-structured-outputs-guidance"></a>
### `OpenAI structured outputs guidance`

- 适用阶段：
  - 结构化输出
  - validator
- 阅读目标：
  - 理解 schema 约束、解析失败、结果稳定性
- 回到仓库后重点问：
  - 当前输出是不是仍在依赖“看起来差不多”
  - schema fail 后有没有明确处理路径

<a id="responses-api-request-lifecycle"></a>
### `Responses API request lifecycle`

- 适用阶段：
  - 第一个单 Agent
  - 结构化输出与第一个 Tool
- 阅读目标：
  - 理解一次请求从输入、模型调用、tool invocation 到结果解析的完整生命周期
- 回到仓库后重点问：
  - 当前请求生命周期里，哪些责任属于 SDK，哪些责任属于 Atlas runtime

<a id="openai-production-best-practices"></a>
### `OpenAI production best practices`

- 适用阶段：
  - tool routing
  - 多 Tool 失败处理
- 阅读目标：
  - 看清 tool 调用、失败回退、可观测性和成本意识
- 回到仓库后重点问：
  - routing 是否可解释
  - tool failure 是否被系统化处理

<a id="openai-evaluation-best-practices"></a>
### `OpenAI evaluation best practices`

- 适用阶段：
  - eval baseline
  - eval loop
- 阅读目标：
  - 理解 dataset、grader、可重复评估和结果解释
- 回到仓库后重点问：
  - 当前质量判断是不是仍然依赖 spot check
  - score 能不能回到 trace 和样本

<a id="current-openai-deprecation-notes-for-evals"></a>
### `Current OpenAI deprecation notes for Evals`

- 适用阶段：
  - eval review
  - 最终评审
- 阅读目标：
  - 训练自己关注“评估工具链会变，contract 怎么保持稳定”
- 回到仓库后重点问：
  - 评估体系的 owner 是不是清楚
  - 依赖变化是否会打断现有证据链

## 协议与运行时类

<a id="model-context-protocol-intro"></a>
### `Model Context Protocol intro`

- 适用阶段：
  - MCP RFC 开篇
- 阅读目标：
  - 建立 tools / resources / transports / clients 的最小协议心智
- 回到仓库后重点问：
  - 当前 Atlas 暴露给外部的能力，哪些适合 MCP，哪些还不适合

<a id="mcp-architecture-and-lifecycle-notes"></a>
### `MCP architecture and lifecycle notes`

- 适用阶段：
  - MCP RFC
  - transport / bootstrap
- 阅读目标：
  - 理解 session lifecycle、capability scope、错误语义
- 回到仓库后重点问：
  - 哪些能力真的值得暴露成 MCP surface
  - transport 选择如何影响生命周期

<a id="mcp-specification-sections-for-toolsresources"></a>
### `MCP specification sections for tools/resources`

- 适用阶段：
  - capability review
  - MCP compatibility
- 阅读目标：
  - 精读 tools / resources 在协议中的职责分工与消费者心智
- 回到仓库后重点问：
  - 当前 capability 设计是否真的尊重了协议分工

<a id="compatibility-notes"></a>
### `Compatibility notes`

- 适用阶段：
  - MCP capability review
  - versioning / deprecation
- 阅读目标：
  - 理解兼容性不是“尽量别坏”，而是要写清楚旧 client 如何继续工作
- 回到仓库后重点问：
  - 当前协议变更是不是仍然默认靠调用方自己适配

<a id="mcp-server-implementation-guide"></a>
### `MCP server implementation guide`

- 适用阶段：
  - MCP bootstrap
  - capability implementation
- 阅读目标：
  - 看清 server 生命周期、dispatch、serialization 和错误返回在实现层如何组织
- 回到仓库后重点问：
  - 当前实现是不是把太多协议责任散在 handler 外面

<a id="mcp-error-behavior-sections"></a>
### `MCP error behavior sections`

- 适用阶段：
  - capability compatibility
  - unsupported / failure path
- 阅读目标：
  - 理解协议层如何表达 bad request、unsupported 和 transient failure
- 回到仓库后重点问：
  - 当前 error shape 是否能让 client 做出正确退化动作

<a id="transport-model-guidance"></a>
### `Transport model guidance`

- 适用阶段：
  - MCP bootstrap
  - transport review
- 阅读目标：
  - 理解 stdio、socket、HTTP wrapper 这些 transport 的生命周期差异
- 回到仓库后重点问：
  - transport 选择是否影响了 Atlas runtime 的 owner 与 shutdown 语义

<a id="tool-schema-design-guidance"></a>
### `Tool schema design guidance`

- 适用阶段：
  - tool registry
  - MCP capability
- 阅读目标：
  - 学会区分“给模型看的 schema”和“给系统 contract 用的 schema”
- 回到仓库后重点问：
  - 当前 tool schema 是在表达能力，还是只是在堆字段

<a id="compatibility-notes"></a>
### `Compatibility notes`

- 适用阶段：
  - capability / versioning review
- 阅读目标：
  - 理解向前兼容、向后兼容、弃用和迁移说明怎么写
- 回到仓库后重点问：
  - 当前协议演进是不是默认靠调用方自己猜

<a id="langgraph-overview"></a>
### `LangGraph overview`

- 适用阶段：
  - durable workflow
  - graph state
- 阅读目标：
  - 建立有状态流程、节点图、恢复语义的直觉
- 回到仓库后重点问：
  - 当前 workflow state 是不是已经能被显式表达
  - checkpoint / resume 边界是否清楚

<a id="langgraph-persistence-guide"></a>
### `LangGraph persistence guide`

- 适用阶段：
  - checkpoint store
  - resume model
  - replay / recovery review
- 阅读目标：
  - 理解持久化流程状态时，checkpoint、thread、state update 和恢复入口各自承担什么
  - 借鉴“状态可恢复”背后的 contract，而不是照搬框架 API
- 回到仓库后重点问：
  - Atlas 的 checkpoint 是否记录了恢复所需的最小状态
  - resume 后是否会重复不可重放副作用

## 治理与安全类

<a id="nist-ai-rmf-genai-profile"></a>
### `NIST AI RMF GenAI profile`

- 适用阶段：
  - threat model
  - governance RFC
- 阅读目标：
  - 理解风险识别、控制、监控和责任分层
- 回到仓库后重点问：
  - Atlas 的高风险动作是不是已经被显式识别
  - 控制点是在执行前、执行中还是执行后

<a id="owasp-llm-top-10-2025"></a>
### `OWASP LLM Top 10 2025`

- 适用阶段：
  - prompt injection
  - data leakage
  - tool misuse
- 阅读目标：
  - 熟悉 LLM 系统最常见攻击面和失效方式
- 回到仓库后重点问：
  - 当前 threat model 是否只盯 prompt，而忽略 tool、trace、memory、audit

<a id="prompt-injection-case-studies"></a>
### `Prompt injection case studies`

- 适用阶段：
  - adversarial tests
- 阅读目标：
  - 训练自己把攻击叙事写成可复现测试场景
- 回到仓库后重点问：
  - 你的测试是在测真实攻击方式，还是只在测玩具字符串

<a id="architecture-review-templates"></a>
### `Architecture review templates`

- 适用阶段：
  - platform RFC
  - semester retrospective
- 阅读目标：
  - 学会用固定结构解释系统边界、取舍和残余风险
- 回到仓库后重点问：
  - 你的架构说明是否仍在堆模块名，而不是在讲系统主线

<a id="internal-platform-design-examples"></a>
### `Internal platform design examples`

- 适用阶段：
  - platform boundary
  - graduation defense
- 阅读目标：
  - 训练自己把功能集合重新讲成平台级 operating model
- 回到仓库后重点问：
  - Atlas 现在更像项目，还是更像平台

## 部署、性能与平台类

<a id="docker-build-basics"></a>
### `Docker build basics`

- 适用阶段：
  - Docker / CI / 第一个月交付卫生
- 阅读目标：
  - 理解镜像分层、构建上下文和可重复构建的最小要求
- 回到仓库后重点问：
  - 当前 Docker 路径是不是在复现本地运行路径，而不是单独维护另一套流程

<a id="health-probe-and-resource-limit-guidance"></a>
### `Health probe and resource limit guidance`

- 适用阶段：
  - manifests / probes / resource controls
- 阅读目标：
  - 理解 readiness / liveness / startup probe 差异
  - 理解 request / limit 其实在写运行假设
- 回到仓库后重点问：
  - probe 是否真的映射到系统状态
  - limit / request 是否有实际依据

<a id="slo-and-rollback-strategy-notes"></a>
### `SLO and rollback strategy notes`

- 适用阶段：
  - alerts / rollback / SLO
- 阅读目标：
  - 理解什么叫值得 page 的告警，以及何时必须 rollback
- 回到仓库后重点问：
  - alert rule 是否对应明确 responder action
  - rollback 条件是否客观且可解释

<a id="capacity-planning-notes"></a>
### `Capacity planning notes`

- 适用阶段：
  - optimization loop
  - capacity plan
- 阅读目标：
  - 理解 workload、并发、饱和点和成本约束怎样一起组成容量判断
- 回到仓库后重点问：
  - 当前容量结论是否建立在稳定 workload 上

<a id="caching-and-optimization-case-studies"></a>
### `Caching and optimization case studies`

- 适用阶段：
  - 优化循环
- 阅读目标：
  - 理解优化不是“到处改一点”，而是围绕主瓶颈做单变量变化
- 回到仓库后重点问：
  - 你是在优化真实瓶颈，还是在优化噪音

<a id="tracing-concepts"></a>
### `Tracing concepts`

- 适用阶段：
  - eval / trace
  - runtime observability
- 阅读目标：
  - 建立 span、event、attribute、causality 的最小直觉
- 回到仓库后重点问：
  - 当前 trace 是不是只记录日志替代品，而没记录决策点

<a id="dataset-curation-and-grading-examples"></a>
### `Dataset curation and grading examples`

- 适用阶段：
  - eval review
  - grader calibration
- 阅读目标：
  - 理解数据集样本、标签、grader 和质量解释如何对齐
- 回到仓库后重点问：
  - 当前 dataset row 是否足以支持复查和 rerun

<a id="json-schema-validation-docs-in-your-stack"></a>
### `JSON Schema validation docs in your stack`

- 适用阶段：
  - schema registry
  - config / contract validation
- 阅读目标：
  - 理解 schema 校验边界、错误表达和兼容性影响
- 回到仓库后重点问：
  - 当前 validation 是在保护 contract，还是只在做格式检查

<a id="registry-design-examples"></a>
### `Registry design examples`

- 适用阶段：
  - tool registry
  - capability inventory
- 阅读目标：
  - 学会区分 discoverability、versioning、ownership 三个维度
- 回到仓库后重点问：
  - registry 是不是只有名称，没有真正承担 contract owner 责任

<a id="guardrails-and-approvals-guide"></a>
### `Guardrails and approvals guide`

- 适用阶段：
  - approval / guardrail
- 阅读目标：
  - 理解哪些动作该被阻断、哪些动作该被确认、哪些动作该被记录
- 回到仓库后重点问：
  - approval path 是不是还只存在于文档

<a id="approval-and-audit-workflow-examples"></a>
### `Approval and audit workflow examples`

- 适用阶段：
  - approval / audit integration
- 阅读目标：
  - 理解一个审批动作如何留下完整审计链
- 回到仓库后重点问：
  - 当前 audit event 是否足以重放审批决策

<a id="audit-event-model-examples"></a>
### `Audit event model examples`

- 适用阶段：
  - policy engine
  - audit event
- 阅读目标：
  - 理解一个审计事件最少该保留哪些字段
- 回到仓库后重点问：
  - 当前事件结构是否能解释“为什么被允许”

<a id="human-in-the-loop-workflow-examples"></a>
### `Human-in-the-loop workflow examples`

- 适用阶段：
  - pause / replay / human intervention
- 阅读目标：
  - 理解人工介入点、批准点和状态合并点在 workflow 中各自的角色
- 回到仓库后重点问：
  - 当前人工介入是否有来源、审批和 merge 语义

<a id="replay-and-recovery-guidance"></a>
### `Replay and recovery guidance`

- 适用阶段：
  - replay / resume / recovery drill
- 阅读目标：
  - 理解 replay、resume、recovery 三者的边界
- 回到仓库后重点问：
  - 当前系统是在恢复状态，还是只是在重新跑一遍流程

<a id="authorization-design-guidance"></a>
### `Authorization design guidance`

- 适用阶段：
  - tool auth
  - policy engine
- 阅读目标：
  - 理解授权设计里 actor、action、resource、scope 和 deny path 的核心元素
- 回到仓库后重点问：
  - 当前授权逻辑是不是还停留在“谁都能调，只是希望他别调”

<a id="secure-logging-and-data-minimization-notes"></a>
### `Secure logging and data minimization notes`

- 适用阶段：
  - audit / governance
  - operator docs
- 阅读目标：
  - 理解哪些字段应该保留，哪些字段不该被无脑打进日志
- 回到仓库后重点问：
  - 当前日志和审计信息是在帮助排障，还是在制造泄漏面

<a id="operator-guide-examples"></a>
### `Operator guide examples`

- 适用阶段：
  - operator docs
  - 平台整合
- 阅读目标：
  - 学会把“系统怎么被运行和排障”写成 operator 语言
- 回到仓库后重点问：
  - 当前文档是在写系统介绍，还是在写真实操作指南

<a id="architecture-documentation-structure"></a>
### `Architecture documentation structure`

- 适用阶段：
  - architecture docs
  - graduation defense
- 阅读目标：
  - 学会边界图、控制流、部署视角和证据索引应该怎样组织
- 回到仓库后重点问：
  - 当前文档是模块清单，还是系统故事

<a id="deployment-and-service-manifests"></a>
### `Deployment and service manifests`

- 适用阶段：
  - deployment topology
  - manifests
- 阅读目标：
  - 理解 deployment、service、probe 和 rollout 这些清单对象各自负责什么
- 回到仓库后重点问：
  - 当前部署清单是不是在表达真实运行边界

<a id="kubernetes-production-environment-docs"></a>
### `Kubernetes production environment docs`

- 适用阶段：
  - 第十个月生产部署
- 阅读目标：
  - 理解生产环境里配置、资源、探针和 rollout 的真实约束
- 回到仓库后重点问：
  - Atlas 当前部署设计是不是仍然停留在本地开发者机器视角

<a id="service-deployment-and-rollback-notes"></a>
### `Service deployment and rollback notes`

- 适用阶段：
  - rollback
  - alert / incident response
- 阅读目标：
  - 看清 rollback 不只是命令，而是判断条件 + 操作路径 + 证据
- 回到仓库后重点问：
  - 当前 rollback 规则是否足够具体到 operator 能执行

<a id="metrics-collection-guidance"></a>
### `Metrics collection guidance`

- 适用阶段：
  - metrics / gate / benchmark
- 阅读目标：
  - 理解指标采集点为什么决定你能否解释系统退化
- 回到仓库后重点问：
  - 当前指标是在支撑决策，还是只在支撑展示

<a id="load-test-tooling-docs"></a>
### `Load test tooling docs`

- 适用阶段：
  - baseline / capacity / optimization
- 阅读目标：
  - 了解如何固定 workload、并发和环境，让测量可以复跑
- 回到仓库后重点问：
  - 当前 benchmark 过程是否已经足够可重复

<a id="release-gate-design"></a>
### `Release gate design`

- 适用阶段：
  - regression gate
  - benchmark review
- 阅读目标：
  - 理解 gate 的作用是做决策，而不是做展示
- 回到仓库后重点问：
  - 当前 gate fail 后，到底谁要做什么动作

<a id="release-gate-design-notes"></a>
### `Release gate design notes`

- 适用阶段：
  - gate 细化
- 阅读目标：
  - 看清 threshold、owner、误报处理和回滚动作的最小结构
- 回到仓库后重点问：
  - 你的 gate 规则有没有 owner 和 fail action

<a id="benchmark-design-notes"></a>
### `Benchmark design notes`

- 适用阶段：
  - benchmark RFC
- 阅读目标：
  - 理解 workload、environment、metric、repeatability 为什么必须先冻结
- 回到仓库后重点问：
  - 当前 benchmark 是不是还缺最小 contract

<a id="alerting-design-basics"></a>
### `Alerting design basics`

- 适用阶段：
  - alert rules
  - SLO / rollback
- 阅读目标：
  - 理解 severity、owner、responder action 和 escalation 的最小设计
- 回到仓库后重点问：
  - 当前告警是能指导动作，还是只是告诉你“出事了”

<a id="opentelemetry-metrics-documentation"></a>
### `OpenTelemetry metrics documentation`

- 适用阶段：
  - metrics / traces / observability
- 阅读目标：
  - 理解 metric instruments、labels 和 aggregation 对后续分析的影响
- 回到仓库后重点问：
  - 当前指标是不是已经足够支撑系统级解释

<a id="full-year-notes"></a>
### `Full-year notes`

- 适用阶段：
  - 年度复盘
- 阅读目标：
  - 把一年产物压成少数稳定判断，而不是重讲过程
- 回到仓库后重点问：
  - 哪三组证据最能证明你已经从“做功能”进入“做平台”

<a id="next-year-specialization-options"></a>
### `Next-year specialization options`

- 适用阶段：
  - 下一阶段路线图
- 阅读目标：
  - 学会把未来路线写成主题、项目和证据，而不是写兴趣方向
- 回到仓库后重点问：
  - 下一阶段路线是不是由现有弱点直接推出来的

## 最后一句

外部阅读最容易犯的错是：

- 看完很多
- 摘了很多
- 回到仓库后什么也没改变

真正有效的读法是：

每读完一份外部材料，都要回到 Atlas 问一句：

它改变了我对哪条边界、哪组证据、哪种失败模式的判断？
