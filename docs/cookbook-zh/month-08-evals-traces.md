# 第八个月：评估、Tracing 与发布门禁

最后更新：2026-06-30

## 本月目标

本月要让 Atlas 从“能运行”升级到“可被度量”。

核心问题是：

- 你怎么知道 Agent 系统变好了？
- 你怎么知道它退化了？
- 你怎么知道错误发生在哪？

## 本月学完后你应该会什么

你应该能解释：

- eval dataset 为什么重要
- grader 为什么重要
- trace 为什么是调试和治理基础设施
- release gate 为什么必须提前定义

## 本月学习重点

### 1. Eval Loop 与 Grader Strategy
理解：
- evaluation unit、grader 类型和评分语义。
- 谁拥有 score interpretation。

### 2. Tracing 与 Dataset Model
理解：
- trace 要记录什么事件。
- dataset 如何 version 和 rerun。

### 3. Metrics、Gate 与发布规则
理解：
- 哪些指标进入 release decision。
- threshold 和 escalation 怎么定。

### 4. Eval Ops 与复盘文档
理解：
- 怎么运行评估、怎么读 trace、怎么更新 dataset。
- measurement 如何改变工程决策。

## 本月 Cookbook 做法

### Recipe 1：定义 eval loop RFC

### Recipe 2：建立 trace pipeline

### Recipe 3：准备 benchmark dataset

### Recipe 4：加入 regression gate

## 本月交付标准

你至少要能给出：

- 一组 eval 样本
- 一组可解释的 trace
- 一条 release gate 规则

## 本月必学术语

- `Eval Loop`：从数据集、grader 到结果解释的完整评估循环。
- `Grader`：负责给一次运行结果打分或分类的评估器。
- `Trace`：记录模型、工具、策略和状态变化过程的可观测链路。
- `Regression Gate`：根据指标阈值决定是否阻断变更的质量门禁。
- `Benchmark Dataset`：用于稳定对比系统表现的一组代表性样本。

## 本月知识地图

- 本月的目标是把“系统变好了没有”从感觉问题变成可度量问题。
- grader、trace 和 dataset 共同构成质量证据链，缺任何一环都无法稳定回归。
- release gate 是工程决策机制，不是多一层 dashboard。
- 好的 eval 体系不仅能给分，还能解释为什么出错、错在模型还是错在系统。

## 本月常见误区

- 拿零散样例冒充 eval loop。
- trace 很多但没有决策价值。
- 门禁很多但没人按它做发布决定。

## 本月最小演示场景

- 背景：一组数据集样本跑完整个 eval loop，trace 记录决策路径，gate 根据指标决定是否允许发布。
- 你至少要能演示：
  - grader 如何打分
  - trace 如何解释一次失败
  - release gate 如何阻断回归
- reviewer 大概率会追问：
  - 为什么 trace 不是“更好看的日志”？
  - score 为什么必须能回溯到数据和 trace？
  - gate fail 后应该怎么决策？

## 本月案例入口

- 先读 [专题案例包](casebook.md) 中的“案例 8：有了 trace 和评估，但没人知道这些数据该怎么用”，把本月知识点放回真实工程场景。
- 再读 [失败场景图谱](failure-patterns.md) 中的“模式三：质量证据失真型故障”，理解这一月最容易反复出现的故障模式。
- 读完后回到本月的周/日文档，检查自己能不能说清楚：错误做法、坏信号和最小证据分别是什么。

## 本月实战清单

- [`RFC`](artifact-playbook.md#artifact-rfc)：写 eval loop RFC，冻结 evaluation unit、grader strategy 和 gate 语义。
- [`Interface`](artifact-playbook.md#artifact-interface)：定义 dataset row、trace event、metric threshold 和 release policy contract。
- [`Implementation`](artifact-playbook.md#artifact-implementation)：接入 grader、trace pipeline 和 regression gate。
- [`Unit Test`](artifact-playbook.md#artifact-unit-test)：覆盖评分逻辑、trace schema、gate fail path 和 dataset 版本一致性。
- [`Review`](artifact-playbook.md#artifact-review)：检查当前质量结论是否仍依赖 spot check 或个人感觉。
- [`Benchmark（按需）`](artifact-playbook.md#artifact-benchmark)：本月的 benchmark 体现在数据集、trace 和指标对比上，至少保留一轮可重跑结果。
- [`Documentation`](artifact-playbook.md#artifact-documentation)：补 eval ops guide、trace interpretation note 和 release gate 说明。
- [`Memory`](artifact-playbook.md#artifact-memory)：沉淀什么才算可信的质量证据。

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

1. 如果要向资深 reviewer 解释“评估、Tracing 与发布门禁”，你会怎样说明这月的架构取舍？
2. 本月哪条主路径最容易失败？你会如何设计验证证据？
3. 本月引入的哪个接口、状态或契约最可能在未来变化？为什么？
4. 如果 Atlas 的调用量增长 10 倍，本月设计会先在哪个点上断裂？
5. 本月消除了哪些技术债，又留下了哪些新债？
6. 哪个周度产物最能证明你的工程判断？为什么？
7. 你会怎样把本月所有周主题串成一个完整系统故事，而不是零散任务？
8. 你会出示哪些测试、trace、日志或运行证据，证明本月成果已经可交付？

## 本月推荐配套阅读

- [Month 08](../months/month-08.md)
- [Month 08 Backlog](../atlas/backlog/month-08.md)
- [Month 08 Issue Set](../atlas/issues/month-08/README.md)
- [第三学期完整日学习索引](day-index-semester-3.md)
