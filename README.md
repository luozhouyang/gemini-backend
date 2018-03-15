# Gemini-backend
README: [ English ](README_EN.md)|[ 中文 ](README.md)
　　

本项目作为 **[gemini-frontend](https://github.com/luozhouyang/gemini-frontend)** 项目的后端，使用HTTP协议进行通信。　　


## 1.创建数据库
本项目使用mysql数据库存储数据，因此在运行项目之前，你应该创建以下几个数据库。　　
　　
```bash
mysql -u root -p

mysql> create user 'gemini'@'localhost' identified by 'usergemini';
mysql> create database gemini_db;
mysql> grant all privileges on gemini_db.* to gemini@localhost identified by 'usergemini';
mysql>
mysql> create database gemini_db_dev;
mysql> grant all privileges on gemini_db_dev.* to gemini@localhost identified by 'usergemini';
mysql>
mysql> create database gemini_db_test;
mysql> grant all privileges on gemini_db_test.* to gemini@localhost identified by 'usergemini';
mysql>
```

注意:
> 如果你在创建数据库的时候，指定了端口号，那么在项目的配置文件中，也需要加上端口号。　　

## 2.接口约定　　
本项目像前端应用提供数据结构服务，因此需要约定接口。主要按照提供服务的类型来详细解释。　　

接口类型可以分为以下几个模块：
* 文章操作
* 待定

### 2.1文章操作
1.**查询**

查询操作可以按照**author**查询、**date**查询和**title**查询，使用GET请求。　　

|方式|参数|解释|示例|
|:---|:---|:---|:---|
|`author`|`author=$AUTHOR_NAME`|作者名字|`author=luozhouyang`|
|`date`|`date=$DATE`|文章发表日期|`date=20180316`|
|`title`|`title=$TITLE`|文章标题（暂不支持模糊）|`title=HelloWorld`|

以上几种方式可以同时进行，也就是可以多个请求类型参数拼接在一起。　　

2.**添加文章**

添加文章操作使用POST请求，需要提交的表单如下：

|键|值|解释|示例|
|:--|:--|:--|:--|
|`title`|`$TITLE`|文章标题|`HelloWorld`|
|`author`|`$AUTHOR`|文章作者|`luozhouyang`|
|`date`|`$DATE`|日期（可能会有调整）|`20180306`|
|`content`|`$CONTENT`|文章内容|`Hello world from gemini!`|

**日期**可能会有调整，因为mysql数据库存储的不是简单的一个日期。　　

