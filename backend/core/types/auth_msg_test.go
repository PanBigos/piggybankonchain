package types

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

func TestRegistrationMatching(t *testing.T) {

	type testCase struct {
		name     string
		message  string
		match    bool
		errorMsg string
	}

	var (
		testCases []testCase = []testCase{
			testCase{
				name:     "InvalidEntity",
				message:  fmt.Sprintf(registerMsg, "fooo", "bar"),
				match:    false,
				errorMsg: "should not match with invalid entity",
			},
			testCase{
				name:     "InvalidAction",
				message:  fmt.Sprintf(registerMsg, "foo", "barr"),
				match:    false,
				errorMsg: "should not match with invalid action",
			},
		}
	)

	for i := range testCases {
		testcase := testCases[i]
		t.Run(testcase.name, func(t *testing.T) {
			matcher := NewRegistrationMatcher("foo", "bar")
			if result := matcher.Matches(testcase.message); result != testcase.match {
				t.Fatalf(
					"msg: %v, does not pass test. error: %v. output: %v",
					testcase.message,
					testcase.errorMsg,
					result,
				)
			}
		})
	}
}

func TestHeaderRecovery(t *testing.T) {

	type testCase struct {
		name     string
		msg      AuthMessage
		validate func(msg string, deadline time.Time, address common.Address, ok bool) error
	}

	ts, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	factory := RegistrationMessageFactory{}
	template := factory.GenerateNewTemplate("foo", "bar")
	var (
		testCases []testCase = []testCase{
			testCase{
				name: "ValidMessage",
				msg:  template.NewMessage(common.Address{0x1}, ts.Add(5*time.Second)),
				validate: func(msg string, deadline time.Time, address common.Address, ok bool) error {
					if !ok {
						return fmt.Errorf("different recover result. got: %v, want: %v", ok, true)
					}
					if address != (common.Address{0x1}) {
						t.Fatalf(
							"different recover address. got: %v, want: %v",
							address,
							common.Address{0x1},
						)
					}
					if !deadline.Equal(ts.Add(5 * time.Second)) {
						t.Fatalf(
							"different recover deadline. got: %v, want: %v",
							ts,
							ts.Add(5*time.Second),
						)
					}
					if !template.IsKnown(msg) {
						return errors.New("template message not verified")
					}
					return nil
				},
			},
			testCase{
				name: "InvalidMessage",
				msg:  AuthMessage(fmt.Sprintf(registerMsg, "xyz", "1")),
				validate: func(_ string, _ time.Time, _ common.Address, ok bool) error {
					if ok {
						return fmt.Errorf("different recover result. got: %v, want: %v", ok, true)
					}
					return nil
				},
			},
		}
	)

	for i := range testCases {
		testcase := testCases[i]
		t.Run(testcase.name, func(t *testing.T) {
			msg, ts, address, ok := template.recover(testcase.msg)
			if err := testcase.validate(msg, ts, address, ok); err != nil {
				t.Fatal(err)
			}
		})
	}

}
