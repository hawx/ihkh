# ihkh

A minimalist Flickr viewer. Single-user port of [jstn/ihardlyknowher][].

```
$ go get hawx.me/code/ihkh
$ ihkh --user-id "75695140@N04" --api-key $FLICKR_API_KEY
...
```

## pages

- `/` public photostream
- `/sets` list of sets
- `/tags` list of tags

## keyboard

- <kbd>J</kbd> next item, at last item double press for next page
- <kbd>K</kbd> previous item, at first item double press for previous page

[jstn/ihardlyknowher]: https://github.com/jstn/ihardlyknowher
