  GNU nano 4.8                                                                                                                                              README.md                                                                                                                                                        # keepcookie
Keepcookie is a middleware plugin for [Traefik](https://github.com/traefik/traefik) which Keeps the cookies that are indicated by deleting all others.

### Configuration

### Static

```yaml
experimental:
  plugins:
    keepcookie:
      moduleName: "github.com/ascpikmin/keepcookie"
      version: "v0.1.0"
```

### Dynamic

```yaml
http:
  middlewares:
      keepcookie:
        cookies:
          - cookieName1
          - cookieName2
```
