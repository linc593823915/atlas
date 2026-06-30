# 日级常见产物示例

最后更新：2026-06-30

## 这份文档解决什么问题

day 文档里反复出现很多“今天最重要交付”类短语：

- `命令如何启动 Atlas`
- `配置从哪里进来`
- `Docker 如何跑`
- `SLO 草案`
- `下一阶段路线图`

这些词比月度或周度产物更小，但更容易写空。

这份文档就是告诉你：

这些“日级产物”最少应该长什么样。

## 第一月常见产物

<a id="how-atlas-starts"></a>
### `命令如何启动 Atlas`

- 最低应该包含：
  - 启动命令
  - 入口文件
  - 启动顺序
  - 成功信号
  - 失败时第一步怎么查

### 一句合格示例

```md
当前 Atlas 通过 `atlas serve` 进入启动主路径，入口依次经过 `cmd/atlas/main.go -> internal/app/root.go -> internal/app/serve.go -> internal/bootstrap/bootstrap.go`，看到成功日志后说明 bootstrap 已完成最小启动。
```

<a id="where-config-enters"></a>
### `配置从哪里进来`

- 最低应该包含：
  - 配置源顺序
  - 谁是唯一入口
  - 错误配置如何失败

### 一句合格示例

```md
当前配置统一从 `config.Manager` 进入，顺序是 `default -> config file -> environment`，坏配置在 `validateConfig` 阶段尽早失败，而不是等业务路径出错。
```

<a id="how-logging-is-configured"></a>
### `日志如何配置`

- 最低应该包含：
  - level / format
  - 初始化位置
  - source 是否保留

<a id="when-dependencies-init"></a>
### `依赖在什么时候初始化`

- 最低应该包含：
  - config / logger / runtime 的顺序
  - db/cache 未来应该接到哪层

<a id="how-docker-runs"></a>
### `Docker 如何跑`

- 最低应该包含：
  - build 命令
  - run 命令
  - 容器内启动路径与本地是否一致
  - 容器失败时最先看哪里

<a id="how-ci-passes"></a>
### `CI 如何过`

- 最低应该包含：
  - 最少跑哪些步骤
  - 哪类错误会让 CI fail
  - CI 结果对应哪类交付质量

<a id="month-one-review-summary"></a>
### `第一月有什么完成了，什么没完成`

- 最低应该包含：
  - 已稳定的边界
  - 仍然脆弱的边界
  - 下月最先要补哪条债

<a id="rfc-draft"></a>
### `一份 RFC 草稿`

- 最低应该包含：
  - 问题定义
  - 目标与非目标
  - 候选方案和取舍
  - 推荐方案
  - 最小验收口径

一句合格标准：别人读完后能复述“为什么要做、为什么这样做、失败时先看哪里”。

<a id="package-structure-draft"></a>
### `一份包结构初稿`

- 最低应该包含：
  - 每个包的职责
  - 调用方向
  - 不允许跨过去的边界
  - 未来可能扩展的位置

不要只画目录树。包结构草稿真正要证明的是 ownership 和依赖方向。

<a id="learning-log-entry"></a>
### `一份学习日志入口`

- 最低应该包含：
  - 今天处理的问题
  - 读过的证据
  - 做出的判断
  - 留下的风险
  - 下一步动作

学习日志不是流水账。它要让未来的你能恢复当时为什么这样判断。

<a id="entry-runtime-boundary-note"></a>
### `CLI 与运行时入口要清楚分离`

- 最低应该包含：
  - CLI 只负责解析命令和参数
  - bootstrap 负责组装依赖
  - runtime 负责运行生命周期
  - `main` 不直接拥有业务依赖

这条笔记要能回答：如果未来增加 worker、MCP server 或后台任务，入口层会不会被迫重写。

<a id="service-shutdown-note"></a>
### `服务如何退出`

- 最低应该包含：
  - 谁接收停止信号
  - 谁停止接收新请求或新任务
  - 谁等待在途工作完成
  - 超时后如何失败
  - 日志里留下什么证据

服务退出不是“进程结束”。它是系统最后一次兑现生命周期 contract。

