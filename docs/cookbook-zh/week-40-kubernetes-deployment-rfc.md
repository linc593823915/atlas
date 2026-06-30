# Week 40：Kubernetes 部署 RFC

最后更新：2026-06-30

## 本周目标

这一周正式进入生产化部署视角。重点不是先配 YAML，而是先想清楚部署架构、资源边界和运行模型。

## 本周学完后你应该会什么

你要先回答“为什么要这样部署”，而不是“怎么把服务塞进 K8s”。

## 本周学习重点

- deployment topology
- resource boundary
- health model

## 本周 Cookbook 做法

### Recipe 1：写部署 RFC

### Recipe 2：画服务拓扑

### Recipe 3：定义健康与发布模型

## 本周交付标准

- 部署 RFC
- 拓扑图
- 健康检查模型

## 本周能力焦点

- 生产所有权（`Production Ownership`）
- SLO 思维（`SLO Thinking`）
- 回滚准备度（`Rollback Readiness`）

## 本周知识清单

- 部署 RFC 先定义 workload topology：API、worker、scheduler、gateway 等单元为什么拆。
- Atlas 离开开发者机器以后，config、secret 和 environment boundary 都会发生变化。
- health model 要区分 startup readiness、steady-state readiness 和 fatal liveness。
- 部署模型会反过来塑造代码边界和依赖 ownership，不是部署结束才回头想的事。

## 本周 Atlas 观察路径

- `[Week 40 英文原文](../weeks/week-40.md)`：对照原始英文任务拆分，确认这一周的 deliverables、source reading 和 daily links。
- `[Atlas 主线 Issue（Month 10）](../atlas/issues/month-10/atlas-m10-01.md)`：把本周主题放回 Atlas 主线问题卡，观察范围、验收项和剩余风险。
- `[Week 40 周报](../reports/weekly/week-40-report.md)`：看本周结果、证据和后续动作是如何被记录的。
- `[docs/reports/snapshots/month-10-snapshot.md](../reports/snapshots/month-10-snapshot.md)`：观察部署与运行模型在快照里如何表述。

## 本周 Atlas 落地检查

- 最小落地动作：围绕“写部署 RFC”推进，让至少一项本周交付标准在仓库里可见。
- 必留证据：deployment RFC、topology diagram 和 health model。
- 最大风险：把 Kubernetes 当成 YAML 翻译练习，完全不去思考系统形状。
- 审查不变量：每个部署单元都必须只有一个清楚的 runtime responsibility。
- 本周最先要盯住的交付：部署 RFC

## 本周最小演示场景

- 背景：这一周你需要把“Kubernetes 部署 RFC”从概念推进成可被别人看懂、验证和质疑的工程结果。
- 你至少要能演示：
  - 部署 RFC、拓扑图、健康检查模型
  - “deployment topology”为什么会直接影响 Atlas 的主路径，而不只是局部实现
  - 风险“把 Kubernetes 当成 YAML 翻译练习，完全不去思考系统形状。”是如何被控制、暴露或留证的
- reviewer 大概率会追问：
  - 为什么这周先做“写部署 RFC”，而不是先做别的？
  - 不变量“每个部署单元都必须只有一个清楚的 runtime responsibility。”如果被破坏，会先造成什么后果？
  - 你打算拿什么证据证明这周成果成立？提示：deployment RFC、topology diagram 和 health model。

## 配套案例

- 本周最接近的案例：读 [专题案例包](casebook.md) 中的“案例 10：Kubernetes 部署了，但一点都不像生产系统”。
- 本周最该对照的失败模式：读 [失败场景图谱](failure-patterns.md) 中的“模式五：平台错觉型故障”。
- 推荐顺序：先读本周知识清单，再读案例，再回本周 Atlas 观察路径和周报，检查自己能不能解释坏信号与最小证据。

## 本周工作顺序

- [`RFC`](artifact-playbook.md#artifact-rfc)：先把“写部署 RFC”对应的问题范围、目标、约束和非目标写清楚。
- [`Interface`](artifact-playbook.md#artifact-interface)：围绕“deployment topology”定义输入、输出、owner、错误语义和边界。
- [`Implementation`](artifact-playbook.md#artifact-implementation)：把最小主路径打通，至少让“部署 RFC、拓扑图、健康检查模型”里的一项开始可验证。
- [`Unit Test`](artifact-playbook.md#artifact-unit-test)：优先覆盖最容易击中风险“把 Kubernetes 当成 YAML 翻译练习，完全不去思考系统形状。”的失败路径、边界条件或 contract。
- [`Review`](artifact-playbook.md#artifact-review)：站在 reviewer 视角检查不变量“每个部署单元都必须只有一个清楚的 runtime responsibility。”是否真的被系统保护，而不是只写在文档里。
- [`Benchmark（按需）`](artifact-playbook.md#artifact-benchmark)：本周如果没有性能主题，不强做 benchmark；重点是留下能证明交付成立的测试、drill 或评估证据。
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

1. 如果要把“Kubernetes 部署 RFC”做成 Atlas 可交付能力，你会如何设计主路径和边界？
2. 本周最值得优先验证的失败模式是什么？你会如何收集证据？
3. 这周的 RFC、实现、验证或复盘里，哪一项最能体现你的工程判断？为什么？
4. 如果时间被砍掉一半，你会保留什么，推迟什么？
5. 本周产物如何自然承接到下周，而不是变成一次性文档或一次性代码？

## 每日入口

- [Day 274](days-semester-4/day-274.md): 理解问题和边界
- [Day 275](days-semester-4/day-275.md): 边界与契约设计
- [Day 276](days-semester-4/day-276.md): 最小主路径实现
- [Day 277](days-semester-4/day-277.md): 补全真实可用路径
- [Day 278](days-semester-4/day-278.md): 验证与实验
- [Day 279](days-semester-4/day-279.md): 文档与评审
- [Day 280](days-semester-4/day-280.md): 复盘与 memory

## 配套阅读

- [Week 40](../weeks/week-40.md)
