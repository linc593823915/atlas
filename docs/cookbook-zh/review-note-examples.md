# Review Note 示例

最后更新：2026-06-30

## 这份示例怎么用

很多学习者知道要做 review，却不知道一个“像样的 review note”到底长什么样。

这份文档不提供唯一模板。

它提供的是几种高频场景下，什么叫：

- 讲清上下文
- 先说关键问题
- 给出证据
- 留下下一步动作

## 示例 1：第一月服务骨架评审

### 适用场景

- 启动主路径
- bootstrap / runtime 分层
- config / logger 接线

### 示例写法

```md
## Review Note

### 结论

当前改动已经把 `main -> app -> bootstrap -> runtime` 这条链路拉直，但 runtime 仍然过空，后续接 db/cache/http 时容易把职责重新推回 `serve.go`。

### 主要发现

1. `serve.go` 仍然直接负责错误打印和生命周期控制，后续应收回统一 logger / error flow。
2. `config.Manager` 的读取入口已经统一，但新增配置项时仍要警惕默认值、文件值、环境变量值三处不同步。
3. `runtime.Runtime` 作为边界已经留对，但现在缺少一条能保护未来扩展方向的测试说明。

### 证据

- `cmd/atlas/main.go` 只负责退出码，说明入口层边界是干净的。
- `internal/bootstrap/bootstrap.go` 已经统一初始化 config、logger 和 runtime。
- `internal/config/loader.go` 明确体现了 `default -> file -> env -> validate` 的顺序。

### 下一步

- 给 runtime 扩展点补一个短说明，明确下一批依赖应该接到哪里。
- 把 `serve.go` 的错误输出逐步统一到 internal/logger。
```

## 示例 2：Tool Gateway contract 评审

### 适用场景

- Tool Gateway
- schema / validation
- timeout / audit hook

### 示例写法

```md
## Review Note

### 结论

这项设计已经把 tool invocation 的控制面收回到了 gateway，但 contract 仍然需要补一条“执行后结果校验”的说明，否则 result normalization 会成为隐含行为。

### 主要发现

1. request-side schema 已经明确，但 response-side schema 仍然偏弱。
2. timeout policy 只写了 duration，没有写谁负责取消和谁负责 fallback。
3. audit hook 已经被列为必要控制，但尚未说明哪些字段是最小留痕集合。

### 证据

- 本周知识清单已经明确 schema、timeout、audit 是统一入口控制面。
- issue / review 证据显示 negative path 是核心验证面，而不是附属测试。

### 下一步

- 在 contract 中补 response validation。
- 给 timeout policy 补 budget owner 和 fail behavior。
- 给 audit event 补最小字段集合。
```

## 示例 3：durable workflow 恢复语义评审

### 适用场景

- checkpoint
- pause / resume / replay
- stateful testing

### 示例写法

```md
## Review Note

### 结论

当前方案已经把 checkpoint 边界列出来，但 resume cursor 仍然不够明确。只要同一个 checkpoint 还能导出多个“下一步”，恢复语义就不算成立。

### 主要发现

1. workflow state 和 memory boundary 已经被区分，这是正确方向。
2. 人工介入点有说明，但 provenance 还不够强。
3. replay 语义如果不区分“重放证据”和“重做副作用”，后续 drill 会出现假恢复。

### 证据

- 本月主题已经明确强调 checkpoint、resume 和 human intervention。
- 失败模式图谱中，恢复语义不确定被视为核心故障模式。

### 下一步

- 补 resume cursor 的唯一性说明。
- 明确人工修改状态后的 provenance 记录格式。
- 给 replay 加一条副作用边界说明。
```

## 示例 4：评估与 gate 设计评审

### 适用场景

- eval loop
- trace
- regression gate

### 示例写法

```md
## Review Note

### 结论

系统已经具备指标、trace 和 gate 的基本结构，但当前最大问题不是“缺数据”，而是“数据还没有形成可操作决策”。

### 主要发现

1. score 与 trace 的可回溯关系需要再写清楚。
2. gate fail 之后的动作分流仍然不够明确：是回滚、调阈值还是追数据质量问题。
3. dataset versioning 已被提到，但还没形成 review 时可快速核对的字段。

### 证据

- 月文档已经强调 gate 不是 dashboard，而是工程决策机制。
- 失败模式图谱已经把“质量证据失真”列为独立故障模式。

### 下一步

- 把 score -> trace -> dataset row 的追溯链写成短文档。
- 给 gate fail 增加动作分流表。
- 给 dataset row 定义 reviewer 能快速核查的 version 字段。
```

## 示例 5：治理与审计评审

### 适用场景

- threat model
- policy engine
- adversarial tests

### 示例写法

```md
## Review Note

### 结论

当前设计已经开始从 threat model 反推 control，但 policy 仍偏文档化，离“执行路径里真的能拦住”还有一段距离。

### 主要发现

1. 高风险动作识别得不错，但 enforcement point 还要继续细化。
2. 审计记录已经能说明“发生了什么”，但还不足以完整解释“为什么被允许”。
3. adversarial test 有了样子，但 benign / malicious 区分证据还不够。

### 证据

- month-09 的目标明确要求 control、audit 和 misuse evidence 同时存在。
- 失败模式图谱把 policy 分散和审计不完整定义为治理失守核心信号。

### 下一步

- 补 enforcement point 和 deny path。
- 给 audit event 补决策输入字段。
- 给测试结果补误伤说明。
```

## 示例 6：平台整合最终评审

### 适用场景

- platform boundary
- end-to-end drill
- graduation defense

### 示例写法

```md
## Review Note

### 结论

当前材料已经足以说明 Atlas 不是零散功能集合，但“平台边界”还需要用更少、更强的证据来表达，否则 reviewer 容易只记住模块名，不记住系统故事。

### 主要发现

1. 子系统列表已经齐，但共享 contract 的表达还不够集中。
2. drill、benchmark、operator guide 都存在，但彼此之间的叙事连接还可以再压缩。
3. 下一阶段路线已经有方向，但还需要和现有证据更强地绑定。

### 证据

- platform month 已经要求 end-to-end drill、architecture pack、review packet 一起出现。
- 年度复盘周强调毕业故事必须同时连接目标、设计、证据和下一步。

### 下一步

- 把共享 contract 单独抽成一页平台边界图。
- 把最强 drill、最强 benchmark 和最强 operator evidence 放到同一组答辩材料里。
- 用现有弱点直接反推 next roadmap，而不是只写兴趣方向。
```

## 一份像样的 review note 最少要有哪几段

如果你不确定怎么写，最少保留这四段：

1. `结论`：一句话说现在到底站在哪
2. `主要发现`：先列最重要问题，不要先讲细节
3. `证据`：说明你的判断不是感觉
4. `下一步`：指出可以执行的改进动作

## 最后一个提醒

review note 不是心得体会。

它的目标是：

- 帮未来的你或别人快速理解当前判断
- 帮系统留下可以被复查的工程证据

如果你接下来要继续补：

- issue card / evidence entry：
  - 看 [Issue 与 Evidence 示例](issue-and-evidence-examples.md)
- monthly scorecard / review packet：
  - 看 [Scorecard 与 Review Packet 示例](scorecard-and-review-examples.md)
