#ifndef ARIPHMETIC_OPERATOR_H
#define ARIPHMETIC_OPERATOR_H

namespace s21 {
template <typename T>
class ariphmetic_operator {
 private:
 public:
  ariphmetic_operator();
  T addition_infix(T, T);
  T addition_prefix(T, T);
  T addition_postfix(T, T);
};
}  // namespace s21

#include "ariphmetic_operator.tpp"

#endif