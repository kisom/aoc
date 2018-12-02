#include <cassert>
#include <fstream>
#include <iostream>
#include <set>
#include <string>
#include <vector>

typedef struct {
	int	two;
	int	three;
} Inventory;

static void
compute_checksum(Inventory &inventory, std::string boxID)
{
	// This function is void because it really needs to return a tuple -
	// which could be done - and it's single-threaded, so it makes sense to
	// just operate on a reference to an Inventory.

	int	counts[26] = {0};	// assumption: there are only lower-case letters.
	int	two = 0;
	int	three = 0;

	for (char c : boxID) {
		// assumption: letters are lowercase.
		c -= 0x61;
		assert(c >= 0);
		assert(c < 26);

		// Poor man's set.
		counts[static_cast<std::size_t>(c)]++;
	}

	for (auto i = 0; i < 26; i++) {
		if (counts[i] == 2)	two++;
		if (counts[i] == 3)	three++;
	}

	if (two > 0) {
		inventory.two++;
	}
	if (three > 0) {
		inventory.three++;
	}
}

static int
compute_checksums(std::vector<std::string> &boxIDs)
{
	Inventory	inventory = {0, 0};

	for (auto boxID : boxIDs) {
		compute_checksum(inventory, boxID);
	}

	return inventory.two * inventory.three;
}

static void
self_check_checksum(void)
{
	std::vector<std::string>	boxIDs = {
		"abcdef",
		"bababc",
		"abbcde",
		"abcccd",
		"aabcdd",
		"abcdee",
		"ababab",
	};

	auto checksum = compute_checksums(boxIDs);
	assert(checksum == 12);
	std::cout << "self check OK" << std::endl;
}

static std::vector<std::string>
load_inventory(std::string path)
{
	std::ifstream			inventory_file(path);
	std::vector<std::string>	boxIDs;

	while (!inventory_file.eof()) {
		// This could be done in a streaming style, but this approach
		// reuses the function that's been tested in the self test. It
		// also allows reusing the box IDs later on.
		std::string	boxID;
		inventory_file >> boxID;
		boxIDs.push_back(boxID);
	}
	inventory_file.close();
	return boxIDs;
}

typedef struct {
	bool		matched;
	std::string	common;
} MatchResult;

static MatchResult
match_boxes(std::string boxID1, std::string boxID2)
{
	MatchResult	result;
	std::string	common = "";
	int		differences = 0;

	for (std::size_t i = 0; i < boxID1.size(); i++) {
		if (boxID1[i] == boxID2[i]) {
			common.push_back(boxID1[i]);
			continue;
		}

		differences++;
	}

	if (differences == 1) {
		result.matched = true;
		result.common = common;
	}
	else {
		result.matched = false;
	}
	return result;
}

static MatchResult
match_inventory(std::vector<std::string> boxIDs)
{
	MatchResult	result;

	for (std::size_t i = 0; i < boxIDs.size() - 1; i++) {
		for (std::size_t j = i+1; j < boxIDs.size(); j++) {
			result = match_boxes(boxIDs[i], boxIDs[j]);
			if (result.matched) {
				return result;
			}
		}
	}

	return result;
}

static void
self_check_match(void)
{
	std::vector<std::string> boxIDs = {
		"abcde",
		"fghij",
		"klmno",
		"pqrst",
		"fguij",
		"axcye",
		"wvxyz",
	};

	auto result = match_inventory(boxIDs);
	assert(result.matched);
	std::cout << "expected fgij, have " << result.common << std::endl;
	assert(result.common == "fgij");
}

static void
self_check(void)
{
	self_check_checksum();
	self_check_match();
}

int
main(int argc, char *argv[])
{
	self_check();

	for (auto i = 1; i < argc; i++) {
		auto path = std::string(argv[i]);
		auto boxIDs = load_inventory(path);
		
		// Part 1
		auto checksum = compute_checksums(boxIDs);
		std::cout << path << ":\t" << checksum << std::endl;

		// Part 2
		auto result = match_inventory(boxIDs);
		if (result.matched) {
			std::cout << "common: " << result.common << std::endl;
		}
		else {
			std::cout << "no match found" << std::endl;
		}
	}
}