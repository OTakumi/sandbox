use assert_cmd::Command;
use std::fs;

#[test]
fn runs() {
    let mut cmd = Command::cargo_bin("echo").unwrap();
    cmd.arg("hello").assert().success();
}

#[test]
#[should_panic]
fn dies_no_args() {
    let mut cmd = Command::cargo_bin("echo").unwrap();
    cmd.assert().stderr(predicates::str::contains("USAGE"));
}

#[test]
fn hello_1() {
    let outfile = "tests/expected/hello_1.txt";
    let expected = fs::read_to_string(outfile).unwrap();

    let mut cmd = Command::cargo_bin("echo").unwrap();
    cmd.arg("Hello there").assert().success().stdout(expected);
}
