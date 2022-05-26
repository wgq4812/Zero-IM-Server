# Path-IM-Server
基于 [Open-IM-Server](https://github.com/OpenIMSDK/Open-IM-Server) 实现的 IM 服务 

## 修改部分
### 服务注册发现
> 使用go-zero微服务框架 开发更方便 自带`链路追踪` `服务发现` `p2c服务负载均衡`

> 不依赖`mysql`所有业务逻辑均请求业务rpc服务 

### 新增超级大群功能
> 类似`QQ`群聊的`读扩散`模式  妈妈再也不用担心mongodb写入性能问题了

## 开源组件依赖
- mongodb (离线消息存储)
- kafka (消息队列)
- redis (存储seq)
- ~~etcd~~ (Path-IM-Server 不依赖etcd)
- ~~mysql~~ (Path-IM-Server 不依赖mysql)

## 系统架构图
![system.svg](https://raw.githubusercontent.com/Path-IM/Path-IM-Docs/main/images/20220517/Path-IM-Server-System.svg)

## 业务架构图
![image1.svg](https://raw.githubusercontent.com/Path-IM/Path-IM-Docs/main/images/20220517/Path-IM-Server-Service.svg)

## 业务流程图
![flow.svg](https://raw.githubusercontent.com/Path-IM/Path-IM-Docs/main/images/20220517/Path-IM-Server-Flow.svg)

# Path-IM-Server-Demo
> 使用 `Path-IM-Server` 开发一个 `IM` 应用 
## 开发计划
- [x] 完成 Path-IM-Server 的 TODO [第一天](https://github.com/Path-IM/Path-IM-Server-Demo/tree/main/docs/day01)
- [x] 完成 用户模块 rpc 接口 编写 [第二天](https://github.com/Path-IM/Path-IM-Server-Demo/tree/main/docs/day02)
- [x] 完成 用户关系模块 rpc 接口 编写 [第三天](https://github.com/Path-IM/Path-IM-Server-Demo/tree/main/docs/day03/relation.md)
- [x] 完成 群聊模块 rpc 接口 编写 [第三天](https://github.com/Path-IM/Path-IM-Server-Demo/tree/main/docs/day03/group.md)
- [x] 完成 用户模块 api 接口 编写 [第四天](https://github.com/Path-IM/Path-IM-Server-Demo/tree/main/docs/day04)
- [x] 完成 用户关系模块 api 接口 编写 [第四天](https://github.com/Path-IM/Path-IM-Server-Demo/tree/main/docs/day04)
- [x] 完成 群聊模块 api 接口 编写 [第四天](https://github.com/Path-IM/Path-IM-Server-Demo/tree/main/docs/day04)
- [x] 完成 k8s 部署 [第五天](https://github.com/Path-IM/Path-IM-Server-Demo/tree/main/deploy/k8s)
- [x] 完成 api 文档 编写 [第六天](https://github.com/Path-IM/Path-IM-Server-Demo/tree/main/docs/api.md)
- [x] 完成 消息持久化存储 文档 编写 [第十天](https://github.com/Path-IM/Path-IM-Server-Demo/tree/main/docs/persistent.md)
- [ ] 完成 离线消息定期清理 

# Path-IM-Client-Go
[Path-IM-Client-Go](https://github.com/Path-IM/Path-IM-Client-Go.git)
> 我们计划编写 `dart` sdk；由于时间问题，暂时放出 `golang` 客户端 测试代码；以供参考！

# 其他
## jaeger
![jaeger.png](https://raw.githubusercontent.com/Path-IM/Path-IM-Docs/main/images/20220517/jaeger.png)
