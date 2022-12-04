use itertools::Itertools;
use std::{error::Error, fs, process::Command};

fn extract_microseconds(output: &str) -> Result<f64, Box<dyn Error>> {
    let out = output.lines().last().unwrap().trim();
    let time = if out.ends_with("ms") {
        out["Time: ".len()..out.len() - 2].parse::<f64>()? * 1000.0
    } else {
        out["Time: ".len()..out.len() - 3].parse::<f64>()?
    };
    Ok(time)
}

fn main() -> Result<(), Box<dyn Error>> {
    let days = fs::read_dir(concat!(env!("CARGO_MANIFEST_DIR"), "/src/bin/"))?
        .filter_map(|p| p.ok()?.path().file_stem()?.to_str().map(str::to_string))
        .sorted()
        .collect::<Vec<_>>();
    let mut total_time = 0.0;
    for day in &days {
        let cmd = Command::new("cargo")
            .args(["run", "--release", "--bin", day])
            .output()?;
        let output = String::from_utf8(cmd.stdout)?;
        println!("Day {}:\n{}", day, output);
        total_time += extract_microseconds(&output)?;
    }
    println!("Total time: {:.5?}ms", total_time / 1000.0);
    Ok(())
}
