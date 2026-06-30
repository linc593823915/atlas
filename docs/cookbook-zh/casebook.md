# 专题案例包

最后更新：2026-06-30

## 这份案例包怎么用

如果你在读 month/week/day 文档时有下面这些感觉：

- 概念好像都懂，但不知道真实工程里怎么出问题
- 知道本周交付标准，但不知道为什么这个标准重要
- 会复述术语，却讲不出一个能说服 reviewer 的系统故事

先来读这份案例包。

每个案例都不是让你背答案。

它的作用是帮你练三件事：

1. 把抽象概念放回真实场景
2. 看懂错误做法会造成什么后果
3. 练习用证据而不是感觉解释工程判断

## 第一学期案例

### 案例 1：服务能启动，但谁都不敢改

- 对应阶段：
  - [第一个月：后端基础与服务骨架](month-01-backend-foundation.md)
  - [Week 01](week-01-learning-system.md)
  - [Week 03](week-03-config-logger-storage.md)
- 场景：
  - Atlas 已经能在一台机器上启动
  - 但 `main`、配置读取、logger 初始化、db/cache 接线全混在一起
  - 每次改入口逻辑都怕把别的路径带崩
- 初学者常见错误：
  - 觉得“先跑起来最重要”，骨架以后再整理
  - 在业务代码里直接读环境变量，直接 new logger 和 client
  - 把 bootstrap 看成目录命名问题，而不是生命周期问题
- 你会看到的坏信号：
  - CLI 入口和 HTTP 入口各自初始化一套依赖
  - README 说不清启动顺序
  - 测试只能测小函数，测不了启动主路径
- 正确拆解方式：
  - 先冻结入口层、bootstrap 层、runtime 层和基础设施层的责任
  - 强制所有配置都从 `config.Manager` 进入
  - 强制所有日志都从统一 logger 抽象进入
  - 把依赖初始化顺序和 shutdown 顺序写进 RFC 与测试
- 最小证据：
  - 一个统一的启动入口
  - 一条可跑通的配置优先级验证
  - 至少一条 bootstrap / runtime 主路径测试
- 回读路径：
  - [术语总索引](glossary.md) 里的 `Bootstrap`、`Runtime`、`config.Manager`
  - [源码与证据阅读路径图](source-reading-map.md) 里的“第一个月”

### 案例 2：任务系统看起来能跑，线上却越跑越乱

- 对应阶段：
  - [第二个月：后台任务与失败处理](month-02-worker-and-failure.md)
  - [Week 05](week-05-job-runner-rfc.md)
  - [Week 07](week-07-retry-timeout-idempotency.md)
- 场景：
  - 你已经有了 queue 和 worker
  - 但某些 job 一超时就被重复执行
  - 有些任务失败后疯狂重试，最后连 operator 都不知道该怎么收口
- 初学者常见错误：
  - 没有先定义状态机和错误语义
  - 觉得 retry 是“更可靠”的同义词
  - 只看任务最终成没成功，不看副作用有没有被放大
- 你会看到的坏信号：
  - 同一个 job 被多次执行，外部系统状态被重复写入
  - queue lag 上升，但没人知道该扩 worker 还是该停重试
  - dead-letter 只是个目录名，没有人真的读它
- 正确拆解方式：
  - 先定义 job state、retryable error 和 terminal error
  - 再定义 lease / visibility timeout、ack / release contract
  - 最后才写 worker loop、graceful shutdown 和 dead-letter flow
- 最小证据：
  - job 状态图
  - retry / timeout / 幂等策略
  - 一组 failure drill 或失败路径测试
- 回读路径：
  - [术语总索引](glossary.md) 里的 `Job State`、`Idempotency Key`、`Dead-letter Queue`
  - [源码与证据阅读路径图](source-reading-map.md) 里的“第二个月”

### 案例 3：做出了会聊天的 Agent，但系统完全不可依赖

- 对应阶段：
  - [第三个月：第一个单 Agent 服务](month-03-first-agent-service.md)
  - [Week 09](week-09-agent-rfc.md)
  - [Week 10](week-10-structured-output.md)
  - [Week 12](week-12-eval-baseline.md)
- 场景：
  - 你做出了一个能调用工具的 Agent
  - 演示时看起来效果很好
  - 但输出没有稳定 schema，tool 失败时整条链会断，没人知道“好”到底怎么衡量
- 初学者常见错误：
  - 把 prompt 文本当成系统 contract
  - 把自然语言输出直接塞给后续代码
  - 没有 eval baseline，只靠演示感觉判断质量
