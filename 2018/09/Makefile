TARGET := 	marbles

CXXFLAGS :=	-Werror -Wall -std=c++14 -lre2
FLAVOR   ?=	DEBUG

ifeq ($(FLAVOR),RELEASE)
CXXFLAGS += -O3
else
CXXFLAGS += -g -O0 -fsanitize=address
endif

.PHONY: all
all: $(TARGET)
	time ./$(TARGET) $(TARGET).txt
	# time ./$(TARGET).py $(TARGET).txt

print-%: ; @echo $*=$($*)

$(TARGET): $(TARGET).cc
	clang++ $(CXXFLAGS) -o $(TARGET) $(TARGET).cc

.PHONY: clean
clean:
	rm -f *.o *.hi

.PHONY: realclean
realclean: clean
	rm -f $(TARGET) $(TARGET)-hs

.PHONY: all clean realclean

