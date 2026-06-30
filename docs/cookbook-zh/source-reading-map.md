# 源码与证据阅读路径图

最后更新：2026-06-30

## 先说清楚：这张图不是“源码目录”

从当前工作树看，仓库里最显性的 Go 代码主要集中在前期基础设施层：

- 入口与启动
- 配置与日志
- 基础 runtime
- programops / scripts

后半年的很多主题，当前更多是通过：

- `docs/atlas/issues/`
- `docs/reports/`
- `docs/operations/`
- `docs/governance/`

来表达设计、验收和治理证据。

所以学习时不要执着于“为什么这里没有一整套对应代码文件”。

这套教材的真实阅读顺序应该是：

1. 先看中文 month/week/day，弄清现在在学什么
2. 再看英文 week 原文，确认原始 deliverable 和 source reading
3. 再看 Atlas issue / 周报 / 复盘，理解范围、证据和残余风险
4. 最后回到真实代码或治理材料，观察这些概念在仓库里怎样落地

## 第一学期：服务地基、任务系统与第一个 Agent

### 第一个月：后端基础与服务骨架

- 中文入口：[第一个月](month-01-backend-foundation.md)、[Week 01](week-01-learning-system.md)、[Week 03](week-03-config-logger-storage.md)
- 必看代码：
  - [cmd/atlas/main.go](../../cmd/atlas/main.go)
  - [internal/bootstrap/bootstrap.go](../../internal/bootstrap/bootstrap.go)
  - [internal/app/root.go](../../internal/app/root.go)
  - [internal/app/serve.go](../../internal/app/serve.go)
  - [internal/config/manager.go](../../internal/config/manager.go)
  - [internal/logger/logger.go](../../internal/logger/logger.go)
- 必看测试：
  - [internal/bootstrap/bootstrap_test.go](../../internal/bootstrap/bootstrap_test.go)
  - [internal/logger/logger_test.go](../../internal/logger/logger_test.go)
  - [internal/config/manager_test.go](../../internal/config/manager_test.go)
- 必看证据：
  - [Month 01 Issue Set](../atlas/issues/month-01/README.md)
  - [Week 01 Report](../reports/weekly/week-01-report.md)
  - [Month 01 Scorecard](../reports/monthly/month-01-scorecard.md)
- 观察重点：
  - 入口层、bootstrap 层和 runtime 层怎么分工
  - 配置、日志和依赖初始化如何避免散在业务代码里
  - 第一月的测试到底在保护哪些边界

### 第二个月：后台任务与失败处理

- 中文入口：[第二个月](month-02-worker-and-failure.md)、[Week 05](week-05-job-runner-rfc.md)、[Week 07](week-07-retry-timeout-idempotency.md)
- 当前仓库更适合看的材料：
  - [Month 02 Issue Set](../atlas/issues/month-02/README.md)
  - [Week 05 Report](../reports/weekly/week-05-report.md)
  - [Month 02 Scorecard](../reports/monthly/month-02-scorecard.md)
  - [Evidence Registry](../operations/evidence-registry.md)
  - [Issue Status Board](../operations/issue-status-board.md)
- 观察重点：
  - 队列、worker、retry、dead-letter 这些能力在当前仓库里更多以设计/运维证据出现
  - 没有现成实现时，要练习从 issue、周报和评分卡反推出 contract 与 failure model

### 第三个月：第一个单 Agent 服务

- 中文入口：[第三个月](month-03-first-agent-service.md)、[Week 09](week-09-agent-rfc.md)、[Week 10](week-10-structured-output.md)
- 必看材料：
  - [Atlas 主线说明](../atlas/atlas.md)
  - [Month 03 Issue Set](../atlas/issues/month-03/README.md)
  - [Week 09 Report](../reports/weekly/week-09-report.md)
  - [Month 03 Scorecard](../reports/monthly/month-03-scorecard.md)
  - [Month 03 Review](../reviews/archive/month-03-review.md)
- 观察重点：
  - 单 Agent 月份在当前仓库里更强调 contract、评估和运维说明，而不是大量现成模型代码
  - 重点读清 tool inventory、structured output 和 eval baseline 是怎样被验收的

## 第二学期：接口、协议与运行时

### 第四个月：Tool Gateway 与接口质量

