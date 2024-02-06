package admin

import (
	"HMS/payment/models"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

type DebitCardAccountController struct {
	BasePaymentControl
}

func (con DebitCardAccountController) Pay(c *gin.Context) {
	fmt.Println("=====debit")
	transferorId, payeeId, amount, _ := con.getParameters(c)
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
	var ul *models.Debit

	tx.Find(&ul, models.DB.Where("Accountnumber = ?", transferorId))
	ul.Balance = ul.Balance - amount
	if err := tx.Save(&ul).Error; err != nil || ul.Balance < 0 {
		tx.Rollback()
		con.error(c, "Payment account abnormal")
		return
	}

	var u3 *models.Debit
	tx.Find(&u3, models.DB.Where("Accountnumber = ?", payeeId))
	u3.Balance = u3.Balance + amount
	if err := tx.Save(&u3).Error; err != nil {
		tx.Rollback()
		con.error(c, "Collection account abnormal")
		return
	}

	tx.Commit()
	con.success(c)
}
func (con DebitCardAccountController) VerifyNumver(c *gin.Context, transferorId int64, amount float32) error {
	// verify transferorId
	var u2 *models.Debit
	u3 := models.DB.First(&u2, models.DB.Where("Accountnumber = ?", transferorId))

	if u3.Error != nil {
		con.error(c, "transferor not found")
		return errors.New("transferor not found")
	}
	return nil
}
