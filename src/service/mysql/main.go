package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"lib/config"
	"lib/system"
	"strconv"
	"strings"
	"time"
)

type Mysql string
type Args struct {
	Database string
	Sql      string
	Bind     []interface{}
}
type ResultRead struct {
	Data      AllLineData
	DataLenth int
}
type ResultWrite struct {
	LastInsertId int64
	RowsAffected int64
}
type ResultError struct {
	Error string
}

type SingleLineData map[string]string
type AllLineData []SingleLineData

var CoonPool = make(map[string]*sql.DB)

func getCoon(DatabaseName string) (db *sql.DB, error error) {
	_, ok := CoonPool[DatabaseName]
	if ok == false {
		db, error = sql.Open("mysql",
			config.DatabaseConfig.User+":"+config.DatabaseConfig.PassWord+"@tcp("+config.DatabaseConfig.Host+":"+strconv.Itoa(config.DatabaseConfig.Port)+")/"+DatabaseName)
		db.SetMaxOpenConns(config.DatabaseConfig.MaxOpenConns)
		db.SetMaxIdleConns(config.DatabaseConfig.MaxIdleConns)
		db.SetConnMaxLifetime(time.Duration(config.DatabaseConfig.ConnMaxLifetime) * time.Second)
		CoonPool[DatabaseName] = db
		return db, error
	} else {
		return CoonPool[DatabaseName], nil
	}
}

func (t *Mysql) Query(args *Args, reply *interface{}) error {
	database := args.Database
	if database == "" {
		*reply = ResultError{Error: "缺少查询数据库"}
		return nil
	}
	//检查是否允许该数据库
	AllowDatabase := make([]interface{}, len(config.DatabaseConfig.AllowDatabase))
	for i, v := range config.DatabaseConfig.AllowDatabase {
		AllowDatabase[i] = v
	}
	if system.IndexOf(database, AllowDatabase) == false {
		*reply = ResultError{Error: "不允许访问该数据库"}
		return nil
	}
	//获取参数里的sql
	QuerySql := args.Sql
	if QuerySql == "" {
		*reply = ResultError{Error: "缺少查询语句"}
		return nil
	}

	//链接数据库
	db, err := getCoon(database)
	//defer db.Close()
	if err != nil {
		*reply = ResultError{Error: err.Error()}
		return nil
	}

	//判断读，写模式
	isRead := false
	if strings.HasPrefix(QuerySql, "select") || strings.HasPrefix(QuerySql, "explain") {
		isRead = true
	}

	//开始执行数据库查询
	if isRead == true {
		rows, err := db.Query(QuerySql, args.Bind...)
		if err != nil {
			*reply = ResultError{Error: err.Error()}
			return nil
		}

		defer rows.Close()

		columns, _ := rows.Columns()
		scanArgs := make([]interface{}, len(columns))
		values := make([]interface{}, len(columns))
		for j := range values {
			scanArgs[j] = &values[j]
		}

		allLineData := AllLineData{}
		singleLineData := SingleLineData{}

		for rows.Next() {
			//将行数据保存到record字典
			err = rows.Scan(scanArgs...)
			for i, col := range values {
				if col != nil {
					singleLineData[columns[i]] = string(col.([]byte))
				}
			}
			allLineData = append(allLineData, singleLineData)
		}
		*reply = ResultRead{Data: allLineData, DataLenth: len(allLineData)}
	} else {
		result, err := db.Exec(QuerySql, args.Bind...)

		if err != nil {
			*reply = ResultError{Error: err.Error()}
			return nil
		}
		LastInsertId, _ := result.LastInsertId()
		RowsAffected, _ := result.RowsAffected()
		*reply = ResultWrite{LastInsertId: LastInsertId, RowsAffected: RowsAffected}
	}

	return nil
}
