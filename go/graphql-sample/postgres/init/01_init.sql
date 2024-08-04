CREATE USER newuser WITH PASSWORD 'password';

CREATE DATABASE graphql_sample;

\c graphql_sample;

-- Create db tables
CREATE TABLE users (
  id text NOT NULL,
  name text NOT NULL,
  project_v2 text,

  PRIMARY KEY (id)
);

CREATE TABLE repositories (
    id text NOT NULL,
    owner text NOT NULL,
    name text NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT current_timestamp,

    PRIMARY KEY (id),
    FOREIGN KEY (owner) REFERENCES users(id)
);

CREATE TABLE issues (
    id text NOT NULL,
    url text NOT NULL,
    title text NOT NULL,
    closed boolean NOT NULL DEFAULT FALSE,
    number integer NOT NULL check (number > 0),
    repository text NOT NULL,

    PRIMARY KEY (id),
    FOREIGN KEY (repository) REFERENCES repositories(id)
);

CREATE TABLE projects (
    id text NOT NULL,
    title text NOT NULL,
    url text NOT NULL,
    owner text NOT NULL,

    PRIMARY KEY (id),
    FOREIGN KEY (owner) REFERENCES users(id)
);

CREATE TABLE pullrequests (
    id text NOT NULL,
    base_ref_name text NOT NULL,
    closed boolean NOT NULL DEFAULT FALSE,
    head_ref_name text NOT NULL,
    url text NOT NULL,
    number integer NOT NULL check (number > 0),
    repository text NOT NULL,

    PRIMARY KEY (id),
    FOREIGN KEY (repository) REFERENCES repositories(id)
);

CREATE TABLE projectcards (
    id text NOT NULL,
    project text NOT NULL,
    issue text,
    pullrequest text,

    PRIMARY KEY (id),
    FOREIGN KEY (project) REFERENCES projects(id),
    FOREIGN KEY (issue) REFERENCES issues(id),
    FOREIGN KEY (pullrequest) REFERENCES pullrequests(id),

    check (
        issue IS NOT NULL OR pullrequest IS NOT NULL
    )
);
