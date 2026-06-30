# Scorecard 与 Review Packet 示例

最后更新：2026-06-30

## 这份文档解决什么问题

如果你已经知道：

- issue card 怎么写
- evidence entry 怎么留
- review note 怎么收口

下一步通常会卡在“月度和季度验收产物到底长什么样”。

这份文档就解决这个问题。

它覆盖三类产物：

1. `Monthly Scorecard`
2. `Monthly Review Packet`
3. `Quarter Replay Pack`

## 示例 1：Monthly Scorecard 应该怎么写

先对照真实文件：

- [Month 03 Scorecard](../reports/monthly/month-03-scorecard.md)
- [Month 10 Scorecard](../reports/monthly/month-10-scorecard.md)

### 这份产物的目标

scorecard 不是再复述一遍本月做了什么。

它的作用是把“本月成果”压缩成一个能快速扫描的判断面：

- 目标达成度
- 交付指标
- Atlas 结果
- 面试/晋升信号
- 下个月是否能继续推进

### 一个更像样的中文写法

```md
## Scorecard

- month: Month 03
- objective: 交付第一个可靠的单 Agent 服务，具备结构化输出、tool orchestration 和最小评估基础
- overall grade: B+
- promotion signal: 已出现平台边界意识，但评估证据仍偏薄

## OKR Summary

- objective score: 8/10
- KR1 score: 8/10
- KR2 score: 7/10
- KR3 score: 6/10
- KR4 score: 8/10

## Delivery Metrics

- planned issues: 5
- done issues: 4
- blocked issues: 0
- completion rate: 80%
- validation coverage: 已有 schema validation 和最小 eval baseline，但 tool failure 面仍偏薄
- documentation freshness: operator note 已更新
- review closure rate: 月内关键 review 均已收口
- evidence completeness: 仍需补一条更强的负路径证据

## Atlas Outcome

- milestone shipped: 单 Agent 主路径可演示
- strongest subsystem improvement: structured output + validator 闭环
- most fragile subsystem: 多 Tool fallback 与质量证据联动

## Interview Readiness

- coding: 能解释主路径，但实现仍偏薄
- design: contract 意识明显提升
- reliability: 已有最小验证，但 failure 面不够厚
- behavioral: 能解释取舍，但还需更强 evidence-first 表达
- documentation: 已出现 operator 视角

## Final Judgment

- strongest evidence: schema + validator + eval baseline 形成第一条可复查质量链
- biggest risk: 质量判断仍容易受 spot check 影响
- next-month carryover: 扩展 tool routing 与 fallback 证据
- mentor recommendation: 可以继续推进，但必须加厚 failure / eval 证据
```

### 写 scorecard 时最容易犯的错

1. 只有分数，没有解释
2. 只有主观判断，没有回到 linked weeks / evidence
3. strongest evidence 和 biggest risk 写成空话

## 示例 2：Monthly Review Packet 应该怎么写

先对照真实文件：

- [Month 03 Review Packet](../reviews/archive/month-03-review.md)
- [Month 09 Review Packet](../reviews/archive/month-09-review.md)
- [Month 12 Review Packet](../reviews/archive/month-12-review.md)

### 这份产物的目标

review packet 不是“更长的 scorecard”。

它的作用是帮助 mentor / reviewer / committee 做结构化判断：

- 本月目标是什么
- 关键证据够不够
- 最强项和最弱项在哪里
- 下个月应该继续、加速还是先补债

### 一个更像样的中文写法

