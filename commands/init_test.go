package commands_test

import (
	"github.com/fatih/color"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"

	"testing"
)

func TestCommands(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "commands")
}

func writeTestConfigFile(contents string) string {
	file, err := ioutil.TempFile("", "config-*.yml")
	Expect(err).ToNot(HaveOccurred())

	err = ioutil.WriteFile(file.Name(), []byte(contents), 0777)
	Expect(err).ToNot(HaveOccurred())
	return file.Name()
}

var _ = BeforeSuite(func() {
	//enable color for this suite, so that colors are tested even in parallel
	//(the color library detects non-tty terminals,
	//which ginkgo uses when running in parallel,
	//so we have to override it)
	color.NoColor = false
})
