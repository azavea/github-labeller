## GitHub Labeller

This is a CLI tool which allows you to manage GitHub labels consistently across many repositories.

Supported actions:
- Ensure a label exists on all configured repositories
- Delete a label, keyed by label text, on all configured repositories

## Getting Started

Go to the [releases page]() and download the zip for the most recent release. Extract it and copy the binary for your platform from the `./bin` directory to a location on your `$PATH`. Consider renaming the binary to remove the platform postfix.

Configure the application by creating a `.github-labeller` file in your HOME dir:
```
touch ~/.github-labeller
```

Copy the following template into `~/.github-labeller` and configure as appropriate:
```toml

token = "your github access token here"

[orgs]
# Multiple [org.foo] blocks can be provided under the [orgs] section
# Replace foo with the name of your GitHub org or account
[orgs.foo]
repositories = [
    "repo-one",
    "repo-two"
]

```

#### Getting your GitHub access token

You can create an API token for use with this app by logging into GitHub with your user and going to https://github.com/settings/tokens --> "Personal Access Token". Create a new token with the "repo" permission and give it a helpful name. Once you have your token, copy it into the `Token` field of the `.github-labeller` config file.

#### Running the tool

Once configured with your list of repositories and GitHub token, you're all set!

Examples:
```
// Create the specified label in all repos if it doesn't exist, or update the color if it does
github-labeller create "backlog" "#AEAEAE"

// Delete a label in all configured repos
github-labeller delete "todo"
```

## Developing

Checkout to `${GOPATH}/src/github.com/azavea/github-labeller`. Build by running `make`. After building, run with `./github-labeller`. Additional commands are available in the Makefile.

#### Building a release

Follow these steps when creating a new release:
- Create a new release branch via git flow
- Update the CHANGELOG.md
- Update the version string in `github-labeller.go`, following SEMVER
- Run `make release`
- Commit the changes and create a new git tag using the same version string
- Finish the release branch via git flow
- Push the develop and master branches, along with the new release tag
