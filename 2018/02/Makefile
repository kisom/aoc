TARGET := inventory
all: $(TARGET) $(TARGET)-hs clean
	./$(TARGET) $(TARGET).txt
	./$(TARGET)-hs

$(TARGET): $(TARGET).cc
	clang++ $(CXXFLAGS) -o $@ $<

$(TARGET)-hs: $(TARGET).hs
	ghc -o $@ $<

clean:
	rm -f *.o *.hi

realclean: clean
	rm -f $(TARGET) $(TARGET)-hs

.PHONY: all clean realclean
