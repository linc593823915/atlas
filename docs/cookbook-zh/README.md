# AI Agent 架构工程师一年自学教材

最后更新：2026-06-30

## 这是什么

这不是一份普通的学习清单。

这是一套为期一年的 `cookbook` 式自学教材，目标是在一年内，从具备后端基础的工程师，成长为能够独立设计、构建、评估、治理、运维 AI Agent 系统的 `AI Agent 架构工程师`。

这套教材的核心特点是：

- 不是只学提示词，而是学完整系统
- 不是只看文档，而是边学边做 Atlas 工程
- 不是只做 Demo，而是按企业级标准推进
- 不是只追求“会用模型”，而是追求“能设计平台”

## 一年后的目标

学完这一年后，你应该能够：

- 设计 Agent Runtime、Tool Schema、MCP Contract、Workflow State
- 构建带有 Tracing、Evals、Approval、Audit 的 Agent 系统
- 解释成本、时延、可靠性、安全性、可演进性的工程权衡
- 把 Atlas 从一个普通代码仓库推进为一个内部 Agent Platform 雏形

## 如何使用这本教材

请先读：

- [如何使用这本教材](how-to-use.md)
- [全年总览](year-overview.md)
- [术语总索引](glossary.md)
- [源码与证据阅读路径图](source-reading-map.md)

然后按顺序进入：

- [第一学期：基础与第一个 Agent](semester-1-foundation.md)
- [第二学期：接口、协议与运行时](semester-2-interface-runtime.md)
- [第三学期：状态、评估与治理](semester-3-evals-governance.md)
- [第四学期：生产、性能与平台整合](semester-4-production-platform.md)

## 学期目录

- [第一学期：基础与第一个 Agent](semester-1-foundation.md)
- [第二学期：接口、协议与运行时](semester-2-interface-runtime.md)
- [第三学期：状态、评估与治理](semester-3-evals-governance.md)
- [第四学期：生产、性能与平台整合](semester-4-production-platform.md)

## 辅助索引

- [术语总索引](glossary.md)
- [源码与证据阅读路径图](source-reading-map.md)
- [主题级源码讲解](code-walkthroughs.md)
- [代码阅读 Patterns](code-reading-patterns.md)
- [外部阅读提示解读](external-reading-guides.md)
- [仓库内阅读线索](internal-reading-cues.md)
- [原子概念补充](atomic-concepts.md)
- [专题案例包](casebook.md)
- [失败场景图谱](failure-patterns.md)
- [评审清单与模式提示](review-checklists.md)
- [Review Note 示例](review-note-examples.md)
- [交付物写作指南](artifact-playbook.md)
- [常见交付物词典](deliverable-reference.md)
- [日级常见产物示例](daily-output-examples.md)
- [日级执行动作手册](daily-execution-playbook.md)
- [延伸追问回答指南](reflection-question-guide.md)
- [Issue 与 Evidence 示例](issue-and-evidence-examples.md)
- [Scorecard 与 Review Packet 示例](scorecard-and-review-examples.md)
- [周报示例](weekly-report-examples.md)
- [Memory 示例](memory-note-examples.md)
- [PDF 生成接口](pdf-generation-interface.md)

## PDF 生成

如果需要生成离线阅读版总书，在仓库根目录执行：

```bash
make cookbook-pdf
```

默认输出文件：

```text
docs/cookbook-zh/atlas-cookbook-zh.pdf
```

生成脚本会递归包含本目录下全部 Markdown 文件，并按 cookbook 阅读顺序合并。详细参数见 [PDF 生成接口](pdf-generation-interface.md)。

## 周度目录

- [第一学期周学习索引](week-index-semester-1.md)
- [第二学期周学习索引](week-index-semester-2.md)
- [第三学期周学习索引](week-index-semester-3.md)
- [第四学期周学习索引](week-index-semester-4.md)

## 日度目录

- [第一学期完整日学习索引](day-index-semester-1.md)
- [第二学期完整日学习索引](day-index-semester-2.md)
- [第三学期完整日学习索引](day-index-semester-3.md)
- [第四学期完整日学习索引](day-index-semester-4.md)

## 月度目录

- [第一个月：后端基础与服务骨架](month-01-backend-foundation.md)
- [第二个月：后台任务与失败处理](month-02-worker-and-failure.md)
- [第三个月：第一个单 Agent 服务](month-03-first-agent-service.md)
- [第四个月：Tool Gateway 与接口质量](month-04-tool-gateway.md)
- [第五个月：MCP 工程化](month-05-mcp-engineering.md)
- [第六个月：多 Agent 运行时与审批](month-06-multi-agent-runtime.md)
- [第七个月：持久化工作流与记忆边界](month-07-durable-workflow.md)
- [第八个月：评估、Tracing 与发布门禁](month-08-evals-traces.md)
- [第九个月：治理、安全与审计](month-09-governance-audit.md)
- [第十个月：生产部署与运行模型](month-10-production-kubernetes.md)
- [第十一个月：性能、成本与容量规划](month-11-benchmark-capacity.md)
- [第十二个月：平台整合与毕业答辩](month-12-platform-integration.md)

## 教材结构

这套教材表面上是中文自学手册，底层仍然连接原来的工程治理系统：

- `Semester`
  四个学期，定义能力阶段
- `Month`
  每月一个主题和 Atlas 里程碑
- `Week`
  每周一个可交付目标
- `Day`
  每天一个最小学习与构建动作
- `Atlas`
  主线工程，所有学习都必须落到 Atlas

## 教材层与管理层

你现在看到的中文 `cookbook` 是“学员端”。

仓库里仍保留完整的“教师端/管理端”资产，供后续评审与治理使用：

- [ROADMAP](../ROADMAP.md)
- [Operations Index](../operations/README.md)
- [Governance Index](../governance/README.md)
- [Reporting Records Index](../reports/README.md)

## 建议学习节奏

- 每周投入：`12-15 小时`
- 每天至少完成一次：`Read -> Write -> Build -> Output`
- 每周必须有一个“可交付结果”
- 每月必须有一个 Atlas 里程碑
- 每季度必须做一次架构复盘

## 不适合的人

如果你只想快速学几个 Prompt 技巧，或者只想做几个聊天机器人 Demo，这套教材会显得过重。

这套教材适合：

- 想在未来 1-3 年进入 AI Agent 平台、基础设施、架构方向的人
- 想用企业级方法自学，而不是碎片化学习的人
- 能接受“学习 = 阅读 + 编码 + 评审 + 复盘”的人
