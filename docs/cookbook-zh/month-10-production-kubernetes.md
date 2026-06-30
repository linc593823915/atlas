# 第十个月：生产部署与运行模型

最后更新：2026-06-30

## 本月目标

本月要把 Atlas 从“开发者机器上的系统”推进到“生产式部署的系统”。

重点不在于 Kubernetes 本身，而在于：

- 健康检查
- 资源边界
- 发布
- 回滚
- SLO

## 本月学完后你应该会什么

你应该能解释：

- 为什么 Agent 平台必须考虑 rollout / rollback
- 为什么 SLO 是架构问题，不只是运维指标
- 为什么部署模型会反过来影响架构边界

## 本月学习重点

### 1. Deployment Topology 与 Runtime Boundary
理解：
- 哪些单元要拆分部署。
- config / secret / environment 如何进入生产。

### 2. Manifests、Probes 与资源控制
理解：
- readiness、liveness、startup probe 的语义差异。
- request/limit 和扩缩容信号。

### 3. Alert、Rollback 与 SLO
理解：
- SLI/SLO 如何定义。
- 什么时候该 page、什么时候该 rollback。

### 4. Benchmark Contract 起点
理解：
- 为什么部署期就要开始定义可比较的 workload contract。
- 生产式部署如何影响后续测量。

## 本月 Cookbook 做法

### Recipe 1：写部署 RFC

### Recipe 2：做 manifests / Helm

### Recipe 3：加 probes 与资源限制

### Recipe 4：定义 alert 与 rollback

## 本月交付标准

你至少要能演示：

- Atlas 如何被部署
- 如何判断它健康
- 什么时候该 rollback

## 本月必学术语

- `Deployment Topology`：服务在生产环境中的部署单元和连接关系。
- `Readiness Probe`：判断实例是否可以接收流量的探针。
- `Liveness Probe`：判断实例是否已进入不可恢复错误状态的探针。
- `SLO`：系统对可用性、时延等关键体验做出的服务级目标承诺。
- `Rollback Criteria`：决定何时必须回退版本的客观信号与规则。

## 本月知识地图

- 生产部署关注的不是 YAML 数量，而是 Atlas 在真实运行环境下的边界和承诺。
- probe、resource 和 rollout 规则会反向塑造代码结构和依赖生命周期。
- SLO、告警和 rollback 是生产系统的操作 contract，不只是运维附录。
- 部署模型和 benchmark contract 是相连的：如果部署环境不稳定，后续性能证据也不可信。

## 本月常见误区

- 把 Kubernetes 当成文件格式迁移。
- probe 全部返回 200 就算健康。
- 没有明确 rollback criteria 就开始部署。

## 本月最小演示场景

- 背景：Atlas 以生产式方式部署：实例健康、资源边界清楚，告警触发后 operator 能判断是否需要 rollback。
- 你至少要能演示：
  - deployment topology 与 probes 如何工作
  - SLO / alert 的含义是什么
  - rollback 触发条件如何被定义
- reviewer 大概率会追问：
  - 为什么 Kubernetes 不是写 YAML 就结束？
  - probe 为什么不能全部返回 200？
  - 谁来决定 rollback？依据是什么？

## 本月案例入口

- 先读 [专题案例包](casebook.md) 中的“案例 10：Kubernetes 部署了，但一点都不像生产系统”，把本月知识点放回真实工程场景。
- 再读 [失败场景图谱](failure-patterns.md) 中的“模式五：平台错觉型故障”，理解这一月最容易反复出现的故障模式。
- 读完后回到本月的周/日文档，检查自己能不能说清楚：错误做法、坏信号和最小证据分别是什么。

## 本月实战清单

- [`RFC`](artifact-playbook.md#artifact-rfc)：写部署 RFC，冻结 deployment topology、health model 和 rollout / rollback 假设。
- [`Interface`](artifact-playbook.md#artifact-interface)：定义 probe、resource boundary、SLI/SLO 和 alert / rollback contract。
- [`Implementation`](artifact-playbook.md#artifact-implementation)：落地 manifests / charts、probes 和最小部署主路径。
- [`Unit Test`](artifact-playbook.md#artifact-unit-test)：覆盖配置渲染、探针语义、回滚条件和部署 smoke 验证。
- [`Review`](artifact-playbook.md#artifact-review)：检查部署模型是否真的反映系统边界，而不是模板复制。
- [`Benchmark（按需）`](artifact-playbook.md#artifact-benchmark)：本月不做完整性能 benchmark，但至少记录部署 / 健康 / 告警演练结果。
- [`Documentation`](artifact-playbook.md#artifact-documentation)：补 deployment runbook、alert policy 和 rollback guide。
- [`Memory`](artifact-playbook.md#artifact-memory)：沉淀生产部署月里最重要的操作 contract。

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

1. 如果要向资深 reviewer 解释“生产部署与运行模型”，你会怎样说明这月的架构取舍？
2. 本月哪条主路径最容易失败？你会如何设计验证证据？
3. 本月引入的哪个接口、状态或契约最可能在未来变化？为什么？
4. 如果 Atlas 的调用量增长 10 倍，本月设计会先在哪个点上断裂？
5. 本月消除了哪些技术债，又留下了哪些新债？
6. 哪个周度产物最能证明你的工程判断？为什么？
7. 你会怎样把本月所有周主题串成一个完整系统故事，而不是零散任务？
8. 你会出示哪些测试、trace、日志或运行证据，证明本月成果已经可交付？

## 本月推荐配套阅读

- [Month 10](../months/month-10.md)
- [Month 10 Backlog](../atlas/backlog/month-10.md)
- [Month 10 Issue Set](../atlas/issues/month-10/README.md)
- [第四学期完整日学习索引](day-index-semester-4.md)
