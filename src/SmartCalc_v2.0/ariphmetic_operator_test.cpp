#include "ariphmetic_operator.h"

#include <gtest/gtest.h>

TEST(summa, infixSum) {
  s21::ariphmetic_operator<int> inf;
  int rez = inf.addition_infix(5, 10);

  EXPECT_EQ(rez, 15);
}

TEST(summa, prefixSum) {
  s21::ariphmetic_operator<int> inf;
  int rez = inf.addition_prefix(5, 10);

  EXPECT_EQ(rez, 15);
}

TEST(summa, postfixSum) {
  s21::ariphmetic_operator<int> inf;
  int rez = inf.addition_postfix(5, 10);

  EXPECT_EQ(rez, 15);
}


int main(int ac, char **av) {
  ::testing::InitGoogleTest(&ac, av);
  int t = RUN_ALL_TESTS();
  return (t);
}