# Stack Stalk
[![Build Status](https://travis-ci.org/skailasa/stackstalk.svg?branch=master)](https://travis-ci.org/skailasa/stackstalk)

## Install

```bash
cd stalk
go install
```

## Usage

```bash
# Print help
stalk

# Query interface relies on 'verbs' and 'adjectives'
stalk <verb> <adjective> --stack, -s <stack>

# Example usage
stalk query top "Can I eat pi?" -s math

# defaults to searching entire site
stalk query new "How tall is obi-wan kenobi?"
```

## API

#### Verbs:

| verb | action  | 
|---|---|
| `query`  |query the api, searching for a string |

#### Adjectives

|  adjective| action  | 
|---|---|
| `top`  | most relevant match  |
|  `new` | newest match |  
|  `hot` | most active match  | 
|  `pop` | most upvoted match |