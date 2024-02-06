package admin

import (
	"HMS/payment/models"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

type CreditCardAccountController struct {
	BasePaymentControl
}

func (con CreditCardAccountController) Pay(c *gin.Context) {
	fmt.Println("-----credit")
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
	// transferor going to transfer mondy to someone
	var ul *models.Credit
	tx.Find(&ul, models.DB.Where("Creditnumber = ?", transferorId))
	ul.Amount = ul.Amount + amount
	if err := tx.Save(&ul).Error; err != nil {
		tx.Rollback()
		con.error(c, "Payment account abnormal")
		return
	}
	//payee get incremented money amount
	var u2 *models.Debit
	tx.Find(&u2, models.DB.Where("Accountnumber = ?", payeeId))
	u2.Balance = u2.Balance + amount
	if err := tx.Save(&u2).Error; err != nil {
		tx.Rollback()
		con.error(c, "Collection account abnormal")
		return
	}

	tx.Commit()
	con.success(c)
}

func (con CreditCardAccountController) VerifyNumver(c *gin.Context, transferorId int64, amount float32) error {
	// verify transferorId
	var u2 *models.Credit
	u3 := models.DB.First(&u2, models.DB.Where("Creditnumber = ?", transferorId))

	if u3.Error != nil {
		con.error(c, "transferor not found")
		return errors.New("transferor not found")
	}
	return nil
}
