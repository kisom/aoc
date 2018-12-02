use std::collections::HashSet;

/// calibrate is used to calibrate the time stream.
pub fn calibrate(deltas: &Vec<i64>, frequency: i64) -> i64 {
    let mut calibrated = frequency;
    for delta in deltas {
        calibrated += delta;
    }
    return calibrated;
}

#[cfg(test)]
fn test_calibrate_case(deltas: &Vec<i64>, expected: i64) {
    let repeated = calibrate(deltas, 0);
    if repeated != expected {
        println!("expected {} but got {}", expected, repeated);
    }
    assert!(repeated == expected);
}

#[test]
fn test_calibrate() {
    test_calibrate_case(&vec![1, -2, 3, 1], 3);
    test_calibrate_case(&vec![1, 1, 1], 3);
    test_calibrate_case(&vec![1, 1, -2], 0);
    test_calibrate_case(&vec![-1, -2, -3], -6);
}

pub fn detect_repeated(deltas: &Vec<i64>, frequency: i64) -> i64 {
    let mut seen: HashSet<i64> = HashSet::new();
    let mut calibrated = frequency;

    seen.insert(calibrated);

    loop {
        for delta in deltas {
            calibrated += delta;
            if seen.contains(&calibrated) {
                return calibrated;
            }
            seen.insert(calibrated);
        }
    }
}

#[cfg(test)]
fn test_detect_repeated_case(deltas: &Vec<i64>, expected: i64) {
    let repeated = detect_repeated(deltas, 0);
    if repeated != expected {
        println!("expected {} but got {}", expected, repeated);
    }
    assert!(repeated == expected);
}

#[test]
fn test_detect_repeated() {
    test_detect_repeated_case(&vec![1, -1], 0);
    test_detect_repeated_case(&vec![1, -2, 3, 1], 2);
    test_detect_repeated_case(&vec![3, 3, 4, -2, -4], 10);
    test_detect_repeated_case(&vec![-6, 3, 8, 5, -6], 5);
    test_detect_repeated_case(&vec![7, 7, -2, -7, -4], 14);
}
