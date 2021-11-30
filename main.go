package main

import (
	"context"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"strings"
	"time"
)

type Link struct {
	url string
	title string
}

// TODO: Write your function here
func CrawlLink(url string) []Link {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, time.Second * 15)
	defer cancel()

	var nodes []*cdp.Node
	if err := chromedp.Run(
		ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible(`.container a`),
		chromedp.Nodes(
			`a`,
			&nodes,
			chromedp.ByQueryAll,
			)); err != nil {
		return nil
	}

	var links []Link
	for _, link := range nodes {
		links = append(links, Link{
			url: link.Attributes[3],
			title: strings.TrimSpace(link.Children[0].NodeValue),
		})
	}

	return links
}