<a id="health-check-proof"></a>
### `健康检查是最小系统可运行性的证明`

- 最低应该包含：
  - health endpoint 或检查命令
  - 它证明了哪些依赖已经就绪
  - 它没有证明哪些依赖
  - 失败时 operator 先看哪里

健康检查不能只返回 `200`。它要表达“当前系统能承接什么类型的工作”。

<a id="engineering-hygiene-note"></a>
### `工程卫生为什么从第一月开始就要建立`

- 最低应该包含：
  - 哪些门禁从第一月就固定
  - 哪些坏习惯以后迁移成本最高
  - 哪条证据说明当前流程可复用

工程卫生不是仪式感。它是在系统还小时保护未来变更速度。

<a id="problem-first-before-work"></a>
### `学会把问题先写清楚，再开工`

- 最低应该包含：
  - 当前问题
  - 不解决会怎样
  - 本轮目标
  - 明确的非目标
  - 最小证据

这类笔记要训练的是工程判断顺序：先定义问题，再进入实现。

<a id="learning-system-setup"></a>
### `建立学习日志、Issue 清单、架构笔记`

- 最低应该包含：
  - 学习日志入口
  - issue 编号和状态
  - 架构笔记位置
  - 每个入口如何被后续复盘引用

学习系统不是文件夹集合，而是让证据能被追回来的工作流。

<a id="atlas-minimal-module-boundary"></a>
### `明确 Atlas 的最小模块边界`

- 最低应该包含：
  - 入口层
  - bootstrap 层
  - runtime 层
  - config / logger 边界
  - 当前不进入业务层的内容

模块边界要能解释调用方向，不能只解释目录命名。

<a id="db-cache-infra-boundary-note"></a>
### `数据库和缓存为什么不能散在业务代码里`

- 最低应该包含：
  - 连接生命周期 owner
  - 超时和重试边界
  - 配置来源
  - shutdown 行为
  - 测试替身或接口位置

db/cache 一旦散在业务代码里，后续测试、恢复和资源治理都会变得不可控。

<a id="logger-abstraction-note"></a>
### `日志为什么必须统一抽象`

- 最低应该包含：
  - logger 初始化位置
  - level / format / source 规则
  - request 或 trace 字段如何注入
  - 禁止各层自建 logger 的理由

统一日志抽象是为了让运行证据可比较，而不是为了统一输出风格。

<a id="tests-are-deliverable-note"></a>
### `测试不是最后补，而是交付的一部分`

- 最低应该包含：
  - 本轮不变量
  - 正向测试
  - 失败路径测试
  - 测试结果如何进入 review

测试如果只在最后补，就很难保护设计时真正关心的边界。

<a id="minimal-run-document"></a>
### `一份最小运行文档`

- 最低应该包含：
  - 启动命令
  - 必需配置
  - 成功信号
  - 常见失败
  - 停止方式

最小运行文档的读者是假设没有作者上下文的人。

<a id="initial-analysis-conclusion"></a>
### `初版分析结论`

- 最低应该包含：
  - 观察到的信号
  - 初步判断
  - 证据强度
  - 不能下结论的部分
  - 下一步验证动作

初版分析结论必须标出不确定性，不能伪装成最终判断。

<a id="integration-document"></a>
### `接入文档`

- 最低应该包含：
  - 接入方是谁
  - 接入入口
  - 输入输出 contract
  - 配置要求
  - 错误和兼容性说明

接入文档要让调用方不看实现也能正确使用。

## Agent / 协议类常见产物

<a id="runnable-mcp-shell"></a>
### `可运行的 MCP 服务壳`

- 最低应该包含：
  - 服务如何启动
  - transport 是什么
  - 至少一个 capability 如何被发现或调用

<a id="mcp-doc-note"></a>
### `MCP 文档`

- 最低应该包含：
  - capability 列表
  - transport 说明
  - compatibility / error shape 简述

<a id="runnable-agent-definitions"></a>
### `可运行的 agent definitions`

- 最低应该包含：
  - 角色名
  - 输入输出
  - 允许的工具
  - handoff / trace 相关字段

