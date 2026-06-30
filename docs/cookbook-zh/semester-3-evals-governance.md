# 第三学期：状态、评估与治理

最后更新：2026-06-30

## 这一学期要解决什么问题

前两学期让系统“能做事”。

第三学期要解决的是：

如何让系统在真实复杂度下仍然可控。

这意味着你必须进入四个更难的主题：

- Durable Workflow
- Memory Boundary
- Evals / Traces
- Governance / Audit

## 这一学期结束时，你应该具备什么能力

你应该能够：

- 解释为什么长流程 Agent 一定要有状态管理
- 解释为什么 Evals 是工程能力，不只是产品能力
- 解释为什么 Tracing 是排障和治理基础设施
- 解释为什么 Prompt Injection、越权 Tool、数据泄漏必须被系统化处理

## 学期结构

### 第 7 个月：持久化工作流与记忆边界

对应阅读：

- [第七个月：持久化工作流与记忆边界](month-07-durable-workflow.md)

### 第 8 个月：评估、Tracing 与发布门禁

对应阅读：

- [第八个月：评估、Tracing 与发布门禁](month-08-evals-traces.md)

### 第 9 个月：治理、安全与审计

对应阅读：

- [第九个月：治理、安全与审计](month-09-governance-audit.md)

## 这一学期的关键误区

### 误区 1：把 Evals 当作“锦上添花”

错。

没有评估，Agent 系统只能靠感觉迭代。

### 误区 2：把 Tracing 当作“日志增强”

错。

Tracing 不是更好看的日志，而是：

- 路由判断可见
- Tool 选择可见
- 失败位置可见
- 责任边界可见

### 误区 3：把治理放到最后

错。

治理不是“上线前补一层规则”，而是：

从设计开始就决定：

- 谁能调用什么
- 什么动作必须审批
- 什么证据必须保留

## 学期完成标准

第三学期完成时，至少应该满足：

- Atlas 的长流程可以 pause / resume
- Atlas 有 trace 和 eval 基础设施
- Atlas 能做最小 regression gate
- Atlas 能记录 audit 和策略决策

## 推荐下一步

从这里进入：

- [第七个月：持久化工作流与记忆边界](month-07-durable-workflow.md)

如果你想按照中文教材逐周学习，请从这里进入：

- [第三学期周学习索引](week-index-semester-3.md)

如果你想按天推进第三学期第一个月，请从这里进入：

- [第三学期日学习索引（第一个月）](day-index-semester-3-part-1.md)

如果你想完整按天推进第三学期，请从这里进入：

- [第三学期完整日学习索引](day-index-semester-3.md)

如果你想先看长流程、评估和治理的典型问题，请配合阅读：

- [专题案例包](casebook.md)
- [失败场景图谱](failure-patterns.md)
