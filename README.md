# piglatin

An HTTP server that translate texts to Pig Latin.

## Setup
Create schema found in ```database/ ```<br>
Set config in ```app.cfg```<br>
Run ```go run . -cfg=app.cfg``` <br>

## API
### Translate
**POST /piglatins**<br>
Request body
``` 
{
  "text": "example input"
}
```
Response
```
{
  "pig_latin": "exampleway inputway"
}
```
### Get all translated texts
**GET /piglatins**<br>
Response
``` 
[
  {
    "text": "example input",
    "pig_latin": "exampleway inputway"
  },
  {
    "text": "one two three",
    "pig_latin": "oneway otway eethray"
  }
]
```

## Test
Run:
```
./test/post.sh example multi word input
./test/get.sh
```
