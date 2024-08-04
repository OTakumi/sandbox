\c graphql_sample;

-- Insert initial data
INSERT INTO users (id, name) VALUES
    ('u_1', 'Alice'),
    ('u_2', 'Bob'),
    ('u_3', 'Charlie');

INSERT INTO repositories (id, owner, name) VALUES
    ('r_1', 'u_1', 'REPO_1'),
    ('r_2', 'u_1', 'REPO_2'),
    ('r_3', 'u_2', 'REPO_3');

INSERT INTO issues (id, url, title, closed, number, repository) VALUES
    ('iss_1', 'http://example.com/repo1/issue/1', 'First Issue', False, 1, 'r_1'),
    ('iss_2', 'http://example.com/repo1/issue/2', 'Second Issue', False, 2, 'r_1'),
    ('iss_3', 'http://example.com/repo2/issue/1', 'First Issue', False, 3, 'r_2'),
    ('iss_4', 'http://example.com/repo3/issue/1', 'First Issue', False, 4, 'r_1');

INSERT INTO projects (id, title, url, owner) VALUES
    ('p_1', 'Project 1', 'http://example.com/project/1', 'u_1'),
    ('p_2', 'Project 2', 'http://example.com/project/2', 'u_2');

INSERT INTO pullrequests (id, base_ref_name, head_ref_name, url, number, repository) VALUES
    ('pr_1', 'master', 'feature', 'http://example.com/repo1/pr/1', 1, 'r_1'),
    ('pr_2', 'master', 'feature', 'http://example.com/repo1/pr/2', 2, 'r_2');