<a id="explainable-handoff-flow"></a>
### `可解释的 handoff 流程`

- 最低应该包含：
  - handoff 条件
  - 上下文如何传播
  - 谁对结果负责

<a id="agent-tool-call-path"></a>
### `Agent 怎么调 Tool`

- 最低应该包含：
  - Agent 如何选择 tool
  - tool 输入 schema 从哪里来
  - tool 执行结果如何回到 Agent
  - schema、timeout、audit 三类失败如何处理

合格写法不是“Agent 调用了工具”，而是能解释 tool invocation 的完整责任链。

<a id="system-usable-output-note"></a>
### `为什么这个输出是“系统可用”的`

- 最低应该包含：
  - 输出契约
  - 解析方式
  - 失败时的降级或报错
  - 一条能复跑的验证证据

“看起来对”不等于系统可用。系统可用意味着调用方能稳定消费它。

<a id="routing-rationale-note"></a>
### `为什么路由逻辑合理`

- 最低应该包含：
  - 路由输入
  - 候选路径
  - 选择条件
  - fallback 或 deny path
  - 错误归因

路由逻辑的价值在于可解释，而不是命中一次 happy path。

<a id="fallback-behavior-note"></a>
### `失败后系统会怎么退回`

- 最低应该包含：
  - 失败类型
  - fallback 条件
  - 是否保留原始错误
  - fallback 后的输出 contract
  - operator 如何识别这次 fallback

fallback 不是吞掉错误。它要让系统在退回后仍然可解释。

<a id="output-contract-draft"></a>
### `输出契约草稿`

- 最低应该包含：
  - 输出字段
  - 必填 / 可选
  - 失败输出
  - 解析方是谁
  - 版本演进假设

输出契约应该先服务调用方，再服务实现方。

<a id="schema-validation-path"></a>
### `schema 校验路径`

- 最低应该包含：
  - schema 来源
  - 校验发生在哪一层
  - 校验失败如何表达
  - schema 版本如何演进

如果校验只在最后补一层，它通常保护不了真正的 contract。

<a id="registered-tool-structure"></a>
### `可注册的 Tool 结构`

- 最低应该包含：
  - tool 名称
  - 输入 / 输出 schema
  - owner
  - timeout / auth / audit 要求
  - 兼容性状态

“可注册”意味着它能被发现、验证和治理，而不只是能被调用。

<a id="error-model-draft"></a>
### `错误模型草稿`

- 最低应该包含：
  - 错误分类
  - 调用方可见字段
  - retryable / terminal 区分
  - 日志或 trace 证据

错误模型要帮助调用方采取动作，而不是只帮助实现方打印异常。

<a id="tool-governance-path"></a>
### `可治理的 Tool 执行路径`

- 最低应该包含：
  - schema 校验
  - timeout
  - auth / approval
  - audit event
  - 失败归因

工具执行一旦进入 Agent 路径，就必须同时考虑能力、风险和证据。

<a id="tool-gateway-boundary-note"></a>
### `工具网关职责图`

- 最低应该包含：
  - gateway 接收什么
  - gateway 不接管什么
  - 下游 tool wrapper 的责任
  - 统一错误面
  - 审计和 trace 字段

职责图要能防止 tool 逻辑散落到 Agent、MCP handler 和业务代码里。

<a id="mcp-capability-entry"></a>
### `最小 capability 入口`

- 最低应该包含：
  - capability 名称
  - 类型
  - 输入输出
  - discovery 方式
  - unsupported 时的返回

最小 capability 不是最小功能，而是最小外部 contract。

<a id="protocol-error-path"></a>
### `协议错误路径`

- 最低应该包含：
  - bad request
  - unsupported capability
  - transient failure
  - internal failure
  - client 应该如何处理

协议错误路径要以 client 的退化动作为中心。

<a id="compatibility-review-note"></a>
### `Compatibility review`

- 最低应该包含：
  - 哪些行为保持兼容
  - 哪些行为弃用
  - 旧 client 如何迁移
  - 测试证据

