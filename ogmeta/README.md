# ogmeta

Scrapes and displays OpenGraph & metadata properties for a website.

## Installation

```
go get github.com/jakewarren/junkdrawer/ogmeta
``` 

## Usage

```
ogmeta <url>
```

## Example

```
❯ ogmeta https://twitter.com/VessOnSecurity/status/1175079589019815936?s=20
robots = NOODP
msapplication-TileImage = //abs.twimg.com/favicons/win8-tile-144.png
msapplication-TileColor = #00aced
swift-page-name = permalink
swift-page-section = permalink
al:ios:url = twitter://status?id=1175079589019815936
al:ios:app_store_id = 333903271
al:ios:app_name = Twitter
al:android:url = twitter://status?status_id=1175079589019815936
al:android:package = com.twitter.android
al:android:app_name = Twitter
og:type = article
og:url = https://twitter.com/VessOnSecurity/status/1175079589019815936
og:title = Vess on Twitter
og:image = https://pbs.twimg.com/profile_images/684776303329935360/7HnClXW4_400x400.jpg
og:description = “OK, so the Crown Sterling/Time AI scammers have been at it again. I'm a bit late to the party and others have debunked their latest stunt, but I just have to comment on this, so here it goes.”
og:site_name = Twitter
fb:app_id = 2231777543
```