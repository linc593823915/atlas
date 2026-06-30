# 术语总索引

最后更新：2026-06-30

## 这份索引怎么用

这不是一本独立教材。

它的作用是：

1. 当你在月/周/日文档里反复看到一个词却还说不清它时，先来这里查最小定义
2. 当你准备写 RFC、接口说明、review note 或复盘时，用它校验自己是不是只会复述名词
3. 当两个相近概念容易混淆时，用这里的“重点阅读”回到对应章节

## 第一学期：服务地基、任务系统与第一个 Agent

- `Bootstrap`：负责依赖初始化、启动和关闭流程的统一入口。重点阅读：[第一个月](month-01-backend-foundation.md)、[Week 01](week-01-learning-system.md)
- `Runtime`：承载服务真实运行状态与依赖对象的最小执行壳。重点阅读：[Week 02](week-02-cli-http-bootstrap.md)
- `config.Manager`：统一默认值、配置文件和环境覆盖规则的配置中心。重点阅读：[Week 03](week-03-config-logger-storage.md)
- `结构化日志`：用稳定字段而不是自由文本记录事件，便于检索、审计和 tracing。重点阅读：[第一个月](month-01-backend-foundation.md)、[Week 03](week-03-config-logger-storage.md)
- `健康检查`：系统对外暴露的最小可运行性信号，不等于“进程还活着”。重点阅读：[Week 02](week-02-cli-http-bootstrap.md)
- `Job State`：后台任务在排队、执行、重试、失败、死信之间的状态语义。重点阅读：[第二个月](month-02-worker-and-failure.md)、[Week 05](week-05-job-runner-rfc.md)
- `Lease / Visibility Timeout`：任务被 worker 暂时占有的时间窗口，用于避免重复消费。重点阅读：[Week 06](week-06-worker-loop.md)
- `Idempotency Key`：保证重复执行同一任务时不会放大副作用的稳定标识。重点阅读：[Week 07](week-07-retry-timeout-idempotency.md)
- `Dead-letter Queue`：承接多次失败或不可恢复任务的隔离区，用于排障和治理。重点阅读：[Week 07](week-07-retry-timeout-idempotency.md)
- `Graceful Shutdown`：停止拉新、排空在途、保存必要状态后再安全退出的流程。重点阅读：[Week 06](week-06-worker-loop.md)
- `Agent Role`：模型在一条链路里承担的职责边界，而不是一段 prompt 文本。重点阅读：[第三个月](month-03-first-agent-service.md)、[Week 09](week-09-agent-rfc.md)
- `Tool Inventory`：一个 Agent 允许调用的工具集合以及各自 contract。重点阅读：[Week 09](week-09-agent-rfc.md)
- `Structured Output`：带 schema 约束的模型输出，用于进入确定性代码路径。重点阅读：[Week 10](week-10-structured-output.md)
- `Validator`：负责检查模型输出是否满足 contract 的校验层。重点阅读：[Week 10](week-10-structured-output.md)
- `Eval Baseline`：后续回归比较所依赖的第一版质量基线。重点阅读：[Week 12](week-12-eval-baseline.md)

## 第二学期：接口、协议与运行时

- `Tool Gateway`：统一接管工具调用的入口层，负责 validation、timeout、audit 等控制。重点阅读：[第四个月](month-04-tool-gateway.md)、[Week 14](week-14-tool-gateway-rfc.md)
- `Registry`：保存 tool schema、version 和 discoverability 信息的注册中心。重点阅读：[Week 15](week-15-schema-registry.md)
- `Schema Validation`：在执行前后检查输入输出是否仍满足约束的校验机制。重点阅读：[Week 15](week-15-schema-registry.md)
- `Timeout Policy`：定义谁负责超时预算、取消传播和失败后行为的规则。重点阅读：[Week 16](week-16-tool-auth-timeout-audit.md)
- `Audit Hook`：在关键调用前后写入可追踪决策证据的钩子。重点阅读：[Week 16](week-16-tool-auth-timeout-audit.md)
- `MCP Surface`：Atlas 向外部 client 暴露的稳定协议能力集合。重点阅读：[第五个月](month-05-mcp-engineering.md)、[Week 18](week-18-mcp-rfc.md)
- `Capability Inventory`：当前 server 提供的 tools/resources 及其元信息清单。重点阅读：[Week 18](week-18-mcp-rfc.md)
- `Tool / Resource`：tool 偏命令式动作，resource 偏只读状态或上下文。重点阅读：[Week 18](week-18-mcp-rfc.md)
- `Transport`：承载 MCP 会话的通信方式，如 stdio、socket 等。重点阅读：[Week 19](week-19-mcp-bootstrap.md)
- `Compatibility`：服务端演进时仍让既有 client 可预测工作的约束。重点阅读：[Week 20](week-20-mcp-capabilities.md)
- `Agent Role Map`：多 Agent 体系中各角色的职责分工图。重点阅读：[第六个月](month-06-multi-agent-runtime.md)、[Week 22](week-22-agent-runtime-rfc.md)
- `Shared Context`：在多个 agent 间传播且仍需保持来源清晰的上下文。重点阅读：[Week 23](week-23-agent-definitions.md)
- `Handoff`：把目标、上下文和责任从一个 agent 转移给另一个 agent。重点阅读：[Week 24](week-24-handoff-approval-guardrails.md)
- `Approval Path`：敏感动作在执行前必须经过的人工或策略审批路径。重点阅读：[Week 24](week-24-handoff-approval-guardrails.md)
- `Guardrail`：限制 agent 行为边界的运行时约束与拦截规则。重点阅读：[Week 24](week-24-handoff-approval-guardrails.md)

