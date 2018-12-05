CXXFLAGS :=	-Werror -Wall -std=c++14 -lre2 -g -O0 -fsanitize=address

TARGET := guards
all: $(TARGET)
	-rm -f core
	./$(TARGET) example.txt
	# ./$(TARGET) $(TARGET).txt

$(TARGET): $(TARGET).cc
	clang++ $(CXXFLAGS) -o $@ $<

$(TARGET)-hs: $(TARGET).hs
	ghc -o $@ $<

clean:
	rm -f *.o *.hi

realclean: clean
	rm -f $(TARGET) $(TARGET)-hs

.PHONY: all clean realclean
