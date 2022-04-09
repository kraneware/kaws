package ks3

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/kraneware/core-go/utils"
	"github.com/kraneware/kaws"
	testsupport "github.com/kraneware/kaws/testingsupport"
	_ "github.com/kraneware/kore-go/helper"
	"github.com/kraneware/lokalstack"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	environ []string
	client  Client
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
	Expect(testsupport.BuildPlatformTestingInfrastructure()).Should(BeNil())
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

			Expect(bucketNames).Should(ContainElement(testsupport.Bucket1))
			Expect(bucketNames).Should(ContainElement(testsupport.Bucket2))
		})
	})

	Context("GetObjects() Test", func() {
		It("should list bucket objects", func() {
			Initialize()
			objects, err := client.GetBucketObjects(testsupport.Bucket1, "test")

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
	client, _ = CreateClient()
}
