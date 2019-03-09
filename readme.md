# Stack Stalk

## Install

```bash
cd stalk
go install
```

## *Usage*

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
