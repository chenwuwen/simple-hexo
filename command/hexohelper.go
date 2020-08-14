package command

import (
	"fmt"
	"os/exec"
	"time"
)

//https://blog.csdn.net/weixin_33896726/article/details/94249815

/**
执行预览命令
*/
func Server(port int) (string, bool) {

	fmt.Println("启动Hexo预览功能")
	//cmd := exec.Command(`hexo  server`)
	//https://www.jianshu.com/p/3c968a390cde
	cmd := exec.Command("cmd.exe", "/c", "hexo server")
	var out []byte
	var err error
	//使用协程执行命令
	go execute(cmd, &out, &err)
	//睡眠5秒
	time.Sleep(time.Duration(5) * time.Second)

	ret := string(out)
	fmt.Println(ret)
	if err != nil {
		return ret, false
	}
	//返回的out是字节切片类型
	return ret, true

}

/**
部署
*/
func Deploy() (string, bool) {
	cmd := exec.Command("cmd.exe", "/c", "start hexo g && hexo d")
	out, err := cmd.CombinedOutput()
	utf8Out := ConvertEncode(out)
	ret := string(utf8Out)
	fmt.Println(ret)
	if err != nil {
		return ret, false

	}
	//返回的out是字节切片类型
	return ret, true
}

/**
部署
*/
func Clean() (string, bool) {
	cmd := exec.Command("cmd.exe", "/c", "hexo clean")
	out, err := cmd.CombinedOutput()
	utf8Out := ConvertEncode(out)
	ret := string(utf8Out)
	fmt.Println(ret)
	if err != nil {
		return ret, false

	}
	//返回的out是字节切片类型
	return ret, true
}

/**
  安装依赖
*/
func InstallDepend() (string, bool) {
	cmd := exec.Command("npm install")
	out, err := cmd.CombinedOutput()
	utf8Out := ConvertEncode(out)
	ret := string(utf8Out)
	if err != nil {
		return ret, false
	}
	fmt.Println(string(out))
	//返回的out是字节切片类型
	return ret, true
}

/**
执行命令,得到输出。
*/
func execute(cmd *exec.Cmd, out *[]byte, err *error) {
	//如果执行的命令是监听类型,因此如果执行命令成功,那么将是阻塞的,如果是阻塞的话,那么下面的代码将不会执行,调用处应做相应判断
	//如果执行的命令肯定不会是阻塞的话,则不应该调用这个execute方法
	o, e := cmd.CombinedOutput()
	//GBK转UTF-8
	*out = ConvertEncode(o)
	*err = e
	fmt.Println("协程执行完毕")
}
