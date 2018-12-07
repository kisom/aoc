#!/usr/bin/env python

import copy
import datetime
import re
import sys
from time import sleep

# { 'a': set('b', 'c') }

STEP_REGEX = re.compile(r"Step (.) must be finished before step (.) can begin.")


def parse_steps(lines):
    steps = {}
    for line in lines:
        m = STEP_REGEX.match(line)
        assert m
        requires, name = m.groups()
        if not name in steps:
            steps[name] = set()
        if not requires in steps:
            steps[requires] = set()
        steps[name].add(requires)

    return steps

def parse_file(path):
    with open(path, 'rt') as file:
        return parse_steps(file)

TEST_STEPS = {
    "A": set(["C"]),
    "B": set(["A"]),
    "C": set(),
    "D": set(["A"]),
    "E": set(["B", "D", "F"]),
    "F": set(["C"]),
}


def self_check_parse_steps():
    lines = """Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin.""".split(
        "\n"
    )
    steps = parse_steps(lines)
    assert steps == TEST_STEPS


def mark_done(steps, name):
    for step in steps.keys():
        if name in steps[step]:
            steps[step].remove(name)
    return steps

def next_task(steps):
    for step in sorted(steps.keys()):
        if len(steps[step]) == 0:
            return step

def sequence(steps):
    result = ""
    while len(steps) > 0:
        step = next_task(steps)
        result += step
        del(steps[step])
        steps = mark_done(steps, step)
    return result


def self_check_sequence():
    steps = copy.deepcopy(TEST_STEPS)
    result = sequence(steps)
    assert result == "CABDFE"

def time_for(name, offset=60):
    return offset + ord(name) - 0x40

def self_check_time_for():
    assert(61 == time_for('A'))
    assert(86 == time_for('Z'))

def find_available(waiting, time):
    return [worker for worker in  waiting if waiting[worker] <= time]

def next_task_sequenced(steps, assignments):
    in_progress = list(assignments.values())
    for step in sorted(steps.keys()):
        if len(steps[step]) == 0 and not step in in_progress:
            return step

def self_check_next_task_sequenced():
    steps = copy.deepcopy(TEST_STEPS)
    assignments = {0: 'C'}
    assert(next_task_sequenced(steps, assignments) == None)

def work_available(steps, assignment, waiting, time):
    available = find_available(waiting, time)
    if not available:
        return False
    steps_copy = copy.deepcopy(steps)
    assignment_copy = copy.deepcopy(assignment)
    for worker in available:
        if worker in assignment_copy:
            step_copy = mark_done(steps_copy, assignment_copy[worker])
            del(assignment_copy[worker])
    next_step = next_task_sequenced(steps_copy, assignment_copy)
    if next_step:
        return True
    return False
    
def sequence_timed(steps, workers, offset):
    time = 0
    waiting = dict.fromkeys(range(workers), 0)
    assignment = {}
    result = ''

    while len(steps) > 0:
        while work_available(steps, assignment, waiting, time):
            for worker in find_available(waiting, time):
                if worker in assignment:
                    step = assignment[worker]
                    result += step
                    steps = mark_done(steps, step)
                    del(assignment[worker])
                    del(steps[step])                
                step = next_task_sequenced(steps, assignment)
                if step:
                    waiting[worker] = time + time_for(step, offset)
                    assignment[worker] = step
                if len(steps) == 0:
                    break
        #s = '{:04}|\t'.format(time)
        #for worker in range(workers):
        #    if worker in assignment:
        #        s += assignment[worker]
        #    else:
        #        s += '.'
        #s += '\t| ' + result
        # print(s)
        time += 1        
    return time-1, result


def self_check_sequence_timed():
    steps = copy.deepcopy(TEST_STEPS)
    time, result = sequence_timed(steps, 2, 0)
    assert(result == "CABFDE")
    print(time)
    assert(time == 15)

def self_check():
    self_check_parse_steps()
    self_check_sequence()
    self_check_time_for()
    self_check_next_task_sequenced()    
    self_check_sequence_timed()

    print('self check OK')
    
def main(args):
    for path in args:
        steps = parse_file(path)
        print(sequence(copy.deepcopy(steps)))
        #time, result = sequence_timed(steps, 5, 60)
        #print('{} in {}'.format(result, time))
                            

if __name__ == "__main__":
    starts = datetime.datetime.now()
    self_check()
    main(sys.argv[1:])
    finishes = datetime.datetime.now()
    print('Finished in {}us'.format((finishes - starts).microseconds))
