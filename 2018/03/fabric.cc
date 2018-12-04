#include <cassert>
#include <fstream>
#include <iostream>
#include <set>
#include <sstream>
#include <string>
#include <vector>

#include <re2/re2.h>

using namespace std; 

class Point {
public:
	int	x;
	int	y;
	Point() : x(0), y(0) {};
	Point(int _x, int _y) : x(_x), y(_y) {};
};

inline bool
operator==(const Point &lhs, const Point &rhs)
{
	return lhs.x == rhs.x && lhs.y == rhs.y;
}

inline bool
operator!=(const Point &lhs, const Point& rhs)
{
	return !(lhs == rhs);
}

inline bool
operator<(const Point &lhs, const Point& rhs)
{
	if (lhs.y < rhs.y) {
		return true;
	}

	if (lhs.y > rhs.y) {
		return false;
	}

	if (lhs.x < rhs.x) {
		return true;
	}

	return false;
}

ostream&
operator<<(ostream& os, const Point& point)
{
	os << "(" << point.x << ", " << point.y << ")";
	return os;
}

static string	claim_regex_str = "^#(\\d+) @ (\\d+),(\\d+): (\\d+)x(\\d+)$";
static re2::RE2	claimre(claim_regex_str);

class Claim{
public:
	int	id;
	Point	upperLeft;
	Point	lowerRight;

	Claim() : id(0), upperLeft(Point(0, 0)), lowerRight(Point(0, 0)) {};
	Claim(int id, Point upperLeft, Point upperRight) : id(id), upperLeft(upperLeft), lowerRight(upperRight) {};
	void dump() { cout << "#" << id << ": (" << upperLeft.x << ", " << upperLeft.y << "), (" << lowerRight.x << ", " << lowerRight.y << ")\n"; }
};

inline bool
operator==(const Claim &lhs, const Claim &rhs)
{
	return (lhs.id == rhs.id) && (lhs.upperLeft == rhs.upperLeft) && (lhs.lowerRight == rhs.lowerRight);
}

inline const bool
operator!=(const Claim &lhs, const Claim& rhs)
{
	return !(lhs == rhs);
}

ostream&
operator<<(ostream &os, const Claim &claim)
{
	os << "#" << claim.id << " " << claim.upperLeft << ", " << claim.lowerRight;
	return os;
}

static Claim
process_claim(string claimString)
{
	Claim	claim;

	RE2::FullMatch(claimString, claimre, &claim.id, &claim.upperLeft.x,
	    &claim.upperLeft.y, &claim.lowerRight.x, &claim.lowerRight.y);
	claim.lowerRight.x += claim.upperLeft.x;
	claim.lowerRight.y += claim.upperLeft.y;

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

static int
count_claim(Claim &claim, set<Point> &counter)
{
	int	overlapping = 0;

	for (auto j = claim.upperLeft.y; j < claim.lowerRight.y; j++) {
		for (auto i = claim.upperLeft.x; i < claim.lowerRight.x; i++) {
			auto point = Point(i, j);
			counter.insert(point);
			if (counter.count(point) == 2) {
				overlapping++;
			}
		}
	}

	return overlapping;
}

static int
count_claims(vector<Claim> &claims, set<Point> &counter)
{
	int	overlapping = 0;

	for (auto claim : claims) {
		overlapping += count_claim(claim, counter);
	}
	cout << "counter has seen " << counter.size() << " points." << endl;
	return overlapping;
}	

static void
self_check()
{
	// Part 1
	self_check_overlapping_claim();
}

static vector<Claim>
load_claims(string path)
{
	ifstream	claim_file(path);
	vector<Claim>	claims;

	for (string line; getline(claim_file, line); ) {
		claims.push_back(process_claim(line));
	}
	claim_file.close();

	return claims;
}

int
main(int argc, char *argv[])
{
	self_check();
	for (auto i = 1; i < argc; i++) {
		auto claims = load_claims(argv[i]);
		assert(claims.size() > 0);
		cout << "loaded " << claims.size() << " claims." << endl;

		set<Point> counter;
		auto count = count_claims(claims, counter);
		cout << "overlapping area: " << count << endl;
	}
}
