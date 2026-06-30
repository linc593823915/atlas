# 第九个月：治理、安全与审计

最后更新：2026-06-30

## 本月目标

本月你要处理 Agent 系统里最容易被忽略、但在企业里最贵的问题：

- Prompt Injection
- 越权 Tool 调用
- 数据泄漏
- 审计可追踪性

## 本月学完后你应该会什么

你应该能解释：

- 为什么治理不能后补
- 为什么策略层是系统设计的一部分
- 为什么 audit log 不是“多打一条日志”

## 本月学习重点

### 1. Threat Model 与 Risk Taxonomy
理解：
- 资产、角色、攻击路径和 blast radius。
- 哪些动作必须被强治理。

### 2. Policy Engine 与 Action Classification
理解：
- policy 输入输出、action class 和 enforcement point。
- 审计事件如何保留决策输入。

### 3. Injection、Leakage 与 Misuse Test
理解：
- 对抗样例怎么设计。
- 安全与可用性如何同时验证。

### 4. Governance 文档与审计叙事
理解：
- review packet 如何组织。
- incident reconstruction 依赖哪些证据。

## 本月 Cookbook 做法

### Recipe 1：写 threat model

### Recipe 2：做 policy engine

### Recipe 3：定义 audit event

### Recipe 4：补 adversarial tests

## 本月交付标准

你至少要能证明：

- 高风险动作可以被拦截
- 审计记录能说明发生了什么
- 至少一类注入/误用场景被系统化验证

## 本月必学术语

- `Threat Model`：描述资产、角色、攻击路径和控制点的风险模型。
- `Policy Engine`：根据上下文和规则决定动作是否允许的决策层。
- `Action Classification`：把系统动作映射成不同风险等级与控制策略的分类过程。
- `Audit Event`：能解释是谁、为何、在何时做了什么的正式留痕单元。
- `Prompt Injection`：外部不可信指令试图改变系统目标或绕过约束的攻击方式。

## 本月知识地图

- 治理不是在系统做完后加几个 deny rule，而是从 threat model 开始设计控制边界。
- policy engine 的价值在于把“感觉危险”翻译成可执行、可重放、可审计的决策。
- prompt injection、data leakage 和 misuse 需要系统化测试，而不是口头提醒。
- 真正的治理产物必须同时包含控制设计、验证证据和 operator 能读懂的说明。

## 本月常见误区

- 把安全问题留到上线前再补。
- policy 散落在各处 if/else。
- 只写治理文档，不做 adversarial 验证。

## 本月最小演示场景

- 背景：一个高风险请求进入系统后，要么被 policy 拦住，要么在 approval 和 audit 证据完备的前提下执行。
- 你至少要能演示：
  - threat model 如何定位高风险动作
  - policy engine 如何作出 allow/deny
  - audit event 如何解释决策原因
- reviewer 大概率会追问：
  - 为什么治理不能后补？
  - prompt injection 之外还有哪些攻击面？
  - 日志为什么不等于可审计？

## 本月案例入口

- 先读 [专题案例包](casebook.md) 中的“案例 9：治理规则很多，但真正出事时一条都兜不住”，把本月知识点放回真实工程场景。
- 再读 [失败场景图谱](failure-patterns.md) 中的“模式四：治理失守型故障”，理解这一月最容易反复出现的故障模式。
- 读完后回到本月的周/日文档，检查自己能不能说清楚：错误做法、坏信号和最小证据分别是什么。

## 本月实战清单

- [`RFC`](artifact-playbook.md#artifact-rfc)：写 governance RFC，冻结 threat model、risk taxonomy 和 control boundary。
- [`Interface`](artifact-playbook.md#artifact-interface)：定义 policy engine、action classification、audit event contract。
- [`Implementation`](artifact-playbook.md#artifact-implementation)：实现最小 policy path、approval / deny path 和 audit 留痕。
- [`Unit Test`](artifact-playbook.md#artifact-unit-test)：覆盖 injection、leakage、misuse 和高风险动作拒绝路径。
- [`Review`](artifact-playbook.md#artifact-review)：检查治理是否仍只存在于文档而不在执行路径中。
- [`Benchmark（按需）`](artifact-playbook.md#artifact-benchmark)：本月不做性能 benchmark，但要保留 adversarial test 结果和控制有效性证据。
- [`Documentation`](artifact-playbook.md#artifact-documentation)：补 governance docs、review packet 和 operator controls 说明。
- [`Memory`](artifact-playbook.md#artifact-memory)：沉淀哪些动作必须被强治理，为什么。

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

1. 如果要向资深 reviewer 解释“治理、安全与审计”，你会怎样说明这月的架构取舍？
2. 本月哪条主路径最容易失败？你会如何设计验证证据？
3. 本月引入的哪个接口、状态或契约最可能在未来变化？为什么？
4. 如果 Atlas 的调用量增长 10 倍，本月设计会先在哪个点上断裂？
5. 本月消除了哪些技术债，又留下了哪些新债？
6. 哪个周度产物最能证明你的工程判断？为什么？
7. 你会怎样把本月所有周主题串成一个完整系统故事，而不是零散任务？
8. 你会出示哪些测试、trace、日志或运行证据，证明本月成果已经可交付？

## 本月推荐配套阅读

- [Month 09](../months/month-09.md)
- [Month 09 Backlog](../atlas/backlog/month-09.md)
- [Month 09 Issue Set](../atlas/issues/month-09/README.md)
- [第三学期完整日学习索引](day-index-semester-3.md)
