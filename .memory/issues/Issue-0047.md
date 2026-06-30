Decision:
在继续把 day 文档提示接成真实链接的过程中，统一加入“嵌套链接压平”修复步骤。

Reason:
批量替换同一批提示词时，如果新旧落点都参与替换，容易生成：
- [[label](inner)](outer)

这会直接影响教材可读性。

因此链接类批量替换之后，必须立即做：
- nested link 扫描
- flatten 修复
- bad link 校验

Future:
继续在每轮批量链接替换后执行同样修复
减少 day 文档中的重复提示词
