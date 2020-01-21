package mysqlt

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
)
import (
	"database/sql"
	"fmt"
)

type Connect struct {

	UserName string
	Password string
	Ip       string
	Port     string
	DbName   string
}

func ConnectMysql(connect *Connect) {
	connect.Ip = "192.168.100.73"
	buildPath := strings.Builder{}
	buildPath.WriteString(connect.UserName)
	buildPath.WriteString(":")
	buildPath.WriteString(connect.Password)
	buildPath.WriteString("@tcp(")
	buildPath.WriteString(connect.Ip)
	buildPath.WriteString(":")
	buildPath.WriteString(connect.Port)
	buildPath.WriteString(")/")
	buildPath.WriteString(connect.DbName)
	path := buildPath.String()
	db, err := sql.Open("mysql", path)
	if err != nil {
		fmt.Println("open DB err")
		return
	}
	//设置数据库最大连接数
	db.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	db.SetMaxIdleConns(10)
	//验证连接
	if err := db.Ping(); err != nil {
		fmt.Println("open database fail")
		return
	}
	//rows, err := db.Query("SELECT TABLE_NAME, column_name, DATA_TYPE, column_comment FROM information_schema.COLUMNS WHERE TABLE_SCHEMA = 'nlstore'")
	//查询数据库所有表名
	//("SELECT distinct TABLE_NAME FROM information_schema.COLUMNS WHERE TABLE_SCHEMA = 'nlstore'")
	//查询表字段名称，字段描述，字段类型
	//("select column_name,column_comment,data_type from information_schema.columns where table_name='address' and table_schema='nlstore'")
	querySql := "select column_name,column_comment,data_type from information_schema.columns where table_name='address' and table_schema='nlstore'"
	fmt.Println("connnect success")
	GetDbData(db, querySql)
}

/**
golang连接查询mysql
*/
func GetDbData(db *sql.DB, querySql string) {
	rows, err := db.Query(querySql)
	if err != nil {
		fmt.Println("Query fail 。。。")
	}
	//获取列名
	columns, _ := rows.Columns()

	//定义一个切片,长度是字段的个数,切片里面的元素类型是sql.RawBytes
	values := make([]sql.RawBytes, len(columns))
	//定义一个切片,元素类型是interface{} 接口
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		//把sql.RawBytes类型的地址存进去了
		scanArgs[i] = &values[i]
	}
	//获取字段值
	var result []map[string]string
	for rows.Next() {
		res := make(map[string]string)
		rows.Scan(scanArgs...)
		for i, col := range values {
			res[columns[i]] = string(col)
		}
		result = append(result, res)
	}

	//遍历结果
	for _, r := range result {
		for k, v := range r {
			log.Printf("%s==%s", k, v)
		}
	}
	fmt.Println("Query success")
	rows.Close()
}
