# pg
用于golang database/sql 的PostgreSQL驱动

[![GoDoc](https://godoc.org/github.com/blusewang/pg?status.svg)](https://godoc.org/github.com/blusewang/pg)
[![Build Status](https://travis-ci.org/blusewang/pg.svg?branch=master)](https://travis-ci.org/blusewang/pg)
[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://github.com/blusewang/pq/blob/master/LICENSE)

## 安装

	go get github.com/blusewang/pq

## 使用
```golang

	db, err := sql.Open("pg", "postgresql://user:password@host_name/database_name?application_name=app_name")
	if err != nil {
		return err
	}
	defer db.Close()
	rows, err := db.Query("select * from bluse where id>$1", 0)
	if err != nil {
		return err
	}
	...

```

## 文档

更多的细节及使用示例，参见： <https://godoc.org/github.com/blusewang/pg>.


## 特性

* 在Scan() null值时，不需要做额外处理，不会报错。不会因为使用 sql.NullString 等类型而破坏结构。
* 常见Array类型直接兼容golang的数组类型。如PG的：integer[]，对应golang的：[]int64
* 数据源格式，既支持键值对，又支持URI。格式遵守：[PG官方规范](https://www.postgresql.org/docs/10/libpq-connect.html#LIBPQ-CONNSTRING)
* 在同一个连接中，使用同一个sql多次sql.Prepare的场景下，自动复用可用的资源。这对web类业务很有用！

## 协议实现
- 此驱动更适合服务于Web

| 状态   | 功能 | 备注 |
| :---- | :---- | :---- |
| <ul><li>- [x] </li></ul> | 启动 | 必备 |
| <ul><li>- [x] </li></ul> | 简单查询 | 必备 |
| <ul><li>- [x] </li></ul> | 扩展查询 | 必备 |
| <ul><li>- [x] </li></ul> | 取消正在处理的请求 | 必备 |
| <ul><li>- [x] </li></ul> | 终止 | 必备 |
| <ul><li>- [ ] </li></ul> | 函数调用 | PG官方推荐使用查询去调用函数 |
| <ul><li>- [ ] </li></ul> | 异步操作 | 实用性低 |
| <ul><li>- [ ] </li></ul> | COPY操作 | 实用性低 |
| <ul><li>- [ ] </li></ul> | SSL会话加密 | 低效 |
| <ul><li>- [ ] </li></ul> | 流复制协议 | 实用性低 |
| <ul><li>- [ ] </li></ul> | 逻辑流式复制协议 | 实用性低 |
