# learning-go 学起来吧！

> 学习golang没有乱七八糟的理由，仅仅是因为它的native编译无任何依赖

一个菜鸟的自我修养，从golang学习开始……
每个文件基本都代表了go语言在传统项目中最常见的一种用法……



## 01-hello-world

程序员入门必备

（可能我不是一个真的程序员，编号从01开始）



## 02-http

如何用go实现一个简单的HTTP Server，简直不要太简单！

想想Java的Servlet-api还要引入包……



## 03-database

用go语言访问MySQL数据库。golang仅定义了SQL接口，没有实现。

这份代码用到`github.com/go-sql-driver/mysql`和`github.com/jmoiron/sqlx`的第三方实现，主要目的还是为了使用方便。



## 04-json

HTTP通信中常用的json序列化与反序列化。golang自带实现。

这份代码中初步体会了golang的反射用法（struct中tag部分）
