package kinterfaceimpl

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
	s3persister S3Persister
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
	Context("CreatePersister() Test", func() {
		It("should create the  s3 persister", func() {
			p, err := CreatePersister(kaws.Name, kaws.Id)

			Expect(err).To(BeNil())
			Expect(p).Should(Not(BeNil()))
		})
	})
	Context("InitPersister() Test", func() {
		It("should init the s3 persister", func() {
			Initialize()
			s3persister, err := s3persister.InitPersister()

			Expect(err).To(BeNil())
			Expect(s3persister.GetConnection()).To(BeAssignableToTypeOf(&s3.S3{}))
		})
	})
	Context("GetName() Test", func() {
		It("should return the name of the s3 persister", func() {
			Initialize()
			s3persister, _ = s3persister.InitPersister()
			pname := s3persister.GetName()

			Expect(pname).To(Equal(kaws.Name))
		})
	})
	Context("ListBuckets() Test", func() {
		It("should return the bucket list", func() {
			Initialize()
			s3persister, _ = s3persister.InitPersister()
			c, err := s3persister.ListBuckets()

			Expect(err).To(BeNil())
			Expect(c).To(BeAssignableToTypeOf([]s3.Bucket{}))
		})
	})
	//Context("ClosePersister() Test", func() {
	//	It("should close the s3 persister", func() {
	//		Initialize()
	//		_ = s3persister.InitPersister()
	//		err := s3persister.ClosePersister()
	//
	//		Expect(err).To(BeNil())
	//	})
	//})

	Context("ID() Test", func() {
		It("should return the id  of the s3 persister", func() {
			Initialize()
			s3persister, _ = s3persister.InitPersister()
			pid := s3persister.ID()

			Expect(pid).To(Equal(kaws.Id))
		})
	})
	Context("GetConnection() Test", func() {
		It("should return the connection of the s3 persister", func() {
			Initialize()
			s3persister, _ = s3persister.InitPersister()
			c, err := s3persister.GetConnection()

			Expect(err).To(BeNil())
			Expect(c).To(BeAssignableToTypeOf(&s3.S3{}))
		})
	})

})

func Initialize() {
	s3persister, _ = CreatePersister(kaws.Name, kaws.Id)
}