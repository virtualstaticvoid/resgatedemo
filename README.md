# Resgate Example

See [Resgate][resgate] for information.

## Local Development Host

This example uses [`ldhdns`][ldhdns] for DNS resolution for the various services.

```
docker run \
  --name ldhdns \
  --detach \
  --network host \
  --security-opt "apparmor=unconfined" \
  --volume "/var/run/docker.sock:/tmp/docker.sock" \
  --volume "/var/run/dbus/system_bus_socket:/var/run/dbus/system_bus_socket" \
  --restart unless-stopped \
  virtualstaticvoid/ldhdns:latest
```

## Usage

### Build

```
docker-compose build
```

### Run

```
docker-compose up
```

### Test

```
curl -v http://api.resgate.ldh.dns:8080/api/example/model
```

Now visit http://www.resgate.ldh.dns in your browser.

## License

MIT License. Copyright (c) 2021 Chris Stefano. See [LICENSE](LICENSE) for details.

<!-- links -->

[ldhdns]: https://github.com/virtualstaticvoid/ldhdns
[resgate]: https://resgate.io/
