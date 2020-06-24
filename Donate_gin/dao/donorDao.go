package dao

import (
	"Donate_gin/db"
	"Donate_gin/entity"
)

//获取用户密码，返回到model层进行判断
func GetDonorPswDao(account string) (donorId int,password string,err error) {
	err = db.DB.QueryRow("SELECT donor_id,password FROM donor WHERE account = ?",account).Scan(&donorId,&password)
	return
}

//获取用户信息
func GetDonorDao(donorID int) (donor entity.Donor,err error) {
	donor.DonorID = donorID

	err = db.DB.QueryRow("SELECT account,nickname,name,id_number,cur_residence,city,avatar,love_value,profile " +
		"FROM donor WHERE donor_id = ?",donorID).Scan(
		&donor.Account,&donor.Nickname,&donor.Name,&donor.IdNumber,&donor.CurResidence,&donor.City,&donor.Avatar,&donor.LoveValue,&donor.Profile)

	return
}

