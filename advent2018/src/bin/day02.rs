extern crate advent;

use std::env;

fn main() {
    let mut args: Vec<String> = env::args().collect();

    args.remove(0); // Remove the path to the binary.
    for path in args {
        println!("Processing inventory: {}", path);
        let inventory = advent::util::read_string_sequence(path);
        let checksum = advent::inventory::compute_checksums(&inventory);

        println!("Inventory checksum: {}", checksum);

        let result = advent::inventory::match_inventory(&inventory);
        match result {
            Some(common) => println!("common: {}", common),
            None => println!("common: not found"),
        }
    }
}
