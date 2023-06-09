# FilterX

FilterX is a command-line tool written in Go that allows you to filter and print URLs based on specified filters. It reads a list of URLs from standard input, compares them against a set of filter words, and prints out the URLs that contain any matching word.

## Installation

## Installation 1

To install FilterX, you can use the following go get command:

`go install -v github.com/Goziem07/filterx@latest`

## Installation 2

1. Clone Filterx-patterns

`git clone https://github.com/Goziem07/Filterx-patterns.git`

2. Clone Filterx

`git clone https://github.com/Goziem07/filterx.git`

3. Move to the directory

`cd filterx`

4. Build it

`go build -o filterx`

5. Move the binary to /usr/local/bin

`sudo mv filterx /usr/local/bin`

## Usage

1. Clone the Filterx-patterns repository:

`git clone https://github.com/Goziem07/Filterx-patterns.git`

3. Prepare a file containing with the list of URLs you want to filter.
  
Run FilterX using the desired command:

To filter URLs using a specific filter, use the -g flag followed by the filter name. For example, to filter URLs using the sqli filter, run:

`cat <file containing urls>.txt | filterx sqli`
  
## Example

Suppose you want to filter sqli related words, here's how you can use FilterX:

`cat <file containing urls>.txt | filterx sqli`

FilterX will process the URLs and print out the matching ones based on the SQL injection filter.
