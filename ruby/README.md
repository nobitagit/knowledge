# Ruby

To update only a single gem:

```sh
bundle update --conservative somegem
```

## Intro

- Interpreted
- Dynamic types
- Multi paradigm (functional and procedural)

## Some basics

```rb
v = "nope"
v.class
=> String
v.class.superclass
=> Object

3.class
=> Integer

def double(e); e * 2; end
double(2)
=> 4
# and even (wtf)
double([2])
=> [2, 2]
# and even...
double(_)
=> [2, 2, 2, 2]
# surprise though
double({})
# NoMethodError (undefined method `*' for {}:Hash)

# Methods can end with a question mark.
# This is convention for when the method returns a boolean
class Radio
   is_on?
     state.active == true
   end
end

# Methods can end with !
# This is convention for when the mthod does something dangerous
# (like mutating) or slightly unexpected
a = " hi  "
a.trim
=> "hi"
puts a
=> " hi  "
a.trim! # mutates
puts a
"hi"

# parens are optional for method invocation
double 2
double(2)
```

## Classes

- start with capital letter
- user camelcase

Use `@` to assign instance variables.

```rb
class Guitar
  def play(note)
    @note = note
    note
  end
end

strato = Guitar.new
strato.play('re')
strato.inspect
# <Guitar:0x00007fb3eb827840 @note="re">
p strato # same as inspect
# <Guitar:0x00007fb3eb827840 @note="re">
```

instance variables are **private** by default.

```rb
g.note # error
```

Instance methods are **public** by default.

To access an instance variable we have to define a method that returns it.

```rb
class Guitar
  #... as above
  def note
    @note
  end
end
```

While this works it can become tedious quickly. The alternative is to use **attribute accessors**.

```rb
class Guitar
  attr_accessor :note # READ AND WRITE
  attr_reader :note # READ ONLY
  attr_writer :note # WRITE ONLY
end
```

This makes it messy when we are defining local variables. Look at this:

```rb
class Train
  attr_accessor :speed, :make, :color, :destination

  def get_state
    speed = 22 # if instance.speed is 35, and we set speed = 22 here...
    [speed, destination] # what is the returned speed here?
  end
end
```

To circumvent it, Ruby wants you to use `self` to use the accessor

```rb
class Train
  attr_accessor :speed, :make, :color, :destination

  def get_state
    speed = 22
    [self.speed, destination] # returned speed is 35
  end
end
```

Initial args are defined within the initialize method:

```rb
class Tube
  def initialize(line, destination)
    @line = line
    @destination = destination
  end
end

Tube.new("Victoria", "Seven Sisters")
```

## Inheritance

A class can only inherit from 1 class at a time. Multiple inheritance is not supported per se.
If not specified, a class inherits from Object.
