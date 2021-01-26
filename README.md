# 用户登录

## 代码仓库

```gitignore
git clone ……
```

## 环境配置

Go version >= 1.13

```
export GO111MODULE=on
export GOPROXY=https://goproxy.io
```

进入项目目录，对项目进行配置,配置文件路径：

    app/config/yaml/config.production.yaml

## 服务启动：
```
SERVER_ENV=production go run main.go
```

输出如下表示项目启动成功:
```
INFO[0000] mysql connection and Init success.
INFO[0000] redis connect ping response:PONG
2021/01/25 15:05:26 Listening and serving HTTP on Port: 9060, Pid: 5731
```

## 线上部署
