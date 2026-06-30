# 主题级源码讲解

最后更新：2026-06-30

## 这份讲解解决什么问题

前面的教材已经告诉你：

- 这周学什么
- 应该看哪些代码和证据
- 典型案例和失败模式是什么

但学习者还有一个常见卡点：

文件我已经打开了，可我不知道“该怎么看”。

这份文档就是解决这个问题。

它不追求把所有文件逐行解释，而是回答：

1. 先看哪几个文件
2. 这些文件在系统里各自承担什么职责
3. 读时最该盯的边界、不变量和危险信号是什么

## 讲解一：第一个月的启动主路径

### 建议阅读顺序

1. [cmd/atlas/main.go](../../cmd/atlas/main.go)
2. [internal/app/root.go](../../internal/app/root.go)
3. [internal/app/serve.go](../../internal/app/serve.go)
4. [internal/bootstrap/bootstrap.go](../../internal/bootstrap/bootstrap.go)
5. [internal/config/manager.go](../../internal/config/manager.go)
6. [internal/config/loader.go](../../internal/config/loader.go)
7. [internal/logger/options.go](../../internal/logger/options.go)
8. [internal/runtime/runtime.go](../../internal/runtime/runtime.go)

### 你在看什么

这条链路不是为了展示“框架用法”。

它是在回答第一月最重要的问题：

Atlas 到底是怎样从一个进程入口，变成一个可启动、可配置、可记录日志的服务壳。

### 逐层看法

#### 1. `cmd/atlas/main.go`

这里故意很薄。

重点不是功能多，而是确认：

- 进程退出码统一由 `app.Run()` 决定
- 真正的启动逻辑没有直接长在 `main`

如果你看到未来有大量依赖初始化、配置读取、业务判断直接写回 `main`，那就是边界开始泄漏。

#### 2. `internal/app/root.go`

这里定义 Cobra 根命令。

读它时重点不是 Cobra API，而是：

- CLI 入口是否被组织成统一 command tree
- `Run()` 只是把 command execution 变成进程退出码，而不是顺手做初始化

这是“入口层”和“启动层”分离的第一步。

#### 3. `internal/app/serve.go`

这是第一条真实启动路径。

读这个文件时要盯住三件事：

1. `bootstrap.New(context.Background())` 负责什么
2. `Start()` 和 `Stop()` 的生命周期是否对称
3. 错误处理是只为了打印信息，还是已经形成可复用 contract

这里目前实现还很薄，甚至还没有优雅信号处理。

这不是问题本身，反而是一个很好的学习点：

你能看到“骨架已经在哪里”，以及“哪些能力还没有补进来”。

#### 4. `internal/bootstrap/bootstrap.go`

这是第一月最值得反复看的文件之一。

它把启动主路径压成了三段：

- build config
- init logger
- create runtime

你应该重点观察：

- 为什么 `buildConfigManager()` 在这里而不是在 `main`
- 为什么 logger 配置跟着 config manager 走
- 为什么 runtime 被当作一个整体对象交给 `Bootstrap`

这体现的是“统一生命周期 owner”。

如果以后 db、cache、worker、http server 都要接进来，最自然的位置仍然应该是 bootstrap / runtime 这一层。

#### 5. `internal/config/manager.go`

这层最重要的不是字段有几个。

重点是：

- `Manager` 自己保存 `path`、`lookupEnv` 和当前 `config`
- `Load()` 是唯一读取和替换配置的入口
- 对外只暴露 `Config()`、`App()`、`Logger()` 这类稳定读取方法

项目规则说不能在业务代码里直接访问环境变量。

这个文件就是该规则的“合法入口”：环境变量读取被收回到了 `internal/config` 内部，而不是散落在业务层。

#### 6. `internal/config/loader.go`

这个文件是读配置时最值得精读的代码。

重点看四个动作：

1. 先拿 `defaultConfig()`
2. 再读配置文件
3. 再应用环境变量覆盖
4. 最后验证

这就是中文教材里反复强调的：

`Default -> config file -> environment`

另外要特别注意：

- `resolveConfigPath()` 会向上查找配置文件
- `validateConfig()` 会把配置错误尽早失败

这两个点直接决定“系统能不能可预测启动”。

#### 7. `internal/logger/options.go`

这个文件说明了 logger 抽象到底在保护什么：

- level 解析
- format 解析
- `AddSource: true`
- `slog.SetDefault(...)`

学习时不要只记住“支持 json/text”。

更重要的是：

- logger 是统一初始化的
- caller source 已经被打开
- 以后接 tracing / audit 时，结构化日志才有意义

#### 8. `internal/runtime/runtime.go`

现在它几乎是空的。

这不是浪费文件。

它恰恰说明：