- 中文入口：[第四个月](month-04-tool-gateway.md)、[Week 14](week-14-tool-gateway-rfc.md)
- 必看材料：
  - [Month 04 Issue Set](../atlas/issues/month-04/README.md)
  - [Week 14 Report](../reports/weekly/week-14-report.md)
  - [Month 04 Scorecard](../reports/monthly/month-04-scorecard.md)
  - [Month 04 Review](../reviews/archive/month-04-review.md)
  - [Google-Style Capability Model](../framework/google-style-capability-model.md)
- 观察重点：
  - gateway 在这里主要作为设计 contract、验证 contract、review contract 的入口
  - 没有完整实现代码时，先学会读 issue、周报、review 里的接口质量语言

### 第五个月：MCP 工程化

- 中文入口：[第五个月](month-05-mcp-engineering.md)、[Week 18](week-18-mcp-rfc.md)
- 必看材料：
  - [Month 05 Issue Set](../atlas/issues/month-05/README.md)
  - [Week 18 Report](../reports/weekly/week-18-report.md)
  - [Month 05 Scorecard](../reports/monthly/month-05-scorecard.md)
  - [Month 05 Review](../reviews/archive/month-05-review.md)
  - [Google Public Sources](../framework/google-public-sources.md)
- 观察重点：
  - 先学协议面 scope、capability inventory 和兼容性，而不是急着找一整套 server 实现
  - 对照 issue 与 scorecard 看什么叫“可消费协议面”

### 第六个月：多 Agent 运行时与审批

- 中文入口：[第六个月](month-06-multi-agent-runtime.md)、[Week 22](week-22-agent-runtime-rfc.md)、[Week 24](week-24-handoff-approval-guardrails.md)
- 必看材料：
  - [Month 06 Issue Set](../atlas/issues/month-06/README.md)
  - [Week 22 Report](../reports/weekly/week-22-report.md)
  - [Month 06 Scorecard](../reports/monthly/month-06-scorecard.md)
  - [Month 06 Review](../reviews/archive/month-06-review.md)
  - [Role Gap Analysis](../governance/role-gap-analysis.md)
  - [Committee Decision Log](../governance/committee-decision-log.md)
- 观察重点：
  - 多 Agent 的阅读核心是 ownership、shared context、approval 和 audit，而不是 agent 数量
  - 治理文档是 runtime 证据的一部分，不是附属阅读

## 第三学期：状态、评估与治理

### 第七个月：持久化工作流与记忆边界

- 中文入口：[第七个月](month-07-durable-workflow.md)、[Week 27](week-27-durable-workflow-rfc.md)
- 必看材料：
  - [Month 07 Issue Set](../atlas/issues/month-07/README.md)
  - [Week 27 Report](../reports/weekly/week-27-report.md)
  - [Month 07 Snapshot](../reports/snapshots/month-07-snapshot.md)
  - [Month 07 Scorecard](../reports/monthly/month-07-scorecard.md)
  - [Month 07 Review](../reviews/archive/month-07-review.md)
- 观察重点：
  - 学会从 RFC、快照和 review 里读出 state model、checkpoint boundary 和恢复语义
  - 注意 workflow state 和 memory boundary 是两类不同东西

### 第八个月：评估、Tracing 与发布门禁

- 中文入口：[第八个月](month-08-evals-traces.md)、[Week 31](week-31-eval-loop-rfc.md)、[Week 33](week-33-metrics-release-gates.md)
- 必看材料：
  - [Month 08 Issue Set](../atlas/issues/month-08/README.md)
  - [Week 31 Report](../reports/weekly/week-31-report.md)
  - [Month 08 Scorecard](../reports/monthly/month-08-scorecard.md)
  - [Month 08 Review](../reviews/archive/month-08-review.md)
  - [Review Metrics](../operations/review-metrics.md)
  - [Dashboard Export](../operations/dashboard-export.md)
- 观察重点：
  - eval loop、trace、gate 在当前仓库里主要通过运行指标、报表和 review 被描述
  - 学会读这些材料里的“质量语言”，它们就是后续平台门禁的基础

### 第九个月：治理、安全与审计

