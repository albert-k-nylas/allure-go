package failure_examples

import (
	"github.com/albert-k-nylas/allure-go"
	"github.com/pkg/errors"
	"testing"
)

func TestFailAllure(t *testing.T) {
	allure.Test(t, allure.Description("Test with Allure error in it"), allure.Action(func() {
		allure.Step(allure.Description("Step 1"), allure.Action(func() {
		}))
		allure.Step(allure.Description("Step 2"), allure.Action(func() {
			allure.Fail(errors.New("Error message"))
		}))
	}))
}
