# 常见交付物词典

最后更新：2026-06-30

## 这份词典解决什么问题

day 文档里经常会出现一些“看起来像产物名词”的短语：

- `job 状态图`
- `queue 抽象`
- `Agent RFC`
- `capability inventory`
- `baseline 数据`
- `trace 样例`

这些词如果只是看字面，很容易产生两个问题：

1. 不知道它至少应该包含什么
2. 不知道它和别的产物有什么区别

这份词典就是把这些高频交付物词条压成最小定义和最低结构。

## 基础与任务系统类

<a id="job-state"></a>
### `job 状态图`

- 它是什么：
  - 用状态机说明后台任务从进入队列到完成/失败/死信的全部状态迁移
- 最低要有：
  - 状态列表
  - 迁移条件
  - retryable / terminal 区分
  - 哪些迁移需要 operator 介入

<a id="queue-abstraction"></a>
### `queue 抽象`

- 它是什么：
  - 对 enqueue、claim、ack、release、retry 这些动作的统一 contract
- 最低要有：
  - 调用方和 owner
  - 输入输出
  - 失败语义
  - visibility / lease / timeout 假设

<a id="drill-record"></a>
### `drill 记录`

- 它是什么：
  - 一次失败演练的输入、过程、结果和结论
- 最低要有：
  - drill 场景
  - 触发方式
  - 观测信号
  - 恢复动作
  - 结论与残余风险

<a id="drill-note"></a>
### `一份 drill 说明` / `drill 文档`

- 它是什么：
  - 把一次故障演练写成别人可以复跑、复查和改进的说明
- 最低要有：
  - 场景
  - 触发方式
  - 预期信号
  - 恢复步骤
  - 实际结果

<a id="failure-path-tests"></a>
### `一组失败路径测试` / `失败路径测试`

- 它是什么：
  - 专门证明坏输入、超时、拒绝、下游失败或状态异常时系统行为可控的测试集合
- 最低要有：
  - 失败场景
  - 触发输入
  - 预期错误或降级
  - 不应发生的副作用
  - 日志、trace 或状态证据

## Agent 与接口类

<a id="agent-rfc"></a>
### `Agent RFC`

- 它是什么：
  - 冻结 Agent role、tool inventory、output contract 和失败语义的设计文档
- 最低要有：
  - 角色说明
  - 输入输出
  - tool 使用边界
  - 失败路径
  - 最小验收口径

<a id="job-runner-rfc"></a>
### `Job Runner RFC`

- 它是什么：
  - 后台任务状态机、queue contract、worker failure model 的设计文档
- 最低要有：
  - job state
  - queue 抽象
  - retry / timeout / dead-letter 语义
  - 最小验收口径

<a id="tool-gateway-rfc-draft"></a>
### `Tool Gateway RFC 草稿`

- 它是什么：
  - 定义统一 tool invocation 入口的一版设计草稿
- 最低要有：
  - request / result envelope
  - validation / timeout / audit 边界
  - 失败路径假设

<a id="tool-gateway-boundary"></a>
### `Tool Gateway 的职责边界`

- 它是什么：
  - 说明 Tool Gateway 统一接管哪些控制面，以及哪些责任仍留给 caller、tool wrapper 或 runtime
- 最低要有：
  - request / result envelope
  - schema validation
  - timeout policy
  - auth / approval / audit hook
  - error mapping

<a id="mcp-rfc"></a>
### `MCP RFC`

- 它是什么：
  - 定义 MCP surface、capability scope、transport 和错误语义的一版协议文档
- 最低要有：
  - surface 范围
  - capability inventory
  - transport 假设
  - compatibility / error shape

<a id="durable-workflow-rfc"></a>
### `Durable Workflow RFC`

- 它是什么：
  - 长流程状态、checkpoint、resume / replay 语义的一版设计文档
- 最低要有：
  - state model
  - checkpoint boundary
  - resume / replay 语义
  - human intervention 入口

<a id="governance-rfc"></a>
### `治理 RFC`

- 它是什么：
  - 定义 policy、approval、audit、risk taxonomy 和高风险动作控制边界的设计文档
- 最低要有：
  - 风险分类
  - 高风险动作
  - 控制点
  - 审批和审计路径
  - 验收证据

<a id="deployment-rfc"></a>
### `部署 RFC`

- 它是什么：
  - 定义部署拓扑、资源限制、健康检查、发布和回滚策略的设计文档
- 最低要有：
  - 部署单元
  - runtime / process model
  - config / secret
  - health model
  - rollback 条件

<a id="eval-loop-rfc"></a>
### `Eval Loop RFC`

