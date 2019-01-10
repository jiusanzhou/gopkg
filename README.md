# gopkg


Gopkg which can be used as a middleware or standalone app allows you to create vanity Go import urls.

## Usage

### Install

**From relese:**
1. Download lasted from [Relese page](https://github.com/jiusanzhou/gopkg/releases) with fitting your platform.
2. Unzip the archive.

**From source code:**
```bash
go get go.zoe.im/gopkg
```

### Start

First, start server with command `./gopkg [-http=localhost:PORT -index=YOUR_SITE HOST/USER]`.
For example:
```
./gopkg -index=https://zoe.im -http=:8080 github.com/jiusanzhou
```
Then, use `caddy` or `Nginx` proxy `YOUR_DOMAIN` to `localhost:PORT`.
**:warning: Must follow upstream headers.**

### Enjoy yourself

```
go get go.zoe.im/gopkg
```

## TODO

- [ ] Caddy pluginable
- [ ] More test
- [ ] Pipe with CI
- [ ] Badges in README
