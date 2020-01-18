# 依赖管理

依赖管理的三个阶段GOPATH, GOVENDOR, go mod

## GOPATH

默认在~/go

所有项目都放在gopath, 所有第三方库也在gopath

### 找的顺序：

​      gorood的src

​      gopath的src



## GOVERDOR

项目中有vendor优先在vendor中找

每个项目有自己的vendor目录, 存放第三方库

大量第三方依赖管理工具: glide, dep, go deep ...

### 找的顺序：

​	 本项目的verndor

​      gorood的src

​      gopath的src



## GO MOD

ide 中构建项目选择go modules, proxy中填写: https://goproxy.cn,direct

```
go get -u go.uber.org/zap@v1.11 #拿指定版本
```

```
go mod tidy #清洁下go.sum, 去掉无用的
```

### 增加第三方库方法

​       通过 go get -u go.uber.org/zap， 或 直接在代码里面写import

### 如何迁移到go mod

1、ide下，点击生成go.mod 文件，之后执行命令

```
go build ./...
```

2、命令行

      ```
go mod init test  
go build ./...
删除vendor目录和glide.yaml
      ```



## 总结:

由go命令统一的管理, 用户不必关心目录结构

初始化: go mod init

增加依赖:  go get

更新依赖: go get [@v...]  go mod tidy

将就项目迁移到go mod:  go mod init,  go build ./...



## 目录的整理

包含main函数的文件必须在一个单独的文件夹下，没有文件加下必须只有一个main函数

```
go build ./... #编译当前目录和子目录的所有文件，不会产生可执行文件
go install ./... #安装，产生可执行文件
go  env GOPATH 
在GOPATH的bin目录

```