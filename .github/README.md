# Furiosa
A Release Management Service and CLI for GitHub.com / JIRA Shops.

* [Read](https://github.com/MXWest/Furiosa/wiki) about the project.
* [Contribute](./CONTRIBUTING.md) to the project!
* [Respect](./CODE_OF_CONDUCT.md) our [Code of Conduct](./CODE_OF_CONDUCT.md)


## Environment Variables

### JIRA Access
Basic Authentication is used. Set the following:
* `FURIOSA_JIRA_USERNAME`
* `FURIOSA_JIRA_PASSWORD`
* `FURIOSA_JIRA_URL`

### GitHub Access
[Create a personal access token](https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/) and assign its value to:
* `FURIOSA_GITHUB_TOKEN`