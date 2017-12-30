package test

type testLoader interface {
	ListFileNames() ([]string, error)
	Load(filename string) ([]testCase, error)
}

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
	Fails     []*testResult
}

type testResultHandler interface {
	OnCaseEnd(result *testResult)
	OnEnd(total *totalTestResult)
}
