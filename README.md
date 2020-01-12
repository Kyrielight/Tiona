# Tiona

Tiona is a simple webserver script that "fixes" MangaDex RSS feeds for IFTTT. Currently, the \<mangaLink\> field isn't recognized by IFTTT, which causes applets to fail. 
Tiona simply strips out all of the \<mangaLink\> fields and returns the rest of the RSS (XML) untouched.

Tiona matches all existing MangaDex.org or MangaDex.cc RSS urls - simply replace the "mangadex.cc" part with your domain.

You can sample it under `https://tiona.best.moe/`.

### Example:

MangaDex "Latest Updates" RSS feed is "https://mangadex.cc/rss/R3umhbeYdUaKy6cfSPVCFQXwtDMEgq7v".
Replace "https://mangadex.cc" with your own domain. 
In the example, this would become:
"https://tiona.best.moe/rss/R3umhbeYdUaKy6cfSPVCFQXwtDMEgq7v"


#### Misc.

Docker build and run scripts are included. They will build this package and run the container under the name "Tiona". Defaults to port 43156.