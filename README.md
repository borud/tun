# tun - multi-hop reverse proxy

`tun` is a multi-hope reverse ssh proxy. When it starts it connects to a remote host via a chain of SSH servers.  It then creates a listen port on the last server in the chain.  Whenever a connection is made to this listen port, the connection is proxied back to the originating host and connected to a *target*.

## Example usage

```shell
tun --key ~/.ssh/tunkey                 \
    --via user@jump1.example.com:22     \
    --via user2@jump2.example.com:22    \
    --target localhost:22               \
    --remote-listen-addr localhost:2222
```

1. This first connects to `jump1.example.com`.
2. Then `jump1.example.com` connects to `jump2.example.com`
3. since `jump2.example.com` is last in the chain it opens a listen port on `localhost:2222`
4. when connections are made to `localhost:2222` on `jump2.example.com` the connection
   is proxied back to the starting host and connected to `localhost:22`

The `--via` command line options are applied in order and you can have as many of them as you like.  You must have at least one.
