#!/usr/bin/env python3

import io
import operator
import sys

def react(stream):
    last = stream.read(1)
    result = '' + last
    current = stream.read(1)

    while len(current) > 0:
        if result != '':
            if last.upper() == current.upper() and last != current:
                result = result[:-1]
                if result != '':
                    last = result[-1]
            else:
                last = current
                result += last
        else:
            last = current
            result += last
        current = stream.read(1)

    return result

def reactall(s):
    results = {}
    x = set(s.upper())
    for base in x:
        s2 = s.replace(base, '')
        s2 = s2.replace(base.lower(), '')
        result = react(io.StringIO(s2))
        results[base] = len(result)
    return min(results.items(), key=operator.itemgetter(1))

def test_string(s, expected):
    stream = io.StringIO(s)
    result = react(stream)
    assert(result == expected)

def self_check_react():
    strings = ['dabAcCaCBAcCcaDA']
    expected = ['dabCBAcaDA']
    for i in range(len(strings)):
        test_string(strings[i], expected[i])

def self_check_reactall():
    polymer = 'dabAcCaCBAcCcaDA'
    expected = ('C', 4)
    result = reactall(polymer)
    assert(result == expected)

def self_check():
    self_check_react()
    self_check_reactall()
    print('self check OK')

def main(args):
    for path in args:
        with open(path, 'rt') as chain:
            polymer = chain.read()
        result = react(io.StringIO(polymer))
        print('{}: {}'.format(path, len(result)))

        result = reactall(polymer)
        print('{}: shortest chain {} by removing {}'.format(path, result[1], result[0]))

if __name__ == '__main__':
    self_check()
    main(sys.argv[1:])