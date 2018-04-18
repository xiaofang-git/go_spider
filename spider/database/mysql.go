package database

import (
	"github.com/gohouse/gorose"
	"os"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"bufio"
	"strings"
)


func Connect() gorose.Connection {
	var host, port, username, password string
	input := bufio.NewReader(os.Stdin)
	fmt.Println("请输入mysql连接地址")
	host, _ = input.ReadString('\n')
	fmt.Println("请输入mysql连接端口")
	port, _ = input.ReadString('\n')
	fmt.Println("请输入mysql连接用户名")
	username, _ = input.ReadString('\n')	
	fmt.Println("请输入mysql连接密码")
	password, _ = input.ReadString('\n')

	need_replace := " \n"
	host = strings.Trim(host, need_replace)
	port = strings.Trim(port, need_replace)
	username = strings.Trim(username, need_replace)
	password = strings.Trim(password, need_replace)

	var dbConfig = map[string]interface{} {
		"Default":         "mysql_dev",// 默认数据库配置
		"SetMaxOpenConns": 300,          // (连接池)最大打开的连接数，默认值为0表示不限制
		"SetMaxIdleConns": 10,          // (连接池)闲置的连接数, 默认1
	
		"Connections":map[string]map[string]string{
			"mysql_dev": {// 定义名为 mysql_dev 的数据库配置
				"host": host, // 数据库地址
				"username": username,       // 数据库用户名
				"password": password,       // 数据库密码
				"port": port,            // 端口
				"database": "spider",        // 链接的数据库名字
				"charset": "utf8",         // 字符集
				"protocol": "tcp",         // 链接协议
				"prefix": "",              // 表前缀
				"driver": "mysql",         // 数据库驱动(mysql,sqlite,postgres,oracle,mssql)
			},
			"sqlite_dev": {
			 "database": "./foo.db",
			 "prefix": "",
			 "driver": "sqlite3",
			},
		},
	}

	conn, err := gorose.Open(dbConfig)
	if err != nil {
		fmt.Println(err)

	}
	return conn
	
}