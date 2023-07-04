#ifdef ARIPHMETIC_OPERATOR_H

template <typename T>
s21::ariphmetic_operator<T>::ariphmetic_operator() {}

template <typename T>
T s21::ariphmetic_operator<T>::addition_infix(T a, T b) {
  return a + b;
}

template <typename T>
T s21::ariphmetic_operator<T>::addition_prefix(T a, T b) {
  return a + b;
}

template <typename T>
T s21::ariphmetic_operator<T>::addition_postfix(T a, T b) {
  return a + b;
}


#endif