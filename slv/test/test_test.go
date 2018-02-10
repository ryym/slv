package test

import (
	"testing"

	"github.com/go-test/deep"
	"github.com/ryym/slv/slv/tp"
)

func Test_testAll(t *testing.T) {
	cases := []testCase{
		{In: "a", Out: "a!"},
		{Name: "hello", In: "bb", Out: "bb!"},
		{Name: "fail", In: "C", Out: "C"},
	}

	wants := []testResult{
		{
			Ok:       true,
			TestCase: testCase{Name: "test.toml[0]", In: "a", Out: "a!"},
			Actual:   "a!",
			Filename: "test.toml",
		},
		{
			Ok:       true,
			TestCase: testCase{Name: "hello", In: "bb", Out: "bb!"},
			Actual:   "bb!",
			Filename: "test.toml",
		},
		{
			Ok:       false,
			TestCase: testCase{Name: "fail", In: "C", Out: "C"},
			Actual:   "C!",
			Filename: "test.toml",
		},
	}

	loader := &testLoaderMock{
		ListFileNamesFunc: func() ([]string, error) {
			return []string{"test.toml"}, nil
		},
		LoadFunc: func(_ string) ([]testCase, error) {
			return cases, nil
		},
	}

	var results []*testResult
	var total *totalTestResult

	handler := &testResultHandlerMock{
		OnCaseEndFunc: func(ret *testResult) {
			results = append(results, ret)
		},
		OnEndFunc: func(t *totalTestResult) {
			total = t
		},
	}

	prg := &tp.ProgramMock{
		CompileFunc: func() (tp.CompileResult, error) {
			return tp.CompileResult{true, []byte{}}, nil
		},
		RunFunc: func(in string) (string, error) {
			return in + "!", nil
		},
	}

	testAll(prg, loader, handler)

	// Check each test result.
	for i, ret := range results {
		if diff := deep.Equal(ret, &wants[i]); diff != nil {
			t.Errorf("[%d]: %s", i, diff)
		}
	}

	// Check total result.
	wantTotal := totalTestResult{
		CaseCnt:   3,
		PassedCnt: 2,
		Fails:     []*testResult{&wants[2]},
	}
	if diff := deep.Equal(total, &wantTotal); diff != nil {
		t.Errorf("wrong total: %s", diff)
	}
}
