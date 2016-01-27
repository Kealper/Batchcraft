#Batchcraft

A command-line RCON tool for Minecraft servers.

-----

###Is this something I want?
Maybe! Batchcraft lets you run most console commands through shell scripts or other programs that can't interact with Minecraft's RCON natively. This includes `.bat` files or Command Prompt on Windows and `.sh` files or your preferred shell/terminal on Linux.

###Cool, how do I get it?
Batchcraft is written in Go, and should run on most remotely-modern systems. If you wish to build Batchcraft from source, just have a working Go installation, clone this repository, and run `go build` in the directory to get a binary that works for your system.

If you'd rather not compile it from source and you're running either Windows or 64-bit Linux, just grab the pre-compiled binary for your server's operating system from the `bin` directory in this repository.

###Okay, I have it, how do I install it and use it?
Because Batchcraft relies on Minecraft's RCON feature, that will need to be enabled before Batchcraft will function.

Easy way: Just have the correct Batchcraft binary for your operating system in the same directory as the shell script or program that will be using it.

Advanced way: Place the correct Batchcraft binary for your operating system in a directory that is available in your system's PATH environment variable. This will enable all your scripts and programs to use the same binary instead of needing a separate copy in any directory that contains something that uses it.

Once you have it "installed" with either method listed above, just use it like this:

`batchcraft -a hostname:port -p password -c "command string"`

Where "hostname:port" is your Minecraft server's RCON address and port (such as `127.0.0.1:25575`), "password" is your configured RCON password, and "command string" is the console command you wish to run. If your command contains spaces, you should put quotes around the command, like this: `"whitelist add SomePlayer"`

###License Information
Batchcraft is licensed under the Apache License version 2.0.

This means you're free to use this code in your own projects, modify and redistribute changes to this code, and even sell products that make use of this code under the conditions that: The unmodified parts of this code are still licensed under this same license and that you don't claim the unmodified parts as your own.

You can find out more about what you can and can not do with this code in the full text of the license here:
https://www.apache.org/licenses/LICENSE-2.0
