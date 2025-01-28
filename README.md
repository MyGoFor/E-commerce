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

## 启动命令
` `