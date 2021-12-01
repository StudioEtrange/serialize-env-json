# serialize-env-json

Parse environment variables and tweak and serialize them into JSON

* Can select var by a regexp filter (all by default)
* Can transform variable name to upper or lower
* Can remove any matching group of the regexp filter from env var name

# Usage

* Command
    ```
    serialize-env-json --help

    serialize-env-json [--filter <regexp>] [--clean] [--upper] [--lower]
    ```

* Options

    ```
    --filter <regexp>
        select env var with a regexp matching env var name (default ".*")
    ```

* Tweak env var name options

    ```
    --clean 
            Remove from env var name all the matching regexp group
    --lower
            Force env var name to lower
    --upper
            Force env var name to upper
    ```

* JSON result file format

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

# Build

## using makefile

```
# build for current platform
make

# build for specific platforms
make PLATFORM=linux/amd64
make PLATFORM=darwin/amd64
make PLATFORM=windows/amd64

# platform list : https://go.dev/doc/install/source#environment

# check code with a linter
make lint
```

## without makefile

```
# build for current platform
DOCKER_BUILDKIT=1 docker build . --target bin --output bin/

# build for specific platforms
DOCKER_BUILDKIT=1 docker build . --target bin --output bin/ --platform linux/amd64
```

# Notes

* Project organisation with docker guide : https://www.docker.com/blog/containerize-your-go-developer-environment-part-1/
* Inspired by https://github.com/joshhsoj1902/parse-env
