# 第一个月：后端基础与服务骨架

最后更新：2026-06-30

## 本月目标

本月不是为了“做出 Agent 功能”。

本月的目标是：

把 Atlas 打造成一个可以长期演进的后端底座。

你要完成的不是一个炫技 Demo，而是一套最小但正确的服务骨架。

## 本月学完后你应该会什么

你应该能独立解释：

- 为什么 Agent 服务必须先有 `bootstrap`
- 为什么 `config.Manager` 必须统一
- 为什么日志入口必须统一
- 为什么数据库、缓存、运行时要在骨架阶段就有边界
- 为什么 Docker、CI、测试属于“第一月”而不是“以后再说”

## 本月学习重点

### 1. Bootstrap 与生命周期
理解：
- 入口层、bootstrap 层和 runtime 层怎么分工。
- 依赖初始化顺序和优雅关闭为什么要先定义。

### 2. 配置分层与唯一入口
理解：
- 默认值、配置文件、环境覆盖的优先级。
- 为什么业务代码不能直接读 `os.Getenv`。

### 3. 统一日志与可观测基础
理解：
- 结构化日志、caller source 和 level 控制。
- 为什么日志抽象要为 trace 和 audit 预留空间。

### 4. 基础设施边界与健康检查
理解：
- 数据库、缓存和 runtime 依赖的 ownership。
- 连接、超时和健康信号为什么属于基础设施契约。

### 5. Docker、CI、测试与交付卫生
理解：
- 容器化、CI 和测试为什么从第一月就必须存在。
- 为什么“能跑”不等于“可交付”。

## 本月 Cookbook 做法

### Recipe 1：搭出最小服务骨架

产物：

- CLI 入口
- 运行时入口
- 启动/停止流程

### Recipe 2：建立统一配置中心

产物：

- `config.Manager`
- `configs/config.yaml`
- 配置默认值、文件、环境变量优先级

### Recipe 3：建立统一日志入口

产物：

- 统一 logger API
- 正确 source 行为
- 统一 level / format 配置

### Recipe 4：接入基础依赖

产物：

- Postgres / Redis 连接边界
- 健康检查
- 最小 smoke test

### Recipe 5：补全交付卫生

产物：

- Dockerfile
- CI
- README 运行说明
- 单元测试

## 本月交付标准

到本月结束时，你至少要能给别人演示：

- Atlas 怎么启动
- 配置在哪里生效
- 日志如何输出
- 依赖怎样初始化
- 测试怎样运行

## 本月常见误区

- 把骨架当目录美化，不处理生命周期和 ownership。
- 在业务代码里直接读环境变量或直接创建 logger / client。
- 代码能跑就算完成，忽略容器、CI 和测试证据。

## 本月必学术语

- `Bootstrap`：负责依赖初始化、启动和关闭流程的统一入口。
- `Runtime`：承载服务真实运行状态与依赖对象的最小执行壳。
- `config.Manager`：统一默认值、配置文件和环境覆盖规则的配置中心。
- `结构化日志`：用稳定字段而不是自由文本记录事件，便于检索、审计和 tracing。
- `健康检查`：系统对外暴露的最小可运行性信号，不等于“进程还活着”。

## 本月知识地图

- 本月真正要打下的是服务地基：统一入口、统一配置、统一日志和统一依赖生命周期。
- 第一月建立的边界会直接决定后面 Agent、MCP、workflow 和 governance 是否还能优雅接上。
- 配置与日志不是“工具细节”，而是所有后续调试、审计和运维能力的前提。
- 数据库、缓存和健康检查应该被看作 runtime contract，而不是零散接线任务。
- Docker、CI、测试和 README 是第一批“证明系统可信”的证据，而不是交付末尾再补的包装。

## 本月最小演示场景

- 背景：一个新同事刚拿到 Atlas 仓库，需要在本地启动服务、切换配置并确认日志与依赖初始化顺序都是可解释的。
- 你至少要能演示：
  - Atlas 如何从统一入口启动和关闭
  - 修改配置后，配置优先级如何生效
  - 日志、db/cache 初始化分别在哪一层发生
- reviewer 大概率会追问：
  - 为什么 `main` 不能直接承担依赖初始化？
  - 哪些边界如果现在不收好，第二个月会先炸？
  - 测试到底在保护哪条启动主路径？

## 本月案例入口

- 先读 [专题案例包](casebook.md) 中的“案例 1：服务能启动，但谁都不敢改”，把本月知识点放回真实工程场景。
- 再读 [失败场景图谱](failure-patterns.md) 中的“模式一：边界缺失型故障”，理解这一月最容易反复出现的故障模式。
- 读完后回到本月的周/日文档，检查自己能不能说清楚：错误做法、坏信号和最小证据分别是什么。

## 本月实战清单

- [`RFC`](artifact-playbook.md#artifact-rfc)：写服务骨架 RFC，冻结入口层、bootstrap 层、runtime 层和基础设施层的责任。
- [`Interface`](artifact-playbook.md#artifact-interface)：定义 `config.Manager`、logger、runtime、db/cache 依赖边界，明确谁能直接接触环境变量和外部依赖。
- [`Implementation`](artifact-playbook.md#artifact-implementation)：打通 CLI / HTTP 共用的启动与关闭主路径。
- [`Unit Test`](artifact-playbook.md#artifact-unit-test)：补 bootstrap、config、logger 和健康检查相关测试或 smoke check。
- [`Review`](artifact-playbook.md#artifact-review)：做一次边界泄漏自查：是否仍有 `os.Getenv`、自建 logger、散落 client 初始化。
- [`Benchmark（按需）`](artifact-playbook.md#artifact-benchmark)：本月不做性能 benchmark，但至少记录启动/健康 smoke 结果，证明主路径稳定。
- [`Documentation`](artifact-playbook.md#artifact-documentation)：完善 README、运行说明和配置优先级说明。
- [`Memory`](artifact-playbook.md#artifact-memory)：沉淀第一月架构边界：入口、配置、日志和依赖初始化为什么这样分。

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

1. 如果要向资深 reviewer 解释“后端基础与服务骨架”，你会怎样说明这月的架构取舍？
2. 本月哪条主路径最容易失败？你会如何设计验证证据？
3. 本月引入的哪个接口、状态或契约最可能在未来变化？为什么？
4. 如果 Atlas 的调用量增长 10 倍，本月设计会先在哪个点上断裂？
5. 本月消除了哪些技术债，又留下了哪些新债？
6. 哪个周度产物最能证明你的工程判断？为什么？
7. 你会怎样把本月所有周主题串成一个完整系统故事，而不是零散任务？
8. 你会出示哪些测试、trace、日志或运行证据，证明本月成果已经可交付？

## 本月推荐配套阅读

- [Month 01](../months/month-01.md)
- [Month 01 Backlog](../atlas/backlog/month-01.md)
- [Month 01 Issue Set](../atlas/issues/month-01/README.md)
- [Week 01 Report](../reports/weekly/week-01-report.md)
- [第一个月日学习索引](day-index-month-01.md)

## 本月结束时问自己

1. Atlas 现在是不是一个干净的后端服务，而不是一堆散代码？
2. 我以后接 Agent Runtime、Tool、MCP 时，现有边界够不够承接？
3. 如果现在让别人接手这个仓库，他能不能启动、理解、扩展？
