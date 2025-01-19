use std::env;
use std::fs;
use std::path::{Path, PathBuf};
use std::process::Command;

pub fn handle_build() -> Result<(), String> {
    // Get the current directory
    let current_dir =
        env::current_dir().map_err(|e| format!("error getting current directory: {}", e))?;

    // Check for Go files in the current directory
    let has_go_file = check_for_go_files(&current_dir)?;

    // Check for specific plain files in the current directory
    let plain_files = vec!["filename1", "filename2"]; // Replace with your actual filenames
    let has_plain_file = check_for_plain_files(&current_dir, &plain_files)?;

    // Call the debug function to print debug information
    debug_build(
        &current_dir.display().to_string(),
        has_go_file,
        has_plain_file,
    );

    // If no Go files in current directory and specific plain files are found, try src directory
    if !has_go_file && has_plain_file {
        // Change to the parent directory twice
        let parent_dir = current_dir.parent().ok_or("No parent directory found")?;
        let src_dir = parent_dir.join("src");

        // Recheck for Go files in src directory
        let has_go_file = check_for_go_files(&src_dir)?;
        if !has_go_file {
            return Err("no Go files found in current directory or src directory".to_string());
        }

        println!("[DEBUG] Switching to src directory for build");
        build_project(&src_dir)?;
    } else if has_go_file {
        // Build the project in the current directory
        build_project(&current_dir)?;
    } else {
        return Err("no Go files found to build".to_string());
    }

    Ok(())
}

fn check_for_go_files(dir: &Path) -> Result<bool, String> {
    let entries = fs::read_dir(dir).map_err(|e| format!("error reading directory: {}", e))?;
    for entry in entries {
        let entry = entry.map_err(|e| format!("error reading entry: {}", e))?;
        if entry.file_name().to_string_lossy().ends_with(".go") {
            println!(
                "[DEBUG] Found Go file in {}: {}",
                dir.display(),
                entry.file_name().to_string_lossy()
            );
            return Ok(true);
        }
    }
    Ok(false)
}

fn check_for_plain_files(dir: &Path, plain_files: &[&str]) -> Result<bool, String> {
    let entries = fs::read_dir(dir).map_err(|e| format!("error reading directory: {}", e))?;
    for entry in entries {
        let entry = entry.map_err(|e| format!("error reading entry: {}", e))?;
        let name = entry.file_name().to_string_lossy().to_string(); // Store in a longer-lived variable
        if plain_files.iter().any(|&file| name == file) {
            println!("[DEBUG] Found plain file in {}: {}", dir.display(), name);
            return Ok(true);
        }
    }
    Ok(false)
}

fn build_project(dir: &Path) -> Result<(), String> {
    let output = Command::new("go")
        .arg("build")
        .current_dir(dir)
        .output()
        .map_err(|e| format!("error building project: {}", e))?;

    if !output.status.success() {
        return Err(format!(
            "Build failed: {}",
            String::from_utf8_lossy(&output.stderr)
        ));
    }

    // Define the target release directory
    let release_dir = dir.join("../target/release");
    fs::create_dir_all(&release_dir)
        .map_err(|e| format!("error creating release directory: {}", e))?;

    // Move the compiled binary to the target/release directory
    let binary_name = dir.file_name().unwrap().to_string_lossy(); // Use the current directory name as the binary name
    let src_binary_path = dir.join(&*binary_name);
    let dest_binary_path = release_dir.join(&*binary_name);

    fs::rename(src_binary_path, dest_binary_path.clone())
        .map_err(|e| format!("error moving binary to release directory: {}", e))?;
    println!("Binary moved to: {}", dest_binary_path.display());

    Ok(())
}

fn debug_build(current_dir: &str, has_go_file: bool, has_plain_file: bool) {
    println!(
        "[DEBUG] Build process started in directory: {}",
        current_dir
    );
    if has_go_file {
        println!("[DEBUG] Go files are present in the directory.");
    }
    if has_plain_file {
        println!("[DEBUG] Plain files are present in the directory.");
    }
}
