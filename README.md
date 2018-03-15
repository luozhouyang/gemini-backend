# Gemini-backend
README: [ English ](README_EN.md)|[ 中文 ](README.md)
　　

本项目作为 **[gemini-frontend](https://github.com/luozhouyang/gemini-frontend)** 项目的后端，使用HTTP协议进行通信。　　


## 创建数据库
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

