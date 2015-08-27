# Remotgo

[![Join the chat at https://gitter.im/msempere/remotgo](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/msempere/remotgo?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

Send commands over ssh to AWS EC2 instances

### Example
Execute "df -H" command in all instances with highway as role.

![Usage example](http://i.imgur.com/69fyPJx.gif)

### Usage

Use one or more -t (-tag) for selecting the tags from the instances that will receive the command specified with -c (-command).
Use -u (-username) and -p (-password) if you need to provide it for the ssh connection.

Examples:

    $ remotgo -t role:core -t environment:test -t group:stack -c "uname -a"
    $ remotgo -t name:webserver01 -t environment:production -t group:eu-servers -c "mount | column -t"

### Configuration

The following environment variables have to be properly set:
 - AWS_ACCESS_KEY_ID
 - AWS_SECRET_ACCESS_KEY
 - AWS_REGION
 
To use environment variables, do the following:

    $ export AWS_ACCESS_KEY_ID=<access_key>
    $ export AWS_SECRET_ACCESS_KEY=<secret_key>
    $ export AWS_REGION=<region>





### License
Distributed under MIT license. See `LICENSE` for more information.

