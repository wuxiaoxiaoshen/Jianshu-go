# Jianshu-go

![](https://img.shields.io/badge/jianshu-v1.0.1-519dd9.svg)
![](https://img.shields.io/badge/language-golang-orange.svg)
[![](https://img.shields.io/badge/weibo-@谢小小路-red.svg)](https://weibo.com/1948244870/profile?topnav=1&wvr=6)
[![](https://img.shields.io/badge/jianshu-@谢小路-F59581.svg)](https://www.jianshu.com/u/58f0817209aa)



本项目尝试解析简书API。



- 编程语言：golang


## Installation

``` 
go get github.com/wuxiaoxiaoshen/jianshu-go
```




主要的接口包括：

- User: 个人主页信息
- Article : 某篇文章的信息
- Home-page: 简书主页的信息
- Home-page-recommend: 简书推荐作者的信息
- Home-page-topic: 简书推荐的专题信息
- Publication： 简书出版信息
- Special-subject: 专题信息


主要包括上面 7 类：


- 支持命令行式

```$xslt
NAME:
   JianShu - An Application of JianShu API.

USAGE:
   jianshu.exe [global options] command [command options] [arguments...]

VERSION:


    ___       ___       ___       ___       ___       ___       ___
   /\  \     /\  \     /\  \     /\__\     /\  \     /\__\     /\__\
  _\:\  \   _\:\  \   /::\  \   /:| _|_   /::\  \   /:/__/_   /:/ _/_
 /\/::\__\ /\/::\__\ /::\:\__\ /::|/\__\ /\:\:\__\ /::\/\__\ /:/_/\__\
 \::/\/__/ \::/\/__/ \/\::/  / \/|::/  / \:\:\/__/ \/\::/  / \:\/:/  /
  \/__/     \:\__\     /:/  /    |:/  /   \::/  /    /:/  /   \::/  /
             \/__/     \/__/     \/__/     \/__/     \/__/     \/__/  v1.0.1




AUTHOR:
   xieWei <wuxiaoxiaoshen@shu.edu.cn>

COMMANDS:
     article      get jianshu article by cli
     home-page    get jianshu home page passage by cli
     recommend    get home page recommend by cli
     topic        get home page topic
     user         get user by cli
     publication  get publication by cli
     help, h      Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version


```


```$xslt

>> jianshu.exe user method --url=https://www.jianshu.com/u/58f0817209aa get-user-id
>> 谢小路

>> jianshu.exe user method --url=https://www.jianshu.com/u/58f0817209aa get-user-gender
>> None

>> jianshu.exe user method --url=https://www.jianshu.com/u/58f0817209aa get-user-link
>> https://www.jianshu.com/u/58f0817209aa

>> jianshu.exe user method --url=https://www.jianshu.com/u/58f0817209aa get-follow-number
>> 9



...


```



## User 

- GetUserID  获取用户ID
- GetUserLink 获取用户主页URL
- GetUserGender 获取用户性别
- GetFollowNumber 获取用户关注数
- GetFollowerNumber 获取用户粉丝数
- GetPassageNumber 获取用户文章书面
- GetWriteNumber 获取用户写的字数
- GetLikeNumber 获取用户得到的喜欢的数目
- GetHomePagePassage 获取用户主页文章信息
- GetPersonalDetail 获取用户个人介绍
- GetTwitterInfo 获取用户微博地址
- GetLikedNotes 获取作者喜欢的文章信息
- GetSubscription 获取作者关注的专题/文集/连载
- GetLatestActice 获取作者最新动态
- GetLatestCommented 获取作者最新评论
- GetHotPassage 获取作者热门文章


![user-one.png](https://upload-images.jianshu.io/upload_images/1818135-b6bdda5bf98d192e.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![user-two.png](https://upload-images.jianshu.io/upload_images/1818135-f68eb4d766af7143.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![user-three.png](https://upload-images.jianshu.io/upload_images/1818135-be2376ad7b93a2c3.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![user-four.png](https://upload-images.jianshu.io/upload_images/1818135-910091c8893091ed.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![user-five.png](https://upload-images.jianshu.io/upload_images/1818135-aaba660c816763ba.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![user-six.png](https://upload-images.jianshu.io/upload_images/1818135-e27de40798147abf.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

#### Tutorial & Usage ==> user

```go
package main

import (
	"fmt"
	"github.com/wuxiaoxiaoshen/jianshu-go"
)

func main() {
    user := jianshu.NewUser("https://www.jianshu.com/u/58f0817209aa", "", "")
    fmt.Println("作者简书ID: ", user.GetUserID())
    fmt.Println("作者简书主页: ", user.GetUserLink())
    fmt.Println("作者性别: ", user.GetUserGender())
    fmt.Println("作者关注数: ", user.GetFollowNumber())
    fmt.Println("作者粉丝数: ", user.GetFollowerNumber())
    fmt.Println("作者文章数: ", user.GetPassageNumber())
    fmt.Println("作者写字数目: ", user.GetWriteNumber())
    fmt.Println("作者收获喜欢数目: ", user.GetLikeNumber())
    fmt.Println("作者主页文章信息: ", user.GetHomepagePassage())
    fmt.Println("作者个人介绍: ", user.GetPersonalDetail())
    fmt.Println("作者微博信息: ", user.GetTwitterInfo())
    fmt.Sprint("作者喜欢的文章信息: ", user.GetLikedNotes())
    fmt.Println("作者关注文集信息: ", user.GetSubscription())
    fmt.Println("作者最新动态: ", user.GetLatestActive())
    fmt.Println("作者最新评论: ", user.GetLatestCommented())
    fmt.Println("作者热门文章: ", user.GetHotPassage())
	
}

```

返回信息：

```

作者简书ID:  谢小路
作者简书主页:  https://www.jianshu.com/u/58f0817209aa
作者性别:  None
作者关注数:  9
作者粉丝数:  562
作者文章数:  83
作者写字数目:  85308
作者收获喜欢数目:  470
作者主页文章信息:  [{2018-03-11T20:12:04+08:00 『简书API : jianshu 基于 golang （1）』 在我眼中，比较崇拜三类人：一类是设计师；一类是作家；一类是程序员。 这三类人都是通过创造、或者改善作品，不断的把世界变的更好。每每看到大师级的作品，总会不禁感叹，人与人的差别... 100 0 2 1 谢小路} {2018-03-10T23:17:32+08:00 『Python 爬虫文集梳理』 过去的几年内，我开始了编程。 过去的一年内，我开始了工作生涯。 我学会的第一个编程技能是『爬虫』，工作后，开始接触Golang。 我开始不断的将编程结合业务， 接触越来越多的... 26 0 2 0 谢小路} {2018-03-10T00:23:23+08:00 『Ansible 上手指南：2』 读一本书最好的时机是什么时候？是你刚买的时候，趁着新鲜劲，先了解这本书，继而马上阅读完这本书。如果错过了最好的时机阅读一本书，那什么时候是合适的时机，是你需要这方面的资料或者... 7 2 1 0 谢小路} {2018-03-06T00:13:44+08:00 『requests-html 源码学习:  1』 大家好，我是谢伟，是一名程序员，熟悉 Pyhton 和 Go。学会的第一个技能是『网络爬虫』。 最近 Python 领域大神  kennethreitz 开源了一个关于网络内... 30 0 3 0 谢小路} {2018-03-05T21:37:43+08:00 『Ansible 上手指南』 Ansible 上手指南 前言 最近在重构一款命令行工具，使用 golang 重新开发，但需要继续维持原有的命令，同时增加新命令。 在重构的过程中，需要对现命令行工具和原命令... 24 0 2 0 谢小路} {2018-02-11T23:03:25+08:00 『Go 语言实现简易爬虫：市值前100数字货币交易信息』 大家好，我是谢伟，一名程序员。之前接触的语言是Python， 编程领域学会的第一个技能是『爬虫』，凭借着爬虫技术先后在两个创业公司从事的是『网络爬虫』这份活。 研究生毕业后，... 80 0 8 0 谢小路} {2018-02-08T22:56:47+08:00 『Beego + Swagger 快速上手』 大纲 Beego 是什么 为什么写这个 如何指导 前几天我写了一个Swagger 上手指南，觉得还是让使用者难以上手。尽管它是一款优秀的API 工具。 但我在编写API 的过... 72 0 0 0 谢小路} {2018-02-03T15:23:06+08:00 『2018年1月知识点合集』 我有一个习惯，就是不断的记录在工作中反复用到的知识点，原本我很喜欢使用印象笔记和有道云笔记，其一是云笔记的同步功能，其二是云笔记的搜索功能，当你输入的笔记多了之后，你才会发现... 100 0 0 0 谢小路} {2018-01-30T21:57:42+08:00 『Swagger 上手』 大纲 问题 RestfulAPI API 动作 请求：Url、Body 返回信息：Status_code、Response 在开发过程中，经常会遇到和其他组件或者服务进行交互... 39 0 1 0 谢小路}]
作者个人介绍:  上海大学2017级研究生毕业.微信公众号：Siwei_Jingjin
作者微博信息:  http://weibo.com/u/1948244870
作者关注文集信息:  [{微信小程序开发于连林520wcf https://www.jianshu.com/c/dfdc2bbd1315 于连林520wcf 收录了660篇文章 3037人关注} {PPT：千页计划谢小路 https://www.jianshu.com/c/d76906de84c5 谢小路 收录了1篇文章 1人关注} {Go 语言学习专栏 谢小路 https://www.jianshu.com/c/0d857ad2e739 谢小路 收录了8篇文章 1人关注} {PPT沈晓马 https://www.jianshu.com/c/9df5862d2a58 沈晓马 收录了1603篇文章 8320人关注} {我的Python自学之路帅灰 https://www.jianshu.com/c/0690e20b7e7d 帅灰 收录了1047篇文章 5183人关注} {强化训练：python谢小路 https://www.jianshu.com/c/03322437b28d 谢小路 收录了19篇文章 20人关注} {爬虫专栏谢小路 https://www.jianshu.com/c/dfcf1390085c 谢小路 收录了262篇文章 950人关注}]
作者最新动态:  [{comment_note 『Ansible 上手指南：2』} {share_note 『简书API : jianshu 基于 golang （1）』} {share_note 『Python 爬虫文集梳理』} {like_user 若与} {share_note 『Ansible 上手指南：2』} {like_user xiaopiu原型设计} {share_note 『requests-html 源码学习:  1』} {share_note 『Ansible 上手指南』} {like_collection } {like_user 刘英滕} {share_note 『Go 语言实现简易爬虫：市值前100数字货币交易信息』} {comment_note 我儿子} {comment_note 修正我的学习} {share_note 『Beego + Swagger 快速上手』} {like_note 【9页幻灯片】阿尔卑斯山} {share_note 『2018年1月知识点合集』} {share_note 『Swagger 上手』} {share_note 『Go 语言学习专栏』-- 第一期} {comment_note 2018『PPT 千页计划』第一期} {share_note 2018『PPT 千页计划』第一期}]
作者最新评论:  [{『Ansible 上手指南：2』 读一本书最好的时机是什么时候？是你刚买的时候，趁着新鲜劲，先了解这本书，继而马上阅读完这本书。如果错过了最好的时机阅读一本书，那什么时候是合适的时机，是你需要这方面的资料或者... https://www.jianshu.com/p/82245bb31bcd} {2018『PPT 千页计划』第一期 2018 年我准备花上一年的业余时间制作 1000 页PPT。 没什么目的。 就是制作 1000 页 PPT。 内容来自一些作业，也来自一些我喜欢作者的原文模仿，也来自一些模... https://www.jianshu.com/p/492875661ab8} {如何成为一名设计师.ppt 我是一名程序员，业余时间玩PPT。 这是根据罗子雄的 TED 演讲而制作的PPT。 很费时间，但习得一门技能从来都不是那么轻易和容易。 https://www.jianshu.com/p/88e861223299} {线性代数：一切为了更好的理解 ![][1][1]: http://latex.codecogs.com/gif.latex?\alpha 线性代数是数学工具掌握它，打开数学的另一扇大门 1：声明 非原创，... https://www.jianshu.com/p/c0e119c76971} {技术文档如何编写？ 关于文档编写的几个思维 近期重新组织了好几篇技术文档，把其中的一些感悟提炼出来。 文档为达到容易理解和操作的程度，对大量的语言重新组织，内容的不同呈现，借助辅助工具等一系列操... https://www.jianshu.com/p/6d66262f6dcd} { 001     |    工作：一周总结 已毕业，已入职，本职工作是云平台部署和管理， PAAS层。 刚入手，系列一周总结，会记录一些工作的学习和思考，目标是：提高业务能力和专业技能，同时成为一名合格的职场人。 为实... https://www.jianshu.com/p/b9eedf83675c} {Python: 一周笔记 主要介绍几个用到的python模块的使用方法。python 含有丰富的内置和第三方库，企图全部掌握并精通那是不可能的。 但当你开发任务需要到的时候，你可以及时的避免重复的一些... https://www.jianshu.com/p/ad450b2126b0} {002 | 工作：职场学习术 这篇文章的灵感来源于暂停玩「王者荣耀」之后，甚至我产生了卸载的冲动，我觉得上班的精神状态和这款手游有着潜移默化的关系。 好，撇开这些不谈，我想谈谈在职场中思考的比较多的两件事... https://www.jianshu.com/p/03c0849eced8} {004 | 工作：月度总结 时间过的很快，距离毕业离校已经过去将近三个月，上班也一个快一个半月，每天都在不断的接触新的东西，新旧知识需要融合，每周公司新人需要进行周报总结，算是一次新旧知识融合，但是输出... https://www.jianshu.com/p/cdb22fa377d1}]
作者热门文章:  [{2017-01-15T09:56:51+08:00 Python:一周笔记 主题 邮件处理 日志模块 pdf处理 md5 mongodb索引和聚合 excel 读写 1. 发送邮件模块 这里指的邮件功能当然不是指的是职场上所谓的邮件，指的是程序运行中... 939 3 30 0 谢小路} {2016-04-29T18:36:33+08:00 专栏：005：Beautiful Soup 的使用 系列爬虫专栏 崇尚的学习思维是：输入，输出平衡，且平衡点不断攀升。 曾经有大神告诫说：没事别瞎写文章；所以，很认真的写的是能力范围内的，看客要是看不懂，不是你的问题，问题在我... 600 2 27 0 谢小路} {2017-08-06T15:23:39+08:00 Python: 一周笔记 主要介绍几个用到的python模块的使用方法。python 含有丰富的内置和第三方库，企图全部掌握并精通那是不可能的。 但当你开发任务需要到的时候，你可以及时的避免重复的一些... 580 3 22 0 谢小路} {2016-04-30T12:23:51+08:00 专栏：006：实战爬取博客 （起什么优雅的名字点击量会增多？） 系列爬虫专栏 崇尚的学习思维是：输入，输出平衡，且平衡点不断攀升。 曾经有大神告诫说：没事别瞎写文章；所以，很认真的写的是能力范围内的，看... 696 1 21 0 谢小路} {2016-05-13T21:15:47+08:00 番外篇：面试总结(1) 初学者 你经历的每一件事都会成为未来的部分，具体看你如何对待了. 0：前言 作为一个初学者，对知识的理解存在着很多的疑惑。同人交流作为学习的方式之一，牛人和兴趣的着眼点的不同... 925 10 18 0 谢小路} {2016-04-27T17:59:40+08:00 专栏：003：正则表达式 系列爬虫专栏 崇尚的学习思维是：输入，输出平衡，且维持平衡点不断精进的地步 曾经有大神告诫说：没事别瞎写文章；为此写的都是，在我能力范围内的 1：框架 2：概念 什么是正则表... 547 5 15 0 谢小路} {2016-05-21T11:43:06+08:00  专栏：016：功能强大的“图片下载器” 用理工科思维看待这个世界 系列爬虫专栏 初学者，尽力实现最小化学习系统 如何实现项目图片的下载 0：学习理念 推荐阅读简书：学习方法论我觉得对我有帮助，多问自己为什么从来不是... 1269 2 14 0 谢小路} {2016-06-02T15:02:33+08:00 线性代数：一切为了更好的理解 ![][1][1]: http://latex.codecogs.com/gif.latex?\alpha 线性代数是数学工具掌握它，打开数学的另一扇大门 1：声明 非原创，... 943 4 14 0 谢小路} {2016-11-13T13:46:36+08:00  Python 强化训练：第十篇 主题： pycharm 开发利器，减少编写错误代码的可能性，遵从PEP8编码规范. 配置 设置 安装模块 快捷键 Git 版本控制 1. 配置 setting--> proj... 437 3 14 0 谢小路}]

```



## Article

- GetAuthor 获取文章作者
- GetDescription 获取作者简介
- GetTitle 获取文章标题
- GetContent 获取文章全文


![article.png](https://upload-images.jianshu.io/upload_images/1818135-f41e684241267c67.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

#### Tutorial & Usage ==> Article


```go
package main

import (
	"fmt"
	"github.com/wuxiaoxiaoshen/jianshu-go"
)

func main(){
    article := jianshu.NewArticle("https://www.jianshu.com/p/4ea9abf7e4e8")
    fmt.Println("文章作者: ", article.GetAuthor())
    fmt.Println("文章作者简介: ", article.GetDescription())
    fmt.Println("文章标题: ", article.GetTitle())
    fmt.Println("文章内容: ", article.GetContent())
}

```

返回信息:

``` 
文章作者:  谢小路
文章作者简介:  上海大学2017级研究生毕业.微信公众号：Siwei_Jingjin
文章标题:  『简书API : jianshu 基于 golang （1）』
文章内容:  在我眼中，比较崇拜三类人：一类是设计师；一类是作家；一类是程序员。
这三类人都是通过创造、或者改善作品，不断的把世界变的更好。每每看到大师级的作品，总会不禁感叹，人与人的差别就是这么大。但是这都不阻碍我们模仿学习他们，向着更好的方向前进。
前几年，偏爱好于作家，总幻想自己能通过作品改变世界。后来证明，这条道路在真实的社会上，需要很大的毅力坚持，而且还需要点天分。
随着毕业、工作。我更偏爱设计师和程序员，而且两者在某些层面上有些共性。编程是我的本职工作，设计领域则是业余时间喜欢关注的点。
这三类人都在通过作品，不断的显现自己的能力。
所以一个程序员，假如没有开源作品，这样显的很格调不高。
开源作品质量其实也参差不齐。
一个好的开源作品：

代码质量优
解决的问题有实际用处
良好的维护
良好的文档

凡是都有第一步，第一步总是有各种各样的缺点，但这并不是不开源的理由。
也许吐槽的多了，或者别人给的意见多了。修改的多了，质量就更好了。

本项目尝试解析简书API。

编程语言：golang

主要的接口包括：

User: 个人主页信息
Article : 某篇文章的信息
Home-page: 简书主页的信息
Home-page-recommend: 简书推荐作者的信息
Home-page-topic: 简书推荐的专题信息
Publication： 简书出版信息
...


```


## Home-page 

- GetHomePagePassages 获取首页文章
- GetNewList 获取首页新上榜
- GetHotSeven 获取首页 7 日热门
- GetHotMonth 获取首页 30 日热门
- GetJianshuSchool 获取首页简书大学堂

![home-page.png](https://upload-images.jianshu.io/upload_images/1818135-dc42bcdd5286ec68.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

#### Tutorial & Usage ==>  Home-page 


```go
package main

import (
	"fmt"
	"github.com/wuxiaoxiaoshen/jianshu-go"
)

func main(){
    homePage := jianshu.NewHomePage(4, 1)
    fmt.Println("** 获取首页文章 **")
    fmt.Println(homePage.GetHomePagePassages())
    fmt.Println("** 获取首页新上榜文章 **")
    homePage.GetNewList()
    fmt.Println("** 获取首页 7 日热门文章 **")
    homePage.GetHotSeven()
    fmt.Println("** 获取首页 30 日热门文章 **")
    homePage.GetHotMonth()
    fmt.Println("** 获取简书大学堂文章 **")
    homePage.GetJianShuSchool()
}

```

返回结果:

``` 
** 获取首页文章 **
[{大雨时行  训练专题 5601 15 329  简宝玉分享会︱读书在精不在多，一起来做拆书达人吧 https://www.jianshu.com/p/60ee009446be 本文主要分享我写拆书稿的经历，并对如何拆书、如何进行有偿拆书作了简单介绍，文末有简宝玉分享会免费分享福利，想要学习如何通过拆书提高读写能力、赚取外快的朋友记得拉到文末获取福利...} {边瑶姑娘  读书 1027 40 46 1 地心游记：一部科幻与科普完美结合的小说 https://www.jianshu.com/p/6d2f71417a90 文|边瑶姑娘❤2018/03/07星期三雨开学三天了，我计划着今年每周要看两本书并写观后感。没想到我看的第一本外国科幻小说，就是这个《地心游记》。说实话，我...} {汐小埋  校园小说 826 47 65  写作课：关于小说布局 结构分析分享 https://www.jianshu.com/p/c6cad3d0672a 地点：有神么工作室第一期写作课免费分享主讲人：校园小说主编宁采野花不采臣参与者：群主暖冬布布糖糖杨杨菁菁红红伊伊浅浅以及有神么工作室管理人员，全体学员。...} {人间最美是清欢  婚姻育儿 1304 59 84 4 你若以为小孩小，你比小孩还要小 https://www.jianshu.com/p/dbf912c595f8 人人都说小孩小小孩人小心不小你若以为小孩小你比小孩还要小这是陶行知先生的一首打油诗。01看到这首诗，突然想起有一天，儿子不停撒泼，我屡劝不止，已经忍无可忍。儿子大概也...}]

** 获取首页新上榜文章 **
云松老师易经风水   苏州的苏，在繁体字中为：蘇。但从字形上看就知道这是一个草木旺盛的鱼米之乡，上以草为头，下有鱼和禾苗，禾苗产出的就是大米。可以想象生活的富足，再结合一句话：仓禀实而知礼节。你就... 11959 1 5 1
碧水山涧   今天一大早的起床拉开窗帘，满满的阳光映在眼帘。亮丽，金色的光芒有无限的温暖。我还没有出去就一身的好心情。这样的天气怎么能呆在家里。可不是，收拾好行装，走在往常行走的路... 11 0 1 0
雪纷   一个早晨，听到应老师家厨房发出轰轰轰轰的声响，什么时候了？应该比较晚了吧？遂即起身轻声走出卧室，来到厨房看看是啥情况？一个熟悉的背影夹带着丝丝缕缕的阳光跳入眼帘，是树妈在抱树... 101 4 11 3
爱如空气2   这两年，母亲一直在弟弟家帮忙带孩子，由于一些育儿观念与生活方式的不一致，原本十分和谐的家庭关系变得越来越紧张。甚至还发生过几次比较大的冲突。起因都是不值一提的琐碎小事，最后却... 15 1 0 0
赫连伯伯   点击关注文集目录7.后凉之覆灭年轻气盛的呼延超不堪家人受辱，怒杀一名后凉军官，无奈，一家人赶紧收拾收拾“细软跑”。呼延平一家其实也没地方可去，况且姑臧... 26 0 0 0
金麟圣兽   目录第二八章富二代的理想我们一口一口地喝着啤酒，酒劲一点点窜上头，身体麻醉了，心理活动倒是多起来。不经意瞟了一眼窗外，身居高楼，窗外繁星点点。要说城市最糟糕的一点便是夜... 36 3 14 0
静心漫时光   文｜静心漫时光1.高洋有个对手，邵刚。他们两个从小一起上学，一起长大，成绩不相上下，旗鼓相当。所有人都爱把他俩放在一起比较，时间久了，他俩心里也开始暗暗较劲，不愿落后于... 2 0 0 0
陪你看雪1678   随着天性的解放，艾文也渐渐脱下了斯文乖巧的一面，恶劣的小孩子本性显露了出来。调皮捣蛋，恶作剧，小把戏层出不穷，让班上的女孩子又爱又恨。而苏琪娜就是艾文实验的对象之一。但艾文... 34 0 0 0
想瘦成烟丝的土豆姐   人们只会再次高兴的看到：袭击过去了，自己还完好无损……《最后的帝国：沉睡与惊醒的满洲国》年轻人往往对战争一无所知，因此便心驰神往。《最残酷的夏天：美国人眼中的越战》打仗到... 27 4 2 0
李野航   多年来，我一直重复做着一个同一主题的梦：门关不上了。我因此长久地思考着这一主题的意义。“门”这一意象一定意味着什么。梦中的“门”肯定不是指现实世界中的那道家门，梦中的“门”乃... 11 3 1 0
凉子菇娘   上个世纪的传奇女性，曾风靡过得装扮和造型成了当今女性一直追求的时尚风标。而这个传奇的女子，皆是由纪梵希先生创作而诞生。赫本曾这样讲过，“如果没有纪梵希先生，就没有赫本。”... 290 4 23 1
Amaorent阿毛的空瓶子   石头开满了花马车驾马离去白石山的石，白泉山的泉眼无影无踪的喟叹，无影无踪我切割后的白云我阻断了的流水飞向我的风，找不到的风口碾碎的齑粉，属于石头的十月一粒粒沙石子儿，长不大... 17 0 7 0
1996专员   「我一直没想过电影里面的狗血剧情会发生在我身上，生活真是比电影还电影」当她讲出这句话的时候，我想，究竟是事情糟糕到了什么地步，哭声一直没有停过。如果你是电影的编剧，最狗血... 138 3 3 1
雅俗儿   每周给你5个小理论，到上周已经12期。整体虽然收到好评，可是发现二个问题，一个是读者方面，觉得这个理论有用，但是一口气看完5个，坚持看完的少，看完的一下消化5个难；另一个方面... 41 0 2 0
娱乐拆穿姐   《这！就是街舞》播到第三期，姐终于复活了。前两期，拉到进度条：哎，怎么还有那么多。第三期拉到进度条：靠，怎么只剩下这么点了，不够看啊。第三期一从海选进入100进49的环节，... 19 1 0 0
windy天意晚晴   今天是女生节。不过，现下大家都不愿意当小鸟依人的女生了，而愿意当独立自主的“女王”。我分析了几部孙俪的大女主戏，总结出“女王”必备的三个特质，快来看看你有没有做“女王”的潜质... 94 1 8 0
丹丘生_9115    72 12 19 0
极小光   简评：Worksmarterinsteadofharder保持9小时的睡眠这听起来似乎有点不太可能，但充足的睡眠能保证你的工作效率，自动弥补你在50%的状... 219 3 4 0

** 获取首页 7 日热门文章 **
sherry时芽   ”前几天，跟一位简书的作者在评论区掐架（我是老好人，一般不掐架），这位作者说简书没有好作品，打算离开简书。我气不过，觉得你离开简书可以，但不能说简书没有好作品，就相互怼了几句... 19611 232 606 10
蔡垒磊   曾经的高考理科状元，本科北大毕业后留美的研究生王猛，是个极其标准的学霸，然而已经12年没有回家过春节了。6年前，他拉黑了父母所有的联系方式。这样的报道，引起了社会上的轩然大... 82275 1208 3328 18
七小葩   如果你还在用百度搜图，那真是“棒棒哒”，像你这么专一的人不多了。经常有人会问我，你一般在哪找的图片？不少朋友都会有自己的素材库，需要时提取就可以了。不过像我这么懒的人，... 98956 589 14244 15
田宝谈写作   01昨晚我失眠了。失眠的原因是看了电影《无问西东》，其中一个场景久久地萦绕于脑中：吴岭澜是影中一位文学天赋极高但理科极差的学生。对于人生，他很少思考过，一味地随着大众的... 97487 487 4406 13
衷曲无闻   01木心的《琼美卡随想录》里有一言，“我好久没有以小步紧跑去迎接一个人的那种快乐了”。每每读到，都有一种孤独之感生出。周国平曾写过，在最内在的精神生活中，我们每个人都是孤... 42950 566 3020 10
凉亦歌   文／凉亦歌小时候我读幼儿园，老师教我们学三字经，第一句是“人之初，性本善”。后来很长一段时间，我都把这句话作为衡量所有人的标准，直到有一天被人嘲笑，才发现自己有多蠢。人... 50111 931 2171 10
道长是名思维贩子    60887 154 付费 0
东篱若尘   近日，忽然看见某人发文，指我为自己的网名“东篱若尘”所作的小诗是抄袭许嵩的山水之间的歌词，本来以为只是好事之人不解其意，也就解释两句罢了，不想后来才知道，这是被简书里的某个小... 8297 0 97 0
尹沽城   我先说三点：其一，这是一次免费课程。既然一年入账〇元，我也确实没脸收费。其二，这不是绝对的原创。有前辈的经验之谈，也有我自己写作十年的体会。其三，能打开这篇文章的，我隐... 30808 284 2551 15
诗音姑娘   01.昨天刷微博看到这样一条新闻，出身“二本”学生质疑：为什么你们不看我的简历？一位普通二本师范类院校的经济专业毕业生自述求职感受：学校好像成为我面对询问时一道尴尬的伤疤... 38249 359 1812 3
Angela在悉尼   最新一期的《十三邀》，主持人许知远与《吐槽大会》策划人李诞对谈，作为功成名就的年轻85后精英代表，李诞一直自信满满，但当被问到毕业学校背景时，他却画风陡转，感慨起来：“我目... 117890 448 4661 6
叫我梅梅呀   文/梅梅01前几天和久别重逢的高中室友去看电影，《无问西东》。看到清华一代代青春洋溢的年轻人，心里却总想着我那个平平常常的母亲。她20岁的目标是考警校，当警察。但最... 26318 424 1446 5
奇奇漫悦读   编者按：这是一个真实的故事。告诉你灾难来临时，你该怎么做……01离海清中学五六个路口远的地方有个叫燕园的废旧水库，一条陈旧的公路环绕着堤坝，那公路窄得仅能容下一辆电动三轮... 82888 520 1532 15
欢喜非喜   1、马上就是农历新年了，春节是一个理所当然的，要回家的日子，因为，那是一个理所当然的，团圆的日子。异乡工作的人抢着车票，身边的人都时时的讨论着小时候在自己的家过年是一个什... 25143 565 682 6
郑小喵   文|郑小喵-1-记得前段时间，有朋友问我:你现在怎么变得这么自律？我仔细想了想，原来，我也在向着自己喜欢的方向，慢慢地靠近啊。以前的我，是一个三分钟热度的人，说好了减... 45946 383 2135 14
而你在明天   上学的时候，很多人都说，现在这个社会，学历并不重要，企业看重的是你的个人能力。直到毕业了开始找工作你才发现，面对人山人海的求职者，公司用于衡量你个人能力的正是学历。大多数HR... 35555 205 2053 6
科技毒瘤   01ColdTurkey老规矩，第一款推荐的软件必须超强是实力派！天降大任于斯人也，必先关其手机，收其平板，拔其网线，这也道出了不少朋友的痛处，当我们想专注做某一件事... 46436 67 3619 3
野蛮人诺基亚   头脑王者被封的原因是违反了《即时通信工具公众信息服务发展管理暂行规定》《规定》里面对即时通讯工具的定义是这样的本规定所称即时通信工具，是指基于互联网面向终端使用者提供即时... 145647 0 688 1

** 获取首页 30 日热门文章 **
sherry时芽   ”前几天，跟一位简书的作者在评论区掐架（我是老好人，一般不掐架），这位作者说简书没有好作品，打算离开简书。我气不过，觉得你离开简书可以，但不能说简书没有好作品，就相互怼了几句... 19611 232 606 10
晓多   01未经审视的人生不值得一提。我特别期待看你的2018年的梦想清单特别想知道你是怎么实现了那么多目标的？有时候我觉得你是很棒的人，工作拿那么多荣誉，遴选考全国岗位第一、成... 126407 1594 9485 15
七小葩   如果你还在用百度搜图，那真是“棒棒哒”，像你这么专一的人不多了。经常有人会问我，你一般在哪找的图片？不少朋友都会有自己的素材库，需要时提取就可以了。不过像我这么懒的人，... 98956 589 14244 15
羊达令   前不久，和几位上学时的小伙伴聊天，话题转到了我们现在的生活上。许久未联系的我们，说着每个人的近况。小陈在家乡考上了事业编，每天朝九晚五，上班忙碌，下班觉得很累，看着别人每... 82895 1073 4536 17
衷曲无闻   01刷朋友圈的时候，看到一条由三感音乐故事拍摄的短视频，被文案戳中了泪点。22岁生日，一个人吃火锅，还好锅底可以点最辣的；187次路过的码头，4次遇到一对情侣，两个人眼中... 59053 725 3176 21
单文   喜欢读书，读书的时候，我是快乐的，如浸泡在阳光中的树苗。喜欢读书，读书的时候，我是自由的，如庄周梦中的蝶。朋友说：我若是有个儿子，我要教他弹钢琴，教他打篮球，教他骑自行车... 51589 480 3969 4
亭主   一、不怕，我们都是普通人借了彭小六的小配图，画很简单，但很形象。这个世界上，天才不多、资质愚笨的人也不多，绝大部分人都是如你我一样的普通人。当我们认清这个事实，反而没什... 42037 201 3978 16
甜到齁的白日梦   01我在上小学，初中，高中，甚至上大学，每每开学前，我的父亲就会告诉我：“孩子，到了学校，要和周围人搞好关系，一定要合群，跟着大家走，不要搞什么特殊化，你这周围的同学，以后... 41751 555 2562 8
MY麦子   文|MY麦子今年三月份的时候辞去了工作，打算好好的放松一下再重新开始，上班的时候就是每天按时去上班，突然不上班了空闲时间也多了起来。平时除了去学车，其他时间都是一个人，... 156524 707 7012 8
Angela在悉尼   最新一期的《十三邀》，主持人许知远与《吐槽大会》策划人李诞对谈，作为功成名就的年轻85后精英代表，李诞一直自信满满，但当被问到毕业学校背景时，他却画风陡转，感慨起来：“我目... 117890 448 4661 6
谈心社   90后法学院副教授陈少威火了。看着他帅气的侧颜，女网友们纷纷坐不住了，惊呼：这简直就是现实版何以琛！不知不觉中，90后已经开始走上各行各业，那些家长眼里“难管教”的孩子，... 71609 431 2361 6
吕白Alex   被《无问西东》感动了。不是因为它人物塑造得有多好，它讲故事的方式有多好，它的剪辑手法有多厉害。事实上有很多地方人物的动机不明确，讲故事的手法有些拙劣，剪辑更是有些头重脚轻。... 54466 321 2632 12
黄不会   -1-前段时间在刷知乎tl的时候看到了一个事情。其实就是一对在杭州萧山的情侣的聊天记录。我贴出来给大家：顺便为了大家方便阅读，我提炼下男女双方的观点出来，给大家看。... 34970 835 1434 9
瓜子跑得快   知乎热榜上有这样一个问题：你还相信读书能改变气质吗？看脸的时代，气质这东西实在太虚无缥缈了，一个胖子和一个经常看书的胖子有什么区别？回复区的几条热评，概括起来就是：读书... 42802 684 2190 11
苏梨衣   如果我的人生标签只能设置3个词汇，那么，一定有一个位置是留给“旅行”的。和很多女生一样，我向往着旅行和远方。在顾少强的“世界那么大，我想去看看”还没火之前，我就已经背上行囊... 41022 858 1783 12
委婉的鱼   今天给大家分享几个国内自学ps的网站，上面的教程大部分还是很不错，希望可以帮助到大家。1.我要自学网网址：http://www.51zxw.net/list.aspx?c... 24167 173 3099 0
蔡垒磊   曾经的高考理科状元，本科北大毕业后留美的研究生王猛，是个极其标准的学霸，然而已经12年没有回家过春节了。6年前，他拉黑了父母所有的联系方式。这样的报道，引起了社会上的轩然大... 82275 1208 3328 18
诗音姑娘   01.前几天看了张艺兴在《偶像练习生》担当全民制作人的宣传片，他很严肃的说，“没准备好，你凭什么上舞台？”播到这句话时，身边一个号称“张艺兴迷妹”的同学，皱了皱眉，喃喃道... 41914 357 2354 6
温佛佳   文/温佛佳硕士毕业，离开书香氤氲的北京师范大学校园，已近三年。2017年全国硕士研究生入学考试拉上帷幕不久，最近，接到不少学弟学妹的电话，他们的问题，大抵围绕两方面:1... 38721 345 1990 8

** 获取简书大学堂文章 **
简书大学堂   过完了少女or女王节，各位漂亮的小姐姐们是不是都收到了来自各方的鲜花与美妆用品（反正在简书上班的小姐姐们都有）~学堂菌在心里默默感叹，如今的节日基本等同于狂欢购物日啊~本... 190 0 9 0
一河漪沫   如果有人告诉你，拍证件照不用再去照相馆，书籍封面名片海报都能做，亲朋好友的忙随时都能帮，超喜欢的图片想用就能用，还能兼职排版设计挣外块，……你想不想试一试？只... 5796 54 100 0
李陌359   很多人的心里，都藏着一个成为小说家的梦想。可同时，这似乎又是个难以企及的梦想。有的人，绞尽脑汁，却构思不出一个精彩有趣的故事，于是，还没等真正落笔，便不得不宣告放弃。有的... 25245 96 693 0
简书大学堂   最近啊，各种不顺~先是开年会，没有中奖，接着连阳光普照奖都丢了，丢就丢了呗，回家路上手机也被偷了......这快过年了，都是什么事啊！！！不行，丧也要拉上一个，这... 3171 14 30 0
心蓝丫头   我从来没有想过，我的人生，会在而立之年，因为一张画，而被彻底改变。少年时，我是那个被语文老师认为有文字天赋，会成为作家的孩子。与此同时，也是那个拿了一张画了一串葡萄的画纸想... 34251 200 867 1
心蓝丫头   18年前，我说我想当个狗仔队，专写明星八卦新闻。于是还未成年的我，跑去当时国内最大的娱乐杂志应聘当编辑，被破格录用，一边念书一边兼职，把娱乐八卦写了个够本。12年前，我说我... 13274 219 408 1
简书大学堂   学习工作不知道有什么意义，找不到人生方向？是在大公司混履历还是出去创业，人生选择太多陷入焦虑？在学校当了N年学霸，毕业以后秒变职场小白，心里落差巨大？困于工作倦怠期，不甘于稳... 3542 5 85 0
BeBeyond   学习工作不知道有什么意义，找不到人生方向？是在大公司混履历还是出去创业，人生选择太多陷入焦虑？在学校当了N年学霸，毕业以后秒变职场小白，心里落差巨大？困于工作倦怠期，不甘于稳... 4590 6 93 0
一鸣   文|南方有路作家余华在作品《兄弟》后记里写过这么一段话“起初我的构思是一部十万字左右的小说，可是叙述统治了我的写作，篇幅超过了五十万字。写作就是这样奇妙，从狭窄开始... 7005 81 119 0
韩大爷的杂货铺   1.各位同学：大家好，我是韩大爷，让大家久等了。称呼屏幕前的您为同学，是我在100天领读营中留下的习惯，“同学”，这里取“同舟共济，共同学习”之意。过去的147天里，... 7549 66 108 0

```


## Home-page-recommend

- GetListRecommendAuthor 获取首页推荐作者

![recommendAuthor.png](https://upload-images.jianshu.io/upload_images/1818135-0c31b9513039cbc1.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

#### Tutorial & Usage ==>  Home-page-recommend

```go

package main

import (
	"fmt"
	"jianshu-go"
)

func main(){
    recommendAuthor := jianshu.NewRecommendAuthor(10)
    fmt.Println("获取简书推荐10位作者: ", recommendAuthor.GetListRecommendAuthor())
}

```

返回结果：

``` 
获取简书推荐10位作者:  [{寻麦 https://www.jianshu.com/users/e0ef486d9b90 None 卖童话的小巫师，强迫症晚期。 《菜菜的奇幻之旅》第十二章：木匠的茶话会致读者：在妖怪的世界里，找回丢失的童真和勇气《菜菜的奇幻之旅》第十一章：自言自语的短鼻象} {遛遛心情的溜妈 https://www.jianshu.com/users/dbfdce352c0d None 溜爸，简书签约作者，舞枪弄棒、舞文弄墨的计算... 清明，约你一起种田，读书……一本吸引我前前后后读了三遍的书，和你分享（下）［奇幻］唐朝那些猫事儿（57）} {驴光掠影 https://www.jianshu.com/users/b3b2c03354f3 None 曾在千亿级A股上市公司任IT高管，现任某创业... 认知系列|知识留存率最高的阅读原则随笔|我去年买了个币云南滇西|你去了一个假的大理} {简书出版 https://www.jianshu.com/users/55b597320c4e None 给我发简信前请先看完我的简介。关于如... 简书2017年的10个好故事（虚构&非虚构）上线整整十部优质付费连载，让你周末看个够更新一小波实用性非常强的优质付费连载} {末晓 https://www.jianshu.com/users/d0f9f963aeb9 None 每一个故事都是写给自己，希望路过的你们恰好喜欢 ［童话］花仙卡卡的魔法修行（7）［童话］花仙卡卡的魔法修行（6）花仙卡卡的魔法修行（5）} {雨落荒原 https://www.jianshu.com/users/7406f22f461e woman 反反香菜联盟会会长合作请联系简书版权... 【3000字转折大赛】两枚硬币《对岸》（10）被男人包围的女人-往事关于写作的一些个人体验} {格列柯南 https://www.jianshu.com/users/ffc565d738a3 man 历史博士、哲学硕士，本科经济学。只有审视生活... 四十而立，五十不惑，六十用伟哥《三言二拍》教你如何写豹纹父母反应过度与失败的孩子教育} {简书大学堂 https://www.jianshu.com/users/c5580cc1c3f4 None 简书自有的学习成长平台。简书大学堂讲... 被忽视的植树节，学堂菌决定来种草上架这个功能，“累死"了一位产品经理这24条职场弯路，你本可以绕开的...} {aloho https://www.jianshu.com/users/1446a350e58a man aloho。名字改编自夏威夷语aloha，你... 有种稳赚不赔的买卖叫环游中国|《绝对光年》番外篇aloho的床头诗601-650一道数学题，解答你能不能环游中国的困惑|《绝对光年》番外篇} {魏童 https://www.jianshu.com/users/5462ec6828f6 man 魏童，知名图书策划人、影视策划人。先后在盛大... 如何发现你的创作天赋和优势？天才的创作者与精明的经营者世间的那些相遇}]


```


## Home-page-topic

- GetTopicCollectionRecommend 获取推荐专题
- GetTopicCollectionHot 获取热门专题
- GetTopicCollectionCity 获取城市专题
- GetTopicCollectionSchoolyard 获取校园专题


![topic.png](https://upload-images.jianshu.io/upload_images/1818135-10880c9494a1afaa.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

#### Tutorial & Usage ==>  Home-page-topic

```go
package main

import (
	"fmt"
	"github.com/wuxiaoxiaoshen/jianshu-go"
)


func main(){
    topic := jianshu.NewTopic(12)
    fmt.Println("获取推荐专题 12 个: ", topic.GetTopicCollectionRecommend())
    fmt.Println("获取热门专题 12 个: ", topic.GetTopicCollectionHot())
    fmt.Println("获取城市专题 12 个: ", topic.GetTopicCollectionCity())
    fmt.Println("获取校园专题 12 个: ", topic.GetTopicCollectionSchoolyard())
}

```


返回结果:

``` 
获取推荐专题 12 个:  [{故事 故事专题，不论是旅行生活中亲身经历的真实故事，还是童话玄幻遐想的虚构故... 162698篇文章 · 764.2K人关注} {旅行·在路上 邂逅一座城池，讲述一段故事，分享你的旅行。欢迎来到「旅行·在路上」专题... 97891篇文章 · 1887.3K人关注} {读书 投稿须知：读书专题仅收录与读书有关的书评、读书笔记、阅读方法、读书... 163277篇文章 · 2565.6K人关注} {人文社科 欢迎关注公众号简宝玉(公众号ID:jianshu4321)后台... 7168篇文章 · 244.5K人关注} {摄影 欢迎关注公众号简宝玉(公众号ID:jianshu4321)后台... 57253篇文章 · 1408.6K人关注} {历史 本专题主要收录历史阐述与分析类文章，也收录历史小说等带一定想象性的文章... 42474篇文章 · 1432.5K人关注} {简书电影 欢迎关注公众号简宝玉(公众号ID:jianshu4321)后台... 77310篇文章 · 1274.9K人关注} {社会热点 欢迎关注公众号简宝玉(公众号ID:jianshu4321)后台... 44300篇文章 · 1242.6K人关注} {@IT·互联网 欢迎关注公众号简宝玉(公众号ID:jianshu4321)后台... 42758篇文章 · 1160.9K人关注} {手绘 欢迎关注公众号简宝玉(公众号ID:jianshu4321)后台... 94646篇文章 · 1194.4K人关注} {自然科普 本专题仅收涉及自然科学（生物、天文、物理等）的科普、介绍、发现类文章，... 3150篇文章 · 665.7K人关注} {短篇小说 借一斑略知全豹，以一目尽传精神。阅读简书最好看的故事，欢迎订... 68914篇文章 · 1758.8K人关注}]
获取热门专题 12 个:  [{首页投稿 玩转简书的第一步，从这个专题开始。想上首页热门榜么？好内容想被... 384659篇文章 · 193.6K人关注} {连载小说 欢迎关注公众号简宝玉(公众号ID:jianshu4321)后台... 134282篇文章 · 945.4K人关注} {故事 故事专题，不论是旅行生活中亲身经历的真实故事，还是童话玄幻遐想的虚构故... 162698篇文章 · 764.2K人关注} {程序员 欢迎关注公众号简宝玉(公众号ID:jianshu4321)后台... 81901篇文章 · 708.3K人关注} {旅行·在路上 邂逅一座城池，讲述一段故事，分享你的旅行。欢迎来到「旅行·在路上」专题... 97891篇文章 · 1887.3K人关注} {短篇小说 借一斑略知全豹，以一目尽传精神。阅读简书最好看的故事，欢迎订... 68914篇文章 · 1758.8K人关注} {哲思 投稿须知：1、字数不限，收入以哲学思辨（中西方哲学观点和作者本人的... 11672篇文章 · 1094.2K人关注} {摄影 欢迎关注公众号简宝玉(公众号ID:jianshu4321)后台... 57253篇文章 · 1408.6K人关注} {@IT·互联网 欢迎关注公众号简宝玉(公众号ID:jianshu4321)后台... 42758篇文章 · 1160.9K人关注} {人物 投稿须知：■人物专题仅收录以人物为主题的文章。例如：历史人物，... 35483篇文章 · 858.4K人关注} {电竞·游戏 少年，一起来玩游戏吗？投稿须知：1.本专题仅收录关于游戏相... 10177篇文章 · 268.7K人关注} {每天写1000字 提高写作能力，成为更好的自己。练习每天写1000字，打卡提... 324196篇文章 · 107.4K人关注}]
获取城市专题 12 个:  [{陕西省 文明起步的地方，你看，炎黄大帝都笑了，你还在等什么呢？投稿须知... 5659篇文章 · 2.0K人关注} {香港 香港，中国的一扇窗，一方面把自己打扮得珠光宝气，一方面又保持古老传统。... 4081篇文章 · 3.3K人关注} {福建省 在鼓浪屿的迷蒙雨中，闲坐在屋檐下，望着不远处的海发呆，也是一种享受。... 1414篇文章 · 751人关注} {辽宁省 简介：共和国的长子、东北的心脏、美丽的辽宁欢迎您。欢迎加入简书... 673篇文章 · 1.0K人关注} {台湾 酷酷的湾仔，嗲嗲的台妹，旖旎的风光……这就是我们的宝岛台湾！... 832篇文章 · 28.2K人关注} {广东省 传承岭南文化，领略南粤风情，在这里遇见最美好的广东。﻿投稿须... 2107篇文章 · 3.0K人关注} {贵州省 『一次“黔”行，终身难忘』 贵州省，简称黔或贵。 这里有世... 1704篇文章 · 867人关注} {云南省 踏入美丽、丰饶、神奇的云南秘境，回到最初的地方，灵魂、身体、美食，总有... 1392篇文章 · 989人关注} {河南省 还你一个真实、全面的河南。投稿须知：http://www.ji... 1534篇文章 · 1.7K人关注} {江西省 投稿须知：http://www.jianshu.com/p/77a35... 500篇文章 · 552人关注} {湖南省 湘江北去，芙蓉天成。 冒有胆子，就不要搞湖湘文化。你要有胆... 1642篇文章 · 1.5K人关注} {黑龙江省 用龙江的坚毅与柔情，谱写我们的未来。投稿须知：http://w... 417篇文章 · 442人关注}]
获取校园专题 12 个:  [{南昌大学 “未完待续就要表白”全国大学生表白母校活动——南昌大学　　格物致... 370篇文章 · 274人关注} {唐山师范学院 表白母校活动，有你参与更精彩! 118篇文章 · 164人关注} {潍坊科技学院 修身博学求索笃行欢迎大家踊跃投稿 195篇文章 · 82人关注} {齐鲁工业大学 于千万人之中，遇见你所要遇见的人，于千万年之中，时间的无涯的荒野里，没... 119篇文章 · 103人关注} {大学生活 本专题收录大学相关文章。要求文内不得带有公众号信息等推广内容。详细规... 75987篇文章 · 706.2K人关注} {南阳理工学院 今厚德博学南工之学子，书求真至善温暖之文字。投稿须知：... 126篇文章 · 256人关注} {南京师范大学泰州学院  534篇文章 · 479人关注} {安徽师范大学 简书师大——专属于师大人的自由创作社区创作、分享、交流，在这里我... 2619篇文章 · 1.6K人关注} {山西师范大学 山西师范大学（ShanxiNormalUniversity），简称... 83篇文章 · 61人关注} {河北地质大学华信学院 河北地质大学华信学院，位于石家庄市，成立于2001年，是经教育部和河北... 560篇文章 · 418人关注} {河北大学 实事求是，笃学诚行 124篇文章 · 170人关注} {南京工业大学浦江学院 1、在1月5号以前，每周可在简书app上发表一篇读书感悟，字数500字... 74篇文章 · 353人关注}]

```


## Publication


- GetPublicizedBook 获取已出版图书
- GetNovelBooks 获取小说
- GetITAndJobMarket 获取IT、理财、职场
- GetCultuereAndHistory 获取文化、历史
- GetMonthlyMagazine 获取专题月刊

![publication-one.png](https://upload-images.jianshu.io/upload_images/1818135-85e4b0cf1305c1eb.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![publication-two.png](https://upload-images.jianshu.io/upload_images/1818135-ffc616be74934c35.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![publication-three.png](https://upload-images.jianshu.io/upload_images/1818135-cfbe6d2f805fb572.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![publication-four.png](https://upload-images.jianshu.io/upload_images/1818135-faea20e3bda89de6.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

![publication-five.png](https://upload-images.jianshu.io/upload_images/1818135-cbc20eea744e722c.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


#### Tutorial & Usage ==>  Publication

```go
package main

import (
	"fmt"
	"github.com/wuxiaoxiaoshen/jianshu-go"
)

func main(){
    publication := jianshu.NewPublication()
    fmt.Println("获取简书已出版图书")
    publication.GetPublicizedBook()
    fmt.Println("获取简书小说")
    publication.GetNovelBooks()
    fmt.Println("获取简书IT、理财、职场图书")
    publication.GetITAndJobMarket()
    fmt.Println("获取简书文化、历史")
    publication.GetCultureAndHistory()
    fmt.Println("获取简书专题月刊")
    publication.GetMonthlyMagazine()
}
```

返回结果：

``` 
获取简书已出版图书
[{陈信诚 转型王道:微商、社交电商实战系统 None ￥ 69.80 https://book.douban.com/subject/27622583/} {顾一宸 管他努力有没有回报，拼过才是人生 None ￥ 38.00 https://book.douban.com/subject/27596903/} {徐成东 从职场小白到团队老大：职场基本思维 None ￥ 39.00 https://book.douban.com/subject/27620324/} {简书 爱上写作，遇见更好的自己 None ￥ 38.00 https://book.douban.com/subject/27620180/} {李陌 失控的布局 出来混，总是要还的！但是，真相并非如此…… ￥ 35.00 https://book.douban.com/subject/27615008/ } {佰稼 别害怕，我们都孤寡 这七篇故事里的人孤独得难以自拔，因为现在就是一个孤独的时代。人与人之间情感越来越淡薄，正因为我们重感情，所以才不敢轻易地付以真情。 ￥ 39.00 https://book.douban.com/subject/27199582/} {尚元用 不要跟踪我 细腻入微的情感，环环相扣、跌宕起伏的情节，出人意料的结局，爱情与悬疑的完美结合。 ￥ 32.00 https://book.douban.com/subject/27113877/} {衷曲无闻 梦想不会辜负努力的你 写给对自己寄予厚望的你 ￥ 38.00 https://book.douban.com/subject/27077802/} {简书 总要为梦想疯狂一次 那些年陪你追梦的人还在身边吗？简书燃情作家和你同追梦共青春。 ￥ 36.00 https://book.douban.com/subject/27065409/} {简书 世界和它的悲欢 这世界永恒演绎着它的悲欢，故事里总能找到你我的影子。 ￥ 42.00 https://book.douban.com/subject/27008382/} {简书 终有人住进你心里 蛾扑火的勇气、奋不顾身的着迷、故作淡定的祝福… ￥ 36.00 https://book.douban.com/subject/26980212/} {沐丞 理财要趁早 理财不是发财，理财也不是简单的是投资、让钱生钱。 ￥ 30.00 https://book.douban.com/subject/26943302/} {简书 越勇敢的女人越幸运 本书是简书出品的第二本书，温暖又犀利地讲述了女人应该勇敢面对生活、爱情、自我成长等方面所遇到的问题和矛盾 ￥ 42.00 https://book.douban.com/subject/26926348/} {简书 你一定要努力，但千万别着急 唯有承担起厚重的经历，才能禁得起岁月的推敲。 ￥ 42.00 https://book.douban.com/subject/26786048/}]
获取简书小说
[{芝小麻 我离婚后的第一年 None ￥ 1.99 http://www.jianshu.com/p/11a941858272} {一元亦有用 神说要有光，我便开了灯 None 免费 http://www.jianshu.com/p/d2303a3b5b41} {离影疏落 读懂你的心——我的心理咨询师手记 None ￥ 1.99 http://www.jianshu.com/p/ab6a63674f62} {潮范 被骗入传销的13天 None ￥ 1.99 http://www.jianshu.com/p/b449e1fd14c9} {B杜 法兰西情人 None ￥ 12.00 http://www.jianshu.com/p/222d00822ea3} {无戒 我的山沟沟，我的女人 None ￥ 1.99 http://www.jianshu.com/p/832453f46086} {佰稼 再见，吾爱 None ￥ 1.99 http://www.jianshu.com/p/cfe52ce819ea} {简书短篇小说专题 你好，2027 None 免费 http://www.jianshu.com/p/855fdff634f0} {佰稼 无爱纪 None ￥ 1.99 http://www.jianshu.com/p/007fe8f6568e} {大肚子猫 爱在郁金香之国 None ￥ 3.99 http://www.jianshu.com/p/53754b3e7492}]
获取简书IT、理财、职场图书
[{齐帆齐 我的草根奋斗故事 None ￥ 1.99 http://www.jianshu.com/p/b2e0a9e9b804} {徐成东 《东哥说职场》第二辑 高效沟通与情绪管理 None ￥ 2.99 http://www.jianshu.com/p/2b7b235492b5} {沐丞 买房要趁早 None ￥ 9.99 http://www.jianshu.com/p/be9de8b0c845} {徐成东 《东哥说职场》第一辑 求职与加薪的正确姿势 None ￥ 4.99 http://www.jianshu.com/p/2e591ca7435b} {蓝桥飞 游戏策划入门修行 None ￥ 1.99 http://www.jianshu.com/p/80e630e81c2a} {庄表伟 开源思索集 None ￥ 1.99 http://www.jianshu.com/p/e5c5af1e45dc} {沐丞 工作十年买下三套房 None ￥ 1.99 http://www.jianshu.com/p/bf306d2725f2} {纯银V 10个案例说明什么是产品模型 None ￥ 1.99 http://www.jianshu.com/p/f74527bd0e04} {bylin 取悦的工序：如何理解游戏 None 免费 http://www.jianshu.com/p/d4be65e0b792} {萌大夫精神病院 我也曾经上班一个礼拜就辞职 None 免费 http://www.jianshu.com/p/d99d0cc4575e}]
获取简书文化、历史
[{芊语芊寻 当你想要学民乐》 None ￥ 3.99 http://www.jianshu.com/p/65c3c8163162} {陈缃眠 故人之书 None ￥ 1.99 http://www.jianshu.com/p/55f735010b37} {王祎 没想到你是这样的美术史 None ￥ 2.99 http://www.jianshu.com/p/d706466fdee3} {马风 宋词·萧红·林徽因/一丘一壑也风流 None ￥ 6.99 http://www.jianshu.com/p/ae719ebb0cbc} {一个历史围观群众 民国人物趣味杂谈 None ￥ 1.99 http://www.jianshu.com/p/c1f68cf064c0} {饱醉豚 新加坡的那些事儿 None ￥ 2.99 http://www.jianshu.com/p/a11911c81ed0}]
获取简书专题月刊
[{简书茶馆叶老板 简书茶点故事精选第二辑 None 免费 http://www.jianshu.com/p/156896ff347e} {简书教育专题 简书教育月刊002 · 当我们谈论教育时，我们谈论的是什么 None 免费 http://www.jianshu.com/p/a368a16c75fe} {简书短篇小说专题 简书短篇小说月刊005 · 传奇故事 世间冷暖 None 免费 http://www.jianshu.com/p/6472808c21bd} {简书谈写作专题 简书谈写作月刊003 · 品读岁月，书写人生 None 免费 http://www.jianshu.com/p/dc8134000909} {简书读书专题 简书读书月刊001 · 穷理尽性 以至于命 None 免费 http://www.jianshu.com/p/78d0b28fe8f8} {简书世间事专题 简书世间事月刊001 · 不动声色 None 免费 http://www.jianshu.com/p/ae6bc1b682d0} {简书奇思妙想专题 简书奇思妙想月刊001 · 有梦为马，随处可栖 None 免费 http://www.jianshu.com/p/1d3f71fadc63} {简书茶馆叶老板 简书茶点故事7月刊 None 免费 http://www.jianshu.com/p/56217ea0170a} {简书谈写作专题 简书谈写作月刊002 · 文字情缘 None 免费 http://www.jianshu.com/p/cf12eedb4536} {简书人物专题 简书人物月刊001 · 再回首绝代芳华 None 免费 http://www.jianshu.com/p/f1dd98c85e13}]


```

## Special-subject

- GetSpecialSubjectURL 获取专题URL
- GetSpecialSubjectTitle 获取专题标题
- GetSpecialSubjectPassageNumber 获取专题文章数目
- GetSpecialSubjectFollowerNumber 获取专题关注数目
- GetSpecialSubjectNotice 获取专题公告
- GetSpecialSubjectTitleAdministrator 获取专题创建者信息
- GetSpecialSubjectNewComment 获取专题最新评论文章信息
- GetSpecialSubjectNewAdd 获取专题最新收录文章信息
- GetSpecialSubjectHot 获取专题热门文章信息

#### Tutorial & Usage ==>  Special-subject

```go


package main

import (
	"fmt"
	"github.com/wuxiaoxiaoshen/jianshu-go"
)

func main(){
	fmt.Println("** Special-subject **")
    	specialSubject := jianshu.NewSpecialSubject("https://www.jianshu.com/c/dfcf1390085c")
    	fmt.Println("获取专题标题: ", specialSubject.GetSpecialSubjectTitle())
    	fmt.Println("获取专题公告: ", specialSubject.GetSpecialSubjectNotice())
    	fmt.Println("获取专题文章数目: ", specialSubject.GetSpecialSubjectPassageNumber())
    	fmt.Println("获取专题文章关注人数: ", specialSubject.GetSpecialSubjectFollowerNumber())
    	fmt.Println("获取专题创建者信息: ")
    	jianshu.JsonPretty(specialSubject.GetSpecialSubjectTitleAdministrator())
    	fmt.Println("获取专题最新评论文章信息: ", specialSubject.GetSpecialSubjectNewComment())
    	fmt.Println("获取专题最新收录文章信息: ", specialSubject.GetSpecialSubjectNewAdd())
    	fmt.Println("获取专题最热门文章信息: ", specialSubject.GetSpecialSubjectHot())
}
```

返回信息:

``` 
** Special-subject **
获取专题标题:  爬虫专栏
获取专题公告:  鉴于越来越多的初学者进行爬虫分享，本着“精进”的原则，将提高收录爬虫文章的门槛。1.编程语言2.爬虫知识3.持续精进
获取专题文章数目:  263
获取专题文章关注人数:  953
获取专题创建者信息: 
{
 "url": "https://www.jianshu.com/u/58f0817209aa",
 "nickname": "谢小路"
}
获取专题最新评论文章信息:  [{2017-02-19T11:10:25+08:00 简书风云榜 由于开学原因，数据爬取中断，共爬取了347294条数据。 爬取时间为2月14号。 以粉丝量进行排序排名，列出简书千人风云榜。 此文章不代表简书官方数据。 签约作者 总共爬取了... 3677 80 58 7 罗罗攀} {2017-07-30T23:10:47+08:00 爬取张佳玮138w+知乎关注者：数据可视化 一、前言 作为简书上第一篇文章，先介绍下小背景，即为什么爬知乎第一大V张公子的138w+关注者信息？ 其实之前也写过不少小爬虫，按照网上各种教程实例去练手，“不可避免”的爬过... 5554 32 130 3 Deserts_X} {2017-02-18T07:16:15+08:00 使用Selenium模拟浏览器，实现自动爬取数据 最近需要在一个网站下载一批数据。但是输入一个查询，返回三四万条结果，每次只能导出500条，而且每次还得输入下载条目的范围！这样点击下载，还不要了我的老命。于是乎想自动化这个过... 3821 1 15 0 teelada} {2018-02-18T21:34:33+08:00 Java网络爬虫实操（4） 上一篇：Java网络爬虫实操（3） 本篇文章继续围绕NetDiscovery框架中pipeline的用法，结合另一个专门爬图片的框架PicCrawler，实现图片的批量下载和... 65 3 7 0 sinkinka} {2017-01-24T12:21:31+08:00 Python 3 爬虫学习笔记 （一） 这是我自己在学习python 3爬虫时的小笔记,做备忘用，难免会有一些错误和疏漏,望指正~~~Python 3 爬虫学习笔记 （二）Python 3 爬虫学习笔记 （三）Py... 4394 10 51 0 Veniendeavor} {2017-04-24T16:58:57+08:00 雪球网沪深全站股票评论爬虫 这个爬虫写得好累，就简单讲一下思路吧。雪球网股票的评论内容是不能直接访问的，必须要携带在第一次访问时雪球网写进本地的cookie（其实你随便打开一次官网就是属于第一次访问了，... 1537 4 11 1 蜗牛仔} {2018-03-01T12:54:05+08:00 python爬虫CCTV2《环球财经连线》 笔者前言：平时比较喜欢看CCTV2的《国际财经报道》（原环球财经连线），但是央视的app没有提供缓存功能，笔者又想上下班和跑步的时候看或者听。于是花了大概一天时间写了pyth... 121 1 5 0 大树先生g} {2016-07-28T11:35:57+08:00 爬取简书全站文章并生成 API（一） 简书中的优质文章非常多，而且我非常喜欢 Markdown 这种语法格式，所以想着能不能爬取简书上面的文章，爬取文章之前先带大家来了解下简书整个网站，简书的首页分为“热门（已推... 5439 46 115 1 田飞雨} {2017-02-18T14:40:15+08:00 使用python抓取美女福利图片 这篇文章干嘛的？ 本屌在上网时偶然看到一个图片网站，网站的尺度是这样的： 里面的美女露骨而不露点，简直是宅男福利。一时兴起，决定将网站上的图片down下来研究研究。正好最近在... 4737 11 39 0 新新格子君} {2016-11-18T13:37:12+08:00 一步一步教你 https 抓包 在 Mac 上常用的抓包软件是 Charles，网上关于 Charles 的教程很多，这里介绍另一个抓包神器 mitmproxy。mitmproxy 是一款可交互式的命令行抓... 16864 119 593 2 hi_xgb}]
获取专题最新收录文章信息:  [{2018-03-13T23:29:43+08:00 『简书API：jianshu 基于golang -- 用法介绍 （2）』 首先我做这个项目的目的是朴素的： 熟悉golang 语法 通过这个项目呢，大家可以分析任意网站， 任意解析网站形成自己的API。 这个项目受项目：zhihu-go 影响。阅读... 18 0 4 0 谢小路} {2018-03-11T20:12:04+08:00 『简书API : jianshu 基于 golang （1）』 在我眼中，比较崇拜三类人：一类是设计师；一类是作家；一类是程序员。 这三类人都是通过创造、或者改善作品，不断的把世界变的更好。每每看到大师级的作品，总会不禁感叹，人与人的差别... 108 0 4 1 谢小路} {2018-03-10T23:17:32+08:00 『Python 爬虫文集梳理』 过去的几年内，我开始了编程。 过去的一年内，我开始了工作生涯。 我学会的第一个编程技能是『爬虫』，工作后，开始接触Golang。 我开始不断的将编程结合业务， 接触越来越多的... 33 0 5 0 谢小路} {2018-03-06T00:13:44+08:00 『requests-html 源码学习:  1』 大家好，我是谢伟，是一名程序员，熟悉 Pyhton 和 Go。学会的第一个技能是『网络爬虫』。 最近 Python 领域大神  kennethreitz 开源了一个关于网络内... 30 0 3 0 谢小路} {2018-03-04T01:02:06+08:00 使用Python自动登录SSO爬取动态页面 最近的项目在做一些数据方面的集成，五花八门的系统对应的接口更是千奇百怪，数据集成的过程总结成八个字就是：逢山开路，遇水架桥。 恰好这两天碰到一个问题，我们要集成的WEB系统没... 25 0 3 0 半夜菊花茶} {2018-03-05T10:54:55+08:00 Java网络爬虫实操（7） 上一篇：Java网络爬虫实操（6） 大家好，我们平常浏览网页经常会看到这样的效果：鼠标滚动到差不多底部的时候，才会加载新内容出来。然后一直滚就一直加载，比如外卖平台上的评价信... 21 0 1 0 sinkinka} {2018-02-28T16:11:33+08:00 Java网络爬虫实操（6） 上一篇：Java网络爬虫实操（5） 大家好，前几篇文章一直提到用xpath去解析html。由于是演示代码，所以看上去都简洁明了的。其实在生产环境下，我们需要获取的数据往往不是... 36 0 1 0 sinkinka} {2018-03-01T12:54:05+08:00 python爬虫CCTV2《环球财经连线》 笔者前言：平时比较喜欢看CCTV2的《国际财经报道》（原环球财经连线），但是央视的app没有提供缓存功能，笔者又想上下班和跑步的时候看或者听。于是花了大概一天时间写了pyth... 121 1 5 0 大树先生g} {2018-02-24T12:14:35+08:00 Java网络爬虫实操（5） 上一篇：Java网络爬虫实操（4） 大家好，前几篇文章介绍的URL都是返回HTML内容的，然后再从HTML字符串里解析出我们想要的数据。但是，随着前端编程技术的发展，至少十多... 30 0 3 0 sinkinka} {2018-02-18T21:34:33+08:00 Java网络爬虫实操（4） 上一篇：Java网络爬虫实操（3） 本篇文章继续围绕NetDiscovery框架中pipeline的用法，结合另一个专门爬图片的框架PicCrawler，实现图片的批量下载和... 65 3 7 0 sinkinka}]
获取专题最热门文章信息:  [{2016-11-18T13:37:12+08:00 一步一步教你 https 抓包 在 Mac 上常用的抓包软件是 Charles，网上关于 Charles 的教程很多，这里介绍另一个抓包神器 mitmproxy。mitmproxy 是一款可交互式的命令行抓... 16864 119 593 2 hi_xgb} {2016-12-09T22:46:50+08:00 爬虫（1）--- Python网络爬虫二三事 1 前言 作为一名合格的数据分析师，其完整的技术知识体系必须贯穿数据获取、数据存储、数据提取、数据分析、数据挖掘、数据可视化等各大部分。在此作为初出茅庐的数据小白，我将会把自... 7665 48 448 5 whenif} {2016-09-11T22:26:20+08:00 Python网络爬虫实战项目代码大全（长期更新，欢迎补充） WechatSogou[1]- 微信公众号爬虫。基于搜狗微信搜索的微信公众号爬虫接口，可以扩展成基于搜狗搜索的爬虫，返回结果是列表，每一项均是公众号具体信息字典。[1]: h... 7336 12 204 0 Python中文社区} {2017-04-02T12:25:48+08:00 如何一小时爬取百万知乎用户信息，并做简单的可视化分析？ 一、使用的技术栈： 爬虫：python27 +requests+json+bs4+time 分析工具： ELK套件 开发工具：pycharm 二、数据成果 三、简单的可视化分... 4146 26 191 1 方志朋} {2017-01-15T14:37:38+08:00 【170115】2016年，我对爬虫的总结 都说年末了，该给自己写写总结了。今天我想谈一谈的是我在公司这一年多里的负责的部分工作---爬虫。做了这么久的爬虫，是该写点什么，留下点什么。在我所负责的这一段时间了。我总结了... 11198 33 164 2 林湾村龙猫} {2017-07-30T23:10:47+08:00 爬取张佳玮138w+知乎关注者：数据可视化 一、前言 作为简书上第一篇文章，先介绍下小背景，即为什么爬知乎第一大V张公子的138w+关注者信息？ 其实之前也写过不少小爬虫，按照网上各种教程实例去练手，“不可避免”的爬过... 5554 32 130 3 Deserts_X} {2017-08-03T09:45:49+08:00 如何写一个简单的分布式知乎爬虫？ 前言 很早就有采集知乎用户数据的想法，要实现这个想法，需要写一个网络爬虫（Web Spider）。因为在学习 python，正好 python 写爬虫也是极好的选择，于是就写... 2898 10 131 0 呓语_yiyu} {2016-12-03T17:12:04+08:00 Python 福利小爬虫，爬取今日头条街拍美女图 先实际感受一下我们要抓取的福利是什么？点击 今日头条，在搜索栏输入街拍 两个字，点开任意一篇文章，里面的图片即是我们要抓取的内容。 可以看到搜索结果默认返回了 20 篇文章，... 8073 43 128 0 追梦人物} {2016-07-28T11:35:57+08:00 爬取简书全站文章并生成 API（一） 简书中的优质文章非常多，而且我非常喜欢 Markdown 这种语法格式，所以想着能不能爬取简书上面的文章，爬取文章之前先带大家来了解下简书整个网站，简书的首页分为“热门（已推... 5439 46 115 1 田飞雨} {2017-06-27T14:57:25+08:00 Python：Scrapy分布式爬虫打造搜索引擎集合篇 -（一）到（八）完整版 Python分布式爬虫打造搜索引擎 基于Scrapy、Redis、elasticsearch和django打造一个完整的搜索引擎网站本教程一共八章：从零开始，直到搭建一个搜索... 3940 21 119 0 天涯明月笙}]

```


---

## ChangeLog

- 使用 Go Module (2019.12.15)
- v1.0.1 (2018.03.13)

增加简书专题接口

- v1.0.0 (2018.03.12) 

增加简书主要的接口





## TODO

- [ ] PhantomJs
- [ ] 增加搜索接口
- [x] 增加命令方式：cli
- [ ] 增加 web 方式：beego
- [ ] 优化输出格式：json
- [ ] 提供输出存储格式：Csv, Txt, SQLite, Mysql, Mongodb

