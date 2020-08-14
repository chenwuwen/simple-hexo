> 默认Windows 构建(go build)时,使用命令符启动程序的,也就是
>先打开命令符窗口再启动程序,命令符窗口不会消失,因此如果需要直接启动程序,则使用命令 go build -ldflags -H=windowsgui