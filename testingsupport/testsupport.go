package testsupport

import (
	"fmt"
	"os"
	"time"

	"github.com/kraneware/core-go/awsutil/services"

	"golang.org/x/sync/errgroup"

	"github.com/kraneware/core-go/awsutil/localstack"
)

const (
	Bucket1               = "bucket1"
	Bucket1Key            = "test"
	Bucket1ObjectContents = "test bucket 1 object contents"

	Bucket2               = "bucket2"
	Bucket2Key            = "test"
	Bucket2ObjectContents = "test bucket 2 object contents"
)

var (
	// DefaultEventuallyTimeout Parameters for eventual checks in tests
	DefaultEventuallyTimeout         = 60 * time.Second // nolint:gochecknoglobals
	DefaultEventuallyPollingInterval = 5 * time.Second  // nolint:gochecknoglobals
)

func createBuckets() error {
	testCtx, td := services.NewTestDaemon()
	defer td.Close()

	err := localstack.NewS3Bucket(testCtx, Bucket1)
	if err != nil {
		return err
	}

	err = localstack.NewS3Bucket(testCtx, Bucket2)
	if err != nil {
		return err
	}

	err = localstack.NewS3BucketObject(testCtx, Bucket1, Bucket1Key, []byte(Bucket1ObjectContents))
	if err != nil {
		return err
	}

	err = localstack.NewS3BucketObject(testCtx, Bucket2, Bucket2Key, []byte(Bucket2ObjectContents))
	if err != nil {
		return err
	}

	return nil
}

// BuildPlatformTestingInfrastructure builds the required infrastructure for testing with localstack
func BuildPlatformTestingInfrastructure() (err error) {
	fmt.Println("Initializing Platform testing infrastructure ... ")

	var errGroup errgroup.Group
	errGroup.Go(createBuckets)
	errGroup.Go(
		func() error {
			return os.Setenv("DATA_S3_BUCKET", Bucket1)
		},
	)

	return errGroup.Wait()
}
