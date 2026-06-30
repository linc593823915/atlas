# 第二学期：接口、协议与运行时

最后更新：2026-06-30

## 这一学期要解决什么问题

第一学期解决的是“把 Agent 服务做出来”。

第二学期要解决的是：

如何把一个能工作的 Agent 服务，升级为一个具备平台边界的系统。

也就是说，你不再只写功能，而要开始设计：

- Tool 接口
- MCP 协议边界
- Multi-Agent Runtime
- Approval Flow

## 这一学期结束时，你应该具备什么能力

你应该能够：

- 解释为什么 Tool 不能散落在业务代码里
- 解释为什么 MCP 对 Agent 平台重要
- 解释多 Agent 为什么需要 handoff、approval、guardrails
- 设计比“Prompt 拼接”更稳定的运行时边界

## 学期结构

### 第 4 个月：统一 Tool Gateway

目标：

- 把工具调用统一到一个入口
- 建立 schema、validation、timeout、audit 的工程边界

对应阅读：

- [第四个月：Tool Gateway 与接口质量](month-04-tool-gateway.md)

### 第 5 个月：MCP 工程化

目标：

- 学会把能力暴露成 MCP Surface
- 理解 tools/resources/contracts/versioning

对应阅读：

- [第五个月：MCP 工程化](month-05-mcp-engineering.md)

### 第 6 个月：多 Agent 运行时

目标：

- 建立 agent role、handoff、approval、guardrail 的运行模型

对应阅读：

- [第六个月：多 Agent 运行时与审批](month-06-multi-agent-runtime.md)

## 这一学期最重要的认知升级

### 从“功能”升级到“边界”

第一学期你在问：

- 怎么把功能做出来？

第二学期你必须开始问：

- 这个能力应该暴露成什么接口？
- 以后别人怎么接？
- 如果以后有 10 个 Tool，而不是 1 个，会不会崩？

### 从“单链路”升级到“系统协作”

这时候你要开始理解：

- Tool 为什么需要统一网关
- Agent 为什么需要角色边界
- 审批为什么不能后补

## 学期完成标准

第二学期完成时，至少应该满足：

- Atlas 有统一 Tool Gateway
- Atlas 能暴露至少一个 MCP Surface
- Atlas 有多 Agent 编排雏形
- 高风险动作可以走 approval path
- 编排与接口变化有对应测试和文档

## 推荐下一步

从这里进入：

- [第四个月：Tool Gateway 与接口质量](month-04-tool-gateway.md)

然后继续配合底层台账推进：

- [Week Index](../weeks/README.md)
- [Day Index](../days/README.md)

如果你想按照中文教材逐周学习，请从这里进入：

- [第二学期周学习索引](week-index-semester-2.md)

如果你想按天推进第二学期第一个月，请从这里进入：

- [第二学期日学习索引（第一个月）](day-index-semester-2-part-1.md)

如果你想完整按天推进第二学期，请从这里进入：

- [第二学期完整日学习索引](day-index-semester-2.md)

如果你想先看接口、协议和多 Agent 最常见的翻车方式，请配合阅读：

- [专题案例包](casebook.md)
- [失败场景图谱](failure-patterns.md)
