# Protocus - Custom Protocol

This is a protocol designed from scratch with the porpouse of transferring files between users of **filecaster**

## Technical Design

Protocus is a frame-based binary protocol built on top of TCP. The main advantage of binary protocols is efficiency while transfering the data.

### Protocol Conventions

* Commands are byte-encoded
* It have **headers**, that contains the args of the commands.
* There is no explicit delimiter, the size of a message is specified beforehand

A command will have the following structure: 

The first 4 bits of the command are dedicated to encode the type of command itself as integers, we will specify the available commands later in this document. The following 16 bits encode as integer the size in bytes of the headers. The next 32 bit encode as integer the size of the file beign transferred. Following the file size, we will attach the headers, each command will have its own required headers, they will be a set of bytes encoding a JSON string where each key/value pair is equivalent to a header. Finally, the file bytes are attached to the end, without any specific delimiter. The chain of bytes for a command could be splitted like this:

```
[command][headers_size][file_size][headers_data][file_data]
```

### Commands

Below is the list of supported commands, with its integer representation and human-readable name:

#### Subscribe - 00
Subscribes to the server for a specific channel, listening to transferred files. Required Headers:
```json
{
  "channel": "ch01"
}
```

#### Broadcast - 01
Order the server to broadcast the attached file to all the subscribers in the specified channel. Required Headers:
```json
{
  "channel": "ch01"
}
```
You must attach a file after the headers.

#### Ok - 02
Ack a command that was processed successfully.

#### Error - 03
Failed to process a command. Required Headers:
```json
{
  "message": "error happened for some reason"
}
```
