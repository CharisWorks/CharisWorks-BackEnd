package admin

import (
	"encoding/json"
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/charisworks/charisworks-backend/internal/transaction"
	"github.com/charisworks/charisworks-backend/internal/utils"
	"github.com/charisworks/charisworks-backend/validation"
)

func SendPurchasedEmail(transactionDetails transaction.TransactionDetails, firebaseApp validation.IFirebaseApp) {
	data, err := os.ReadFile("./auth_address.json")
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

注文ID： %v

--------------------------

【ご注文情報】
商品名		値段		数量
		`, transactionDetails.UserAddress.RealName, transactionDetails.TransactionId)
		for _, item := range transactionDetails.Items {
			body += fmt.Sprintf(`
%v		%v円		%v個
			`, item.Name, item.Price, item.Quantity)
		}
		body += fmt.Sprintf(`
送料： %v円
合計金額： %v 円
購入日時： %v 
		`, 350, transactionDetails.TotalPrice+350, utils.ConvertToJST(transactionDetails.TransactionAt))
		body += fmt.Sprintf(`
--------------------------

【お届け先】
お名前： %v様
住所： %v`, transactionDetails.UserAddress.RealName, transactionDetails.UserAddress.Address)
		body +=
			`

--------------------------

商品の発送準備が整いましたら、別途メールにてご連絡いたします。通常、商品の発送には、2,3日程度かかりますので、ご了承ください。

ご質問やご不明な点がございましたら、いつでもお気軽にお客様相談室からお問い合わせください。お手続きや配送に関する詳細情報は、ご注文IDを教えていただくとスムーズに対応できます。

また、商品の受け取り後に何かお気づきの点やご意見がございましたら、お知らせいただけると幸いです。お客様のご意見は、弊社のサービス向上につながりますので、ぜひお聞かせください。

改めまして、ご購入いただきありがとうございます。今後とも、より良い商品とサービスをご提供できるよう努めてまいりますので、どうぞよろしくお願い申し上げます。


--------------------------
CharisWorks

お客様相談室:contact@charis.works

`
		SendEmail(transactionDetails.Email, "【決済完了通知】ご購入ありがとうございます。", body)

	}
	{
		margin := 0.05
		itemsByManufacturer := make(map[string][]transaction.TransactionItem)
		for _, item := range transactionDetails.Items {
			itemsByManufacturer[item.ManufacturerUserId] = append(itemsByManufacturer[item.ManufacturerUserId], item)
		}

		for manufacturerUserId, item := range itemsByManufacturer {
			body := fmt.Sprintf(`
%v 様
出品された商品が購入されました。
以下に、ご注文の詳細情報を記載いたします。

--------------------------

【商品情報】
商品名		値段		数量
`, item[0].ManufacturerName)
			email, err := firebaseApp.GetEmail(manufacturerUserId)
			if err != nil {
				log.Fatalf("メールアドレスの取得に失敗しました：%v", err)
				return
			}
			benefit := 0.0
			for _, item := range item {

				body += fmt.Sprintf(`
%v		%v円		%v個
`, item.Name, item.Price, item.Quantity)
				benefit += float64(item.Price) * float64(item.Quantity) * float64(1-margin)
			}

			body += fmt.Sprintf(`
売上： %v 円
購入日時： %v
`, benefit, utils.ConvertToJST(transactionDetails.TransactionAt))
			body += `
--------------------------
CharisWorks

お客様相談室:contact@charis.works
`
			SendEmail(email, "商品が購入されました", body)

		}
	}
	log.Print("Email sent successfully!")
}

func SendShippedEmail(transactionDetails transaction.TransactionDetails) {
	data, err := os.ReadFile("./auth_address.json")
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
		body := `発送が完了しました。`
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
		SendEmail(to, "発送完了通知", body)
	}
	{
		body := fmt.Sprintf(`
%v 様 


この度はお買い上げいただき、誠にありがとうございます。
お客様から注文のあった商品を発送しましたので、ご連絡いたします。
以下に、ご注文の詳細情報を記載いたします。

注文ID： %v
追跡番号： %v
--------------------------

【ご注文情報】
商品名		値段		数量
		`, transactionDetails.UserAddress.RealName, transactionDetails.TransactionId, transactionDetails.TrackingId)
		for _, item := range transactionDetails.Items {
			body += fmt.Sprintf(`
%v		%v円		%v個
			`, item.Name, item.Price, item.Quantity)
		}
		body += fmt.Sprintf(`
送料： %v円
合計金額： %v 円
購入日時： %v 
		`, 350, transactionDetails.TotalPrice+350, utils.ConvertToJST(transactionDetails.TransactionAt))
		body += fmt.Sprintf(`
--------------------------

【お届け先】
お名前： %v様
住所： %v`, transactionDetails.UserAddress.RealName, transactionDetails.UserAddress.Address)
		body +=
			`

--------------------------

商品の返品・返金に致しましては、商品到着後7日以内にお問い合わせフォームよりご連絡ください。商品の状態を確認の上、返品・返金の手続きをさせていただきます。

また、お客様自身の都合による返品・返金については、致しかねる場合がございますので、予めご了承ください。

--------------------------
CharisWorks

お客様相談室:contact@charis.works
お問い合わせフォーム:[link]
`
		SendEmail(transactionDetails.Email, "【決済完了通知】ご購入ありがとうございます。", body)

	}

	log.Print("Email sent successfully!")
}

func SendPrivilegedEmail(userId string, firebaseApp validation.IFirebaseApp) {
	data, err := os.ReadFile("./auth_address.json")
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
	email, err := firebaseApp.GetEmail(userId)
	if err != nil {
		log.Fatalf("メールアドレスの取得に失敗しました：%v", err)
		return
	}
	for _, to := range *authAddress {
		body := `出品者権限を付与しました。`
		body += "\nemail: " + email
		SendEmail(to, "出品者権限付与通知", body)
	}
	{
		body := fmt.Sprintf(`
%v 様


出品者権限を付与しました。
以下リンクより口座登録をお願いします。

[link]
`, email)
		SendEmail(email, "出品者権限を付与しました。", body)

	}

	log.Print("Email sent successfully!")
}

func SendEmail(to string, subject string, body string) error {
	data, err := os.ReadFile("./email_credentials.json")
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
	authmail := emailCredentials["auth_mail"]
	from := emailCredentials["mail_from"]
	password := emailCredentials["mail_auth_pass"]
	smtp_server_addr := emailCredentials["smtp_server_addr"]
	smtp_server := emailCredentials["smtp_server"]
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
