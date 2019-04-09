package common

import (
	"database/sql"
	"github.com/Deansquirrel/goMonitorV4/global"
	log "github.com/Deansquirrel/goToolLog"
	"github.com/Deansquirrel/goToolMSSql"
)

type Common struct {
}

//获取配置库连接配置
func (c *Common) getConfigDBConfig() *goToolMSSql.MSSqlConfig {
	return &goToolMSSql.MSSqlConfig{
		Server: global.SysConfig.ConfigDBConfig.Server,
		Port:   global.SysConfig.ConfigDBConfig.Port,
		DbName: global.SysConfig.ConfigDBConfig.DbName,
		User:   global.SysConfig.ConfigDBConfig.User,
		Pwd:    global.SysConfig.ConfigDBConfig.Pwd,
	}
}

func (c *Common) GetRowsBySQL(sql string, args ...interface{}) (*sql.Rows, error) {
	conn, err := goToolMSSql.GetConn(c.getConfigDBConfig())
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	if args == nil {
		rows, err := conn.Query(sql)
		if err != nil {
			log.Error(err.Error())
			return nil, err
		}
		return rows, nil
	} else {
		rows, err := conn.Query(sql, args...)
		if err != nil {
			log.Error(err.Error())
			return nil, err
		}
		return rows, nil
	}
}

func (c *Common) SetRowsBySQL(sql string, args ...interface{}) error {
	conn, err := goToolMSSql.GetConn(c.getConfigDBConfig())
	if err != nil {
		log.Error(err.Error())
		return err
	}
	if args == nil {
		_, err = conn.Exec(sql)
		if err != nil {
			log.Error(err.Error())
			return err
		}
		return nil
	} else {
		_, err := conn.Exec(sql, args...)
		if err != nil {
			log.Error(err.Error())
			return err
		}
		return nil
	}
}
