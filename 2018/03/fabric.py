#!/usr/bin/env python3

import collections
import re
import sys
from typing import List


class Claim:
    __slots__ = ["id", "upper_left", "lower_right"]

    def __init__(self, id: int, x0: int, y0: int, x1: int, y1: int):
        self.id = int(id)
        self.upper_left = (x0, y0)
        self.lower_right = (x1, y1)

    def __eq__(self, other):
        if self.id != other.id:
            return False
        if self.upper_left != other.upper_left:
            return False
        if self.lower_right != other.lower_right:
            return False
        return True

    def points(self):
        points = []
        for j in range(self.upper_left[1], self.lower_right[1]):
            for i in range(self.upper_left[0], self.lower_right[0]):
                points.append((i, j))
        return points


def create_claim_regex():
    return re.compile(r"^#(\d+) @ (\d+),(\d+): (\d+)x(\d+)$")


claim_regex = create_claim_regex()


def process_claim(claim: str) -> Claim:
    m = claim_regex.match(claim.strip())
    x0 = int(m.group(2))
    y0 = int(m.group(3))
    x1 = x0 + int(m.group(4))
    y1 = y0 + int(m.group(5))
    return Claim(m.group(1), x0, y0, x1, y1)


def self_check_process_claim():
    claim_strings = ["#7 @ 922,250: 13x26", "#8 @ 256,742: 18x14"]
    expected = [Claim(7, 922, 250, 935, 276), Claim(8, 256, 742, 274, 756)]

    for i in range(len(claim_strings)):
        claim = process_claim(claim_strings[i])
        assert claim == expected[i]


def load_claims(path: str) -> List[Claim]:
    with open(path, "rt") as claim_file:
        return list(map(process_claim, claim_file.readlines()))


def map_points(claims):
    counter = collections.Counter()
    for claim in claims:
        points = claim.points()
        for point in points:
            counter[point] += 1

    return counter


def overlapping(points):
    return len([pt for pt in points.keys() if points[pt] > 1])


def alone(claims, counter):
    for claim in claims:
        non_overlapping = 0
        points = claim.points()
        for point in points:
            if counter[point] == 1:
                non_overlapping += 1
        if len(points) == non_overlapping:
            return claim.id
    return None


def self_check_overlapping():
    claims = list(
        map(process_claim, ["#1 @ 1,3: 4x4", "#2 @ 3,1: 4x4", "#3 @ 5,5: 2x2"])
    )
    area = overlapping(map_points(claims))
    assert area == 4


def self_check():
    self_check_process_claim()
    # self_check_points()
    self_check_overlapping
    print("self check OK")


def main(args):
    for arg in args:
        claims = load_claims(arg)
        print("Loaded {} claims".format(len(claims)))
        counter = map_points(claims)
        area = overlapping(counter)
        print("Overlapping area:", area)

        lone = alone(claims, counter)
        print("Lone claim:", lone)


if __name__ == "__main__":
    self_check()
    main(sys.argv[1:])
