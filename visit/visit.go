package visit

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"auto-visit-blog/link"
	"auto-visit-blog/tools"
)

var wg sync.WaitGroup

// 访问准备好的博客链接
func Work(blog_info link.BlogInfo) {
	fmt.Println("Working...")
	if blog_info.VisitLinkLen == 0 {
		fmt.Println("visit link list is update now...")
		return
	}

	// 根据比例随机生成需要访问的链接
	lg := listGenerator(blog_info.VisitLinkLen-1, int(float64(blog_info.VisitLinkLen)*blog_info.CfgInfo.Visitrate))

	// 访问链接
	for _, v := range lg {
		go httpGet(blog_info.CfgInfo.Visitdelay, blog_info.VisitLink[v])
		wg.Add(1)
	}

	wg.Wait()
}

// 发起 GET 请求
func httpGet(delay int, url string) {
	rand_delay, err := tools.GetRandDurationInt(0, delay)
	time.Sleep(rand_delay * time.Second)

	resp, err := http.Get(url)
	tools.CheckError(err)

	defer resp.Body.Close()

	wg.Done()
	fmt.Println("visit: ", url, "done.")
}

// 随机生成设置比例的标记号
func listGenerator(max int, cnt int) (lg map[int]int) {
	lg = make(map[int]int)

	for {
		if len(lg) < cnt {
			tmp_num, err := tools.GetRandInt(0, max)
			if err != nil {
				fmt.Println(err)
				return
			}
			lg[tmp_num] = tmp_num
		} else {
			break
		}
	}

	return lg
}
