# Trie

A trie aimed at being memory efficient and fast. It currently stores 2.5 million items, having an average length of 16 characters, using 6.7MB. A query for a key takes 23μs.

## Status

Currently only Insert and Find work. I plan to:

- Add unit tests around insert and find
- Refactor and clean up the code
- Improve performance / memory (some ideas I want to try, not sure if any will work)
- Handle more common cases (two different entries with the same value)
- Work on delete/update
