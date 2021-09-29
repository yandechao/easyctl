package ulimit

import (
	_ "embed"
	"github.com/sirupsen/logrus"
	"github.com/weiliang-ms/easyctl/pkg/set"
	"os"
	"testing"
)

func TestUlimit(t *testing.T) {
	b, readErr := os.ReadFile("../../../asset/config.yaml")
	if readErr != nil {
		panic(readErr)
	}

	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	err := set.Ulimit(b, logger)
	if err != nil {
		panic(err)
	}

}