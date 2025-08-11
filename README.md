# chorely
Daily chores app -- learning golang + htmx

## Setup / run
uses air to detect changes

```bash
# get htmx static
curl -o static/htmx.min.js https://unpkg.com/htmx.org@2.0.6/dist/htmx.min.js

go install github.com/air-verse/air@v1.62.0
# to set up new .air.toml file
~/go/bin/air init

# to run
~/go/bin/air
```

