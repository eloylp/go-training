#include <stdio.h>
#include "lib/lib.h"

int main(int argc, char **argv){

    GoString message;
    message.p = "message \n";
    message.n = 10;
    PrintMessage(message);
}
