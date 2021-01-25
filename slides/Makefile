SOURCES = begin.slide day1.slide day2.slide day3.slide end.slide
TARGET = mastering-go.slide


run: $(TARGET)
	present

$(TARGET): $(SOURCES)
	cat $(SOURCES) > $(TARGET)

clean:
	-rm $(TARGET)
