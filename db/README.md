#db

* 高效的sql表达式翻译器,对于复杂的查询条件,可以简化sql.
* 内置多种表达式,满足大多数需求.
* 支持大多数数据库,oracle,mysql,sqlite,postgres


## 示例

##### 示例数据表(students)

| id     | name  | gender| id_num | college | major | create_at |
| ----- | ------ |-------------|--------|---------|-------|-----------|
| 20180101 | 杜子腾 |男|158177199901044792|计算机学院|计算机科学与工程|2019-09-01 08:23:34|
| 20180102 | 杜琦燕 |女|151008199801178529|外国语学院|英语|2019-09-02 14:53:04|
| 20180103 | 范统 |男|17156319980116959X|计算机学院|计算机科学与工程|2019-09-03 09:00:09|
| 20180104 | 史珍香 |女|141992199701078600|外国语学院|日语|2019-09-04 10:23:34|



##### 创建db对象
```
    db, err := NewDB("mysql","root:123456@tcp(127.0.0.1:3306)/go?charset=utf8mb4&parseTime=True&loc=PRC",20,10,600)
    if err != nil {
        panic(err)
    }
    
    // user db to do
    .....
```

##### `@` 
替换为参数化字符
> 比如我们现在要查询 gender=男 且 college = 计算机学院 的学生姓名

    db.Query(`select name from students where gender=@gender and college = @college`,map[string]interface{}{
        "gender": "男",
        "college": "计算机学院",
    })

> 那么我们的解析器会将sql和参数解析为
>> sql: `select name from students where gender=? and college = ? `
>> args:` [男 计算机学院]`

##### `#`
原样替换为指定值，值为空时返回 NULL

> 我们现在使用in查询,查询 id 在 `[20180101,20180103,20180105]` 中的学生姓名

    db.Query(`select name from students where id in #ids`,map[string]interface{}{
       "ids":"(20180101,20180103,20180105)",
    })

* 那么我们的解析器会将sql解析为:
* sql: `select name from students where id in (20180101,20180103,20180105)`
* 如果 `ids` 参数没有传,或者为空. sql将解析为: `select name from students where id in NULL` ,这会导致sql执行错误

##### `$` 
原样替换为指定值，值为空时返回 ""

* 同上, 如果 `ids` 参数没有传,或者为空.sql将解析为: `select name from students where id in ` ,这会导致sql执行错误

##### `[` 
`>=` 表达式,值为空时返回"",否则返回: and name >= value
> 我们现在要查询创建时间在 9月2号到9月4号之间的女同学学生姓名

    db.Query(`select name from students where gender = @gender [start_time ]end_time`,map[string]interface{}{
        "gender":"女",
       "start_time":"2019-09-02 00:00:00",
       "end_time":"2019-09-04 23:59:59",
    })
    
* 那么我们的解析器会将sql解析为:
* sql: `select name from students where gender = ? and start_time>=? and end_time<=?` args: `[女 2019-09-02 00:00:00 2019-09-04 23:59:59]`
* 如果 `end_time` 参数没有传,或者为空. sql将解析为:   `select name from students where gender = ? and start_time>=? ` args: `[女 2019-09-02 00:00:00]`

##### `]` 
`<=` 表达式,值为空时返回"",否则返回: and name <= value

* 使用方法同上

##### `~`
检查值，值为空时返回"",否则返回: , name=value
> 当我们要更新 `杜子腾` 的专业时,而性别又是一个可选更新项目

    db.Query(`update students set major = @major ~gender where name = @name`,map[string]interface{}{
       "gender":"女",
       "major":"网络工程",
       "name":"杜子腾",
    })

* 那么我们的解析器会将sql解析为:
* sql: `update students set major = ? ,gender=? where name = ?` args: `[网络工程 女 杜子腾]`
* 如果 `gender` 参数没有传,或者为空. sql将解析为: `update students set major = ? where name = ?` args: `[网络工程 杜子腾]`

##### `&` 
and 条件表达式，检查值，值为空时返回"",否则返回: and name=value
> 要查询计算机学院的学生姓名,但是性别是可选条件时
    
    db.Query(`select name from students where collage = @collage &gender`,map[string]interface{}{
       "gender":"女",
       "collage":"计算机学院",
    })

* 那么我们的解析器会将sql解析为:
* sql: `select name from students where collage = ? and gender=?` args: `[计算机学院 女]`
* 如果 `gender` 参数没有传,或者为空. sql将解析为: `select name from students where collage = ?` args:`[计算机学院]`

##### `|` 
or 条件表达式，检查值，值为空时返回"", 否则返回: or name=value

* 使用方法同上

##### `?` 
like 条件表达式
> 你模糊记得有一个叫`珍香`的女生,但是具体叫什么记不起了,可用使用模糊查询

    db.Query(`select * from students where gender = @gender ?name`,map[string]interface{}{
           "gender":"女",
           "name":"珍香",
    })
    
* 那么我们的解析器会将sql解析为:
* sql: `select * from students where gender = ? and name like '%?%'` args `[女 珍香]`
* 如果 `name` 参数没有传,或者为空. sql将解析为: `select * from students where gender = ?` args: `[女]`



## 复杂sql

* 对于复杂sql,分页,join,子查询,别名,都有比较好的支持.


