package test

import (
	"fmt"
	"strings"

	"github.com/ryym/slv/slv/tp"
)

func TestAll(prg tp.Program, testDir string) (bool, error) {
	printer := newResultPrinter()
	loader := newTestLoader(testDir)
	return testAll(prg, loader, printer)
}

func testAll(prg tp.Program, loader testLoader, printer testResultPrinter) (bool, error) {
	testFiles, err := loader.ListFileNames()
	if err != nil {
		return false, err
	}

	totalResult := totalTestResult{}
	for _, filename := range testFiles {
		cases, err := loader.Load(filename)
		if err != nil {
			return false, err
		}

		totalResult.CaseCnt += len(cases)
		for i, tcase := range cases {
			out, err := prg.Run(tcase.In)
			if err != nil {
				return false, err
			}

			if !strings.HasSuffix(tcase.Out, "\n") {
				tcase.Out += "\n"
			}
			if tcase.Name == "" {
				tcase.Name = fmt.Sprintf("%s[%d]", filename, i)
			}

			result := testResult{
				Ok:       tcase.Out == out,
				TestCase: tcase,
				Actual:   out,
				Filename: filename,
			}

			printer.ShowResult(&result)

			if result.Ok {
				totalResult.PassedCnt += 1
			} else {
				totalResult.Fails = append(totalResult.Fails, result)
			}
		}
	}

	printer.ShowFailures(totalResult.Fails)
	printer.ShowSummary(&totalResult)

	return len(totalResult.Fails) == 0, nil
}
