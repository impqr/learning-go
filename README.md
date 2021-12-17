# learning-go 学起来吧！

> 学习golang没有乱七八糟的理由，仅仅是因为它的native编译无任何依赖

一个菜鸟的自我修养，从golang学习开始……
每个文件基本都代表了go语言在传统项目中最常见的一种用法……



## 01_hello-world

程序员入门必备

（可能我不是一个真的程序员，编号从01开始）



## 02_http

如何用go实现一个简单的HTTP Server，简直不要太简单！

想想Java的Servlet-api还要引入包……



## 03_database

用go语言访问MySQL数据库。golang仅定义了SQL接口，没有实现。

这份代码用到`github.com/go-sql-driver/mysql`和`github.com/jmoiron/sqlx`的第三方实现，主要目的还是为了使用方便。



## 04_json

HTTP通信中常用的json序列化与反序列化。golang自带实现。

这份代码中初步体会了golang的反射用法（struct中tag部分）



## 05_elasticsearch

对ElasticSearch的访问，用到`github.com/elastic/go-elasticsearch/v7`

（我应该把这部分放到database的）



## 06_files

文件读写。这里演示了从本地的日志文件中读取，数据清洗后写入MySQL数据库。



## 07_jwt

JWT的go语言实现。

这里抛弃了JWT标准中的`Audience / ExpiresAt / Id / IssuedAt / Issuer / NotBefore / Subject`参数，而是额外定义了`Payload`对象用来记录自定义数据。



## 08_snowflake

雪花算法的实现。

从网络中找了一份资料参考（抄袭）：[https://www.cnblogs.com/dwxt/p/12876041.html](https://www.cnblogs.com/dwxt/p/12876041.html) ，在其基础上做了一定的修改，如全局共用一个雪花ID生成器对象等。

（侵权麻烦告知删除）



## 09_algorithm

RSA秘钥对生成和加解密示例。