- 你会看到的坏信号：
  - 同样输入今天能跑，明天却换了格式
  - tool routing 是黑箱，operator 解释不了为什么选这个工具
  - 出现坏结果时只能说“模型又抽风了”
- 正确拆解方式：
  - 先冻结 Agent role、tool inventory、output contract
  - 再给结构化输出加 validator 和失败处理路径
  - 最后建立最小 eval baseline 和 operator note
- 最小证据：
  - Agent RFC
  - schema + validator
  - 至少一组可重复的 eval 样本
- 回读路径：
  - [术语总索引](glossary.md) 里的 `Agent Role`、`Structured Output`、`Eval Baseline`
  - [源码与证据阅读路径图](source-reading-map.md) 里的“第三个月”

## 第二学期案例

### 案例 4：工具越来越多，但系统越来越难治理

- 对应阶段：
  - [第四个月：Tool Gateway 与接口质量](month-04-tool-gateway.md)
  - [Week 14](week-14-tool-gateway-rfc.md)
  - [Week 16](week-16-tool-auth-timeout-audit.md)
- 场景：
  - Agent 已经能调多个工具
  - 但每个工具各自处理参数、超时、日志和错误
  - 新接一个工具就像复制一遍旧混乱
- 初学者常见错误：
  - 觉得 gateway 只是多套一层 routing
  - schema 只在调用方校验，不在执行层校验
  - auth、timeout、audit 等以后再补
- 你会看到的坏信号：
  - 同样的 tool error，在不同调用链里表现不一致
  - reviewer 讲不清“高风险工具到底在哪一层被拦住”
  - contract tests 几乎不存在
- 正确拆解方式：
  - 把 validation、timeout、auth、audit 收回到统一 gateway
  - 建 registry，明确 schema owner 和 versioning
  - 用 negative tests 验证错误输入、拒绝路径和超时路径
- 最小证据：
  - Tool Gateway RFC
  - registry + schema contract
  - 一组 contract / negative tests
- 回读路径：
  - [术语总索引](glossary.md) 里的 `Tool Gateway`、`Registry`、`Audit Hook`
  - [源码与证据阅读路径图](source-reading-map.md) 里的“第四个月”

### 案例 5：MCP 服务能启动，但根本不是可消费协议

- 对应阶段：
  - [第五个月：MCP 工程化](month-05-mcp-engineering.md)
  - [Week 18](week-18-mcp-rfc.md)
  - [Week 20](week-20-mcp-capabilities.md)
- 场景：
  - 你已经把 Atlas 某些能力暴露成 MCP server
  - 但 client 不知道 capability 列表是否稳定
  - 一次小改动就会打碎旧 client
- 初学者常见错误：
  - 直接把内部 helper 当公开 surface
  - 没有 capability inventory 和兼容策略
  - 只有本地手调，没有 protocol-level tests
- 你会看到的坏信号：
  - 不同 client 对同一个 capability 理解不一致
  - 旧版本 client 升级 server 后静默出错
  - 文档只说明“怎么启动”，不说明 contract
- 正确拆解方式：
  - 先明确 scope：什么值得暴露，什么只能留在内部
  - 再定义 tools/resources 的分工与 capability discoverability
  - 最后补兼容性说明和协议级测试
- 最小证据：
  - MCP RFC
  - capability inventory
  - protocol tests + compatibility note
- 回读路径：
  - [术语总索引](glossary.md) 里的 `MCP Surface`、`Capability Inventory`、`Compatibility`
  - [源码与证据阅读路径图](source-reading-map.md) 里的“第五个月”

### 案例 6：多 Agent 看起来更强，实际上没人知道谁负责

- 对应阶段：
  - [第六个月：多 Agent 运行时与审批](month-06-multi-agent-runtime.md)
  - [Week 22](week-22-agent-runtime-rfc.md)
  - [Week 24](week-24-handoff-approval-guardrails.md)
- 场景：
  - 系统里有 planner、executor、reviewer 等多个角色
  - 但 handoff 条件不清晰，shared context 来源不清晰，高风险动作也没有真实审批路径
- 初学者常见错误：
  - 把多 Agent 设计成“谁都能接着聊”
  - 共享所有上下文，根本不区分 owner
  - 把 approval 只写进文档，不写进 runtime
- 你会看到的坏信号：
  - 失败后没人能说明哪个 agent 做了错误决策
  - 审计记录只有最终结果，没有中间责任链
  - agent 角色越多，系统越难解释
- 正确拆解方式：
  - 先写 role map，明确 owner、输入和输出 contract
  - 再定义 handoff、shared context 和 trace propagation
  - 最后接 approval / guardrail 与 orchestration tests
