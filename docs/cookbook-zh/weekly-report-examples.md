# 周报示例

最后更新：2026-06-30

## 这份文档解决什么问题

学习者通常知道：

- 每周要有交付
- 每周要做复盘
- 每周要留证据

但到了真正写周报时，最常见的问题是：

- 只会把本周做过的事重新列一遍
- 写不出“最强结果”和“最大风险”
- 看不出一周工作的工程价值

这份文档就是把周报写法讲清楚。

先对照真实模板：

- [Program Weekly Report](../operations/program-weekly-report.md)
- [Week 01 Report](../reports/weekly/week-01-report.md)
- [Week 35 Report](../reports/weekly/week-35-report.md)

## 周报不是流水账

周报真正要回答的，是下面四个问题：

1. 这周原计划做什么，实际做到了什么
2. 这周留下了哪些证据
3. 这周最大的风险和最弱点是什么
4. 下周最值得优先继续什么

如果一份周报只记录“做了哪些动作”，而没有回答这四个问题，它通常就不够强。

## 示例 1：第一周周报怎么写

### 场景

- 主题：学习系统搭建与服务骨架 RFC
- 特点：代码量可能不大，但边界定义和学习系统建立非常关键

### 一份更像样的中文周报

```md
## Weekly Report

- week: Week 01
- theme: 学习系统搭建与服务骨架 RFC
- monthly objective: 建立 Atlas 的后端服务地基
- Atlas milestone touched: backend foundation and service skeleton

## Planned vs Done

- planned issue cards: ATLAS-M01-01, ATLAS-M01-02
- done issue cards: ATLAS-M01-01（RFC 草稿）、ATLAS-M01-02（入口层草图）
- moved to review: bootstrap boundary note
- blocked: 无

## Evidence Produced

- code or docs: 服务骨架 RFC、包结构草图、学习日志入口
- tests or validation: 启动主路径 smoke 想定与验证口径说明
- traces, drills, or benchmark: 本周不做 benchmark；已记录最小启动验证计划
- review notes: 已完成一版“为什么 main 不应该承担初始化”的 review note
- memory updates: 已沉淀第一周边界判断

## Risks

- top blocker: runtime 仍然过薄，后续依赖接入位置需要更明确说明
- aging blocked issue: 无
- design risk: serve 入口未来可能重新吸回过多生命周期责任
- delivery risk: 如果不尽快补 smoke path，后续 Demo 会依赖作者机器状态

## Review Summary

- strongest result: 入口层 / bootstrap / runtime 三层边界已经开始成形
- weakest result: 证据仍偏文档化，运行验证还不够厚
- code-health note: main 保持很薄，这是正确方向
- reliability note: 启动路径 contract 已出现，但 failure 语义仍需继续补
- next-week focus: 打通 CLI / HTTP 外壳与最小健康检查路径

## Weekly Metrics Snapshot

- completion rate: 80%
- validation coverage: 低到中，已有口径但运行证据偏薄
- evidence completeness: 中
- blocked aging: 0
```

### 这份写法为什么更强

因为它不是在写“我很忙”，而是在写：

- 本周系统边界发生了什么变化
- 当前最弱环节是什么
- 下周应该优先收哪条债

## 示例 2：治理周周报怎么写

### 场景

- 主题：Governance RFC 与 threat model
- 特点：这周很多产物可能是 RFC、threat model、policy 说明和 adversarial evidence，而不是功能代码

### 一份更像样的中文周报

