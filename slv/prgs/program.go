package prgs

import (
	"io"
	"os"
	"os/exec"

	"github.com/pkg/errors"
	"github.com/ryym/slv/slv/tp"
)

type programImpl struct {
	srcPath string
	destDir string
	def     programDef
}

func (p *programImpl) Compile() (ret tp.CompileResult, err error) {
	_, err = os.Stat(p.destDir)
	if os.IsNotExist(err) {
		err = os.Mkdir(p.destDir, 0755)
	}
	if err != nil {
		return ret, errors.Wrap(err, "Failed to create work dir")
	}

	cmds := p.def.GetCompileCmds(p.srcPath, p.destDir)

	var out []byte
	if cmds != nil {
		out, err = exec.Command(cmds[0], cmds[1:]...).CombinedOutput()
		if err != nil {
			return ret, errors.Wrap(err, string(out))
		}
	}

	return tp.CompileResult{
		Compiled: cmds != nil,
		Output:   out,
	}, nil
}

func (p *programImpl) Run(input string) (string, error) {
	_, err := p.Compile()
	if err != nil {
		return "", err
	}

	execCmds := p.def.GetExecCmds(p.srcPath, p.destDir)
	cmd := exec.Command(execCmds[0], execCmds[1:]...)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return "", errors.Wrap(err, "Failed to pipe stdin")
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, input)
	}()

	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", errors.Wrap(err, string(out))
	}

	return string(out), nil
}

func (p *programImpl) RunWithPipes(stdin io.ReadCloser, stdout io.WriteCloser) error {
	_, err := p.Compile()
	if err != nil {
		return err
	}

	execCmds := p.def.GetExecCmds(p.srcPath, p.destDir)
	cmd := exec.Command(execCmds[0], execCmds[1:]...)
	cmd.Stdin = stdin
	cmd.Stdout = stdout
	return cmd.Run()
}
