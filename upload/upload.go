package upload

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"os"
	"regexp"

	"github.com/google/uuid"
	"github.com/joho/godotenv"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func UploadImage(base64image string) (*string, error) {
	mode := os.Getenv("MODE")
	if mode == "production" {
		err := godotenv.Load(fmt.Sprintf("./%s.env", os.Getenv("GO_ENV")))
		if err != nil {
			return nil, err
		}
	}

	AccessKeyID := os.Getenv("AccessKeyID")
	SecretAccessKey := os.Getenv("SecretAccessKey")

	regex := `^data:\w+\/\w+;base64,`
	re := regexp.MustCompile(regex)

	base64image = re.ReplaceAllString(base64image, "")
	data, err := base64.StdEncoding.DecodeString(base64image)
	file := new(bytes.Buffer)
	file.Write(data)

	// img, err := jpeg.Decode(file)
	// if err != nil {
	// 	return nil, err
	// }

	// m := resize.Resize(600, 0, img, resize.Bilinear)
	// pr, pw := io()
	// go func() {
	// 	if err := jpeg.Encode(pw, m, nil); err != nil {
	// 		fmt.Println("error:jpeg\n", err)
	// 		return
	// 	}
	// 	pw.Close()
	// }()

	ext := ".jpeg"
	filename := uuid.New().String() + ext
	filename = "/tweet/" + filename

	sess := session.Must(session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(AccessKeyID, SecretAccessKey, ""),
		Region:      aws.String("us-east-2"),
	}))

	bucketName := "go-app-bucket"

	uploader := s3manager.NewUploader(sess)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(filename),
		Body:        file,
		ContentType: aws.String("image/jpeg"),
	})
	if err != nil {
		return nil, err
	}

	return &filename, nil
}

func UploadIcon(base64image string) (*string, error) {
	AccessKeyID := os.Getenv("AccessKeyID")
	SecretAccessKey := os.Getenv("SecretAccessKey")

	regex := `^data:\w+\/\w+;base64,`
	re := regexp.MustCompile(regex)

	base64image = re.ReplaceAllString(base64image, "")
	data, err := base64.StdEncoding.DecodeString(base64image)
	file := new(bytes.Buffer)
	file.Write(data)

	ext := ".jpeg"
	filename := uuid.New().String() + ext
	filename = "/user/" + filename

	sess := session.Must(session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(AccessKeyID, SecretAccessKey, ""),
		Region:      aws.String("us-east-2"),
	}))

	bucketName := "go-app-bucket"

	uploader := s3manager.NewUploader(sess)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(filename),
		Body:        file,
		ContentType: aws.String("image/jpeg"),
	})
	if err != nil {
		return nil, err
	}

	return &filename, nil
}
