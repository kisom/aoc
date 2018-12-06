#include <algorithm>
#include <cassert>
#include <fstream>
#include <iostream>
#include <set>
#include <sstream>
#include <string>
#include <stack>

using namespace std;


static int
react(istream &in)
{
	char		current = 0;
	char		last = 0;
	stack<char>	result;

	in >> last;
	result.push(last);

	in >> current;
	while (!in.eof()) {
		if (!result.empty()) {
			if ((current ^ 0x20) == last) {
				result.pop();
				if (!result.empty()) {
					last = result.top();
				}
			}
			else {
				last = current;
				result.push(last);
			}
		}
		else {
			last = current;
			result.push(last);
		}

		in >> current;
	}

	return result.size();
}


static void
self_check_react()
{
	string		chain = "dabAcCaCBAcCcaDA";
	string		expected = "dabCBAcaDA";
	stringstream	chainStream(chain);
	
	assert(react(chainStream) == static_cast<int>(expected.size()));
}


static string
stripBase(const string s, char base)
{
	string sCopy;
	auto lcase = static_cast<char>(base ^ 0x20);

	sCopy.reserve(s.size());
	copy(s.begin(), s.end(), back_inserter(sCopy));
	sCopy.erase(remove(sCopy.begin(), sCopy.end(), base), sCopy.end());
	sCopy.erase(remove(sCopy.begin(), sCopy.end(), lcase), sCopy.end());
	return sCopy;
}


static int
reactAll(string chain)
{
	set<char>	bases;
	int		minChain = -1;
	char		minBase;

	for (auto x : chain) {
		if (x >= 'a') {
			x ^= 0x20;	
		}
		bases.insert(x);
	}

	for (auto base : bases) {
		string	chainCopy = stripBase(chain, base);

		stringstream chainStream(chainCopy);
		auto result = react(chainStream);
		if (minChain == -1 || result < minChain) {
			minChain = result;
			minBase = base;
		}
	}

	return minChain;
}


static void
self_check_reactAll()
{
	string chain = "dabAcCaCBAcCcaDA";

	auto result = reactAll(chain);
	assert(result == 4);
}


static string
readFile(string path)
{
	ifstream	file(path);
	
	string	chain;
	file.seekg(0, ios::end);
	chain.reserve(file.tellg());
	file.seekg(0, ios::beg);
	std::copy((std::istreambuf_iterator<char>(file)),
	    std::istreambuf_iterator<char>(),
	    std::back_inserter(chain));
	file.close();

	return chain;
}


static void
self_check()
{
	self_check_react();
	self_check_reactAll();
	cout << "self check OK" << endl;
}


static int
part1(string chain)
{
	stringstream	chainFile(chain);

	return react(chainFile);
}


int
main(int argc, char *argv[])
{
	self_check();

	for (auto i = 1; i < argc; i++) {
		auto chain = readFile(string(argv[i]));
		auto result = part1(chain);

		cout << "part1: " << result << endl;

		result = reactAll(chain);
		cout << "part2: " << result << endl;
	}
}
