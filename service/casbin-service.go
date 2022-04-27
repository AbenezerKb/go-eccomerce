package service

import (	
	"gin-exercise/db"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)




type CasbinRule struct {
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Ptype string `gorm:"size:512"`
	V0    string `gorm:"size:512"`
	V1    string `gorm:"size:512"`
	V2    string `gorm:"size:512"`
	V3    string `gorm:"size:512"`
	V4    string `gorm:"size:512"`
	V5    string `gorm:"size:512"`
	V6    string `gorm:"size:512"`
	V7    string `gorm:"size:512"`
	V8    string `gorm:"size:512"`
}


// var enforcer *casbin.Enforcer


func Enforcer() (*casbin.Enforcer){

	conn, err := db.ConnectDB()
	if err != nil {		
		return nil
	}

	// ROOT := `root`
	// CUSTOM=`custom`

	adapter, err := gormadapter.NewAdapterByDBWithCustomTable(conn, &CasbinRule{})
	if err != nil {		
		return nil
	}

	enforcer,err := casbin.NewEnforcer(`C:\Users\Administrator\Documents\E-commerce\abenezer-dev-prep\Config\rbac_model.conf`, adapter)
	
	if err != nil {		
		return nil
	}
	
	err =enforcer.LoadPolicy()
	if err!=nil{	
		return nil
	}
	return enforcer

}
