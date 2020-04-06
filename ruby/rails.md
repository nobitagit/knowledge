# Rails

An empty project to follow along can be found [here](https://github.com/nobitagit/docker-rails-postgres)

## Controllers

```sh
rails generate controller [name] [action, action, action]
# for instance
rails generate controller users index
```

This will generate a new controller in `controller/home_controller.rb`:

```rb
class HomeController < ApplicationController
  def index
  end
end
```

We can change this to be our homepage. In routes, change this:

```rb
# from
get 'home/index'
# to
root 'home#index'
```

We can add a new route manually.

```rb
get 'about' => 'home#about'
```

For that to work, our home controller will need to adapt:

```rb
class HomeController < ApplicationController
  def index
  end

  def about
  end
end
```

## Models

```sh
rails generate model [model_name] [field_name:field_type, field_name:field_type,]
rails generate model question email:string body:text
# if we also want to add restful routes to handle the model a better way is to
rails generate resource question {same as before}
```

This will generate a bunch of stuff, namely the controller and the migration for the new model.
To run the migration.

```rb
rake db:migrate
```

To fill the db:

```sh
rails console
Question.create email: 'andy@gray.com', body: 'some text'
# #<Question id: 1, email: "andy@gray.com", body: "some text", created_at: "2020-04-05 00:08:46", updated_at: "2020-04-05 00:08:46">
Question.first
```
