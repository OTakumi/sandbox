use clap::Parser;

/// echor - Rust echo
/// version: 0.1.0
#[derive(Parser, Debug)]
#[command(version, about)]
struct Cli {
    /// Input text
    text: Vec<String>,

    /// Do not print newline
    #[arg(short = 'n')]
    omit_newline: bool,
}

fn main() {
    let cli = Cli::parse();

    println!("{:?}", std::env::args());

    // Check if omit_newline is true
    // If true, set ending to empty string
    let ending = if cli.omit_newline { "" } else { "\n" };

    print!("{}{}", cli.text.join(" "), ending);
}
