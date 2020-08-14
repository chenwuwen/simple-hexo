package main

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/flopp/go-findfont"
	"os"
	"simple-hexo/command"
	c "simple-hexo/const"
	"simple-hexo/ui"
	"strings"
)

var layoutContainer *fyne.Container

func init() {
	//设置中文字体
	//os.Setenv("FYNE_FONT","Alibaba-PuHuiTi-Medium.ttf")
	fontPaths := findfont.List()
	for _, path := range fontPaths {
		//fmt.Println(path)
		//楷体:simkai.ttf
		//黑体:simhei.ttf
		//微软雅黑:msyh.ttc
		if strings.Contains(path, "msyh.ttc") {
			os.Setenv("FYNE_FONT", path)
			break
		}
	}
	fmt.Println("=============")

}

func main() {
	app := app.New()
	//设置应用的主题，默认为 DarkTheme
	app.Settings().SetTheme(theme.LightTheme())
	//os.Setenv("FYNE_FONT","E:\\simple-hexo\\YaHeiConsolasHybrid1.12.ttf")
	//设置程序工作路径,默认在当前文件夹下
	os.Chdir("E:\\snowgrace.github.io")

	//这一步是阻塞的
	mainLayout(app).ShowAndRun()

	fmt.Println("*******应用关闭*******")
	os.Unsetenv("FYNE_FONT")
}

/**
主布局
整体布局分为上下两个布局,上布局为普通布局,下布局为grid布局
*/
func mainLayout(app fyne.App) fyne.Window {
	w := app.NewWindow("Hexo Assistant")
	w.Resize(fyne.NewSize(600, 500))

	//创建一个2列的布局
	gridLayout := layout.NewGridLayoutWithColumns(2)

	//布局中加入组件
	layoutContainer = fyne.NewContainerWithLayout(
		gridLayout,
		ui.CreateReviewButton(w),
		ui.CreateDeployButton(w),
		widget.NewButton(c.BUTTON_NAME_CLEAN, func() {
			ret, state := command.Clean()
			ui.CleanResult(ret, state, w)
		}),
		widget.NewButton(c.BUTTON_NAME_INSTALL, func() {
			ret, state := command.InstallDepend()
			ui.InstallResult(ret, state, w)
		}),
	)

	//设置主窗体内容
	w.SetContent(widget.NewVBox(
		ui.ShowCurDir(),
		widget.NewButton(c.BUTTON_NAME_SETWORKDIR, func() {
			ui.SetWorkDir(w)
		}),
		widget.NewButton("Quit", func() {
			app.Quit()
		}),
		layoutContainer,
	))
	return w
}
