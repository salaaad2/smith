/*********************************/
/*   SMITH            (  //      */
/*   smith             ( )/      */
/*   by salade         )(/       */
/*  ________________  ( /)       */
/* ()__)____________)))))   :^}  */
/*********************************/

#include "smith.h"
#include <sys/time.h>

char *
mr_smith()
{
    return ("---SMITH_V0.0.0---");
}

long
c_get_timestamp()
{
    struct timeval tv;
    long ct;

    gettimeofday(&tv, NULL);
    ct = ((tv.tv_sec * 1000) + (tv.tv_usec / 1000));
    return (ct);
}
