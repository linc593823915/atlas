# 第四个月：Tool Gateway 与接口质量

最后更新：2026-06-30

## 本月目标

从这个月开始，你不再满足于“工具能调用”。

你要让工具调用具备平台边界。

这意味着要把 Tool 统一成一个可治理的入口。

## 本月学完后你应该会什么

你应该能解释：

- 为什么 Tool 不能散在不同业务模块里
- 为什么 schema、validation、timeout、audit 必须统一
- 为什么 Tool Gateway 是平台边界，而不是简单 wrapper

## 本月学习重点

### 1. Tool Gateway 边界
理解：
- 为什么所有工具调用都应该统一入口。
- request / result envelope 应该包含什么。

### 2. Registry 与 Schema Validation
理解：
- tool schema 的 owner 是谁。
- 执行前和执行后分别验证什么。

### 3. Timeout、Retry、Audit Hook
理解：
- 统一策略为什么属于 gateway contract。
- gateway 如何成为治理和观测面。

### 4. Contract Test 与接口质量
理解：
- 为什么负路径比 happy path 更能证明 gateway 价值。
- review 应该盯哪些 failure mode。

## 本月 Cookbook 做法

### Recipe 1：定义 Tool Gateway RFC

### Recipe 2：做出 Registry 与 Schema Validation

### Recipe 3：加入 Timeout / Retry / Audit Hook

### Recipe 4：补 Contract Tests

## 本月交付标准

到本月结束时，你至少要能解释：

- 新工具如何接入
- 工具如何被统一验证
- 工具超时或失败时系统如何处理
- 为什么这比“直接写函数调用”更适合平台化

## 本月必学术语

- `Tool Gateway`：统一接管工具调用的入口层，负责 validation、timeout、audit 等控制。
- `Registry`：保存 tool schema、version 和 discoverability 信息的注册中心。
- `Schema Validation`：在执行前后检查输入输出是否仍满足约束的校验机制。
- `Timeout Policy`：定义谁负责超时预算、取消传播和失败后行为的规则。
- `Audit Hook`：在关键调用前后写入可追踪决策证据的钩子。

## 本月知识地图

- 第四个月的关键不是多一层抽象，而是把工具调用真正收回到一个可治理的入口。
- gateway 存在的意义是统一 schema、timeout、auth、audit 和 result normalization。
- registry 决定 contract 的 discoverability、versioning 和 ownership。
- 接口质量要靠 contract tests 和 negative tests 证明，而不是靠“设计上应该没问题”。

## 本月常见误区

- 把 gateway 写成薄路由器。
- 让 schema 分散在调用方和实现方两边。
- 只测工具能成功调用，不测拒绝和失败路径。

## 本月最小演示场景

- 背景：同一个工具请求从统一 gateway 进入：正常输入通过，非法输入被拒绝，关键调用带上 timeout 和 audit 元数据。
- 你至少要能演示：
  - tool 请求如何经过 gateway
  - schema/validation 在哪一层执行
  - timeout / audit 为什么必须统一
- reviewer 大概率会追问：
  - gateway 真正收回了哪些控制权？
  - 哪些校验应该在执行前做，哪些在执行后做？
  - 负路径测试是否比 happy path 更说明问题？

## 本月案例入口

- 先读 [专题案例包](casebook.md) 中的“案例 4：工具越来越多，但系统越来越难治理”，把本月知识点放回真实工程场景。
- 再读 [失败场景图谱](failure-patterns.md) 中的“模式一：边界缺失型故障”，理解这一月最容易反复出现的故障模式。
- 读完后回到本月的周/日文档，检查自己能不能说清楚：错误做法、坏信号和最小证据分别是什么。

## 本月实战清单

- [`RFC`](artifact-playbook.md#artifact-rfc)：写 Tool Gateway RFC，冻结 request/result envelope 和控制面职责。
- [`Interface`](artifact-playbook.md#artifact-interface)：定义 registry、schema、validation、timeout、audit hook 接口。
- [`Implementation`](artifact-playbook.md#artifact-implementation)：把工具调用统一接入 gateway，并打通最小调用链。
- [`Unit Test`](artifact-playbook.md#artifact-unit-test)：覆盖 contract tests、negative tests、timeout 与 deny path。
- [`Review`](artifact-playbook.md#artifact-review)：检查 tool contract 是否仍在 gateway 外散落。
- [`Benchmark（按需）`](artifact-playbook.md#artifact-benchmark)：本月不要求性能 benchmark，但要有接口稳定性和失败路径验证证据。
- [`Documentation`](artifact-playbook.md#artifact-documentation)：补 tool 接入指南、gateway contract 和 review note。
- [`Memory`](artifact-playbook.md#artifact-memory)：沉淀为什么 schema、timeout、audit 要在同一入口收口。

## 本月评分标准

- 满分 `100` 分。
- 结果对齐本月目标：20 分。
- 技术深度与判断：20 分。
- 实现质量与代码健康：15 分。
- 测试、可靠性与失败思维：15 分。
- 系统设计清晰度：10 分。
- 文档与沟通质量：10 分。
- 学习、反馈与复盘闭环：10 分。
- 及格线：`80+`，且核心产物无缺项。
- 优秀线：`90+`，并能证明 Atlas 的能力超出了本月最小范围。

## 本月证据清单

- 本月目标和关键结果已经写清楚，并能回头打分。
- 本月每周交付物都存在、可定位、可复查。
- 本月核心改动有测试、trace、drill 或其他验证证据。
- 文档、运行说明和复盘结论已经反映当前 Atlas 状态。
- 至少有一份自评或评审总结，能指出后续改进动作。

## 本月面试追问

1. 如果要向资深 reviewer 解释“Tool Gateway 与接口质量”，你会怎样说明这月的架构取舍？
2. 本月哪条主路径最容易失败？你会如何设计验证证据？
3. 本月引入的哪个接口、状态或契约最可能在未来变化？为什么？
4. 如果 Atlas 的调用量增长 10 倍，本月设计会先在哪个点上断裂？
5. 本月消除了哪些技术债，又留下了哪些新债？
6. 哪个周度产物最能证明你的工程判断？为什么？
7. 你会怎样把本月所有周主题串成一个完整系统故事，而不是零散任务？
8. 你会出示哪些测试、trace、日志或运行证据，证明本月成果已经可交付？

## 本月推荐配套阅读

- [Month 04](../months/month-04.md)
- [Month 04 Backlog](../atlas/backlog/month-04.md)
- [Month 04 Issue Set](../atlas/issues/month-04/README.md)
- [第二学期日学习索引（第一个月）](day-index-semester-2-part-1.md)
