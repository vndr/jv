# jv

JV Tool is a command-line tool written in Go that helps you fetch your local and public IP addresses. It is designed to be cross-platform and can be easily released and distributed using GoReleaser.

## Features

- Fetch local IP address for a specific network interface:
`jv ip local`
- Fetch public IP address using an external service: `jv ip public`
- Cross-platform support (Linux, macOS, Windows).
- GitHub Actions workflow to automate releases.
- Branch protection to ensure code quality.

```
$jv help

A CLI tool to fetch IP addresses, both local and public.

Usage:
  jv [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  ip          IP related commands

Flags:
  -h, --help   help for jv

Use "jv [command] --help" for more information about a command.
```


```
$jv ip help

IP related commands

Usage:
  jv ip [command]

Available Commands:
  local       Get the local IP address
  public      Get the public IP address

Flags:
  -h, --help   help for ip

Use "jv ip [command] --help" for more information about a command.
```
## How to contribute

### 1. Workflow for Making Changes
To make changes to the project, follow this workflow:

1. **Create a new branch:**
   ```bash
   git checkout -b feature/my-feature

2. **Make your changes and commit them:**
   ```bash
   git add .
   git commit -m "Implement my new feature"

3. **Push the branch to GitHub:**
   ```bash
   git push origin feature/my-feature

4. **Open a pull request** on GitHub from your branch into the **main** branch
5. **Review and merge** the pull request after it meets all the branch protection rules.

### 2. Set Up GoReleaser

1. **Install GoReleaser:**
   ```bash
   brew install goreleaser/tap/goreleaser

2. **Create a .goreleaser.yml Configuration File:**
    ```yml
    version: 2

    project_name: jv-ip-tool

    builds:
    - id: jv
        env:
        - CGO_ENABLED=0
        goos:
        - linux
        - darwin
        - windows
        goarch:
        - amd64
        - arm64
        binary: jv
        ldflags:
        - -s -w
        flags:
        - -trimpath
        artifacts:
        - type: binary
        - type: archive
            format: tar.gz
            name_template: "{{ .ProjectName }}_{{ .Version }}_{{ .Os }}_{{ .Arch }}"
            replacements:
            darwin: macOS
            amd64: x86_64
            arm64: ARM64
        - type: checksum
            name_template: "{{ .ProjectName }}_checksums.txt"

    release:
    github:
        owner: your-username
        name: go-ip-tool

3. **Create a GitHub Release:**
   ```bash
   git tag -a v1.0.0 -m "Initial release"
   git push origin v1.0.0

4. **Run GoReleaser:**
   ```bash
   goreleaser release


### 3. Configure GitHub Actions Workflow

1. **Create a Workflow File:**
   
   In your repository, create a file at .github/workflows/release.yml:
   ```yml
   name: Release

   on:
   push:
       tags:
       - "v*"

   jobs:
   release:
       runs-on: self-hosted

       steps:
       - uses: actions/checkout@v2
       - name: Set up Go
           uses: actions/setup-go@v2
           with:
           go-version: 1.18

       - name: Install Goreleaser
           run: |
           curl -sSfL https://goreleaser.com/install.sh | sh

       - name: Run Goreleaser
           run: goreleaser release
           env:
           GITHUB_TOKEN: ${{ secrets.RELEASE_TOKEN }}

2. **Set Up a GitHub Token:**

   - Create a GitHub personal access token with **repo** scope.
   - Add the token to your repository secrets under **Settings** -> **Secrets and variables** -> **Actions** -> **New repository secret**. 
   - Name it **RELEASE_TOKEN**.


### 4. Protect the `main` Branch

To ensure that the **main** branch is protected and all changes go through a pull request:

1. **Navigate to Branch Protection Rules**:
    - Go to **Settings** -> **Branches** -> **Add branch protection rule**.
2. **Configure Protection:**
    - Branch name pattern: **main**
    - Enable **Require a pull request before merging**.
    - (Optional) Enable **Require approvals** and set the number of required approvals.
    - (Optional) Enable **Require status checks to pass before merging**.
    - (Optional) Enable **Include administrators** to apply these rules to admins.

3. **Save the Protection Rule**

Click **Create** or **Save changes**.


## License

```
MIT License

Copyright (c) [year] [fullname]

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
