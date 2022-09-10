# Example

Prevent containers from stopping immediately.
```
// -i --interactive Keep STDIN open even if not attached
// -t --tty         Allocte a pseudo-TTY
// --rm             Automatically remove the container when it exits
$ docker run -it --rm alpine:3.16
```

Create non-root user.
```
// -h --home DIR           Home directory
// -g --gecos GECOS        GECOS field
// -s --shell SHELL        Login shell named SHELL by example /bin/bash
// -D --disabled-password  Don't assign a password, so cannot login
// -H --no-create-home     Don't create home directory
// -G --ingroup GRP        Group (by name)
// -u --uid UID             User id
$ adduser \
  -h "/dev/null" \
  -g "" \
  -s "/sbin/nologin" \
  -D \
  -H \
  -u 1000 \
  playerone
```
# Linux File Permissions

```
$ ls -l
total 60
drwxr-xr-x    2 root     root          4096 Aug  9 08:49 bin
```

File Types

```
- = Regular File
d = Directory
l = Symbolic Link
etc..
```