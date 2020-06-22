#include <stdio.h>
#include "cLib.h"

void  printHello(){
    printf("Hello from C \n");
}

void printMessage(char* message) {
    printf("Message is: %s (called from C)\n", message);
}
