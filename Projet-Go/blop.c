#include<stdio.h>

int f(int x) {
  if (x==0) return 1;
  return f(x-1)+f(x-2);
}

int main() {
  int x=10;
  int a=5;

  printf("f(%d)+f(%d)=%d\n",a,x,f(a)+f(x));
}
