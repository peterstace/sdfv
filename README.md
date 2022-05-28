# sdfv

SDF Viewer

## Go Package Structure

- `main`: command line parsing, invoke the engine function, manage exit status.

- `engine`: SDF rendering engine.

- `scenes`: Defines the scenes to render.

- `sdf`: defines the signed distance function DSL.

```

main --> engine
 |           |
 |           |
 v           v
scenes ---> sdf

```
