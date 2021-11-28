package s3

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/kraneware/core-go/awsutil/services"
	"github.com/mholt/archiver"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func DownloadFile(bucket string, key string){
	var dl *s3manager.Downloader = services.S3Downloader()
	dl.Download(io.wr)
}

func downloadAndUnzipFile(s3Object events.S3Entity) (localFile string, err error) {
	absPath := os.TempDir() + "/" + s3Object.Object.Key

	filename := filepath.Base(absPath)

	downloadPath := os.TempDir() + "/" + filename

	var file *os.File
	file, err = os.Create(downloadPath)
	defer closeFile(file)

	var unzippedPath string
	if err == nil {
		//log.Infof("Downloading s3://%s/%s to %s", s3Object.Bucket.Name, s3Object.Object.Key, downloadPath)

		downloader := services.S3Downloader()

		var numBytes int64
		numBytes, err = downloader.Download(file,
			&s3.GetObjectInput{
				Bucket: aws.String(s3Object.Bucket.Name),
				Key:    aws.String(s3Object.Object.Key),
			})

		_ = fmt.Sprintf("Downloadeded %s %d bytes", file.Name(), numBytes)

		filename = strings.TrimSuffix(filename, ".zip")
		unzippedPath = os.TempDir() + "/" + filename + strconv.FormatInt(time.Now().UnixNano(), 10)
		err = archiver.Unarchive(downloadPath, unzippedPath)

	}

	return unzippedPath, err
}

func closeFile(f *os.File) {
	err := f.Close()

	if err != nil {
		//log.With(zap.Error(err)).Error(fmt.Sprintf("Error with closing file: %s", f.Name()))
	}
}