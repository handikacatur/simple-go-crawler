# simple-go-crawler
This is a simple web crawler to collect all links inside a website. This crawler support Multi Page Application as well as Single Page Application.

---
### Basic Usage
To use this app, you just need to run the server and then post your url to `/api/crawl`. For Single Page Applications you might need to consider html element to be visible before you start scrapping. <br>
Change `chromedp.WaitVisible('your-element')` to match your requirement. It will wait until those particular element appear before start scrapping.

### Dependencies
1. chromedp [see here](https://github.com/chromedp/chromedp)
2. GoFiber [see here](https://github.com/gofiber/fiber)