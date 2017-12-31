package slv

// TOML data
const DEFAULT_CONF = `
[lang.go]
exts = [".go"]
compile = "go build -o {{.Dest}}/out {{.Src}}"
run = "{{.Dest}}/out"

[lang.ruby]
exts = [".rb"]
run = "ruby {{.Src}}"

[lang.scala]
exts = [".scala"]
compile = "scalac -d {{.Dest}} {{.Src}}"
run = "scala -cp {{.Dest}} Main"
`
