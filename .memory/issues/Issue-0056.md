Decision:
中文 cookbook 增加 `reflection-question-guide.md`，并把 day 文档里的 `延伸追问` 全部接到可回答的判断框架。

Reason:
`今日阅读` 和日级执行动作已经有稳定落点后，仍然有一类学习空洞：
学习者看到追问时知道问题本身，但不知道应该怎样回答才算工程判断。

本轮新增延伸追问回答指南，要求每个追问至少回答：
1. 我的判断是什么
2. 我依据什么证据
3. 这个判断的风险或边界在哪里
4. 下一步最小动作是什么

并将 393 个日页中的 1179 个延伸追问链接到指南。顺手补齐 `atomic-concepts.md` 的 `os.Getenv` 与 `main()` 显式锚点，收口核心知识点里的裸反引号。

Future:
继续抽查 `今日自问` 与 `每日评分标准` 是否足够区分“读过、做过、真正理解”
继续保持 raw core/reflection code terms、bad_links、nested_link_files、missing_explicit_anchors 为 0
后续重点从“链接覆盖”转向“学习路径是否能独立闭环”
