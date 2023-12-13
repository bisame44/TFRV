CPP=mpic++
CFLAGS=-O0 -Wall -std=c++17
LDFLAGS=
SOURCES=src/main.cpp
OBJECTS=$(SOURCES:.cpp=.o)
EXECUTABLE=conntest

all: $(SOURCES) $(EXECUTABLE)

$(EXECUTABLE): $(OBJECTS)
	$(CPP) $(LDFLAGS) $(OBJECTS) -o $@

.cpp.o:
	$(CPP) -c $(CFLAGS) $< -o $@
