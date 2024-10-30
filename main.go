//

package main

import (
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/go-git/go-git/v5"
)

type URL string

func (u URL) rebuild() (string, error) {
	parsedUrl, err := url.Parse(string(u))
	if err != nil {
		return "", fmt.Errorf("URLの解析中にエラーが発生しました: %s", err)
	}

	// 再構築したURLを表示
	return fmt.Sprintf(
		"https://github.com%s",
		strings.TrimSuffix(parsedUrl.Path, ".git"),
	), nil
}

var path string

func init() {
	path = "."

	args := os.Args
	if len(args) > 1 {
		path = args[1]
	}
}

func getUrlMap() (map[string]URL, error) {
	// リポジトリをオープン
	repo, err := git.PlainOpen(path)
	if err != nil {
		return nil, err
	}

	// 現在のブランチを取得
	remotes, err := repo.Remotes()
	if err != nil {
		return nil, err
	}

	urls := map[string]URL{}
	for _, r := range remotes {
		urls[r.Config().Name] = URL(r.Config().URLs[0])
	}

	return urls, nil
}

func main() {
	urls, err := getUrlMap()
	if err != nil {
		fmt.Println("エラー:", err)
		os.Exit(1)
	}

	u, ok := urls["origin"]
	if !ok {
		fmt.Println("エラー: 'origin' リモートが見つかりませんでした。")
		return
	}

	// 再構築したURLを表示
	url, err := u.rebuild()
	if err != nil {
		fmt.Println("エラー:", err)
		os.Exit(1)
	}
	openBrowser(url)
}

func openBrowser(url string) error {
	var err error
	switch {
	case runtime.GOOS == "linux":
		err = exec.Command("xdg-open", url).Start()
	case runtime.GOOS == "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case runtime.GOOS == "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	return err
}
