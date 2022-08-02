package kinterfaceimpl_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/kraneware/kaws"
<<<<<<< HEAD:kinterfaceimpl/s3persister_test.go
	testsupport "github.com/kraneware/kaws/testingsupport"
=======
	"github.com/kraneware/kaws/kinterfaceimpl"
	utils "github.com/kraneware/kore-go/helper"
	localstack "github.com/kraneware/lokalstack"
	"golang.org/x/sync/errgroup"
	"os"
	"testing"
>>>>>>> b24503ac92f974fca1167eba3d1c3a05b643ee39:kinterfaceimpl/kinterfaceimpl_test.go

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	Bucket1               = "bucket1"
	Bucket1Key            = "test"
	Bucket1ObjectContents = "test bucket 1 object contents"
)

var (
	environ       []string
	s3persister   kinterfaceimpl.S3Persister
	secretskeeper kinterfaceimpl.Skeeper
)

func TestClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "kinterfaceimpl Test Suite")
}

var _ = BeforeSuite(func() {
	environ := os.Environ()

	if kaws.DEBUG {
		for _, v := range environ {
			fmt.Printf("%s\n", v)
		}
	}

	Expect(localstack.StartContainer()).Should(BeNil())
	Expect(buildTestingInfrastructure()).Should(BeNil())
})

var _ = AfterSuite(func() {
	defer Expect(utils.ResetEnv(environ)).Should(BeNil())
	Expect(localstack.StopContainer()).Should(BeNil())
})

var _ = Describe("Kaws/s3 tests", func() {
	Context("CreatePersister() Test", func() {
		It("should create the  s3 persister", func() {
			p, err := kinterfaceimpl.CreatePersister(kaws.Name, kaws.Id)

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
	s3persister, _ = kinterfaceimpl.CreatePersister(kaws.Name, kaws.Id)
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
	testCtx, td := localstack.NewTestDaemon()
	defer td.Close()

	err := localstack.NewS3Bucket(testCtx, Bucket1)
	if err != nil {
		return err
	}

	err = localstack.NewS3BucketObject(testCtx, Bucket1, Bucket1Key, []byte(Bucket1ObjectContents))
	if err != nil {
		return err
	}

	return nil
}
