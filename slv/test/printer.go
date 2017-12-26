package test

import "fmt"

type defaultPrinter struct{}

func NewResultPrinter() TestResultPrinter {
	return &defaultPrinter{}
}

func (p *defaultPrinter) ShowResult(ret *testResult) {
	if ret.Ok {
		fmt.Print(".")
	} else {
		fmt.Print("F")
	}
}

func (p *defaultPrinter) ShowFailures(results []testResult) {
	fmt.Println("")
	for _, r := range results {
		tc := r.TestCase

		// XXX: Should output diff text.
		tmpl := `
Case: %s
- Expected:
%s
- Actual:
%s`
		fmt.Printf(tmpl, tc.Name, tc.Out, r.Actual)
	}
}

func (p *defaultPrinter) ShowSummary(tr *totalTestResult) {
	var title string
	if len(tr.Fails) == 0 {
		title = "OK"
	} else {
		title = "FAILED"
	}
	fmt.Printf(
		"\n[%s] All: %d, Passed: %d, Failed: %d\n",
		title,
		tr.CaseCnt,
		tr.PassedCnt,
		len(tr.Fails),
	)
}
