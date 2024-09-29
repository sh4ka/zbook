package operations

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/zizdlp/zbook/util"
)

// Clone clones a git repository from the specified URL into the specified directory.
// Optionally, a specific branch can be cloned if provided.
func Clone(gitURL, dir, branch string) error {
	// Create the git clone command with branch and directory parameter
	args := []string{"clone", gitURL, dir}
	if branch != "" {
		args = append(args, "--branch", branch)
	}

	cmd := exec.Command("git", args...)

	// Run the command and capture its output
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf(util.ParserGitCloneError(string(output)))
	}
	return nil
}

// CloneWithPassword clones a git repository from the specified URL into the specified directory.
// It supports cloning private repositories using either a personal access token (token)
// or basic authentication (username and password). Optionally, a specific branch can be cloned if provided.
func CloneWithPassword(gitURL, dir, username, password, branch string) error {
	// Construct the clone URL with username and password embedded
	urlWithCredentials := embedCredentialsInURL(gitURL, username, password)

	// Create the git clone command with branch and directory parameter
	args := []string{"clone", urlWithCredentials, dir}
	if branch != "" {
		args = append(args, "--branch", branch)
	}

	cmd := exec.Command("git", args...)

	// Run the command and capture its output
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf(util.ParserGitCloneError(string(output)))
	}

	return nil
}

// CloneWithToken clones a git repository from the specified URL into the specified directory.
// It supports cloning private repositories using a personal access token. Optionally, a specific branch can be cloned if provided.
func CloneWithToken(gitURL, dir, token, branch string) error {
	// Construct the clone URL with the token embedded
	urlWithToken := embedTokenInURL(gitURL, token)

	// Create the git clone command with branch and directory parameter
	args := []string{"clone", urlWithToken, dir}
	if branch != "" {
		args = append(args, "--branch", branch)
	}

	cmd := exec.Command("git", args...)

	// Run the command and capture its output
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf(util.ParserGitCloneError(string(output)))
	}

	return nil
}

// embedCredentialsInURL embeds the username and password into the git URL for basic authentication.
func embedCredentialsInURL(gitURL, username, password string) string {
	// Split the URL at "//" to insert username and password
	parts := strings.Split(gitURL, "//")
	if len(parts) < 2 {
		return gitURL // Invalid URL format, return as is
	}

	// Insert username and password after "//"
	return fmt.Sprintf("%s//%s:%s@%s", parts[0], username, password, parts[1])
}

// embedTokenInURL embeds the personal access token into the git URL for token authentication.
func embedTokenInURL(gitURL, token string) string {
	// Split the URL at "//" to insert the token
	parts := strings.Split(gitURL, "//")
	if len(parts) < 2 {
		return gitURL // Invalid URL format, return as is
	}

	// Insert token after "//" using the format "token@"
	return fmt.Sprintf("%s//%s@%s", parts[0], token, parts[1])
}
