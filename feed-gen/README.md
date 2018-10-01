# feed-gen

This API is intended to create a terrible feed experience. For the purpose of exploring various tools and techniques for normalising data.

The general outline is;

| field | description | quirks |
| - | - |
| fixture | This field lists who is playing | It is string field delimited by a random delimiter. It may also appear in an array of length 1 |
| startDate | start date of the event | none (yet) |
| sport | describes the sport of the feed | This field may randomly be encapsulated in an array of length 1 |
| markets | betting markets object | none |
| markets.id | id of market | Can be int or string |
| markets.type | type of market | No real quirks |
| markets.options | options of market | none |
| markets.options.id | if od options | can be int or string |
| markets.options.name | name of competitor (or "draw") | none |
| markets.options.odds | odds for option | can be in one of 4 odds formats |
| id | id of message | may be called `id` of `uuid`, may be a string or int, may be encapsulated in an array of length 1 |

## Usage

```sh
DEBUG=feed-gen:* npm start
```
