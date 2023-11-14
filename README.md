> **note**
> The main purpose of this project is for me to learn Go, feedback is much appreciated!

# Gezondheid

Gezondheid `/ɣəˈzɔntˌɦɛi̯t/` (Dutch for "Health") is a simple CLI tool to periodically check the health of URLs.

👷‍♂️ This project (and readme) is under construction

## Plugins

Behaviour can be extended with 3rd party plugins like [gezondheid-hook](https://github.com/LiamEderzeel/gezondheid-hook) to add webhook support when health checks fail.

```yaml
- name: test.test
  url: https://test.test
  interval: 10s
  plugins:
    - name: "gezondheid-hook.so"
      config:
        method: "POST"
        url: "https://webhook.test"
        statusCodeMinimum: 200
```




