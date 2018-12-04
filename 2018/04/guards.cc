#include <algorithm>
#include <fstream>
#include <iostream>
#include <vector>
#include <re2/re2.h>

using namespace std;

class Guard {
public:
	int	id;
	int	sleeps;
	Guard() : id(0), sleeps(0) {};
} Guard;

const int EVENT_ONSHIFT = 1;
const int EVENT_SLEEP   = 2;
const int EVENT_AWAKE   = 3;

class Event {
public:
	int		id;
	uint32_t	timestamp;
	int		event;
	int		sleep;

	Event() : id(0), timestamp(0), event(0), sleep(0) {};

	friend bool operator<(const Event& l, const Event& r) {
		return l.timestamp < r.timestamp;
	}
};



static uint32_t
get_timestamp(int year, int month, int day, int hour, int minute)
{
	uint32_t	timestamp;

	timestamp = year * 31536000;
	switch (month) {
		case 1:
		case 3:
		case 5:
		case 7:
		case 8:
		case 10:
		case 12:
			timestamp += 31 * 86400;
			break;
		case 2:
			timestamp += 28 * 86400;
			break;
		default:
			timestamp += 30 * 86400;
	}

	timestamp += day * 86400;
	timestamp += hour * 3600;
	timestamp += minute * 60;
	return timestamp;
}

static re2::RE2	wakesup("^[(\\d+)\\-(\\d+)\\-(\\d+) (\\d+):(\\d+)] wakes up$");
static re2::RE2	fallasleep("^[(\\d+)\\-(\\d+)\\-(\\d+) (\\d+):(\\d+)] falls asleep$");
static re2::RE2	onshift("[(\\d+)\\-(\\d+)\\-(\\d+) (\\d+):(\\d+)] Guard #(\\d+) begins shift");

static Event
parse_event(string record)
{
	int year, month, day, hour, minute;
	int id;
	Event event;

	if (re2::RE2::FullMatch(record, onshift, &year, &month, &day, &hour, &minute, &id)) {
		auto timestamp = get_timestamp(year, month, day, hour, minute);
		event.id = id;
		event.timestamp = timestamp;
		event.event = EVENT_ONSHIFT;
	}
	else if (re2::RE2::FullMatch(record, wakesup, &year, &month, &day, &hour, &minute)) {
		auto timestamp = get_timestamp(year, month, day, hour, minute);
		event.timestamp = timestamp;
		event.event = EVENT_AWAKE;
	}
	else if (re2::RE2::FullMatch(record, fallasleep, &year, &month, &day, &hour, &minute)) {
		auto timestamp = get_timestamp(year, month, day, hour, minute);
		event.timestamp = timestamp;
		event.event = EVENT_SLEEP;
	}

	return event;
}

static vector<Event>
read_events(string path)
{
	vector<Event>	events;
	ifstream	eventLog(path);

	for (string record; getline(eventLog, record); ) {
		auto event = parse_event(record);
		events.push_back(event);
	}
	
	sort(events.begin(), events.end());
	int id = 0;
	uint32_t last_time = 0;

	for (auto i = 0; i < events.size(); i++) {
		auto event = events.at(i);
		if (event.event == EVENT_ONSHIFT) {
			id = event.id;
			continue;
		}

		event.id = id;
		if (event.event == EVENT_SLEEP) {
			last_time = event.timestamp;
		} else {
			event.sleep = event.timestamp - last_time;
		}

		events.at(i) = event;
		cout << event.id << " " << event.event << " @ " << event.timestamp << endl;
	}

	return events;
}

static int
find_sleep_guard(vector<Event> events)
{
	map<int, int>	sleepers;

	int id = 0;
	int max_sleep = 0;

	for (auto event : events) {
		if (event.event != EVENT_AWAKE) {
			continue;
		}

		auto it = sleepers.find(event.id);
		auto sleep = event.sleep;

		if (it == sleepers.end()) {
			sleepers[event.id] = sleep;
		}
		else {
			sleep += it->second;
			sleepers[event.id] = sleep;
		}

		if (sleep > max_sleep) {
			id = event.id;
			max_sleep = sleep;
		}
	}

	return id * (max_sleep / 60);
}

int
main(int argc, char *argv[])
{
	for (auto i = 1; i < argc; i++) {
		auto events = read_events(string(argv[i]));
		cout << events.size() << " events loaded.\n";
		
		auto answer = find_sleep_guard(events);
		cout << answer << endl;
	}
}