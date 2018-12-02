fn compute_checksum(box_id: String) -> (i64, i64) {
	let mut set: [u8; 26] = [
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	];
	let mut two: i64 = 0;
	let mut three: i64 = 0;

	for c in box_id.chars() {
		let i = c as usize;
		set[i - 0x61] += 1;
	}

	for i in 0..set.len() {
		if set[(i as usize)] == 2 {
			two += 1;
		}
		if set[i as usize] == 3 {
			three += 1;
		}
	}

	if two > 0 {
		two = 1;
	}

	if three > 0 {
		three = 1;
	}

	return (two, three);
}

#[test]
fn test_compute_checksum() {
	let (two, three) = compute_checksum("abcdef".to_string());
	assert!(two == 0 && three == 0);

	let (two, three) = compute_checksum("bababc".to_string());
	assert!(two == 1 && three == 1);
}

pub fn compute_checksums(box_ids: &Vec<String>) -> i64 {
	let mut two: i64 = 0;
	let mut three: i64 = 0;

	for box_id in box_ids.iter() {
		let (two_, three_) = compute_checksum(box_id.to_string());
		two += two_;
		three += three_;
	}

	return two * three;
}

#[test]
fn test_compute_checksums() {
	let box_ids: Vec<String> = vec![
		"abcdef".to_string(),
		"bababc".to_string(),
		"abbcde".to_string(),
		"abcccd".to_string(),
		"aabcdd".to_string(),
		"abcdee".to_string(),
		"ababab".to_string(),
	];
	let checksum = compute_checksums(&box_ids);
	assert!(checksum == 12);
}

fn match_boxes(box_id1: String, box_id2: String) -> Option<String> {
	let mut differences = 0;
	let mut common: String = String::new();

	let box_id1_chars = box_id1.as_bytes();
	let box_id2_chars = box_id2.as_bytes();

	for i in 0..box_id1.len() {
		if box_id1_chars[i] == box_id2_chars[i] {
			common.push(box_id1_chars[i] as char);
			continue;
		}

		differences += 1;
		if differences > 1 {
			break;
		}
	}

	if differences == 1 {
		return Some(common);
	}

	None
}

#[test]
fn test_match_boxes() {
	assert!(
		match_boxes("abcde".to_string(), "fghij".to_string()).is_none(),
		"boxes shouldn't match"
	);
	assert!(
		match_boxes("abcde".to_string(), "axcye".to_string()).is_none(),
		"boxes shouldn't match"
	);
	let common = match_boxes("fghij".to_string(), "fguij".to_string());
	match common {
		Some(s) => {
			println!("common: {}", s);
			assert!(s == "fgij".to_string());
		}
		_ => assert!(false),
	};
}

pub fn match_inventory(inventory: &Vec<String>) -> Option<String> {
	for i in 0..(inventory.len() - 1) {
		for j in i + 1..inventory.len() {
			let result = match_boxes(inventory[i].clone(), inventory[j].clone());
			match result {
				Some(common) => return Some(common),
				_ => continue,
			}
		}
	}

	None
}

#[test]
fn test_match_inventory() {
	let inventory = &vec![
		"abcde".to_string(),
		"fghij".to_string(),
		"klmno".to_string(),
		"pqrst".to_string(),
		"fguij".to_string(),
		"axcye".to_string(),
		"wvxyz".to_string(),
	];
	let result = match_inventory(&inventory);
	match result {
		Some(common) => assert!(common == "fgij".to_string()),
		_ => assert!(false),
	};
}
