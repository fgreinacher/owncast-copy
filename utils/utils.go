package utils

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/mssola/user_agent"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
	"mvdan.cc/xurls"
)

//GetTemporaryPipePath gets the temporary path for the streampipe.flv file
func GetTemporaryPipePath() string {
	return filepath.Join(os.TempDir(), "streampipe.flv")
}

//DoesFileExists checks if the file exists
func DoesFileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}

	return true
}

//GetRelativePathFromAbsolutePath gets the relative path from the provided absolute path
func GetRelativePathFromAbsolutePath(path string) string {
	pathComponents := strings.Split(path, "/")
	variant := pathComponents[len(pathComponents)-2]
	file := pathComponents[len(pathComponents)-1]

	return filepath.Join(variant, file)
}

//Copy copies the
func Copy(source, destination string) error {
	input, err := ioutil.ReadFile(source)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(destination, input, 0644)
}

// IsUserAgentABot returns if a web client user-agent is seen as a bot
func IsUserAgentABot(userAgent string) bool {
	if userAgent == "" {
		return false
	}

	botStrings := []string{
		"mastodon",
		"pleroma",
		"applebot",
	}

	for _, botString := range botStrings {
		if strings.Contains(strings.ToLower(userAgent), botString) {
			return true
		}
	}

	ua := user_agent.New(userAgent)
	return ua.Bot()
}

func RenderSimpleMarkdown(raw string) string {
	markdown := goldmark.New(
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
		goldmark.WithExtensions(
			extension.NewLinkify(
				extension.WithLinkifyAllowedProtocols([][]byte{
					[]byte("http:"),
					[]byte("https:"),
				}),
				extension.WithLinkifyURLRegexp(
					xurls.Strict,
				),
			),
		),
	)

	trimmed := strings.TrimSpace(raw)
	var buf bytes.Buffer
	if err := markdown.Convert([]byte(trimmed), &buf); err != nil {
		panic(err)
	}

	return buf.String()
}
