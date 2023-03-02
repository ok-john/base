### base

Encode to any base, 2-62, in your shell.


### Usage

```bash
    $ base -h
```

base works great with brace expansions, here you can
generate bash to view every base between 2-62.

```bash
    $ echo -e "echo ok-john|./base "{2..62}"\n"
```

now run it

```bash
    $ echo -e "echo ok-john|./base "{2..62}"\n"|bash
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

