## GitHub Labeller

This is a CLI tool which allows you to manage GitHub labels consistently across many repositories.

Supported actions:
- Ensure a label exists on all configured repositories
- Delete a label, keyed by label text, on all configured repositories

## Getting Started

Download the binary for your platform by going to the releases page and place it somewhere on your `$PATH`.

Configure the application by copying the example .github-labeller config file to your home dir:
```
cp .github-labeller.example ~/.github-labeller
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

Checkout to `goworkspace/src/github.com/azavea/github-labeller`. Then use the standard `go` commands to build and install the tool to your go bin dir.
