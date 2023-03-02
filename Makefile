CC := CGO_ENABLED=0 go
NAME := base
INSTALL_DIR := /usr/local/bin

build ::
	$(CC) build -o $(NAME)

install ::
	mv $(NAME) $(INSTALL_DIR)

