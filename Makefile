NAME = ft_otp

GO = go
SRC_DIR = cmd
MAIN = $(SRC_DIR)/main.go

all: $(NAME)

$(NAME): $(MAIN)
	$(GO) build -o $(NAME) $(MAIN)

clean:
	rm -f $(NAME)

fclean: clean
	go clean -cache

re: fclean all

.PHONY: all clean fclean re 