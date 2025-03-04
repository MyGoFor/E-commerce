# E-commerce
## 字节寒假青训营后端电商项目

## 技术栈
| technology                                                                   | introduce                   |
|------------------------------------------------------------------------------|-----------------------------|
| cwgo                                                                         | 自动生成代码                      |
| kitex                                                                        | rpc框架                       |
| Hertz                                                                        | http框架                      |
| MySQL                                                                        | 数据库                         |
| Session(redis)                                                               | 保持登陆状态                      |
| Docker                                                                       | 容器化以及部署                     |
| [bootstrap](https://getbootstrap.com/docs/5.3/getting-started/introduction/) | 前端框架                        |
| air                                                                          | 热加载                         |
| nats                                                                         | 消息中间件(结帐成功发送)               |
| casbin(rbac)                                                                 | 权限控制(用户与商家)                 |
| Prometheus                                                                   | 指标                          |
| otel                                                                         | 链路追踪(kitex,hertz,gorm,nats) |
| grafana                                                                      | 数据面板                        |
| logrus/GLP                                                                   | 日志系统                        |
| kind                                                                         | docker内部k8s集群部署             |
| 熔断和fallback                                                                         | for product                 |

## 项目文档
<https://redrock.feishu.cn/docx/MkIOdIvpuorOCTxoYA4cvI0knge>

## 启动命令
不同情况先修改`common/mtl/tracing.go`

1. 本地启动</br>
`make docker` </br>
`make user_run` </br>
`make product_run` </br>
`make cart_run` </br>
`make order_run` </br>
`make payment_run` </br>
`make checkout_run` </br>
`make email_run` </br>
`make casbin_run` </br>
`make eino_run` </br>
`make frontend_run` </br>
`make web` </br>

2. docker容器化部署<br>
`git clone https://github.com/MyGoFor/E-commerce.git`<br>
`make docker`<br>


3. kind部署k8s集群后启动<br>
