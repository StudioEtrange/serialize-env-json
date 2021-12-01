# serialize-env-json

Parse environment variables and tweak and serialize them into JSON

* Can select var by a regexp filter (all by default)
* Can transform variable name to upper or lower
* Can remove any matching group of the regexp filter from env var name

# Usage

```
serialize-env-json [--filter <regexp>] [--clean] [--upper] [--lower]
```

return a json string with this json struct

```
fullname : unmodified env var name
name : resulting env var name
value : env var value
```

# Samples

```
serialize-env-json --filter "^P(A)TH" --clean --lower


{"envs":[{"fullname":"PATH","name":"pth","value":"/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"}]}
```


# Fork

Forked and modified from https://github.com/StudioEtrange/parse-env
