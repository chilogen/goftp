package auth

import (
	"errors"
	"github.com/goftp/server/config"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"time"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Auth struct {
	ipWhiteList []string
	ipBlackList []string
}

type userModel struct {
	Id            int64     `gorm:"column:id" json:"id"`
	UserName      string    `gorm:"column:user_name" json:"user_name"`
	Password      string    `gorm:"column:password" json:"password"`
	CreatorId     int64     `gorm:"column:creator_id" json:"creator_id"`
	UserType      int64     `gorm:"column:user_type" json:"user_type"`
	AvailableFlag int64     `gorm:"column:available_flag" json:"available_flag"`
	Email         string    `gorm:"column:email" json:"email"`
	UpdateTime    time.Time `gorm:"column:update_time" json:"update_time"`
	CreateTime    time.Time `gorm:"column:create_time" json:"create_time"`
}

type ipListModel struct {
	Id            int64     `gorm:"column:id" json:"id"`
	IPType        int64     `gorm:"column:ip_type" json:"ip_type"`
	Version       int64     `gorm:"column:version" json:"version"`
	CreatorID     int64     `gorm:"column:creator_id" json:"creator_id"`
	Addr          string    `gorm:"column:addr" json:"addr"`
	AvailableFlag int64     `gorm:"column:available_flag" json:"available_flag"`
	UpdateTime    time.Time `gorm:"column:update_time" json:"update_time"`
	CreateTime    time.Time `gorm:"column:create_time" json:"create_time"`
}

var (
	auth *Auth
	db *gorm.DB
)

func init() {
	var err error
	db,err=gorm.Open("mysql",config.GetConfig().Mysql.Infodb)
	if err!=nil{
		logrus.Error(err)
		return
	}
}


func (auth *Auth)queryUser(userName string)(res userModel,err error){
	resList:=make([]userModel,0)
	db.Table("user").Where(map[string]interface{}{"user_name":userName}).Find(&resList)
	if len(resList)==0{
		return res,errors.New("no such user")
	}
	res=resList[0]
	return res,nil
}

func (auth *Auth)queryIP(addr string)(res []ipListModel){
	res=make([]ipListModel,0)
	db.Table("ip_list").Where(map[string]interface{}{"addr":addr}).Find(&res)
	return
}

func (auth *Auth) CheckUser(userName, passwd string) (userPath string, err error) {
	res,err:=auth.queryUser(userName)
	if err!=nil{
		return "", errors.New("Can not find user : " + userName)
	}
	if res.Password!=passwd{
		return "", errors.New("Wrong password for user : " + userName)
	}
	return "/"+res.UserName, nil
}

//黑白名单策略
func (auth *Auth) CheckIP(ip string) (ok bool, err error) {
	res:=auth.queryIP(ip)

	//默认允许链接（白名单失效）
	if len(res)==0{
		return true,nil
	}
	//黑名单
	if res[0].IPType==1{
		return false,nil
	}
	return true,nil
}

func GetAuth() *Auth {
	return auth
}

