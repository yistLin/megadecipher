# MegaDownloader Link Decipher

This project is inspired by a gist [dinos80152/mega-links-decrypter.html](https://gist.github.com/dinos80152/fa09c00b7befbfce07a7471b042665dc) a lot.

### Install & Usage

To get the package

```bash
go get github.com/yistLin/megadecipher
```

To install the command-line tool

```bash
go install github.com/yistLin/megadecipher/cmd/mega-decipher
```

To decipher a link (make sure you have `$GOPATH` configured and have `$GOPATH/bin` in your `$PATH`)

```bash
mega-decipher "mega://enc2?a_ciphered_url_here_is_just_an_example"
```