## 第三学期：状态、评估与治理

- `Workflow State`：长流程运行时必须被显式保存和恢复的状态集合。重点阅读：[第七个月](month-07-durable-workflow.md)、[Week 27](week-27-durable-workflow-rfc.md)
- `Checkpoint`：在流程关键节点持久化的可恢复快照。重点阅读：[Week 28](week-28-graph-state-checkpoint.md)
- `Resume Cursor`：告诉 runtime 应从哪个状态节点继续执行的定位信息。重点阅读：[Week 28](week-28-graph-state-checkpoint.md)
- `Replay`：基于历史状态和证据重新演示或重放流程的机制。重点阅读：[Week 29](week-29-pause-replay-human-loop.md)
- `Human Intervention`：人工在流程中插入决策、修正或批准的受控接入点。重点阅读：[Week 29](week-29-pause-replay-human-loop.md)
- `Eval Loop`：从数据集、grader 到结果解释的完整评估循环。重点阅读：[第八个月](month-08-evals-traces.md)、[Week 31](week-31-eval-loop-rfc.md)
- `Grader`：负责给一次运行结果打分或分类的评估器。重点阅读：[Week 32](week-32-graders-traces-dataset.md)
- `Trace`：记录模型、工具、策略和状态变化过程的可观测链路。重点阅读：[Week 32](week-32-graders-traces-dataset.md)
- `Regression Gate`：根据指标阈值决定是否阻断变更的质量门禁。重点阅读：[Week 33](week-33-metrics-release-gates.md)
- `Benchmark Dataset`：用于稳定对比系统表现的一组代表性样本。重点阅读：[Week 32](week-32-graders-traces-dataset.md)
- `Threat Model`：描述资产、角色、攻击路径和控制点的风险模型。重点阅读：[第九个月](month-09-governance-audit.md)、[Week 35](week-35-governance-rfc.md)
- `Policy Engine`：根据上下文和规则决定动作是否允许的决策层。重点阅读：[Week 36](week-36-policy-engine.md)
- `Action Classification`：把系统动作映射成不同风险等级与控制策略的分类过程。重点阅读：[Week 36](week-36-policy-engine.md)
- `Audit Event`：能解释是谁、为何、在何时做了什么的正式留痕单元。重点阅读：[Week 36](week-36-policy-engine.md)
- `Prompt Injection`：外部不可信指令试图改变系统目标或绕过约束的攻击方式。重点阅读：[Week 37](week-37-injection-leakage-tests.md)

## 第四学期：生产、性能与平台化

- `Deployment Topology`：服务在生产环境中的部署单元和连接关系。重点阅读：[第十个月](month-10-production-kubernetes.md)、[Week 40](week-40-kubernetes-deployment-rfc.md)
- `Readiness Probe`：判断实例是否可以接收流量的探针。重点阅读：[Week 41](week-41-manifests-healthchecks.md)
- `Liveness Probe`：判断实例是否已进入不可恢复错误状态的探针。重点阅读：[Week 41](week-41-manifests-healthchecks.md)
- `SLO`：系统对可用性、时延等关键体验做出的服务级目标承诺。重点阅读：[Week 42](week-42-alerts-rollback-slo.md)
- `Rollback Criteria`：决定何时必须回退版本的客观信号与规则。重点阅读：[Week 42](week-42-alerts-rollback-slo.md)
- `Baseline`：后续所有优化对比都依赖的第一版可信测量结果。重点阅读：[第十一个月](month-11-benchmark-capacity.md)、[Week 44](week-44-baseline-latency-cost.md)
- `P95 Latency`：95% 请求不应超过的时延位置，用于看尾延迟。重点阅读：[Week 44](week-44-baseline-latency-cost.md)
- `Throughput`：单位时间内系统能够完成的请求或任务量。重点阅读：[Week 44](week-44-baseline-latency-cost.md)
- `Token Cost`：模型调用因 token 消耗而带来的直接成本约束。重点阅读：[Week 44](week-44-baseline-latency-cost.md)
- `Capacity Assumption`：关于并发、负载和饱和点的容量前提假设。重点阅读：[Week 45](week-45-optimization-capacity.md)
- `Platform Boundary`：把零散子系统组织成一套平台时必须冻结的边界。重点阅读：[第十二个月](month-12-platform-integration.md)、[Week 46](week-46-platform-rfc.md)
- `End-to-end Drill`：跨多个子系统执行完整用户旅程和故障演练的验证方式。重点阅读：[Week 49](week-49-end-to-end-drills.md)
- `Evidence Pack`：支持评审或答辩的一组结构化证据材料。重点阅读：[Week 50](week-50-architecture-operator-docs.md)、[Week 51](week-51-final-benchmark-review.md)
- `Operator Guide`：教接手者和运维者如何运行、排障和回退系统的说明。重点阅读：[Week 50](week-50-architecture-operator-docs.md)
- `Graduation Defense`：围绕系统设计、证据和未来路线进行的毕业级答辩。重点阅读：[Week 52](week-52-year-end-retrospective.md)

## 怎么判断自己真的学会了

如果你已经学会一个术语，你应该至少能同时回答下面三个问题：

1. 它在 Atlas 的哪一层出现？
2. 它解决的真实工程问题是什么？
3. 如果这个概念被做错了，系统会先在哪里坏掉？