- 它是什么：
  - 评估单元、grader、dataset 和 gate 语义的一版设计文档
- 最低要有：
  - evaluation unit
  - grader strategy
  - dataset structure
  - gate 规则

<a id="multi-agent-rfc"></a>
### `多 Agent RFC`

- 它是什么：
  - 多 Agent role map、handoff、shared context 和 approval 语义的一版设计文档
- 最低要有：
  - role map
  - handoff boundary
  - shared context
  - approval / guardrail 说明

<a id="benchmark-rfc"></a>
### `benchmark RFC`

- 它是什么：
  - 定义 workload、environment、metric 和 repeatability 的 benchmark 设计文档
- 最低要有：
  - workload
  - 环境
  - 指标口径
  - repeatability 假设

<a id="tool-list"></a>
### `Tool 列表`

- 它是什么：
  - 当前主题里允许暴露给模型或 runtime 的工具集合
- 最低要有：
  - tool 名称
  - 用途
  - 输入/输出 schema
  - 风险等级

<a id="capability-inventory"></a>
### `capability inventory`

- 它是什么：
  - MCP 或工具系统当前公开能力的注册清单
- 最低要有：
  - capability 名称
  - 类型（tool / resource / 其他）
  - 用途
  - 版本或兼容性说明

<a id="transport-path"></a>
### `transport 路径`

- 它是什么：
  - MCP 或服务协议里，从连接建立到请求分发再到返回的路径说明
- 最低要有：
  - transport 类型
  - 生命周期
  - dispatch 入口
  - shutdown 语义

<a id="mcp-doc-note"></a>
### `MCP 文档`

- 它是什么：
  - 面向使用者说明 capability、transport 和错误语义的文档
- 最低要有：
  - capability 列表
  - transport 说明
  - compatibility / error shape 简述

<a id="mcp-bootstrap-note"></a>
### `MCP bootstrap`

- 它是什么：
  - MCP server 从进程启动、transport 初始化到 capability 注册完成的最小启动链路说明
- 最低要有：
  - 启动入口
  - transport 初始化
  - capability registry
  - shutdown 或错误返回路径

<a id="mcp-integration-note"></a>
### `MCP integration`

- 它是什么：
  - MCP surface 与内部 Tool Gateway、config、runtime 之间的整合说明
- 最低要有：
  - 共享 contract
  - capability 映射
  - timeout / error / audit 对齐方式
  - 兼容性检查

<a id="mcp-capability-boundary"></a>
### `MCP 能力边界`

- 它是什么：
  - 说明哪些能力适合暴露给 MCP client，哪些不应该暴露
- 最低要有：
  - capability 范围
  - 不暴露项
  - 风险理由
  - 后续扩展条件

<a id="tools-resources-capability"></a>
### `tools/resources capability`

- 它是什么：
  - MCP 中 tool 与 resource 两类 capability 的划分说明
- 最低要有：
  - tool 负责什么
  - resource 负责什么
  - 输入输出差异
  - client 消费方式

<a id="protocol-tests"></a>
### `Protocol tests`

- 它是什么：
  - 验证协议请求、响应、错误和兼容性行为的测试集合
- 最低要有：
  - 合法请求
  - 非法请求
  - unsupported capability
  - error shape
  - version / compatibility 样本

<a id="contract-tests"></a>
### `Contract tests`

- 它是什么：
  - 从调用方视角保护接口或协议稳定性的测试集合
- 最低要有：
  - 输入样本
  - 输出断言
  - 错误语义
  - 兼容性或版本假设

<a id="negative-tests"></a>
### `Negative tests`

- 它是什么：
  - 专门证明错误输入、拒绝路径和失败语义可控的测试集合
- 最低要有：
  - 坏输入
  - 越权或 unsupported 行为
  - 预期错误
  - 不应发生的副作用

<a id="protocol-boundary-note"></a>
### `协议边界说明`

- 它是什么：
  - 说明外部协议面与内部实现边界的一份短文档
- 最低要有：
  - 暴露什么
  - 不暴露什么
  - 哪些行为对外可见

<a id="compatibility-note"></a>
### `兼容性说明`

- 它是什么：
  - 面向调用方说明版本演进和不兼容变更策略的文档
- 最低要有：
  - 兼容范围
  - 弃用策略
  - 升级建议

<a id="output-contract"></a>
### `Output contract`

- 它是什么：
  - 定义模型、Agent、tool 或 API 输出结构和失败输出的契约
- 最低要有：
  - 字段
  - 必填 / 可选
  - 成功与失败形态
  - 解析方
  - 版本策略

