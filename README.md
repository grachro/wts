# wts

wts is 'translate Windows path to Samba path' command.

## how to use

```
$ wts '\\share-server\foo\bar'
smb://share-server/foo/bar
```

## build

```
go build wts.go
```
