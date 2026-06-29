Decision:
统一 Configuration 入口。

Reason:
采用三层配置模型：Default -> config.yaml -> Environment。
优先级：ENV > config.yaml > Default。

Future:
TOML
