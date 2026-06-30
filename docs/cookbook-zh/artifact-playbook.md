# 交付物写作指南

最后更新：2026-06-30

## 这份指南解决什么问题

现在教材已经告诉你：

- 这一月学什么
- 这一周做什么
- 这一天先读什么再写什么

但学习者还是经常卡在一个实际问题上：

我知道要交 `RFC`、`Interface`、`Review`、`Memory`，可我不知道每一种产物到底该写成什么样。

这份指南就是把仓库要求的顺序写成可执行的交付物说明：

`RFC -> Interface -> Implementation -> Unit Test -> Review -> Benchmark（按需） -> Documentation -> Memory`

## 总原则

每一种交付物都应该同时满足三件事：

1. 能说明你在解决什么问题
2. 能说明你为什么这么做
3. 能留下别人以后还能复查的证据

如果一份产物只能回答“我做了什么”，却回答不了“为什么这样做”和“怎么证明成立”，那它通常还不够。

## 配套示例怎么搭配读

当你需要从“知道名称”进入“能真的写出来”，建议按这个顺序搭配：

1. 先看这份指南，确认每类产物最少要包含什么
2. 再看 [常见交付物词典](deliverable-reference.md)，确认 day 文档里那些产物名词至少应该长什么样
3. 再看 [Review Note 示例](review-note-examples.md)，学习怎么收口判断
4. 再看 [Issue 与 Evidence 示例](issue-and-evidence-examples.md)，学习怎么留下治理与证据产物
5. 如果已经进入月度或季度验收，再看 [Scorecard 与 Review Packet 示例](scorecard-and-review-examples.md)
6. 如果你在写每周收口材料，再看 [周报示例](weekly-report-examples.md)

<a id="artifact-rfc"></a>
## 1. RFC

### 什么时候写

- 当本周主题涉及新边界、新协议、新状态机或新运行时责任时
- 当你还在定义问题和取舍，而不是已经开始写大量实现时

### 最低结构

1. 问题定义
2. 目标与非目标
3. 当前约束
4. 候选方案与取舍
5. 推荐方案
6. 失败模式 / 风险
7. 本周或本月验收口径

### 一句话自检

读完这份 RFC，另一个工程师能不能说清：

- 我们为什么要做它
- 为什么不选别的方案
- 失败时最容易先在哪坏掉

### 最常见错误

- 把 RFC 写成功能愿望清单
- 只有“要做什么”，没有“为什么这样做”
- 不写非目标，导致范围无限膨胀

<a id="artifact-interface"></a>
## 2. Interface

### 什么时候写

- 当你已经基本确认方向，要把边界冻结成接口、schema、event 或 contract 时

### 最低结构

1. 输入
2. 输出
3. owner / 调用方
4. 错误语义
5. 版本或演进假设
6. 超时 / 重试 / 幂等 / 审计等关键控制点

### 一句话自检

如果只把 interface 文档发给调用方，他能不能在不看实现的情况下正确接入？

### 最常见错误

- schema 只写 happy path
- 错误语义不明确
- 接口定义散落在多个地方

<a id="artifact-implementation"></a>
## 3. Implementation

### 什么时候写

- 当 RFC 和 Interface 已经足够清楚，可以打第一条最小主路径时

### 最低交付

1. 一条最小可运行主路径
2. 一份简短实现说明
3. 至少一个明确的不变量

### 实现说明最少要回答

- 输入从哪里来
- 输出到哪里去
- 关键状态在哪里变化
- 哪条边界最不能被打破

### 最常见错误

- 还没冻结边界就开始堆实现
- 顺手把次要复杂度一起做了
- 实现能跑，但没人说得清控制流

<a id="artifact-unit-test"></a>
## 4. Unit Test

### 什么时候写

- 当最小主路径已经跑通，需要把“不变量”和“失败路径”固定下来时

### 最低交付

1. 至少一条正向验证
2. 至少一条失败路径验证
3. 一句说明：这组测试到底在保护什么

