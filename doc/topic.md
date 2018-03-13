## HOME-PAGE-TOPIC


### 接口：

- GetTopicCollectionRecommend 获取推荐专题
- GetTopicCollectionHot 获取热门专题
- GetTopicCollectionCity 获取城市专题
- GetTopicCollectionSchoolyard 获取校园专题

### 示例:

```go
package main

import (
	"fmt"
	"jianshu-go"
)


func main(){
    topic := jianshu.NewTopic(12)
    fmt.Println("获取推荐专题 12 个: ", topic.GetTopicCollectionRecommend())
    fmt.Println("获取热门专题 12 个: ", topic.GetTopicCollectionHot())
    fmt.Println("获取城市专题 12 个: ", topic.GetTopicCollectionCity())
    fmt.Println("获取校园专题 12 个: ", topic.GetTopicCollectionSchoolyard())
}


```


``` 
获取推荐专题 12 个:  [{故事 故事专题，不论是旅行生活中亲身经历的真实故事，还是童话玄幻遐想的虚构故... 162698篇文章 · 764.2K人关注} {旅行·在路上 邂逅一座城池，讲述一段故事，分享你的旅行。欢迎来到「旅行·在路上」专题... 97891篇文章 · 1887.3K人关注} {读书 投稿须知：读书专题仅收录与读书有关的书评、读书笔记、阅读方法、读书... 163277篇文章 · 2565.6K人关注} {人文社科 欢迎关注公众号简宝玉(公众号ID:jianshu4321)后台... 7168篇文章 · 244.5K人关注} {摄影 欢迎关注公众号简宝玉(公众号ID:jianshu4321)后台... 57253篇文章 · 1408.6K人关注} {历史 本专题主要收录历史阐述与分析类文章，也收录历史小说等带一定想象性的文章... 42474篇文章 · 1432.5K人关注} {简书电影 欢迎关注公众号简宝玉(公众号ID:jianshu4321)后台... 77310篇文章 · 1274.9K人关注} {社会热点 欢迎关注公众号简宝玉(公众号ID:jianshu4321)后台... 44300篇文章 · 1242.6K人关注} {@IT·互联网 欢迎关注公众号简宝玉(公众号ID:jianshu4321)后台... 42758篇文章 · 1160.9K人关注} {手绘 欢迎关注公众号简宝玉(公众号ID:jianshu4321)后台... 94646篇文章 · 1194.4K人关注} {自然科普 本专题仅收涉及自然科学（生物、天文、物理等）的科普、介绍、发现类文章，... 3150篇文章 · 665.7K人关注} {短篇小说 借一斑略知全豹，以一目尽传精神。阅读简书最好看的故事，欢迎订... 68914篇文章 · 1758.8K人关注}]
获取热门专题 12 个:  [{首页投稿 玩转简书的第一步，从这个专题开始。想上首页热门榜么？好内容想被... 384659篇文章 · 193.6K人关注} {连载小说 欢迎关注公众号简宝玉(公众号ID:jianshu4321)后台... 134282篇文章 · 945.4K人关注} {故事 故事专题，不论是旅行生活中亲身经历的真实故事，还是童话玄幻遐想的虚构故... 162698篇文章 · 764.2K人关注} {程序员 欢迎关注公众号简宝玉(公众号ID:jianshu4321)后台... 81901篇文章 · 708.3K人关注} {旅行·在路上 邂逅一座城池，讲述一段故事，分享你的旅行。欢迎来到「旅行·在路上」专题... 97891篇文章 · 1887.3K人关注} {短篇小说 借一斑略知全豹，以一目尽传精神。阅读简书最好看的故事，欢迎订... 68914篇文章 · 1758.8K人关注} {哲思 投稿须知：1、字数不限，收入以哲学思辨（中西方哲学观点和作者本人的... 11672篇文章 · 1094.2K人关注} {摄影 欢迎关注公众号简宝玉(公众号ID:jianshu4321)后台... 57253篇文章 · 1408.6K人关注} {@IT·互联网 欢迎关注公众号简宝玉(公众号ID:jianshu4321)后台... 42758篇文章 · 1160.9K人关注} {人物 投稿须知：■人物专题仅收录以人物为主题的文章。例如：历史人物，... 35483篇文章 · 858.4K人关注} {电竞·游戏 少年，一起来玩游戏吗？投稿须知：1.本专题仅收录关于游戏相... 10177篇文章 · 268.7K人关注} {每天写1000字 提高写作能力，成为更好的自己。练习每天写1000字，打卡提... 324196篇文章 · 107.4K人关注}]
获取城市专题 12 个:  [{陕西省 文明起步的地方，你看，炎黄大帝都笑了，你还在等什么呢？投稿须知... 5659篇文章 · 2.0K人关注} {香港 香港，中国的一扇窗，一方面把自己打扮得珠光宝气，一方面又保持古老传统。... 4081篇文章 · 3.3K人关注} {福建省 在鼓浪屿的迷蒙雨中，闲坐在屋檐下，望着不远处的海发呆，也是一种享受。... 1414篇文章 · 751人关注} {辽宁省 简介：共和国的长子、东北的心脏、美丽的辽宁欢迎您。欢迎加入简书... 673篇文章 · 1.0K人关注} {台湾 酷酷的湾仔，嗲嗲的台妹，旖旎的风光……这就是我们的宝岛台湾！... 832篇文章 · 28.2K人关注} {广东省 传承岭南文化，领略南粤风情，在这里遇见最美好的广东。﻿投稿须... 2107篇文章 · 3.0K人关注} {贵州省 『一次“黔”行，终身难忘』 贵州省，简称黔或贵。 这里有世... 1704篇文章 · 867人关注} {云南省 踏入美丽、丰饶、神奇的云南秘境，回到最初的地方，灵魂、身体、美食，总有... 1392篇文章 · 989人关注} {河南省 还你一个真实、全面的河南。投稿须知：http://www.ji... 1534篇文章 · 1.7K人关注} {江西省 投稿须知：http://www.jianshu.com/p/77a35... 500篇文章 · 552人关注} {湖南省 湘江北去，芙蓉天成。 冒有胆子，就不要搞湖湘文化。你要有胆... 1642篇文章 · 1.5K人关注} {黑龙江省 用龙江的坚毅与柔情，谱写我们的未来。投稿须知：http://w... 417篇文章 · 442人关注}]
获取校园专题 12 个:  [{南昌大学 “未完待续就要表白”全国大学生表白母校活动——南昌大学　　格物致... 370篇文章 · 274人关注} {唐山师范学院 表白母校活动，有你参与更精彩! 118篇文章 · 164人关注} {潍坊科技学院 修身博学求索笃行欢迎大家踊跃投稿 195篇文章 · 82人关注} {齐鲁工业大学 于千万人之中，遇见你所要遇见的人，于千万年之中，时间的无涯的荒野里，没... 119篇文章 · 103人关注} {大学生活 本专题收录大学相关文章。要求文内不得带有公众号信息等推广内容。详细规... 75987篇文章 · 706.2K人关注} {南阳理工学院 今厚德博学南工之学子，书求真至善温暖之文字。投稿须知：... 126篇文章 · 256人关注} {南京师范大学泰州学院  534篇文章 · 479人关注} {安徽师范大学 简书师大——专属于师大人的自由创作社区创作、分享、交流，在这里我... 2619篇文章 · 1.6K人关注} {山西师范大学 山西师范大学（ShanxiNormalUniversity），简称... 83篇文章 · 61人关注} {河北地质大学华信学院 河北地质大学华信学院，位于石家庄市，成立于2001年，是经教育部和河北... 560篇文章 · 418人关注} {河北大学 实事求是，笃学诚行 124篇文章 · 170人关注} {南京工业大学浦江学院 1、在1月5号以前，每周可在简书app上发表一篇读书感悟，字数500字... 74篇文章 · 353人关注}]

```