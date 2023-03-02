### base

Encode to any base, 2-62, in your shell.


### Install

Prebuilt binaries are available for the following archs - [grab one here](https://github.com/ok-john/base/releases/tag/v1.0.0).

```
darwin-amd64  freebsd-amd64  freebsd-arm64  linux-amd64  linux-arm64  linux-mips64    linux-ppc64    linux-s390x    windows-arm5   windows-arm7
darwin-arm64  freebsd-arm5   freebsd-arm7   linux-arm5   linux-arm7   linux-mips64le  linux-ppc64le  windows-386    windows-arm6
freebsd-386   freebsd-arm6   linux-386      linux-arm6   linux-mips   linux-mipsle    linux-riscv64  windows-amd64  windows-arm64
```

Or compile from source, after git cloning run:

```bash
$ make
$ make install
```

### Usage

```bash
$ base -h
```

base works great with brace expansions, here you can
generate bash to view every base between 2-62.

```bash
$ echo -e "echo ok-john|base "{2..62}"\n"
```

now run it

```bash
$ echo -e "echo ok-john|base "{2..62}"\n"|bash
```

For example, hexadecimal:

```bash
$ echo 'ok-john'|base 16
> 6f6b2d6a6f686e0a
```

Verify with:

```bash
$ echo 'ok-john'|xxd -ps
> 6f6b2d6a6f686e0a
```

