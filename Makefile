NAME = ft_otp

GO = go
GOFMT = gofmt
GOFLAGS = -v
SRC_DIR = cmd
MAIN = $(SRC_DIR)/main.go

all: $(NAME)

$(NAME): $(MAIN)
	$(GO) build $(GOFLAGS) -o $(NAME) $(MAIN)

clean:
	rm -f $(NAME)

fclean: clean
	go clean -cache

re: fclean all

fmt:
	$(GOFMT) -w .

.PHONY: all clean fclean re fmt 