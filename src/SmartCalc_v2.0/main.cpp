#include <iostream>

#include "ariphmetic_operator.h"

int main() {
  s21::ariphmetic_operator<int> sum;
  int rez = sum.addition_infix(5, 10);
  std::cout << rez << std::endl;
  return (0);
}