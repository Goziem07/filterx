# FilterX

FilterX is a command-line tool written in Go that allows you to filter and print URLs based on specified filters. It reads a list of URLs from standard input, compares them against a set of filter words, and prints out the URLs that contain any matching word.

## Installation

To install FilterX, you can use the following go get command:

`go get github.com/Goziem07/findip`

## Commands

FilterX supports the following commands:

- `-g <filter-name>`: Specify the name of the filter to use. Filter names correspond to JSON files containing a list of filter words.
- `-d <path/to/filter/directory>`: Specify the directory where your json files are.

## Usage

1. Create a directory using the command (I am using urlfilter as my filename):

`mkdir urlfilter`

2. Create a JSON filter file: In the `/<directory>/urlfilter` directory, create a JSON file with a list of filter words. Each JSON file should have the following structure, for example, if you want to create sqli.json:

```
{
  "words": [
    "id=",
    "orderby="
  ]
}
```

3. Prepare a file named <file containing urls>.txt with the list of URLs you want to filter.
  
4. Prepare a file named urls.txt with the list of URLs you want to filter.
  
Run FilterX using the desired command:

To filter URLs using a specific filter, use the -g flag followed by the filter name. For example, to filter URLs using the sqli filter, run:

`cat <file containing urls>.txt | filterx -g sqli -d /<directory>/urlfilter`
  
## Example

Suppose you have a filter named sqli in the /<directory>/urlfilter directory, which contains a list of SQL injection-related words. You want to filter URLs from the urls.txt file using this filter. Here's how you can use FilterX:

Create the JSON filter file <path to directory>/urlfilter/sqli.json with the following content:

```
{
  "words": [
    "id=",
    "select=",
    "report=",
    ...
  ]
}
```
Create the <file containing urls>.txt file with a list of URLs to filter.

Run FilterX with the desired command:

`cat <file containing urls>.txt | filterx -g sqli -d <path to directory>/urlfilter`

FilterX will process the URLs and print out the matching ones based on the SQL injection filter.