兼容性 review 的目标不是承诺永远不变，而是让变化可预期。

<a id="transport-model-note"></a>
### `transport model`

- 最低应该包含：
  - transport 类型
  - 连接生命周期
  - request dispatch
  - shutdown 语义
  - 错误如何返回给 client

transport model 会影响 runtime owner，因此不能只当通信细节。

## 工作流 / 治理类常见产物

<a id="pausable-recoverable-flow"></a>
### `可暂停恢复流程`

- 最低应该包含：
  - pause 条件
  - resume 从哪继续
  - 人工介入如何并回

<a id="approval-chain"></a>
### `审批链路`

- 最低应该包含：
  - 谁发起
  - 谁批准
  - 什么时候会被拒绝
  - 审批记录如何留痕

<a id="audit-chain"></a>
### `审计链路`

- 最低应该包含：
  - actor
  - action
  - decision
  - trace / approval 关联

<a id="pausable-recoverable-flow-note"></a>
### `可暂停恢复流程`

- 最低应该包含：
  - pause 条件
  - checkpoint 保存内容
  - resume cursor
  - 人工介入如何合并
  - 不能重放的副作用

可恢复不是“失败后再跑一次”，而是恢复到可解释的下一步。

<a id="human-intervention-path"></a>
### `人工介入路径`

- 最低应该包含：
  - 触发人工介入的条件
  - 人的输入被记录在哪里
  - 审批或修改如何进入 workflow state
  - trace / audit 如何保留 provenance

人工介入如果没有 merge 语义，就会变成系统外的手工补丁。

<a id="approval-boundary-draft"></a>
### `审批边界草稿`

- 最低应该包含：
  - 哪些动作必须审批
  - 审批前系统停在哪里
  - deny path 如何返回
  - 审批结论如何进入 audit

审批边界要保护高风险动作，不是给所有动作增加人工确认。

<a id="approval-example"></a>
### `approval 示例`

- 最低应该包含：
  - 一个具体动作
  - 审批请求字段
  - allow / deny 两种结果
  - 对 workflow 的影响

示例必须包含拒绝路径，否则它只能证明 happy path。

<a id="audit-field-plan"></a>
### `审计字段方案`

- 最低应该包含：
  - actor
  - action
  - resource / scope
  - decision
  - reason
  - trace / approval 关联

审计字段要支持事后重建决策，而不是只记录动作发生过。

<a id="audit-note"></a>
### `审计说明`

- 最低应该包含：
  - 哪些事件会被记录
  - 哪些字段被脱敏或省略
  - operator 如何查询
  - 审计记录不能证明什么

合格的审计说明要同时服务安全、排障和合规复查。

<a id="action-classification-table"></a>
### `动作分类表`

- 最低应该包含：
  - action 名称
  - 风险等级
  - allow / deny / approval 规则
  - audit 要求

动作分类表是 policy engine 的输入，不是安全文档里的装饰。

<a id="risk-boundary-note"></a>
### `风险边界说明`

- 最低应该包含：
  - 风险发生在哪条路径
  - 当前已有控制点
  - 仍然缺的证据
  - 下一步降低风险的动作

风险边界说明要落到系统路径，不能只列抽象风险名。

<a id="high-risk-action-list"></a>
### `高风险动作清单`

- 最低应该包含：
  - 动作
  - 风险原因
  - 需要的控制点
  - 审批和审计要求

高风险动作必须能映射到真实执行路径。

<a id="governance-document"></a>
### `治理文档`

- 最低应该包含：
  - policy 范围
  - 高风险动作
  - approval / audit 关系
  - 测试或演练证据

治理文档要说明系统如何被约束，而不是只表达原则。

<a id="recovery-document"></a>
### `恢复文档`

- 最低应该包含：
  - failure mode
  - 检测信号
  - 恢复步骤
  - 验证方式
  - 残余风险

恢复文档要让 operator 能执行，而不是让作者回忆。

## 评估 / 性能 / 平台类常见产物

<a id="slo-draft"></a>
### `SLO 草案`

- 最低应该包含：
  - 指标名
  - 目标值
  - 谁关心它
  - fail 后动作