<a id="tool-inventory"></a>
### `Tool inventory`

- 它是什么：
  - 当前可用工具的能力、schema、owner 和风险等级清单
- 最低要有：
  - tool 名称
  - 用途
  - schema
  - owner
  - 风险等级

<a id="registry-model"></a>
### `Registry 模型`

- 它是什么：
  - 注册、发现、版本和状态管理的模型
- 最低要有：
  - key
  - metadata
  - version
  - supported / deprecated / experimental 状态
  - owner

<a id="schema-validation"></a>
### `Schema 校验`

- 它是什么：
  - 用 schema 保护输入输出 contract 的校验机制
- 最低要有：
  - schema 来源
  - 校验时机
  - 错误表达
  - 版本兼容策略

<a id="shared-context-design"></a>
### `shared context 设计`

- 它是什么：
  - 多 Agent 之间传播上下文的最小 contract
- 最低要有：
  - 哪些字段共享
  - 哪些字段不共享
  - 来源说明
  - trace / approval / audit 关联字段

<a id="shared-context"></a>
### `shared context`

- 它是什么：
  - 多 Agent、workflow 或审批链路中需要跨 owner 传递的最小上下文
- 最低要有：
  - 必传字段
  - 禁止扩散字段
  - 来源
  - 生命周期
  - trace / approval 关联

<a id="agent-definition"></a>
### `Agent definition`

- 它是什么：
  - 单个 Agent 的角色、输入输出、工具范围和运行约束定义
- 最低要有：
  - role
  - instructions / constraints
  - input / output contract
  - allowed tools
  - failure behavior

<a id="agent-role-map"></a>
### `Agent role map`

- 它是什么：
  - 多 Agent 系统中各角色职责和交接关系的地图
- 最低要有：
  - role 列表
  - owner
  - handoff 条件
  - shared context 字段

<a id="runtime-responsibility"></a>
### `Runtime responsibility`

- 它是什么：
  - 说明 runtime 拥有哪些生命周期、状态、调度和失败处理责任
- 最低要有：
  - 启动 / 停止
  - context / cancellation
  - state owner
  - error propagation

<a id="tool-auth-hook"></a>
### `Tool auth hook`

- 它是什么：
  - tool invocation 前进行授权、策略或审批检查的扩展点
- 最低要有：
  - 输入上下文
  - actor / action / resource
  - allow / deny / approval
  - audit 字段

<a id="audit-metadata"></a>
### `audit metadata`

- 它是什么：
  - 跟随请求、tool call 或审批动作传播的审计元数据
- 最低要有：
  - actor
  - action
  - trace id
  - approval id
  - reason / decision

<a id="guardrail-draft"></a>
### `guardrail 规则初稿`

- 它是什么：
  - 一版明确哪些动作禁止、哪些动作需确认、哪些动作必须留痕的规则集
- 最低要有：
  - 动作类型
  - 允许 / 拒绝 / 需审批
  - 触发条件
  - 留痕要求

## 工作流与治理类

<a id="checkpoint-design-draft"></a>
### `checkpoint 设计草稿`

- 它是什么：
  - 状态持久化和恢复点的最小设计说明
- 最低要有：
  - 保存什么
  - 什么时候保存
  - 恢复从哪里继续
  - 哪些副作用不能直接重放

<a id="state-diagram"></a>
### `状态图`

- 它是什么：
  - 描述 workflow、job、agent run 或部署状态如何迁移的图或表
- 最低要有：
  - 状态列表
  - 迁移条件
  - terminal 状态
  - 非法迁移
  - owner

<a id="graph-state"></a>
### `graph state`

- 它是什么：
  - 图式 workflow 中节点共享或传递的状态模型
- 最低要有：
  - state 字段
  - 节点读写权限
  - 更新规则
  - checkpoint 关系

<a id="workflow-state"></a>
### `workflow state`

- 它是什么：
  - workflow 执行过程中用于恢复、审计和判断下一步的状态集合
- 最低要有：
  - 当前节点
  - 输入摘要
  - 决策结果
  - checkpoint id
  - 副作用记录

<a id="resume-model"></a>
### `resume model`

- 它是什么：
  - workflow 从 checkpoint 恢复时如何选择下一步的模型
- 最低要有：
  - resume cursor
  - 恢复输入
  - 状态校验
  - 不重复副作用机制

<a id="checkpoint-storage-model"></a>
### `checkpoint 存储模型`

- 它是什么：
  - checkpoint 在物理或逻辑存储层的组织方式
