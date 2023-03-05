///////////////
// PORCUPINE //
///////////////

// todo: explore writing tree like data structures in rust
// todo: revise Map ds, k-v store theory (hashmaps, search trees etc )
// https://en.wikipedia.org/wiki/Hash_table
// todo: avl, red-black tree, B-Tree practice
// todo: explore Clap api

use clap::Parser;


#[derive(Debug, Parser)]
struct Cli {
    pattern: String,
    path: std::path::PathBuf
}

fn main() {
    let args = Cli::parse();

    println!("{:?}", args);
}
