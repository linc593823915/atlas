# 第十二个月：平台整合与毕业答辩

最后更新：2026-06-30

## 本月目标

最后一个月不是简单“收尾”。

而是要回答：

Atlas 到底是不是一个平台雏形？

## 本月学完后你应该会什么

你应该能：

- 把 config、tool gateway、MCP、workflow、eval、approval、audit、deployment 串成一个系统故事
- 用架构语言解释取舍
- 用证据支撑平台级判断

## 本月学习重点

### 1. 跨子系统合同整合
理解：
- config、tool gateway、MCP、workflow、eval、approval、audit 如何对齐。
- 哪里应该共享 contract。

### 2. 端到端证明与失败演练
理解：
- 真正的平台级用户旅程有哪些。
- 哪些故障必须被 drill。

### 3. 架构文档、运维文档与 Benchmark 包
理解：
- 边界图、operator guide、review packet 怎么组织。
- 哪些证据最适合答辩。

### 4. 毕业答辩与下一阶段路线
理解：
- 如何解释 trade-off、限制和未来 6-12 个月下注。
- 怎样证明 Atlas 已经像平台，而不只是项目。

## 本月 Cookbook 做法

### Recipe 1：写 platform RFC

### Recipe 2：整合核心子系统

### Recipe 3：做 end-to-end drill

### Recipe 4：准备毕业答辩材料

## 本月交付标准

你至少要能完成：

- 平台级演示
- 架构文档
- 关键运行文档
- 基于证据的答辩

## 最后要回答的问题

1. Atlas 现在是不是一个“平台”，而不只是“项目”？
2. 我能不能清楚解释每个核心边界？
3. 我有没有足够证据说明这个系统具备可运营性？

## 本月必学术语

- `Platform Boundary`：把零散子系统组织成一套平台时必须冻结的边界。
- `End-to-end Drill`：跨多个子系统执行完整用户旅程和故障演练的验证方式。
- `Evidence Pack`：支持评审或答辩的一组结构化证据材料。
- `Operator Guide`：教接手者和运维者如何运行、排障和回退系统的说明。
- `Graduation Defense`：围绕系统设计、证据和未来路线进行的毕业级答辩。

## 本月知识地图

- 最后一个月的任务不是再做新功能，而是证明前 11 个月形成的是一套平台级系统。
- 平台整合的关键是让各子系统共享一套 contract、evidence 和 operator language。
- 端到端 drill、架构文档、operator 文档和最终 benchmark 共同组成毕业答辩的证据包。
- 真正成熟的毕业材料会同时讲清楚系统能做什么、为什么可信、还有哪些限制以及下一步往哪走。

## 本月常见误区

- 把收尾月做成功能大杂烩。
- 只堆文档，不做端到端证明。
- 只讲成绩，不讲限制和下一步下注。

## 本月最小演示场景

- 背景：Atlas 以平台视角完整演示一条跨子系统旅程，并能用文档、drill、benchmark 和 review packet 完成毕业答辩。
- 你至少要能演示：
  - config、tool gateway、MCP、workflow、eval、approval、audit 如何协同
  - operator 和 reviewer 分别看什么证据
  - 下一阶段路线如何基于现有证据提出
- reviewer 大概率会追问：
  - 为什么功能全都有仍然不等于平台成立？
  - 哪些边界是平台级边界，而不是模块级边界？
  - 如果只选三组证据做答辩，你会选什么？

## 本月案例入口

- 先读 [专题案例包](casebook.md) 中的“案例 12：功能都在，为什么还是说不清这是个平台”，把本月知识点放回真实工程场景。
- 再读 [失败场景图谱](failure-patterns.md) 中的“模式五：平台错觉型故障”，理解这一月最容易反复出现的故障模式。
- 读完后回到本月的周/日文档，检查自己能不能说清楚：错误做法、坏信号和最小证据分别是什么。

## 本月实战清单

- [`RFC`](artifact-playbook.md#artifact-rfc)：写平台整合 RFC，冻结平台边界、共享 contract 和整合顺序。
- [`Interface`](artifact-playbook.md#artifact-interface)：定义跨子系统共享的 capability / evidence / operator contract。
- [`Implementation`](artifact-playbook.md#artifact-implementation)：打通 end-to-end 集成主路径与至少一条故障演练路径。
- [`Unit Test`](artifact-playbook.md#artifact-unit-test)：覆盖关键整合点的一致性验证和端到端 smoke path。
- [`Review`](artifact-playbook.md#artifact-review)：从 reviewer 视角检查平台故事是否完整、证据是否足够。
- [`Benchmark（按需）`](artifact-playbook.md#artifact-benchmark)：本月 benchmark 体现在最终 benchmark 包和平台级 health summary 上。
- [`Documentation`](artifact-playbook.md#artifact-documentation)：完成 architecture pack、operator guide、graduation defense 材料和 next roadmap。
- [`Memory`](artifact-playbook.md#artifact-memory)：沉淀全年最稳定的架构判断、证据组织方式和下一步下注。

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

1. 如果要向资深 reviewer 解释“平台整合与毕业答辩”，你会怎样说明这月的架构取舍？
2. 本月哪条主路径最容易失败？你会如何设计验证证据？
3. 本月引入的哪个接口、状态或契约最可能在未来变化？为什么？
4. 如果 Atlas 的调用量增长 10 倍，本月设计会先在哪个点上断裂？
5. 本月消除了哪些技术债，又留下了哪些新债？
6. 哪个周度产物最能证明你的工程判断？为什么？
7. 你会怎样把本月所有周主题串成一个完整系统故事，而不是零散任务？
8. 你会出示哪些测试、trace、日志或运行证据，证明本月成果已经可交付？

## 本月推荐配套阅读

- [Month 12](../months/month-12.md)
- [Month 12 Backlog](../atlas/backlog/month-12.md)
- [Month 12 Issue Set](../atlas/issues/month-12/README.md)
- [第四学期完整日学习索引](day-index-semester-4.md)