- 最低要有：
  - key / version
  - 原子性假设
  - 读取 / 覆盖 / 回滚策略

<a id="replay-note"></a>
### `replay 说明`

- 它是什么：
  - 描述一次 replay 会重放什么、不会重放什么
- 最低要有：
  - replay 目标
  - 证据与副作用边界
  - 与 resume 的区别

<a id="replay"></a>
### `replay`

- 它是什么：
  - 基于已有输入、状态和证据重新演示或重放流程，用来验证恢复或解释决策
- 最低要有：
  - replay 起点
  - 使用的 checkpoint 或 trace
  - 会重放的内容
  - 不会重做的副作用
  - replay 后结论

<a id="policy-structure"></a>
### `policy 结构`

- 它是什么：
  - policy engine 里的规则组织方式
- 最低要有：
  - 输入上下文
  - 动作分类
  - 决策输出
  - deny / allow / escalate 路径

<a id="policy-engine"></a>
### `policy engine`

- 它是什么：
  - 根据 actor、action、resource、context 输出 allow / deny / approval 的控制组件
- 最低要有：
  - 输入上下文
  - 规则来源
  - 决策输出
  - audit 字段
  - 测试样本

<a id="risk-taxonomy"></a>
### `risk taxonomy`

- 它是什么：
  - 用于分类高风险动作、误用、数据泄漏和治理失败的风险分类表
- 最低要有：
  - 风险类别
  - 示例路径
  - 控制点
  - 检测信号

<a id="governance-boundary"></a>
### `governance boundary`

- 它是什么：
  - 说明治理逻辑与业务逻辑、tool 执行和协议 surface 的边界
- 最低要有：
  - governance owner
  - 接入点
  - 不接管的行为
  - audit / approval 关系

<a id="guardrail-constraint"></a>
### `guardrail 约束`

- 它是什么：
  - 对高风险动作的预防性或执行时约束
- 最低要有：
  - 约束对象
  - 触发条件
  - 阻断或降级行为
  - 证据记录

<a id="action-classification"></a>
### `action classification`

- 它是什么：
  - 把动作按风险、审批需求和审计需求分类的产物
- 最低要有：
  - action
  - risk level
  - allow / deny / approval
  - audit requirement

<a id="control-explanation"></a>
### `control explanation`

- 它是什么：
  - 对某个控制点为什么存在、如何触发、如何证明有效的说明
- 最低要有：
  - 控制目标
  - 触发路径
  - 系统行为
  - 证据

<a id="audit-event-draft"></a>
### `audit event 草稿`

- 它是什么：
  - 一版审计事件字段定义
- 最低要有：
  - actor
  - action
  - resource / scope
  - decision
  - why allowed / denied
  - trace / approval 关联信息

<a id="approval-path"></a>
### `approval path`

- 它是什么：
  - 高风险动作从请求审批到允许或拒绝的完整路径
- 最低要有：
  - 触发条件
  - 请求字段
  - approver
  - allow / deny 结果
  - workflow 与 audit 影响

<a id="handoff-boundary"></a>
### `handoff 边界`

- 它是什么：
  - 一个 Agent、workflow 节点或子系统把控制权交给另一个 owner 的边界
- 最低要有：
  - 触发条件
  - 交接上下文
  - 不交接的内容
  - 结果 owner

<a id="checkpoint-boundary"></a>
### `checkpoint 边界`

- 它是什么：
  - 定义在哪些决策点保存状态，以及保存到什么粒度
- 最低要有：
  - 保存时机
  - 保存字段
  - 恢复 cursor
  - 不能重放的副作用

<a id="checkpoint-store"></a>
### `checkpoint store`

- 它是什么：
  - 保存 workflow checkpoint 的存储抽象或实现
- 最低要有：
  - key 模型
  - version
  - 原子写入假设
  - 读取与回滚语义

<a id="resume-tests"></a>
### `resume tests`

- 它是什么：
  - 验证恢复后控制流、状态和副作用边界的测试集合
- 最低要有：
  - 正常 resume
  - checkpoint 缺失或损坏
  - human intervention 后 resume
  - 不重复副作用断言

<a id="stateful-testing"></a>
### `stateful testing`

- 它是什么：
  - 围绕状态迁移、不变量和恢复语义设计的测试方式
- 最低要有：
  - 初始状态
  - 操作序列
  - 预期状态
  - 非法迁移断言

## 评估与性能类

<a id="dataset-structure-draft"></a>
### `dataset 结构草稿`

- 它是什么：
  - eval 或 benchmark 样本集的一版结构定义
