TARGET := calibration
all: $(TARGET) $(TARGET)-hs clean
	./$(TARGET) < calibration.txt
	./$(TARGET)-hs

$(TARGET): $(TARGET).cc
	clang++ $(CXXFLAGS) -o $@ $<

$(TARGET)-hs: $(TARGET).hs
	ghci -o $@ $<

clean:
	rm -f *.o *.hi

realclean: clean
	rm -f $(TARGET) $(TARGET)-hs

.PHONY: all clean realclean
