# Ruby debugging

## Show the stacktrace when at a breakpoint

```rb
binding.pry
pry(#<Application::Account>)> caller
```

Shows:

```
/Users/aurelio/.rbenv/versions/2.5.9/lib/ruby/gems/2.5.0/gems/pry-0.13.1/lib/pry/pry_instance.rb:290:in `eval'",
 "/Users/aurelio/.rbenv/versions/2.5.9/lib/ruby/gems/2.5.0/gems/pry-0.13.1/lib/pry/pry_instance.rb:290:in `evaluate_ruby'",
 "/Users/aurelio/.rbenv/versions/2.5.9/lib/ruby/gems/2.5.0/gems/pry-0.13.1/lib/pry/pry_instance.rb:659:in `handle_line'",
 "/Users/aurelio/.rbenv/versions/2.5.9/lib/ruby/gems/2.5.0/gems/pry-0.13.1/lib/pry/pry_instance.rb:261:in `block (2 levels) in eval'",
```

You can filter the trace by filename or path:

```rb
caller.select {|x| x["/app/jobs/update_newsletter"] }
```

See [here](https://stackoverflow.com/a/21620257/1446845)
