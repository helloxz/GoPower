# GoPower
使用Golang实现HTTP WOL唤醒和关机功能。

**注意：该项目仅支持`Linux/Windows` X64系统**

## 修改配置文件

下载此项目后，请将`config.ini.simple`复制一份并命名为`config.ini`，各参数含义如下：

* `port`：监听端口，默认使用`886`，一般不需要修改
* `RunMode`：运行模式，可选参数有
  * release：生产模式
  * debug：开发调试模式
  * test：测试环境模式
* `webui`：是否开启WEBUI，通常只需要在其中一个节点开启即可，`on`：开启，`off`：关闭（默认）
* `key`：请求的key密钥，请修改为一个随机字符串（务必修改）

## 在Windows上安装

首先在需要远程关机的windows操作系统上下载此项目，然后在此目录下执行命令：

```powershell
#安装GoPower
.\run.exe install
#启动GoPower
.\run.exe start
```

**注意：下次开机会自动运行此服务**

## 在Linux上安装

在需要执行WOL命令的机器上下载此项目，然后进入项目目录执行：

```bash
#添加执行权限
chmod +x ./bin/*
chmod +x GoPower
#运行此项目
nohup ./GoPower &
```

## 二次开发（编译）

GoPower是前后端分离项目，后端使用Golang，前端使用Vue 3，如果您需要二次开发，可以克隆此项目后再自行编译。

**后端编译：**

```bash
#安装依赖包
go mod tidy
#编译
go build -ldflags -w main.go
```

**前端编译：**

前端源文件位于项目的`front`文件夹下，您需要先安装Node.js环境，然后进入此目录下执行：

```bash
#安装依赖
npm install
#运行
npm run dev
#构建(编译)
npm run build
```

编译后的文件位于：`front/dist`目录
