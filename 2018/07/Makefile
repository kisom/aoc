CXXFLAGS :=	-Werror -Wall -std=c++14 -lre2 -g -O0 -fsanitize=address

TARGET := steps

.PHONY: all
all: $(TARGET)
	time ./$(TARGET) $(TARGET).txt

print-%: ; @echo $*=$($*)

$(TARGET): $(TARGET).cc
	clang++ $(CXXFLAGS) -o $(TARGET) $(TARGET).cc

# $(TARGET)-hs: $(TARGET).hs
# 	ghc -o $@ $<

.PHONY: clean
clean:
	rm -f *.o *.hi

.PHONY: realclean
realclean: clean
	rm -f $(TARGET) $(TARGET)-hs

.PHONY: all clean realclean

