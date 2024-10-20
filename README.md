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

## Installation

- [GitHub Releases](https://github.com/ryym/slv/releases)
- `go get -u github.com/ryym/slv`

## Overview

```
NAME:
   slv - Helps you solve programming problems

USAGE:
   slv [global options] command [command options] [arguments...]

VERSION:
   0.2.0

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

## Usage

First, create a new problem directory.

```bash
$ slv new hello
$ ls hello
src/ test/
$ cd hello
```

Next, write your solution and save it in the `src` directory.
Any file name is fine as long as the extension is correct.

```ruby
# src/hello.rb

name = gets.chomp
puts "hello, #{name}."
```

Then prepare the test cases in the `test` directory.
Any file name is fine again.

```toml
# test/test.toml

[[test]]
in = "alice"
out = "hello, alice."

[[test]]
in = "bob"
out = "hello, bob."
```

Now you can test your solution!

```bash
$ slv test src/hello.rb
testing hello.rb...
..

[OK] All: 2, Passed: 2, Failed: 0 
```

You can also specify the solution by the language name or its extension.

```bash
$ slv test ruby
testing hello.rb...
..

[OK] All: 2, Passed: 2, Failed: 0 

# Or

$ slv test rb
testing hello.rb...
..

[OK] All: 2, Passed: 2, Failed: 0 
```

When your test fails, Slv displays the output diff.
For example, add a new test case which expects a different output:

```toml
# test/test2.toml
# (You can store any number of test case files in the `test` directory)

[[test]]
in = "anonymous"
out = "what your name?"
```

This results in the error like this:

```diff
$ slv test rb
testing hello.rb...
..F

test2.toml[0]:

-what your name?
+hello, anonymous.


[FAILED] All: 3, Passed: 2, Failed: 1
```

## How to write test cases

Test cases are loaded from all TOML files in the `test` directory.
See the [sample test cases](sample_cases.toml) as a reference.

## Customize

Slv supports some major languages by default as [TOML config][default-langs].
You can use any language with Slv by adding the configuration to `.slv.toml`.
Slv searches `.slv.toml` from current directory and its all parent directories.


[default-langs]: https://github.com/ryym/slv/blob/master/slv/config.go

## Development

```sh
git clone https://github.com/ryym/slv
go test ./...
```
