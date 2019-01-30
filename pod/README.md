## Pod

Pod is a language for Automata job, which borrowed a lot from [HCL](https://github.com/hashicorp/hcl) and [hil](https://github.com/hashicorp/hil).

## Example

Pod job will be like:

```
name = "jenkins_error"

triggers {
  type = "webhook/timer/emit"
  options {
    method = "POST"
    type = "application/json"
  }
}

steps [
  {
    layer = "kv"
    statements [
      q = get ${trigger.name}
    ]
    export [
      q
    ]
  }
  {
    condition = ${steps.1.q} != nil && ${trigger.build.status} == "SUCCESS"
    layer = "trello"
    statements [
      delete card ${steps.1.q.card_id}
    ]
  }
  {
    condition = ${steps.1.q} == nil && ${trigger.build.status} != "SUCCESS"
    layer = "trello"
    statements [
      card = create card ${trigger.name} xxxx
    ]
    export [
      card
    ]
  }
  {
    condition = ${steps.3.card} != nil
    layer = "kv"
    statements [
      pus ${trigger.name} card
    ]
  }
]

hooks [
  on_error {
    layer = "trello"
      statements [
      add card ${steps.1.q.card_id}
    ]
  }
]

layers [
  trello {
    source = "git://github.com/NieR/nier-trello@v1.0.0"
    options {
      xxx: ${secrets.trello_api_key}
      yyy: yyy
    }
  }
  kv {
    options {
      type: persistent/temporary
    }
  }
]

variables {
  trello_board_id: "xxx"
  long_string: <<EOF
xxxxxx
xxxxx
xxxx
EOF
}

secrets {
  trello_api_key: xxx
}
```