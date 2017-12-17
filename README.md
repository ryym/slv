# slv

- Manage programming problem solution sources with test cases
- The programs take input from stdin and output to stdout
- solution can be written in several languages

```bash
$ slv new hello
$ cd hello
$ cat << EOF > src/hello.rb
name = gets
puts "hello, #{name}."
EOF
$ cat << EOF > test/test.toml
[[test]]
in = "alice"
out = "hello, alice."
[[test]]
in = "bob"
out = "hello, bob."
EOF
$ slv test src/hello.rb
[OK] All: 2, Passed: 2, Failed: 0 
```