- 最小证据：
  - multi-agent runtime RFC
  - handoff / approval 规则
  - orchestration tests 与 operator docs
- 回读路径：
  - [术语总索引](glossary.md) 里的 `Agent Role Map`、`Handoff`、`Approval Path`
  - [源码与证据阅读路径图](source-reading-map.md) 里的“第六个月”

## 第三学期案例

### 案例 7：长流程跑得很长，但根本恢复不了

- 对应阶段：
  - [第七个月：持久化工作流与记忆边界](month-07-durable-workflow.md)
  - [Week 27](week-27-durable-workflow-rfc.md)
  - [Week 28](week-28-graph-state-checkpoint.md)
- 场景：
  - 工作流已经能跨多个步骤
  - 但一旦进程退出、状态损坏或人工介入，就没人知道应该从哪里继续
- 初学者常见错误：
  - 把 checkpoint 当“顺手存个快照”
  - 把 workflow state 和 memory 混为一谈
  - pause / resume 语义不确定，每次恢复都像新开一条流程
- 你会看到的坏信号：
  - 同一个 checkpoint 读出来能推导出多个“下一步”
  - 人工改过一次状态后，系统再也讲不清 provenance
  - recovery drill 只能靠作者手工解释
- 正确拆解方式：
  - 先冻结 state model 和 checkpoint boundary
  - 再定义 resume cursor、replay 语义和人工 merge 语义
  - 最后用 stateful tests 和 operator docs 证明可恢复
- 最小证据：
  - workflow RFC
  - checkpoint design
  - recovery drill / stateful tests
- 回读路径：
  - [术语总索引](glossary.md) 里的 `Workflow State`、`Checkpoint`、`Replay`
  - [源码与证据阅读路径图](source-reading-map.md) 里的“第七个月”

### 案例 8：有了 trace 和评估，但没人知道这些数据该怎么用

- 对应阶段：
  - [第八个月：评估、Tracing 与发布门禁](month-08-evals-traces.md)
  - [Week 31](week-31-eval-loop-rfc.md)
  - [Week 33](week-33-metrics-release-gates.md)
- 场景：
  - 你收集了 trace，做了 grader，建了几条指标
  - 但 team 仍然无法回答“这次改动是不是更好了”
- 初学者常见错误：
  - 用零散样例充当 eval dataset
  - trace 事件太多，但没有决策语义
  - gate 阈值只是写在表里，没人真拿它做发布判断
- 你会看到的坏信号：
  - 每次回归都要人工重新解释指标
  - 分数和 trace 之间对不上
  - release gate fail 了也没人知道该回退还是该忽略
- 正确拆解方式：
  - 先定义 eval loop 的 unit 和 grader strategy
  - 再定义 trace event、dataset version 和 regression threshold
  - 最后把 metric / gate 接进操作流程和复盘文档
- 最小证据：
  - eval loop RFC
  - trace / dataset 模型
  - release gate policy 与 review note
- 回读路径：
  - [术语总索引](glossary.md) 里的 `Eval Loop`、`Trace`、`Regression Gate`
  - [源码与证据阅读路径图](source-reading-map.md) 里的“第八个月”

### 案例 9：治理规则很多，但真正出事时一条都兜不住

- 对应阶段：
  - [第九个月：治理、安全与审计](month-09-governance-audit.md)
  - [Week 35](week-35-governance-rfc.md)
  - [Week 37](week-37-injection-leakage-tests.md)
- 场景：
  - 系统里已经有一些 deny rule、审核说明和审计日志
  - 但面对 prompt injection、越权 tool 调用或数据泄漏时，系统还是靠人兜底
- 初学者常见错误：
  - 先写规则，再去找这些规则到底防什么威胁
  - policy 逻辑分散在不同组件里
  - 对抗测试只是几个玩具字符串
- 你会看到的坏信号：
  - 同一类高风险动作在不同链路里控制强度不同
  - audit 只有“发生了什么”，没有“为什么会被允许”
  - 安全 review 只能靠口头解释
- 正确拆解方式：
  - 从 threat model 和 risk taxonomy 反推 policy boundary
  - 把 action classification 和 audit event 统一到 policy engine
  - 用 injection / leakage / misuse tests 做系统化验证
- 最小证据：
  - governance RFC
  - threat model + policy engine design
  - adversarial tests + review packet
- 回读路径：
  - [术语总索引](glossary.md) 里的 `Threat Model`、`Policy Engine`、`Prompt Injection`
  - [源码与证据阅读路径图](source-reading-map.md) 里的“第九个月”

## 第四学期案例

### 案例 10：Kubernetes 部署了，但一点都不像生产系统

