# go-unzip

#### 介绍
使用go语言，实现zip包的解压
解决了各种平台下解压包的乱码问题

#### 软件架构
软件架构说明


#### 安装教程

1.  `go get -u github.com/mikeside/go-unzip`
2.  `go get -u gitee.com/mikehub/go-unzip`

#### Examples

    package main
    
    import (
        "github.com/artdarek/go-unzip"
        "fmt"
    )
    
    func main() {
        uz := unzip.New("file.zip", "directory/")
        err := uz.Extract()
        if err != nil {
            fmt.Println(err)
        }
    }


#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request


#### 特技

1.  使用 Readme\_XXX.md 来支持不同的语言，例如 Readme\_en.md, Readme\_zh.md
2.  Gitee 官方博客 [blog.gitee.com](https://blog.gitee.com)
3.  你可以 [https://gitee.com/explore](https://gitee.com/explore) 这个地址来了解 Gitee 上的优秀开源项目
4.  [GVP](https://gitee.com/gvp) 全称是 Gitee 最有价值开源项目，是综合评定出的优秀开源项目
5.  Gitee 官方提供的使用手册 [https://gitee.com/help](https://gitee.com/help)
6.  Gitee 封面人物是一档用来展示 Gitee 会员风采的栏目 [https://gitee.com/gitee-stars/](https://gitee.com/gitee-stars/)
