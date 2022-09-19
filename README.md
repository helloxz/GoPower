# http-shutdown
使用golang实现HTTP请求关闭Windows计算机。

**注意：改项目仅支持X64系统**

## 修改配置文件

下载此项目后，请将`config.ini.simple`复制一份并命名为`config.ini`，各参数含义如下：

* `port`：监听端口，默认使用`886`，一般不需要修改
* `RunMode`：运行模式，可选参数有
  * release：生产模式
  * debug：开发调试模式
  * test：测试环境模式
* `key`：请求的key密钥，请修改为一个随机字符串（务必修改）

## 在Windows上安装

首先在需要远程关机的windows操作系统上下载此项目，然后在此目录下执行命令：

```powershell
#安装http-shutdown
.\run.exe install
#启动http-shutdown
.\run.exe install
```

## 在Linux上安装

在需要执行WOL命令的机器上下载此项目，然后进入项目目录执行：

```bash
#添加执行权限
chmod +x ./bin/*
chmod +x http-shutdown
#运行此项目
nohup ./http-shutdown &
```