- 对应阶段：
  - [第十个月：生产部署与运行模型](month-10-production-kubernetes.md)
  - [Week 40](week-40-kubernetes-deployment-rfc.md)
  - [Week 42](week-42-alerts-rollback-slo.md)
- 场景：
  - Atlas 已经有 manifest
  - 但 probe 只是返回 200，资源限制是复制别人的默认值，告警和 rollback 也没有真实触发条件
- 初学者常见错误：
  - 把 Kubernetes 当成“把服务塞进 YAML”
  - 没有把 deployment topology 和 runtime responsibility 对齐
  - rollback 只存在于想象里
- 你会看到的坏信号：
  - readiness / liveness 混成一个探针
  - operator 不知道告警后先扩容还是先回滚
  - release 后出问题却没有客观 rollback signal
- 正确拆解方式：
  - 先定义 deployment topology 和 health model
  - 再定义 probes、resource control 和 rollout / rollback 规则
  - 最后把 SLO / alert / operator runbook 串起来
- 最小证据：
  - deployment RFC
  - probe / resource / alert 策略
  - rollback 规则和演练记录
- 回读路径：
  - [术语总索引](glossary.md) 里的 `Deployment Topology`、`Readiness Probe`、`SLO`
  - [源码与证据阅读路径图](source-reading-map.md) 里的“第十个月”

### 案例 11：你拿到一堆性能数字，但一个能信的结论都没有

- 对应阶段：
  - [第十一个月：性能、成本与容量规划](month-11-benchmark-capacity.md)
  - [Week 44](week-44-baseline-latency-cost.md)
  - [Week 45](week-45-optimization-capacity.md)
- 场景：
  - 你测了 latency、throughput、token cost
  - 也做了 before / after 对比
  - 但 workload 不一致、环境不一致、样本量不一致，最后谁也不相信这些数字
- 初学者常见错误：
  - 先优化，再补 benchmark contract
  - 把平均值当全部故事
  - 把容量规划写成拍脑袋估算
- 你会看到的坏信号：
  - 每次 rerun 结果都变
  - before / after 对比看似提升，但没人能解释瓶颈是否真的变化
  - 讨论成本时只说 token，不说 workload mix 和架构约束
- 正确拆解方式：
  - 先冻结 benchmark contract 和 baseline
  - 再做单变量优化循环
  - 最后写 capacity assumption 和 trade-off note
- 最小证据：
  - benchmark contract
  - baseline 数据
  - before / after 结果与瓶颈分析
- 回读路径：
  - [术语总索引](glossary.md) 里的 `Baseline`、`P95 Latency`、`Capacity Assumption`
  - [源码与证据阅读路径图](source-reading-map.md) 里的“第十一个月”

### 案例 12：功能都在，为什么还是说不清这是个平台

- 对应阶段：
  - [第十二个月：平台整合与毕业答辩](month-12-platform-integration.md)
  - [Week 47](week-47-config-tool-mcp-integration.md)
  - [Week 49](week-49-end-to-end-drills.md)
  - [Week 52](week-52-year-end-retrospective.md)
- 场景：
  - Atlas 已经有配置、tool gateway、MCP、workflow、eval、approval、audit、部署和 benchmark
  - 但面对 reviewer，你还是讲不清这些东西怎样组成一套平台
- 初学者常见错误：
  - 把所有成果列成清单，而不是讲成一条系统故事
  - 只谈成功能力，不谈限制和 trade-off
  - 只拿文档说话，不拿 end-to-end drill 和 evidence pack 说话
- 你会看到的坏信号：
  - reviewer 能记住功能名，却记不住你的平台边界
  - 同一个概念在不同子系统里名字和 contract 不一致
  - 答辩时只能说“都已经接上了”，却讲不清 why now / why this design
- 正确拆解方式：
  - 先画清平台边界和共享 contract
  - 再用跨子系统 drill 和 benchmark 证明系统真的协同工作
  - 最后把证据整理成 architecture pack、operator guide 和 graduation defense
- 最小证据：
  - platform RFC
  - end-to-end drill
  - evidence pack + next roadmap
- 回读路径：
  - [术语总索引](glossary.md) 里的 `Platform Boundary`、`Evidence Pack`、`Graduation Defense`
  - [源码与证据阅读路径图](source-reading-map.md) 里的“第十二个月”

## 读完案例后应该做什么

不要把案例包当成“标准答案集”。

正确用法是：

1. 找到你当前正在学的月份
2. 先读对应案例，确认真实风险和错误做法
3. 再回 month/week/day 文档，重新看交付标准为什么重要
4. 最后回 Atlas issue、周报和代码/治理材料，把案例里的判断对到仓库证据上