- 最低要有：
  - 输入样本
  - 预期行为
  - 标签 / grader 所需字段
  - version 字段

<a id="grader-structure"></a>
### `grader 结构`

- 它是什么：
  - grader 接受什么输入、输出什么分数或判断的说明
- 最低要有：
  - 输入字段
  - 评分逻辑
  - 失败或不确定时怎么处理

<a id="eval-baseline"></a>
### `eval baseline`

- 它是什么：
  - eval loop 第一版稳定结果，用于后续比较质量变化
- 最低要有：
  - dataset 版本
  - grader 版本
  - score
  - 失败样本
  - 解释

<a id="eval-loop"></a>
### `eval loop`

- 它是什么：
  - 从样本、运行、grader、结果解释到 gate 的评估闭环
- 最低要有：
  - dataset
  - runner
  - grader
  - score
  - gate or review action

<a id="grader-strategy"></a>
### `grader strategy`

- 它是什么：
  - grader 如何判断通过、失败、不确定和需要人工复查的策略
- 最低要有：
  - 判定维度
  - 阈值
  - 不确定处理
  - 校准样本

<a id="graders"></a>
### `graders`

- 它是什么：
  - 一组用于不同质量维度的评分器或判定器
- 最低要有：
  - grader 名称
  - 适用样本
  - 输入输出
  - 版本

<a id="dataset-structure"></a>
### `dataset structure`

- 它是什么：
  - eval 或 benchmark 数据集每一行和整体元数据的结构
- 最低要有：
  - input
  - expected behavior
  - labels
  - metadata
  - version

<a id="trace-capture"></a>
### `trace capture`

- 它是什么：
  - 捕获 workflow、tool、policy 和模型调用关键事件的 trace 方案
- 最低要有：
  - span 边界
  - event 字段
  - trace id 传播
  - 采样或保存策略

<a id="trace-id"></a>
### `trace id`

- 它是什么：
  - 串联请求、tool call、approval、audit 和 eval 结果的唯一追踪标识
- 最低要有：
  - 生成位置
  - 传播规则
  - 日志字段
  - 下游调用携带方式

<a id="trace-interpretation"></a>
### `trace interpretation`

- 它是什么：
  - 解释 trace 中关键 span、event 和失败节点的阅读说明
- 最低要有：
  - 正常路径
  - 异常路径
  - 决策点
  - 结论

<a id="trace-path-draft"></a>
### `trace path 草图`

- 它是什么：
  - 一条请求 / workflow / tool path 应被记录哪些 span、event 和 attribute 的草图
- 最低要有：
  - 起点
  - 关键决策点
  - tool / policy / result 节点
  - trace id 传播方式

<a id="trace-sample"></a>
### `trace 样例`

- 它是什么：
  - 一条真实或伪造的 trace 输出示例
- 最低要有：
  - 关键事件顺序
  - 至少一个失败或分支点
  - 一句解释：为什么这条 trace 有判断价值

<a id="baseline-data"></a>
### `baseline 数据`

- 它是什么：
  - 第一轮可信测量结果
- 最低要有：
  - workload
  - 环境
  - 样本量
  - 核心指标
  - 方差或稳定性说明

<a id="before-after-data"></a>
### `before/after 数据`

- 它是什么：
  - 优化前后使用同一口径测得的可比较结果
- 最低要有：
  - 同一 benchmark contract
  - before
  - after
  - 差异解释

<a id="capacity-note"></a>
### `capacity 说明`

- 它是什么：
  - 关于并发、负载、饱和点和成本边界的最小判断
- 最低要有：
  - 假设
  - 数据依据
  - 最脆弱点
  - 何时需要重测

<a id="sli-slo"></a>
### `SLI/SLO`

- 它是什么：
  - 服务指标和目标值的定义
- 最低要有：
  - SLI 名称
  - SLO 目标
  - 测量窗口
  - fail 后动作

<a id="metrics"></a>
### `metrics`

- 它是什么：
  - 系统运行、质量、成本或性能指标集合
- 最低要有：
  - 指标名
  - 采集点
  - 标签
  - 聚合方式
  - 使用场景

<a id="thresholds"></a>
### `thresholds`

- 它是什么：
  - 指标触发 gate、alert 或 review 的阈值集合
- 最低要有：
  - 指标
  - 阈值
  - 依据
  - 动作

<a id="alert-rules"></a>
### `alert rules`

- 它是什么：
  - 指标或事件触发告警的规则
- 最低要有：
  - 条件
  - severity
  - owner
  - responder action

<a id="rollback-decision"></a>
### `rollback decision`

