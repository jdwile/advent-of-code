use std::{fs, time::Instant};

pub fn read_input_file(day: i32) -> String {
    let file_path = format!("src\\inputs\\{}.txt", day);
    fs::read_to_string(file_path).expect("Should have been able to read the file")
}

pub fn start_timer() -> Instant {
    Instant::now()
}
