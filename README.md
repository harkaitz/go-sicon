SICON (Systray Icon Command Line Utility)
=========================================

Create a systray icon for your shell script. You can find an
example [here](./bin/sicon-example).

You can download a MS Windows prebuilt binary [here]().

## Go programs

    Usage: sicon CMDS... < CMDS > IDS
    
    Create a system tray icon as defined by the commands received from
    the standard input and print the selected menu item IDs to the
    standard output.
    
    Options:
    
      -n : Do not read commands from standard input.
    
    Commands:
    
      title=TITLE                : Set the icon title.
      tooltip=TOOLTIP            : Set the icon tooltip.
      icon=(b64:BASE64_ICO|FILE) : Set the icon image.
    
      menu_add=ID|LABEL|TOOLTIP          : Add a menu item.
      menu_add_quit=LABEL|TOOLTIP        : Add a quit menu item (quit ID).
      menu_add_separator                 : Add a menu separator.
      menu_hide=ID                       : Hide a menu item.
      menu_show=ID                       : Show a menu item.
      menu_icon=ID|(b64:BASE64_ICO|FILE) : Add icon to menu item.
    
    Copyright (c) 2024 - Harkaitz Agirre - All rights reserved.

## Collaborating

For making bug reports, feature requests and donations visit
one of the following links:

1. [gemini://harkadev.com/oss/](gemini://harkadev.com/oss/)
2. [https://harkadev.com/oss/](https://harkadev.com/oss/)
