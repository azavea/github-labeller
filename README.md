## GitHub Labeller

**This tool is still WIP and does not have all advertised functionality**

This is a CLI tool which allows you to manage GitHub labels consistently across many repositories.

Supported actions:
- Create a new label on all configured repositories
- Update an existing label on all configured repositories
- Delete a label, keyed by label text, on all configured repositories

## Getting Started

Download the binary for your platform by going to the releases page and place it somewhere on your `$PATH`.

Configure the application by copying the example .github-labeller config file to your home dir:
```
cp .github-labeller.example ~/.github-labeller
```

#### Getting your GitHub access token

You can create an API token for use with this app by logging into GitHub with your user and going to https://github.com/settings/tokens --> "Personal Access Token". Create a new token with the "repo" permission and give it a helpful name. Once you have your token, copy it into the `Token` filed of the `.github-labeller` config file.

#### Running the tool

Once configured with your list of repositories and GitHub token, you're all set!

Examples:
```
// Create a new label in all repos and update an repos with the existing label name
github-labeller -create "backlog" "#AEAEAE"

// Delete a label in all repos
github-labeller -delete "todo" "#000"
```

## Developing

Checkout to `goworkspace/src/github.com/azavea/github-labeller`. Then use the standard `go` commands to build and install the tool to your go bin dir.
