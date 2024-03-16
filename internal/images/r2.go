package images

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/charisworks/charisworks-backend/internal/utils"
	r2client "github.com/whatacotton/cloudflare-go/client"
	"github.com/whatacotton/cloudflare-go/r2"
)

var (
	Bucket = os.Getenv("BUCKET_NAME")
)

type R2Conns struct {
	Crud r2.R2crud
	ctx  context.Context
}

var (
	EndPoint        string = os.Getenv("ENDPOINT")
	AccountID       string = os.Getenv("ACCOUNT_ID")
	AccessKeyID     string = os.Getenv("ACCESS_KEY_ID")
	accessKeySecret string = os.Getenv("ACCESS_KEY_SECRET")
)

func (r *R2Conns) Init() {
	flag.StringVar(&EndPoint, "endpoint", EndPoint, "endpoint")
	flag.StringVar(&AccountID, "account-id", AccountID, "account-id")
	flag.StringVar(&AccessKeyID, "access-key-id", AccessKeyID, "access-key-id")
	flag.StringVar(&accessKeySecret, "account-key-secret", accessKeySecret, "account-key-secret")
	flag.Parse()

	if EndPoint == "" || AccountID == "" || AccessKeyID == "" || accessKeySecret == "" {
		panic("missing required parameters")
	}
	client, err := r2client.New(
		AccountID,
		EndPoint,
		AccessKeyID,
		accessKeySecret,
	).Connect(context.TODO())
	if err != nil {
		log.Fatalf("r2 client conneciton error :%v\n", err)
	}
	r.ctx = context.Background()
	r.Crud = r2.NewR2CRUD(Bucket, client, 60)
	log.Println("r2 client connection success")
}

func (r *R2Conns) UploadImage(file *os.File, path string) error {
	filedata, err := fileToByte(file)
	if err != nil {
		log.Print(err)
		err = &utils.InternalError{Message: utils.InternalErrorR2}
		return err
	}

	return r.Crud.UploadObject(r.ctx, filedata, path)
}

func (r *R2Conns) GetImages(path string) ([]string, error) {
	objects, err := r.Crud.ListObjects(r.ctx, path)
	if err != nil {
		log.Print(err)
		err = &utils.InternalError{Message: utils.InternalErrorR2}
		return nil, err
	}
	log.Print(objects)
	var images []string
	for _, obj := range objects.Contents {
		log.Print(*obj.Key)
		images = append(images, *obj.Key)
	}
	return images, nil
}

func (r *R2Conns) DeleteImage(path string) error {
	err := r.Crud.DeleteObject(r.ctx, path)
	if err != nil {
		log.Print(err)
		err = &utils.InternalError{Message: utils.InternalErrorR2}
		return err
	}
	return nil
}

// 画像ファイルをバイト列に変換
func fileToByte(file *os.File) ([]byte, error) {
	fileInfo, err := file.Stat()
	if err != nil {
		log.Print(err)
		err = &utils.InternalError{Message: utils.InternalErrorR2}
		return nil, err
	}
	fileBytes := make([]byte, fileInfo.Size())
	_, err = file.Read(fileBytes)
	if err != nil {
		log.Print(err)
		err = &utils.InternalError{Message: utils.InternalErrorR2}
		return nil, err
	}

	return fileBytes, nil
}

/*
func main() {
	client, err := r2client.New(
		AccountID,
		EndPoint,
		AccessKeyID,
		accessKeySecret,
	).Connect(context.TODO())
	if err != nil {
		log.Fatalf("r2 client conneciton error :%v\n", err)
	}

	r2 := r2.NewR2CRUD(Bucket, client, 60)

	// 画像ファイルを開いておく
	file, err := os.OpenFile("sample.png", os.O_RDONLY, 0666)
	if err != nil {
		log.Fatalf("file open error :%v\n", err)
	}

	filedata, err := fileToByte(file)
	if err != nil {
		log.Fatalf("file transcate error :%v\n", err)
	}

	// 画像ファイルをアップロード
	if err := r2.UploadObject(context.Background(), filedata, "folder2/sample2.png"); err != nil {
		log.Fatalf("r2 upload error :%v\n", err)
	}

	log.Println("upload success")

	// 全てのオブジェクトを取得
	objects, err := r2.ListObjects(context.Background(), "folder2/")
	if err != nil {
		log.Fatalf("r2 list objects error :%v\n", err)
	}

	for i, obj := range objects.Contents {
		log.Printf("%d: %s", i, *obj.Key) // オブジェクトのキー表示
	}

	// オブジェクトのURLを取得
	url, err := r2.PublishPresignedObjectURL(context.Background(), "folder/sample.png")
	if err != nil {
		log.Fatalf("r2 publish presigned object url error :%v\n", err)
	}

	log.Println("sample.png Access URL :", url)

	// オブジェクトを削除
	if err := r2.DeleteObject(context.Background(), "sample.png"); err != nil {
		log.Fatalf("r2 delete object error :%v\n", err)
	}

	log.Println("delete success")
}
*/
