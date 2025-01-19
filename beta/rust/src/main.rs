// src/main.rs

use std::env;
use std::process::exit;

mod build;
mod help;
mod init;
mod version; // Ensure this is included if you have it // Add this line to include the help module

fn main() {
    if env::args().len() != 2
        || (env::args().nth(1).unwrap() != "init"
            && env::args().nth(1).unwrap() != "build"
            && env::args().nth(1).unwrap() != "-v"
            && env::args().nth(1).unwrap() != "help")
    {
        help::display_help(); // Call the help function if the command is invalid
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
        "help" => {
            help::display_help(); // Call the help function
            exit(0);
        }
        _ => {
            help::display_help(); // Call the help function if the command is invalid
            exit(1);
        }
    }
}
