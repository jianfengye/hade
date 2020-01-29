package suite2

import "testing"
import . "github.com/smartystreets/goconvey/convey"

func TestRunSuite(t *testing.T) {
	SetUp()
	defer TearDown()
	Convey("初始化", t, nil)

	runCase(t, NormalCase1)

}

func runCase(t *testing.T, testCase func(*testing.T)) {
	Before()
	defer After()

	testCase(t)
}