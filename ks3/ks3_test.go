package ks3_test

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/kraneware/kaws"
	"github.com/kraneware/kaws/ks3"
	_ "github.com/kraneware/kore-go/helper"
	utils "github.com/kraneware/kore-go/helper"
	"github.com/kraneware/lokalstack"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"golang.org/x/sync/errgroup"
	"os"
	"testing"
	"time"
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
	environ []string
	client  ks3.Client

	DefaultEventuallyTimeout         = 60 * time.Second // nolint:gochecknoglobals
	DefaultEventuallyPollingInterval = 5 * time.Second  // nolint:gochecknoglobals
)

func TestClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Kaws/s3 Test Suite")
}

var _ = BeforeSuite(func() {
	environ := os.Environ()

	if kaws.DEBUG {
		for _, v := range environ {
			fmt.Printf("%s\n", v)
		}
	}

	Expect(lokalstack.StartContainer()).Should(BeNil())
	Expect(buildTestingInfrastructure()).Should(BeNil())
})

var _ = AfterSuite(func() {
	defer Expect(utils.ResetEnv(environ)).Should(BeNil())
	Expect(lokalstack.StopContainer()).Should(BeNil())
})

var _ = Describe("Kaws/ks3 tests", func() {

	Context("ListBuckets() Test", func() {
		It("should list buckets", func() {
			Initialize()
			buckets, err := client.GetBucketList()

			Expect(err).To(BeNil())
			Expect(buckets).To(Not(BeNil()))
			Expect(buckets).To(BeAssignableToTypeOf([]s3.Bucket{}))

			var bucketNames []string
			for _, v := range buckets {
				bucketNames = append(bucketNames, *v.Name)
			}

			Expect(bucketNames).Should(ContainElement(Bucket1))
			Expect(bucketNames).Should(ContainElement(Bucket2))
		})
	})

	Context("GetObjects() Test", func() {
		It("should list bucket objects", func() {
			Initialize()
			objects, err := client.GetBucketObjects(Bucket1, "test")

			Expect(err).To(BeNil())
			Expect(objects).To(Not(BeNil()))
			Expect(objects).To(BeAssignableToTypeOf([]s3.Object{}))

			for _, v := range objects {
				fmt.Sprintf("%v", v)
			}
		})
	})
})

func Initialize() {
	client, _ = ks3.CreateClient()
}

// buildPlatformTestingInfrastructure builds the required infrastructure for testing with lokalstack
func buildTestingInfrastructure() (err error) {
	fmt.Println("initializing testing infrastructure ... ")

	var errGroup errgroup.Group
	errGroup.Go(createBuckets)
	errGroup.Go(
		func() error {
			return os.Setenv("DATA_S3_BUCKET", Bucket1)
		},
	)

	return errGroup.Wait()
}

func createBuckets() error {
	testCtx, td := lokalstack.NewTestDaemon()
	defer td.Close()

	err := lokalstack.NewS3Bucket(testCtx, Bucket1)
	if err != nil {
		return err
	}

	err = lokalstack.NewS3Bucket(testCtx, Bucket2)
	if err != nil {
		return err
	}

	err = lokalstack.NewS3BucketObject(testCtx, Bucket1, Bucket1Key, []byte(Bucket1ObjectContents))
	if err != nil {
		return err
	}

	err = lokalstack.NewS3BucketObject(testCtx, Bucket2, Bucket2Key, []byte(Bucket2ObjectContents))
	if err != nil {
		return err
	}

	return nil
}
