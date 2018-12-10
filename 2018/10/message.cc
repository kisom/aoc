#include <cassert>
#include <fstream>
#include <iostream>
#include <vector>

using namespace std;

#include <re2/re2.h>

static re2::RE2	coordinate("position= ?<(\\d+),  ?(\\d+)> velocity=< ?(\\d+),  ?(\\d+)>");

class Point {
public:
	Point(string line) {
		auto match = re2::RE2::FullMatch(line, coordinate, &x, &y, &dx, &dy);
		assert(match);
	}

	void step() {
		x += dx;
		y += dy;
	}

	int	x;
	int	y;
private:
	int	dx;
	int	dy;
};

vector<string>
readLines(const string path)
{
       ifstream        file(path);
       vector<string>  lines;

       for (string line; getline(file, line); ) {
               lines.push_back(line);
       }
       file.close();

       return lines;
}


vector<Point>
readPoint(const string path)
{
	auto lines = readLines(path);

	vector<Point>	points;
	points.reserve(lines.size());

	for (auto line : lines) {
		Point	point(line);
		points.push_back(point);
	}

	return points;
}


int
main(int argc, char *argv[])
{
	for (auto i = 1; i < argc; i++) {
		auto lines = readLines(string(argv[i]));
	}
}
