package admin

import (
	"encoding/json"
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/charisworks/charisworks-backend/internal/transaction"
)

func SendPurchasedEmail(transactionDetails transaction.TransactionDetails) {
	data, err := os.ReadFile("../auth_address.json")
	if err != nil {
		log.Fatalf("JSONファイルの読み込みに失敗しました：%v", err)
		return
	}
	authAddress := new([]string)
	err = json.Unmarshal(data, &authAddress)
	if err != nil {
		log.Fatalf("JSONデータの解析に失敗しました：%v", err)
		return
	}
	for _, to := range *authAddress {
		body := `購入が完了しました。`
		body += fmt.Sprintf(`
購入者情報：
名前： %v  
メールアドレス： %v 
住所： %v  `,
			transactionDetails.UserAddress.RealName,
			transactionDetails.Email,
			transactionDetails.UserAddress.Address)
		for i, item := range transactionDetails.Items {
			body += fmt.Sprintf(`
--------------------------
%d 品目：
商品情報：
商品名： %v 
値段： %v 
数量： %v 
合計金額： %v 
			`, i+1, item.Name, item.Price, item.Quantity, item.Price*item.Quantity)
		}
		body += fmt.Sprintf(`
--------------------------
合計売上： %v 
購入日時： %v 
		`, transactionDetails.TotalPrice, transactionDetails.TransactionAt)
		SendEmail(to, "決済完了通知", body)
	}
	{
		body := fmt.Sprintf(`
%v 様 
この度はお買い上げいただき、誠にありがとうございます。
お客様のご注文を確認いたしましたので、ご連絡いたします。
以下に、ご注文の詳細情報を記載いたします。
		`, transactionDetails.UserAddress.RealName)
		for _, item := range transactionDetails.Items {
			body += fmt.Sprintf(`
--------------------------
【ご注文情報】
商品名： %v 
値段： %v 
数量： %v 
			`, item.Name, item.Price, item.Quantity)
		}
		body += fmt.Sprintf(`
--------------------------
合計金額： %v 
購入日時： %v 
		`, transactionDetails.TotalAmount, transactionDetails.TransactionAt)
		SendEmail(transactionDetails.Email, "【決済完了通知】ご購入ありがとうございます。", body)

	}
	log.Print("Email sent successfully!")
}

func SendEmail(to string, subject string, body string) error {
	data, err := os.ReadFile("../email_credentials.json")
	if err != nil {
		log.Fatalf("JSONファイルの読み込みに失敗しました：%v", err)
		return err
	}
	emailCredentials := make(map[string]string)
	err = json.Unmarshal(data, &emailCredentials)
	if err != nil {
		log.Fatalf("JSONデータの解析に失敗しました：%v", err)
		return err
	}
	log.Print(emailCredentials)
	authmail := emailCredentials["auth_mail"]
	from := emailCredentials["mail_from"]
	password := emailCredentials["mail_auth_pass"]
	smtp_server_addr := emailCredentials["smtp_server_addr"]
	smtp_server := emailCredentials["smtp_server"]
	log.Print(authmail, from, password, smtp_server_addr, smtp_server)
	// メールヘッダーの作成
	header := make(map[string]string)
	header["From"] = "CharisWorks本部" + " <" + from + ">"
	header["To"] = to
	header["Subject"] = subject

	// メール本文の作成
	message := ""
	for key, value := range header {
		message += key + ": " + value + "\r\n"
	}
	message += "\r\n" + body
	err = smtp.SendMail(smtp_server_addr,
		smtp.PlainAuth("", authmail, password, smtp_server),
		from, []string{to}, []byte(message))
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Email sent successfully!")
	}
	return nil
}
