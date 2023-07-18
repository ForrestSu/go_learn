## go_learn

学习golang，编写的一些例子


## tools 小工具

- 1 HKD 汇率
> go install github.com/ForrestSu/go_learn/tools/hkd@latest

- 2 🔋电池健康百分比 (MacOS)
> go install github.com/ForrestSu/go_learn/tools/battery@latest

## go module 说明
go.mod 文件在树的根目录中，
有四种指令：`module`，`require`，`replace`，`exclude`。

创建一个go.mod:  
> go mod init

使用当前模块路径，创建一个go.mod:  
> go mod init github.com/ForrestSu/go_learn


Q1 我是否应该提交`go.sum`文件以及`go.mod`文件？

A: 通常，模块的go.sum文件应与go.mod文件一起提交。  
1. go.sum 包含特定模块版本内容的预期密码校验和; 
2. `go mod verify`检查磁盘下载的模块下载在磁盘上的缓存副本是否仍与中的条目匹配; 
3. 请注意，go.sum它不是某些替代性依赖项管理系统中使用的锁定文件。 
 > $ go mod verify  
 > all modules verified

Q2 如果我没有任何依赖关系，还应该添加一个`go.mod`文件吗？

A: 是。这支持在GOPATH之外进行工作，帮助与您选择模块的生态系统进行通信，
此外module，您的指令go.mod还用作代码身份的明确声明。

## 我的 Blog 二维码

![blog](qrcode.png)
