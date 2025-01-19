pub fn debug_build(current_dir: &str, has_go_file: bool, has_plain_file: bool) {
    println!(
        "[DEBUG] Build process started in directory: {}",
        current_dir
    );

    if has_go_file {
        println!("[DEBUG] Go files are present in the directory.");
    } else {
        println!("[DEBUG] No Go files found in the directory.");
    }

    if has_plain_file {
        println!("[DEBUG] Plain files are present in the directory.");
    } else {
        println!("[DEBUG] No plain files found in the directory.");
    }
}

pub fn debug_build_success(binary_path: &str) {
    println!(
        "[DEBUG] Build successful! Binary created at: {}",
        binary_path
    );
}

pub fn debug_build_error(err: &dyn std::error::Error) {
    println!("[DEBUG] Build error: {}", err);
}
