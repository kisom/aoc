#include <cassert>
#include <chrono>
#include <fstream>
#include <iostream>
#include <numeric>
#include <sstream>
#include <vector>

using namespace std;


class Node {
public:
	vector<Node>	children;
	vector<int>	metadata;
};


static Node
readNode(istream &ins)
{
	Node	node;

	int	nchildren;
	int	nmetadata;
	int	metadata;

	ins >> nchildren;
	ins >> nmetadata;

	node.children.reserve(nchildren);
	node.metadata.reserve(nmetadata);
	for (auto i = 0; i < nchildren; i++) {
		auto child = readNode(ins);
		node.children.push_back(child);
	}

	for (auto i = 0; i < nmetadata; i++) {
		ins >> metadata;
		node.metadata.push_back(metadata);
	}

	return node;
}


static string TEST_LICENSE = "2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2";


static void
selfCheckReadNode()
{
	stringstream	str(TEST_LICENSE);

	auto node = readNode(str);
	assert(node.children.size() == 2);
	assert(node.metadata.size() == 3);

	int	expected[3] = {1, 1, 2};
	for (auto i = 0; i < 3; i++) {
		assert(node.metadata[i] == expected[i]);
	}
}


static int
sumMetadata(Node &node)
{
	int	sum = 0;

	sum = accumulate(node.metadata.begin(), node.metadata.end(), 0);
	for (auto i = 0; i < node.children.size(); i++) {
		sum += sumMetadata(node.children[i]);
	}

	return sum;
}

static void
selfCheckSumMetadata()
{
	stringstream	str(TEST_LICENSE);
	int		expected = 138;

	auto	node = readNode(str);
	auto	sum = sumMetadata(node);

	assert(sum == expected);
}


static int
nodeValue(Node &node)
{
	if (node.children.size() == 0) {
		return accumulate(node.metadata.begin(), node.metadata.end(), 0);
	}

	int	value = 0;
	for (auto i = 0; i < node.metadata.size(); i++) {
		auto ref = node.metadata[i] - 1;

		if (ref == -1) {
			continue;
		}

		if (ref >= node.children.size()) {
			continue;
		}

		// Could memoize with a map but with as small of a map
		// as we have, it's not worth it
		value += nodeValue(node.children[ref]);
	}

	return value;
}


static void
selfCheckNodeValue()
{
	stringstream	str(TEST_LICENSE);
	int		expected = 66;

	auto	node = readNode(str);
	auto	value = nodeValue(node);
	assert(expected == value);
}


static void
selfCheck()
{
	selfCheckReadNode();
	selfCheckSumMetadata();
	selfCheckNodeValue();
}

int
main(int argc, char *argv[])
{
	auto start = chrono::system_clock::now();
	selfCheck();

	for (int i = 1; i < argc; i++) {
		ifstream	licenseFile(argv[i]);

		auto node = readNode(licenseFile);
		auto sum = sumMetadata(node);
		cout << argv[i] << ": sum=" << sum;

		auto value = nodeValue(node);
		cout << ", root value=" << value << endl;
	}

	auto finished = chrono::system_clock::now();
	auto span = chrono::duration_cast<chrono::microseconds>(finished - start);
	cout << "solution(s) generated in " << span.count() << "µs" << endl;
}
