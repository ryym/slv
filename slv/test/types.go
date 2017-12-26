package test

type testCase struct {
	Name string
	In   string
	Out  string
}

type testCases struct {
	Test []testCase
}

type testResult struct {
	Ok       bool
	TestCase testCase
	Actual   string
	Filename string
}

type totalTestResult struct {
	CaseCnt   int
	PassedCnt int
	Fails     []testResult
}

type TestResultPrinter interface {
	ShowResult(result *testResult)
	ShowFailures(cases []testResult)
	ShowSummary(total *totalTestResult)
}
