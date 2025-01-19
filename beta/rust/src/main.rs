use std::env;
use std::process::exit;

mod build;
mod init;

fn main() {
    if env::args().len() != 2
        || (env::args().nth(1).unwrap() != "init" && env::args().nth(1).unwrap() != "build")
    {
        println!("Usage: ./main <command>");
        println!("Commands: init, build");
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
        _ => {
            println!("Invalid command. Use 'init' or 'build'.");
            exit(1);
        }
    }
}
