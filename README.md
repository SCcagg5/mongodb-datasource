# Grafana MongoDB Data Source

![ci](https://github.com/haohanyang/mongodb-datasource/actions/workflows/ci.yml/badge.svg?branch=master)

A free, open source plugin to integrate MongoDB on Grafana, an alternative to Grafana Lab's MongoDB enterprise plugin and MongoDB Atlas Charts.

![](docs/assets/screenshot-editor.png)

![](docs/assets/screenshot-editor-modal2.png)

Visit the [documents](https://haohanyang.github.io/mongodb-datasource/) to get started.
## Full connection URI override

This fork supports an optional `connectionUri` field in the datasource JSON settings.
When `connectionUri` is set, the backend passes the URI directly to the MongoDB Go driver and ignores the decomposed Scheme, Host, Connection String Options, Authentication, and TLS/SSL fields.

Example X.509 URI:

```text
mongodb://database:27017/prefect?authMechanism=MONGODB-X509&authSource=%24external&tls=true&tlsCAFile=/certs/authority_ca.crt&tlsCertificateKeyFile=/certs/client.pem&directConnection=true&serverSelectionTimeoutMS=5000
```

The database is inferred from the URI path (`prefect` in the example). If the URI has no database path, the `database` field is used as a fallback.
