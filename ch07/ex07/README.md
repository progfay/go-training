# Why default value of `-temp` option in help message has `°C` suffix

- Set returned value of `flag.Value.String()` to `flag.FlagSet.DefValue` on `flag.CommandLine.Var(flag.Value, string, string)`
- `--help` option call `FlagSet.PrintDefaults()` and print `flag.FlagSet.DefValue` as default value
