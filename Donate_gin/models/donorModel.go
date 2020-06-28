package models

import (
	"Donate_gin/dao"
	"Donate_gin/util"
	"errors"
	"fmt"
)

//捐赠方登录
func DonorLoginModel(account string,password string)(donormap map[string]string,err error)  {
	donorID ,password_, err := dao.GetDonorPswDao(account)
	//todo 加密模块
	if err != nil{
		return
	}

	donormap = make(map[string]string)
	if password == password_{
		donor,_ := dao.GetDonorDao(donorID)

		donormap["account"] = donor.Account
		donormap["nickname"] = donor.Nickname
		donormap["id_number"] = donor.IdNumber
		donormap["cur_residence"] = donor.CurResidence
		donormap["city"] = donor.City
		donormap["avatar"] = donor.Avatar
		donormap["love_value"] = donor.LoveValue
		donormap["profile"] = donor.Profile
		return

	}else {
		return nil,errors.New("账号或密码错误")
	}

}

var Code map[string]string

func SendSMSModel(account string) (err error)  {
	donorID,_ := dao.GetDonorAccountDao(account)
	if donorID == 0{
		code := util.GenValidateCode(6)
		fmt.Printf("验证码为:%v\n:",code)
		Code[account] = code
		return nil
	}else {
		return errors.New("账号已存在")
	}
}


//发送验证码
//获取验证码
//判断验证码
func DonorRegister(account string,code string,pwd string,name string,IdNumber string,city string) (donationID int64,err error) {
	gencode := Code[account]

	if code == gencode{
		//todo password加密
		password := pwd
		donationID ,err = dao.AddDonorDao(account,password,name,IdNumber,city)
		if err!= nil{
			return 0,err
		}
		return
	}else {
		return 0,errors.New("验证码错误")
	}

}
