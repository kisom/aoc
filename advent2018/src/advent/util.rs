use std::fs::File;
use std::io::BufRead;
use std::io::BufReader;

pub fn read_int_sequence(filename: String) -> Vec<i64> {
	let mut deltas = Vec::new();
	let file = File::open(filename).expect("file not found");
	let reader = BufReader::new(&file);

	for line in reader.lines() {
		let delta: i64 = line.unwrap().parse().unwrap();
		deltas.push(delta);
	}

	return deltas;
}

pub fn read_string_sequence(filename: String) -> Vec<String> {
	let mut strings = Vec::new();
	let file = File::open(filename).expect("file not found");
	let reader = BufReader::new(&file);

	for line in reader.lines() {
		strings.push(line.unwrap());
	}

	return strings;
}
