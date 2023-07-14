package failure_examples

import (
	"github.com/albert-k-nylas/allure-go"
	"github.com/pkg/errors"
	"testing"
)

func TestFailNowAllure(t *testing.T) {
	allure.Test(t, allure.Description("Test with Allure error in it"), allure.Action(func() {
		allure.FailNow(errors.New("A more serious error"))
		allure.Step(allure.Description("Step you're not supposed to see"), allure.Action(func() {}))
	}))
}
