
# auto-visit-blog 

> 利用 go colly 自动刷博客访问量（目前只做了 CSDN 部分）
  
## 策略
- 配置文件中设置用户ID、间隔、延时、随机访问率
- 8小时自动更新一次博文访问链接列表
- 20~25分钟（配置文件中设置随机数）自动遍历访问一次随机选择的博文链接

## 注意
- 由于采用正常 HTTP GET 方法访问，所以需要设置一段合理的访问间隔，防止被网站检测

## Mac下交叉编译
```shell
# linux 版本
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
# windows 版本
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
```
> 注意：fish shell 下会失败，需要切换到 bash 环境下；

## 版本
> 1.0.0