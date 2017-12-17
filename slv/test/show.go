package test

import (
	"errors"
	"fmt"
)

// XXX: It will be better if each case result is printed as soon as
// the test case is run. Current implementation shows all results
// after all test cases are done.

func ShowResult(ret *testResult) error {
	if ret.Ok {
		fmt.Println(makeSummary(ret))
		return nil
	}

	s := make([]byte, 0)
	for _, fc := range ret.FailedCases {
		s = append(s, makeFailedResult(&fc)...)
	}
	s = append(s, makeSummary(ret)...)
	return errors.New(string(s))
}

func makeSummary(ret *testResult) string {
	var title string
	if ret.Ok {
		title = "OK"
	} else {
		title = "FAILED"
	}
	return fmt.Sprintf(
		"[%s] All: %d, Passed: %d, Failed: %d",
		title,
		ret.CaseCnt,
		ret.PassedCnt,
		len(ret.FailedCases),
	)
}

func makeFailedResult(fc *failedTestCase) string {
	// XXX: Should output diff text.
	return fmt.Sprintf(`case: %s
- Expected:
	%s
- Actual:
	%s`, fc.Name, fc.Out, fc.Actual)
}