- 它是什么：
  - 判断是否回滚以及如何回滚的决策记录
- 最低要有：
  - 触发信号
  - 影响范围
  - 回滚动作
  - 回滚后验证

<a id="release-policy"></a>
### `release policy`

- 它是什么：
  - 发布前必须满足哪些质量、运行、治理和性能条件的规则
- 最低要有：
  - gate
  - owner
  - fail action
  - override 条件

<a id="optimization-loop"></a>
### `optimization loop`

- 它是什么：
  - 围绕瓶颈提出假设、做单变量实验、测量并决策的优化闭环
- 最低要有：
  - baseline
  - hypothesis
  - change
  - after result
  - decision

<a id="release-gate-rule"></a>
### `release gate 规则`

- 它是什么：
  - 哪些指标退化会阻断发布的规则
- 最低要有：
  - 指标
  - 阈值
  - fail 后动作
  - owner

<a id="benchmark-contract"></a>
### `benchmark contract`

- 它是什么：
  - 冻结 benchmark workload、环境、指标和可重复性假设的契约
- 最低要有：
  - workload
  - environment
  - metrics
  - sample size
  - rerun 条件

<a id="benchmark-dataset"></a>
### `benchmark dataset`

- 它是什么：
  - benchmark 使用的输入样本或负载数据集
- 最低要有：
  - 样本来源
  - 数据规模
  - 版本
  - 代表性说明

<a id="latency-baseline"></a>
### `latency baseline`

- 它是什么：
  - 在固定 workload 和环境下测得的延迟基准
- 最低要有：
  - p50 / p95 / p99
  - workload
  - 环境
  - 采集方式

<a id="throughput-baseline"></a>
### `throughput baseline`

- 它是什么：
  - 系统在固定条件下能稳定处理的吞吐基准
- 最低要有：
  - QPS / jobs per minute
  - 并发
  - 饱和信号
  - error rate

<a id="token-cost"></a>
### `token cost`

- 它是什么：
  - 模型调用或 Agent workflow 的 token 消耗记录
- 最低要有：
  - 输入 token
  - 输出 token
  - tool 或 retry 影响
  - workload 绑定

<a id="capacity-assumption"></a>
### `capacity assumption`

- 它是什么：
  - 容量规划背后的负载、并发、成本和增长假设
- 最低要有：
  - 目标 workload
  - 增长假设
  - 瓶颈假设
  - 何时需要重测

<a id="final-benchmark"></a>
### `final benchmark`

- 它是什么：
  - 年度或阶段收尾时用于证明平台性能边界的 benchmark 证据包
- 最低要有：
  - benchmark contract
  - baseline
  - final result
  - 解释
  - 残余风险

<a id="year-retro"></a>
### `年度复盘`

- 它是什么：
  - 对全年成果、最强证据、最大残余风险和下一阶段重点的压缩总结
- 最低要有：
  - 最稳定的判断
  - 最强证据
  - 最大风险
  - 下一阶段重点

<a id="risk-list"></a>
### `风险清单`

- 它是什么：
  - 对当前阶段最关键风险的排序和说明
- 最低要有：
  - 风险项
  - 为什么重要
  - 当前证据状态
  - 下一步动作

## 文档与测试集类

<a id="tool-gateway-test-suite"></a>
### `Tool Gateway 测试集`

- 它是什么：
  - gateway 的正向、负向、timeout、deny path 相关测试集合
- 最低要有：
  - contract 正例
  - contract 负例
  - timeout / auth / audit 关键路径

<a id="mcp-test-suite"></a>
### `MCP 测试集`

- 它是什么：
  - MCP discovery、invocation、unsupported capability 和 error shape 相关测试集合
- 最低要有：
  - discovery
  - invocation
  - unsupported
  - compatibility or error shape

<a id="e2e-test-suite"></a>
### `E2E 测试集`

- 它是什么：
  - 跨多个子系统的真实旅程验证集合
- 最低要有：
  - 正常旅程
  - 至少一个失败旅程
  - 关键恢复路径

<a id="orchestration-tests"></a>
### `Orchestration tests`

- 它是什么：
  - 验证多 Agent、handoff、fallback、approval 和 shared context 的编排测试
- 最低要有：
  - 正常 handoff
  - fallback
  - approval allow / deny
  - context propagation

<a id="multi-agent-test-suite"></a>
### `多 Agent 测试集`

- 它是什么：
  - 验证多 Agent role、handoff、shared context、fallback 和 approval 的测试集合
- 最低要有：
  - role routing
  - handoff
  - shared context propagation
  - approval allow / deny
  - failure or fallback path

