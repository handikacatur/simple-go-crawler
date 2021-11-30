# simple-go-crawler
This is a simple web crawler to collect all links inside a website. This crawler support Multi Page Application as well as Single Page Application.

---
### Basic Usage
To use this app, you just need to run the crawler function with url parameter. For Single Page Applications you might need to consider html element to be visible before you start scrapping. <br>
Change `chromedp.WaitVisible('your-element')` to match your requirement. It will wait until those particular element appear before start scrapping. 