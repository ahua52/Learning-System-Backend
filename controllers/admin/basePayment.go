package admin

import (
	"HMS/payment/models"
	"errors"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type payment interface {
	Pay()
}

type BasePaymentControl struct {
}

func (con BasePaymentControl) SignIn(ctx *gin.Context) {
	// console.log(ctx)
	// fmt.Println(1111, ctx)
	ctx.JSON(http.StatusOK, gin.H{
		"message":            "success",
		"paymentReferenceId": rand.Intn(10000000),
	})
}

func (con BasePaymentControl) Pay(c *gin.Context) {
	_, payeeId, amount, err := con.getParameters(c)
	if err != nil {
		return
	}
	verifyerr := con.VerifyNumver(c, payeeId, amount)
	if verifyerr != nil {
		return
	}
	payMethod := c.Query("payMethod")
	/*
	* {'DEBIT_CARD':debit, 'CREDIT_CARD':credit, 'MOBILE_PAY':mobile}
	 */
	if payMethod == "DEBIT_CARD" {
		DebitCardAccountController{}.Pay(c)
	} else if payMethod == "CREDIT_CARD" {
		CreditCardAccountController{}.Pay(c)
	} else if payMethod == "MOBILE_PAY" {
		MobileAccountController{}.Pay(c)
	} else {
		c.String(http.StatusOK, "wrong payment method")
	}
}
func (con BasePaymentControl) getParameters(c *gin.Context) (int64, int64, float32, error) {
	transferorId, _ := strconv.ParseInt(c.Query("transferorId"), 10, 64)
	payeeId, _ := strconv.ParseInt(c.Query("payeeId"), 10, 64)
	value, err := strconv.ParseFloat(c.Query("amount"), 32)
	amount := float32(value)
	if err != nil || amount <= 0 {
		// do something sensible
		con.error(c, "wrong amount")
		return 0, 0, 0, errors.New("Wrong amount")
	}
	return transferorId, payeeId, amount, nil
}

func (con BasePaymentControl) VerifyNumver(c *gin.Context, payeeId int64, amount float32) error {
	// verify payeeId
	var u2 *models.Debit
	u3 := models.DB.First(&u2, models.DB.Where("Accountnumber = ?", payeeId))

	if u3.Error != nil {
		con.error(c, "payee not found")
		return errors.New("payee not found")
	}
	return nil
}

func (con BasePaymentControl) success(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message":            "success",
		"paymentReferenceId": rand.Intn(10000000),
	})
}
func (con BasePaymentControl) error(c *gin.Context, errInfo string) {
	if errInfo == "" {
		errInfo = "error"
	}
	c.JSON(http.StatusOK, gin.H{
		"message": errInfo,
	})
}
