#include <stdio.h>

int globale = 5 ;
int max(int a, int b)
{
int c=a;
if (a<b){ c=b; };
return c;
}

int f(int x)
{
printf("blop!\n");
return 2*x+max(x+4,2*x);
}

int main()
{
printf("%d et %d\n",f(3),f(5));
return 0;
}

