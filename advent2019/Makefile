DAYS := $(shell ls | grep day..)
all:
	for day in $(DAYS) ;			\
	do 					\
		( cd $$day && go run $$day.go );	\
	done

test:
	go test ./...


.PHONY: all test
