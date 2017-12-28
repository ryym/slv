package test

import (
	"fmt"

	"github.com/fatih/color"
	diffmp "github.com/sergi/go-diff/diffmatchpatch"
)

type colorizer func(a ...interface{}) string

type defaultPrinter struct {
	colorBad  colorizer
	colorGood colorizer
}

func NewResultPrinter() TestResultPrinter {
	return &defaultPrinter{
		colorBad:  color.New(color.FgRed).SprintFunc(),
		colorGood: color.New(color.FgGreen).SprintFunc(),
	}
}

func (p *defaultPrinter) ShowResult(ret *testResult) {
	if ret.Ok {
		fmt.Print(".")
	} else {
		fmt.Print(p.colorBad("F"))
	}
}

func (p *defaultPrinter) ShowFailures(results []testResult) {
	fmt.Println("\n")
	for _, r := range results {
		tc := r.TestCase

		fmt.Printf("%s:\n\n", tc.Name)

		diffs := makeLineDiff(tc.Out, r.Actual)
		for _, d := range diffs {
			switch d.Type {
			case diffmp.DiffDelete:
				fmt.Print(p.colorBad("-" + d.Text))
			case diffmp.DiffInsert:
				fmt.Print(p.colorGood("+" + d.Text))
			case diffmp.DiffEqual:
				fmt.Print(d.Text)
			}
		}
		fmt.Println("")
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

// https://qiita.com/shibukawa/items/dd75ad01e623c4c1166b
func makeLineDiff(s1, s2 string) []diffmp.Diff {
	dmp := diffmp.New()
	a, b, c := dmp.DiffLinesToChars(s1, s2)
	diffs := dmp.DiffMain(a, b, false)
	return dmp.DiffCharsToLines(diffs, c)
}
