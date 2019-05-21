package link

import (
	"bytes"
	"fmt"
	"log"
	"text/template"
	"time"

	"auto-visit-blog/config"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

// 博客信息结构体
type BlogInfo struct {
	CfgInfo           config.ConfigInfo
	Bpagelinktemplate string
	VisitLink         []string
	VisitLinkLen      int
}

// 查询单页中请求到的文章ID
func (blog_info *BlogInfo) CrawlCSDNOnePageLinkList(link string) (final bool) {
	c := colly.NewCollector()

	var link_list []string

	c.UserAgent = "User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.75 Safari/537.36"

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("crawl: ", r.URL.String())
	})

	c.OnHTML("div.article-list", func(e *colly.HTMLElement) {
		e.DOM.Find("div.article-item-box").Each(func(i int, s *goquery.Selection) {
			if tmp_id, ok := s.Attr("data-articleid"); ok {
				link_list = append(link_list, "https://blog.csdn.net/"+blog_info.CfgInfo.Uname+"/article/details/"+tmp_id)
			}
		})
	})

	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		RandomDelay: 1 * time.Second,
	})

	c.OnScraped(func(_ *colly.Response) {
		if list_len := len(link_list); list_len > 0 {
			final = false
			blog_info.VisitLink = append(blog_info.VisitLink, link_list...)
			blog_info.VisitLinkLen += list_len
		} else {
			final = true
		}
	})

	err := c.Visit(link)
	if err != nil {
		// colly 返回的错误信息
		log.Println("Visit error: ", err)
		final = true
	}

	return final
}

// 查询所有博客链接列表
func (blog_info *BlogInfo) CrawlAllBlogLinkList() {
	switch blog_info.CfgInfo.Bname {
	case "csdn":
		blog_info.Bpagelinktemplate = "https://blog.csdn.net/{{.BlogName}}/article/list/{{.BlogPageNum}}"
		log.Println("match csdn.")
		// 循环得到所有链接
		// 由于请求URL时，服务器返回非所有链接，所以此处需要循环请求
		for i := 1; ; i++ {
			if final := blog_info.CrawlCSDNOnePageLinkList(blog_info.getBlogVisitLink(i)); final {
				return
			}
		}
	default:
		log.Println("Unmatched...")
	}

	return
}

// 更新所有博文链接
func (blog_info *BlogInfo) UpdateAllBlogLinkList() {
	// 清空所有链接
	blog_info.VisitLinkLen = 0
	blog_info.VisitLink = blog_info.VisitLink[0:0]

	// 重新爬取所有访问链接
	blog_info.CrawlAllBlogLinkList()
}

// 根据模版拼接博客列表页的访问链接
func (blog_info *BlogInfo) getBlogVisitLink(page int) string {
	var link bytes.Buffer

	vl := struct {
		BlogName    string
		BlogPageNum int
	}{
		blog_info.CfgInfo.Uname,
		page,
	}

	tmpl, _ := template.New("link").Parse(blog_info.Bpagelinktemplate)
	tmpl.Execute(&link, vl)

	return link.String()
}
