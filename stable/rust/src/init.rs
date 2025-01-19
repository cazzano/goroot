use std::env;
use std::fs;
use std::path::{Path, PathBuf};
use std::string::String;

pub fn handle_init() -> Result<(), String> {
    // Get current directory
    let current_dir =
        env::current_dir().map_err(|e| format!("error getting current directory: {}", e))?;
    println!("[DEBUG] Current directory: {}", current_dir.display());

    // Check if required files exist in current directory
    let mut has_go_file = false;
    let mut has_plain_file = false;

    let entries =
        fs::read_dir(&current_dir).map_err(|e| format!("error reading directory: {}", e))?;

    for entry in entries {
        let entry = entry.map_err(|e| format!("error reading entry: {}", e))?;
        let name = entry.file_name();
        let name_str = name.to_string_lossy();
        let base_name = name_str.trim_end_matches(".go").to_string();

        if name_str.ends_with(".go") {
            has_go_file = true;
            println!("[DEBUG] Found Go file: {}", name_str);
        }
        // Skip LICENSE, README.md, and .git
        if name_str != "LICENSE"
            && name_str != "README.md"
            && name_str != ".git"
            && name_str == base_name
        {
            has_plain_file = true;
            println!("[DEBUG] Found plain file: {}", name_str);
        }
    }

    // Determine target directory for creating folders
    let target_dir: PathBuf = if has_go_file || has_plain_file {
        current_dir
            .parent()
            .ok_or("No parent directory found")?
            .to_path_buf()
    } else {
        current_dir.clone()
    };

    if has_go_file || has_plain_file {
        let mut file_status = String::new();
        if has_go_file {
            file_status.push_str(".go file present");
        }
        if has_plain_file {
            if !file_status.is_empty() {
                file_status.push_str(", ");
            }
            file_status.push_str("plain file present");
        }
        println!("Found required file(s): {}", file_status);
    } else {
        println!("No required files found, creating directories in current location");
    }

    // Create src and target directories
    let dirs = ["src", "target"];
    for dir in &dirs {
        let dir_path = target_dir.join(dir);
        fs::create_dir_all(&dir_path)
            .map_err(|e| format!("error creating directory {}: {}", dir, e))?;
        println!("[DEBUG] Created directory: {}", dir_path.display());
    }

    Ok(())
}
