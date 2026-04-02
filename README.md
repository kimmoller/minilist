# minilist

Minimalistic todo list

Run minilist -h to see available commands. All individual commands have a help function that will print out requirements and available aliases.

### Compatibility

So far the application has only been tested on a Linux system but assuming that the ~/.config directory exists or `XDG_CONFIG_HOME` is set to an existing path, everything should work as expected.

## Data

All data for the minilist application will by default be stored in a json file located in `~/.config/minilist/data.json`

If another value has been set in `XDG_CONFIG_HOME` that value will be used instead.

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

### Start

Set an items status to in progress

```
minilist start "${your_description_here}"
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
