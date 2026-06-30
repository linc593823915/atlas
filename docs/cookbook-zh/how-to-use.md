# 如何使用这本教材

最后更新：2026-06-30

## 核心原则

这本教材不是让你“看完”。

它是让你：

1. 学一块
2. 做一块
3. 评一块
4. 留下证据

也就是说，你的学习必须同时满足：

- 有输入
- 有输出
- 有工程结果
- 有复盘证据

如果你在阅读中经常出现下面两种卡点：

- 词都认识，但不知道它们在书里怎么串起来
- 知道这一周学什么，但不知道回仓库该看哪些文件

先配合使用这两个辅助索引：

- [术语总索引](glossary.md)
- [源码与证据阅读路径图](source-reading-map.md)

如果 day 文档里出现了很多外部阅读提示词，但你不知道“读完到底该提炼什么”，再配合：

- [外部阅读提示解读](external-reading-guides.md)

如果提示词看起来像“仓库内应该一起看哪些线索”，再配合：

- [仓库内阅读线索](internal-reading-cues.md)

如果某个词很短，但你总觉得“好像懂了又说不清”，再配合：

- [原子概念补充](atomic-concepts.md)

如果你已经知道该看哪些文件，但还不知道“怎么读”，再加上：

- [主题级源码讲解](code-walkthroughs.md)
- [代码阅读 Patterns](code-reading-patterns.md)

如果你不是缺定义，而是缺“真实感觉”，再加上这两个材料：

- [专题案例包](casebook.md)
- [失败场景图谱](failure-patterns.md)

如果你做到 review、复盘或排错阶段，直接配合这份材料：

- [评审清单与模式提示](review-checklists.md)
- [Review Note 示例](review-note-examples.md)

如果你已经知道“要交什么”，但还不知道“这些交付物到底怎么写”，再加上：

- [交付物写作指南](artifact-playbook.md)
- [常见交付物词典](deliverable-reference.md)
- [日级常见产物示例](daily-output-examples.md)
- [日级执行动作手册](daily-execution-playbook.md)
- [延伸追问回答指南](reflection-question-guide.md)
- [Issue 与 Evidence 示例](issue-and-evidence-examples.md)
- [Scorecard 与 Review Packet 示例](scorecard-and-review-examples.md)
- [周报示例](weekly-report-examples.md)
- [Memory 示例](memory-note-examples.md)

## 每日学习法

每天按四步走：

### 1. Read

读当天指定的概念、官方文档、源码或现有设计。

目标不是“多读”，而是回答两个问题：

- 今天要解决什么问题？
- 这个问题在真实系统里为什么重要？

### 2. Write

写下你的理解、约束、问题边界、方案比较，或者最小 RFC 草稿。

如果你读了很多却没有写出自己的判断，那通常说明你还没有真正理解。

### 3. Build

把当天的学习落到 Atlas。

哪怕当天只做一件很小的事，也必须尽量让它是一个真实工程动作，比如：

- 定义接口
- 写一个 Tool Schema
- 增加一个测试
- 补一段配置或日志治理

### 4. Output

当天结束时必须留下一个最小产物：

- 一段注释过的代码
- 一份设计说明
- 一条验证记录
- 一个 review note
- 一份 memory 草稿

## 每周学习法

每周都要回答：

- 这周我要完成什么？
- 这周我交付了什么？
- 这周 Atlas 被我改变了什么？

一周结束时，至少应该有：

- 一个可扫描的交付结果
- 一组可解释的技术决策
- 一条面向下周的调整建议

## 每月学习法

每个月都必须形成一个小闭环：

- 目标
- 实现
- 测试/验证
- 文档
- 复盘

如果某个月只积累了想法，没有形成 Atlas 变化，那么这个月不算完成。

## 如何对待难点

遇到以下情况时，不要停留在“我懂了”：

- 概念太多
- 架构太大
- 文档太长
- 不知道从哪开始

处理方法：

1. 只抓最小闭环
2. 只做一条主路径
3. 只写一个最小可验证结果
4. 把复杂问题拆成 Issue

## 如何使用底层治理文档

你不需要每天都读治理文档，但你应该知道它们存在。

当你需要：

- 看标准
- 做复盘
- 检查缺口
- 组织评审

就回到这些索引：

- [Operations Index](../operations/README.md)
- [Governance Index](../governance/README.md)
- [Reporting Records Index](../reports/README.md)

## 如何把“概念”变成“看得懂”

如果某个主题还是觉得抽象，建议固定按这个顺序回读：

1. 先看对应中文周文档，确认本周知识点和交付标准
2. 再看对应英文周文档，确认原始 deliverables 和 source reading
3. 再看 Atlas issue 和周报，理解范围、风险和证据
4. 最后再回真实代码、operations 或 governance 材料

不要一上来直接翻代码。

如果没有问题边界和证据上下文，代码通常只会显得更抽象。

## 一句话判断今天学得对不对

如果今天结束后你能同时回答下面三个问题，方向通常就是对的：

1. 今天我学了什么概念？
2. 今天 Atlas 增加了什么真实能力？
3. 今天留下了什么证据供未来复盘？
