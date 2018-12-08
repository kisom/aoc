#include <algorithm>
#include <cassert>
#include <chrono>
#include <iostream>
#include <fstream>
#include <map>
#include <set>
#include <vector>

using namespace std;


class Problem {
public:
	Problem() {};
	Problem(vector<string> lines)
	{
		for (auto line : lines) {
			char	dependency = line[5];
			char	name = line[36];
			this->constraint(name, dependency);
		}
	}

	void
	constraint(char name, char dependency)
	{
		// Make sure both the name and constraint are accounted
		// for. Without this, the starting step won't be recorded.
		if (constraints.count(name) == 0) {
			constraints[name] = set<char>();
		};

		if (constraints.count(dependency) == 0) {
			constraints[dependency] = set<char>();
		}

		constraints[name].insert(dependency);
		// auto temp = constraints[name];
		// temp.insert(dependency);
		// constraints[name] = temp;
	}

	void
	complete(char dependency)
	{
		for (auto step : constraints) {
			if (step.second.count(dependency) == 0) {
				continue;
			}
			constraints[step.first].erase(dependency);
		}
	}

	string
	solve()
	{
		string	result;
		while (constraints.size() > 0) {
			for (auto it : constraints) {
				if (it.second.size() != 0) {
					continue;
				}

				constraints.erase(it.first);
				complete(it.first);
				result.push_back(it.first);
				break;
			}
		}

		return result;
	}

	size_t
	size()
	{
		return this->constraints.size();
	}

private:
	map<char, set<char>>	constraints;
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


const vector<string>	TestSteps = {
	"Step C must be finished before step A can begin.",
	"Step C must be finished before step F can begin.",
	"Step A must be finished before step B can begin.",
	"Step A must be finished before step D can begin.",
	"Step B must be finished before step E can begin.",
	"Step D must be finished before step E can begin.",
	"Step F must be finished before step E can begin."
};


void
self_check_problem_ctor()
{
	Problem		problem(TestSteps);

	assert(problem.size() == 6);
}


void
self_check_solve()
{
	auto expected = "CABDFE";
	Problem		testCase(TestSteps);

	auto result = testCase.solve();
	assert(expected == result);
}


static void
self_check()
{
	self_check_problem_ctor();
	self_check_solve();
	cerr << "self check: OK" << endl;
}


int
main(int argc, char *argv[])
{
	auto start = chrono::system_clock::now();
	self_check();

	for (auto i = 1; i < argc; i++) {
		auto lines = readLines(argv[i]);

		Problem	problem(lines);
		cout << problem.size() << " tasks recorded." << endl;

		auto result = problem.solve();
		cout << "solution: " << result << endl;
	}

	auto finished = chrono::system_clock::now();
	auto span = chrono::duration_cast<chrono::microseconds>(finished - start);
	cout << "solution(s) generated in " << span.count() << "µs" << endl;
}
