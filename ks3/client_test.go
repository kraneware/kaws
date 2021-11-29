package ks3

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/kraneware/core-go/awsutil/localstack"
	"github.com/kraneware/core-go/utils"
	"github.com/kraneware/kaws"
	testsupport "github.com/kraneware/kaws/testingsupport"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	environ [] string
	client Client
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

	Expect(localstack.StartContainer()).Should(BeNil())
	Expect(testsupport.BuildPlatformTestingInfrastructure()).Should(BeNil())
})

var _ = AfterSuite(func() {
	defer Expect(utils.ResetEnv(environ)).Should(BeNil())
	Expect(localstack.StopContainer()).Should(BeNil())
})

var _ = Describe("Kaws/s3 tests", func() {

	Context("ListBuckets() Test", func() {
		It("should list buckets", func() {
			Initialize()
			buckets, err := client.GetBucketList()

			Expect(err).To(BeNil())
			Expect(buckets).To(Not(BeNil()))
			Expect(buckets).To(BeAssignableToTypeOf([] s3.Bucket{}))

			var bucketNames [] string
			for _, v := range buckets {
				bucketNames = append(bucketNames, *v.Name)
			}

			Expect(bucketNames).Should(ContainElement(testsupport.Bucket1))
			Expect(bucketNames).Should(ContainElement(testsupport.Bucket2))
		})
	})
})

func Initialize() {
	client, _ = CreateClient()
}