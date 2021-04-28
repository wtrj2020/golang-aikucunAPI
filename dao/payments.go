package dao

import "appApi/models"

func Payments(pay_from string) (payments models.Payments) {
	DB.Find(&payments, "pay_from=?", pay_from)
	return
}
