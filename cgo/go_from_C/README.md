
# Calling Go from C
Must build a c shared lib

```bash
pushd .
cd lib
go build -o lib.o -buildmode=c-shared main.go
popd
gcc -o program main.c ./lib.o
```