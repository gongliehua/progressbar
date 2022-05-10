package main

import (
	"github.com/liukaishui/progressbar"
	"time"
)

func main() {
	// 定义一个ProgressBar
	progressBar := progressbar.NewProgressBar(100)
	// 显示ProgressBar
	progressBar.Show()

	for i := 1; i <= 100; i++ {
		// 增加进度条
		progressBar.Add(1)
		// 这里不用延时，这里延时仅是为了模拟环境，方便查看进度条效果
		time.Sleep(10 * time.Millisecond)
	}

	// 实际使用中，建议延迟1/4秒后退出主程序
	// 如果运算速度太快，有可能还没来得及 fmt.Printf 100% 主程序就退出了
	// 如果你不在乎或下方有执行其他代码，那么这里也就没必要写延时了
	time.Sleep(250 * time.Millisecond)

	// ^_^ 祝你好运
}
