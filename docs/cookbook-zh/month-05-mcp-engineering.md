# 第五个月：MCP 工程化

最后更新：2026-06-30

## 本月目标

本月的目标是理解并实现 MCP 工程化能力。

你不只是“知道 MCP 是什么”，而要能在 Atlas 里把能力暴露成可协作协议面。

## 本月学完后你应该会什么

你应该能解释：

- MCP 为什么重要
- tools/resources/prompts 各自适合什么职责
- 为什么协议兼容性和 versioning 是工程问题

## 本月学习重点

### 1. MCP Surface 选择
理解：
- 哪些能力值得对外暴露。
- tools 和 resources 的边界。

### 2. Server Bootstrap 与 Transport
理解：
- session lifecycle 和 dependency lifecycle 如何对齐。
- stdio / socket / HTTP wrapper 的差异。

### 3. Capability 设计与兼容性
理解：
- 兼容性、discoverability 和协议演进。
- deprecated capability 如何处理。

### 4. Protocol Test 与文档产品化
理解：
- 为什么协议文档本身就是产品。
- 如何证明旧 client 仍能工作。

## 本月 Cookbook 做法

### Recipe 1：定义 MCP RFC

### Recipe 2：实现最小 MCP Server

### Recipe 3：暴露至少两类 capability

### Recipe 4：补协议测试与兼容说明

## 本月交付标准

到本月结束时，你至少要能演示：

- MCP 服务怎么启动
- 一个 capability 怎样被消费
- 兼容性问题怎样被提前考虑

## 本月必学术语

- `MCP Surface`：Atlas 向外部 client 暴露的稳定协议能力集合。
- `Capability Inventory`：当前 server 提供的 tools/resources 及其元信息清单。
- `Tool / Resource`：tool 偏命令式动作，resource 偏只读状态或上下文。
- `Transport`：承载 MCP 会话的通信方式，如 stdio、socket 等。
- `Compatibility`：服务端演进时仍让既有 client 可预测工作的约束。

## 本月知识地图

- MCP 不是把内部功能“开放一下”，而是定义 Atlas 的第一个外部协议面。
- surface scope、session lifecycle 和 error shape 会长期影响后续 client 的使用成本。
- 兼容性是 MCP 工程化的中心问题，任何对外能力都要有 discoverability 和演进策略。
- 协议级测试和文档是 MCP 是否真正可消费的核心证据。

## 本月常见误区

- 把内部 helper 暴露成长期协议 contract。
- 没有 capability discoverability 就开始谈兼容。
- 只有手工调试，没有 protocol-level tests。

## 本月最小演示场景

- 背景：一个外部 client 通过 MCP 发现 capability、调用 capability，并在不支持能力时拿到可解释反馈。
- 你至少要能演示：
  - MCP 服务如何启动
  - capability inventory 如何被发现
  - 兼容性和 unsupported path 如何表现
- reviewer 大概率会追问：
  - 哪些能力真的值得暴露成 MCP surface？
  - tools 和 resources 的边界为什么重要？
  - 如何避免把内部 helper 锁死成外部 contract？

## 本月案例入口

- 先读 [专题案例包](casebook.md) 中的“案例 5：MCP 服务能启动，但根本不是可消费协议”，把本月知识点放回真实工程场景。
- 再读 [失败场景图谱](failure-patterns.md) 中的“模式一：边界缺失型故障”，理解这一月最容易反复出现的故障模式。
- 读完后回到本月的周/日文档，检查自己能不能说清楚：错误做法、坏信号和最小证据分别是什么。

## 本月实战清单

- [`RFC`](artifact-playbook.md#artifact-rfc)：写 MCP RFC，冻结 surface scope、session lifecycle 和 error shape。
- [`Interface`](artifact-playbook.md#artifact-interface)：定义 capability inventory、transport 行为和 compatibility 规则。
- [`Implementation`](artifact-playbook.md#artifact-implementation)：实现最小 MCP server 与至少一类 capability。
- [`Unit Test`](artifact-playbook.md#artifact-unit-test)：覆盖 discovery、invocation、unsupported capability 和 error shape。
- [`Review`](artifact-playbook.md#artifact-review)：检查暴露给 client 的 surface 是否已经脱离内部实现细节。
- [`Benchmark（按需）`](artifact-playbook.md#artifact-benchmark)：本月不要求性能 benchmark，但至少做协议层回归验证。
- [`Documentation`](artifact-playbook.md#artifact-documentation)：补 capability 文档、兼容说明和 client 使用说明。
- [`Memory`](artifact-playbook.md#artifact-memory)：沉淀 MCP 工程化里最难变更、最该保守的 contract。

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

1. 如果要向资深 reviewer 解释“MCP 工程化”，你会怎样说明这月的架构取舍？
2. 本月哪条主路径最容易失败？你会如何设计验证证据？
3. 本月引入的哪个接口、状态或契约最可能在未来变化？为什么？
4. 如果 Atlas 的调用量增长 10 倍，本月设计会先在哪个点上断裂？
5. 本月消除了哪些技术债，又留下了哪些新债？
6. 哪个周度产物最能证明你的工程判断？为什么？
7. 你会怎样把本月所有周主题串成一个完整系统故事，而不是零散任务？
8. 你会出示哪些测试、trace、日志或运行证据，证明本月成果已经可交付？

## 本月推荐配套阅读

- [Month 05](../months/month-05.md)
- [Month 05 Backlog](../atlas/backlog/month-05.md)
- [Month 05 Issue Set](../atlas/issues/month-05/README.md)
- [第二学期完整日学习索引](day-index-semester-2.md)
