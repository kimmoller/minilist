# minilist

Minimalistic todo list

Run minilist -h to see available commands. All individual commands have a help function that will print out requirements and available aliases.

### Compatibility

So far the application has only been tested on a Linux system but assuming that the ~/.config directory exists or `XDG_CONFIG_HOME` is set to an existing path, everything should work as expected.

## Data

All data for the minilist application will by default be stored in a json file located in `~/.config/minilist/data.json`

If another value has been set in `XDG_CONFIG_HOME` that value will be used instead.

## Commands

### Migrate

Migrate old v0.1.1 data format into the new data format introduced in v.0.2.0.

This command will be removed in the future as it is only necessary when upgrading from v0.1.x to v.0.2.0. There will be a release note when the command will be removed to notify users that they need to upgrade to a version prior to that to avoid having to migrate data by hand.

Running the command is a no-op if you have already migrated to the new data format.

```
minilist migrate
```

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
