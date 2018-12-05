#!/usr/bin/env python3

import collections
import datetime
import pdb
import re
import sys

date_regex = re.compile(r'^\[(\d+)-0?(\d+)-0?(\d+) 0?(\d+):0?(\d+)]')
guard_regex = re.compile(r'.+Guard #(\d+) begins shift$')

def parse_date(s):
    m = date_regex.match(s)
    assert(m)

    # return Date(m.group(1), m.group(2), m.group(3), m.group(4), m.group(5))
    return datetime.datetime(int(m.group(1)), int(m.group(2)), int(m.group(3)), int(m.group(4)), int(m.group(5), 0))

def self_check_dates():
    log = """[1518-04-21 00:57] wakes up
[1518-09-03 00:12] falls asleep
[1518-04-21 00:04] Guard #3331 begins shift""".split('\n')
    dates = [parse_date(s) for s in log]
    expected = [datetime.datetime(1518,4,21,0,57,0), datetime.datetime(1518,9,3,0,12,0), datetime.datetime(1518,4,21,0,4,0)]
    for i in range(3):
        assert(dates[i] == expected[i])

    expected = [datetime.datetime(1518,4,21,0,4,0), datetime.datetime(1518,4,21,0,57,0), datetime.datetime(1518,9,3,0,12,0)]
    assert(expected[2] > expected[1])
    dates = sorted(dates)
    for i in range(3):
        assert(dates[i] == expected[i])

EVENT_START = 1
EVENT_SLEEP = 2
EVENT_AWAKE = 3

class Event:
    def __init__(self, s):
        s = s.strip()
        self.date = parse_date(s)
        m = guard_regex.match(s)
        if m:
            self.id = int(m.group(1))
            self.what = EVENT_START
        elif s.endswith('falls asleep'):
            self.what = EVENT_SLEEP
        elif s.endswith('wakes up'):
            self.what = EVENT_AWAKE
        else:
            raise ValueError("invalid event: " + s)
        
    def __str__(self):
        s = '[{}] {}'.format(self.date, self.id)
        if self.what == EVENT_AWAKE:
            s += ' wakes up'
        elif self.what == EVENT_SLEEP:
            s += ' falls asleep'
        elif self.what == EVENT_START:
            s += ' starts shift'
        return s
    
    def __lt__(self, other):
        return self.date < other.date
    
def parse_log(log):
    events = []
    for line in log.split('\n'):
        line = line.strip()
        if not line:
            continue
        events.append(Event(line))
    
    events.sort()
    last_id = 0
    for i in range(len(events)):
        if events[i].what == EVENT_START:
            last_id = events[i].id
        else:
            events[i].id = last_id
    
    return events

def self_check_parse_log():
    log = """[1518-04-21 00:57] wakes up
[1518-09-03 00:12] falls asleep
[1518-04-21 00:04] Guard #3331 begins shift"""
    events = parse_log(log)
    expected = [EVENT_START, EVENT_AWAKE, EVENT_SLEEP]
    for i in range(3):
        assert(events[i].what == expected[i])
        
def sleepers(events):
    total_sleep = {}
    minutes = {}

    last_minute = 0
    for event in events:
        if event.what == EVENT_START:
            continue
        elif event.what == EVENT_SLEEP:
            last_minute = event.date.minute
        elif event.what == EVENT_AWAKE:
            slept = event.date.minute - last_minute
            if not event.id in total_sleep:
                total_sleep[event.id] = 0
            total_sleep[event.id] += slept
            if not event.id in minutes:
                minutes[event.id] = collections.Counter()
            for minute in range(last_minute, event.date.minute):
                minutes[event.id][minute] += 1
    
    return (total_sleep, minutes)

def find_part1(total_sleep, minutes):
    print(total_sleep)
    max_sleep = 0
    max_guard = 0
    for guard in total_sleep:
        if total_sleep[guard] > max_sleep:
            max_guard = guard
            max_sleep = total_sleep[guard]

    max_sleep = 0
    max_minute = 0
    for minute in minutes[max_guard]:
        if minutes[max_guard][minute] > max_sleep:
            max_sleep = minutes[max_guard][minute]
            max_minute = minute
    print(max_guard, max_minute)
    return max_guard * max_minute

def find_part2(total_sleep, minutes):
    max_guard = 0
    max_minute = 0
    max_sleep = 0

    for guard in minutes:
        for i in range(0, 60):
            if minutes[guard][i] > max_sleep:
                max_sleep = minutes[guard][i]
                max_guard = guard
                max_minute = i
    
    return max_guard * max_minute
                

def self_check():
    self_check_dates()
    self_check_parse_log()

def main(args):
    for arg in args:
        with open(arg, 'rt') as log:
            events = parse_log(log.read())
        total, minutes = sleepers(events)
        answer1 = find_part1(total, minutes)
        answer2 = find_part2(total, minutes)
        print(arg, answer1, answer2)

if __name__ == '__main__':
    self_check()
    main(sys.argv[1:])