- runtime 这个抽象已经被预留
- 未来真正的 db/cache/http/worker 生命周期可以继续往这里塞

读这种“薄实现”时，训练重点是：

不要只看“现在做了多少”，要看“边界是不是已经留对位置了”。

### 这一组代码最该提的 review 问题

1. `serve.go` 里是否已经承担了过多 bootstrap 责任？
2. `runtime.Runtime` 现在过空，未来第一批能力应该先接什么？
3. config 和 logger 的初始化顺序是否已经成为默认 contract？
4. 哪些错误现在只是 `fmt.Printf`，后续应收回统一 logger / error flow？

## 讲解二：配置层为什么值得精读

### 相关文件

- [internal/config/model.go](../../internal/config/model.go)
- [internal/config/manager.go](../../internal/config/manager.go)
- [internal/config/loader.go](../../internal/config/loader.go)
- [configs/config.yaml](../../configs/config.yaml)
- [internal/config/manager_test.go](../../internal/config/manager_test.go)

### 重点不是“会不会读 YAML”

配置层真正重要的是：

- 统一默认值
- 统一覆盖顺序
- 统一校验
- 统一对外读取方式

你应该重点检查：

1. 新增一个配置项时，默认值、文件值、环境变量值是否都要同步考虑
2. 校验失败时，系统是否能在启动期尽早失败
3. 业务代码有没有绕过 manager，自己去找配置

### 读测试时该看什么

看 [internal/config/manager_test.go](../../internal/config/manager_test.go) 时，不要只看“是否 green”。

要问：

- 测试是不是覆盖了优先级和校验规则
- 有没有覆盖缺配置、坏格式、坏 level 这类错误输入
- 测试是否在保护真正的 contract，而不是实现细节

## 讲解三：logger 抽象在保护什么

### 相关文件

- [internal/logger/logger.go](../../internal/logger/logger.go)
- [internal/logger/options.go](../../internal/logger/options.go)
- [internal/logger/record.go](../../internal/logger/record.go)
- [internal/logger/logger_test.go](../../internal/logger/logger_test.go)

### 最重要的不是打印日志，而是统一语义

logger 抽象真正保护的是：

- 统一 level
- 统一 format
- 统一 source 行为
- 统一默认 logger

学习时要注意区分两层：

- `logger.go` 暴露的是业务可调用的 API
- `options.go` 决定的是默认 logger 的初始化和 handler 行为

### 你该盯住的风险

- 业务层绕过 `internal/logger` 直接写别的日志实现
- logger 初始化顺序和 config 顺序脱节
- source 丢失，后面做排障时只能猜是谁打出来的

## 讲解四：programops 脚本为什么值得后半程反复看

### 相关文件

- [scripts/programops/main.go](../../scripts/programops/main.go)
- [scripts/programops/validate.go](../../scripts/programops/validate.go)
- [scripts/programops/evidence.go](../../scripts/programops/evidence.go)

### 为什么它重要

后半年的很多主题在当前仓库里不是通过一大坨产品代码体现，而是通过“治理代码 + 证据代码”体现。

`programops` 就是这种风格的代表。

### 读法

#### 1. `main.go`

它做的事情很简单：

- parse command
- route 到不同子命令

这说明脚本层也在追求稳定 command surface，而不是随便拼 shell。

#### 2. `validate.go`

这个文件最适合拿来理解“文档不是附件，而是结构化系统输入”。

它在做的事情是：

- 校验 issue card 是否缺 section
- 校验状态值是否合法
- 校验 owner 和 risk level 是否存在

换句话说，它把文档约束变成了可执行规则。

这正是中文教材后半段一直强调的：

- 治理
- 审计
- 发布门禁

它们不只是开会话术，应该尽量变成脚本和结构校验。

#### 3. `evidence.go`

这个文件最值得拿来理解“证据闭环”。

它会按月份统计：

- current progress 缺口
- next action 缺口
- validation checkpoint 缺口
- documentation checkpoint 缺口

这里的重点不是 JSON 输出怎么拼。

而是它说明：

仓库已经把“证据不完整”当成一个可以被扫描出来的问题。

### 这一组脚本最适合服务哪些学习主题

- 第八个月：评估、Tracing 与发布门禁
- 第九个月：治理、安全与审计
- 第十二个月：平台整合与毕业答辩

如果你读到后半程还在问“为什么老是强调证据、review packet、scorecard、review note”，就回来看这一组脚本。

## 最后一个提醒

看代码时不要只追“现在能跑什么”。

真正有价值的问题是：

1. 哪条边界已经被代码固定下来了？
2. 哪个不变量已经被测试或脚本保护了？
3. 哪些地方现在还只是空壳，但边界已经留对了？
