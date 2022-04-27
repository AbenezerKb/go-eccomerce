package service

import (
	//	_ "github.com/go-sql-driver/mysql"
	//	Config "stripe.com/s/src/config"
	"context"
	 "gin-exercise/db"
	. "gin-exercise/entity"
)

func SavePayment(ctx context.Context, charge *Charge) error { //(err error) {
	//if err = Config.
	err :=db.PaymentSave(ctx,*charge) //.Error; err != nil {
	
	if err !=nil{
		return err
	}	// 	return err
	// }
	return nil

}
