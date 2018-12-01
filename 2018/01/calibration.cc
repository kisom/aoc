#include <cstdlib>
#include <iostream>
#include <set>
#include <sstream>
#include <vector>

static std::vector<int>
read_deltas(std::istream &stream)
{
	std::vector<int> deltas;
	while (!stream.eof()) {
		int delta = 0;
		stream >> delta;
		deltas.push_back(delta);
	}
	return deltas;	
}

static int
calibrate_vec(std::vector<int> &deltas, int frequency)
{
	for (auto delta : deltas) {
		frequency += delta;
	}

	return frequency;
}

static int
calibrate(std::istream &stream, int frequency)
{
	std::vector<int> deltas;
	while (!stream.eof()) {
		int delta = 0;
		stream >> delta;
		deltas.push_back(delta);
	}
	
	return calibrate_vec(deltas, frequency);
}

static int
calibrate_detecting_repeats_vec(std::vector<int> &deltas, int frequency)
{
	std::set<int> frequencies;
	frequencies.insert(frequency);

	while (true) {
		for (auto delta : deltas) {
			frequency += delta;
			if (frequencies.count(frequency) > 0) {
				return frequency;
			}

			frequencies.insert(frequency);
		}
	}

	return frequency;
}

static int
calibrate_detecting_repeats(std::istream &stream, int frequency)
{
	std::vector<int> deltas;

	while (!stream.eof()) {
		int delta = 0;
		stream >> delta;
		deltas.push_back(delta);
	}

	return calibrate_detecting_repeats_vec(deltas, frequency);
}

static int
self_check_calibrate_string(std::string input)
{
	std::stringstream stream(input);

	auto out = calibrate(stream, 0);
	return out;
}

static int
self_check_calibrate_detect_repeats_string(std::string input)
{
	std::stringstream stream(input);

	auto out = calibrate_detecting_repeats(stream, 0);
	return out;
}

static void
assert(std::string msg, int expected, int actual)
{
	if (expected == actual) {
		return;
	}

	std::cerr << msg << ": expected " << expected << " but have " 
		  << actual << std::endl;
	std::abort();
}

static void self_check_calibrate(void)
{
	auto test_case_1 = "+1\n-2\n+3\n+1";
	auto test_case_2 = "+1\n+1\n+1\n";
	auto test_case_3 = "+1\n+1\n-2";
	auto test_case_4 = "-1\n-2\n-3";
	assert("test 1", 3, self_check_calibrate_string(test_case_1));
	assert("test 2", 3, self_check_calibrate_string(test_case_2));
	assert("test 3", 0, self_check_calibrate_string(test_case_3));
	assert("test 4", -6, self_check_calibrate_string(test_case_4));
	std::cout << "calibration self check: OK\n";
}

static void self_check_calibrate_detect_repeating(void)
{
	auto test_case_1 = "+1\n-2\n+3\n+1\n+1\n-2\n";
	auto test_case_2 = "+1\n-1\n";
	auto test_case_3 = "+3\n+3\n+4\n-2\n-4";
	auto test_case_4 = "-6\n+3\n+8\n+5\n-6";
	auto test_case_5 = "+7\n+7\n-2\n-7\n-4";
	assert("test 1", 2, self_check_calibrate_detect_repeats_string(test_case_1));
	assert("test 2", 0, self_check_calibrate_detect_repeats_string(test_case_2));
	assert("test 3", 10, self_check_calibrate_detect_repeats_string(test_case_3));
	assert("test 4", 5, self_check_calibrate_detect_repeats_string(test_case_4));
	assert("test 5", 14, self_check_calibrate_detect_repeats_string(test_case_5));
	std::cout << "calibration detecting repeats self check: OK\n";
}

static void self_check(void)
{
	self_check_calibrate();
	self_check_calibrate_detect_repeating();
}

int main(void)
{
	self_check();
	
	auto deltas = read_deltas(std::cin);
	auto out = calibrate_vec(deltas, 0);
	auto repeated = calibrate_detecting_repeats_vec(deltas, 0);
	std::cout << "calibration: " << out << std::endl;
	std::cout << "detected repeated frequency: " << repeated << std::endl;
}