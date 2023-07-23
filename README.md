# go-wallhaven

<div align="center">

[![Build Status](https://img.shields.io/github/checks-status/coderj001/go-wallhaven/main?color=black&style=for-the-badge&logo=github)][github-actions]
[![Code Coverage](https://img.shields.io/codecov/c/github/coderj001/go-wallhaven?color=blue&logo=codecov&style=for-the-badge)][github-actions-tests]
[![Security: bandit](https://img.shields.io/badge/Security-GoSec-lightgrey?style=for-the-badge&logo=springsecurity)](https://github.com/securego/gosec)
[![Dependencies Status](https://img.shields.io/badge/Dependencies-Up%20to%20Date-brightgreen?style=for-the-badge&logo=dependabot)][dependabot-pulls]
[![Semantic Versioning](https://img.shields.io/badge/versioning-semantic-black?style=for-the-badge&logo=semver)][github-releases]
[![Pre-Commit Enabled](https://img.shields.io/badge/Pre--Commit-Enabled-blue?style=for-the-badge&logo=pre-commit)][precommit-config]
[![License](https://img.shields.io/github/license/coderj001/go-wallhaven?color=red&style=for-the-badge)][project-license]
[![Go v1.18](https://img.shields.io/badge/Go-%20v1.18-black?style=for-the-badge&logo=go)][gomod-file]

</div>

## Installation (Binary)
Download binary from [releases](https://github.com/coderj001/go-wallhaven/releases) or build with `go build .`.
If you have api key, it free and will suggest to create a free account, create a file `.go-wallhaven` at home dir
```json
{
  "API_KEY" : "<API_KEY>",
  "DIR" : "<DIRECTORY YOU WANTED TO BE SAVE>",
  "USER_NAME" : "<USER_NAME>"
}
```


## Initial Setup

This section is intended to help developers and contributors get a working copy of
`go-wallhaven` on their end

<details>
<summary>
    1. Clone this repository
</summary><br>

```sh
git clone https://github.com/coderj001/go-wallhaven
cd go-wallhaven
```
</details>

<details>
<summary>
    2. Install `golangci-lint`
</summary><br>

Install `golangci-lint` from the [official website][golangci-install] for your OS
</details>
<br>


## Local Development

This section will guide you to setup a fully-functional local copy of `go-wallhaven`
on your end and teach you how to use it! Make sure you have installed
[golangci-lint][golangci-install] before following this section!

**Note:** This section relies on the usage of [Makefile][makefile-official]. If you
can't (or don't) use Makefile, you can follow along by running the internal commands
from [`go-wallhaven's` Makefile][makefile-file] (most of which are
OS-independent)!

### Installing dependencies

To install all dependencies associated with `go-wallhaven`, run the
command

```sh
make install
```



### Using Code Formatters

Code formatters format your code to match pre-decided conventions. To run automated code
formatters, use the Makefile command

```sh
make codestyle
```

### Using Code Linters

Linters are tools that analyze source code for possible errors. This includes typos,
code formatting, syntax errors, calls to deprecated functions, potential security
vulnerabilities, and more!

To run pre-configured linters, use the command

```sh
make lint
```

### Running Tests

Tests in `go-wallhaven` are classified as *fast* and *slow* - depending
on how quick they are to execute.

To selectively run tests from either test group, use the Makefile command

```sh
make fast-test

OR

make slow-test
```

Alternatively, to run the complete test-suite -- i.e. *fast* and *slow* tests at one
go, use the command

```sh
make test
```

### Running the Test-Suite

The *`test-suite`* is simply a wrapper to run linters, stylecheckers and **all** tests
at once!

To run the test-suite, use the command

```sh
make test-suite
```

In simpler terms, running the test-suite is a combination of running [linters](#using-code-linters)
and [all tests](#running-tests) one after the other!
<br>


## Additional Resources

### Makefile help

<details>
<summary>
    Tap for a list of Makefile commands
</summary><br>

|     Command    	|                               Description                               	| Prerequisites 	|
|:--------------:	|:-----------------------------------------------------------------------:	|:-------------:	|
|     `help`     	| Generate help dialog listing all Makefile commands with description     	|       NA      	|
|    `install`   	| Fetch project dependencies                                              	|       NA      	|
|   `codestyle`  	| Run code-formatters                                                     	| golangci-lint 	|
|     `lint`     	| Check codestyle and run linters                                         	| golangci-lint 	|
|     `test`     	| Run **all** tests                                                       	|       NA      	|
|  `fast-tests`  	| Selectively run *fast* tests                                            	|       NA      	|
|  `slow-tests`  	| Selectively run *slow* tests                                            	|       NA      	|
|  `test-suite`  	| Check codestyle, run linters and **all** tests                        	| golangci-lint 	|
|      `run`     	| Run *go-wallhaven*                           	|       NA      	|
|  `docker-gen`  	| Create production docker image for *go-wallhaven* 	|     docker    	|
|  `docker-debug`  	| Create debug-friendly docker image for *go-wallhaven* 	|     docker    	|
| `clean-docker` 	| Remove docker image generated by `docker-gen`                           	|     docker    	|

<br>
</details>

Optionally, to see a list of all Makefile commands, and a short description of what they
do, you can simply run

```sh
make
```

Which is equivalent to;

```sh
make help
```

Both of which will list out all Makefile commands available, and a short description
of what they do!

### Generating Binaries

To generate binaries for multiple OS/architectures, simply run

```sh
bash build-script.sh
```

The command will generate binaries for Linux, Windows and Mac targetting multiple
architectures at once! The binaries, once generated will be stored in the `bin`
directory inside the project directory.

The binaries generated will be named in the format

```text
go-wallhaven_<version>_<target-os>_<architecture>.<extension>
```

The `<extension>` is optional. By default, `version` is an empty string. A custom
version can be passed as an argument while running the script. As an example;

```sh
bash build-script.sh v1.2.1
```

An example of the files generated by the previous command will be;

```sh
go-wallhaven_v1.2.1_windows_x86_64.exe
```

### Setting up Codecov

Follow [Codecov Docs][codecov-docs] to activate Codecov for your project repository.

Once you've activated Codecov for your project, you should be able to see the
`Repository Upload Token`. Copy this token, and add it as a secret to your Github
repository. Checkout [Creating secrets for a Repository][creating-secrets] for info
on how to add secrets on Github.

For the secret, the name of the secret should be "`CODECOV_TOKEN`" (no spaces,
copy-paste the string as it is). The value of the secret would be the
`Repository Upload Token` obtained from Codecov.

Save the secret, you should be able to a secret named `CODECOV_TOKEN` in the
*Settings > Secrets* section of your project repository. If this field is visible,
you are done with setting up Codecov, and should be able to see code coverage reports
after the next run of CI pipeline!

### Using Docker

To run `go-wallhaven` in a docker container, read the instructions in
[docker section](./docker).

### Running `go-wallhaven`

To run go-wallhaven, use the command

```sh
make run
```

Additionally, you can pass any additional command-line arguments (if needed) as the
argument "`q`". For example;

```sh
make run q="--help"

OR

make run q="--version"
```
<br>


## Releases

You can check out a list of previous releases on the [Github Releases][github-releases]
page.

### Semantic versioning with Release Drafter

<details>
    <summary>
        What is Semantic Versioning?
    </summary><br>

Semantic versioning is a versioning scheme aimed at making software management easier.
Following semantic versioning, version identifiers are divided into three parts;

```sh
    <major>.<minor>.<patch>
```

> MAJOR version when you make incompatible API changes [breaking changes]<br>
> MINOR version when you add functionality in a backwards compatible manner [more features]<br>
> PATCH version when you make backwards compatible bug fixes [bug fixes and stuff]<br>

For a more detailed description, head over to [semver.org][semver-link]

</details>

[Release Drafter][release-drafter] automatically updates the release version as pull
requests are merged.

Labels allowed;

 - `major`: Affects the `<major>` version number for semantic versioning
 - `minor`, `patch`: Affects the `<patch>` version number for semantic versioning

Whenever a pull request with one of these labels is merged to the `master` branch,
the corresponding version number will be bumped by one digit!

### List of Labels

Pull requests once merged, will be classified into categories by
[release-drafter][release-drafter] based on pull request labels

This is managed by the [`release-drafter.yml`][release-drafter-config] config file.

|                        **Label**                        	|      **Title in Releases**      	|
|:-------------------------------------------------------:	|:-------------------------------:	|
|                        `security`                       	|         :lock: Security         	|
|           `enhancement`, `feature`,  `update`           	|         :rocket: Updates        	|
|                  `bug`, `bugfix`, `fix`                 	|         :bug: Bug Fixes         	|
|                 `documentation`, `docs`                 	|       :memo: Documentation      	|
| `wip`, `in-progress`, `incomplete`, `partial`, `hotfix` 	| :construction: Work in Progress 	|
|               `dependencies`, `dependency`              	|      :package: Dependencies     	|
|      `refactoring`, `refactor`, `tests`, `testing`      	|  :test_tube: Tests and Refactor 	|
|                `build`, `ci`, `pipeline`                	|   :robot: CI/CD and Pipelines   	|

The labels `bug`, `enhancement`, and `documentation` are automatically created by Github
for repositories. [Dependabot][dependabot-link] will implicitly create the
`dependencies` label with the first pull request raised by it.

The remaining labels can be created as needed!
<br>

## TODO:
1. Limit the number of go worker should be user input
2. Page input should be a range of pages (need to rewrite apifech function & others)
3. Resolution shlould be fzf style multi-select (will try to apply on other fields)
: 
: 
