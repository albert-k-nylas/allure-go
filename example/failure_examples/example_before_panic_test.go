package failure_examples

import (
	"github.com/albert-k-nylas/allure-go"
	"testing"
)

func TestBeforePanic(t *testing.T) {
	allure.BeforeTest(t, allure.Action(func() {
		panic("panic at the before statement! (disco)")
	}))

	allure.Test(t, allure.Action(func() {}))
}
