Decision:
中文 cookbook 将 day 文档 `今日阅读` 中剩余的反引号提示全部接到稳定中文落点，并补强现有支撑文档的长尾条目。

Reason:
此前 day 层仍有大量裸提示，例如：
- `一份 drill 说明`
- `多 Agent 测试集`
- `Tool Gateway 的职责边界`
- `健康检查配置`
- `部署清单`
- `一组 adversarial case`

这些提示如果停留在纯文本，学习者仍然需要猜“最低应该写什么、读什么、怎么验收”。本轮把它们归并到现有文档：
- `daily-output-examples.md`
- `deliverable-reference.md`
- `internal-reading-cues.md`
- `external-reading-guides.md`

并将 `今日阅读` 中未链接反引号项从 256 个唯一项逐步收口到 0。

Future:
后续继续从“有没有链接”转向“链接落点是否足够教学化”
优先抽查 month/week/day 的实际学习链路是否连贯
继续保持 bad_links、nested_link_files、missing_explicit_anchors 三项为 0