<a id="fallback-tests"></a>
### `fallback tests`

- 它是什么：
  - 验证系统在下游失败、模型失败或 tool 失败后如何退回的测试集合
- 最低要有：
  - 触发失败
  - fallback 条件
  - fallback 输出
  - 原始错误证据

<a id="misuse-tests"></a>
### `misuse tests`

- 它是什么：
  - 验证系统面对误用、越权或 prompt injection 类攻击时的控制效果
- 最低要有：
  - 攻击或误用样本
  - 预期拒绝或降级
  - audit / trace 证据
  - 未发生副作用断言

<a id="adversarial-cases"></a>
### `一组 adversarial case`

- 它是什么：
  - 用来验证 prompt injection、越权、数据泄漏或 tool misuse 防护的攻击/误用样本
- 最低要有：
  - 攻击意图
  - 输入样本
  - 预期系统行为
  - control / guardrail 证据
  - 误伤风险说明

<a id="failure-drills"></a>
### `failure drills`

- 它是什么：
  - 用真实或近似真实的故障场景验证恢复、告警和运行说明
- 最低要有：
  - 故障注入方式
  - 观测信号
  - 恢复动作
  - drill 结论

## 平台、运维与交接类

<a id="operator-guide"></a>
### `operator guide`

- 它是什么：
  - 面向运行维护者说明如何启动、观察、排障和恢复系统的文档
- 最低要有：
  - 启动 / 停止
  - 健康检查
  - 常见故障
  - 指标与日志入口
  - 恢复步骤

<a id="architecture-docs"></a>
### `architecture docs`

- 它是什么：
  - 说明系统边界、控制流、数据流、部署视角和关键取舍的架构文档集合
- 最低要有：
  - 边界图
  - 主链路
  - 关键 contract
  - 风险与证据索引

<a id="developer-onboarding"></a>
### `developer onboarding`

- 它是什么：
  - 新开发者接手代码、运行测试、理解边界和开始贡献的入门说明
- 最低要有：
  - 环境准备
  - 启动命令
  - 测试命令
  - 目录导览
  - 第一项安全改动建议

<a id="runtime-docs"></a>
### `runtime docs`

- 它是什么：
  - runtime 生命周期、状态、取消、错误传播和资源 owner 的说明
- 最低要有：
  - start / stop
  - context owner
  - dependency owner
  - shutdown 行为

<a id="graceful-shutdown"></a>
### `graceful shutdown`

- 它是什么：
  - 系统停止时先停止接新工作、再处理在途工作、最后释放资源的生命周期语义
- 最低要有：
  - 停止信号来源
  - 停拉新
  - 在途处理策略
  - 超时行为
  - 资源释放顺序

<a id="worker-lifecycle"></a>
### `worker 生命周期`

- 它是什么：
  - worker 从启动、领取任务、执行、ack/release 到退出的完整生命周期
- 最低要有：
  - start
  - claim
  - execute
  - ack / release / retry
  - shutdown

<a id="dead-letter-appearance"></a>
### `dead-letter 何时出现`

- 它是什么：
  - 说明任务在什么条件下不再重试，而进入 dead-letter 隔离区
- 最低要有：
  - terminal failure 条件
  - 最大重试或不可恢复错误
  - operator 可见信号
  - 后续处置方式

<a id="timeout-policy"></a>
### `timeout policy`

- 它是什么：
  - 定义调用、任务或 tool 执行的时间预算和超时后行为的策略
- 最低要有：
  - timeout owner
  - duration / budget
  - cancel propagation
  - error mapping
  - fallback or fail behavior

<a id="service-startup-path"></a>
### `服务启动路径`

- 它是什么：
  - 服务从命令入口到 bootstrap、runtime 和 protocol server ready 的启动链路
- 最低要有：
  - command entry
  - config / logger 初始化
  - runtime 初始化
  - server bind
  - ready signal

<a id="workflow-docs"></a>
### `workflow docs`

- 它是什么：
  - workflow state、节点、checkpoint、resume 和人工介入路径的说明
- 最低要有：
  - state model
  - transition
  - checkpoint
  - resume / replay
  - failure mode

<a id="governance-docs"></a>
### `governance docs`

- 它是什么：
  - policy、approval、audit、risk taxonomy 和 misuse controls 的治理文档集合
- 最低要有：
  - 高风险动作
  - 控制点
  - 审批路径
  - 审计事件
  - 测试或 review 证据

<a id="eval-ops"></a>
### `eval ops`

- 它是什么：
  - eval 如何运行、解释、归档和接入 release decision 的运维说明
