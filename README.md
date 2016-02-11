Gyazo-Server
==

Gyazo server for home use.

Getting Started
--

### Setup AWS account

Write AWS account in your shell configuration;

```sh
export AWS_ACCESS_KEY_ID=YOUR_AKID
export AWS_SECRET_ACCESS_KEY=YOUR_SECRET_KEY
```

Then run server daemon;

```sh
$ gyazo_server -b YOUR_BUCKET -r YOUR_REGION
```

Bucket name and region name option are required.

Options
--

```
Application Options:
  -p, --port=       Port number for listening (9090)
  -b, --bucket=     Bucket name for AWS
  -r, --region=     Region name for AWS
  -h, --host=       Host name (http://localhost:9090)
  -P, --passphrase= Passphrase for upload

Help Options:
  -h, --help        Show this help message
```

Tips
--

### Authentication for upload

This gyazo server has a mechanism of simple authentication for image uploading.

#### Server

If you want to enable authentication, please run server with `-P` (also `--passphrase`) option.

#### Client

Add `X-Gyazo-Auth` request header. If header value equals passphrase of server, it allows to upload image. Otherwise, it will be denied.

License
--

MIT

