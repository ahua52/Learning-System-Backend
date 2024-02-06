package admin

import (
	"HMS/payment/models"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

type MobileAccountController struct {
	BasePaymentControl
}

func (con MobileAccountController) Pay(c *gin.Context) {
	fmt.Println("+++++mobile")
	transferorId := c.Query("transferorId")
	_, payeeId, amount, _ := con.getParameters(c)
	err := con.VerifyNumver(c, transferorId, amount)
	if err != nil {
		return
	}
	// begin a transaction
	tx := models.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			con.error(c, "failure of transaction")
		}
	}()
	//user transfer money to hospital
	var ul *models.Mobile
	tx.Find(&ul, models.DB.Where("Phonenumber = ?", transferorId))
	ul.Balance = ul.Balance - amount
	if err := tx.Save(&ul).Error; err != nil || ul.Balance < 0 {
		tx.Rollback()
		con.error(c, "Payment account abnormal")
		return
	}
	//hospital account increasedy by amount
	var u2 *models.Debit
	tx.Find(&u2, models.DB.Where("Accountnumber = ?", payeeId))
	u2.Balance = u2.Balance + amount
	if err := tx.Save(&u2).Error; err != nil {
		tx.Rollback()
		con.error(c, "Collection account abnormal1")
		return
	}

	tx.Commit()
	con.success(c)
}
func (con MobileAccountController) VerifyNumver(c *gin.Context, transferorId string, amount float32) error {
	// verify transferorId
	var u2 *models.Mobile
	u3 := models.DB.First(&u2, models.DB.Where("Phonenumber = ?", transferorId))

	if u3.Error != nil {
		con.error(c, "transferor not found")
		return errors.New("transferor not found")
	}
	return nil
}
