CXXFLAGS :=	-Werror -Wall -std=c++14

TARGET := fabric
all: $(TARGET)
	./$(TARGET) $(TARGET).txt

$(TARGET): $(TARGET).cc
	clang++ $(CXXFLAGS) -o $@ $< -lre2

$(TARGET)-hs: $(TARGET).hs
	ghc -o $@ $<

clean:
	rm -f *.o *.hi

realclean: clean
	rm -f $(TARGET) $(TARGET)-hs

.PHONY: all clean realclean