```md
## Monthly Objective

建立 policy controls、threat model、misuse defenses 和可审计高风险动作。

## Review Inputs

- 第 35-39 周交付物
- Month 09 issue set 进展
- adversarial tests、audit evidence、review notes
- self-review note

## Monthly OKR Grading

- Objective score: 7/10
- Key Result 1 score: 8/10
- Key Result 2 score: 7/10
- Key Result 3 score: 6/10
- Key Result 4 score: 7/10
- Overall monthly grade: B

## Evidence Checklist

- roadmap artifacts exist: 是
- Atlas backlog issues moved meaningfully: 是
- technical decision notes exist: 是
- tests or validation evidence exist: 有，但 benign/malicious 区分还不够
- docs or runbooks were updated: governance docs 已更新

## Mentor Notes

- strongest result: threat model 已经开始真实约束 policy boundary
- weakest area: audit 仍偏“发生了什么”，不够解释“为什么被允许”
- code health note: 控制面正在集中，但执行点还需收口
- reliability note: adversarial evidence 有起点，但还不够系统
- next-month development focus: 强化 enforcement point 与误伤证据

## Mock Interview Summary

- coding signal: 中
- design signal: 强
- reliability signal: 中偏弱
- behavioral signal: 能解释风险，但证据表达还可更精炼
- overall recommendation: 可以继续下一阶段，但治理控制仍需更可执行

## Promotion Signal

- current level signal: 已有系统性治理意识
- next-level blockers: control 到 execution path 的落地仍不够强
- evidence that changed the rating: threat model 与 adversarial tests 已形成第一组可信证据
```

### 写 review packet 时最容易犯的错

1. 把 strongest / weakest 写成情绪判断
2. Mock interview summary 只写“不错”“一般”
3. reviewer 看完之后仍然不知道 next-month focus 是什么

## 示例 3：Quarter Replay Pack 应该怎么写

先对照真实文件：

- [Quarter 4 Replay Pack](../reviews/archive/quarter-4-replay.md)

### 这份产物的目标

季度 replay 的作用不是再做一次月报。

它是在回答：

- 这个季度 Atlas 到底哪一点真正抬高了工程条线
- 哪个决策最体现架构判断
- 哪个 failure mode 仍然没被保护好
- 下一季度应该继续、加速还是先修债

### 一个更像样的中文写法

```md
## Quarter Theme

生产部署、性能证据与平台整合。

## Replay Questions

- What changed in Atlas this quarter that materially raised the engineering bar?
  生产式部署 contract、baseline / optimization 证据，以及平台级答辩材料开始形成闭环。
- Which decision best showed technical judgment?
  先冻结 benchmark contract，再做优化，而不是直接追数字。
- Which failure mode is still under-protected?
  跨子系统 contract drift 仍然是平台整合里的首要风险。
- What evidence proves the quarter was more than documentation work?
  end-to-end drills、Month 11 benchmark evidence、Month 12 review packet。

## Quarter Evidence Pack

- best RFC: platform integration RFC
- best implementation milestone: 配置、gateway、MCP 主路径被整合到同一平台边界中
- strongest test or validation artifact: end-to-end drills + benchmark regression checks
- strongest operations or reliability artifact: alert / rollback / operator guide evidence
- strongest documentation or review artifact: Month 12 review packet + graduation defense materials

## Calibration Summary

- results rating: B+
- technical judgment rating: A-
- code health rating: B
- reliability rating: B
- communication rating: A-
- learning rating: A
- quarter overall: B+

## Replay Decision

- continue as planned: 是
- accelerate: 平台边界与证据组织表达
- remediate before continuing: 跨子系统 contract drift 和残余 failure mode
- focus for next quarter: 更强的运行时整合和证据压缩能力
```

### 写季度 replay 时最容易犯的错

1. 只是把三个月内容机械拼起来
2. 只有成果，没有主线
3. 写不出“这个季度真正改变了什么”

## 一套最小验收产物链长什么样

如果你要把一个月真正收口，最少应该形成这条链：

1. issue cards
2. execution logs
3. review notes
4. evidence entries
5. monthly scorecard
6. monthly review packet

如果要把一个季度真正收口，再加：

7. quarter replay pack

## 最后一句

scorecard、review packet 和 quarter replay 的目标都不是“让材料变多”。

它们的目标是让更高层级的判断，仍然能回到具体 artifact，而不是停留在印象里。
