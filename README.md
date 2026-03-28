# minilist

Minimalistic todo list

Run minilist -h to see available commands. All individual commands have a help function that will print out requirements and available aliases.

## Data

All data for the minilist application will be stored in a json file located in `~/.config/minilist/data.json`

## Commands

### List

List all existing in progress items

```
minilist list
```

To list all items including completed items, run the commands with the --all flag.

```
minilist list --all
```

### Add

Add a new item

```
minilist add "${your_description_here}"
```

### Complete

Mark an item as complete

```
minilist complete ${item_id}
```

### Delete

Delete an existing item

```
minilist delete ${item_id}
```
