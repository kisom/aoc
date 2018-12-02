extern crate advent;

use std::env;

fn main() {
    let mut args: Vec<String> = env::args().collect();

    args.remove(0); // Remove the path to the binary.
    for path in args {
        println!("Processing calibrations: {}", path);
        let deltas = advent::util::read_int_sequence(path);
        let frequency = advent::calibrate::calibrate(&deltas, 0);
        let repeated = advent::calibrate::detect_repeated(&deltas, 0);

        println!("Calibrated frequency: {}", frequency);
        println!("Detected repeated frequency: {}", repeated);
    }
}
