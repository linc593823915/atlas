# 第三个月：第一个单 Agent 服务

最后更新：2026-06-30

## 本月目标

本月你第一次真正进入 Agent 层。

目标不是做一个花哨聊天机器人，而是做一个：

- 有工具调用
- 有结构化输出
- 有失败处理
- 有最小评估基础

的单 Agent 服务。

## 本月学完后你应该会什么

你应该能解释：

- Responses API 在系统里扮演什么角色
- 什么是 Tool Calling 的工程边界
- 为什么 Structured Output 对可靠性重要
- 为什么输出“差不多正确”不等于系统可用

## 本月学习重点

### 1. Agent 角色与主路径
理解：
- 模型负责什么，runtime 负责什么，tool 负责什么。
- 第一个 Agent RFC 应该冻结哪些 contract。

### 2. Tool Inventory 与调用边界
理解：
- 为什么不是所有 helper 都应该暴露成 tool。
- tool schema 和 capability 组织方式。

### 3. Structured Output 与结果校验
理解：
- schema、validator 和 repair / retry 语义。
- 输出“差不多对”为什么仍然不安全。

### 4. 多 Tool 路由与最小评估
理解：
- routing policy、fallback path 和 eval baseline。
- operator note 为什么从第一个 Agent 就要写。

## 本月 Cookbook 做法

### Recipe 1：定义第一个 Agent RFC

产物：

- agent role
- tool inventory
- output contract

### Recipe 2：打通第一个 Tool

产物：

- end-to-end tool call
- parsed output

### Recipe 3：扩成多 Tool

产物：

- 至少三个 tool
- routing path
- fallback path

### Recipe 4：补评估底座

产物：

- 小型 eval baseline
- operator note

## 本月交付标准

到本月结束时，你至少要能演示：

- Agent 如何选工具
- 输出如何被验证
- 工具失败如何回退
- 为什么这个流程是“可演进”的

## 本月必学术语

- `Agent Role`：模型在一条链路里承担的职责边界，而不是一段 prompt 文本。
- `Tool Inventory`：一个 Agent 允许调用的工具集合以及各自 contract。
- `Structured Output`：带 schema 约束的模型输出，用于进入确定性代码路径。
- `Validator`：负责检查模型输出是否满足 contract 的校验层。
- `Eval Baseline`：后续回归比较所依赖的第一版质量基线。

## 本月知识地图

- 本月是从普通后端服务进入 Agent 系统设计的第一步，关键不是“能聊天”，而是 contract 是否稳定。
- tool calling 的真正难点不在调用 API，而在 tool schema、错误语义和 routing 决策是否可解释。
- structured output 决定模型结果能否进入确定性代码路径。
- 第一个 Agent 的质量证明必须依赖小型 eval、文档和 operator evidence，而不是演示时刚好效果不错。

## 本月常见误区

- 把 prompt 当系统设计。
- 把模型自然语言当可靠结构化数据。
- 先做多 tool 炫技，后补 contract 和 eval。

## 本月最小演示场景

- 背景：用户发来一个任务请求，Agent 需要选工具、得到结构化结果，并在工具失败时给出可解释回退。
- 你至少要能演示：
  - Agent 如何选择 tool
  - 结构化输出如何被 validator 接住
  - 质量如何用最小 eval baseline 被证明
- reviewer 大概率会追问：
  - 为什么 prompt 不能等于 contract？
  - tool inventory 为什么不能无限加？
  - 坏结果出现时如何判断是模型错还是系统错？

## 本月案例入口

- 先读 [专题案例包](casebook.md) 中的“案例 3：做出了会聊天的 Agent，但系统完全不可依赖”，把本月知识点放回真实工程场景。
- 再读 [失败场景图谱](failure-patterns.md) 中的“模式三：质量证据失真型故障”，理解这一月最容易反复出现的故障模式。
- 读完后回到本月的周/日文档，检查自己能不能说清楚：错误做法、坏信号和最小证据分别是什么。

## 本月实战清单

- [`RFC`](artifact-playbook.md#artifact-rfc)：写第一个 Agent RFC，冻结 role、tool inventory 和 output contract。
- [`Interface`](artifact-playbook.md#artifact-interface)：定义 tool schema、structured output schema 和 validator 行为。
- [`Implementation`](artifact-playbook.md#artifact-implementation)：打通首条 tool path，再扩为多 tool routing 与 fallback。
- [`Unit Test`](artifact-playbook.md#artifact-unit-test)：覆盖 schema 校验、tool failure 和输出解析失败。
- [`Review`](artifact-playbook.md#artifact-review)：从 reviewer 视角检查系统是不是还在依赖“看起来差不多”的自然语言。
- [`Benchmark（按需）`](artifact-playbook.md#artifact-benchmark)：本月不做性能 benchmark，但必须建立可重复的小型 eval baseline。
- [`Documentation`](artifact-playbook.md#artifact-documentation)：补 operator note、运行说明和 eval 使用说明。
- [`Memory`](artifact-playbook.md#artifact-memory)：沉淀单 Agent 最容易失控的边界：tool contract、schema 和 fallback。

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

1. 如果要向资深 reviewer 解释“第一个单 Agent 服务”，你会怎样说明这月的架构取舍？
2. 本月哪条主路径最容易失败？你会如何设计验证证据？
3. 本月引入的哪个接口、状态或契约最可能在未来变化？为什么？
4. 如果 Atlas 的调用量增长 10 倍，本月设计会先在哪个点上断裂？
5. 本月消除了哪些技术债，又留下了哪些新债？
6. 哪个周度产物最能证明你的工程判断？为什么？
7. 你会怎样把本月所有周主题串成一个完整系统故事，而不是零散任务？
8. 你会出示哪些测试、trace、日志或运行证据，证明本月成果已经可交付？

## 本月推荐配套阅读

- [Month 03](../months/month-03.md)
- [Month 03 Backlog](../atlas/backlog/month-03.md)
- [Month 03 Issue Set](../atlas/issues/month-03/README.md)
- [Week 09 Report](../reports/weekly/week-09-report.md)
- [第一学期完整日学习索引](day-index-semester-1.md)
