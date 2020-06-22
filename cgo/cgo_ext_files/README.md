
# CGO from external files

Simulating a C library called from Go.

Execute the following commands from the root of this folder
```bash
## Compile the c code
gcc -c cLib/*.c

## Generate C library archive file
ar rs cLib.a *.o
```