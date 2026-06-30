# 失败场景图谱

最后更新：2026-06-30

## 这份图谱解决什么问题

专题案例包讲的是“一个完整故事”。

这份图谱讲的是另一件事：

当系统开始变复杂时，失败往往不是按月份发生的，而是按模式重复发生。

你可以把它理解成一本：

- 常见故障模式表
- 排障起点表
- 复盘时的速查表

## 模式一：边界缺失型故障

### 典型表现

- `main`、bootstrap、runtime 和业务初始化混在一起
- tool schema 在调用方和实现方两边各写一份
- agent role、handoff 和审批路径靠口头共识维持

### 根因

- 没有先冻结 contract，就直接写实现
- 没有明确 owner，任何人都能随手改边界
- 把“以后会整理”当成默认策略

### 学习者最容易误判的地方

- 看到系统能跑，就以为边界已经足够清楚
- 把目录结构当成架构边界
- 以为文档稍后补也来得及

### 先回读哪里

- [第一个月：后端基础与服务骨架](month-01-backend-foundation.md)
- [第四个月：Tool Gateway 与接口质量](month-04-tool-gateway.md)
- [第六个月：多 Agent 运行时与审批](month-06-multi-agent-runtime.md)
- [术语总索引](glossary.md) 中的 `Bootstrap`、`Tool Gateway`、`Agent Role Map`

## 模式二：失败语义缺失型故障

### 典型表现

- queue 在重试，但没人知道该不该继续重试
- workflow 可以 resume，但 resume 后结果和原先不一致
- tool timeout 之后，系统既没有 fallback，也没有 operator signal

### 根因

- 没有区分 retryable failure 和 terminal failure
- 没有把 pause / resume / replay 设计成确定性语义
- 没有在 runtime 中显式定义 timeout、取消和回退动作

### 学习者最容易误判的地方

- 觉得“系统很努力地继续执行”就是可靠
- 以为只要存了 checkpoint 就一定能恢复
- 以为 failure handling 是实现细节，不是设计核心

### 先回读哪里

- [第二个月：后台任务与失败处理](month-02-worker-and-failure.md)
- [第七个月：持久化工作流与记忆边界](month-07-durable-workflow.md)
- [Week 07](week-07-retry-timeout-idempotency.md)
- [Week 29](week-29-pause-replay-human-loop.md)

## 模式三：质量证据失真型故障

### 典型表现

- 演示效果不错，但一周后说不清系统到底有没有退化
- 有 trace，但没人能从 trace 里定位为什么分数掉了
- gate fail 了，却不知道该 rollback、该调阈值还是该改 grader

### 根因

- 没有稳定 eval loop
- trace 事件设计没有围绕决策和责任边界
- benchmark workload、环境和阈值定义不稳定

### 学习者最容易误判的地方

- 把 spot check 当 eval
- 把日志量当可观测性
- 把一张性能截图当性能工程

### 先回读哪里

- [第三个月：第一个单 Agent 服务](month-03-first-agent-service.md)
- [第八个月：评估、Tracing 与发布门禁](month-08-evals-traces.md)
- [第十一个月：性能、成本与容量规划](month-11-benchmark-capacity.md)
- [术语总索引](glossary.md) 中的 `Eval Baseline`、`Trace`、`Baseline`

## 模式四：治理失守型故障

### 典型表现

- 高风险 tool 调用路径在不同入口上控制强度不同
- approval 只出现在 PPT 和文档里，不出现在 runtime
- 审计只能回答“发生了什么”，回答不了“为什么被允许”

### 根因

- 没有 threat model 和 risk taxonomy
- policy engine 不集中，规则分散在系统各角落
- 审批和审计没有和执行路径真正连起来

### 学习者最容易误判的地方

- 把治理理解成“上线前加一层 deny rule”
- 把 prompt injection 当成唯一安全问题
- 觉得日志里有一条记录就等于可审计

### 先回读哪里

- [第六个月：多 Agent 运行时与审批](month-06-multi-agent-runtime.md)
- [第九个月：治理、安全与审计](month-09-governance-audit.md)
- [Week 24](week-24-handoff-approval-guardrails.md)
- [Week 35](week-35-governance-rfc.md)
- [Week 37](week-37-injection-leakage-tests.md)

## 模式五：平台错觉型故障

### 典型表现

- 有很多功能，但没有统一 contract
- 可以跑几个 end-to-end demo，但没有 operator guide 和 benchmark evidence
- reviewer 记住了一堆名词，却记不住平台边界

### 根因

- 子系统是逐个长出来的，但从未被重新整合
- 没有跨子系统 drill 和一致性检查
- 缺少 evidence pack 和架构叙事

### 学习者最容易误判的地方

- 把“功能很多”误认为“平台成立”
- 把文档很多误认为“可交接”
- 把一次成功演示误认为“系统可以被运营”

### 先回读哪里

- [第十个月：生产部署与运行模型](month-10-production-kubernetes.md)
- [第十二个月：平台整合与毕业答辩](month-12-platform-integration.md)
- [Week 47](week-47-config-tool-mcp-integration.md)
- [Week 49](week-49-end-to-end-drills.md)
- [Week 52](week-52-year-end-retrospective.md)

## 遇到真实问题时怎么用这份图谱

如果你正在排一个实际问题，建议按这个顺序：

1. 先判断它最像哪一种失败模式
2. 回到对应月份和周文档，看本该建立但没有建立的概念是什么
3. 再去 [源码与证据阅读路径图](source-reading-map.md) 里找对应代码、issue、周报和 review
4. 最后把修复动作写回自己的 RFC、测试、review note 和 memory
