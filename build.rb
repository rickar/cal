#!/usr/bin/ruby

require 'colorize'

# Run a Command
def run(label, command)
  failures = `#{command}`
  if failures.size > 0
    puts "#{label}: ".cyan << 'Failed'.red
    puts failures
    exit(1)
  end
  puts "#{label}: ".cyan << 'Passed'.green
end

# Format the code
run('GoFormat', 'gofmt -w .')

# Run Linter
run('GoLinter', 'golint ./... 2>&1')

# Run Unit-Tests
run('RunTests', "go test ./... 2>&1 | grep -v ^ok | grep -v '^?'")

# Build
run('RunBuild', 'go build ./...')

# Install
run('Installs', 'go install ./...')
