`elixir file.exs` to run things
`mix format file.exs` to format


`|> tap(fn result -> require IEx; IEx.pry; end)`   Add to end of statements
Run with `iex file.exs`

- `&Kernel.tap/2`
- `&IO.inspect/2`