<a id="alert-rules"></a>
### `告警规则`

- 最低应该包含：
  - 触发条件
  - severity
  - responder action

<a id="health-check-config"></a>
### `健康检查配置`

- 最低应该包含：
  - readiness / liveness / startup 各自检查什么
  - 检查间隔
  - timeout
  - failure threshold
  - 失败后的 operator 动作

健康检查配置要反映系统状态，而不是照抄平台默认值。

<a id="resource-limit-note"></a>
### `资源限制说明`

- 最低应该包含：
  - CPU / memory / connection / queue 等资源
  - request / limit 或容量假设
  - 超限行为
  - 监控信号
  - 何时需要重新评估

资源限制说明是在写运行假设，不是写部署参数。

<a id="rollback-note"></a>
### `回滚说明`

- 最低应该包含：
  - 回滚条件
  - 回滚动作
  - 回滚后验证

<a id="next-stage-roadmap"></a>
### `下一阶段路线图`

- 最低应该包含：
  - 下阶段主题
  - 关键项目
  - 证据目标
  - 为什么是这些，而不是别的

<a id="dataset-draft"></a>
### `数据集初稿`

- 最低应该包含：
  - 样本输入
  - 期望行为
  - 标签或分类
  - grader 所需字段
  - 版本

数据集初稿要先支持复查，再追求规模。

<a id="basic-eval-samples"></a>
### `一组基础 eval 样本`

- 最低应该包含：
  - 覆盖主路径的样本
  - 至少一个负例
  - 预期输出或判定标准
  - 为什么这些样本能代表当前风险

基础 eval 样本不是展示模型能力，而是固定系统质量口径。

<a id="grader-design-draft"></a>
### `grader 设计草稿`

- 最低应该包含：
  - grader 输入
  - 输出分数或标签
  - 通过 / 失败阈值
  - 不确定时如何处理

grader 要能解释为什么判定失败，而不是只给一个分数。

<a id="regression-judgment-logic"></a>
### `回归判断逻辑`

- 最低应该包含：
  - baseline
  - 当前结果
  - 阈值
  - 允许波动范围
  - fail 后动作

回归判断必须服务发布或改动决策。

<a id="metric-threshold-note"></a>
### `指标阈值说明`

- 最低应该包含：
  - 指标名
  - 阈值
  - 设定依据
  - 超阈值后的动作

阈值没有动作，就只是展示数字。

<a id="workload-definition"></a>
### `工作负载定义`

- 最低应该包含：
  - 请求类型
  - 并发
  - 数据规模
  - 运行时长
  - 环境假设

workload 决定 benchmark 是否能复跑和比较。

<a id="metric-definition"></a>
### `指标口径`

- 最低应该包含：
  - 指标名
  - 计算方式
  - 采集位置
  - 排除项
  - owner

指标口径要先稳定，数字才有比较意义。

<a id="cost-record"></a>
### `成本记录`

- 最低应该包含：
  - 计费维度
  - workload
  - 单次或单位成本
  - 成本变化原因
  - 优化建议

成本记录必须和真实 workload 绑定。

<a id="bottleneck-analysis"></a>
### `瓶颈分析`

- 最低应该包含：
  - 热路径
  - 观测证据
  - 候选瓶颈
  - 最可能瓶颈
  - 下一次单变量实验

瓶颈分析不是猜最慢的模块，而是用证据缩小范围。

<a id="platform-rfc"></a>
### `平台 RFC`

- 最低应该包含：
  - 平台边界
  - 子系统责任
  - 共享 contract
  - 运行与治理模型
  - 最小证据包

平台 RFC 要把功能集合讲成 operating model。

<a id="subsystem-map"></a>
### `子系统图`

- 最低应该包含：
  - 子系统名称
  - 责任
  - 输入输出
  - 依赖方向
  - 共享 contract

子系统图不是架构装饰图，而是责任边界图。

<a id="integration-order-note"></a>
### `整合顺序说明`

- 最低应该包含：
  - 先整合什么
  - 后整合什么
  - 每一步的验证信号
  - 回退条件

