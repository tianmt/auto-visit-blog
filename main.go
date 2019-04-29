package main

import (
	"log"
	"time"

	"auto-visit-blog/link"
	"auto-visit-blog/tools"
	"auto-visit-blog/visit"
)

func main() {
	// 初始化
	var blog_info link.BlogInfo
	// 读取配置文件
	blog_info.CfgInfo.ReadConfigInfo()
	// 爬取用户所有博客链接
	blog_info.CrawlAllBlogLinkList()

	// 访问
	for {
		// 遍历存储链接，发起请求...
		visit.Work(blog_info)
		// 每次发起一波访问的时间间隔为随机生成（此处设置随机+5的范围）
		rand_delay, err := tools.GetRandDurationInt(blog_info.CfgInfo.Loopdelay, blog_info.CfgInfo.Loopdelay+5)
		if err != nil {
			rand_delay = 10
		}
		time.Sleep(rand_delay * time.Minute)
	}

	// 添加定时更新访问列表
	go func(blog_info *link.BlogInfo) {
		log.Println("Now update the blog link list.")
		blog_info.UpdateAllBlogLinkList()
		time.Sleep(time.Duration(8) * time.Hour)
	}(&blog_info)
}
