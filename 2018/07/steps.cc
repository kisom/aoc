#include <iostream>
#include <fstream>
#include <map>
#include <set>
#include <vector>

#include <re2/re2.h>

using namespace std;


static re2::RE2	stepRegex("Step (.) must be finished before step (.) can begin.");

class Step {
public:
	Step(char name) : name(name) {}

	void	dependsOn(char x) {
		prerequisites.insert(x);
	}

	void	completed(char x) {
		prerequisites.erase(x);
	}

	bool	ready() {
		return prerequisites.empty();
	}

	bool	requires(char x) {
		return prerequisites.count(x) > 0;
	}

	friend bool operator<(const Step &l, const Step &r) {
		return l.name < r.name;
	}

	friend ostream& operator<<(ostream& outs, const Step &step) {
		outs << "Step " << step.name << " requires:";
		for (auto x : step.prerequisites) {
			outs << " " << x;
		}
		return outs;
	}
private:
	char		name;
	set<char>	prerequisites;
};

class Sum {
	Sum() {};
	void	add(char name, char requisite) {
		if (steps.count(name) == 0) {
			auto step = new Step(name);
			step->dependsOn(requisite);
			steps[name] = step;
		} else {
			auto step = steps[name];
			step->dependsOn(requisite);
			steps[name] = step;
		}
	}
	void	complete(char name) {
		vector<char>	removals;
		for (auto it = steps.begin(); it != steps.end(); it++) {
			auto step = steps[*it];
			if (!step->requires(name)) {
				continue;
			}

			step->completed(name);
			if (step->ready()) {
				removals.push_back(*it);
			}
		}
	}
private:
	map<char, Step*>	steps;
	string			sequence;
};


vector<string>
readLines(const string path)
{
	ifstream	file(path);
	vector<string>	lines;

	for (string line; getline(file, line); ) {
		lines.push_back(line);
	}
	file.close();

	return lines;
}

void
printLines(vector<string> &lines)
{
	for (auto line : lines) {
		cout << line << endl;
	}
}


Sum *
parseLines


static void
selfCheck()
{

}

int
main(int argc, char *argv[])
{
	selfCheck();
	for (auto i = 1; i < argc; i++) {
		auto lines = readLines(argv[i]);
		printLines(lines);
	}
}
