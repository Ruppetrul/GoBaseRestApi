# A simple example of a website with real-time cryptocurrency rates

## Here I have tested the speed of such a site on Go. I compared several approaches.

### Test routes:
`current` - classical list retrieval from the database. (20-25ms)

`test1` - Read the prepared html doc. from the file. (20-25ms)

`test2` - Read prepared html doc. from the database. (10-25ms)

`test3` - From RAM. (10-18ms)