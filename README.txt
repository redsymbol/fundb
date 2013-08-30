simple SQL database server

This is a relational database server I am making for fun - hence, the
name. It is:

 * uses a file-based data store (like sqlite3)
 * designed for highly concurrent reads and writes
 * implemented in Go

You almost certainly want to use a different DB server for your
application - postgres, mysql, sqlite3, etc.  The only reason to look
at this is if you are learning Go, or like the language, and would
enjoy learning from my code (and/or learning from its flaws).

LICENSE

Public domain.

AUTHOR

Aaron Maxwell - amax.at.redsymbol.dot.net

