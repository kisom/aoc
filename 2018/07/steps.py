#!/usr/bin/env python

import copy
import re
import sys

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


def sequence(steps):
    result = ""
    while len(steps) > 0:
        for step in sorted(steps.keys()):
            if len(steps[step]) == 0:
                result += step
                del(steps[step])
                steps = mark_done(steps, step)
                break
    return result


def self_check_sequence():
    steps = copy.deepcopy(TEST_STEPS)
    result = sequence(steps)
    assert result == "CABDFE"


def self_check():
    self_check_parse_steps()
    self_check_sequence()

    print('self check OK')

def sequence_file(path):
    with open(path, 'rt') as file:
        return sequence(parse_steps(file))
    
def main(args):
    for path in args:
        print(sequence_file(path))

if __name__ == "__main__":
    self_check()
    main(sys.argv[1:])
