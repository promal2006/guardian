package gqt_test

import (
	"fmt"
	"path"
	"path/filepath"
	"strconv"
	"time"

	"code.cloudfoundry.org/garden"
	"code.cloudfoundry.org/guardian/gqt/runner"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Info", func() {
	var (
		client          *runner.RunningGarden
		container       garden.Container
		containerLimits garden.Limits
	)

	BeforeEach(func() {
		client = runner.Start(config)
		containerLimits = garden.Limits{}
	})

	JustBeforeEach(func() {
		var err error
		container, err = client.Create(garden.ContainerSpec{
			Network: "10.252.0.2",
			Properties: garden.Properties{
				"foo": "bar",
			},
			Limits: containerLimits,
		})
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		Expect(client.DestroyAndStop()).To(Succeed())
	})

	It("can return the state", func() {
		info, err := container.Info()
		Expect(err).NotTo(HaveOccurred())
		Expect(info.State).To(Equal("active"))
	})

	It("can return the network information", func() {
		info, err := container.Info()
		Expect(err).NotTo(HaveOccurred())
		Expect(info.ContainerIP).To(Equal("10.252.0.2"))
		Expect(info.HostIP).To(Equal("10.252.0.1"))
	})

	It("can return the container path", func() {
		info, err := container.Info()
		Expect(err).NotTo(HaveOccurred())
		Expect(info.ContainerPath).To(Equal(path.Join(client.DepotDir, container.Handle())))
	})

	It("can return the list of properties", func() {
		Expect(container.SetProperty("abc", "xyz")).To(Succeed())

		info, err := container.Info()
		Expect(err).NotTo(HaveOccurred())

		Expect(info.Properties).To(HaveKeyWithValue("foo", "bar"))
		Expect(info.Properties).To(HaveKeyWithValue("abc", "xyz"))
	})

	It("can return port mappings", func() {
		hostPort, containerPort, err := container.NetIn(0, 0)
		Expect(err).NotTo(HaveOccurred())

		info, err := container.Info()
		Expect(err).NotTo(HaveOccurred())

		portMapping := info.MappedPorts[0]
		Expect(portMapping.HostPort).To(Equal(hostPort))
		Expect(portMapping.ContainerPort).To(Equal(containerPort))
	})

	When("the container has a memory limit applied", func() {
		BeforeEach(func() {
			containerLimits = garden.Limits{Memory: garden.MemoryLimits{LimitInBytes: 30 * mb}}
		})

		It("adds an out of memory event", func() {
			process, err := container.Run(garden.ProcessSpec{
				Path: "dd",
				Args: []string{"if=/dev/urandom", "of=/dev/shm/foo", "bs=1M", "count=32"},
			}, garden.ProcessIO{})
			Expect(err).NotTo(HaveOccurred())

			_, err = process.Wait()
			Expect(err).NotTo(HaveOccurred())
			expectedMemoryCgroupPath := client.CgroupSubsystemPath("memory", container.Handle())
			Eventually(getEventsForContainer(container), time.Minute).Should(
				ContainElement("Out of memory"),
				fmt.Sprintf("%#v", map[string]string{
					"Container PID":                        getContainerPid(container.Handle()),
					"Expected memory cgroup path":          expectedMemoryCgroupPath,
					"Pids in the container memory cgroup":  listPidsInCgroup(expectedMemoryCgroupPath),
					"Memory limit as listed in the cgroup": readFileString(filepath.Join(expectedMemoryCgroupPath, "memory.limit_in_bytes")),
					"Expected limit":                       strconv.FormatUint(containerLimits.Memory.LimitInBytes, 10),
				}),
			)
		})
	})
})

func getEventsForContainer(container garden.Container) func() []string {
	return func() []string {
		info, err := container.Info()
		Expect(err).NotTo(HaveOccurred())
		return info.Events
	}
}

var _ = Describe("BulkInfo", func() {
	var (
		client *runner.RunningGarden
	)

	BeforeEach(func() {
		client = runner.Start(config)
		_, err := client.Create(garden.ContainerSpec{
			Handle: "first",
		})
		Expect(err).NotTo(HaveOccurred())

		_, err = client.Create(garden.ContainerSpec{
			Handle: "second",
		})
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		Expect(client.DestroyAndStop()).To(Succeed())
	})

	It("can return info for each handle", func() {
		infos, err := client.BulkInfo([]string{"first", "second"})
		Expect(err).NotTo(HaveOccurred())

		Expect(infos).To(HaveLen(2))

		Expect(infos["first"].Info).NotTo(BeNil())
		Expect(infos["first"].Err).NotTo(HaveOccurred())

		Expect(infos["second"].Info).NotTo(BeNil())
		Expect(infos["second"].Err).NotTo(HaveOccurred())
	})
})
