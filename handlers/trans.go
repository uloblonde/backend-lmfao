package handlers

import (
	dto "backendtask/dto/result"
	transdto "backendtask/dto/transaction"
	"backendtask/models"
	"backendtask/repositories"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"gopkg.in/gomail.v2"
)

type handlerTrans struct { //membuat struct handler dengan isi sebagai berikut
	TransRepository repositories.TransRepository // field UserRepository berisikan package dari repositories dan memanggil si interface UserRepository dari package repositories
}

func HandlersTrans(TransRepository repositories.TransRepository) *handlerTrans {
	return &handlerTrans{TransRepository}
}

func (h *handlerTrans) CariTrans(c echo.Context) error {
	trans, err := h.TransRepository.CariTrans()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, trans)
}

func (h *handlerTrans) DapatTrans(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	trans, err := h.TransRepository.DapatTrans(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: trans})
}

func (h *handlerTrans) MembuatTrans(c echo.Context) error {
	request := new(transdto.TransactionRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	var transactionIsMatch = false
	var transactionId int
	for !transactionIsMatch {
		transactionId = int(time.Now().Unix())
		transactionData, _ := h.TransRepository.DapatTrans(transactionId)
		if transactionData.Id == 0 {
			transactionIsMatch = true
		}
	}
	// data form pattern submit to pattern entity db Transaction
	Transactions := models.Transaction{
		Id:        transactionId,
		StartDate: time.Now(),
		EndDate:   time.Now().Add(time.Hour * 24 * time.Duration(30)),
		UserId:    int(userId),
		Price:     request.Price,
		Status:    "pending",
	}

	dataTransactions, err := h.TransRepository.MembuatTrans(Transactions)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	// 1. Initiate Snap client
	var s = snap.Client{}
	s.New(os.Getenv("SERVER_KEY"), midtrans.Sandbox)
	// Use to midtrans.Production if you want Production Environment (accept real transaction).

	// 2. Initiate Snap request param
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(dataTransactions.Id),
			GrossAmt: int64(dataTransactions.Price),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: dataTransactions.User.FullName,
			Email: dataTransactions.User.Email,
		},
	}

	// 3. Execute request create Snap transaction to Midtrans Snap API
	snapResp, _ := s.CreateTransaction(req)

	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: snapResp})

}

// func (h *handlerTrans) UpdateTrans(c echo.Context) error {
// 	meminta := new(transdto.CreateTransRequest)
// 	if err := c.Bind(meminta); err != nil {
// 		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
// 	}
// 	id, _ := strconv.Atoi(c.Param("id"))

// 	user, err := h.TransRepository.DapatTrans(id)

// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
// 	}
// 	if meminta.StartDate != "" {
// 		user.StartDate = meminta.StartDate
// 	}
// 	if meminta.DueDate != "" {
// 		user.DueDate = meminta.DueDate
// 	}
// 	if meminta.UserId != 0 {
// 		user.UserId = meminta.UserId
// 	}
// 	if meminta.Attache != "" {
// 		user.Attache = meminta.Attache
// 	}
// 	if meminta.Status != "" {
// 		user.Status = meminta.Status
// 	}

// 	data, err := h.TransRepository.UpdateTrans(user, id)

// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
// 	}
// 	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: data})
// }

func (h *handlerTrans) HapusTrans(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	trans, err := h.TransRepository.DapatTrans(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.TransRepository.HapusTrans(trans, id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: data})
}

func (h *handlerTrans) Notification(c echo.Context) error {
	var notificationPayload map[string]interface{}

	if err := c.Bind(&notificationPayload); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	transactionStatus := notificationPayload["transaction_status"].(string)
	fraudStatus := notificationPayload["fraud_status"].(string)
	orderId := notificationPayload["order_id"].(string)
	order_id, _ := strconv.Atoi(orderId)
	transaction, err := h.TransRepository.DapatTrans(order_id)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Print("ini payloadnya", notificationPayload)

	if transactionStatus == "capture" {
		if fraudStatus == "challenge" {
			// TODO set transaction status on your database to 'challenge'
			// e.g: 'Payment status challenged. Please take action on your Merchant Administration Portal
			h.TransRepository.UpdateTrans("pending", order_id)
		} else if fraudStatus == "accept" {
			SendMail("success", transaction)
			// TODO set transaction status on your database to 'success'
			h.TransRepository.UpdateTrans("success", order_id)
		}
	} else if transactionStatus == "settlement" {
		SendMail("success", transaction)
		// TODO set transaction status on your databaase to 'success'
		h.TransRepository.UpdateTrans("success", order_id)
	} else if transactionStatus == "deny" {
		// TODO you can ignore 'deny', because most of the time it allows payment retries
		// and later can become success
		h.TransRepository.UpdateTrans("failed", order_id)
	} else if transactionStatus == "cancel" || transactionStatus == "expire" {
		// TODO set transaction status on your databaase to 'failure'
		h.TransRepository.UpdateTrans("failed", order_id)
	} else if transactionStatus == "pending" {
		// TODO set transaction status on your databaase to 'pending' / waiting payment
		h.TransRepository.UpdateTrans("pending", order_id)
	}

	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: notificationPayload})
}
func SendMail(status string, transaction models.Transaction) {

	if status != transaction.Status && (status == "success") {
		var CONFIG_SMTP_HOST = "smtp.gmail.com"
		var CONFIG_SMTP_PORT = 587
		var CONFIG_SENDER_NAME = "Dumbflix <demo.dumbflix@gmail.com>"
		var CONFIG_AUTH_EMAIL = os.Getenv("EMAIL_SYSTEM")
		var CONFIG_AUTH_PASSWORD = os.Getenv("PASSWORD_SYSTEM")

		var productName = "Subscription Dumbflix"
		var price = strconv.Itoa(transaction.Price)

		mailer := gomail.NewMessage()
		mailer.SetHeader("From", CONFIG_SENDER_NAME)
		mailer.SetHeader("To", transaction.User.Email)
		mailer.SetHeader("Subject", "Transaction Status")
		mailer.SetBody("text/html", fmt.Sprintf(`<!DOCTYPE html>
	  <html lang="en">
		<head>
		<meta charset="UTF-8" />
		<meta http-equiv="X-UA-Compatible" content="IE=edge" />
		<meta name="viewport" content="width=device-width, initial-scale=1.0" />
		<title>Document</title>
		<style>
		  h1 {
		  color: brown;
		  }
		</style>
		</head>
		<body>
		<h2>Product payment :</h2>
		<ul style="list-style-type:none;">
		  <li>Name : %s</li>
		  <li>Total payment: Rp.%s</li>
		  <li>Status : <b>%s</b></li>
		</ul>
		</body>
	  </html>`, productName, price, status))

		dialer := gomail.NewDialer(
			CONFIG_SMTP_HOST,
			CONFIG_SMTP_PORT,
			CONFIG_AUTH_EMAIL,
			CONFIG_AUTH_PASSWORD,
		)

		err := dialer.DialAndSend(mailer)
		if err != nil {
			log.Fatal(err.Error())
		}

		log.Println("Mail sent! to " + transaction.User.Email)
	}
}

// func convertResTrans(u models.Transaction) transdto.CreateTransResponse {
// 	return transdto.CreateTransResponse{
// 		Id:        u.Id,
// 		StartDate: u.StartDate,
// 		DueDate:   u.DueDate,
// 		UserId:    u.UserId,
// 		Attache:   u.Attache,
// 		Status:    u.Status,
// 	}
// }
