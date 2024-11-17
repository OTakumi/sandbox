use assert_cmd::Command;
use predicates::prelude::*;

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
