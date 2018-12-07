CXXFLAGS :=	-Werror -Wall -std=c++14 -lre2 

ifdef NDEBUG
CXXFLAGS += -O2
else
CXXFLAGS += -g -O0 -fsanitize=address
endif

TARGET := steps

.PHONY: all
all: $(TARGET)
	time ./$(TARGET) $(TARGET).txt
	time ./$(TARGET).py $(TARGET).txt

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

