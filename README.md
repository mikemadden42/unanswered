# unanswered

Show unanswered posts in [discuss.elastic.co](https://discuss.elastic.co).

A topic counts as unanswered when it has only the original post and is not
closed, pinned, or marked solved.

## Install

    go install github.com/mikemadden42/unanswered@latest

## Usage

    unanswered [-category elastic-stack/beats/28] [-pages 10]

- `-category`: the category path as it appears in `/c/<category>/l/latest`
  (default `elastic-stack/beats/28`)
- `-pages`: maximum number of pages to fetch (default `10`); fetching stops
  early when there are no more topics

## Development

    make check     # golangci-lint
    go test ./...
    make build     # goreleaser snapshot build
