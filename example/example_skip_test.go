package example

import (
	"github.com/albert-k-nylas/allure-go"
	"testing"
)

func TestSkipStep(t *testing.T) {
	allure.Test(t,
		allure.Description("Skip test"),
		allure.Action(func() {
			allure.SkipStep(allure.Description("Skip"), allure.Action(func() {}), allure.Reason("reason"))
		}))
}

func TestSkipTest(t *testing.T) {
	allure.SkipTest(t,
		allure.Reason("Skipping the test"),
		allure.Description("Skip test"),
		allure.Action(func() {
			allure.SkipStep(allure.Description("Skip"), allure.Action(func() {}), allure.Reason("reason"))
		}))
}
