### OnTheGo Chat Server

OnTheGo is a chat server that enables connected clients to broadcast messages to other connected clients.
Clients cannot view older messages that were sent before they joined the chat room.

Currently, it supports only one chat room.

<br/>

#### What it does
- On start up, the server listens for new client connections
- After establishing a connection to the server, the client must provide a user name. This user name is used while broadcasting the user's messages to other connected users and logging the message
- The server listens to a common channel. The user messages are passed to the server via this channel
- A new "read" channel is created for every new user that connects to the server. The user passes messages to the server via the common channel
- On recieving a new message on the common channel, the server publishes this message to every other user's "read" channel. The server also, passes this message to the MessageLogger which logs this message to a file
- Information like server host, port and log folder location are loaded from the configuration (config/default.json)

<br/>

#### Running the server
- Clone the repository to your go workspace or add the workspace with your local repository to your GOPATH
``` bash
$ export GOPATH=$GOPATH:<workspace-path>
```
- Change the server configuration to specify the port, host, [OPTIONAL]
``` bash
$ vi <project-path>/config/default.json
```
Configuration is as follows:
``` json
{
    "LogFilePath": "<path to the folder where the chat messages are logged>",
    "ServerHost": "localhost|0.0.0.0|127.0.0.1<If left empty, the server listens on IPs provided by all the interfaces on the machine>",
    "ServerPort": "8554"
}
```
- Build the server
``` bash
$ go build
```
- Run the server
``` bash
$ ./chat_server
```

<br/>

#### Connecting to the server
- Connect to the server using telnet
``` bash
$ telnet <server-ip> <server-port>
```


<br/>

#### Tech Debt / Improvements
- Message Log file: Message log file should be rotated based on size/date. In the current approach, the message file keeps growing. The message file name is currently not configurable keeping in mind the file rotation implementation
- Configuration: There should be an option to pass the configuration file location should be passed to the server on start up.
- User Name: Currently, after connecting to the server, the user is allowed to join the chat room even if the user name entered by the user is already in use by another active user in the chat room. This should be prevented lest it cause confusion
- Application Logging: Currently, all application logging is only println statements directed to the stdout. This output should be directed to a log file using a logger and a log format should be followed.
- User Messages: User's own messages appear without any information such as timestamp on the user console currently. The user's own message should also have the same format as other user's messages (example: replacing \<user-name\> with "me")
- Scaling: Since this chat server implementation is solely based on go-routines, it cannot be scaled horizontally.

<br/>

#### Development Roadmap
* Allow users to create and join different chat rooms
    * Users should be able to list all available chat rooms using a command like  `/list`
    * User should be able to join a chat room using a command like  `/join <chat-room-name>`
    * User should be able to create a chat room using a command like `/create <chat-room-name>`
* A user should be able to ignore messages from another user by a command like `/mute <user-name>` and reverse the action by a command like `/unmute <user-name>`
* A user should be able to upload images, videos and documents to the chat room (for when you really must share your cat videos :) )
* An API to post messages to a specific channel or the default broadcast channel `POST /api/chat-room/{chat-room-name}`
* An API to get all of a chat-room's messages `GET /api/chat-room/{chat-room-name}`
