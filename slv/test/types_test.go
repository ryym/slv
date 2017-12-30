package test

type FkTestLoader struct {
	FkListFileNames func() ([]string, error)
	FkLoad          func(filename string) ([]testCase, error)
}

func (tl *FkTestLoader) ListFileNames() ([]string, error) {
	return tl.FkListFileNames()
}
func (tl *FkTestLoader) Load(filename string) ([]testCase, error) {
	return tl.FkLoad(filename)
}

type FkTestResultHandler struct {
	FkOnCaseEnd func(result *testResult)
	FkOnEnd     func(total *totalTestResult)
}

func (h *FkTestResultHandler) OnCaseEnd(result *testResult) {
	h.FkOnCaseEnd(result)
}
func (h *FkTestResultHandler) OnEnd(total *totalTestResult) {
	h.FkOnEnd(total)
}
