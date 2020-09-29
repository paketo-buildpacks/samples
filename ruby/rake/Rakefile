task :main => :main_dependent do
  puts "I am the main rake task"
end

task :default => [:main]

task :main_dependent do
  puts "I am the rake task that has to run before main"
end
