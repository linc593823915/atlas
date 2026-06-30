# 第七个月：持久化工作流与记忆边界

最后更新：2026-06-30

## 本月目标

本月的目标是让系统具备“长流程”能力。

也就是说，Agent 不再只做一次请求，而是能：

- checkpoint
- pause
- resume
- replay
- human intervention

## 本月学完后你应该会什么

你应该能解释：

- 为什么长流程一定需要状态
- 为什么 pause/resume 不是“附加能力”
- 为什么 memory boundary 要先想清楚

## 本月学习重点

### 1. Workflow State Model
理解：
- run、step、checkpoint 和 terminal state 怎么命名。
- 什么信息属于 workflow state。

### 2. Checkpoint 与 Resume Model
理解：
- 状态图、checkpoint store 和 resume cursor。
- 部分写入和部分副作用如何恢复。

### 3. Pause、Replay 与 Human Loop
理解：
- pause 是什么语义。
- 人工介入如何并回工作流。

### 4. Stateful Test 与运维说明
理解：
- 重启、重复 resume、坏 checkpoint 怎么测。
- operator 如何决定 resume / replay / abort。

## 本月 Cookbook 做法

### Recipe 1：定义状态图

### Recipe 2：引入 checkpoint

### Recipe 3：实现 pause / resume

### Recipe 4：加入 replay 与人工介入

## 本月交付标准

你至少要能演示：

- 一个流程如何中断后恢复
- 状态保存在什么地方
- 什么信息属于 workflow state，什么信息不属于

## 本月必学术语

- `Workflow State`：长流程运行时必须被显式保存和恢复的状态集合。
- `Checkpoint`：在流程关键节点持久化的可恢复快照。
- `Resume Cursor`：告诉 runtime 应从哪个状态节点继续执行的定位信息。
- `Replay`：基于历史状态和证据重新演示或重放流程的机制。
- `Human Intervention`：人工在流程中插入决策、修正或批准的受控接入点。

## 本月知识地图

- 长流程能力的核心是状态边界和恢复语义，而不是多写几个步骤。
- checkpoint 是否可信，决定了 pause / resume / replay 是否只是概念词。
- human intervention 需要 schema、audit 和 merge 语义，否则人工操作会破坏自动化诚实性。
- stateful system 必须用故障恢复测试和 operator 文档证明自己真的可运行。

## 本月常见误区

- 把长流程当普通请求循环。
- 保存 snapshot 却无法推出下一步。
- 人工修改状态但不记录 provenance。

## 本月最小演示场景

- 背景：一个长流程中途暂停，进程重启后从 checkpoint 恢复，必要时人工介入后继续完成。
- 你至少要能演示：
  - workflow state 保存了什么
  - resume 从哪里继续
  - 人工介入后如何保持状态诚实
- reviewer 大概率会追问：
  - checkpoint 和 memory boundary 为什么不是一回事？
  - 为什么 resume 必须是确定性的？
  - 哪些副作用绝不能在 replay 中被无脑重放？

## 本月案例入口

- 先读 [专题案例包](casebook.md) 中的“案例 7：长流程跑得很长，但根本恢复不了”，把本月知识点放回真实工程场景。
- 再读 [失败场景图谱](failure-patterns.md) 中的“模式二：失败语义缺失型故障”，理解这一月最容易反复出现的故障模式。
- 读完后回到本月的周/日文档，检查自己能不能说清楚：错误做法、坏信号和最小证据分别是什么。

## 本月实战清单

- [`RFC`](artifact-playbook.md#artifact-rfc)：写 durable workflow RFC，冻结 state model、checkpoint boundary 和 resume 语义。
- [`Interface`](artifact-playbook.md#artifact-interface)：定义 checkpoint store、resume cursor、human intervention contract。
- [`Implementation`](artifact-playbook.md#artifact-implementation)：实现 pause / resume / replay 的最小闭环。
- [`Unit Test`](artifact-playbook.md#artifact-unit-test)：覆盖重复 resume、checkpoint 损坏、人工修改状态等场景。
- [`Review`](artifact-playbook.md#artifact-review)：检查长流程是否仍依赖作者记忆，而不是状态和证据。
- [`Benchmark（按需）`](artifact-playbook.md#artifact-benchmark)：本月不做性能 benchmark，但要做恢复 drill 和状态一致性验证。
- [`Documentation`](artifact-playbook.md#artifact-documentation)：补 operator 恢复手册和 workflow 状态说明。
- [`Memory`](artifact-playbook.md#artifact-memory)：沉淀哪些事实属于 workflow state，哪些属于外部长期记忆。

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

1. 如果要向资深 reviewer 解释“持久化工作流与记忆边界”，你会怎样说明这月的架构取舍？
2. 本月哪条主路径最容易失败？你会如何设计验证证据？
3. 本月引入的哪个接口、状态或契约最可能在未来变化？为什么？
4. 如果 Atlas 的调用量增长 10 倍，本月设计会先在哪个点上断裂？
5. 本月消除了哪些技术债，又留下了哪些新债？
6. 哪个周度产物最能证明你的工程判断？为什么？
7. 你会怎样把本月所有周主题串成一个完整系统故事，而不是零散任务？
8. 你会出示哪些测试、trace、日志或运行证据，证明本月成果已经可交付？

## 本月推荐配套阅读

- [Month 07](../months/month-07.md)
- [Month 07 Backlog](../atlas/backlog/month-07.md)
- [Month 07 Issue Set](../atlas/issues/month-07/README.md)
- [第三学期日学习索引（第一个月）](day-index-semester-3-part-1.md)
