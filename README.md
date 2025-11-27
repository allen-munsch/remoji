# Remoji: remove emojis from text


```
go build remoji.go
./test.sh | ./remoji

# example output
 File → stdout
 Pipe input → stdout
 In-place edit
 No emoji untouched
 Mixed content

Summary:
   Passed: 5
   Failed: 0

# ultra mega test

cat vibes.txt | ./remoji
./remoji -i vibes.txt

```
