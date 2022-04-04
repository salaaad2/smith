/*********************************/
/*   SMITH            (  //      */
/*   smith             ( )/      */
/*   by salade         )(/       */
/*  ________________  ( /)       */
/* ()__)____________)))))   :^}  */
/*********************************/

#include "smith.h"
#include <sys/time.h>
#include <stdio.h>
#include <stdlib.h>

char *
mr_smith()
{
    return ("---SMITH_V0.0.0---");
}

char*
get_timestamp()
{
    struct timeval tv;
    long ct;
    char * buffer;

    buffer = (char*)malloc(30);
    gettimeofday(&tv, NULL);
    ct = ((tv.tv_sec * 1000) + (tv.tv_usec / 1000));
    sprintf(buffer, "%ld", ct);
    return buffer;
}