- 最低要有：
  - 运行命令
  - 数据集版本
  - grader 版本
  - 结果解释
  - gate 关系

<a id="recovery-readiness"></a>
### `recovery readiness`

- 它是什么：
  - 系统是否具备恢复演练和真实恢复能力的准备度说明
- 最低要有：
  - 已覆盖 failure mode
  - 未覆盖 failure mode
  - drill 证据
  - operator 缺口

<a id="deployment-topology"></a>
### `deployment topology`

- 它是什么：
  - 系统部署单元、进程、网络入口和依赖关系的拓扑说明
- 最低要有：
  - 部署单元
  - service / endpoint
  - config / secret
  - health probe
  - rollout / rollback 假设

<a id="manifests"></a>
### `manifests`

- 它是什么：
  - 部署、服务、配置、探针和资源限制等运行清单
- 最低要有：
  - deployment
  - service
  - config / secret 引用
  - probes
  - resources

<a id="probes"></a>
### `probes`

- 它是什么：
  - startup、readiness、liveness 等健康探针定义
- 最低要有：
  - probe 类型
  - 检查路径
  - timeout
  - failure threshold
  - 对应系统状态

<a id="resource-limits"></a>
### `resource limits`

- 它是什么：
  - 部署时对 CPU、内存或其他资源的请求和限制
- 最低要有：
  - resource 类型
  - request
  - limit
  - 依据
  - 超限行为

<a id="platform-boundary"></a>
### `platform boundary`

- 它是什么：
  - 平台对外提供什么能力、对内拥有哪些控制面的边界说明
- 最低要有：
  - 对外能力
  - 内部子系统
  - shared contracts
  - 不负责的范围

<a id="subsystem-map"></a>
### `subsystem map`

- 它是什么：
  - 平台子系统、职责和依赖方向的地图
- 最低要有：
  - 子系统
  - owner
  - 输入输出
  - 依赖

<a id="integration-order"></a>
### `integration order`

- 它是什么：
  - 多子系统整合时的顺序和验证计划
- 最低要有：
  - step
  - dependency
  - validation
  - rollback point

<a id="config-integration"></a>
### `config integration`

- 它是什么：
  - config center 与 runtime、tool gateway、MCP 或部署清单的配置整合说明
- 最低要有：
  - 配置来源
  - consumer
  - validation
  - reload or restart 假设

<a id="tool-gateway-integration"></a>
### `tool gateway integration`

- 它是什么：
  - Tool Gateway 与 Agent、MCP、config、audit、trace 的整合说明
- 最低要有：
  - invocation path
  - schema source
  - timeout / auth / audit
  - error mapping

<a id="platform-health"></a>
### `platform health`

- 它是什么：
  - 从运行、质量、治理、性能四个维度判断平台成熟度的状态
- 最低要有：
  - 运行信号
  - eval / gate 信号
  - governance 信号
  - benchmark 信号
  - 最大风险

<a id="year-retrospective"></a>
### `year retrospective`

- 它是什么：
  - 年度复盘，压缩全年系统能力、证据和下一阶段方向
- 最低要有：
  - 成熟能力
  - 最强证据
  - 主要风险
  - 下一阶段重点

<a id="graduation-evidence"></a>
### `graduation evidence`

- 它是什么：
  - 证明学习者已经能独立解释和交付平台级系统的证据集合
- 最低要有：
  - 架构证据
  - 运行证据
  - 质量证据
  - 治理证据
  - 答辩结论

<a id="next-roadmap"></a>
### `next roadmap`

- 它是什么：
  - 下一阶段学习或建设路线图
- 最低要有：
  - 主题
  - 项目
  - 证据目标
  - 选择理由

<a id="resource-boundary"></a>
### `resource boundary`

- 它是什么：
  - 说明 CPU、内存、连接、队列、模型调用等资源由谁拥有和限制
- 最低要有：
  - 资源类型
  - owner
  - limit / request
  - 超限行为

<a id="health-model"></a>
### `health model`

- 它是什么：
  - 系统健康状态和探针含义的模型
- 最低要有：
  - startup / readiness / liveness 区分
  - 依赖状态
  - degraded 状态
  - operator action

<a id="review-package"></a>
### `review package`

- 它是什么：
  - 一次阶段性评审所需的 RFC、实现、测试、证据、风险和下一步集合
- 最低要有：
  - 结论
  - 证据链接
  - 主要风险
  - 阻断项
  - 下一步

## 最后一句

如果一个交付物名字你能念出来，但说不清“它至少包含什么”，那你还没有真正理解它。
