#include <algorithm>
#include <cassert>
#include <fstream>
#include <iostream>
#include <map>
#include <vector>

#include <re2/re2.h>

using namespace std;


static re2::RE2	LineRegex("(\\d+) players; last marble is worth (\\d+) points");


class Game {
public:
	Game(string line) {
		auto match = re2::RE2::FullMatch(line, LineRegex, &nplayer, &nmarbles);
		assert(match);
	}

private:
	int		nplayer;
	int		nmarbles;
	vector<int>	marbles;
	map<int, int>	scores;
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


int
main(int argc, char *argv[])
{
	for (auto i = 1; i < argc; i++) {
		auto lines = readLines(string(argv[i]));
	}
}
