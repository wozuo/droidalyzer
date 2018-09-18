# droidalyzer

- Multithreaded scanning of Android projects for networking libraries
- Scanning of Android projects for networking code

## Usage

```
droidalyzer [flags] [path]
```

### Flags

The -sp flag indicates if a single project should be scanned
The -lib flag indicates if projects should be scanned
for networking libraries
The -printLib flag indicates if information about libraries
should be printed for every scanned project

### Examples

```
droidalyzer -lib "~/documents/Android Projects"
droidalyzer -lib -printLib "~/documents/Android Projects"
droidalyzer -sp "~/documents/Android Project"
```

## License

MIT