```md
## Weekly Report

- week: Week 35
- theme: Governance RFC 与 threat model
- monthly objective: 建立 policy controls、threat modeling、misuse defenses 与 auditable actions
- Atlas milestone touched: governance, safety, and audit

## Planned vs Done

- planned issue cards: ATLAS-M09-01, ATLAS-M09-02
- done issue cards: threat model 草稿、governance RFC 一版
- moved to review: action classification 设计说明
- blocked: 无

## Evidence Produced

- code or docs: governance RFC、risk taxonomy、初版控制边界图
- tests or validation: adversarial scenario 列表、注入与越权路径测试计划
- traces, drills, or benchmark: 本周无性能 benchmark；已列出治理证据链和审计留痕要求
- review notes: 已完成一版 policy enforcement point review
- memory updates: 已沉淀“治理不能后补”的稳定结论

- Risks
- top blocker: policy 仍偏文档化，execution path 的 enforce point 还不够强
- aging blocked issue: 无
- design risk: threat model 如果不覆盖真实高风险动作，后续 control 都会失焦
- delivery risk: 只写治理文档、不做 adversarial evidence 会让整月结论失真

## Review Summary

- strongest result: 风险已经开始从“抽象安全词”变成具体动作分类
- weakest result: 审计仍偏解释“发生了什么”，不足以解释“为什么被允许”
- code-health note: 控制面开始收束，但还未完全进入执行路径
- reliability note: governance 结论已有方向，证据仍需加厚
- next-week focus: 把 policy / audit / misuse tests 接进真实执行路径

## Weekly Metrics Snapshot

- completion rate: 75%
- validation coverage: 低到中，已有 adversarial plan，执行证据不足
- evidence completeness: 中
- blocked aging: 0
```

### 治理周最容易写错的地方

1. 把文档工作写成“没有工程价值”的附属项
2. 没有把风险和控制动作配对
3. 没有诚实写出 execution path 仍然偏弱

## 示例 3：年底周报怎么写

### 场景

- 主题：年度复盘与下一阶段路线图
- 特点：重点不是新功能，而是压缩证据和形成稳定判断

### 一份更像样的中文周报

```md
## Weekly Report

- week: Week 52
- theme: 年度复盘与下一阶段路线图
- monthly objective: 完成平台级答辩与下一阶段规划
- Atlas milestone touched: platform integration and graduation defense

## Planned vs Done

- planned issue cards: ATLAS-M12-05
- done issue cards: 年度复盘草稿、下一阶段路线图、毕业证据包整理
- moved to review: graduation defense materials
- blocked: 无

## Evidence Produced

- code or docs: 年度复盘、平台边界总结、路线图草稿
- tests or validation: 本周不新增主验证逻辑，重点整合已有 drill / benchmark / review evidence
- traces, drills, or benchmark: 已汇总最终 benchmark 和 end-to-end drill
- review notes: 已完成一版平台答辩 review note
- memory updates: 已沉淀全年最稳定的架构判断

## Risks

- top blocker: 证据很多，但平台故事如果不压缩，reviewer 仍可能只记住零散模块
- aging blocked issue: 无
- design risk: 下一阶段路线若不直接基于现有弱点，会变成空泛规划
- delivery risk: 只讲成绩不讲 trade-off，会削弱毕业答辩说服力

## Review Summary

- strongest result: 端到端 drill、benchmark、operator guide 开始组成平台级证据包
- weakest result: 共享 contract 的表达仍可更集中
- code-health note: 最终判断更多依赖证据组织，而不是新增代码
- reliability note: 关键 failure mode 已被识别，但保护强度仍不均匀
- next-week focus: 无；下一阶段从平台边界、残余风险和证据最弱点反推

## Weekly Metrics Snapshot

- completion rate: 100%
- validation coverage: 以既有 benchmark / drill / review evidence 为主
- evidence completeness: 高
- blocked aging: 0
```

## 周报里最常见的坏味道

### 1. strongest result 写成“做了很多”

这不算结果。

应该写：

- 哪条边界更清楚了
- 哪组证据更完整了
- 哪个 failure mode 被真正收住了

### 2. biggest risk 写成“时间有点紧”

这不算工程风险。

更好的写法应该指向：

- 哪条主路径仍然脆弱
- 哪组证据仍然不够
- 哪个设计判断还没有站稳

### 3. next-week focus 写成“继续推进”

这没有信息量。

应该写：

- 下一周最该优先推进哪一条边界
- 为什么是它，不是别的

## 一份周报最少要能支持什么

至少要支持：

1. 别人快速接手本周上下文
2. 月度 scorecard 能从这里抽证据
3. review packet 能从这里抽 strongest / weakest / next focus

## 最后一句

好的周报不是为了证明你很努力。

而是为了让下一周、下个月和后续 reviewer 都能直接站在你已经留下的判断上继续前进。
