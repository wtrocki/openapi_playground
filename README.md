# openapi_playground

OpenAPI playground used to reproduce OpenAPI Generator issues and improvements.

## Usage

1. Create new openapi file in .openapi folder
2. Add new generator step to ./scripts/generate.sh
3. Run `make generate`
4. Run mock `npx @stoplight/prism-cli mock .openapi/your_file.json|
5. Create and run sdk example 

## Runniung tests

1.Run mock

```bash
make mock
```

2.In separate window run example

```bash
make example
```

## Generating sdks

```bash
make generate
```
