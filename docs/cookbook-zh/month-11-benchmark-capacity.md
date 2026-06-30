# 第十一个月：性能、成本与容量规划

最后更新：2026-06-30

## 本月目标

本月你要开始回答：

- 这个系统有多快？
- 有多贵？
- 在多大规模下会出问题？

## 本月学完后你应该会什么

你应该能解释：

- 为什么 benchmark 不等于压测截图
- 为什么 token cost 也是架构约束
- 为什么 capacity planning 是平台能力的一部分

## 本月学习重点

### 1. Baseline Measurement
理解：
- cold / warm、latency / throughput / token cost 如何分开测。
- 什么 workload 才有代表性。

### 2. Optimization Loop
理解：
- 如何提出瓶颈假设。
- before / after 比较如何保持可比。

### 3. Cost 与 Capacity Planning
理解：
- token cost 和 infra cost 怎么进入架构判断。
- 容量假设怎么写。

### 4. Graduation Platform RFC
理解：
- 为什么性能月末会过渡到平台边界。
- 子系统组合能力如何被重新审视。

## 本月 Cookbook 做法

### Recipe 1：定义 benchmark contract

### Recipe 2：拿到 baseline

### Recipe 3：做一次优化循环

### Recipe 4：写 capacity assumption

## 本月交付标准

你至少要能拿出：

- baseline 数据
- before/after 优化结果
- 成本/容量判断依据

## 本月必学术语

- `Baseline`：后续所有优化对比都依赖的第一版可信测量结果。
- `P95 Latency`：95% 请求不应超过的时延位置，用于看尾延迟。
- `Throughput`：单位时间内系统能够完成的请求或任务量。
- `Token Cost`：模型调用因 token 消耗而带来的直接成本约束。
- `Capacity Assumption`：关于并发、负载和饱和点的容量前提假设。

## 本月知识地图

- 性能工程的第一步是拿到可信基线，而不是直接开始优化。
- 优化必须围绕主瓶颈和可比较 workload 进行，否则数字变化没有意义。
- 成本和容量并不是运营尾活，它们会直接限制架构和产品形态。
- 第十一个月后半段开始进入平台化思考：性能结论必须能回到平台边界与集成顺序。

## 本月常见误区

- 基线本身不稳定就开始优化。
- 只看平均值，不看 workload 组成和方差。
- 把 capacity planning 写成拍脑袋估算。

## 本月最小演示场景

- 背景：同一份 benchmark contract 下，先拿 baseline，再做一次优化，最后给出容量和成本判断。
- 你至少要能演示：
  - baseline 如何定义
  - before / after 为什么可比较
  - 容量假设和 token cost 怎样进入架构判断
- reviewer 大概率会追问：
  - 为什么 benchmark 不等于一张压测截图？
  - 哪个瓶颈最值得先优化？
  - 如果 workload 变了，这组数字还可信吗？

## 本月案例入口

- 先读 [专题案例包](casebook.md) 中的“案例 11：你拿到一堆性能数字，但一个能信的结论都没有”，把本月知识点放回真实工程场景。
- 再读 [失败场景图谱](failure-patterns.md) 中的“模式三：质量证据失真型故障”，理解这一月最容易反复出现的故障模式。
- 读完后回到本月的周/日文档，检查自己能不能说清楚：错误做法、坏信号和最小证据分别是什么。

## 本月实战清单

- [`RFC`](artifact-playbook.md#artifact-rfc)：写 benchmark contract 或 measurement plan，冻结 workload、环境和指标口径。
- [`Interface`](artifact-playbook.md#artifact-interface)：定义 latency / throughput / token cost 统计口径和 capacity assumption 模板。
- [`Implementation`](artifact-playbook.md#artifact-implementation)：完成 baseline 测量并执行一轮单变量优化。
- [`Unit Test`](artifact-playbook.md#artifact-unit-test)：验证统计脚本、结果解析或阈值计算逻辑。
- [`Review`](artifact-playbook.md#artifact-review)：检查性能结论是否真的建立在可比 workload 和证据之上。
- [`Benchmark（按需）`](artifact-playbook.md#artifact-benchmark)：本月 benchmark 是必做项：保留 baseline、after 和方差说明。
- [`Documentation`](artifact-playbook.md#artifact-documentation)：补 benchmark note、capacity plan 和 trade-off summary。
- [`Memory`](artifact-playbook.md#artifact-memory)：沉淀最值得复用的测量方法和最容易误导的数字。

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

1. 如果要向资深 reviewer 解释“性能、成本与容量规划”，你会怎样说明这月的架构取舍？
2. 本月哪条主路径最容易失败？你会如何设计验证证据？
3. 本月引入的哪个接口、状态或契约最可能在未来变化？为什么？
4. 如果 Atlas 的调用量增长 10 倍，本月设计会先在哪个点上断裂？
5. 本月消除了哪些技术债，又留下了哪些新债？
6. 哪个周度产物最能证明你的工程判断？为什么？
7. 你会怎样把本月所有周主题串成一个完整系统故事，而不是零散任务？
8. 你会出示哪些测试、trace、日志或运行证据，证明本月成果已经可交付？

## 本月推荐配套阅读

- [Month 11](../months/month-11.md)
- [Month 11 Backlog](../atlas/backlog/month-11.md)
- [Month 11 Issue Set](../atlas/issues/month-11/README.md)
- [第四学期完整日学习索引](day-index-semester-4.md)
