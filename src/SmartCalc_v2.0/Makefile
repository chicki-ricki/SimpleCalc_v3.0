# NAME	=	calculator.a
NAME	=	calculator

SRC		=	main.cpp

SRC_TEST = ariphmetic_operator_test.cpp

HEADER	=	ariphmetic_operator.h ariphmetic_operator.tpp

OBJS		=	${SRC:%.cpp=%.o}

CC			=	c++

CFLAGS	=	-Wall -Wextra -Werror

RM			=	rm -f

.PHONY	:	all re clean fclean test

# %.o	:	%.cpp ${HEADER}
# 	${CC} ${CFLAGS} -c $< -o $@

${NAME}: ${OBJS} ${HEADER}
# ar rcs ${NAME} ${OBJS}
	${CC} ${CFLAGS} ${SRC} -o ${NAME}

all:	${NAME}

test:	${OBJS}
	${CC} ${SRC_TEST} -o test -std=c++17 -lgtest -lgmock

style:
	@clang-format -i *.cpp *.h *.tpp
	clang-format -n *.cpp *.h *.tpp