- 中文入口：[第九个月](month-09-governance-audit.md)、[Week 35](week-35-governance-rfc.md)、[Week 37](week-37-injection-leakage-tests.md)
- 必看材料：
  - [Month 09 Issue Set](../atlas/issues/month-09/README.md)
  - [Week 35 Report](../reports/weekly/week-35-report.md)
  - [Month 09 Scorecard](../reports/monthly/month-09-scorecard.md)
  - [Month 09 Review](../reviews/archive/month-09-review.md)
  - [Committee Decision Examples](../governance/committee-decision-examples.md)
  - [Sample Evidence Entries](../governance/sample-evidence-entries.md)
- 观察重点：
  - 治理月的“源码”很多时候就是 policy、decision log、review packet 和 adversarial evidence
  - 不要只盯安全名词，要看控制如何被证明

## 第四学期：生产、性能与平台整合

### 第十个月：生产部署与运行模型

- 中文入口：[第十个月](month-10-production-kubernetes.md)、[Week 40](week-40-kubernetes-deployment-rfc.md)、[Week 42](week-42-alerts-rollback-slo.md)
- 必看材料：
  - [Month 10 Issue Set](../atlas/issues/month-10/README.md)
  - [Week 40 Report](../reports/weekly/week-40-report.md)
  - [Month 10 Snapshot](../reports/snapshots/month-10-snapshot.md)
  - [Month 10 Scorecard](../reports/monthly/month-10-scorecard.md)
  - [Month 10 Review](../reviews/archive/month-10-review.md)
- 观察重点：
  - 这一月的重点是 deployment topology、probe、rollback 和 SLO contract
  - 结合 snapshot 和 review 看“生产式部署”到底意味着什么

### 第十一个月：性能、成本与容量规划

- 中文入口：[第十一个月](month-11-benchmark-capacity.md)、[Week 44](week-44-baseline-latency-cost.md)、[Week 45](week-45-optimization-capacity.md)
- 必看材料：
  - [Month 11 Issue Set](../atlas/issues/month-11/README.md)
  - [Week 44 Report](../reports/weekly/week-44-report.md)
  - [Month 11 Snapshot](../reports/snapshots/month-11-snapshot.md)
  - [Month 11 Scorecard](../reports/monthly/month-11-scorecard.md)
  - [Month 11 Review](../reviews/archive/month-11-review.md)
- 观察重点：
  - 先看 baseline 如何被定义和记录，再看优化如何被解释
  - 性能、成本和容量要一起读，不能拆开理解

### 第十二个月：平台整合与毕业答辩

- 中文入口：[第十二个月](month-12-platform-integration.md)、[Week 47](week-47-config-tool-mcp-integration.md)、[Week 49](week-49-end-to-end-drills.md)、[Week 52](week-52-year-end-retrospective.md)
- 必看材料：
  - [Month 12 Issue Set](../atlas/issues/month-12/README.md)
  - [Week 47 Report](../reports/weekly/week-47-report.md)
  - [Month 12 Snapshot](../reports/snapshots/month-12-snapshot.md)
  - [Month 12 Scorecard](../reports/monthly/month-12-scorecard.md)
  - [Month 12 Review](../reviews/archive/month-12-review.md)
  - [Exec Dashboard Pack](../governance/exec-dashboard-pack.md)
  - [Steering Committee Pack](../governance/steering-committee-pack.md)
  - [Annual Review Memo](../governance/annual-review-memo.md)
- 观察重点：
  - 平台整合月最重要的是把零散能力串成一条可 defend 的系统故事
  - 文档、benchmark、drill 和 review packet 共同组成毕业证据包

## 每周固定阅读顺序

如果你不想每次都重新思考从哪开始，直接按这个顺序读：

1. 中文周文档
2. 对应英文周文档
3. 对应 Atlas issue
4. 对应周报
5. 对应月度 scorecard / snapshot / review
6. 最后再回到真实代码或治理材料

这样做的好处是：

- 先建立上下文
- 再看实现或证据
- 不会一上来就被代码细节或材料堆淹没

## 做 review 或排错时再回什么

如果你已经读完源码、issue、周报，但还是不知道该怎么挑错，直接再配合：

- [评审清单与模式提示](review-checklists.md)
- [Review Note 示例](review-note-examples.md)

它更适合在下面这些时候打开：

- 准备写 review note
- 准备做接口评审
- 准备复盘失败场景
- 准备判断 benchmark 结果到底能不能信

如果你已经找到文件，但还是不知道“这一组代码到底在保护什么边界”，继续看：

- [主题级源码讲解](code-walkthroughs.md)
- [代码阅读 Patterns](code-reading-patterns.md)
