# cloudstack-server-agent

[![GitHub](https://camo.githubusercontent.com/1e4b67c04239c3dd635207b75c841e801dcf80d5d4a4bdc815672b4c50549ea5/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f2d4769746875622d3030303f7374796c653d666c6174266c6f676f3d476974687562266c6f676f436f6c6f723d7768697465)](https://github.com/cloudstack-tech/cloudstack-server-agent)
![GitHub go.mod Go version (branch)](https://img.shields.io/github/go-mod/go-version/cloudstack-tech/cloudstack-server-agent/master)
![GitHub License](https://img.shields.io/github/license/cloudstack-tech/cloudstack-server-agent)
![GitHub Issues or Pull Requests](https://img.shields.io/github/issues/cloudstack-tech/cloudstack-server-agent)
![GitHub Issues or Pull Requests](https://img.shields.io/github/issues-pr/cloudstack-tech/cloudstack-server-agent)
![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/cloudstack-tech/cloudstack-server-agent/go.yml?color=%2325c0a1)

CloudStack Server Agent 是由云栈互联（CloudStack）开发的服务器被控端服务，为云管理平台提供 Windows 服务器的远程管理、监控和日志采集功能。

该服务支持多种通信方式，包括 gRPC、HTTP 和 WebSocket，确保在不同网络环境下的稳定通信。服务具备以下核心特性：

- 自动数据上报：定时收集并上报系统性能指标、日志事件等数据
- 按需数据获取：支持控制端实时请求获取系统信息和执行管理操作
- 断线重连：自动检测网络状态并进行重连，确保服务可用性
- 数据本地缓存：在网络异常时缓存待上报数据，恢复后自动补传
- 低资源占用：优化的数据采集和处理机制，最小化对被控服务器的影响
- 插件化设计：支持动态加载自定义监控和管理模块

## 通信特性

- gRPC：基于 HTTP/2 的高性能 RPC 通信
- HTTP：标准 RESTful API 接口
- WebSocket：实时双向数据传输

## 功能特性

### 自动上报

- [ ] 运行状态

### 远程调用 RPC

- [ ] 命令行执行
  - [ ] CMD 命令执行
  - [ ] PowerShell 脚本执行
  - [ ] WMI 查询接口
  - [ ] 超时控制和错误处理

### 日志采集

- [ ] Windows 事件日志
  - [ ] 系统日志
  - [ ] 应用程序日志
  - [ ] 安全日志
  - [ ] 自定义日志通道

### 系统监控

- [x] CPU 监控
  - [x] CPU 使用率
  - [x] CPU 频率
  - [x] CPU 核心数
  - [x] CPU 核心数据(型号)
- [ ] 内存监控
  - [ ] 内存使用率
  - [ ] 可用物理内存
  <!-- - [ ] 页面交换率
  - [ ] 页面文件使用情况 -->
- [ ] 磁盘监控
  - [ ] 磁盘使用率
  - [ ] 磁盘 IO 性能
  - [ ] 磁盘读写延迟
- [ ] 磁盘分区监控
  - [ ] 分区使用率
- [ ] 网络监控
  - [ ] 网络接口流量
  - [ ] 网络连接状态
  - [ ] TCP/IP 性能指标
  - [ ] 网络延迟和丢包率
  - [ ] 网络接口配置信息
- [ ] 进程监控
  - [ ] 进程 CPU 使用率
  - [ ] 进程内存使用
  - [ ] 进程句柄数
  - [ ] 进程线程数
  - [ ] 进程启动参数
  - [ ] 进程依赖项
- [ ] 系统
  - [ ] 系统运行时长
  - [ ] 系统版本信息
  - [ ] 已安装更新
  - [ ] 系统服务状态
  - [ ] 开机启动项

### 系统管理

- [ ] 服务管理
  - [ ] 服务启停控制
  - [ ] 服务状态查询
  - [ ] 服务配置修改
- [ ] 进程管理
  - [ ] 进程启停
  - [ ] 进程优先级调整
  - [ ] 进程内存 dump
- [ ] 系统操作
  - [ ] 系统重启/关机
  - [ ] 远程桌面会话管理
  - [ ] 系统时间同步
- [ ] 文件管理
  - [ ] 文件上传下载
  - [ ] 文件删除/重命名
  - [ ] 文件权限管理
  - [ ] 文件压缩/解压

### 自我管理

- [ ] 自动更新
  - [ ] 版本检测
  - [ ] 增量更新
  - [ ] 回滚机制
- [ ] 健康检查
  - [ ] 服务状态监控
  - [ ] 资源占用控制
  - [ ] 故障自恢复
- [ ] 日志管理
  - [ ] 本地日志轮转
  - [ ] 日志级别动态调整
  - [ ] 诊断信息收集

## 快速开始

### 系统要求

- Windows Server 2012 R2 及以上版本

### 安装部署

```bash
# 待补充
```

### 配置说明

```yaml
# 待补充
```

## 开发说明

### 构建环境

- Go 1.24.5 或更高版本
- Windows 开发环境

### 编译构建

```bash
# 待补充
```

## 许可证

本项目采用 [LICENSE](LICENSE) 许可证。
