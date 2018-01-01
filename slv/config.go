package slv

// TOML data
const DEFAULT_CONF = `
[lang.go]
exts = ['.go']
compile = 'go build -o {{.Join .Dest .Out}} {{.Src}}'
run = '{{.Join .Dest .Out}}'

[lang.ruby]
exts = ['.rb']
run = 'ruby {{.Src}}'

[lang.scala]
exts = ['.scala']
compile = 'scalac -d {{.Dest}} {{.Src}}'
run = 'scala -cp {{.Dest}} Main'

[lang.java]
exts = ['.java']
compile = 'javac -d {{.Dest}} {{.Src}}'
run = 'java -cp {{.Dest}} Main'

[lang.rust]
exts = ['.rs']
compile = 'rustc -o {{.Join .Dest .Out}} {{.Src}}'
run = '{{.Join .Dest .Out}}'

# XXX: What is a most common C++ compiler?
[lang.'c++']
exts = ['.cpp']
compile = 'c++ -o {{.Join .Dest .Out}} {{.Src}}'
run = '{{.Join .Dest .Out}}'

[lang.javascript]
exts = ['.js']
run = 'node {{.Src}}'
`
