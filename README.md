# Jianshu-go


本项目尝试解析简书API。

- 编程语言：golang

主要的接口包括：

- User: 个人主页信息
- Article : 某篇文章的信息
- Home-page: 简书主页的信息
- Home-page-recommend: 简书推荐作者的信息
- Home-page-topic: 简书推荐的专题信息
- Publication： 简书出版信息


主要包括上面 6 类：


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

## Article

- GetAuthor 获取文章作者
- GetDescription 获取作者简介
- GetTitle 获取文章标题
- GetContent 获取文章全文


![article.png](https://upload-images.jianshu.io/upload_images/1818135-f41e684241267c67.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


## Home-page 

- GetHomePagePassages 获取首页文章
- GetNewList 获取首页新上榜
- GetHotSeven 获取首页 7 日热门
- GetHotMonth 获取首页 30 日热门
- GetJianshuSchool 获取首页简书大学堂

![home-page.png](https://upload-images.jianshu.io/upload_images/1818135-dc42bcdd5286ec68.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)



## Home-page-recommend

- GetListRecommendAuthor 获取首页推荐作者

![recommendAuthor.png](https://upload-images.jianshu.io/upload_images/1818135-0c31b9513039cbc1.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)



## Home-page-topic

- GetTopicCollectionRecommend 获取推荐专题
- GetTopicCollectionHot 获取热门专题
- GetTopicCollectionCity 获取城市专题
- GetTopicCollectionSchoolyard 获取校园专题


![topic.png](https://upload-images.jianshu.io/upload_images/1818135-10880c9494a1afaa.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)




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



---

捐赠：

如果项目对你有帮助，欢迎打赏，鼓励作者出更好的开源项目。





---

TODO

- [ ] PhantomJs
- [ ] Search
- [ ] Learn JavaScript
- [x] 增加命令方式：cli
- [ ] 增加 web 方式：beego