### 优先覆盖哪些东西

- contract 被破坏时系统怎么表现
- 状态迁移是否合法
- timeout / retry / fallback 是否符合设计
- 配置或输入错误时是否尽早失败

### 最常见错误

- 只写 happy path
- 追覆盖率，不追风险
- 测试通过了，但没保护真正的不变量

<a id="artifact-review"></a>
## 5. Review

### 什么时候写

- 当你已经有实现和验证证据，要把判断收口成 review note 时

### 最低结构

1. 结论
2. 主要发现
3. 证据
4. 下一步

### 写 review 时最该问的三件事

1. 这项工作真实降低了什么复杂度
2. 增加了哪些维护成本
3. “已经完成”的说法到底有没有足够证据

### 参考

- [Review Note 示例](review-note-examples.md)
- [评审清单与模式提示](review-checklists.md)

### 最常见错误

- 把 review 写成心得体会
- 不先列关键问题，直接讲细节
- 只给判断，不给证据

<a id="artifact-benchmark"></a>
## 6. Benchmark（按需）

### 什么时候写

- 当这一周或这一月的主题直接涉及性能、成本、吞吐、容量或发布门禁时

### 最低结构

1. workload
2. 环境
3. 指标口径
4. before / after 或 baseline
5. 结果解释

### 一句话自检

别人能不能在相同条件下重跑出可比较结果？

### 最常见错误

- workload 不一致就开始比较数字
- 只有数字，没有解释
- 把单次压测截图当作 benchmark

<a id="artifact-documentation"></a>
## 7. Documentation

### 什么时候写

- 当实现已经成立，要让别人接手、运行、排障或继续扩展时

### 常见文档类型

- README 更新
- runbook
- operator guide
- developer guide
- review packet

### 最低结构

1. 这是什么
2. 为什么存在
3. 怎么运行 / 怎么接入 / 怎么排障
4. 已知限制
5. 相关证据入口

### 最常见错误

- 文档只说“怎么做”，不说“为什么这样做”
- 文档里不写限制，导致读者误以为能力已经完整
- 文档和当前实现状态脱节

<a id="artifact-memory"></a>
## 8. Memory

### 什么时候写

- 当这一周或这一月已经沉淀出可复用结论时

### 最低结构

1. `Decision`
2. `Reason`
3. `Future`

### 应该记什么

- 可复用的结论
- 后续仍然会影响同类工作的边界判断
- 典型坑和应对方式

### 不应该记什么

- 流水账
- 一次性命令输出
- 没有复用价值的个人情绪

## 哪类问题去看哪份示例

- “review note 该怎么写”：
  - [Review Note 示例](review-note-examples.md)
- “issue card / evidence entry 该怎么写”：
  - [Issue 与 Evidence 示例](issue-and-evidence-examples.md)
- “scorecard / monthly review / quarter replay 该怎么写”：
  - [Scorecard 与 Review Packet 示例](scorecard-and-review-examples.md)
- “weekly report 该怎么写”：
  - [周报示例](weekly-report-examples.md)
- “memory 该怎么写得更强”：
  - [Memory 示例](memory-note-examples.md)
- “day 文档里这些具体交付物名词到底是什么意思”：
  - [常见交付物词典](deliverable-reference.md)
- “day 层最小交付物到底该怎么写”：
  - [日级常见产物示例](daily-output-examples.md)

## 一套最小闭环长什么样

如果你一周时间有限，至少也要做出下面这个闭环：

1. RFC 草稿
2. 一条最小实现
3. 一条失败路径验证
4. 一份短 review note
5. 一条 memory

这五样加在一起，才更接近“学会了一块”，而不是“看过了一块”。

## 最后一个提醒

不要把这些交付物理解成文书负担。

在这套教材里，它们本身就是学习内容的一部分。

因为 AI Agent 架构工程师要学的，从来不只是把代码写出来，而是把系统判断、证据和交接能力一起写出来。
