#include <algorithm>
#include <cassert>
#include <fstream>
#include <iostream>
#include <vector>
#include <re2/re2.h>

using namespace std;

static re2::RE2	wakesup("^\\[(\\d+)-(\\d+)-(\\d+) (\\d+):(\\d+)] wakes up$");
static re2::RE2	fallasleep("^\\[(\\d+)-(\\d+)-(\\d+) (\\d+):(\\d+)] falls asleep$");
static re2::RE2	onshift("^\\[(\\d+)-(\\d+)-(\\d+) (\\d+):(\\d+)] Guard #(\\d+) begins shift");

class Date {
public:
	int	year;
	int	month;
	int	day;
	int	hour;
	int	minute;

	bool is_zero() { return year == 0 && month == 0 && day == 0 && hour == 0 && minute == 0; }

	Date() : year(0), month(0), day(0), hour(0), minute(0) {};
	Date(int year, int month, int day, int hour, int minute) : year(year), month(month), day(day), hour(hour), minute(minute) {};
	friend ostream& operator<<(ostream &outs, const Date& d) {
		outs << d.year << "-" << d.month << "-" << d.day << " " << d.hour << ":" << d.minute;
		return outs;
	}
	friend bool operator<(const Date& lhs, const Date &rhs) {
		if (lhs.year < rhs.year) return true;
		if (lhs.month < rhs.month) return true;
		if (lhs.day < rhs.day) return true;
		if (lhs.hour < rhs.hour) return true;
		if (lhs.minute < rhs.minute) return true;
		return false;
	}
	friend bool operator==(const Date& lhs, const Date &rhs) {
		if (lhs.year != rhs.year) return false;
		if (lhs.month != rhs.month) return false;
		if (lhs.day != rhs.day) return false;
		if (lhs.hour != rhs.hour) return false;
		if (lhs.minute != rhs.minute) return false;
		return true;
	}
	friend bool operator!=(const Date& lhs, const Date& rhs) {
		return (!(lhs == rhs));
	}
};

static void
self_check_dates()
{
	// 1518-11-1 0:0 10 3
	// 1518-11-4 0:2 99 3
	// 1518-11-5 0:3 99 3
	Date a = Date(1518,11,1,0,0);
	Date b = Date(1518,11,4,0,2);
	Date c = Date(1518,11,5,0,3);
	Date d = Date(1518, 11, 1, 0, 5);
	assert(d < c);
	assert(a < b);
	assert(a < c);
	assert(b < c);
	assert(d < b);
	assert(a < d);
}

class Event {
public:
	int	id;
	Date	d;
	int	what;
	Event() : id(0), d(), what(0) {};
	friend bool operator<(const Event& l, const Event& r) {
		return l.d < r.d;
	}
	friend bool operator==(const Event& a, const Event& b) {
		return a.id == b.id && a.d == b.d && a.what == b.what;
	}
	friend bool operator!=(const Event& a, const Event& b) {
		return !(a == b);
	}
};

static int ev_start = 3;
static int ev_sleep = 1;
static int ev_wake  = 2;

static void
dumpev(Event &e)
{
	cerr << e.d.year << "-" << e.d.month << "-" << e.d.day << " " << e.d.hour << ":" << e.d.minute << " " << e.what << " " << e.id << endl;
}


static Event
parse(string record)
{
	Event	e;

	int year, month, day, hour, minute, id;
	id = 0;

	if (re2::RE2::FullMatch(record, wakesup, &year, &month, &day, &hour, &minute)) {
		e.what = ev_wake;
	}
	else if (re2::RE2::FullMatch(record, fallasleep, &year, &month, &day, &hour, &minute)) {
		e.what = ev_sleep;
	}
	else if (re2::RE2::FullMatch(record, onshift, &year, &month, &day, &hour, &minute, &id)) { 
		e.what = ev_start;
	}
	else {
		abort();
	}

	Date d(year, month, day, hour, minute);
	e.id = id;
	e.d = d;
	assert(e.d.year == 1518);
	assert(e.what != 0);
	return e;
}

static int
read_events(string path)
{
	vector<Event>	records;
	ifstream	eventLog(path);
	map<int, int*>	sleepers;
	map<int, int>   max_sleepers;

	for (string record; getline(eventLog, record); ) {
		if (record == "") continue;
		auto event = parse(record);
		records.push_back(event);
	}

	for (auto event : records) {
		if (event.d.is_zero()) {
			cout << "found a zero date\n";
			abort();
		}
	}
	
	dumpev(records[0]);
	cout << records.size() << endl;
	sort(records.begin(), records.end(), [](Event a, Event b) { return a < b; });
	cout << "sorting complete\n";

	int last_id = 0;
	int last_minute = 0;

	int max_id = 0;
	int max_sleep = 0;

	for (auto i = 0; i < records.size(); i++) {
		auto record = records[i];
		cout << record.d << " " << record.id << " " << record.what << endl;
		if (record.what == ev_start) {
			last_id = record.id;
			continue;
		}
		if (record.what == ev_sleep) {
			last_minute = record.d.minute;
			continue;
		}
		if (record.what == ev_wake) {
			auto it = sleepers.find(last_id);

			if (it == sleepers.end()) {
				sleepers[last_id] = new int[60];
				for (int i = 0; i < 60; i++) {
					sleepers[last_id][i] = 0;
				}
				max_sleepers[last_id] = 0;
			}

			cout << last_id << " " << last_minute << " " << record.d.minute << endl;
			for (int i = last_minute; i < record.d.minute; i++) {
				sleepers[last_id][i] = sleepers[last_id][i] + 1;
			}

			max_sleepers[last_id] += (record.d.minute - last_minute);
			if (max_sleepers[last_id] > max_sleep) {
				max_id = last_id;
				max_sleep = max_sleepers[last_id];
				cout << "max sleeper: " << max_id << " " << max_sleep << endl;
			}
		}

	}

	last_minute = 0;
	max_sleep = 0;
	auto schedule = sleepers[max_id];
	for (auto i = 0; i < 60; i++) {
		if (schedule[i] > max_sleep) {
			max_sleep = schedule[i];
			last_minute = i;
			cout << "new best: " << max_id << " " << max_sleep << " " << last_minute << endl;
		}
	}

	cout << ">" << max_id << " " << last_minute << endl;
	return max_id * last_minute;
}

int
main(int argc, char *argv[])
{
	self_check_dates();
	for (auto i = 1; i < argc; i++) {
		auto answer = read_events(string(argv[i]));
		cout << answer << endl;
	}
}
