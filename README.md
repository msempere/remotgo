# Remotgo

[![Join the chat at https://gitter.im/msempere/remotgo](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/msempere/remotgo?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

Send commands over ssh to AWS EC2 instances

### Example
Execute "df -H" command in all instances with highway as role.

![Usage example](http://i.imgur.com/69fyPJx.gif)

### Usage

Use one or more -t (-tag) for selecting the tags from the instances that will receive the command specified with -c (-command).
Use -u (-username) and -p (-password) if you need to provide it for the ssh connection.
    
    COMMANDS:
        help, h	Shows a list of commands or help for one command
       
    GLOBAL OPTIONS:
       --username, -u "user"                    Ssh username (default: current user)
       --password, -p 				            Ssh password (default: empty)
       --command, -c "uname -a"			        Command to execute.
       --quiet, -q					            Quiet mode (default: false)
       --timeout, -o "200"				        Shh command timeout (default: 200)
       --tags, -t [--tags option --tags option]	EC2 instance tags
       --rsa, -r "/home/user/.ssh/id_rsa"		Path to RSA file (default ~/.ssh/id_rsa)
       --dsa, -d "/home/user/.ssh/id_dsa"		Path to DSA file (default ~/.ssh/id_dsa)
       --help, -h					            show help
       --version, -v				            print the version


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

