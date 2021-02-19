# restvirt-client

A client for the restvirt service.

## Configuration

When this client is used as a part of another application, the configuration from environment variable and config files might come in handy.
By default, `NewClientFromEnvironment()` will look for a configuration file in `~/.restvirt/config.yaml` that has the following structure:
```yaml
default:
  host: https://myhost.com:8091
  username: username
  password: password
  ca: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K
```
By default, it will use the `default` profile. If you want the client to use a different profile, you can set the `RESTVIRT_PROFILE` environment variable to the desired profile name.
