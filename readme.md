### orm client for tldb in go

------------

See the example at  https://tlnet.top/tlormgo

```go
type UserOrm struct {
        Id      int64
        Name    string  `idx`  //  create an index for field `name`
        Age     int
        Level   byte
        Balance float64
}

   //orm 测试
 func Test_orm(t *testing.T) {
        orm.RegisterDefaultResource(false, "db.tlnet.top:7001", "mycli=123")
        //create table
        orm.Create[UserOrm]()
        //insert data
        seq, _ := orm.Insert(&UserOrm{Name: "tom", Age: 20, Level: 1, Balance: 99.2})
        //select
        u, _ := orm.SelectById[UserOrm](seq)
        logging.Debug(u)
    }


/***
        创建表对应的struct：UserOrm，要求必须有Id int64字段，该字段无需赋值，有数据库自动生成自增序号
        字段 如果加 tag 如：Name string `idx`,则表示对 name字段创建索引，也可以写成： Name string `idx:"1"`,结果一致。
        调用 Create[UserOrm](), 创建表，如果已经存在，则返回已经存在错误码
        调用 AlterTable   修改表结构，如果表不存在，则新建表
        调用 Drop         删除表及表数据
        调用 Insert(&UserOrm{Name: "tom", Age: 20, Level: 1, Balance: 99.2}) 插入数据，Id无需赋值，返回创建的Id值。
        调用 SelectById[UserOrm](seq)    根据Id值返回UserOrm对象
        调用 SelectsByIdLimit[UserOrm]() 根据Id范围值返回UserOrm数组
        调用 SelectByIdx[UserOrm]()      根据索引字段查询 返回 UserOrm数组
        调用 SelectAllByIdx[UserOrm]()   根据索引字段查询 返回 UserOrm数组
        调用 SelectByIdxLimit[UserOrm]() 根据索引字段查询 返回 UserOrm数组
             Update 修改数据
             Delete 删除表数据
***/

```