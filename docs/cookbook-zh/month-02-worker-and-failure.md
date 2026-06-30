# 第二个月：后台任务与失败处理

最后更新：2026-06-30

## 本月目标

本月的目标是让 Atlas 具备“后台执行能力”。

你要从一个只能同步跑主流程的服务，升级成一个能处理：

- queue
- worker
- retry
- timeout
- idempotency
- dead-letter

的工程系统。

## 本月学完后你应该会什么

你应该能解释：

- 为什么 Agent 系统也离不开后台任务模型
- 为什么 retry 不是“多试几次”这么简单
- 为什么 timeout 和 cancellation 是可靠性设计的一部分
- 为什么幂等在调用外部工具和长链任务里尤其重要

## 本月学习重点

### 1. 任务状态模型
理解：
- queued、running、retryable、failed、dead-letter 各是什么意思。
- 状态迁移如何影响 operator 决策。

### 2. Worker 生命周期
理解：
- claim -> execute -> ack/release 主循环。
- graceful shutdown、lease/heartbeat 和并发控制。

### 3. Retry、Timeout 与幂等
理解：
- 暂时性失败和终局失败怎么区分。
- timeout 预算、幂等 key 和副作用保护。

### 4. 可靠性验证与运维证据
理解：
- failure drill、queue lag、retry count、dead-letter 证据。
- 为什么可靠性要用测试和演练证明。

## 本月 Cookbook 做法

### Recipe 1：定义任务模型

产物：

- job state
- queue abstraction
- worker contract

### Recipe 2：跑通 worker 主循环

产物：

- enqueue
- consume
- graceful shutdown

### Recipe 3：补上失败路径

产物：

- retry policy
- timeout behavior
- dead-letter path

### Recipe 4：验证可靠性

产物：

- worker tests
- drill notes
- failure review

## 本月交付标准

到本月结束时，你至少要能演示：

- 一个任务如何入队和执行
- 一个失败任务如何重试
- 一个硬失败任务如何进入 dead-letter
- 一个重复任务如何被正确处理

## 本月必学术语

- `Job State`：后台任务在排队、执行、重试、失败、死信之间的状态语义。
- `Lease / Visibility Timeout`：任务被 worker 暂时占有的时间窗口，用于避免重复消费。
- `Idempotency Key`：保证重复执行同一任务时不会放大副作用的稳定标识。
- `Dead-letter Queue`：承接多次失败或不可恢复任务的隔离区，用于排障和治理。
- `Graceful Shutdown`：停止拉新、排空在途、保存必要状态后再安全退出的流程。

## 本月知识地图

- 后台任务系统首先是状态和失败语义系统，worker 代码只是这些语义的执行器。
- 没有清楚的重试、超时和幂等语义，任何队列都会把偶发故障放大成系统性事故。
- worker 生命周期必须同时考虑并发、退出、重启和 ownership，否则重复执行与任务丢失会同时出现。
- 第二个月的核心不是“让任务跑起来”，而是让任务在失败时依然可解释、可恢复、可运维。

## 本月常见误区

- 先写 worker 后补状态机。
- 把 retry 当默认正确答案，不区分错误类型。
- 只做 happy-path 验证，不做 failure drill。

## 本月最小演示场景

- 背景：系统收到一个后台任务：正常时一次成功，异常时会超时、重试，最终失败任务被送入 dead-letter。
- 你至少要能演示：
  - 一个 job 的状态迁移路径
  - 超时和重试时系统如何避免副作用被放大
  - operator 如何看到 dead-letter 和恢复动作
- reviewer 大概率会追问：
  - 为什么 retry 不等于可靠？
  - 哪个状态迁移最容易被漏测？
  - 如果 worker 异常退出，哪条 contract 先救场？

## 本月案例入口

- 先读 [专题案例包](casebook.md) 中的“案例 2：任务系统看起来能跑，线上却越跑越乱”，把本月知识点放回真实工程场景。
- 再读 [失败场景图谱](failure-patterns.md) 中的“模式二：失败语义缺失型故障”，理解这一月最容易反复出现的故障模式。
- 读完后回到本月的周/日文档，检查自己能不能说清楚：错误做法、坏信号和最小证据分别是什么。

## 本月实战清单

- [`RFC`](artifact-playbook.md#artifact-rfc)：写任务模型 / Job Runner RFC，冻结状态机、错误语义和 queue contract。
- [`Interface`](artifact-playbook.md#artifact-interface)：定义 enqueue、claim、ack、release、retry、dead-letter 的接口和 owner。
- [`Implementation`](artifact-playbook.md#artifact-implementation)：实现 worker 主循环、graceful shutdown 与重试/超时主路径。
- [`Unit Test`](artifact-playbook.md#artifact-unit-test)：覆盖状态迁移、重复执行、超时与幂等保护。
- [`Review`](artifact-playbook.md#artifact-review)：检查 failure handling 是否仍依赖隐式约定，而不是显式 contract。
- [`Benchmark（按需）`](artifact-playbook.md#artifact-benchmark)：本月不做性能 benchmark，但要做 failure drill 或至少记录 queue / retry 证据。
- [`Documentation`](artifact-playbook.md#artifact-documentation)：补 worker runbook、失败路径说明和 dead-letter 处理手册。
- [`Memory`](artifact-playbook.md#artifact-memory)：沉淀哪些失败必须重试、哪些失败应该直接终止。

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

1. 如果要向资深 reviewer 解释“后台任务与失败处理”，你会怎样说明这月的架构取舍？
2. 本月哪条主路径最容易失败？你会如何设计验证证据？
3. 本月引入的哪个接口、状态或契约最可能在未来变化？为什么？
4. 如果 Atlas 的调用量增长 10 倍，本月设计会先在哪个点上断裂？
5. 本月消除了哪些技术债，又留下了哪些新债？
6. 哪个周度产物最能证明你的工程判断？为什么？
7. 你会怎样把本月所有周主题串成一个完整系统故事，而不是零散任务？
8. 你会出示哪些测试、trace、日志或运行证据，证明本月成果已经可交付？

## 本月推荐配套阅读

- [Month 02](../months/month-02.md)
- [Month 02 Backlog](../atlas/backlog/month-02.md)
- [Month 02 Issue Set](../atlas/issues/month-02/README.md)
- [Week 05 Report](../reports/weekly/week-05-report.md)
- [第一学期日学习索引（前两个月）](day-index-semester-1-part-1.md)
