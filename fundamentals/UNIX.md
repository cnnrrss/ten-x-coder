# Unix Cheatsheet
http://www.mathcs.emory.edu/~valerie/courses/fall10/155/resources/unix_cheatsheet.html

`man {command` read the manual
`whatis {command}` short desc

### Directories
List a directory
`ls {path}`	It's ok to combine attributes, eg ls -laF gets a long listing of all files with types.
`ls {path_1} {path_2}`	List both {path_1} and {path_2}.
`ls -l {path}`	Long listing, with date, size and permisions.
`ls -a {path}`	Show all files, including important .dot files that don't otherwise show.
`ls -F {path}`	Show type of each file. "/" = directory, "*" = executable.
`ls -R {path}`	Recursive listing, with all subdirs.
`ls {path} | more`	Show listing one screen at a time.

`cd {dir}`
`cd ~`
`cd ..`
`mkdir`
`pwd`

### Copy, Move or rename a file
`cp {file1} {file2}`
`cp -r {dir1} {dir2}` recursive dirs
`cat {newfile} >> {oldfile}` append new to end of old
`mv {old} {new}`

### Delete a file
`rm {filespec}`
`ls {filespec}`
`rm {filespec}`

### View a text file
`more {filename}`
`less {filename}`
`cat {filename}`
`cat {filename} | more`

### Edit a text file
`gedit {filename}` Basic editor
`nano {filename}`
`vi {filename}`
`vim {filename}`
`emacs {filename}`

### Create a textfile
`cat > {filename}`
`gedit {filename}` create some text and save it

### Compare two files
`diff {file1} {file2}`
`sdiff {file1} {file2}`

### Other text commands
`grep '{pattern}' {file}` Find regexp in file
`spell {file}` Display misspelled words
`wc {file}` Count words in file
`wc -l {file}` Count lines in file

### Make an alias
`alias {name}='{command}'` (e.g. alias python=python3)

### Pipes and Redirection

`{command} > {file}` Redirect output to file
`{command} >> {file}` Append output to a file
`{command} < {file}` Get input from a file
`{command} < {file1} > {file2}` Get input from file1 and write to file2
`{command} | {command}` Pipe one command to another


### System info
`date`	    Show date and time.
`df`	        Check system disk capacity.
`du`	        Check your disk usage and show bytes in each directory.
`du -h`	    Check your disk usage in a human readable format
`printenv`	Show all environmental variables
`uptime`	    Find out system load.
`w`	        Who's online and what are they doing?
`top`	        Real time processor and memory usage


### Unix Directory Format


### DOS and UNIX commands
|Action | DOS |	UNIX |
|-------|-----|------|
|change directory |	cd |cd|
|change file protection	|attrib	|chmod|
|compare files|	comp	|diff|
|copy file|	copy	|cp|
|delete file|	del	|rm|
|delete directory|	rd	|rmdir|
|directory list|	dir	|ls|
|edit a file|	edit	|pico|
|environment|	set	|printenv|
|find string in file|	find	|grep|
|help|	help|   man|
|make directory|	md	|mkdir|
|move file|	move	|mv|
|rename file|	ren	|mv|
|show date and time| date, time	|date|
|show disk space|	chkdsk	|df|
|show file|	type|	cat|
|show file by screens	|type filename \| more|	more
|sort data|	sort|	sort|