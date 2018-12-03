#include <cassert>
#include <iostream>
#include <sstream>
#include <string>

class Point {
public:
	int	x;
	int	y;
	Point() : x(0), y(0) {};
	Point(int _x, int _y) : x(_x), y(_y) {};
};
inline bool operator==(const Point &lhs, const Point &rhs) {
	return lhs.x == rhs.x && lhs.y == rhs.y;
}
inline const bool operator!=(const Point &lhs, const Point& rhs){ return !(lhs == rhs); }


class Claim{
public:
	int	id;
	Point	upperLeft;
	Point	lowerRight;

	Claim() : id(0), upperLeft(Point(0, 0)), lowerRight(Point(0, 0)) {};
	Claim(int id, Point upperLeft, Point upperRight) : id(id), upperLeft(upperLeft), lowerRight(upperRight) {};
};
inline bool operator==(const Claim &lhs, const Claim &rhs) {
	return (lhs.id == rhs.id) && (lhs.upperLeft == rhs.upperLeft) && (lhs.lowerRight == rhs.lowerRight);
}
inline const bool operator!=(const Claim &lhs, const Claim& rhs){ return !(lhs == rhs); }


static Claim
process_claim(std::string claimString)
{
	Claim	claim;

	claimString = claimString.substr(1, claimString.size());
	std::stringstream	claimStream(claimString);

	claimStream >> claim.id;
	std::string	offset;
	claimStream >> offset;
	
	std::cerr << "offset\n";
	offset = offset.substr(0, offset.length() - 1);
	auto separator = offset.find(",");
	std::cerr << offset << "\t" << separator << std::endl;
	offset.replace(separator, separator+1, " ");
	
	std::cerr << "upperLeft\n";
	Point upperLeft;
	std::stringstream	offsetStream(offset);
	offsetStream >> upperLeft.x;
	offsetStream >> upperLeft.y;
	claim.upperLeft = upperLeft;

	std::cerr << "edge\n";
	std::string	edge;
	claimStream >> edge;
	separator = edge.find("x");
	edge.replace(separator, separator+1, " ");
	std::stringstream	edgeStream(edge);

	std::cerr << "lowerRight\n";
	Point lowerRight;
	edgeStream >> lowerRight.x;
	edgeStream >> lowerRight.y;
	claim.lowerRight = lowerRight;

	return claim;
}

static void
self_check_process_claim()
{
	auto claimString = "#1 @ 257,829: 10x23";
	Claim expected(1, Point(257, 829), Point(267, 852));
	auto claim = process_claim(claimString);
	assert(claim == expected);
}

static void
self_check_overlapping_claim()
{
	self_check_process_claim();
}

static void
self_check()
{
	// Part 1
	self_check_overlapping_claim();
}

int
main()
{
	self_check();
}