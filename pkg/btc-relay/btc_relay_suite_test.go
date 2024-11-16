package btcrelay_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestBtcRelay(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BtcRelay Suite")
}
