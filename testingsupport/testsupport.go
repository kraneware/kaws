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
	OneSecondTimeout = 1
	Bucket1 = "bucket1"
	Bucket2 = "bucket2"
)

var (
	// DefaultEventuallyTimeout Parameters for eventual checks in tests
	DefaultEventuallyTimeout         = 60 * time.Second // nolint:gochecknoglobals
	DefaultEventuallyPollingInterval = 5 * time.Second  // nolint:gochecknoglobals
)

func createBuckets() error {
	testCtx, td := services.NewTestDaemon()
	defer td.Close()

	err := localstack.NewS3Bucket(testCtx, Bucket2)
	if err != nil {
		return err
	}

	return localstack.NewS3Bucket(testCtx, Bucket1)
}

// BuildPlatformTestingInfrastructure builds the required infrastructure for testing with localstack
func BuildPlatformTestingInfrastructure() (err error) {
	fmt.Println("Initializing Platform testing infrastructure ... ")

	var errGroup errgroup.Group
	errGroup.Go(createBuckets)
	errGroup.Go(func() error { return os.Setenv("DATA_S3_BUCKET", Bucket1) })

	return errGroup.Wait()
}
