#include <unistd.h>
#include "ignore_sigchild.h"

int main(void) {
  set_not_wait_on_child();

  int year = 3600 * 24 * 365;
  while (1) {
    sleep(year);
  }
}
