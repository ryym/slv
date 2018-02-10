# Slv

[![circleci](https://circleci.com/gh/ryym/slv.svg?style=svg)](https://circleci.com/gh/ryym/slv)
[![appveyor](https://ci.appveyor.com/api/projects/status/8e2o0r8bgcfobxmi?svg=true)](https://ci.appveyor.com/project/ryym/slv)

Slv is a command line tool for managing source files of common programming contests, along with test cases.
Slv can manage any programs that take inputs from stdin and output results to stdout regardless of language. 
You can:

- define test cases as TOML files
- write source code in multiple languages

For example, Slv can be used to manage solution code for
[AOJ][aoj] and [AtCoder][at-coder] holding online programming contests,
or for [CodeIQ][code-iq] which provides various programming problems,
because they require a program that uses stdin / stdout.

[aoj]: http://judge.u-aizu.ac.jp/onlinejudge/index.jsp
[at-coder]: https://atcoder.jp/?lang=en
[code-iq]: https://codeiq.jp/

## Usage

```
NAME:
   slv - Helps you solve programming problems

USAGE:
   slv [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
     new, n   Create new problem directory
     test, t  Run tests for the specified source code
     compile  Compile without running
     run, r   Run the specified source code
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

Example:

```bash
# First, create new problem directory.
$ slv new hello
$ cd hello
$ ls hello
src/ test/

# Next, write your solution.
$ cat << EOF > src/hello.rb
name = gets.chomp
puts "hello, #{name}."
EOF

# And prepare test cases.
$ cat << EOF > test/test.toml
[[test]]
in = "alice"
out = "hello, alice."
[[test]]
in = "bob"
out = "hello, bob."
EOF

# Now you can test your solution!
$ slv test src/hello.rb
..

[OK] All: 2, Passed: 2, Failed: 0 

# You can also specify the source file by the language name.
$ slv test ruby
..

[OK] All: 2, Passed: 2, Failed: 0 
```

Test cases are loaded from all TOML files in the `test` directory.

## Customize

Slv supports some major languages by default as [TOML config][default-langs].
You can use any language with Slv by adding the configuration to `.slv.toml`.
Slv searches `.slv.toml` from current directory and its all parent directories.


[default-langs]: https://github.com/ryym/slv/blob/master/slv/config.go

## Installation

[GitHub Releases](https://github.com/ryym/slv/releases)

## Development

```sh
git clone https://github.com/ryym/slv
go get -u github.com/golang/dep/cmd/dep
dep ensure
```
