package ui

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
	"github.com/atotto/clipboard"
	c "simple-hexo/const"
)

func ServerResult(result string, state bool, parent fyne.Window) {
	if state {
		dialog.ShowInformation(c.TITLE_SUCCESS, "启动预览成功", parent)
		ReviewButton.Disable()
	} else {
		content := widget.NewLabel(result)
		scroll := widget.NewScrollContainer(content)
		scroll.SetMinSize(fyne.NewSize(400, 300))
		fmt.Printf("启动预览出错 %#v \n", result)
		dialog.NewCustom(c.TITLE_FAIL, "关闭", scroll, parent).SetOnClosed(func() {
			//复制内容到剪切板
			clipboard.WriteAll(result)
		})

	}
}

func DeployResult(result string, state bool, parent fyne.Window) {
	if state {
		dialog.ShowInformation(c.TITLE_SUCCESS, "Deploy Complete", parent)
	} else {
		fmt.Printf("启动预览出错 %#v", result)
	}
}

func CleanResult(result string, state bool, parent fyne.Window) {
	if state {
		dialog.ShowInformation(c.TITLE_SUCCESS, "Clean Complete", parent)
	} else {
		fmt.Printf("启动预览出错 %#v", result)
	}
}

func InstallResult(result string, state bool, parent fyne.Window) {
	if state {
		dialog.ShowInformation(c.TITLE_SUCCESS, "Install Complete", parent)
	} else {
		content := widget.NewLabel(result)
		scroll := widget.NewScrollContainer(content)
		scroll.SetMinSize(fyne.NewSize(400, 300))
		fmt.Printf("安装依赖出错 %#v \n", result)
		dialog.NewCustom(c.TITLE_FAIL, "关闭", scroll, parent).SetOnClosed(func() {
			//复制内容到剪切板
			clipboard.WriteAll(result)
		})
	}
}

func SetWorkDir(parent fyne.Window) {

	//fDialog := dialog.NewFileOpen(func(closer fyne.URIReadCloser, err error) {
	//	path := closer.URI().String()
	//	//截取字符串,从第几个字符开始截取
	//	path = path[7:]
	//	fmt.Println("新选择的路径: ", path)
	//	//更改Label中的值
	//	CurDirLabel.SetText(path)
	//	//更改工作目录
	//	os.Chdir(path)
	//}, parent)
	//
	//
	//f := []string{""}
	//u := storage.NewURI("")
	//fmt.Printf("type:%s,ext: %s,sche: %s",u.MimeType(),u.Extension(),u.Scheme())
	//filter := storage.ExtensionFileFilter{Extensions: f}
	//filter.Matches(u)
	//fDialog.SetFilter(&filter)

	dialog.ShowFileOpen(func(dir fyne.ListableURI, err error) {
		if err != nil {
			dialog.ShowError(err, parent)
			return
		}
		//loadDir(dir)
		fmt.Println("")
	}, parent)
	//fDialog.Show()
}