整合顺序说明要降低系统同时变化的风险。

<a id="integration-main-path"></a>
### `整合后的主链路`

- 最低应该包含：
  - 入口
  - 关键子系统
  - 共享 contract
  - 失败路径
  - 端到端证据

主链路要能解释请求如何穿过 config、gateway、runtime 和外部 surface。

<a id="dependency-order-note"></a>
### `依赖顺序说明`

- 最低应该包含：
  - 初始化顺序
  - 谁拥有依赖生命周期
  - 失败时如何停止后续初始化
  - shutdown 逆序

依赖顺序如果写不清，系统退出和恢复通常也写不清。

<a id="platform-skeleton"></a>
### `平台骨架`

- 最低应该包含：
  - 入口层
  - runtime
  - capability / tool 层
  - protocol surface
  - observability / governance 位置

平台骨架要说明未来扩展落在哪里，而不是只说明当前有哪些目录。

<a id="platform-health-summary"></a>
### `平台健康总结`

- 最低应该包含：
  - 运行健康
  - 质量健康
  - 治理健康
  - 性能健康
  - 最大残余风险

平台健康总结应该是证据索引，不是年终感想。

<a id="deployment-checklist"></a>
### `部署清单`

- 最低应该包含：
  - 镜像或制品
  - 配置和 secret
  - 资源限制
  - health probe
  - rollout / rollback 步骤
  - 验证信号

部署清单要能被 operator 执行，不能只列“准备部署”。

<a id="ops-document"></a>
### `运维文档`

- 最低应该包含：
  - 启停方式
  - 健康检查
  - 日志和指标入口
  - 常见故障
  - 恢复与回滚步骤

运维文档的目标是减少线上判断对作者记忆的依赖。

<a id="handover-note"></a>
### `接手说明`

- 最低应该包含：
  - 当前系统怎么启动
  - 主要子系统在哪里
  - 常见失败如何排查
  - 不能随便改的 contract
  - 下一步工作入口

接手说明的读者是没有上下文的新维护者。

<a id="graduation-evidence-package"></a>
### `毕业证据包`

- 最低应该包含：
  - 最强架构证据
  - 最强质量证据
  - 最强运行证据
  - 最强治理证据
  - 下一阶段路线图

毕业证据包要证明系统成熟度，而不是堆全年产物链接。

## 复盘与阶段收口类常见产物

<a id="retrospective-conclusion"></a>
### `复盘结论`

- 最低应该包含：
  - 本阶段最稳定的判断
  - 最强证据
  - 最大残余风险
  - 下一阶段第一件事

复盘不是总结心情，而是校准方向。

<a id="semester-review-note"></a>
### `学期复盘`

- 最低应该包含：
  - 本学期形成的边界
  - 哪些能力已经可复用
  - 哪些风险会影响下一学期
  - 下一学期准备点

学期复盘要回答“系统能力发生了什么变化”。

<a id="next-semester-prep-note"></a>
### `下一学期准备说明`

- 最低应该包含：
  - 下学期主题
  - 必须先补的债
  - 可直接复用的资产
  - 验收证据目标

准备说明要把上一阶段复盘转成下一阶段执行入口。

<a id="technical-debt-list"></a>
### `技术债清单`

- 最低应该包含：
  - 债务项
  - 影响路径
  - 当前证据
  - 优先级
  - 处理建议

技术债不是“不完美列表”，而是会影响后续交付的风险队列。

<a id="risk-note"></a>
### `风险说明`

- 最低应该包含：
  - 风险是什么
  - 触发条件
  - 影响范围
  - 现有缓解措施
  - 下一步证据

风险说明必须能映射到真实系统路径。

<a id="test-result-note"></a>
### `测试结果`

- 最低应该包含：
  - 测试范围
  - 通过和失败项
  - 失败解释
  - 是否阻断后续动作

测试结果要能支撑决策，而不是只说明命令跑过。

## 最后一句

如果一个日级产物你只能写成一句口号，那它还不是交付物。

它至少要回答：

1. 这是什么
2. 为什么重要
3. 怎么判断它是否成立
