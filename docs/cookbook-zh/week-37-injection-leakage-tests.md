# Week 37：注入、泄漏与误用测试

最后更新：2026-06-30

## 本周目标

这一周不再停留在“我觉得安全”，而是开始用测试逼出系统真实问题。

## 本周学完后你应该会什么

重点是 adversarial thinking。

## 本周学习重点

- prompt injection
- data leakage
- misuse tests

## 本周 Cookbook 做法

### Recipe 1：设计注入测试

### Recipe 2：设计泄漏测试

### Recipe 3：设计误用测试

## 本周交付标准

- 一组 adversarial case
- 风险说明
- 测试结果

## 本周能力焦点

- 安全判断（`Security Judgment`）
- 策略控制（`Policy Controls`）
- 可审计性（`Auditability`）

## 本周知识清单

- prompt injection 防御首先要识别不可信指令从哪里进入系统。
- 数据泄漏路径不仅在模型输出，也在 tool result、trace、log、memory 和人工 review 面。
- misuse tests 要尽量贴近真实攻击者会做的事，而不是只写几个玩具字符串。
- 好的防御证据既要证明恶意流量被拦住，也要证明正常流量没有被误伤。

## 本周 Atlas 观察路径

- `[Week 37 英文原文](../weeks/week-37.md)`：对照原始英文任务拆分，确认这一周的 deliverables、source reading 和 daily links。
- `[Atlas 主线 Issue（Month 09）](../atlas/issues/month-09/atlas-m09-03.md)`：把本周主题放回 Atlas 主线问题卡，观察范围、验收项和剩余风险。
- `[Week 37 周报](../reports/weekly/week-37-report.md)`：看本周结果、证据和后续动作是如何被记录的。
- `[docs/governance/sample-evidence-entries.md](../governance/sample-evidence-entries.md)`：把 adversarial test 结果放回正式证据格式里看。

## 本周 Atlas 落地检查

- 最小落地动作：围绕“设计注入测试”推进，让至少一项本周交付标准在仓库里可见。
- 必留证据：injection tests、leakage checks 和 misuse results。
- 最大风险：以为 prompt hygiene 已经等于治理，忽略了 tool、trace 和数据面的攻击面。
- 审查不变量：恶意场景被阻断的同时，等价 benign 场景也必须能被系统区分和放行。
- 本周最先要盯住的交付：一组 adversarial case

## 本周最小演示场景

- 背景：这一周你需要把“注入、泄漏与误用测试”从概念推进成可被别人看懂、验证和质疑的工程结果。
- 你至少要能演示：
  - 一组 adversarial case、风险说明、测试结果
  - “prompt injection”为什么会直接影响 Atlas 的主路径，而不只是局部实现
  - 风险“以为 prompt hygiene 已经等于治理，忽略了 tool、trace 和数据面的攻击面。”是如何被控制、暴露或留证的
- reviewer 大概率会追问：
  - 为什么这周先做“设计注入测试”，而不是先做别的？
  - 不变量“恶意场景被阻断的同时，等价 benign 场景也必须能被系统区分和放行。”如果被破坏，会先造成什么后果？
  - 你打算拿什么证据证明这周成果成立？提示：injection tests、leakage checks 和 misuse results。

## 配套案例

- 本周最接近的案例：读 [专题案例包](casebook.md) 中的“案例 9：治理规则很多，但真正出事时一条都兜不住”。
- 本周最该对照的失败模式：读 [失败场景图谱](failure-patterns.md) 中的“模式四：治理失守型故障”。
- 推荐顺序：先读本周知识清单，再读案例，再回本周 Atlas 观察路径和周报，检查自己能不能解释坏信号与最小证据。

## 本周工作顺序

- [`RFC`](artifact-playbook.md#artifact-rfc)：先把“设计注入测试”对应的问题范围、目标、约束和非目标写清楚。
- [`Interface`](artifact-playbook.md#artifact-interface)：围绕“prompt injection”定义输入、输出、owner、错误语义和边界。
- [`Implementation`](artifact-playbook.md#artifact-implementation)：把最小主路径打通，至少让“一组 adversarial case、风险说明、测试结果”里的一项开始可验证。
- [`Unit Test`](artifact-playbook.md#artifact-unit-test)：优先覆盖最容易击中风险“以为 prompt hygiene 已经等于治理，忽略了 tool、trace 和数据面的攻击面。”的失败路径、边界条件或 contract。
- [`Review`](artifact-playbook.md#artifact-review)：站在 reviewer 视角检查不变量“恶意场景被阻断的同时，等价 benign 场景也必须能被系统区分和放行。”是否真的被系统保护，而不是只写在文档里。
- [`Benchmark（按需）`](artifact-playbook.md#artifact-benchmark)：本周不做性能 benchmark，但要保留 adversarial test 结果和误伤说明。
- [`Documentation`](artifact-playbook.md#artifact-documentation)：补运行说明、review note 或 operator 指引，让别人能复用本周结论。
- [`Memory`](artifact-playbook.md#artifact-memory)：沉淀一条可复用结论，说明为什么本周的边界、证据或失败判断值得保留。

## 本周评分标准

- 满分 `100` 分。
- 交付物完成度：25 分。
- 技术清晰度与取舍解释：20 分。
- 代码健康、测试或验证深度：15 分。
- 可靠性与边界条件思考：15 分。
- 文档与评审质量：10 分。
- 复盘、memory 和下一步质量：15 分。
- 及格线：`75+`，且本周计划产物全部存在。
- 优秀线：`90+`，并能证明本周工作真实改善了 Atlas，而不只是完成表面任务。

## 本周面试追问

1. 如果要把“注入、泄漏与误用测试”做成 Atlas 可交付能力，你会如何设计主路径和边界？
2. 本周最值得优先验证的失败模式是什么？你会如何收集证据？
3. 这周的 RFC、实现、验证或复盘里，哪一项最能体现你的工程判断？为什么？
4. 如果时间被砍掉一半，你会保留什么，推迟什么？
5. 本周产物如何自然承接到下周，而不是变成一次性文档或一次性代码？

## 每日入口

- [Day 253](days-semester-3/day-253.md): 理解问题和边界
- [Day 254](days-semester-3/day-254.md): 边界与契约设计
- [Day 255](days-semester-3/day-255.md): 最小主路径实现
- [Day 256](days-semester-3/day-256.md): 补全真实可用路径
- [Day 257](days-semester-3/day-257.md): 验证与实验
- [Day 258](days-semester-3/day-258.md): 文档与评审
- [Day 259](days-semester-3/day-259.md): 复盘与 memory

## 配套阅读

- [Week 37](../weeks/week-37.md)
