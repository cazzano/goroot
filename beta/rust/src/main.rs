// src/main.rs

use std::env;
use std::process::exit;

mod build;
mod init;
mod version; // Add this line to include the version module

fn main() {
    if env::args().len() != 2
        || (env::args().nth(1).unwrap() != "init"
            && env::args().nth(1).unwrap() != "build"
            && env::args().nth(1).unwrap() != "-v")
    {
        println!("Usage: ./main <command>");
        println!("Commands: init, build, -v");
        exit(1);
    }

    match env::args().nth(1).unwrap().as_str() {
        "init" => {
            if let Err(err) = init::handle_init() {
                eprintln!("Error: {}", err);
                exit(1);
            }
        }
        "build" => {
            if let Err(err) = build::handle_build() {
                eprintln!("Error: {}", err);
                exit(1);
            }
        }
        "-v" => {
            version::display_version(); // Call the version display function
            exit(0);
        }
        _ => {
            println!("Invalid command. Use 'init', 'build', or '-v'.");
            exit(1);
        }
    }
}
