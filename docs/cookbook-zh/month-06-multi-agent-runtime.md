# 第六个月：多 Agent 运行时与审批

最后更新：2026-06-30

## 本月目标

本月要从“一个 Agent 会调用工具”升级到：

“多个 Agent 如何协作、交接、审批、受控运行”。

## 本月学完后你应该会什么

你应该能解释：

- Agent role 为什么必须清楚
- handoff 为什么不能隐式发生
- approval 为什么不是 UI 需求，而是治理需求
- guardrail 为什么应内建到 runtime

## 本月学习重点

### 1. Agent Role Map 与 Runtime Ownership
理解：
- planner、executor、reviewer 等角色如何分工。
- 谁对最终结果负责。

### 2. Shared Context 与 Trace Propagation
理解：
- 哪些上下文应共享，哪些不应共享。
- trace / thread id 如何贯穿多 agent。

### 3. Handoff、Approval 与 Guardrail
理解：
- 什么条件下 handoff。
- 敏感动作如何审批和约束。

### 4. Orchestration Test 与学期复盘
理解：
- 委托链、fallback 和人类接管怎么验证。
- 第二学期到底获得了什么能力。

## 本月 Cookbook 做法

### Recipe 1：画出角色图

### Recipe 2：实现 shared context 与 trace path

### Recipe 3：加入 handoff 与 approval

### Recipe 4：补多 Agent 测试

## 本月交付标准

到本月结束时，你至少要能说明：

- 多 Agent 为什么比单 Agent 更难
- 系统如何知道什么时候 handoff
- 为什么高风险动作必须审批

## 本月必学术语

- `Agent Role Map`：多 Agent 体系中各角色的职责分工图。
- `Shared Context`：在多个 agent 间传播且仍需保持来源清晰的上下文。
- `Handoff`：把目标、上下文和责任从一个 agent 转移给另一个 agent。
- `Approval Path`：敏感动作在执行前必须经过的人工或策略审批路径。
- `Guardrail`：限制 agent 行为边界的运行时约束与拦截规则。

## 本月知识地图

- 多 Agent 运行时的核心不是“多起几个模型”，而是设计清楚 delegation、ownership 和 evidence flow。
- shared context 必须被筛选和标注来源，否则 specialization 会被上下文噪音淹没。
- approval 和 guardrail 不是附属功能，而是多 Agent 系统进入真实使用前的基本控制层。
- 第二学期的完成标准是接口和 runtime 边界开始稳定，而不是角色数量变多。

## 本月常见误区

- 把多 Agent 设计成失控委托。
- 所有 agent 共享一坨不可追溯的上下文。
- 审批和 guardrail 只存在于文档，不存在于 runtime。

## 本月最小演示场景

- 背景：planner 把任务 handoff 给 executor，敏感动作进入 approval path，整个链路留下 trace 和 audit。
- 你至少要能演示：
  - 多 Agent 如何分工
  - handoff 时哪些上下文会被传播
  - 高风险动作如何被审批或阻断
- reviewer 大概率会追问：
  - 多 Agent 为什么不是多写几个 prompt？
  - 哪个 agent 对最终结果负责？
  - approval 如果只写在文档里会怎样？

## 本月案例入口

- 先读 [专题案例包](casebook.md) 中的“案例 6：多 Agent 看起来更强，实际上没人知道谁负责”，把本月知识点放回真实工程场景。
- 再读 [失败场景图谱](failure-patterns.md) 中的“模式四：治理失守型故障”，理解这一月最容易反复出现的故障模式。
- 读完后回到本月的周/日文档，检查自己能不能说清楚：错误做法、坏信号和最小证据分别是什么。

## 本月实战清单

- [`RFC`](artifact-playbook.md#artifact-rfc)：写多 Agent runtime RFC，冻结 role map、ownership 和 orchestration boundary。
- [`Interface`](artifact-playbook.md#artifact-interface)：定义 handoff、shared context、approval 和 guardrail contract。
- [`Implementation`](artifact-playbook.md#artifact-implementation)：打通最小多 Agent 编排链和审批主路径。
- [`Unit Test`](artifact-playbook.md#artifact-unit-test)：覆盖 orchestration、fallback、context propagation 和 approval deny path。
- [`Review`](artifact-playbook.md#artifact-review)：检查多 Agent 是否仍依赖隐式人工理解，而不是 runtime contract。
- [`Benchmark（按需）`](artifact-playbook.md#artifact-benchmark)：本月不做性能 benchmark，但要保留 orchestration / approval 证据。
- [`Documentation`](artifact-playbook.md#artifact-documentation)：补 runtime docs、operator path 和高风险动作说明。
- [`Memory`](artifact-playbook.md#artifact-memory)：沉淀 role map、handoff 和 approval 最容易失控的边界。

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

1. 如果要向资深 reviewer 解释“多 Agent 运行时与审批”，你会怎样说明这月的架构取舍？
2. 本月哪条主路径最容易失败？你会如何设计验证证据？
3. 本月引入的哪个接口、状态或契约最可能在未来变化？为什么？
4. 如果 Atlas 的调用量增长 10 倍，本月设计会先在哪个点上断裂？
5. 本月消除了哪些技术债，又留下了哪些新债？
6. 哪个周度产物最能证明你的工程判断？为什么？
7. 你会怎样把本月所有周主题串成一个完整系统故事，而不是零散任务？
8. 你会出示哪些测试、trace、日志或运行证据，证明本月成果已经可交付？

## 本月推荐配套阅读

- [Month 06](../months/month-06.md)
- [Month 06 Backlog](../atlas/backlog/month-06.md)
- [Month 06 Issue Set](../atlas/issues/month-06/README.md)
- [第二学期完整日学习索引](day-index-semester-2.md)
