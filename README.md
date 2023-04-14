# mdt
touch with mkdir -p

# motivation
if you want to create a file that does not exist, you can use `touch` command.

```bash
touch filename
```

but touch command cannot create file with directory.

```bash
# this command fails if directory does not exist.
touch path/to/file
```

so I created this command. this command can create file with directory.


```bash
mdt path/to/file
```

# install

```bash
go install github.com/mkusaka/mdt@latest
```

# usage

```bash
mdt path/to/file
```

```bash
mdt filename
```
